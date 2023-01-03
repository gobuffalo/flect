package flect

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type tt struct {
	act string
	exp string
}

func Test_LoadInflections(t *testing.T) {
	r := require.New(t)
	m := map[string]string{
		"baby": "bebe",
		"xyz":  "zyx",
	}

	b, err := json.Marshal(m)
	r.NoError(err)

	r.NoError(LoadInflections(bytes.NewReader(b)))

	for k, v := range m {
		r.Equal(v, Pluralize(k))
		r.Equal(v, Pluralize(v))
		r.Equal(k, Singularize(k))
		r.Equal(k, Singularize(v))
	}
}

func Test_LoadInflectionsWrongSingular(t *testing.T) {
	r := require.New(t)
	m := map[string]string{
		"a file": "files",
	}

	b, err := json.Marshal(m)
	r.NoError(err)

	r.Error(LoadInflections(bytes.NewReader(b)))
}

func Test_LoadInflectionsWrongPlural(t *testing.T) {
	r := require.New(t)
	m := map[string]string{
		"beatle": "the beatles",
	}

	b, err := json.Marshal(m)
	r.NoError(err)

	r.Error(LoadInflections(bytes.NewReader(b)))
}

func Test_LoadAcronyms(t *testing.T) {
	r := require.New(t)
	m := []string{
		"ACC",
		"TLC",
		"LSA",
	}

	b, err := json.Marshal(m)
	r.NoError(err)

	r.NoError(LoadAcronyms(bytes.NewReader(b)))

	for _, acronym := range m {
		r.True(baseAcronyms[acronym])
	}
}

type dict struct {
	singular          string
	plural            string
	doSingularizeTest bool
	doPluralizeTest   bool
}

var singlePluralAssertions = []dict{
	{"", "", true, true},
	{"Car", "Cars", true, true},
	{"Boy", "Boys", true, true},
	{"GoodBoy", "GoodBoys", true, true},
	{"Axis", "Axes", true, true},
	{"Child", "Children", true, true},
	{"GoodChild", "GoodChildren", true, true},
	{"node_child", "node_children", true, true},
	{"SmartPerson", "SmartPeople", true, true},
	{"great_person", "great_people", true, true},
	{"salesperson", "salespeople", true, true},
	{"SuperbOx", "SuperbOxen", true, true},
	{"WildOx", "WildOxen", true, true},
	{"wild_ox", "wild_oxen", true, true},
	{"address", "addresses", true, true},
	{"alias", "aliases", true, true},
	{"analysis", "analyses", true, true},
	{"axis", "axes", true, true},
	{"blitz", "blitzes", true, true},
	{"box", "boxes", true, true},
	{"buffalo", "buffaloes", true, true},
	{"case", "cases", true, true},
	{"cat", "cats", true, true},
	{"collapse", "collapses", true, true},
	{"comment", "comments", true, true},
	{"crisis", "crises", true, true},
	{"custom_field", "custom_fields", true, true},
	{"database", "databases", true, true},
	{"dear", "dears", true, true},
	{"device", "devices", true, true},
	{"diagnosis_a", "diagnosis_as", true, true},
	{"diagnosis", "diagnoses", true, true},
	{"eclipse", "eclipses", true, true},
	{"edge", "edges", true, true},
	{"ellipsis", "ellipses", true, true},
	{"equipment", "equipment", true, true},
	{"experience", "experiences", true, true},
	{"fez", "fezzes", true, true},
	{"field", "fields", true, true},
	{"fix", "fixes", true, true},
	{"fleet", "fleets", true, true},
	{"foobar", "foobars", true, true},
	{"fox", "foxes", true, true},
	{"funky jeans", "funky jeans", true, true},
	{"glimpse", "glimpses", true, true},
	{"halo", "halos", true, true},
	{"horse", "horses", true, true},
	{"index", "indices", true, true},
	{"jeans", "jeans", true, true},
	{"lapse", "lapses", true, true},
	{"lunch", "lunches", true, true},
	{"marsh", "marshes", true, true},
	{"matrix", "matrices", true, true},
	{"mouse", "mice", true, true},
	{"news", "news", true, true},
	{"newsletter", "newsletters", true, true},
	{"payment_information", "payment_information", true, true},
	{"pepsi", "pepsis", true, true},
	{"photo", "photos", true, true},
	{"piano", "pianos", true, true},
	{"portfolio", "portfolios", true, true},
	{"potato", "potatoes", true, true},
	{"prize", "prizes", true, true},
	{"process", "processes", true, true},
	{"psi", "psis", true, true},
	{"quiz", "quizzes", true, true},
	{"rice", "rice", true, true},
	{"search", "searches", true, true},
	{"service", "services", true, true},
	{"shoe", "shoes", true, true},
	{"sportsEquipment", "sportsEquipment", true, true},
	{"stack", "stacks", true, true},
	{"status_code", "status_codes", true, true},
	{"switch", "switches", true, true},
	{"tax", "taxes", true, true},
	{"testis", "testes", true, true},
	{"tomato", "tomatoes", true, true},
	{"truss", "trusses", true, true},
	{"user", "users", true, true},
	{"user_custom_field", "user_custom_fields", true, true},
	{"wish", "wishes", true, true},
	{"woman", "women", true, true},
	{"human", "humans", true, true},
	{"spokesman", "spokesmen", true, true},

	// Words that end in -f or -fe change -f or -fe to -ves
	// https://en.wiktionary.org/wiki/Category:English_irregular_plurals_ending_in_"-ves"
	{"calf", "calves", true, true},
	{"dwarf", "dwarves", true, false}, // dwarfs looks popular than dwarves
	{"dwarf", "dwarfs", true, true},
	{"elf", "elves", true, true},
	{"half", "halves", true, true},
	{"knife", "knives", true, true},
	{"leaf", "leaves", true, true},
	{"life", "lives", true, true},
	{"loaf", "loaves", true, true},
	{"safe", "saves", true, true},
	{"scarf", "scarves", true, true},
	{"self", "selves", true, true},
	{"sheaf", "sheaves", true, true},
	{"shelf", "shelves", true, true},
	{"thief", "thieves", true, true},
	{"wharf", "wharves", true, true},
	{"wife", "wives", true, true},
	{"wolf", "wolves", true, true},
	{"beef", "beef", true, true},
	{"belief", "beliefs", true, true},
	{"chef", "chefs", true, true},
	{"chief", "chiefs", true, true},
	{"hoof", "hoofs", true, true},
	{"kerchief", "kerchiefs", true, true},
	{"roof", "roofs", true, true},
	{"archive", "archives", true, true},
	{"perspective", "perspectives", true, true},

	// Words that end in -y preceded by a consonant change -y to -ies
	{"day", "days", true, true},
	{"hobby", "hobbies", true, true},
	{"agency", "agencies", true, true},
	{"body", "bodies", true, true},
	{"abbey", "abbeys", true, true},
	{"jiffy", "jiffies", true, true},
	{"geology", "geologies", true, true},
	{"geography", "geographies", true, true},
	{"cookie", "cookies", true, true},
	{"sky", "skies", true, true},
	{"supply", "supplies", true, true},
	{"academy", "academies", true, true},
	{"ebony", "ebonies", true, true},
	{"boy", "boys", true, true},
	{"copy", "copies", true, true},
	{"category", "categories", true, true},
	{"embassy", "embassies", true, true},
	{"entity", "entities", true, true},
	{"guy", "guys", true, true},
	{"obsequy", "obsequies", true, true},
	{"navy", "navies", true, true},
	{"proxy", "proxies", true, true},
	{"crazy", "crazies", true, true},

	// Words from French that end in -u add an x
	{"aboideau", "aboideaux", true, true},
	{"beau", "beaux", true, true},
	{"château", "châteaux", true, true},
	{"chateau", "chateaux", true, true},
	{"fabliau", "fabliaux", true, true},
	{"tableau", "tableaux", true, true},
	{"bureau", "bureaus", true, true},
	{"adieu", "adieux", true, true},

	// Words from Greek that end in -on change -on to -a
	{"tetrahedron", "tetrahedra", true, true},

	// Words from Latin that end in -um change -um to -a
	{"stadium", "stadiums", true, true},
	{"stadium", "stadia", true, false},
	{"aquarium", "aquaria", true, true},
	{"auditorium", "auditoria", true, true},
	{"bacterium", "bacteria", true, true},
	{"pretorium", "pretoriums", true, true},
	{"symposium", "symposia", true, true},
	{"symposium", "symposiums", true, false},
	{"amoebaeum", "amoebaea", true, true},
	{"coliseum", "coliseums", true, true},
	{"museum", "museums", true, true},
	{"agenda", "agendas", true, true},
	{"curriculum", "curriculums", true, true},
	{"collum", "colla", true, true},
	{"datum", "data", true, true},
	{"erratum", "errata", true, true},
	{"maximum", "maxima", true, true},
	{"platinum", "platinums", true, true},
	{"serum", "sera", true, true},
	{"spectrum", "spectra", true, true},

	// Words from Latin that end in -us change -us to -i or -era
	{"opera", "operas", true, true},

	// Words from Latin that end in -a change -a to -ae
	{"alumna", "alumnae", true, true},
	{"larva", "larvae", true, true},
	{"minutia", "minutiae", true, true},
	{"nebula", "nebulae", true, true},
	{"vertebra", "vertebrae", true, true},
	{"vita", "vitae", true, true},
	{"antenna", "antennas", true, true},
	{"formula", "formulas", true, true},
	{"tuna", "tuna", true, true},
	{"quota", "quotas", true, true},
	{"vedalia", "vedalias", true, true},
	{"media", "media", true, true}, // instead of mediae, popular case
	{"multimedia", "multimedia", true, true},

	// Words that end in -ch, -o, -s, -sh, -x, -z
	{"bus", "buses", true, true}, // end with u + s, but no -us rule
	{"campus", "campuses", true, true},
	{"caucus", "caucuses", true, true},
	{"circus", "circuses", true, true},
	{"plus", "pluses", true, true},
	{"prometheus", "prometheuses", true, true},
	{"status", "statuses", true, true},
	{"virus", "viruses", true, true},
	{"use", "uses", true, true}, // -use
	{"fuse", "fuses", true, true},
	{"house", "houses", true, true},
	{"spouse", "spouses", true, true},
}

func init() {
	for _, wd := range dictionary {
		if wd.uncountable && wd.plural == "" {
			wd.plural = wd.singular
		}

		singlePluralAssertions = append(singlePluralAssertions, dict{
			singular:          wd.singular,
			plural:            wd.plural,
			doSingularizeTest: !wd.unidirectional,
		})

		if wd.alternative != "" {
			singlePluralAssertions = append(singlePluralAssertions, dict{
				singular:        wd.singular,
				plural:          wd.alternative,
				doPluralizeTest: false,
			})
		}
	}
}
