package main

import (
	"testing"
	"reflect")
type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}
func TestWalk(t *testing.T) {
	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls [] string
	} { 
		{ "Struct with one string field",
		struct {
			Name string
		}{ "danizavtz"},
			[]string{"danizavtz"},
		},
		{ "struct with two string fields",
		struct {
			Name string
			City string
		}{"danizavtz","brazil"},
		[]string{"danizavtz","brazil"},
		},
		{ "struct with non string field",
		struct {
			Name string
			Age int
		}{"danizavtz", 35},
		[]string {"danizavtz"},
		},
		{ "Nested fields",
		Person{
			"danizavtz",
			Profile{35, "brazil"},
		},
		[]string{"danizavtz","brazil"},
		},
		{ "pointers to things",
		&Person{
			"danizavtz",
			Profile{35,"brazil"},
		},
		[]string{"danizavtz","brazil"},
		},
		{ "Slices",
		[]Profile{
			{35, "brazil"},
			{34, "Reykjavík"},
		},
		[]string{"brazil","Reykjavík"},
		},
		{ "arrays",
		[2]Profile{
			{35,"danizavtz"},
			{34,"Reykjavík"},
		},
		[]string{"danizavtz", "Reykjavík"},
		},
		{ "Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
		[]string{"Bar","Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string){
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}