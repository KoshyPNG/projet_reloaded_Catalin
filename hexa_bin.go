package projet

import "strconv"

func Convhexa(tab []string, i int) []string {
	val,_ := strconv.ParseInt(tab[i-1], 16, 64)
	tab[i-1] = strconv.FormatInt(val, 10)
	return Reecrire(tab, i)
}

func Convbin(tab []string, i int) []string {
	val,_ := strconv.ParseInt(tab[i-1], 2, 64)
	tab[i-1] = strconv.FormatInt(val, 10)
	return Reecrire(tab, i)
}
