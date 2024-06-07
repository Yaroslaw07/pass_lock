package main

import (
	"bytes"
	"io"
	"os"

	"filippo.io/age"
	"filippo.io/age/armor"
)

const secFile = "exp.age"

func writeEnc(txt string, pass string) error {
	recipient, err := age.NewScryptRecipient(pass)

	if err != nil {
		return err
	}

	out, err := os.OpenFile(secFile, os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return err
	}

	defer out.Close()

	armorWrite := armor.NewWriter(out)

	defer armorWrite.Close()
	w, err := age.Encrypt(armorWrite, recipient)

	if err != nil {
		return err
	}

	defer w.Close()

	if _, err := w.Write([]byte(txt)); err != nil {
		return err
	}

	return nil
}

func readEnc(pass string) (string, error) {
	indentify, err := age.NewScryptIdentity(pass)

	if err != nil {
		return "", err
	}

	out := &bytes.Buffer{}

	in, err := os.Open(secFile)

	if err != nil {
		return "", err
	}

	defer in.Close()

	armorRead := armor.NewReader(in)
	r, err := age.Decrypt(armorRead, indentify)

	if err != nil {
		return "", err
	}

	if _, err := io.Copy(out, r); err != nil {
		return "", err
	}

	return out.String(), nil
}
