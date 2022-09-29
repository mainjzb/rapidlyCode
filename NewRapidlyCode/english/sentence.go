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

	words := strings.Split(sentence, " ")

	// 判断首个单词是否应该为小写
	firstWord := strings.ToLower(words[0])
	_, ok := Words[firstWord]
	if ok {
		words[0] = firstWord
	}

	types := make([]WordType, len(words))

	for i := 0; i < len(words); i++ {
		types[i].Word = words[i]
		if IsOneType(Words[words[i]]) {
			types[i].Real = Words[words[i]].Type
		}
	}

	return types
}
