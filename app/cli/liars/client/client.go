// Package client provides access to the game engine.
package client

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
)

// Client provides access to the game engine API.
type Client struct {
	url string
}

// New constructs a client that provides access to the game engine.
func New(url string) *Client {
	url = strings.TrimSuffix(url, "/")
	return &Client{
		url: url,
	}
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

// Connect authenticates the use to the game engine.
func (c *Client) Connect(keyFile string, passPhrase string) (Token, error) {
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

	return token, nil
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

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("client do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	return nil
}
