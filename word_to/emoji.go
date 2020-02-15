package word_to

import "strings"

func Emoji(words string) string {
	var new_string string
	for _, word := range strings.Fields(words) {
		switch strings.ToLower(word) {
		case "hello":
			new_string += "👋"
		case "world":
			new_string += "🌎"
		default:
			new_string += ""
		}
	}
	return new_string
}
