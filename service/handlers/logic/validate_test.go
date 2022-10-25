package logic

import (
	"reflect"
	"testing"

	"github.com/google/go-github/v48/github"
	"github.com/sonnht1409/scanning/service/models"
)

func TestServiceLogic_CheckRule(t *testing.T) {
	type args struct {
		content string
		rule    models.RegexRule
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []int
	}{
		{
			name: "test pass public",
			args: args{
				content: "package main.go",
				rule:    PUBLIC_RULE,
			},
			want:  false,
			want1: []int{},
		},
		{
			name: "test pass private",
			args: args{
				content: "package main.go",
				rule:    PRIVATE_RULE,
			},
			want:  false,
			want1: []int{},
		},
		{
			name: "test failed public",
			args: args{
				content: "const public_key = \" \"",
				rule:    PUBLIC_RULE,
			},
			want:  true,
			want1: []int{1},
		},
		{
			name: "test failed public multi line",
			args: args{
				content: `const public_key = \" \"
				
				const public_key = \" \"`,
				rule: PUBLIC_RULE,
			},
			want:  true,
			want1: []int{1, 3},
		},
		{
			name: "test failed private",
			args: args{
				content: "const private_key = \" \"",
				rule:    PRIVATE_RULE,
			},
			want:  true,
			want1: []int{1},
		},
		{
			name: "test failed private multi line",
			args: args{
				content: `const private_key = \" \"
				
				const private_key = \" \"`,
				rule: PRIVATE_RULE,
			},
			want:  true,
			want1: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ServiceLogic{
				githubClient: github.NewClient(nil),
			}
			got, got1 := s.CheckRule(tt.args.content, tt.args.rule)
			if got != tt.want {
				t.Errorf("ServiceLogic.CheckRule() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ServiceLogic.CheckRule() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
