package main

import (
	"dice/src"
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	Base           int    `json:"base"`
	Length         int    `json:"length"`
	CharKinds      string `json:"char_kinds"`
	ForbiddenChars string `json:"forbidden_chars"`
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

	key, dice, err := src.KeyGen(uint64(settings.Base), "./data/dice.txt")
	if err != nil {
		panic(err)
	}

	g := src.NewStringGenerator(settings.CharKinds)
	g.ForbidCharacters(settings.ForbiddenChars)

	necessary, full := g.NecessarySize(settings.Length, settings.Base)
	if dice < necessary {
		fmt.Printf(
			"[ERROR] Not eough dice, needed %d, actual: %d \n", necessary, dice,
		)
		return
	}
	fmt.Printf("Full size: %d\n", full)

	secret := g.Gen(key, settings.Length)

	fmt.Println(secret)
}
