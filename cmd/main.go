package main

import (
	"fmt"
	"github.com/ansu-francis/envloader"
	"os"
)

func main() {
	if err := envloader.LoadEnv("../config.env"); err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(os.Getenv("TEST"))	
	fmt.Println(os.Getenv("ABC"))
}

