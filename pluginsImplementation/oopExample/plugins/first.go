package main

import "fmt"

type myPlugin string

func (h myPlugin) Talk() {
	fmt.Println("Hello FROM PLUGIN!!!")
}

var MyPlugin myPlugin

