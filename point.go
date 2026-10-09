package projet

import "strings"

func Ponctuation(tab []string) []string {
	return FormatPoint(tab)
}

func FormatPoint(tab []string) []string {
	newtab := tab
	var new string
	for word := 0; word < len(newtab); word++ {
		if strings.HasPrefix(newtab[word], ".") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], ".")
				newtab[word] = new
				newtab[word-1] += "."
				word--
			}
		} else if strings.HasPrefix(newtab[word], ",") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], ",")
				newtab[word] = new
				newtab[word-1] += ","
				word--
			}
		} else if strings.HasPrefix(newtab[word], "!") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], "!")
				newtab[word] = new
				newtab[word-1] += "!"
				word--
			}
		} else if strings.HasPrefix(newtab[word], "?") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], "?")
				newtab[word] = new
				newtab[word-1] += "?"
				word--
			}
		} else if strings.HasPrefix(newtab[word], ":") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], ":")
				newtab[word] = new
				newtab[word-1] += ":"
				word--
			}
		} else if strings.HasPrefix(newtab[word], ";") {
			if word > 0 {
				new = strings.TrimPrefix(newtab[word], ";")
				newtab[word] = new
				newtab[word-1] += ";"
				word--
			}
		}
	}
	newtab = Format(newtab)
	return newtab
}
