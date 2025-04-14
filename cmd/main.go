package main

import "mystore/internal/config"

func main() {
	_, err := config.ConnectDB()
	if err != nil {
		return
	}
}
