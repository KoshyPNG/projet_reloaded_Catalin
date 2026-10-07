package main

import (
	"os"
	"projet"
	"strconv"
)

func main() {

	entre := os.Args[1]
	sortie := os.Args[2]

	txt,_ := os.ReadFile(entre)
	var tab []string
	word := ""
	compri := false
	for _, lettre := range txt {
		if lettre == '(' {
			compri = true
		}
		if lettre == ')' {
			compri = false
		}
		if lettre == ' ' && !compri {
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
	newtab := tab

	for word := 0; word < len(newtab); word++ {
		switch newtab[word] {
		case "(hex)":
			newtab = projet.Convhexa(newtab, word)
			word--
		case "(bin)":
			newtab = projet.Convbin(newtab, word)
			word--
		}
		decoupe := decoupe(newtab[word])
		switch decoupe[0] {
		case "(up)":
			newtab = projet.Convup(newtab,decoupe[1], word)
			val,_ := strconv.Atoi(decoupe[1])
			word -= val
		case "(low)":
			newtab = projet.Convlow(newtab,decoupe[1], word)
			val,_ := strconv.Atoi(decoupe[1])
			word -= val
		case "(cap)":
			newtab = projet.Convcap(newtab,decoupe[1], word)
			val,_ := strconv.Atoi(decoupe[1])
			word -= val	
		}
	}
	return newtab
}

func decoupe(mot string) []string {
	decoupe := []string{}
	val := ""
	for i, lettre := range mot {
		if i == 0 && lettre != '(' {
			decoupe = append(decoupe, "rien")
			return decoupe
		}
		val += string(lettre)
		switch val {
		case "(up":
			decoupe = append(decoupe, "(up)")
		case "(low":
			decoupe = append(decoupe, "(low)")
		case "(cap":
			decoupe = append(decoupe, "(cap)")
		}
		if lettre <= '9' && lettre >= '2' {
			decoupe = append(decoupe, string(lettre))
		}
	}
	if len(decoupe) == 1 {
		decoupe = append(decoupe, "1")
	}
	return decoupe
}