package flect

import "fmt"

var pluralRules = []rule{}

// AddPlural adds a rule that will replace the given suffix with the replacement suffix.
// The name is confusing. This function will be deprecated in the next release.
func AddPlural(suffix string, repl string) {
	InsertPluralRule(suffix, repl)
}

// InsertPluralRule inserts a rule that will replace the given suffix with
// the repl(acement) at the begining of the list of the pluralize rules.
func InsertPluralRule(suffix, repl string) {
	pluralMoot.Lock()
	defer pluralMoot.Unlock()

	pluralRules = append([]rule{{
		suffix: suffix,
		fn:     simpleRuleFunc(suffix, repl),
	}}, pluralRules...)

	pluralRules = append([]rule{{
		suffix: repl,
		fn:     noop,
	}}, pluralRules...)
}

type word struct {
	singular       string
	plural         string
	alternative    string
	unidirectional bool // plural to singular is not possible (or bad)
	uncountable    bool
	exact          bool
}

// dictionary is the main table for singularize and pluralize.
// All words in the dictionary will be added to singleToPlural, pluralToSingle
// and singlePluralAssertions by init() functions.
var dictionary = []word{
	// identicals https://en.wikipedia.org/wiki/English_plurals#Nouns_with_identical_singular_and_plural
	{singular: "aircraft", plural: "aircraft"},
	{singular: "beef", plural: "beef", alternative: "beefs"},
	{singular: "bison", plural: "bison"},
	{singular: "blues", plural: "blues", unidirectional: true},
	{singular: "chassis", plural: "chassis"},
	{singular: "deer", plural: "deer"},
	{singular: "fish", plural: "fish", alternative: "fishes"},
	{singular: "moose", plural: "moose"},
	{singular: "police", plural: "police"},
	{singular: "salmon", plural: "salmon", alternative: "salmons"},
	{singular: "series", plural: "series"},
	{singular: "sheep", plural: "sheep"},
	{singular: "shrimp", plural: "shrimp", alternative: "shrimps"},
	{singular: "species", plural: "species"},
	{singular: "trout", plural: "trout", alternative: "trouts"},
	// -en https://en.wikipedia.org/wiki/English_plurals#Plurals_in_-(e)n
	{singular: "child", plural: "children"},
	{singular: "ox", plural: "oxen", exact: true},
	// apophonic https://en.wikipedia.org/wiki/English_plurals#Apophonic_plurals
	{singular: "foot", plural: "feet"},
	{singular: "goose", plural: "geese"},
	{singular: "louse", plural: "lice"},
	{singular: "man", plural: "men"},
	{singular: "human", plural: "humans"}, // not humen
	{singular: "mouse", plural: "mice"},
	{singular: "tooth", plural: "teeth"},
	{singular: "woman", plural: "women"},
	// misc https://en.wikipedia.org/wiki/English_plurals#Miscellaneous_irregular_plurals
	{singular: "die", plural: "dice"},
	{singular: "person", plural: "people"},

	// Words from French that end in -u add an x; in addition to eau to eaux rule
	{singular: "adieu", plural: "adieux", alternative: "adieus"},
	{singular: "fabliau", plural: "fabliaux"},
	{singular: "bureau", plural: "bureaus", alternative: "bureaux"}, // popular

	{singular: "base", plural: "bases"}, // popular case
	{singular: "basis", plural: "bases", unidirectional: true},

	{singular: "media", plural: "media"}, // popular case: media -> media
	{singular: "medium", plural: "media", alternative: "mediums", unidirectional: true},
	{singular: "stadium", plural: "stadiums", alternative: "stadia"},

	// uncountables
	{singular: "money", plural: "money", uncountable: true},

	// exceptions: -f to -ves, not -fe
	{singular: "dwarf", plural: "dwarfs", alternative: "dwarves"},
	{singular: "hoof", plural: "hoofs", alternative: "hooves"},
	{singular: "thief", plural: "thieves"},
	// exceptions: instead of -f(e) to -ves
	{singular: "chive", plural: "chives"},
	{singular: "hive", plural: "hives"},
	{singular: "move", plural: "moves"},

	// exceptions: instead of -y to -ies
	{singular: "movie", plural: "movies"},
	{singular: "cookie", plural: "cookies"},
}

// singleToPlural is the highest priority map for Pluralize().
// singularToPluralSuffixList is used to build pluralRules for suffixes and
// compound words.
var singleToPlural = map[string]string{
	"alias":       "aliases",
	"alumna":      "alumnae",
	"alumnus":     "alumni",
	"analysis":    "analyses",
	"antenna":     "antennas",
	"antithesis":  "antitheses",
	"apex":        "apexes",
	"appendix":    "appendices",
	"axis":        "axes",
	"bacillus":    "bacilli",
	"bacterium":   "bacteria",
	"bus":         "buses",
	"campus":      "campuses",
	"caucus":      "caucuses",
	"circus":      "circuses",
	"codex":       "codices",
	"concerto":    "concertos",
	"corpus":      "corpora",
	"crisis":      "crises",
	"criterion":   "criteria",
	"curriculum":  "curriculums",
	"datum":       "data",
	"diagnosis":   "diagnoses",
	"ellipsis":    "ellipses",
	"equipment":   "equipment",
	"erratum":     "errata",
	"fez":         "fezzes",
	"focus":       "foci",
	"foo":         "foos",
	"formula":     "formulas",
	"fungus":      "fungi",
	"genus":       "genera",
	"graffito":    "graffiti",
	"grouse":      "grouse",
	"halo":        "halos",
	"hypothesis":  "hypotheses",
	"index":       "indices",
	"information": "information",
	"jeans":       "jeans",
	"larva":       "larvae",
	"libretto":    "librettos",
	"locus":       "loci",
	"matrix":      "matrices",
	"minutia":     "minutiae",
	"nebula":      "nebulae",
	"news":        "news",
	"nucleus":     "nuclei",
	"oasis":       "oases",
	"octopus":     "octopi",
	"offspring":   "offspring",
	"opus":        "opera",
	"ovum":        "ova",
	"parenthesis": "parentheses",
	"phenomenon":  "phenomena",
	"photo":       "photos",
	"phylum":      "phyla",
	"piano":       "pianos",
	"plus":        "pluses",
	"prognosis":   "prognoses",
	"prometheus":  "prometheuses",
	"quiz":        "quizzes",
	"quota":       "quotas",
	"radius":      "radiuses",
	"referendum":  "referendums",
	"ress":        "resses",
	"rice":        "rice",
	"sex":         "sexes",
	"shoe":        "shoes",
	"stimulus":    "stimuli",
	"stratum":     "strata",
	"swine":       "swine",
	"syllabus":    "syllabi",
	"symposium":   "symposiums",
	"synapse":     "synapses",
	"synopsis":    "synopses",
	"testis":      "testes",
	"thesis":      "theses",
	"tuna":        "tuna",
	"vedalia":     "vedalias",
	"vertebra":    "vertebrae",
	"vertix":      "vertices",
	"vita":        "vitae",
	"vortex":      "vortices",
	"you":         "you",
}

// pluralToSingle is the highest priority map for Singularize().
// singularToPluralSuffixList is used to build singularRules for suffixes and
// compound words.
var pluralToSingle = map[string]string{}

// NOTE: This map should not be built as reverse map of singleToPlural since
// there are words that has the same plurals.
func init() {
	// FIXME: remove when the data fully migrated to dictionary.
	for k, v := range singleToPlural {
		pluralToSingle[v] = k
	}
}

// build singleToPlural and pluralToSingle with dictionary
func init() {
	for _, wd := range dictionary {
		if singleToPlural[wd.singular] != "" {
			panic(fmt.Errorf("map singleToPlural already has an entry for %s", wd.singular))
		}
		singleToPlural[wd.singular] = wd.plural

		if !wd.unidirectional {
			if pluralToSingle[wd.plural] != "" {
				panic(fmt.Errorf("map pluralToSingle already has an entry for %s", wd.plural))
			}
			pluralToSingle[wd.plural] = wd.singular
		}

		if wd.alternative != "" {
			if pluralToSingle[wd.alternative] != "" {
				panic(fmt.Errorf("map pluralToSingle already has an entry for %s", wd.alternative))
			}
			pluralToSingle[wd.alternative] = wd.singular
		}
	}
}

type singularToPluralSuffix struct {
	singular string
	plural   string
}

var singularToPluralSuffixList = []singularToPluralSuffix{
	// https://en.wiktionary.org/wiki/Appendix:English_irregular_nouns#Rules
	// Words that end in -f or -fe change -f or -fe to -ves
	{"tive", "tives"}, // exception
	{"eaf", "eaves"},
	{"oaf", "oaves"},
	{"afe", "aves"},
	{"arf", "arves"},
	{"rfe", "rves"},
	{"rf", "rves"},
	{"lf", "lves"},
	{"fe", "ves"}, // previously '[a-eg-km-z]fe' TODO: regex support

	// Words that end in -y preceded by a consonant change -y to -ies
	{"ay", "ays"},
	{"ey", "eys"},
	{"oy", "oys"},
	{"quy", "quies"},
	{"uy", "uys"},
	{"y", "ies"}, // '[^aeiou]y'

	// Words from French that end in -u add an x (eg, château becomes châteaux)
	{"eau", "eaux"}, // it seems like 'eau' is the most popular form of this rule

	{"campus", "campuses"},
	{"person", "people"},
	{"phylum", "phyla"},
	{"randum", "randa"},
	{"actus", "acti"},
	{"adium", "adia"},
	{"basis", "basis"},
	{"child", "children"},
	{"focus", "foci"},
	{"genus", "genera"},
	{"hello", "hellos"},
	{"jeans", "jeans"},
	{"louse", "lice"},
	{"media", "media"},
	{"mouse", "mice"},
	{"oasis", "oasis"},
	{"atum", "ata"},
	{"atus", "atuses"},
	{"base", "bases"},
	{"cess", "cesses"},
	{"dium", "diums"},
	{"eses", "esis"},
	{"iano", "ianos"},
	{"irus", "iri"},
	{"isis", "ises"},
	{"leus", "li"},
	{"mnus", "mni"},
	{"news", "news"},
	{"odex", "odice"},
	{"oose", "eese"},
	{"ouse", "ouses"},
	{"ovum", "ova"},
	{"shoe", "shoes"},
	{"stis", "stes"},
	{"vice", "vices"},
	{"box", "boxes"},
	{"dge", "dges"},
	{"itz", "itzes"},
	{"ium", "ia"},
	{"ize", "izes"},
	{"man", "men"},
	{"nna", "nnas"},
	{"oci", "ocus"},
	{"ode", "odes"},
	{"tum", "ta"},
	{"tus", "tuses"},
	{"ula", "ulae"},
	{"ula", "ulas"},
	{"uli", "ulus"},
	{"use", "uses"},
	{"uss", "usses"},
	{"you", "you"},
	{"ch", "ches"},
	{"ex", "ices"},
	{"io", "ios"},
	{"sh", "shes"},
	{"ss", "sses"},
	{"tz", "tzes"},
	{"va", "vae"},
	{"zz", "zzes"},
	{"o", "oes"},
	{"x", "xes"},
}

func init() {
	for i := len(singularToPluralSuffixList) - 1; i >= 0; i-- {
		InsertPluralRule(singularToPluralSuffixList[i].singular, singularToPluralSuffixList[i].plural)
		InsertSingularRule(singularToPluralSuffixList[i].plural, singularToPluralSuffixList[i].singular)
	}

	// build pluralRule and singularRule with dictionary for compound words
	for _, wd := range dictionary {
		if !wd.exact {
			InsertPluralRule(wd.singular, wd.plural)
			if !wd.unidirectional {
				InsertSingularRule(wd.plural, wd.singular)
			}

			if wd.alternative != "" {
				InsertSingularRule(wd.alternative, wd.singular)
			}
		}
	}
}
