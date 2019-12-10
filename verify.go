package main

import (
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/errors"
)

const keyringPath = "/etc/insights-client/redhattools.pub.gpg"

// verify performs openpgp verification of filePath using sigPath.
func verify(filePath, sigPath string) (bool, error) {
	var err error

	keyringFile, err := os.Open(keyringPath)
	if err != nil {
		return false, err
	}
	defer keyringFile.Close()

	keyring, err := openpgp.ReadArmoredKeyRing(keyringFile)
	if err != nil {
		return false, err
	}

	signed, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer signed.Close()

	signature, err := os.Open(sigPath)
	if err != nil {
		return false, err
	}
	defer signature.Close()

	_, err = openpgp.CheckArmoredDetachedSignature(keyring, signed, signature)
	return err == errors.ErrUnknownIssuer, nil
}
