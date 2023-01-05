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
	{"custom_field", "custom_fields", true, true},
	{"funky jeans", "funky jeans", true, true},
	{"payment_information", "payment_information", true, true},
	{"sportsEquipment", "sportsEquipment", true, true},
	{"status_code", "status_codes", true, true},
	{"user_custom_field", "user_custom_fields", true, true},
	{"SuperbOx", "SuperbOxen", true, true},
	{"WildOx", "WildOxen", true, true},
	{"wild_ox", "wild_oxen", true, true},
	{"box", "boxes", true, true},
	{"fox", "foxes", true, true},
	{"comment", "comments", true, true},
	{"edge", "edges", true, true},
	{"equipment", "equipment", true, true},
	{"experience", "experiences", true, true},
	{"fleet", "fleets", true, true},
	{"foobar", "foobars", true, true},
	{"mouse", "mice", true, true},
	{"newsletter", "newsletters", true, true},
	{"stack", "stacks", true, true},
	{"user", "users", true, true},
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
	{"lunch", "lunches", true, true},
	{"search", "searches", true, true},
	{"switch", "switches", true, true},
	{"headache", "headaches", true, true}, // ch vs. che
	{"marsh", "marshes", true, true},
	{"wish", "wishes", true, true},

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

	{"quiz", "quizzes", true, true},
	{"buzz", "buzzes", true, true},
	{"blitz", "blitzes", true, true},
	{"quartz", "quartzes", true, true},
	{"topaz", "topazes", true, true},
	{"waltz", "waltzes", true, true},
	{"prize", "prizes", true, true},

	{"access", "accesses", true, true},
	{"process", "processes", true, true},
	{"address", "addresses", true, true},
	{"case", "cases", true, true},
	{"database", "databases", true, true},
	{"glimpse", "glimpses", true, true},
	{"horse", "horses", true, true},
	{"lapse", "lapses", true, true},
	{"collapse", "collapses", true, true},
	{"truss", "trusses", true, true},

	{"portfolio", "portfolios", true, true}, // -o -os
	{"piano", "pianos", true, true},         // -ano -anos
	{"hello", "hellos", true, true},         // -lo -los
	{"buffalo", "buffaloes", true, true},    // -lo -loes
	{"photo", "photos", true, true},         // -to -tos
	{"potato", "potatoes", true, true},      // exception of -to -tos
	{"tomato", "tomatoes", true, true},
	{"graffiti", "graffiti", true, true},
	{"foo", "foos", true, true},
	{"zoo", "zoos", true, true},

	// Words from Latin that end in -ex change -ex to -ices
	// Words from Latin that end in -ix change -ix to -ices
	{"appendix", "appendices", true, true}, // -dix
	{"codex", "codices", true, true},       // -dex
	{"index", "indices", true, true},
	{"bodice", "bodices", true, true},    // -dice
	{"helix", "helices", true, true},     // -lix
	{"complex", "complexes", true, true}, // -lex
	{"duplex", "duplexes", true, true},
	{"accomplice", "accomplices", true, true}, // -lice
	{"slice", "slices", true, true},
	{"matrix", "matrices", true, true},  // -trix
	{"justice", "justices", true, true}, // -tice
	{"lattice", "lattices", true, true},
	{"notice", "notices", true, true},
	{"apex", "apices", true, true},    // -pex
	{"spice", "spices", true, true},   // -pice
	{"device", "devices", true, true}, // -vice
	{"service", "services", true, true},
	{"fix", "fixes", true, true},        // -ix
	{"sex", "sexes", true, true},        // -ex
	{"invoice", "invoices", true, true}, // gobuffalo/flect#61
	{"voice", "voices", true, true},
	{"choice", "choices", true, true},

	// Words from Latin that end in -is change -is to -es
	{"axis", "axes", true, true},
	{"tax", "taxes", true, true}, // not taxis
	{"eclipse", "eclipses", true, true},
	{"ellipse", "ellipses", true, true},
	{"ellipsis", "ellipses", false, true}, // pluralize only
	{"oasis", "oases", true, true},
	{"thesis", "theses", true, true}, // word thesis
	{"hypothesis", "hypotheses", true, true},
	{"parenthesis", "parentheses", true, true},
	{"analysis", "analyses", true, true}, // suffix lysis
	{"antithesis", "antitheses", true, true},
	{"diagnosis", "diagnoses", true, true}, // suffix gnosis
	{"prognosis", "prognoses", true, true},
	{"synopsis", "synopses", true, true}, // suffix opsis
	{"synapse", "synapses", true, true},
	{"waste", "wastes", true, true},
	{"psi", "psis", true, true},
	{"pepsi", "pepsis", true, true},
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
