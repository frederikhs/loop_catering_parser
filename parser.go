package loop_catering_parser

import (
	"strings"
)

type Day struct {
	IsClosed bool
	Main     string
	Sides    []string
}

func parseToMeal(in string) Day {
	parts := strings.Split(in, "- ")

	main := Clean(parts[0])

	if main == "Lukket" {
		return Day{
			IsClosed: true,
		}
	}

	m := Day{
		Main:  main,
		Sides: nil,
	}

	for _, a := range parts[0:] {
		m.Sides = append(m.Sides, Clean(a))
	}

	return m
}

type WeekMenu map[string]Day

func ParseToWeek(in string) WeekMenu {
	wm := make(WeekMenu)

	for _, dr := range DayRegexes {
		body := dr.Regex.FindStringSubmatch(in)
		meal := parseToMeal(body[1])
		wm[dr.NameOfDay] = meal
	}

	return wm
}
