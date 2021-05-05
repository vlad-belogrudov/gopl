// Package uniq helps to remove duplicated strings in a slice
package uniq

func Uniq(words []string) []string {
	i := 0
	var previous string
	for _, word := range words {
		if word != previous {
			words[i] = word
			previous = word
			i++
		}
	}
	return words[:i]
}
