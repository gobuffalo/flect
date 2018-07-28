package flect

// source for grammar rules: https://www.grammarly.com/blog/plural-nouns/
import (
	"strings"
	"sync"
)

var singularMoot = &sync.RWMutex{}

func Singularize(s string) string {
	return New(s).Singularize()
}

func (i Ident) Singularize() string {
	s := i.original
	if len(s) == 0 {
		return ""
	}

	singularMoot.RLock()
	defer singularMoot.RUnlock()
	ls := strings.ToLower(s)
	if p, ok := pluralToSingle[ls]; ok {
		return p
	}
	if _, ok := singleToPlural[ls]; ok {
		return s
	}
	for _, r := range singularRules {
		if strings.HasSuffix(ls, r.suffix) {
			return r.fn(s)
		}
	}

	return s
}

var singularRules = []rule{}

func AddSingular(ext string, repl string) {
	singularMoot.Lock()
	defer singularMoot.Unlock()
	singularRules = append(singularRules, rule{
		suffix: ext,
		fn: func(s string) string {
			s = s[:len(s)-len(ext)]
			return s + repl
		},
	})

	singularRules = append(singularRules, rule{
		suffix: repl,
		fn: func(s string) string {
			return s
		},
	})
}

func init() {
	AddSingular("ria", "rion")
	AddSingular("news", "news")
	AddSingular("halves", "half")
	AddSingular("appendix", "appendix")
	AddSingular("zzes", "zz")
	AddSingular("ulas", "ula")
	AddSingular("psis", "pse")
	AddSingular("genus", "genera")
	AddSingular("phyla", "phylum")
	AddSingular("odice", "odex")
	AddSingular("oxen", "ox")
	AddSingular("ianos", "iano")
	AddSingular("ulus", "uli")
	AddSingular("mice", "mouse")
	AddSingular("ouses", "ouse")
	AddSingular("mni", "mnus")
	AddSingular("ocus", "oci")
	AddSingular("shoes", "shoe")
	AddSingular("oasis", "oasis")
	AddSingular("lice", "louse")
	AddSingular("men", "man")
	AddSingular("ta", "tum")
	AddSingular("ia", "ium")
	AddSingular("tives", "tive")
	AddSingular("ldren", "ld")
	AddSingular("people", "person")
	AddSingular("aves", "afe")
	AddSingular("uses", "us")
	AddSingular("bves", "bfe")
	AddSingular("cves", "cfe")
	AddSingular("dves", "dfe")
	AddSingular("eves", "efe")
	AddSingular("gves", "gfe")
	AddSingular("hves", "hfe")
	AddSingular("chives", "chive")
	AddSingular("ives", "ife")
	AddSingular("movies", "movie")
	AddSingular("jeans", "jeans")
	AddSingular("cesses", "cess")
	AddSingular("cess", "cess")
	AddSingular("acti", "actus")
	AddSingular("itzes", "itz")
	AddSingular("usses", "uss")
	AddSingular("uss", "uss")
	AddSingular("jves", "jfe")
	AddSingular("kves", "kfe")
	AddSingular("mves", "mfe")
	AddSingular("nves", "nfe")
	AddSingular("moves", "move")
	AddSingular("oves", "ofe")
	AddSingular("pves", "pfe")
	AddSingular("qves", "qfe")
	AddSingular("sves", "sfe")
	AddSingular("tves", "tfe")
	AddSingular("uves", "ufe")
	AddSingular("vves", "vfe")
	AddSingular("wves", "wfe")
	AddSingular("xves", "xfe")
	AddSingular("yves", "yfe")
	AddSingular("zves", "zfe")
	AddSingular("hives", "hive")
	AddSingular("lves", "lf")
	AddSingular("rves", "rf")
	AddSingular("quies", "quy")
	AddSingular("bies", "by")
	AddSingular("cies", "cy")
	AddSingular("dies", "dy")
	AddSingular("fies", "fy")
	AddSingular("gies", "gy")
	AddSingular("hies", "hy")
	AddSingular("jies", "jy")
	AddSingular("kies", "ky")
	AddSingular("lies", "ly")
	AddSingular("mies", "my")
	AddSingular("nies", "ny")
	AddSingular("pies", "py")
	AddSingular("qies", "qy")
	AddSingular("ries", "ry")
	AddSingular("sies", "sy")
	AddSingular("ties", "ty")
	AddSingular("vies", "vy")
	AddSingular("wies", "wy")
	AddSingular("xies", "xy")
	AddSingular("zies", "zy")
	AddSingular("xes", "x")
	AddSingular("ches", "ch")
	AddSingular("sses", "ss")
	AddSingular("shes", "sh")
	AddSingular("oes", "o")
	AddSingular("ress", "ress")
	AddSingular("iri", "irus")
	AddSingular("irus", "irus")
	AddSingular("tuses", "tus")
	AddSingular("tus", "tus")
	AddSingular("s", "")
	AddSingular("ss", "ss")
}
