// Package signature provides helper functions for handling the blockchain
// signature needs.
package signature

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

// ZeroHash represents a hash code of zeros.
const ZeroHash string = "0x0000000000000000000000000000000000000000000000000000000000000000"

// ardanID is an arbitrary number for signing messages. This will make it
// clear that the signature comes from the Ardan blockchain.
// Ethereum and Bitcoin do this as well, but they use the value of 27.
const EthID = 27

// =============================================================================

// VerifySignature verifies the signature conforms to our standards.
func VerifySignature(value any, v, r, s *big.Int) error {

	// Check the recovery id is either 0 or 1.
	uintV := v.Uint64() - EthID
	if uintV != 0 && uintV != 1 {
		return errors.New("invalid recovery id")
	}

	// Check the signature values are valid.
	if !crypto.ValidateSignatureValues(byte(uintV), r, s, false) {
		return errors.New("invalid signature values")
	}

	return nil
}

// FromAddress extracts the address for the account that signed the data.
func FromAddress(value any, signature string) (string, error) {

	// NOTE: If the same exact data for the given signature is not provided
	// we will get the wrong from address for this transaction. There is no
	// way to check this on the node since we don't have a copy of the public
	// key used. The public key is being extracted from the data and signature.

	// Prepare the data for public key extraction.
	data, err := stamp(value)
	if err != nil {
		return "", err
	}

	sig, err := hex.DecodeString(signature[2:])
	if err != nil {
		return "", err
	}

	sig[64] = sig[64] - EthID

	// Capture the public key associated with this data and signature.
	publicKey, err := crypto.SigToPub(data, sig)
	if err != nil {
		return "", err
	}

	// Extract the account address from the public key.
	return crypto.PubkeyToAddress(*publicKey).String(), nil
}

// =============================================================================

// stamp returns a hash of 32 bytes that represents this data with
// the Ardan stamp embedded into the final hash.
func stamp(value any) ([]byte, error) {

	// Marshal the data.
	v, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	// Hash the data data into a 32 byte array. This will provide
	// a data length consistency with all data.
	txHash := crypto.Keccak256(v)

	// Convert the stamp into a slice of bytes. This stamp is
	// used so signatures we produce when signing data
	// are always unique to the Ardan blockchain.
	stamp := []byte("\x19Ethereum Signed Message:\n32")

	// Hash the stamp and txHash together in a final 32 byte array
	// that represents the data.
	data := crypto.Keccak256(stamp, txHash)

	return data, nil
}
