// Package engine provides access to the game engine.
package engine

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

// Engine provides access to the game engine API.
type Engine struct {
	url   string
	token string
}

// New constructs a client that provides access to the game engine.
func New(url string) *Engine {
	url = strings.TrimSuffix(url, "/")
	return &Engine{
		url: url,
	}
}

// URL returns the url of the game engine.
func (e *Engine) URL() string {
	return e.url
}

// Connect authenticates the use to the game engine.
func (e *Engine) Connect(keyStorePath string, address string, passPhrase string) (Token, error) {
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

	url := fmt.Sprintf("%s/v1/game/connect", e.url)

	var token Token
	if err := e.do(url, &token, data); err != nil {
		return Token{}, err
	}

	e.token = token.Token

	return token, nil
}

// Events establishes a web socket connection to the game engine.
func (e *Engine) Events(f func(event string, address string)) (func(), error) {
	url := strings.Replace(e.url, "http", "ws", 1)
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
func (e *Engine) Configuration() (Config, error) {
	url := fmt.Sprintf("%s/v1/game/config", e.url)

	var config Config
	if err := e.do(url, &config, nil); err != nil {
		return Config{}, err
	}

	return config, nil
}

// QueryStatus starts a new game on the game engine.
func (e *Engine) QueryStatus() (Status, error) {
	url := fmt.Sprintf("%s/v1/game/status", e.url)

	var status Status
	if err := e.do(url, &status, nil); err != nil {
		return Status{}, err
	}

	return status, nil
}

// NewGame starts a new game on the game engine.
func (e *Engine) NewGame() (Status, error) {
	url := fmt.Sprintf("%s/v1/game/new", e.url)

	var status Status
	if err := e.do(url, &status, nil); err != nil {
		return Status{}, err
	}

	return status, nil
}

// Balance returns the accounts balance.
func (e *Engine) Balance() (string, error) {
	url := fmt.Sprintf("%s/v1/game/balance", e.url)

	var account struct {
		Balance string `json:"balance"`
	}
	if err := e.do(url, &account, nil); err != nil {
		return "", err
	}

	return account.Balance, nil
}

// StartGame generates the five dice for the player.
func (e *Engine) StartGame() (Status, error) {
	url := fmt.Sprintf("%s/v1/game/start", e.url)

	var status Status
	if err := e.do(url, &status, nil); err != nil {
		return Status{}, err
	}

	return status, nil
}

// RollDice generates the five dice for the player.
func (e *Engine) RollDice() ([]int, error) {
	url := fmt.Sprintf("%s/v1/game/rolldice", e.url)

	var dice struct {
		Dice []int `json:"dice"`
	}

	if err := e.do(url, &dice, nil); err != nil {
		return nil, err
	}

	return dice.Dice, nil
}

// JoinGame adds a player to the current game.
func (e *Engine) JoinGame() (Status, error) {
	url := fmt.Sprintf("%s/v1/game/join", e.url)

	var status Status
	if err := e.do(url, &status, nil); err != nil {
		return Status{}, err
	}

	return status, nil
}

// Bet submits a bet to the game engine.
func (e *Engine) Bet(number int, suite rune) (Status, error) {
	url := fmt.Sprintf("%s/v1/game/bet/%d/%c", e.url, number, suite)

	var status Status
	if err := e.do(url, &status, nil); err != nil {
		return Status{}, err
	}

	return status, nil
}

// Liar submits a liar call to the game engine.
func (e *Engine) Liar() (Status, error) {
	url := fmt.Sprintf("%s/v1/game/liar", e.url)

	var status Status
	if err := e.do(url, &status, nil); err != nil {
		return Status{}, err
	}

	return status, nil
}

// Reconcile submits a reconcile call when the game is over.
func (e *Engine) Reconcile() (Status, error) {
	url := fmt.Sprintf("%s/v1/game/reconcile", e.url)

	var status Status
	if err := e.do(url, &status, nil); err != nil {
		return Status{}, err
	}

	return status, nil
}

// =============================================================================

// do makes the actual http call to the engine.
func (e Engine) do(url string, result interface{}, input []byte) error {
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

	req.Header.Add("authorization", fmt.Sprintf("bearer %s", e.token))

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

		if strings.Contains(fileName, address[2:]) {
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
