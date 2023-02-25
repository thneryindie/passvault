package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Credential struct {
	Name		string
	Login		string
	Url			string
	Password	string
}

type CredentialsFile struct {
	Data map[string]map[string]string `yaml:",inline"`
}

func createCredentials(name, login, url, password string) {
	fmt.Println("Create Credentials For:")
	cred := Credential{
		Name: name, 
		Login: login, 
		Url: url, 
		Password: password,
	}

	filePath := writeCredential(cred)

	fmt.Println(filePath)
} 

func writeCredential(cred Credential) string {
	fmt.Println(cred)
	return "file path"
}

func readCredentials(key string) string {
	data, err := ioutil.ReadFile(".passwords.yml")

	if err != nil {
		panic(err)
	}

	var credFile CredentialsFile
	err = yaml.Unmarshal(data, &credFile)

	if err != nil {
		panic(err)
	}

	value, ok := credFile.Data[key]

	if !ok {
		panic(fmt.Sprintf("key %q not found", key))
	}

	for subkey, subvalue := range value {
		fmt.Printf("%s: %s\n", subkey, subvalue)
	}

	return ""
}

func main() {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	credName := createCmd.String("name", "", "name")
	credLogin := createCmd.String("login", "", "login")
	credUrl := createCmd.String("url", "", "url")
	credPassword := createCmd.String("password", "", "password")

	readCmd := flag.NewFlagSet("read", flag.ExitOnError)
	key := readCmd.String("key", "", "key")

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		createCredentials(*credName, *credLogin, *credUrl, *credPassword)
	case "read":
		readCmd.Parse(os.Args[2:])
		readCredentials(*key)
	}
}
