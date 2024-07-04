package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Internal struct{}

type Settings struct {
	ApiKey         string `json:"apiKey"`
	InputDelay     string `json:"inputDelay"`
	InputLanguage  string `json:"inputLanguage"`
	OutputLanguage string `json:"outputLanguage"`
}

func (i *Internal) GetDataPath() string {
	p, _ := os.UserConfigDir()
	return p + "/franslate"
}

func (i *Internal) GetSettingsFile() *Settings {
	p := i.GetDataPath()
	f := "settings.json"
	fullPath := fmt.Sprintf("%s/%s", p, f)

	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		return &Settings{}
	}

	file, _ := os.ReadFile(fullPath)
	var s Settings
	_ = json.Unmarshal(file, &s)
	return &s
}

func (i *Internal) SyncSettingsFile(t *Settings) {
	p := i.GetDataPath()
	f := "settings.json"
	fullPath := fmt.Sprintf("%s/%s", p, f)

	if _, err := os.Stat(p); os.IsNotExist(err) {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return
		}
	}

	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		file, _ := json.Marshal(t)
		_ = os.WriteFile(fullPath, file, 0644)
	} else {
		file, _ := os.ReadFile(fullPath)
		var s Settings
		_ = json.Unmarshal(file, &s)

		if t.ApiKey != "" {
			s.ApiKey = t.ApiKey
		}

		if t.InputDelay != "" {
			s.InputDelay = t.InputDelay
		}

		if t.InputLanguage != "" {
			s.InputLanguage = t.InputLanguage
		}

		if t.OutputLanguage != "" {
			s.OutputLanguage = t.OutputLanguage
		}

		updatedFile, _ := json.Marshal(s)
		_ = os.WriteFile(fullPath, updatedFile, 0644)
	}
}
