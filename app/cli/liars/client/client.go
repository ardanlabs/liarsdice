// Package client provides access to the game engine.
package client

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
	"github.com/gorilla/websocket"
)

// Client provides access to the game engine API.
type Client struct {
	url   string
	token string
}

// New constructs a client that provides access to the game engine.
func New(url string) *Client {
	url = strings.TrimSuffix(url, "/")
	return &Client{
		url: url,
	}
}

// Connect authenticates the use to the game engine.
func (c *Client) Connect(keyStorePath string, address string, passPhrase string) (Token, error) {
	keyFile, err := findKeyFile(keyStorePath, address)
	if err != nil {
		return Token{}, fmt.Errorf("find key file: %w", err)
	}

	privateKey, err := contract.PrivateKeyByKeyFile(keyFile, passPhrase)
	if err != nil {
		return Token{}, fmt.Errorf("get private key: %w", err)
	}

	dt := struct {
		DateTime string `json:"date_time"` // YYYYMMDDHHMMSS
	}{
		DateTime: time.Now().Format("20060102150405"),
	}

	sig, err := contract.Sign(dt, privateKey)
	if err != nil {
		return Token{}, fmt.Errorf("sign: %w", err)
	}

	dts := struct {
		DateTime  string `json:"date_time"`
		Signature string `json:"sig"`
	}{
		DateTime:  dt.DateTime,
		Signature: fmt.Sprintf("0x%s", hex.EncodeToString(sig)),
	}

	data, err := json.Marshal(dts)
	if err != nil {
		return Token{}, fmt.Errorf("marshal: %w", err)
	}

	url := fmt.Sprintf("%s/v1/game/connect", c.url)

	var token Token
	if err := c.do(url, &token, data); err != nil {
		return Token{}, err
	}

	c.token = token.Token

	return token, nil
}

// Events establishes a web socket connection to the game engine.
func (c *Client) Events(f func(event string, address string)) (func(), error) {
	url := strings.Replace(c.url, "http", "ws", 1)
	url = fmt.Sprintf("%s/v1/game/events", url)

	socket, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, fmt.Errorf("dial: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	teardown := func() {
		socket.Close()
		wg.Wait()
	}

	go func() {
		defer wg.Done()
		for {
			_, message, err := socket.ReadMessage()
			if err != nil {
				f("error", err.Error()[:30])
				return
			}

			var event struct {
				Type    string `json:"type"`
				Address string `json:"address"`
			}
			if err := json.Unmarshal(message, &event); err != nil {
				f("error", err.Error()[:30])
				continue
			}
			f(event.Type, event.Address)
		}
	}()

	return teardown, nil
}

// Configuration returns the configuration of the game engine.
func (c *Client) Configuration() (Config, error) {
	url := fmt.Sprintf("%s/v1/game/config", c.url)

	var config Config
	if err := c.do(url, &config, nil); err != nil {
		return Config{}, err
	}

	return config, nil
}

// Status starts a new game on the game engine.
func (c *Client) Status() (Status, error) {
	url := fmt.Sprintf("%s/v1/game/status", c.url)

	var status Status
	if err := c.do(url, &status, nil); err != nil {
		return Status{}, err
	}

	return status, nil
}

// NewGame starts a new game on the game engine.
func (c *Client) NewGame() (Status, error) {
	url := fmt.Sprintf("%s/v1/game/new", c.url)

	var status Status
	if err := c.do(url, &status, nil); err != nil {
		return Status{}, err
	}

	return status, nil
}

// Balance returns the accounts balance.
func (c *Client) Balance() (string, error) {
	url := fmt.Sprintf("%s/v1/game/balance", c.url)

	var account struct {
		Balance string `json:"balance"`
	}
	if err := c.do(url, &account, nil); err != nil {
		return "", err
	}

	return account.Balance, nil
}

// StartGame generates the five dice for the player.
func (c *Client) StartGame() (Status, error) {
	url := fmt.Sprintf("%s/v1/game/start", c.url)

	var status Status
	if err := c.do(url, &status, nil); err != nil {
		return Status{}, err
	}

	return status, nil
}

// RollDice generates the five dice for the player.
func (c *Client) RollDice() ([]int, error) {
	url := fmt.Sprintf("%s/v1/game/rolldice", c.url)

	var dice struct {
		Dice []int `json:"dice"`
	}

	if err := c.do(url, &dice, nil); err != nil {
		return nil, err
	}

	return dice.Dice, nil
}

// =============================================================================

func (c Client) do(url string, result interface{}, input []byte) error {
	var req *http.Request
	var err error

	if input == nil {
		req, err = http.NewRequest(http.MethodGet, url, nil)
	} else {
		reader := bytes.NewReader(input)
		req, err = http.NewRequest(http.MethodPost, url, reader)
	}

	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}

	req.Header.Add("authorization", fmt.Sprintf("bearer %s", c.token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("client do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var er ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&er); err != nil {
			return fmt.Errorf("status: %s, decode error: %w", resp.Status, err)
		}
		return fmt.Errorf("status: %s, error: %s", resp.Status, er.Error)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	return nil
}

// findKeyFile searches the keystore for the specified address key file.
func findKeyFile(keyStorePath string, address string) (string, error) {
	keyStorePath = strings.TrimSuffix(keyStorePath, "/")
	errFound := errors.New("found")

	var filePath string
	fn := func(fileName string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walkdir failure: %w", err)
		}

		if dirEntry.IsDir() {
			return nil
		}

		if strings.Contains(fileName, fileName) {
			filePath = fmt.Sprintf("%s/%s", keyStorePath, fileName)
			return errFound
		}

		return nil
	}

	if err := fs.WalkDir(os.DirFS(keyStorePath), ".", fn); err != nil {
		if errors.Is(err, errFound) {
			return filePath, nil
		}
		return "", fmt.Errorf("walking directory: %w", err)
	}

	return "", errors.New("not found")
}
