package piscine

func Atoi(s string) int {
	a := 0
	b := 1
	if len(s) == 0 {
		return 0
	}
	if int(s[0]) == 45 {
		b = -1
	}
	for i := 0; i < len(s); i++ {
		if i == 0 && int(s[0]) == 45 || i == 0 && int(s[0]) == 43 {
			continue
		}
		if int(s[i]) < 48 || int(s[i]) > 57 {
			return 0
		} else {
			a = a*10 + (int(s[i]) - 48)
		}
	}
	return a * b
}
