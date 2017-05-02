package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sideshow/apns2/payload"
)

func InitializePayload() error {
	dir := filepath.Join(os.Getenv("HOME"), ".config", "pusher")

	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("cannot create directoty: %v", err)
	}
	file := filepath.Join(dir, "payload.json")
	payload := payload.NewPayload()

	_, err := os.Stat(file)
	if err == nil {
		bs, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bs, payload)
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
	payload.Alert("Alert")
	payload.Badge(10)
	payload.Sound("sound-file")
	payload.Custom("custom-key", "custom-value")

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(payload)
}
