package main

import (
	"github.com/nikolaevnikita/go-api-test-app/internal/app"
	"fmt"
)

func main() {
	if err := app.NewApp().Start(); err != nil {
		fmt.Printf("=== App was not started due to an error: %s ===\n", err.Error())
	}
}

// TODO:
// 1. Добавить логирование
// 2. Добавить авторизацию
// 3. Связать Task c User