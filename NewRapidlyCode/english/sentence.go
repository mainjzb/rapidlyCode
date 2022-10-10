package english

import (
	"strings"
	"unicode"
)

func SplitParagraph(paragraph string) []string {
	sentences := make([]string, 0, 16)

	index := 0
	start := 0
	end := 0
	searchStart := 0
	for {
		index = strings.IndexAny(paragraph[searchStart:], ".!?")
		if index == -1 {
			sentences = append(sentences, paragraph[searchStart:])
			break
		}

		// 下一个单词以大写开头
		nextContent := strings.TrimSpace(paragraph[searchStart+index+1:])
		if len(nextContent) == 0 {
			sentences = append(sentences, strings.TrimSpace(paragraph[end:searchStart+index+1]))
			break
		}
		if unicode.IsLower(rune(nextContent[0])) {
			searchStart = searchStart + index + 1
			continue
		}

		start = end
		end = searchStart + index + 1
		add := paragraph[start:end]
		sentences = append(sentences, strings.TrimSpace(add))

		searchStart = end
	}

	return sentences
}

type WordType struct {
	Word    string
	Suppose PoS
	Real    PoS
}

func Segment(sentence string) []WordType {
	// New events and major system changes are arriving this update!
	sentence = strings.TrimRight(sentence, ".!?")
	words := strings.Split(sentence, " ")

	// 判断首个单词是否应该为小写
	firstWord := strings.ToLower(words[0])
	_, ok := Words[firstWord]
	if ok {
		words[0] = firstWord
	}
	types := make([]WordType, len(words))

	RealNum := 0
	for i := 0; i < len(words); i++ {
		types[i].Word = words[i]
		types[i].Suppose = Words[words[i]].Type
		if IsOneType(Words[words[i]]) {
			types[i].Real = Words[words[i]].Type
			RealNum++
			if types[i].Real == Pron && i+1 < len(words) && IsNoun(Words[words[i+1]].Type) {
				types[i+1].Word = words[i+1]
				types[i+1].Real = Noun
				i++
				RealNum++
			}
			continue
		}
	}

	startNode := &Node{Type: "^"}
	prev := startNode
	for i := 0; i < len(types); {
		node := &Node{prev: prev}
		prev.next = node
		if types[i].Real == Unknown {
			list := make([]WordType, 0)
			list = append(list, types[i])
			nextType := types[i]
			for j := 0; i+j < len(types); j++ {
				t := types[i+j].Real
				if t == Conj || t == AuxVerb {
					nextType = types[i+j]
					break
				}
				list = append(list, types[i+j])
			}
			judge(prev, list, nextType)
		}
	}

	// node := &Node{}
	// // 识别率大于60 开始判断句型
	// if float64(RealNum)/float64(len(words)) > 0.6 {
	// 	for i := 0; i < len(types); i++ {
	// 		if i == 0 && IsOnlyCompoundNoun(types[i].Suppose) {
	// 			node.
	// 		}
	// 	}
	// }

	return types
}

func judge(pre *Node, list []WordType, next WordType) {

}
