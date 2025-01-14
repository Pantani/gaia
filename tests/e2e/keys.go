package e2e

import (
	"github.com/cosmos/go-bip39"
)

// createMnemonic creates a random string mnemonic
func createMnemonic() (string, error) { //nolint:unused // this is called during e2e tests
	entropySeed, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}
