package main

func reverseWords(s []byte) {
	swap := func(l []byte) {
		for i, j := 0, len(l)-1; i < j; i++ {
			l[i], l[j] = l[j], l[i]
			j--
		}
	}

	swap(s)
	start := 0
	for i, e := range s {
		if e == 32 { // []byte(" ")[0] == 32
			swap(s[start:i])
			start = i + 1
		}
	}
	swap(s[start:len(s)])

}

func main() {

}
