package main

import (
	"os"
	"projet"
	"strings"
	"fmt"
)

func main() {

	entre := os.Args[1]
	sortie := os.Args[2]

	txt,_ := os.ReadFile(entre)
	var tab []string

	tab = strings.Split(string(txt), " ")
	tab = projet.Format(tab)
	fmt.Println(tab)
	nouveau := projet.Erreur(tab)
	var tabbyte []byte

	for _,mot := range nouveau {
		tabbyte = append(tabbyte, ' ')
		for _,lettre := range mot {
			tabbyte = append(tabbyte, byte(lettre))
		}
	}

	os.WriteFile(sortie,tabbyte, os.FileMode(os.O_CREATE))
}