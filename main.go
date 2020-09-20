package main

import (
	"github.com/geraldofigueiredo/mypage/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server := server.NewServer()
	if err := server.Start(); err != nil {
		panic(err)
	}
}
