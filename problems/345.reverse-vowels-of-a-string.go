package main

func reverseVowels(s string) string {
	c := []int{}
	for i, v := range s {
		if string(v) == "a" ||
			string(v) == "i" ||
			string(v) == "u" ||
			string(v) == "e" ||
			string(v) == "o" ||
			string(v) == "A" ||
			string(v) == "I" ||
			string(v) == "U" ||
			string(v) == "E" ||
			string(v) == "O" {
			c = append(c, i)
		}
	}
	b := []byte(s)

	l := len(c)
	for i := 0; i < l/2; i++ {
		b[c[i]] = s[c[l-1-i]]
		b[c[l-1-i]] = s[c[i]]
	}
	return string(b)
}
