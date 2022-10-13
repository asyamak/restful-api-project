package config

import (
	"encoding/json"
	"os"
	"time"
)

const (
	// var add string
	filename   = "config/config.json"
	maxHeader  = 1 >> 20
	writeTO    = 10 * time.Second
	shutDownTO = 3 * time.Second
)

type Config struct {
	Port            string `json:"port"`
	MaxHeaderBytes  int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeOut time.Duration
	Host            string `json:"host"`
	Dbname          string `json:"dbname"`
	Ssl             string `json:"ssl"`
	Password        string `json:"password"`
	User            string `json:"user"`
}

func NewConfig() (*Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, err
	}
	defer configFile.Close()
	return &Config{
		Port:            config.Port,
		MaxHeaderBytes:  maxHeader,
		ReadTimeout:     writeTO,
		WriteTimeout:    writeTO,
		ShutdownTimeOut: shutDownTO,
		Host:            config.Host,
		Dbname:          config.Dbname,
		Ssl:             config.Ssl,
		Password:        config.Password,
		User:            config.User,
	}, nil
}
