package linkycore

import "github.com/microcosm-cc/bluemonday"

var stripTagsPolicy = bluemonday.StripTagsPolicy()

// TruncateString - Truncate one string with with x words
func TruncateString(str string, length int, omission string) string {
	if length <= 0 {
		return ""
	}

	orgLen := len(str)
	if orgLen <= length {
		return str
	}

	if orgLen > length {
		return str[:length] + omission
	}

	return str[:length]

	// // Support Japanese
	// // Ref: Range loops https://blog.golang.org/strings
	// truncated := ""
	// count := 0
	// for _, char := range str {
	// 	truncated += string(char)
	// 	count++
	// 	if count >= length {
	// 		break
	// 	}
	// }

	// return truncated
}

// StripTags - Remove tags from html text
func StripTags(str string) string {
	return stripTagsPolicy.Sanitize(str)
}

// StripTagsAndTruncate - Remove all tags from string and then truncate if need with omission
func StripTagsAndTruncate(str string, length int, omission string) string {
	return TruncateString(StripTags(str), length, omission)
}
