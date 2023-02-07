package trie

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestDocParse(t *testing.T) {
	file, err := os.Open("test.html")
	if err != nil {
		t.Error("can't open test.html")
	}
	content, err := io.ReadAll(file)
	if err != nil {
		t.Error("can't read test.html")
	}

	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want *Doc
	}{
		{
			name: "test.html",
			args: args{string(content)},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DocParse(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocParse() = %v, want %v", got, tt.want)
			}
		})
	}
}
