// Package engine provides access to the game engine.
package engine

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ardanlabs/ethereum"
	"github.com/ethereum/go-ethereum/common"
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
func (e *Engine) Connect(keyStorePath string, address common.Address, passPhrase string) (Token, error) {
	keyFile, err := findKeyFile(keyStorePath, address)
	if err != nil {
		return Token{}, fmt.Errorf("find key file: %w", err)
	}

	privateKey, err := ethereum.PrivateKeyByKeyFile(keyFile, passPhrase)
	if err != nil {
		return Token{}, fmt.Errorf("get private key: %w", err)
	}

	dt := struct {
		Address  common.Address `json:"address"`
		ChainID  int            `json:"chainId"`  // 1337
		DateTime string         `json:"dateTime"` // YYYYMMDDHHMMSS
	}{
		Address:  address,
		ChainID:  1337,
		DateTime: time.Now().UTC().Format("20060102150405"),
	}

	sig, err := ethereum.SignAny(dt, privateKey)
	if err != nil {
		return Token{}, fmt.Errorf("sign: %w", err)
	}

	dts := struct {
		Address   common.Address `json:"address"`
		ChainID   int            `json:"chainId"`
		DateTime  string         `json:"dateTime"`
		Signature string         `json:"sig"`
	}{
		Address:   dt.Address,
		ChainID:   dt.ChainID,
		DateTime:  dt.DateTime,
		Signature: sig,
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
func (e *Engine) Events(f func(event string, address common.Address)) (func(), error) {
	url := strings.Replace(e.url, "http", "ws", 1)
	url = fmt.Sprintf("%s/v1/game/events", url)

	req := make(http.Header)
	req.Add("authorization", fmt.Sprintf("Bearer %s", e.token))

	socket, _, err := websocket.DefaultDialer.Dial(url, req)
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
				f("error: "+err.Error()[:30], common.Address{})
				return
			}

			var event struct {
				Type    string         `json:"type"`
				Address common.Address `json:"address"`
			}
			if err := json.Unmarshal(message, &event); err != nil {
				f("error: "+err.Error()[:30], common.Address{})
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

// QueryState returns the current state of the specified game.
func (e *Engine) QueryState(gameID string) (State, error) {
	url := fmt.Sprintf("%s/v1/game/%s/state", e.url, gameID)

	var state State
	if err := e.do(url, &state, nil); err != nil {
		return State{}, err
	}

	return state, nil
}

// Tables returns the current set of tables.
func (e *Engine) Tables(gameID string) (Tables, error) {
	url := fmt.Sprintf("%s/v1/game/tables", e.url)

	var tables Tables
	if err := e.do(url, &tables, nil); err != nil {
		return Tables{}, err
	}

	return tables, nil
}

// NewGame starts a new game on the game engine.
func (e *Engine) NewGame() (State, error) {
	url := fmt.Sprintf("%s/v1/game/new", e.url)

	var state State
	if err := e.do(url, &state, nil); err != nil {
		return State{}, err
	}

	return state, nil
}

// StartGame generates the five dice for the player.
func (e *Engine) StartGame(gameID string) (State, error) {
	url := fmt.Sprintf("%s/v1/game/%s/start", e.url, gameID)

	var state State
	if err := e.do(url, &state, nil); err != nil {
		return State{}, err
	}

	return state, nil
}

// RollDice generates the five dice for the player.
func (e *Engine) RollDice(gameID string) (State, error) {
	url := fmt.Sprintf("%s/v1/game/%s/rolldice", e.url, gameID)

	var state State
	if err := e.do(url, &state, nil); err != nil {
		return State{}, err
	}

	return state, nil
}

// JoinGame adds a player to the current game.
func (e *Engine) JoinGame(gameID string) (State, error) {
	url := fmt.Sprintf("%s/v1/game/%s/join", e.url, gameID)

	var state State
	if err := e.do(url, &state, nil); err != nil {
		return State{}, err
	}

	return state, nil
}

// Bet submits a bet to the game engine.
func (e *Engine) Bet(gameID string, number int, suit rune) (State, error) {
	url := fmt.Sprintf("%s/v1/game/%s/bet/%d/%c", e.url, gameID, number, suit)

	var state State
	if err := e.do(url, &state, nil); err != nil {
		return State{}, err
	}

	return state, nil
}

// Liar submits a liar call to the game engine.
func (e *Engine) Liar(gameID string) (State, error) {
	url := fmt.Sprintf("%s/v1/game/%s/liar", e.url, gameID)

	var state State
	if err := e.do(url, &state, nil); err != nil {
		return State{}, err
	}

	return state, nil
}

// Reconcile submits a reconcile call when the game is over.
func (e *Engine) Reconcile(gameID string) (State, error) {
	url := fmt.Sprintf("%s/v1/game/%s/reconcile", e.url, gameID)

	var state State
	if err := e.do(url, &state, nil); err != nil {
		return State{}, err
	}

	return state, nil
}

// do makes the actual http call to the engine.
func (e *Engine) do(url string, result interface{}, input []byte) error {
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

	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", e.token))

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
		return fmt.Errorf("%s", er.Error)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	return nil
}

// findKeyFile searches the keystore for the specified address key file.
func findKeyFile(keyStorePath string, address common.Address) (string, error) {
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

		if strings.Contains(strings.ToLower(fileName), strings.ToLower(address.Hex()[2:])) {
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
