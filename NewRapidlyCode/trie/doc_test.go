package trie

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestDocParse(t *testing.T) {
	list3 := &StreamNode{
		Data:      "ccc",
		Appending: "",
		Type:      Word,
	}
	list2 := &StreamNode{
		Data:      "bbb",
		Appending: " ",
		Type:      Word,
		Next:      list3,
	}
	list1 := &StreamNode{
		Data:      "aaa",
		Appending: " ",
		Type:      Word,
		Next:      list2,
	}

	tag6 := &StreamNode{
		Data: "</p>",
		Type: Tag,
	}
	tag5 := &StreamNode{
		Data:      "ccc",
		Appending: "",
		Type:      Punctuation,
		Next:      tag6,
	}
	tag4 := &StreamNode{
		Data:      "ccc",
		Appending: "",
		Type:      Word,
		Next:      tag5,
	}
	tag3 := &StreamNode{
		Data:      "bbb",
		Appending: " ",
		Type:      Word,
		Next:      tag4,
	}
	tag2 := &StreamNode{
		Data:      "aaa",
		Appending: " ",
		Type:      Word,
		Next:      tag3,
	}
	tag1 := &StreamNode{
		Data: "<p>",
		Type: Tag,
		Next: tag2,
	}

	tests := []struct {
		name string
		args string
		want *StreamNode
	}{
		{
			name: "list",
			args: "aaa bbb ccc",
			want: list1,
		}, {
			name: "tags",
			args: "<p>aaa bbb ccc.</p>",
			want: tag1,
		},
	}

	for _, tt := range tests {
		root := tt.want
		result := DocParse(tt.args)
		for result.Next == nil && root.Next == nil {
			if result.Next != nil || root.Next != nil {
				t.Error("Unequal length")
			}
			if result.Type != root.Type {
				t.Error("Unequal type")
			}
			if result.Data != root.Data {
				t.Error("unequal data")
			}
			if result.Appending != root.Appending {
				t.Error("unequal appending")
			}

			result = result.Next
			root = root.Next
		}
	}

}

func TestDocParse2(t *testing.T) {
	file, err := os.Open("test.html")
	if err != nil {
		t.Error("can't open test.html")
	}
	content, err := io.ReadAll(file)
	if err != nil {
		t.Error("can't read test.html")
	}
	result := DocParse(string(content))
	for result.Next != nil {
		fmt.Println(result.Data + "|" + result.Appending + "||")
		result = result.Next
	}

}
