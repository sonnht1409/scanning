package common

import "testing"

func TestStringify(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestStringify success",
			args: args{
				in: map[string]interface{}{
					"abc": "123",
				},
			},
			want: "{\"abc\":\"123\"}",
		},
		{
			name: "TestStringify empty",
			args: args{
				in: map[string]interface{}{},
			},
			want: "{}",
		},
		{
			name: "TestStringify failed ",
			args: args{
				in: make(chan int),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Stringify(tt.args.in); got != tt.want {
				t.Errorf("Stringify() = %v, want %v", got, tt.want)
			}
		})
	}
}
