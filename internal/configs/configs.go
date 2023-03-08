package configs

import (
	"encoding/json"
	"io/ioutil"

	"github.com/stori/internal/schema"
)

// PopulateConfigs get all configurations from localconfig folder
func PopulateConfigs() (*schema.PopulatedConfigs, error) {
	configFile, err := ioutil.ReadFile("./localconfig/config.json")
	if err != nil {
		return nil, err
	}
	secretFile, err := ioutil.ReadFile("./localconfig/secrets.json")
	if err != nil {
		return nil, err
	}
	var config schema.Config
	var secret schema.Secret

	err = json.Unmarshal([]byte(configFile), &config)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(secretFile), &secret)
	if err != nil {
		return nil, err
	}

	return &schema.PopulatedConfigs{Config: config, Secret: secret}, nil
}
