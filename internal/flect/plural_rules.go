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

type Word struct {
	Singular       string
	Plural         string
	Alternative    string
	Unidirectional bool // plural to singular is not possible (or bad)
	Uncountable    bool
	Exact          bool
}

// dictionary is the main table for singularize and pluralize.
// All words in the dictionary will be added to singleToPlural, pluralToSingle
// and singlePluralAssertions by init() functions.
var Dictionary = []Word{
	// identicals https://en.wikipedia.org/wiki/English_plurals#Nouns_with_identical_singular_and_plural
	{Singular: "aircraft", Plural: "aircraft"},
	{Singular: "beef", Plural: "beef", Alternative: "beefs"},
	{Singular: "bison", Plural: "bison"},
	{Singular: "blues", Plural: "blues", Unidirectional: true},
	{Singular: "chassis", Plural: "chassis"},
	{Singular: "deer", Plural: "deer"},
	{Singular: "fish", Plural: "fish", Alternative: "fishes"},
	{Singular: "moose", Plural: "moose"},
	{Singular: "police", Plural: "police"},
	{Singular: "salmon", Plural: "salmon", Alternative: "salmons"},
	{Singular: "series", Plural: "series"},
	{Singular: "sheep", Plural: "sheep"},
	{Singular: "shrimp", Plural: "shrimp", Alternative: "shrimps"},
	{Singular: "species", Plural: "species"},
	{Singular: "swine", Plural: "swine", Alternative: "swines"},
	{Singular: "trout", Plural: "trout", Alternative: "trouts"},
	{Singular: "tuna", Plural: "tuna", Alternative: "tunas"},
	{Singular: "you", Plural: "you"},
	// -en https://en.wikipedia.org/wiki/English_plurals#Plurals_in_-(e)n
	{Singular: "child", Plural: "children"},
	{Singular: "ox", Plural: "oxen", Exact: true},
	// apophonic https://en.wikipedia.org/wiki/English_plurals#Apophonic_plurals
	{Singular: "foot", Plural: "feet"},
	{Singular: "goose", Plural: "geese"},
	{Singular: "man", Plural: "men"},
	{Singular: "human", Plural: "humans"}, // not humen
	{Singular: "louse", Plural: "lice", Exact: true},
	{Singular: "mouse", Plural: "mice"},
	{Singular: "tooth", Plural: "teeth"},
	{Singular: "woman", Plural: "women"},
	// misc https://en.wikipedia.org/wiki/English_plurals#Miscellaneous_irregular_plurals
	{Singular: "die", Plural: "dice", Exact: true},
	{Singular: "person", Plural: "people"},

	// Words from French that end in -u add an x; in addition to eau to eaux rule
	{Singular: "adieu", Plural: "adieux", Alternative: "adieus"},
	{Singular: "fabliau", Plural: "fabliaux"},
	{Singular: "bureau", Plural: "bureaus", Alternative: "bureaux"}, // popular

	// Words from Greek that end in -on change -on to -a; in addition to hedron rule
	{Singular: "criterion", Plural: "criteria"},
	{Singular: "ganglion", Plural: "ganglia", Alternative: "ganglions"},
	{Singular: "lexicon", Plural: "lexica", Alternative: "lexicons"},
	{Singular: "mitochondrion", Plural: "mitochondria", Alternative: "mitochondrions"},
	{Singular: "noumenon", Plural: "noumena"},
	{Singular: "phenomenon", Plural: "phenomena"},
	{Singular: "taxon", Plural: "taxa"},

	// Words from Latin that end in -um change -um to -a; in addition to some rules
	{Singular: "media", Plural: "media"}, // popular case: media -> media
	{Singular: "medium", Plural: "media", Alternative: "mediums", Unidirectional: true},
	{Singular: "stadium", Plural: "stadiums", Alternative: "stadia"},
	{Singular: "aquarium", Plural: "aquaria", Alternative: "aquariums"},
	{Singular: "auditorium", Plural: "auditoria", Alternative: "auditoriums"},
	{Singular: "symposium", Plural: "symposia", Alternative: "symposiums"},
	{Singular: "curriculum", Plural: "curriculums", Alternative: "curricula"}, // ulum
	{Singular: "quota", Plural: "quotas"},

	// Words from Latin that end in -us change -us to -i or -era
	{Singular: "alumnus", Plural: "alumni", Alternative: "alumnuses"}, // -i
	{Singular: "bacillus", Plural: "bacilli"},
	{Singular: "cactus", Plural: "cacti", Alternative: "cactuses"},
	{Singular: "coccus", Plural: "cocci"},
	{Singular: "focus", Plural: "foci", Alternative: "focuses"},
	{Singular: "locus", Plural: "loci", Alternative: "locuses"},
	{Singular: "nucleus", Plural: "nuclei", Alternative: "nucleuses"},
	{Singular: "octopus", Plural: "octupuses", Alternative: "octopi"},
	{Singular: "radius", Plural: "radii", Alternative: "radiuses"},
	{Singular: "syllabus", Plural: "syllabi"},
	{Singular: "corpus", Plural: "corpora", Alternative: "corpuses"}, // -ra
	{Singular: "genus", Plural: "genera"},

	// Words from Latin that end in -a change -a to -ae
	{Singular: "alumna", Plural: "alumnae"},
	{Singular: "vertebra", Plural: "vertebrae"},
	{Singular: "differentia", Plural: "differentiae"}, // -tia
	{Singular: "minutia", Plural: "minutiae"},
	{Singular: "vita", Plural: "vitae"},   // -ita
	{Singular: "larva", Plural: "larvae"}, // -va
	{Singular: "postcava", Plural: "postcavae"},
	{Singular: "praecava", Plural: "praecavae"},
	{Singular: "uva", Plural: "uvae"},

	// Words from Latin that end in -ex change -ex to -ices
	{Singular: "apex", Plural: "apices", Alternative: "apexes"},
	{Singular: "codex", Plural: "codices", Alternative: "codexes"},
	{Singular: "index", Plural: "indices", Alternative: "indexes"},
	{Singular: "latex", Plural: "latices", Alternative: "latexes"},
	{Singular: "vertex", Plural: "vertices", Alternative: "vertexes"},
	{Singular: "vortex", Plural: "vortices", Alternative: "vortexes"},

	// Words from Latin that end in -ix change -ix to -ices (eg, matrix becomes matrices)
	{Singular: "appendix", Plural: "appendices", Alternative: "appendixes"},
	{Singular: "radix", Plural: "radices", Alternative: "radixes"},
	{Singular: "helix", Plural: "helices", Alternative: "helixes"},

	// Words from Latin that end in -is change -is to -es
	{Singular: "axis", Plural: "axes", Exact: true},
	{Singular: "crisis", Plural: "crises"},
	{Singular: "ellipsis", Plural: "ellipses", Unidirectional: true}, // ellipse
	{Singular: "genesis", Plural: "geneses"},
	{Singular: "oasis", Plural: "oases"},
	{Singular: "thesis", Plural: "theses"},
	{Singular: "testis", Plural: "testes"},
	{Singular: "base", Plural: "bases"}, // popular case
	{Singular: "basis", Plural: "bases", Unidirectional: true},

	{Singular: "alias", Plural: "aliases", Exact: true}, // no alia, no aliasis
	{Singular: "vedalia", Plural: "vedalias"},           // no vedalium, no vedaliases

	// Words that end in -ch, -o, -s, -sh, -x, -z (can be conflict with the others)
	{Singular: "use", Plural: "uses", Exact: true}, // us vs use
	{Singular: "abuse", Plural: "abuses"},
	{Singular: "cause", Plural: "causes"},
	{Singular: "clause", Plural: "clauses"},
	{Singular: "cruse", Plural: "cruses"},
	{Singular: "excuse", Plural: "excuses"},
	{Singular: "fuse", Plural: "fuses"},
	{Singular: "house", Plural: "houses"},
	{Singular: "misuse", Plural: "misuses"},
	{Singular: "muse", Plural: "muses"},
	{Singular: "pause", Plural: "pauses"},
	{Singular: "ache", Plural: "aches"},
	{Singular: "topaz", Plural: "topazes"},
	{Singular: "buffalo", Plural: "buffaloes", Alternative: "buffalos"},
	{Singular: "potato", Plural: "potatoes"},
	{Singular: "tomato", Plural: "tomatoes"},

	// uncountables
	{Singular: "equipment", Uncountable: true},
	{Singular: "information", Uncountable: true},
	{Singular: "jeans", Uncountable: true},
	{Singular: "money", Uncountable: true},
	{Singular: "news", Uncountable: true},
	{Singular: "rice", Uncountable: true},

	// exceptions: -f to -ves, not -fe
	{Singular: "dwarf", Plural: "dwarfs", Alternative: "dwarves"},
	{Singular: "hoof", Plural: "hoofs", Alternative: "hooves"},
	{Singular: "thief", Plural: "thieves"},
	// exceptions: instead of -f(e) to -ves
	{Singular: "chive", Plural: "chives"},
	{Singular: "hive", Plural: "hives"},
	{Singular: "move", Plural: "moves"},

	// exceptions: instead of -y to -ies
	{Singular: "movie", Plural: "movies"},
	{Singular: "cookie", Plural: "cookies"},

	// exceptions: instead of -um to -a
	{Singular: "pretorium", Plural: "pretoriums"},
	{Singular: "agenda", Plural: "agendas"}, // instead of plural of agendum
	// exceptions: instead of -um to -a (chemical element names)

	// Words from Latin that end in -a change -a to -ae
	{Singular: "formula", Plural: "formulas", Alternative: "formulae"}, // also -um/-a

	// exceptions: instead of -o to -oes
	{Singular: "shoe", Plural: "shoes"},
	{Singular: "toe", Plural: "toes", Exact: true},
	{Singular: "graffiti", Plural: "graffiti"},

	// abbreviations
	{Singular: "ID", Plural: "IDs", Exact: true},
}

// singleToPlural is the highest priority map for Pluralize().
// singularToPluralSuffixList is used to build pluralRules for suffixes and
// compound words.
var singleToPlural = map[string]string{}

// pluralToSingle is the highest priority map for Singularize().
// singularToPluralSuffixList is used to build singularRules for suffixes and
// compound words.
var pluralToSingle = map[string]string{}

// NOTE: This map should not be built as reverse map of singleToPlural since
// there are words that has the same plurals.

// build singleToPlural and pluralToSingle with dictionary
func init() {
	for _, wd := range Dictionary {
		if singleToPlural[wd.Singular] != "" {
			panic(fmt.Errorf("map singleToPlural already has an entry for %s", wd.Singular))
		}

		if wd.Uncountable && wd.Plural == "" {
			wd.Plural = wd.Singular
		}

		if wd.Plural == "" {
			panic(fmt.Errorf("plural for %s is not provided", wd.Singular))
		}

		singleToPlural[wd.Singular] = wd.Plural

		if !wd.Unidirectional {
			if pluralToSingle[wd.Plural] != "" {
				panic(fmt.Errorf("map pluralToSingle already has an entry for %s", wd.Plural))
			}
			pluralToSingle[wd.Plural] = wd.Singular

			if wd.Alternative != "" {
				if pluralToSingle[wd.Alternative] != "" {
					panic(fmt.Errorf("map pluralToSingle already has an entry for %s", wd.Alternative))
				}
				pluralToSingle[wd.Alternative] = wd.Singular
			}
		}
	}
}

type singularToPluralSuffix struct {
	singular string
	plural   string
}

// singularToPluralSuffixList is a list of "bidirectional" suffix rules for
// the irregular plurals follow such rules.
//
// NOTE: IMPORTANT! The order of items in this list is the rule priority, not
// alphabet order. The first match will be used to inflect.
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

	// Words from Latin that end in -a change -a to -ae; before -on to -a and -um to -a
	{"bula", "bulae"},
	{"dula", "bulae"},
	{"lula", "bulae"},
	{"nula", "bulae"},
	{"vula", "bulae"},

	// Words from Greek that end in -on change -on to -a (eg, polyhedron becomes polyhedra)
	// https://en.wiktionary.org/wiki/Category:English_irregular_plurals_ending_in_"-a"
	{"hedron", "hedra"},

	// Words from Latin that end in -um change -um to -a (eg, minimum becomes minima)
	// https://en.wiktionary.org/wiki/Category:English_irregular_plurals_ending_in_"-a"
	{"ium", "ia"}, // some exceptions especially chemical element names
	{"seum", "seums"},
	{"eum", "ea"},
	{"oum", "oa"},
	{"stracum", "straca"},
	{"dum", "da"},
	{"elum", "ela"},
	{"ilum", "ila"},
	{"olum", "ola"},
	{"ulum", "ula"},
	{"llum", "lla"},
	{"ylum", "yla"},
	{"imum", "ima"},
	{"ernum", "erna"},
	{"gnum", "gna"},
	{"brum", "bra"},
	{"crum", "cra"},
	{"terum", "tera"},
	{"serum", "sera"},
	{"trum", "tra"},
	{"antum", "anta"},
	{"atum", "ata"},
	{"entum", "enta"},
	{"etum", "eta"},
	{"itum", "ita"},
	{"otum", "ota"},
	{"utum", "uta"},
	{"ctum", "cta"},
	{"ovum", "ova"},

	// Words from Latin that end in -us change -us to -i or -era
	// not easy to make a simple rule. just add them all to the dictionary

	// Words from Latin that end in -ex change -ex to -ices (eg, vortex becomes vortices)
	// Words from Latin that end in -ix change -ix to -ices (eg, matrix becomes matrices)
	//    for example, -dix, -dex, and -dice will have the same plural form so
	//    making a simple rule is not possible for them
	{"trix", "trices"}, // ignore a few words end in trice

	// Words from Latin that end in -is change -is to -es (eg, thesis becomes theses)
	// -sis and -se has the same plural -ses so making a rule is not easy too.
	{"iasis", "iases"},
	{"mesis", "meses"},
	{"kinesis", "kineses"},
	{"resis", "reses"},
	{"gnosis", "gnoses"}, // e.g. diagnosis
	{"opsis", "opses"},   // e.g. synopsis
	{"ysis", "yses"},     // e.g. analysis

	// Words that end in -ch, -o, -s, -sh, -x, -z
	{"ouse", "ouses"},
	{"lause", "lauses"},
	{"us", "uses"}, // use/uses is in the dictionary

	{"ch", "ches"},
	{"io", "ios"},
	{"sh", "shes"},
	{"ss", "sses"},
	{"ez", "ezzes"},
	{"iz", "izzes"},
	{"tz", "tzes"},
	{"zz", "zzes"},
	{"ano", "anos"},
	{"lo", "los"},
	{"to", "tos"},
	{"oo", "oos"},
	{"o", "oes"},
	{"x", "xes"},

	// for abbreviations
	{"S", "Ses"},
}

func init() {
	nSuffix := len(singularToPluralSuffixList)
	nDict := len(Dictionary)

	newPlural := make([]rule, 0, 2*(nDict+nSuffix))
	newSingular := make([]rule, 0, 3*(nDict+nSuffix)) // 3× to account for Alternative entries

	// Dict compound rules (higher priority).
	// Original code prepended in forward order → last entry ended at index 0.
	// Replicate with append by iterating backward so last entry is appended first.
	for i := nDict - 1; i >= 0; i-- {
		wd := Dictionary[i]
		if wd.Exact {
			continue
		}
		if wd.Uncountable && wd.Plural == "" {
			wd.Plural = wd.Singular
		}
		newPlural = append(newPlural,
			rule{suffix: wd.Plural, fn: noop},
			rule{suffix: wd.Singular, fn: simpleRuleFunc(wd.Singular, wd.Plural)},
		)
		if !wd.Unidirectional {
			// Alternative was inserted last (highest priority among the two),
			// so it must be appended first here.
			if wd.Alternative != "" {
				newSingular = append(newSingular,
					rule{suffix: wd.Singular, fn: noop},
					rule{suffix: wd.Alternative, fn: simpleRuleFunc(wd.Alternative, wd.Singular)},
				)
			}
			newSingular = append(newSingular,
				rule{suffix: wd.Singular, fn: noop},
				rule{suffix: wd.Plural, fn: simpleRuleFunc(wd.Plural, wd.Singular)},
			)
		}
	}

	// Suffix rules (lower priority).
	// Original code iterated backward and prepended → element 0 ended at index 0.
	// Replicate with append by iterating forward.
	for _, s := range singularToPluralSuffixList {
		newPlural = append(newPlural,
			rule{suffix: s.plural, fn: noop},
			rule{suffix: s.singular, fn: simpleRuleFunc(s.singular, s.plural)},
		)
		newSingular = append(newSingular,
			rule{suffix: s.singular, fn: noop},
			rule{suffix: s.plural, fn: simpleRuleFunc(s.plural, s.singular)},
		)
	}

	pluralRules = newPlural
	singularRules = newSingular
}
