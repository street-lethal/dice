package main

import (
	"dice/src"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	Base      int    `yaml:"base"`
	Length    int    `yaml:"length"`
	CharKinds string `yaml:"char_kinds"`
}

func main() {
	var settings Settings
	bin, _ := os.ReadFile("./config/settings.yml")
	err := yaml.Unmarshal(bin, &settings)
	if err != nil {
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
	}

	g := src.NewStringGenerator(settings.CharKinds)

	necessary, full := g.NecessarySize(settings.Length, settings.Base)
	fmt.Printf("Necessary dice: %d\n", necessary)
	fmt.Printf("Full size: %d\n", full)

	secret := g.Gen(key, settings.Length)

	fmt.Println(secret)
}
