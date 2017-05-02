package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	P12FilePath     string `toml:"p12"`
	P12Password     string `toml:"password"`
	DeviceToken     string `toml:"devicetoken"`
	Topic           string `toml:"topic"`
	PayloadFilePath string `toml:"payload"`
	Env             string `toml:"env"`
}

func (cfg *Config) Initialize() error {
	dir := filepath.Join(os.Getenv("HOME"), ".config", "pusher")

	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("cannot create directoty: %v", err)
	}

	file := filepath.Join(dir, "config.toml")

	_, err := os.Stat(file)
	if err == nil {
		_, err = toml.DecodeFile(file, cfg)
		if err != nil {
			return err
		}

		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	cfg.Env = "production"
	cfg.P12FilePath = "/path/to/file.p12"
	cfg.PayloadFilePath = "/path/to/payload.json"

	return toml.NewEncoder(f).Encode(cfg)
}
