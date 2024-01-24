package leetcode30

import (
	"reflect"
	"testing"
)

func Test_getGlobalCombination(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		{
			name: "1. abc",
			args: args{words: []string{"ab", "cd", "ef"}},
			want: map[string]bool{
				"abcdef": true, // a b c
				"abefcd": true, // a c b
				"cdabef": true, // b a c
				"cdefab": true, // b c a
				"efabcd": true, // c a b
				"efcdab": true, // c b a
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGlobalCombination(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGlobalCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "good"},
			},
			want: []int{8},
		},
		{
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "word"},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
