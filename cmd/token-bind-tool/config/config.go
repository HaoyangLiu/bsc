package config

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	ContractData  string `json:"contract_data"`
	Symbol        string `json:"symbol"`
	BEP2Symbol    string `json:"bep2_symbol"`
	LedgerAccount string `json:"ledger_account"`
}

func (bindConfig *Config) validate() error {
	_, err := hex.DecodeString(bindConfig.ContractData)
	if err != nil {
		return fmt.Errorf("invalid contract byte code: %s", err.Error())
	}
	if len(bindConfig.BEP2Symbol) == 0 {
		return fmt.Errorf("missing bep2 token symbol")
	}
	if !strings.HasPrefix(bindConfig.LedgerAccount, "0x") || len(bindConfig.LedgerAccount) != 42 {
		return fmt.Errorf("invalid ledger account, expect bsc address, like 0x4E656459ed25bF986Eea1196Bc1B00665401645d")
	}
	return nil
}

func ReadConfigData(configPath string) (Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open config file: %s",err.Error())
	}
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return Config{}, err
	}
	err = config.validate()
	if err != nil {
		return Config{}, err
	}
	return config, nil
}


