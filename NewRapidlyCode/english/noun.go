package english

// PoS from https://en.wikipedia.org/wiki/Part_of_speech
type PoS int

const (
	Unknown PoS = 0
	Noun    PoS = 1 << iota // 名词
	Verb                    // 动词
	Adj                     // 形容词
	Adv                     // 副词
	Prep                    // 介词 ( at / in / on / to / above / over / below / under )
	Conj                    // 连词
	AuxVerb                 // 助动词
	Pron                    // 代名词

	Ad = Adj | Adv
)

type Word struct {
	source string
	parent *Word // 根词汇
	Type   PoS
}
