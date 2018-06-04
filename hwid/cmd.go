package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/foxcpp/go-hwid"
)

func main() {
	id, err := hwid.HWID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(hex.EncodeToString(id))
}
