// Package cred Make twitter credential with anaconda.
// cred file is located $HOME/.tw
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

// Credential type has tokens.
type Credential struct {
	AccessToken  string
	AccessSecret string
}

func parseConf() (Credential, error) {
	confData, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Println("Credential file not found")
		return Credential{}, err
	}
	var c Credential
	err = json.Unmarshal(confData, &c)
	if err != nil {
		log.Println("Credential file can't parsed")
		return Credential{}, err
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
	c := Credential{AccessToken: cred.Token, AccessSecret: cred.Secret}
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

// New make Credential type which has tokens.
// if $HOME/.tw is not exist, make .
func New() (Credential, error) {
	_, err := os.Stat(confPath)
	if os.IsNotExist(err) {
		err := makeCredential()
		if err != nil {
			return Credential{}, nil
		}
	}
	c, err := parseConf()
	return c, err
}
