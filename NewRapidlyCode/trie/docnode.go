package trie

import (
	"strings"
)

type Type int

const (
	Compound Type = iota
	Word
	Img
)

type DocNode struct {
	Value string
	Index int
	Type  Type
}

func NewDoc(doc string) []DocNode {
	docNode := make([]DocNode, 0, 32)
	for part, i := t.segmenter(doc, 0); part != ""; part, i = t.segmenter(doc, i) {
		docNode = append(docNode, DocNode{Value: part, Index: i - 1 - len(part)})
	}

}

func Segmenter(str string, start int) (segment string, next int) {
	if len(str) == 0 || start < 0 || start > len(str)-1 {
		return "", -1
	}
	end := strings.IndexRune(str[start+1:], ' ') // next '/' after 0th rune
	if end == -1 {
		return str[start:], -1
	}
	return str[start : start+end+1], start + end + 1
}
