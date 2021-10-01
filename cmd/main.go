package main

import (
	"fmt"

	"github.com/jkodirov/wallet/v1/pkg/wallet"
)

func main() {
	svc := &wallet.Service{}
	svc.RegisterAccount("+998998029829")
	fmt.Println(svc)
}