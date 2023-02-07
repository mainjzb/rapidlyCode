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
		if firstRune == '<' {
			right = strings.IndexRune(content, '>') + 1
			node := &StreamNode{
				Data: content[:right],
				Type: Tag,
				Prev: prev,
			}
			prev.Next = node
			prev = node
		} else if firstRune == ' ' {
			continue
		} else if size == 1 && isPunctuation(byte(firstRune)) {
			right = left + 1
			for isPunctuation(content[right]) {
				right++
			}
			node := &StreamNode{
				Data: content[left:right],
				Type: Punctuation,
				Prev: prev,
			}
			prev.Next = node
			prev = node
		} else {

		}
		left = right
	}
	return root.Next
}

func isPunctuation(r byte) bool {
	if r == '.' || r == '!' || r == ']' || r == '[' || r == '(' || r == ')' {
		return true
	}
	return false
}

// func parseTag(content string) (Tag, error) {
// 	if content[0] != '<' || content[len(content-1)] != '>' {
// 		returns
// 	}
// }
