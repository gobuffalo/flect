package flect

// source for grammar rules: https://www.grammarly.com/blog/plural-nouns/
import (
	"strings"
)

func Pluralize(s string) string {
	return New(s).Pluralize()
}

func (i Ident) Pluralize() string {
	if len(i.parts) == 0 {
		return ""
	}
	li := len(i.parts) - 1
	last := i.parts[li]
	if _, ok := pluralToSingle[strings.ToLower(last)]; ok {
		return i.original
	}
	parts := i.parts[:li]
	if p, ok := singleToPlural[strings.ToLower(last)]; ok {
		parts = append(parts, p)
		return strings.Join(parts, " ")
	}

	for _, r := range pluralRules {
		if strings.HasSuffix(last, r.suffix) {
			parts = append(parts, r.fn(last))
			return strings.Join(parts, " ")
		}
	}

	//To make regular nouns plural, add ‑s to the end.
	if !strings.HasSuffix(last, "s") {
		last += "s"
	}
	parts = append(parts, last)
	return strings.Join(parts, " ")
}

//  If the singular noun ends in ‑s, -ss, -sh, -ch, -x, or -z, add ‑es to the end to make it plural.
func esPlural(s string) string {
	return s + "es"
}

//  If the noun ends with ‑f or ‑fe, the f is often changed to ‑ve before adding the -s to form the plural version.
func fPlural(s string) string {
	if len(s) > 1 {
		c := s[len(s)-2]
		if isVowel(rune(c)) {
			return s + "s"
		}
	}
	s = strings.TrimSuffix(s, "f")
	return s + "ves"
}

func fePlural(s string) string {
	return fPlural(strings.TrimSuffix(s, "e"))
}

// If a singular noun ends in ‑y and the letter before the -y is a consonant, change the ending to ‑ies to make the noun plural.
func yPlural(s string) string {
	if len(s) > 1 {
		c := s[len(s)-2]
		if isVowel(rune(c)) {
			return s + "s"
		}
	}
	s = strings.TrimSuffix(s, "y")
	return s + "ies"
}

// If the singular noun ends in ‑us, the plural ending is frequently ‑i.
func usPlural(s string) string {
	s = strings.TrimSuffix(s, "us")
	return s + "i"
}

// If the singular noun ends in ‑is, the plural ending is ‑es.
func isPlural(s string) string {
	s = strings.TrimSuffix(s, "is")
	return s + "es"
}

// If the singular noun ends in ‑on, the plural ending is ‑a.
func onPlural(s string) string {
	s = strings.TrimSuffix(s, "on")
	return s + "a"
}

func umPlural(s string) string {
	s = strings.TrimSuffix(s, "um")
	return s + "a"
}

var pluralRules = []rule{
	{"es", noop},
	{"ves", noop},
	{"ies", noop},
	{"i", noop},
	{"es", noop},
	{"a", noop},
	{"ss", esPlural},
	{"sh", esPlural},
	{"ch", esPlural},
	{"x", esPlural},
	{"z", esPlural},
	{"o", esPlural},
	{"f", fPlural},
	{"fe", fPlural},
	{"y", yPlural},
	{"us", usPlural},
	{"is", isPlural},
	{"on", onPlural},
	{"um", umPlural},
}

var singleToPlural = map[string]string{
	"matrix":      "matrices",
	"vertix":      "vertices",
	"index":       "indices",
	"mouse":       "mice",
	"louse":       "lice",
	"ress":        "resses",
	"ox":          "oxen",
	"quiz":        "quizzes",
	"series":      "series",
	"octopus":     "octopi",
	"person":      "people",
	"man":         "men",
	"child":       "children",
	"equipment":   "equipment",
	"information": "information",
	"rice":        "rice",
	"money":       "money",
	"species":     "species",
	"fish":        "fish",
	"sheep":       "sheep",
	"jeans":       "jeans",
	"police":      "police",
	"dear":        "dear",
	"goose":       "geese",
	"woman":       "women",
	"tooth":       "teeth",
	"foot":        "feet",
	"bus":         "busses",
	"fez":         "fezzes",
	"piano":       "pianos",
	"halo":        "halos",
	"photo":       "photos",
	"aircraft":    "aircraft",
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
	"basis":       "bases",
	"beau":        "beaus",
	"bison":       "bison",
	"bureau":      "bureaus",
	"château":     "châteaux",
	"codex":       "codices",
	"concerto":    "concertos",
	"corpus":      "corpora",
	"crisis":      "crises",
	"curriculum":  "curriculums",
	"deer":        "deer",
	"diagnosis":   "diagnoses",
	"die":         "dice",
	"dwarf":       "dwarves",
	"ellipsis":    "ellipses",
	"erratum":     "errata",
	"faux pas":    "faux pas",
	"focus":       "foci",
	"formula":     "formulas",
	"fungus":      "fungi",
	"genus":       "genera",
	"graffito":    "graffiti",
	"grouse":      "grouse",
	"half":        "halves",
	"hoof":        "hooves",
	"hypothesis":  "hypotheses",
	"larva":       "larvae",
	"libretto":    "librettos",
	"loaf":        "loaves",
	"locus":       "loci",
	"medium":      "mediums",
	"memorandum":  "memoranda",
	"minutia":     "minutiae",
	"moose":       "moose",
	"nebula":      "nebulae",
	"nucleus":     "nuclei",
	"oasis":       "oases",
	"offspring":   "offspring",
	"opus":        "opera",
	"ovum":        "ova",
	"parenthesis": "parentheses",
	"phenomenon":  "phenomena",
	"phylum":      "phyla",
	"prognosis":   "prognoses",
	"radius":      "radiuses",
	"referendum":  "referendums",
	"salmon":      "salmon",
	"shrimp":      "shrimp",
	"stimulus":    "stimuli",
	"stratum":     "strata",
	"swine":       "swine",
	"syllabus":    "syllabi",
	"symposium":   "symposiums",
	"synopsis":    "synopses",
	"tableau":     "tableaus",
	"thesis":      "theses",
	"thief":       "thieves",
	"trout":       "trout",
	"tuna":        "tuna",
	"vertebra":    "vertebrae",
	"vita":        "vitae",
	"vortex":      "vortices",
	"wharf":       "wharves",
	"wife":        "wives",
	"wolf":        "wolves",
	"datum":       "data",
}

var pluralToSingle = map[string]string{}

func init() {
	for k, v := range singleToPlural {
		pluralToSingle[v] = k
	}
}
