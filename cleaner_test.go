package loop_catering_parser

import (
	"fmt"
	"testing"
)

func TestReplaceCharacterMistakes(t *testing.T) {
	testCases := []struct {
		Before string
		After  string
	}{
		{
			Before: "–",
			After:  "-",
		},
		{
			Before: "– Pankora, dahl og riata (V)(L)(LØ)(H)",
			After:  "- Pankora, dahl og riata (V)(L)(LØ)(H)",
		},
		{
			Before: "–Pankora, dahl og riata (V)(L)(LØ)(H)",
			After:  "-Pankora, dahl og riata (V)(L)(LØ)(H)",
		},
		{
			Before: "Pankora, dahl og riata (V)(L)(LØ)(H)–",
			After:  "Pankora, dahl og riata (V)(L)(LØ)(H)-",
		},
		{
			Before: "Pankora, dahl og riata –(V)(L)(LØ)(H)",
			After:  "Pankora, dahl og riata -(V)(L)(LØ)(H)",
		},
	}

	for _, testCase := range testCases {
		shorted := ReplaceCharacterMistakes(testCase.Before)
		if shorted != testCase.After {
			t.Errorf("expected \"%s\", got \"%s\"", testCase.After, shorted)
		}
	}
}

func TestReplaceDietaryInformationBadges(t *testing.T) {
	testCases := []struct {
		Before string
		After  string
	}{
		{
			Before: "Kylling tikka masala med dampede jasmin ris (LØ)(H)",
			After:  "Kylling tikka masala med dampede jasmin ris",
		},
		{
			Before: "Kage: Gulerods kage(L)(G)(Æ)",
			After:  "Kage: Gulerods kage",
		},
		{
			Before: "Laksesalat (L)(LØ) ",
			After:  "Laksesalat",
		},
		{
			Before: "Skagensalat (L)(G)(LØ)(Æ)(SK)",
			After:  "Skagensalat",
		},
		{
			Before: "Skagensalat (med ekstra skagen)",
			After:  "Skagensalat (med ekstra skagen)",
		},
		{
			Before: "Pasta pesto med nye grønsager fra årstiden ((V)G)(H)",
			After:  "Pasta pesto med nye grønsager fra årstiden",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			shorted := ReplaceDietaryInformationBadges(testCase.Before)
			if shorted != testCase.After {
				t.Errorf("expected \"%s\", got \"%s\"", testCase.After, shorted)
			}
		})
	}
}
