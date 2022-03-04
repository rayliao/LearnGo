package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/rayliao/hello/app"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	poker.NewCLI(store, os.Stdin).PlayPoker()
}
