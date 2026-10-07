package projet

func Reecrire(tab []string,i int) []string {
	var new []string
	for y, word := range tab {
		if y != i {
			new = append(new, word)
		}
	}
	return new
}