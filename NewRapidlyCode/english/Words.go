package english

var Words = map[string]Word{
	// 辅助动词
	"are": {Type: AuxVerb},

	"news":     {Type: Noun},
	"new":      {Type: Adj | Adv | Noun},
	"events":   {Type: Noun},
	"and":      {Type: Conj},
	"major":    {Type: Noun | Adj | Verb},
	"system":   {Type: Noun},
	"changes":  {Type: Verb | Noun},
	"arriving": {Type: Verb}, // 需要现在分词
	"this":     {Type: Pron},
}

func IsOneType(word Word) bool {
	if word.Type == Noun || word.Type == Verb ||
		word.Type == Adj || word.Type == Adv ||
		word.Type == Prep || word.Type == Conj {
		return true
	}
	return false
}
