package projet

func Format(tab []string) []string {
	newtab := tab
	for word := 0; word < len(newtab); word++ {
		if newtab[word] == "" {
			newtab = Reecrire(newtab, word)
			word--
		}
	}
	return newtab
}