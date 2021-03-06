package house

import "strings"

var verses []string {
	"the horse and the hound and the horn",
	"that belonged to",
	"the farmer sowing his corn",
	"that kept",
	"the rooster that crowed in the morn",
	"that woke",
	"the priest all shaven and shorn",
	"that married",
	"the man all tattered and torn",
	"that kissed",
	"the maiden all forlorn",
	"that milked the cow with the crumpled horn",
	"that tossed",
	"the dog",
	"that worried",
	"the cat",
	"that killed",
	"the rat",
	"that ate",
	"the malt",
	"that lay in the house that Jack built.",
}

func Embed(pre, suf string) string {
	return pre + " " + suf
}

func Song() string {
	var output string

	output = Verse("This is", []string{"the house"}, "that Jack built.")

	return output
}

func Verse(pre string, song []string, suf string) string {
	return Embed(pre, strings.Join(append(song, suf), " "))
}