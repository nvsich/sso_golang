package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	cfg := config.MustLoadConfig()
	fmt.Printf("%#v\n", cfg)

	// TODO: initialize logger

	// TODO: initialize app

	// TODO: run grpc server
}
