/*
Package trie implements several types of performant Tries (e.g. rune-wise,
path-wise).

The implementations are optimized for Get performance and to allocate 0 bytes
of heap memory (i.e. garbage) per Get.

The Tries do not synchronize access (not thread-safe). A typical use case is
to perform Puts and Deletes upfront to populate the Trie, then perform Gets
very quickly.
*/
package trie

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type StreamNodeType int

const (
	Tag StreamNodeType = iota
	Word
	Punctuation
)

type StreamNode struct {
	Data       string
	Appending  string //  跟随在单词后面的空格，在复原文本时有用
	Type       StreamNodeType
	Prev, Next *StreamNode
}

func DocParse(content string) *StreamNode {
	root := &StreamNode{}
	strings.Trim(content, " \n\t\r")

	left := 0
	right := 0
	prev := root

	for {
		if left == len(content) {
			break
		}
		firstRune, size := utf8.DecodeRuneInString(content[left:])
		if size == 1 && firstRune == '<' {
			right = strings.IndexRune(content[left:], '>') + left + 1
			prev, right = AddStreamNode(content, left, right, prev, Tag)
		} else if size == 1 && (isPunctuation(byte(firstRune)) || firstRune == '\'') {
			right = left + 1
			for isPunctuation(content[right]) {
				right++
			}
			prev, right = AddStreamNode(content, left, right, prev, Punctuation)
		} else {
			for i := left + 1; i <= len(content); i++ {
				if (content[i] == '.' || content[i] == '\'') && i+1 < len(content) && isLetterOrNumber(content[i+1]) {
					continue
				}

				if i == len(content) || isPunctuation(content[i]) || content[i] == ' ' || content[i] == '<' {
					right = i
					break
				}
			}
			prev, right = AddStreamNode(content, left, right, prev, Word)
		}
		left = right
	}
	return root.Next
}

func AddStreamNode(content string, left int, right int, prev *StreamNode, t StreamNodeType) (*StreamNode, int) {
	node := &StreamNode{
		Data: content[left:right],
		Type: t,
		Prev: prev,
	}

	// find space to appending
	for i := right; i < len(content); i++ {
		if content[i] != ' ' {
			right = i
			break
		}
		node.Appending += string(content[i])
	}

	prev.Next = node
	return node, right
}

func isLetterOrNumber(c byte) bool {
	if unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) {
		return true
	}
	return false
}

func isPunctuation(r byte) bool {
	if r == '.' || r == ',' || r == '!' || r == ']' || r == '[' || r == '(' || r == ')' || r == '-' || r == ':' || r == '"' || r == '?' || r == '\'' {
		return true
	}
	return false
}

// func parseTag(content string) (Tag, error) {
// 	if content[0] != '<' || content[len(content-1)] != '>' {
// 		returns
// 	}
// }
