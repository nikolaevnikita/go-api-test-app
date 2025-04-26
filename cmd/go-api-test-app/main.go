package main

import (
	"fmt"

	"github.com/nikolaevnikita/go-api-test-app/internal/app"
	"github.com/nikolaevnikita/go-api-test-app/internal/config"
)

func main() {
	config := config.ReadConfig()
	if err := app.NewApp(config).Start(); err != nil {
		fmt.Printf("=== App was not started due to an error: %s ===\n", err.Error())
	}
}

// TODO:
// 1. Добавить логирование
// 2. Добавить авторизацию
// 3. Связать Task c User