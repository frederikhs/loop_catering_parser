package loop_catering_parser

import (
	"regexp"
	"strings"
)

var characterReplacers = map[string]string{
	"–": "-",
}

func ReplaceCharacterMistakes(in string) string {
	for from, to := range characterReplacers {
		in = strings.ReplaceAll(in, from, to)
	}

	return in
}

var dietaryMatcher = regexp.MustCompile(` ?\(*[A-ZÆØÅ]+\) ?`)

func ReplaceDietaryInformationBadges(in string) string {
	return dietaryMatcher.ReplaceAllString(in, "")
}

func Clean(in string) string {
	in = ReplaceCharacterMistakes(in)
	in = ReplaceDietaryInformationBadges(in)

	return in
}
