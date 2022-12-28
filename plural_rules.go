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
	{singular: "bison", plural: "bison"},
	{singular: "deer", plural: "deer"},
	{singular: "moose", plural: "moose"},
	{singular: "fish", plural: "fish", alternative: "fishes"},
	{singular: "salmon", plural: "salmon", alternative: "salmons"},
	{singular: "sheep", plural: "sheep"},
	{singular: "shrimp", plural: "shrimp", alternative: "shrimps"},
	{singular: "trout", plural: "trout", alternative: "trouts"},
	{singular: "aircraft", plural: "aircraft"},
	{singular: "blues", plural: "blues", unidirectional: true},
	{singular: "chassis", plural: "chassis"},
	{singular: "series", plural: "series"},
	{singular: "species", plural: "species"},
	{singular: "police", plural: "police"},
	// -en https://en.wikipedia.org/wiki/English_plurals#Plurals_in_-(e)n
	{singular: "ox", plural: "oxen", exact: true},
	{singular: "child", plural: "children"},
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
	{singular: "person", plural: "people"},
	{singular: "die", plural: "dice"},

	{singular: "base", plural: "bases"}, // popular case
	{singular: "basis", plural: "bases", unidirectional: true},

	{singular: "media", plural: "media"}, // popular case: media -> media
	{singular: "medium", plural: "media", alternative: "mediums", unidirectional: true},
	{singular: "stadium", plural: "stadiums", alternative: "stadia"},
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
	"beau":        "beaus",
	"bureau":      "bureaus",
	"bus":         "buses",
	"campus":      "campuses",
	"caucus":      "caucuses",
	"château":     "châteaux",
	"circus":      "circuses",
	"codex":       "codices",
	"concerto":    "concertos",
	"corpus":      "corpora",
	"crisis":      "crises",
	"criterion":   "criteria",
	"curriculum":  "curriculums",
	"datum":       "data",
	"diagnosis":   "diagnoses",
	"dwarf":       "dwarves",
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
	"half":        "halves",
	"halo":        "halos",
	"hoof":        "hooves",
	"hypothesis":  "hypotheses",
	"index":       "indices",
	"information": "information",
	"jeans":       "jeans",
	"larva":       "larvae",
	"libretto":    "librettos",
	"loaf":        "loaves",
	"locus":       "loci",
	"matrix":      "matrices",
	"minutia":     "minutiae",
	"money":       "money",
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
	"tableau":     "tableaus",
	"testis":      "testes",
	"thesis":      "theses",
	"thief":       "thieves",
	"tuna":        "tuna",
	"vedalia":     "vedalias",
	"vertebra":    "vertebrae",
	"vertix":      "vertices",
	"vita":        "vitae",
	"vortex":      "vortices",
	"wharf":       "wharves",
	"wife":        "wives",
	"wolf":        "wolves",
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
	{"campus", "campuses"},
	{"person", "people"},
	{"phylum", "phyla"},
	{"randum", "randa"},
	{"actus", "acti"},
	{"adium", "adia"},
	{"basis", "basis"},
	{"child", "children"},
	{"chive", "chives"},
	{"focus", "foci"},
	{"genus", "genera"},
	{"hello", "hellos"},
	{"jeans", "jeans"},
	{"louse", "lice"},
	{"media", "media"},
	{"mouse", "mice"},
	{"movie", "movies"},
	{"oasis", "oasis"},
	{"atum", "ata"},
	{"atus", "atuses"},
	{"base", "bases"},
	{"cess", "cesses"},
	{"dium", "diums"},
	{"eses", "esis"},
	{"half", "halves"},
	{"hive", "hives"},
	{"iano", "ianos"},
	{"irus", "iri"},
	{"isis", "ises"},
	{"leus", "li"},
	{"mnus", "mni"},
	{"move", "moves"},
	{"news", "news"},
	{"odex", "odice"},
	{"oose", "eese"},
	{"ouse", "ouses"},
	{"ovum", "ova"},
	{"shoe", "shoes"},
	{"stis", "stes"},
	{"tive", "tives"},
	{"vice", "vices"},
	{"wife", "wives"},
	{"afe", "aves"},
	{"bfe", "bves"},
	{"box", "boxes"},
	{"cfe", "cves"},
	{"dfe", "dves"},
	{"dge", "dges"},
	{"efe", "eves"},
	{"gfe", "gves"},
	{"hfe", "hves"},
	{"ife", "ives"},
	{"itz", "itzes"},
	{"ium", "ia"},
	{"ize", "izes"},
	{"jfe", "jves"},
	{"kfe", "kves"},
	{"man", "men"},
	{"mfe", "mves"},
	{"nfe", "nves"},
	{"nna", "nnas"},
	{"oaf", "oaves"},
	{"oci", "ocus"},
	{"ode", "odes"},
	{"ofe", "oves"},
	{"pfe", "pves"},
	{"qfe", "qves"},
	{"quy", "quies"},
	{"rfe", "rves"},
	{"sfe", "sves"},
	{"tfe", "tves"},
	{"tum", "ta"},
	{"tus", "tuses"},
	{"ufe", "uves"},
	{"ula", "ulae"},
	{"ula", "ulas"},
	{"uli", "ulus"},
	{"use", "uses"},
	{"uss", "usses"},
	{"vfe", "vves"},
	{"wfe", "wves"},
	{"xfe", "xves"},
	{"yfe", "yves"},
	{"you", "you"},
	{"zfe", "zves"},
	{"by", "bies"},
	{"ch", "ches"},
	{"cy", "cies"},
	{"dy", "dies"},
	{"ex", "ices"},
	{"fy", "fies"},
	{"gy", "gies"},
	{"hy", "hies"},
	{"io", "ios"},
	{"jy", "jies"},
	{"ky", "kies"},
	{"lf", "lves"},
	{"ly", "lies"},
	{"my", "mies"},
	{"ny", "nies"},
	{"py", "pies"},
	{"qy", "qies"},
	{"rf", "rves"},
	{"ry", "ries"},
	{"sh", "shes"},
	{"ss", "sses"},
	{"sy", "sies"},
	{"ty", "ties"},
	{"tz", "tzes"},
	{"va", "vae"},
	{"vy", "vies"},
	{"wy", "wies"},
	{"xy", "xies"},
	{"zy", "zies"},
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
