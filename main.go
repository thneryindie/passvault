package main

import (
	"flag"
	"fmt"
	"os"
)

type credential struct {
	name		string
	login		string
	url			string
	password	string
}

func createCredentials(name, login, url, password string) {
	fmt.Println("Create Credentials For:")
	cred := credential{
		name: name, 
		login: login, 
		url: url, 
		password: password,
	}

	filePath := writeCredential(cred)

	fmt.Println(filePath)
} 

func writeCredential(cred credential) string {
	fmt.Println(cred)
	return "file path"
}

func main() {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	credName := createCmd.String("name", "", "name")
	credLogin := createCmd.String("login", "", "login")
	credUrl := createCmd.String("url", "", "url")
	credPassword := createCmd.String("password", "", "password")

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		createCredentials(*credName, *credLogin, *credUrl, *credPassword)
	}
}
