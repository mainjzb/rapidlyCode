package english

// PoS from https://en.wikipedia.org/wiki/Part_of_speech
type PoS int
type VerbsType int
type Word struct {
	source string
	parent *Word // 根词汇
	Type   PoS
}

const (
	Unknown PoS = 0
	Noun    PoS = 1 << iota // 名词
	Adj                     // 形容词
	Adv                     // 副词
	Prep                    // 介词 ( at / in / on / to / above / over / below / under )
	Conj                    // 连词
	AuxVerb                 // 助动词
	Pron                    // 代名词

	VerbBase                       // 动词
	VerbPastTense                  // 动词过去式
	VerbPastParticiple             // 动词过去分词
	VerbPresentParticiple          // 动词现在分词
	VerbThirdPersonSingularPresent // 动词第三人称单数

	Ad   = Adj | Adv
	Verb = VerbBase | VerbPastTense | VerbPastParticiple | VerbPresentParticiple | VerbThirdPersonSingularPresent
)

var Words = map[string]Word{
	// 辅助动词
	"are": {Type: AuxVerb},

	"and":  {Type: Conj},
	"this": {Type: Pron},

	"news":   {Type: Noun},
	"events": {Type: Noun},
	"system": {Type: Noun},

	"new": {Type: Adj | Adv | Noun},

	"update":   {Type: Noun | VerbBase},
	"major":    {Type: Noun | Adj | VerbBase},
	"changes":  {Type: VerbThirdPersonSingularPresent | Noun},
	"arriving": {Type: VerbPresentParticiple}, // 需要现在分词
}

func IsOneType(word Word) bool {
	if word.Type == Noun ||
		word.Type == Adj || word.Type == Adv ||
		word.Type == Prep || word.Type == Conj ||
		word.Type == AuxVerb || word.Type == Pron {
		return true
	}
	v := word.Type & Verb
	if v > 0 && word.Type == v {
		return true
	}
	return false
}

func IsVerb(t PoS) bool {
	if t&Verb > 0 {
		return true
	}
	return false
}

func IsNoun(t PoS) bool {
	if t&Noun > 0 {
		return true
	}
	return false
}

func IsOnlyCompoundNoun(t PoS) bool {
	if t == Adj|Adv|Noun {
		return true
	}
	if t == Noun {
		return true
	}

	return false
}
