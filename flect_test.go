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
	singular string
	plural   string
}

var singlePluralAssertions = []dict{
	{"", ""},
	{"Car", "Cars"},
	{"Boy", "Boys"},
	{"GoodBoy", "GoodBoys"},
	{"Axis", "Axes"},
	{"Child", "Children"},
	{"GoodChild", "GoodChildren"},
	{"SmartPerson", "SmartPeople"},
	{"SuperbOx", "SuperbOxen"},
	{"WildOx", "WildOxen"},
	{"wild_ox", "wild_oxen"},
	{"ability", "abilities"},
	{"address", "addresses"},
	{"agency", "agencies"},
	{"alias", "aliases"},
	{"analysis", "analyses"},
	{"archive", "archives"},
	{"axis", "axes"},
	{"basis", "bases"},
	{"belief", "beliefs"},
	{"blitz", "blitzes"},
	{"box", "boxes"},
	{"boy", "boys"},
	{"buffalo", "buffaloes"},
	{"bus", "buses"},
	{"cactus", "cacti"},
	{"case", "cases"},
	{"cat", "cats"},
	{"category", "categories"},
	{"chef", "chefs"},
	{"chief", "chiefs"},
	{"child", "children"},
	{"circus", "circuses"},
	{"city", "cities"},
	{"collapse", "collapses"},
	{"comment", "comments"},
	{"crisis", "crises"},
	{"criterion", "criteria"},
	{"custom_field", "custom_fields"},
	{"database", "databases"},
	{"datum", "data"},
	{"day", "days"},
	{"dear", "dears"},
	{"deer", "deer"},
	{"device", "devices"},
	{"diagnosis_a", "diagnosis_as"},
	{"diagnosis", "diagnoses"},
	{"dwarf", "dwarves"},
	{"eclipse", "eclipses"},
	{"edge", "edges"},
	{"elf", "elves"},
	{"ellipsis", "ellipses"},
	{"equipment", "equipment"},
	{"experience", "experiences"},
	{"fez", "fezzes"},
	{"field", "fields"},
	{"fish", "fish"},
	{"fix", "fixes"},
	{"fleet", "fleets"},
	{"focus", "foci"},
	{"foobar", "foobars"},
	{"foot", "feet"},
	{"fox", "foxes"},
	{"funky jeans", "funky jeans"},
	{"fuse", "fuses"},
	{"glimpse", "glimpses"},
	{"goose", "geese"},
	{"great_person", "great_people"},
	{"half", "halves"},
	{"halo", "halos"},
	{"horse", "horses"},
	{"house", "houses"},
	{"human", "humans"},
	{"index", "indices"},
	{"information", "information"},
	{"jeans", "jeans"},
	{"lapse", "lapses"},
	{"louse", "lice"},
	{"lunch", "lunches"},
	{"marsh", "marshes"},
	{"matrix", "matrices"},
	{"medium", "media"},
	{"multimedia", "multimedia"},
	{"mouse", "mice"},
	{"move", "moves"},
	{"movie", "movies"},
	{"news", "news"},
	{"newsletter", "newsletters"},
	{"node_child", "node_children"},
	{"octopus", "octopi"},
	{"ovum", "ova"},
	{"ox", "oxen"},
	{"payment_information", "payment_information"},
	{"pepsi", "pepsis"},
	{"person", "people"},
	{"perspective", "perspectives"},
	{"phenomenon", "phenomena"},
	{"photo", "photos"},
	{"piano", "pianos"},
	{"plus", "pluses"},
	{"portfolio", "portfolios"},
	{"potato", "potatoes"},
	{"prize", "prizes"},
	{"process", "processes"},
	{"prometheus", "prometheuses"},
	{"psi", "psis"},
	{"puppy", "puppies"},
	{"query", "queries"},
	{"quiz", "quizzes"},
	{"ray", "rays"},
	{"rice", "rice"},
	{"roof", "roofs"},
	{"safe", "saves"},
	{"salesperson", "salespeople"},
	{"search", "searches"},
	{"series", "series"},
	{"service", "services"},
	{"sheep", "sheep"},
	{"shoe", "shoes"},
	{"species", "species"},
	{"spokesman", "spokesmen"},
	{"sportsEquipment", "sportsEquipment"},
	{"spouse", "spouses"},
	{"stack", "stacks"},
	{"stadium", "stadia"},
	{"status_code", "status_codes"},
	{"status", "statuses"},
	{"Status", "Statuses"},
	{"switch", "switches"},
	{"tax", "taxes"},
	{"testis", "testes"},
	{"tomato", "tomatoes"},
	{"tooth", "teeth"},
	{"truss", "trusses"},
	{"user", "users"},
	{"user_custom_field", "user_custom_fields"},
	{"vedalia", "vedalias"},
	{"virus", "viri"},
	{"wife", "wives"},
	{"wish", "wishes"},
	{"wolf", "wolves"},
	{"woman", "women"},
}

func init() {
	for k, v := range singleToPlural {
		singlePluralAssertions = append(singlePluralAssertions, dict{k, v})

		// add some variations
		// singlePluralAssertions = append(singlePluralAssertions, tt{strings.ToUpper(k), v})
		// singlePluralAssertions = append(singlePluralAssertions, tt{strings.ToLower(k), v})
		// for i, x := range k {
		// 	n := k[:i] + strings.ToLower(string(x)) + k[i+1:]
		// 	singlePluralAssertions = append(singlePluralAssertions, tt{n, v})
		//
		// 	n = k[:i] + strings.ToUpper(string(x)) + k[i+1:]
		// 	singlePluralAssertions = append(singlePluralAssertions, tt{n, v})
		// }
	}
}
