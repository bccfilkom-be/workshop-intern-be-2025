package main

import (
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/bootstrap"
)

func main() {

	if err := bootstrap.Start(); err != nil {
		panic(err)
	}

}
