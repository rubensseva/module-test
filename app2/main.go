package main

import (
	"fmt"
    // "cloud.google.com/go/storage"
	// "context"

     "github.com/rubensseva/module-test/internal/shared"
)

func main() {
	fmt.Println("hello from app2")
	// fmt.Println(storage.NewClient(context.Background()))
	shared.PrintTheThing()
}
