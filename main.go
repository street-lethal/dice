package main

import (
	"dice/src"
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	Base      int    `json:"base"`
	Length    int    `json:"length"`
	CharKinds string `json:"char_kinds"`
}

func main() {
	var settings Settings
	bin, err := os.ReadFile("./config/settings.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bin, &settings)
	if err != nil {
		panic(err)
	}

	if settings.Base == 0 {
		fmt.Println("Base must not be empty!")
		return
	}

	if settings.Length == 0 {
		fmt.Println("Length must not be empty!")
		return
	}

	if settings.CharKinds == "" {
		fmt.Println("Char kind must not be empty!")
		return
	}

	key, err := src.KeyGen(uint64(settings.Base), "./data/dice.txt")
	if err != nil {
		panic(err)
	}

	g := src.NewStringGenerator(settings.CharKinds)

	necessary, full := g.NecessarySize(settings.Length, settings.Base)
	fmt.Printf("Necessary dice: %d\n", necessary)
	fmt.Printf("Full size: %d\n", full)

	secret := g.Gen(key, settings.Length)

	fmt.Println(secret)
}
