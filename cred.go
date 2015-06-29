package cred

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ChimeraCoder/anaconda"
	"github.com/mitchellh/go-homedir"
)

var confPath string

// Config type has token information.
type Config struct {
	AccessToken  string
	AccessSecret string
}

func parseConf() (Config, error) {
	confData, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Println("config file not found")
		return Config{}, err
	}
	var c Config
	err = json.Unmarshal(confData, &c)
	if err != nil {
		log.Println("config file can't parsed")
		return Config{}, err
	}
	return c, nil
}

func makeCredential() error {
	url, cred, err := anaconda.AuthorizationURL("")
	if err != nil {
		log.Println("something wrong happen")
		return err
	}
	fmt.Println("please access this url:", url)
	var vel string
	fmt.Print("enter verifier: ")
	fmt.Scanln(&vel)
	cred, _, err = anaconda.GetCredentials(cred, vel)
	if err != nil {
		log.Println("authentification failed")
		return err
	}
	c := Config{AccessToken: cred.Token, AccessSecret: cred.Secret}
	b, err := json.Marshal(c)
	f, err := os.Create(confPath)
	if err != nil {
		log.Println("can't make conf file")
		return err
	}
	defer f.Close()
	err = ioutil.WriteFile(confPath, b, 0600)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Println("homedir can't get")
		os.Exit(1)
	}
	confPath = filepath.Join(home, ".tw")
}

// New make Config type which has tokens
func New() (Config, error) {
	_, err := os.Stat(confPath)
	if os.IsNotExist(err) {
		err := makeCredential()
		if err != nil {
			return Config{}, nil
		}
	}
	c, err := parseConf()
	return c, err
}
