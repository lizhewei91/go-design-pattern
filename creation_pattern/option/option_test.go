package main

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		name     string
		location string
		level    string
		age      int
	}
	tests := []struct {
		name string
		args args
		want *Person
	}{
		{
			name: "test1",
			args: args{
				name:     "lzw",
				age:      18,
				location: "beijing",
				level:    "high",
			},
			want: &Person{
				Name: "lzw",
				Age:  18,
				Opt: options{
					location: "beijing",
					level:    "high",
				},
			},
		},

		{
			name: "test2",
			args: args{
				name:     "lzw",
				age:      18,
				location: "beijing",
				level:    "collage",
			},
			want: &Person{
				Name: "lzw",
				Age:  18,
				Opt: options{
					location: "beijing",
					level:    "collage",
				},
			},
		},

		{
			name: "test3",
			args: args{
				name:     "lzw",
				age:      18,
				location: "shanghai",
				level:    "high",
			},
			want: &Person{
				Name: "lzw",
				Age:  18,
				Opt: options{
					location: "shanghai",
					level:    "high",
				},
			},
		},

		{
			name: "test4",
			args: args{
				name:     "lzw",
				age:      18,
				location: "shanghai",
				level:    "collage",
			},
			want: &Person{
				Name: "lzw",
				Age:  18,
				Opt: options{
					location: "shanghai",
					level:    "collage",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.name, tt.args.location, tt.args.level, tt.args.age); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
