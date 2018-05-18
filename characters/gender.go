package characters

import "strings"

type CGender byte

const (
	ENBY CGender = iota
	MALE
	FEMALE
)

var (
	Spouse            = [...]string{"spouse", "husband", "wife"}
	Child             = [...]string{"child", "son", "daughter"}
	GenderedPerson    = [...]string{"enby", "male", "female"}
	PronounAbsolute   = [...]string{"theirs", "his", "hers"}
	PronounSubject    = [...]string{"they", "he", "she"}
	PronounObject     = [...]string{"them", "him", "her"}
	PronounPossessive = [...]string{"their", "his", "her"}
	PronounReflexive  = [...]string{"themself", "himself", "herself"}
)

func GenderFmt(personOneGender, personTwoGender CGender, str string) string {
	ret := str
	ret = strings.Replace(ret, "%P", GenderedPerson[personOneGender], -1)
	ret = strings.Replace(ret, "%wP", GenderedPerson[personTwoGender], -1)
	ret = strings.Replace(ret, "%S", Spouse[personOneGender], -1)
	ret = strings.Replace(ret, "%wS", Spouse[personTwoGender], -1)
	ret = strings.Replace(ret, "%c", Child[personOneGender], -1)
	ret = strings.Replace(ret, "%wc", Child[personTwoGender], -1)
	ret = strings.Replace(ret, "%a", PronounAbsolute[personOneGender], -1)
	ret = strings.Replace(ret, "%wa", PronounAbsolute[personTwoGender], -1)
	ret = strings.Replace(ret, "%s", PronounSubject[personOneGender], -1)
	ret = strings.Replace(ret, "%ws", PronounSubject[personTwoGender], -1)
	ret = strings.Replace(ret, "%o", PronounObject[personOneGender], -1)
	ret = strings.Replace(ret, "%wo", PronounObject[personTwoGender], -1)
	ret = strings.Replace(ret, "%p", PronounPossessive[personOneGender], -1)
	ret = strings.Replace(ret, "%wp", PronounPossessive[personTwoGender], -1)
	ret = strings.Replace(ret, "%r", PronounReflexive[personOneGender], -1)
	ret = strings.Replace(ret, "%wr", PronounReflexive[personTwoGender], -1)
	return ret
}
