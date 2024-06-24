package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/cnnrznn/authenticator/model"
)

const (
	FN = ".authenticator"
)

func Save(tokens []model.Token) error {
	path, err := genPath()
	if err != nil {
		return err
	}

	f, err := open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	bs, err := json.Marshal(tokens)
	if err != nil {
		return err
	}

	n, err := f.Write(bs)
	if err != nil {
		return err
	}
	if n != len(bs) {
		return fmt.Errorf("failed to write all bytes")
	}

	if err := f.Truncate(int64(n)); err != nil {
		return err
	}

	return nil
}

func Load() ([]model.Token, error) {
	path, err := genPath()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tokens := []model.Token{}

	if err := json.NewDecoder(f).Decode(&tokens); err != nil {
		return nil, err
	}

	return tokens, nil
}

func open(fn string) (*os.File, error) {
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func genPath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(dir, FN), nil
}
