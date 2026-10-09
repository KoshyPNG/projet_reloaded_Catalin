package projet

func Erreur(tab []string) []string {
	newtab := tab

	for word := 0; word < len(newtab); word++ {
		switch newtab[word] {
		case "(hex)":
			newtab = Convhexa(newtab, word)
			word--
		case "(bin)":
			newtab = Convbin(newtab, word)
			word--
		case "(up)" :
			newtab = Convup(newtab, "1", word)
			word--
		case "(up,":
			rep := Repetion(newtab[word+1])
			newtab = Reecrire(newtab, word+1)
			newtab = Convup(newtab, rep, word)
			word--
			word--
		case "(low)" :
			newtab = Convlow(newtab, "1", word)
			word--
		case "(low,":
			rep := Repetion(newtab[word+1])
			newtab = Reecrire(newtab, word+1)
			newtab = Convlow(newtab, rep, word)
			word--
			word--
		case "(cap)" :
			newtab = Convcap(newtab, "1", word)
			word--
		case "(cap,":
			rep := Repetion(newtab[word+1])
			newtab = Reecrire(newtab, word+1)
			newtab = Convcap(newtab, rep, word)
			word--
			word--
		}
	}
	newtab = Format(newtab)
	newtab = Ponctuation(newtab)
	return newtab
}

func Repetion(word string) string {
	for _, lettre := range word {
		if lettre >= '1' && lettre <= '9' {
			return string(lettre)
		}
	}
	return "1"
}