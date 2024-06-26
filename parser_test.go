package loop_catering_parser

import (
	"github.com/frederikhs/goword"
	"testing"
)

func TestParseMenus(t *testing.T) {
	testCases := []struct {
		Filepath string
		Mains    map[string]string
	}{
		{
			Filepath: "testfiles/Menu Uge 20.docx",
			Mains: map[string]string{
				"Mandag":  "Kylling i karry med grønne asparges og blomkål hertil løse ris",
				"Tirsdag": "Krebinetter med persille kartofler og sauce tatar",
				"Onsdag":  "Kalkun overlår med urter hertil ovn stegte gnocchi, mozzarella og tomat",
				"Torsdag": "Pasta pesto med nye grønsager fra årstiden",
				"Fredag":  "Tacos med langtids braiseret gris",
			},
		},
		{
			Filepath: "testfiles/Menu Uge 21.docx",
			Mains: map[string]string{
				"Mandag":  "",
				"Tirsdag": "Kylling tikka masala med dampede jasmin ris",
				"Onsdag":  "Ungarsk gullasch med luftig mos",
				"Torsdag": "Pankora, dahl og riata",
				"Fredag":  "Kalkun lasagne",
			},
		},
		{
			Filepath: "testfiles/Menu Uge 22.docx",
			Mains: map[string]string{
				"Mandag":  "Paprika kylling med pasta ragout",
				"Tirsdag": "Nakkefilet stegte med hvidløg og salvie hertil ovnstegte kartofler, champion, løg og aubergine",
				"Onsdag":  "Kalkun vindaloo med dampede ris",
				"Torsdag": "Wraps med marineret soya protein salt og dressing",
				"Fredag":  "Græske kødboller med tzatziki og lun bulgur med grønt og urter",
			},
		},
		{
			Filepath: "testfiles/Menu Uge 23.docx",
			Mains: map[string]string{
				"Mandag":  "Kylling i karry og asparges og ris",
				"Tirsdag": "Sprængt kalvespidsbryst med peberrods sauce og nye kartofler og ærter og dild",
				"Onsdag":  "Marineret kalkun overlår med ovnstegte jordskokker, kartofler vendt med urtesalsa hertil tzatziki",
				"Torsdag": "Korma med sødekartofler, blomkål og kikærter hertil fladbrød",
				"Fredag":  "Hamburgerryg med flødekartofler",
			},
		},
		{
			Filepath: "testfiles/Menu Uge 24.docx",
			Mains: map[string]string{
				"Mandag":  "Libanesisk kyllingeret med kold tabbouleh",
				"Tirsdag": "Stegt gris med oyster sauce, ingefær, hvidløg, champion, broccoli og gulerod hertil dampede jasmin ris",
				"Onsdag":  "Frikadeller med pasta pesto, spinat og tomater",
				"Torsdag": "Chili zin carne med brune ris",
				"Fredag":  "Fredags burger",
			},
		},
		{
			Filepath: "testfiles/Menu Uge 25.docx",
			Mains: map[string]string{
				"Mandag":  "Tarteletter med høns i asparges",
				"Tirsdag": "Karbonader med persille kartofler og sauce tartar",
				"Onsdag":  "Kødboller med tomat, svampe og fløde hertil pasta",
				"Torsdag": "Rød karry med frisk grønt og dampede ris",
				"Fredag":  "Pitabrød med krydret hakket okse, salat og dressing",
			},
		},
		{
			Filepath: "testfiles/Menu Uge 26.docx",
			Mains: map[string]string{
				"Mandag":  "Pasta carbonára toppet med ærter og urter",
				"Tirsdag": "Ovnstegt kyllingeoverlår med ovnstegte kartofler, løg, peberfrugt, champion vendt med rød pesto",
				"Onsdag":  "Thailandsk beef green curry med dampede jasmin ris",
				"Torsdag": "Grøntsags pankora med sauce tatar og urtekartofler",
				"Fredag":  "Kylling wraps med salat og dressing",
			},
		},
		{
			Filepath: "testfiles/Menu Uge 27.docx",
			Mains: map[string]string{
				"Mandag":  "Japanske kødboller med japansk nuddelsalat",
				"Tirsdag": "Lørdags kylling med ovnstegte kartofler, spidskål, tomat og courgette",
				"Onsdag":  "Boeuf bourguignon med luftig mos",
				"Torsdag": "Cremet norditaliensk pastaret og urter",
				"Fredag":  "Pita brød med krydret kalkun, salat og dressing",
			},
		},
	}

	for _, testCase := range testCases {
		test := testCase
		t.Run(test.Filepath, func(t *testing.T) {
			t.Parallel()
			content, err := goword.ParseText(test.Filepath)
			if err != nil {
				t.Errorf("could not parse text: %v", err)
			}

			weekMenu := ParseToWeek(content)

			for day, main := range test.Mains {
				if weekMenu[day].Main != main {
					t.Errorf("expected \"%s\", got \"%s\" for file \"%s\" and day \"%s\"", main, weekMenu[day].Main, test.Filepath, day)
				}
			}
		})
	}
}
