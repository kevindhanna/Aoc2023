package helpers

import "regexp"

/**
 * Parses s with the given regular expression and returns the
 * group values defined in the expression.
 *
 */
func GetCaptureGroupMap(r, s string) (result map[string]string) {
	var compRegEx = regexp.MustCompile(r)
	match := compRegEx.FindStringSubmatch(s)

	result = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			result[name] = match[i]
		}
	}
	return result
}
