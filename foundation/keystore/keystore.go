// Package keystore implements the auth.KeyLookup interface. This implements
// an in-memory keystore for JWT support.
package keystore

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	keyTypeRSA   = "rsa"
	keyTypeECDSA = "ecdsa"
)

type key struct {
	keyType    string
	privatePEM string
	publicPEM  string
}

// KeyStore represents an in memory store implementation of the
// KeyLookup interface for use with the auth package.
type KeyStore struct {
	log   *logger.Logger
	store map[string]key
}

// New constructs an empty KeyStore ready for use.
func New(log *logger.Logger) *KeyStore {
	return &KeyStore{
		log:   log,
		store: make(map[string]key),
	}
}

// LoadAuthKeys loads a set of RSA PEM files rooted inside of a directory. The
// name of each PEM file will be used as the key id.
func (ks *KeyStore) LoadAuthKeys(folder string) error {
	fsys := os.DirFS(folder)

	fn := func(fileName string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walkdir failure: %w", err)
		}

		if dirEntry.IsDir() {
			return nil
		}

		if path.Ext(fileName) != ".pem" {
			return nil
		}

		file, err := fsys.Open(fileName)
		if err != nil {
			return fmt.Errorf("opening key file: %w", err)
		}
		defer file.Close()

		// limit PEM file size to 1 megabyte. This should be reasonable for
		// almost any PEM file and prevents shenanigans like linking the file
		// to /dev/random or something like that.
		pem, err := io.ReadAll(io.LimitReader(file, 1024*1024))
		if err != nil {
			return fmt.Errorf("reading auth private key: %w", err)
		}

		privatePEM := string(pem)
		publicPEM, err := toPublicPEM(privatePEM)
		if err != nil {
			return fmt.Errorf("reading auth private key: %w", err)
		}

		key := key{
			keyType:    keyTypeRSA,
			privatePEM: privatePEM,
			publicPEM:  publicPEM,
		}

		kid := strings.TrimSuffix(dirEntry.Name(), ".pem")
		ks.store[kid] = key

		ks.log.Info(context.Background(), "Loading Auth Keys", "KID", kid)

		return nil
	}

	if err := fs.WalkDir(fsys, ".", fn); err != nil {
		return fmt.Errorf("walking directory: %w", err)
	}

	return nil
}

// LoadBankKeys loads a set of password protected ECDSA key files rooted inside
// of a directory. The last section of the name for each file will be used as
// the key id.
func (ks *KeyStore) LoadBankKeys(folder string, passPhrase string) error {
	fsys := os.DirFS(folder)

	fn := func(fileName string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walkdir failure: %w", err)
		}

		if dirEntry.IsDir() {
			return nil
		}

		fileName = fmt.Sprintf("%s/%s", folder, fileName)

		pk, err := ethereum.PrivateKeyByKeyFile(fileName, passPhrase)
		if err != nil {
			return fmt.Errorf("capture private key: %s", err)
		}

		kid := strings.Split(dirEntry.Name(), "Z--")
		if len(kid) != 2 {
			return fmt.Errorf("misformed file name: %s", dirEntry.Name())
		}

		pem := pem.EncodeToMemory(&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: crypto.FromECDSA(pk),
		})

		key := key{
			keyType:    keyTypeECDSA,
			privatePEM: string(pem),
		}

		ks.store[kid[1]] = key

		ks.log.Info(context.Background(), "Loading Bank Keys", "KID", kid[1])

		return nil
	}

	if err := fs.WalkDir(fsys, ".", fn); err != nil {
		return fmt.Errorf("walking directory: %w", err)
	}

	return nil
}

// PrivateKey searches the key store for a given kid and returns the private key.
func (ks *KeyStore) PrivateKey(kid string) (string, error) {
	privateKey, found := ks.store[kid]
	if !found {
		return "", errors.New("kid lookup failed")
	}

	return string(privateKey.privatePEM), nil
}

// PublicKey searches the key store for a given kid and returns the public key.
func (ks *KeyStore) PublicKey(kid string) (string, error) {
	key, found := ks.store[kid]
	if !found {
		return "", errors.New("kid lookup failed")
	}

	switch key.keyType {
	case keyTypeRSA:
		return key.publicPEM, nil

	case keyTypeECDSA:
		return "", errors.New("Unsupported")
	}

	return "", errors.New("Unsupported")
}

// toPublicPEM was taken from the JWT package to reduce the dependency. It
// accepts a PEM encoding of a RSA private key and converts to a PEM encoded
// public key.
func toPublicPEM(privateKeyPEM string) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	}

	var parsedKey any
	parsedKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
	}

	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("key is not a valid RSA private key")
	}

	asn1Bytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", fmt.Errorf("marshaling public key: %w", err)
	}

	publicBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	var buf bytes.Buffer
	if err := pem.Encode(&buf, &publicBlock); err != nil {
		return "", fmt.Errorf("encoding to public PEM: %w", err)
	}

	return buf.String(), nil
}
