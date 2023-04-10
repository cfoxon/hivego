package hivego_test

import (
	"bytes"
	"testing"

	"github.com/cfoxon/hivego"
	"github.com/decred/dcrd/dcrec/secp256k1/v2"
)

func TestKeyPairFromWif(t *testing.T) {
	keyPair, err := hivego.KeyPairFromWif("5JUvJcF6rQvFbZLtDFagreKCYWWcHpHApy7sbRHZ6PeZYNftLh6")

	if err != nil {
		t.Error("Failed to decode valid private key")
	}

	var expectedPubKey = []byte{3, 106, 48, 22, 243, 45, 96, 255, 51, 197, 8, 179, 85, 147, 131, 32, 165, 214, 76, 64, 90, 168, 63, 67, 124, 7, 139, 26, 114, 145, 144, 94, 153}
	var actualPubKey = keyPair.PublicKey.Serialize()

	if !bytes.Equal(expectedPubKey, actualPubKey) {
		t.Errorf("Public Key %v does not match expected key %v", actualPubKey, expectedPubKey)
	}

	var expectedPrivKey = []byte{87, 184, 167, 38, 175, 19, 89, 57, 56, 199, 44, 31, 237, 202, 159, 200, 87, 40, 158, 247, 154, 118, 181, 226, 213, 12, 41, 131, 159, 122, 80, 230}
	var actualPrivKey = keyPair.PrivateKey.Serialize()

	if !bytes.Equal(expectedPrivKey, actualPrivKey) {
		t.Errorf("Private Key %v does not match expected key %v", actualPrivKey, expectedPrivKey)
	}
}

func TestDecodePublicKey(t *testing.T) {
	pubKey, err := hivego.DecodePublicKey("STM7dzxQo2aaav9weydSVAwqewcUz2GbUwyWrAVqkdiKsD6V1uX8B")

	if err != nil {
		t.Errorf("Error Decoding valid public key")
	}

	var expectedPubKey = []byte{3, 106, 48, 22, 243, 45, 96, 255, 51, 197, 8, 179, 85, 147, 131, 32, 165, 214, 76, 64, 90, 168, 63, 67, 124, 7, 139, 26, 114, 145, 144, 94, 153}
	var actualPubKey = pubKey.Serialize()

	if !bytes.Equal(expectedPubKey, actualPubKey) {
		t.Errorf("Public Key %v does not match expected key %v", actualPubKey, expectedPubKey)
	}
}

func TestGetPublicKeyString(t *testing.T) {
	var pubKeyData = []byte{3, 106, 48, 22, 243, 45, 96, 255, 51, 197, 8, 179, 85, 147, 131, 32, 165, 214, 76, 64, 90, 168, 63, 67, 124, 7, 139, 26, 114, 145, 144, 94, 153}
	const pubKeyStringExpected = "STM7dzxQo2aaav9weydSVAwqewcUz2GbUwyWrAVqkdiKsD6V1uX8B"

	pubKey, err := secp256k1.ParsePubKey(pubKeyData)

	if err != nil {
		t.Error("Couldn't parse public key to run test")
	}

	pubKeyString := hivego.GetPublicKeyString(pubKey)

	if *pubKeyString != pubKeyStringExpected {
		t.Errorf("Public Key string %s does not match expected string %s", *pubKeyString, pubKeyStringExpected)
	}
}
