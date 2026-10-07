package main

import (
	"os"
	"projet"
)

func main() {

	entre := os.Args[1]
	sortie := os.Args[2]

	txt,_ := os.ReadFile(entre)
	var tab []string
	word := ""
	for _, lettre := range txt {
		if lettre == ' ' {
			tab = append(tab, string(word))
			word = ""
		} else {
			word += string(lettre)
		}
	}

	tab = append(tab, string(word))
	nouveau := erreur(tab)
	var tabbyte []byte

	for _,mot := range nouveau {
		tabbyte = append(tabbyte, ' ')
		for _,lettre := range mot {
			tabbyte = append(tabbyte, byte(lettre))
		}
	} 
	os.WriteFile(sortie,tabbyte, os.FileMode(os.O_CREATE))
}

func erreur(tab []string) []string {
	hex := -1
	bin := -1
	up := -1
	low := -1
	cap := -1
	newtab := tab

	for i, word := range tab {
		switch word {
		case "(hex)":
			hex = i
		case "(bin)":
			bin = i
		case "(up)":
			up = i
		case "(low)":
			low = i
		case "(cap)":
			cap = i
		}
	}
	if hex != -1{
		newtab = projet.Convhexa(newtab,hex)
	}
	if bin != -1 {
		if hex != -1 && hex < bin {
			bin--
		}
		newtab = projet.Convbin(newtab,bin)
	}
	if up != -1 {
		newtab = projet.Convup(newtab,up)
	}
	if low != -1 {
		newtab = projet.Convlow(newtab,low)
	}
	if cap != -1 {
		newtab = projet.Convcap(newtab,cap)
	}

	return newtab
}

