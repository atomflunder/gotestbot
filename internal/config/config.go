package config

import (
	"encoding/json"
	"os"
)

//gets the token
type TokenConfig struct {
	Token string `json:"token"`
}

func GetToken(fileName string) (c *TokenConfig, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}

	c = new(TokenConfig)
	err = json.NewDecoder(f).Decode(c)

	return
}

//gets the prefix
type PrefixConfig struct {
	Prefix string `json:"prefix"`
}

func GetPrefix(fileName string) (c *PrefixConfig, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}

	c = new(PrefixConfig)
	err = json.NewDecoder(f).Decode(c)

	return
}
