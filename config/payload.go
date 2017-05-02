package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sideshow/apns2/payload"
)

func (c *Config) InitializePayload() error {
	dir := filepath.Dir(c.PayloadFilePath)

	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("cannot create directoty: %v", err)
	}

	payload := payload.NewPayload()

	_, err := os.Stat(c.PayloadFilePath)
	if err == nil {
		bs, err := ioutil.ReadFile(c.PayloadFilePath)
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

	f, err := os.Create(c.PayloadFilePath)
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
