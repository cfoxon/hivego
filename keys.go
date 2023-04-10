package hivego

import (
	"bytes"
	"errors"

	"github.com/decred/base58"
	"github.com/decred/dcrd/dcrec/secp256k1/v2"

	//lint:ignore SA1019 ripemd160 is used for checksums of public keys and is required for compatibility with Hive
	"golang.org/x/crypto/ripemd160"
)

var PublicKeyPrefix = "STM"

type KeyPair struct {
	PrivateKey *secp256k1.PrivateKey
	PublicKey  *secp256k1.PublicKey
}

// Gets a KeyPair from a given WIF String
func KeyPairFromWif(wif string) (*KeyPair, error) {
	privKey, _, err := GphBase58CheckDecode(wif)

	if err != nil {
		return nil, err
	}

	prvKey, pubKey := secp256k1.PrivKeyFromBytes(privKey)

	return &KeyPair{prvKey, pubKey}, nil
}

// Decodes a base58 Hive public key to secp256k1 public key
func DecodePublicKey(pubKey string) (*secp256k1.PublicKey, error) {
	// check prefix matches
	if pubKey[:len(PublicKeyPrefix)] != PublicKeyPrefix {
		return nil, errors.New("invalid prefix")
	}

	// remove prefix
	pubKey = pubKey[len(PublicKeyPrefix):]

	// decode base58
	decoded := base58.Decode(pubKey)

	// get checksum
	checksum := decoded[len(decoded)-4:]

	// get public key
	pubKeyBytes := decoded[:len(decoded)-4]

	// get ripemd160 hash (checksum)
	hasher := ripemd160.New()
	_, err := hasher.Write(pubKeyBytes)
	newChecksum := hasher.Sum(nil)[:4]

	if err != nil {
		return nil, err
	}

	// check if checksums match
	if !bytes.Equal(checksum, newChecksum) {
		return nil, errors.New("checksums do not match")
	}

	parsedKey, err := secp256k1.ParsePubKey(pubKeyBytes)

	if err != nil {
		return nil, err
	}

	return parsedKey, nil
}

func (kp *KeyPair) GetPublicKeyString() *string {
	return GetPublicKeyString(kp.PublicKey)
}

func GetPublicKeyString(pubKey *secp256k1.PublicKey) *string {
	if pubKey == nil {
		return nil
	}

	pubKeyBytes := pubKey.SerializeCompressed()

	// get ripemd160 hash
	hasher := ripemd160.New()
	_, err := hasher.Write(pubKeyBytes)

	if err != nil {
		return nil
	}

	// get checksum
	checksum := hasher.Sum(nil)[:4]

	// append checksum to public key

	pubKeyBytes = append(pubKeyBytes, checksum...)

	// encode to base58
	encoded := base58.Encode(pubKeyBytes)

	// add prefix
	encoded = PublicKeyPrefix + encoded

	return &encoded
}
