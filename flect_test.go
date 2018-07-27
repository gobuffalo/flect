package flect

type tt struct {
	act string
	exp string
}

var singlePluralAssertions = []tt{
	{"", ""},
	{"user", "users"},
	{"cat", "cats"},
	{"truss", "trusses"},
	{"bus", "busses"},
	{"marsh", "marshes"},
	{"lunch", "lunches"},
	{"tax", "taxes"},
	{"blitz", "blitzes"},
	{"fez", "fezzes"},
	{"wolf", "wolves"},
	{"roof", "roofs"},
	{"belief", "beliefs"},
	{"chef", "chefs"},
	{"chief", "chiefs"},
	{"city", "cities"},
	{"puppy", "puppies"},
	{"ray", "rays"},
	{"boy", "boys"},
	{"potato", "potatoes"},
	{"tomato", "tomatoes"},
	{"photo", "photos"},
	{"piano", "pianos"},
	{"halo", "halos"},
	{"cactus", "cacti"},
	{"focus", "foci"},
	{"datum", "data"},
	{"analysis", "analyses"},
	{"ellipsis", "ellipses"},
	{"phenomenon", "phenomena"},
	{"criterion", "criteria"},
	{"sheep", "sheep"},
	{"series", "series"},
	{"species", "species"},
	{"dear", "dear"},
	{"child", "children"},
	{"goose", "geese"},
	{"man", "men"},
	{"woman", "women"},
	{"tooth", "teeth"},
	{"foot", "feet"},
	{"mouse", "mice"},
	{"person", "people"},
}

var pluralSingularAssertions = []tt{}

func init() {
	for k, v := range singleToPlural {
		singlePluralAssertions = append(singlePluralAssertions, tt{k, v})
	}

	for _, a := range singlePluralAssertions {
		pluralSingularAssertions = append(pluralSingularAssertions, tt{act: a.exp, exp: a.act})
	}
}
