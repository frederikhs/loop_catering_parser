package loop_catering_parser

import "regexp"

type DayMatcher struct {
	NameOfDay string
	Regex     *regexp.Regexp
}

var DayRegexes = []DayMatcher{
	{
		NameOfDay: "Mandag",
		Regex:     regexp.MustCompile(`Mandag[-|–]+ ?(.*)`),
	},
	{
		NameOfDay: "Tirsdag",
		Regex:     regexp.MustCompile(`Tirsdag[-|–]+ ?(.*)`),
	},
	{
		NameOfDay: "Onsdag",
		Regex:     regexp.MustCompile(`Onsdag[-|–]+ ?(.*)`),
	},
	{
		NameOfDay: "Torsdag",
		Regex:     regexp.MustCompile(`Torsdag[-|–]+ ?(.*)`),
	},
	{
		NameOfDay: "Fredag",
		Regex:     regexp.MustCompile(`Fredag[-|–]+ ?(.*)`),
	},
}
