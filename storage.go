package main

import (
	"encoding/json"
	"os"
)
type Storage[L any] struct {
	FileName string
}

func NewStorage[L any](fileName string) *Storage[L] {
	return &Storage[L]{FileName: fileName}
}

func (s *Storage[L]) Save(data L) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	
	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[L]) Load(data *L) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data)
}
