package form

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFields(t *testing.T) {

	var nilTestTypePtr *struct {
		Name string
		Age  int
	}

	tests := map[string]struct {
		strct interface{}
		want  []field
	}{
		"Simplest Use Case": {
			strct: struct {
				Name string
			}{},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "",
				},
			},
		}, "Field names should be determined from the struct": {
			strct: struct {
				Name string
			}{},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "",
				},
			},
		}, "Multiple fields should be supported": {
			strct: struct {
				Name  string
				Email string
				Age   int
			}{},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "",
				}, {
					Label:       "Email",
					Name:        "Email",
					Type:        "text",
					Placeholder: "Email",
					Value:       "",
				}, {
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       0,
				},
			},
		}, "Values should be parsed": {
			strct: struct {
				Name  string
				Email string
				Age   int
			}{
				Name:  "Zack",
				Email: "zack@test.com",
				Age:   123,
			},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "Zack",
				}, {
					Label:       "Email",
					Name:        "Email",
					Type:        "text",
					Placeholder: "Email",
					Value:       "zack@test.com",
				}, {
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       123,
				},
			},
		}, "Pointers to structs should still work": {
			strct: &struct {
				Name string
				Age  int
			}{
				Name: "Zack",
				Age:  123,
			},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "Zack",
				}, {
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       123,
				},
			},
		}, "Nil pointers with a struct type should be supported": {
			strct: nilTestTypePtr,
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "",
				}, {
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       0,
				},
			},
		}, "Pointer fields should be supported": {
			strct: &struct {
				Name *string
				Age  *int
			}{},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "",
				}, {
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       0,
				},
			},
		}, "Nested structs should be supported": {
			strct: struct {
				Name    string
				Address struct {
					Street string
					Zip    int
				}
			}{
				Name: "Zack",
				Address: struct {
					Street string
					Zip    int
				}{
					Street: "123 Elm St",
					Zip:    90210,
				},
			},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "Zack",
				},
				{
					Label:       "Street",
					Name:        "Address.Street",
					Type:        "text",
					Placeholder: "Street",
					Value:       "123 Elm St",
				},
				{
					Label:       "Zip",
					Name:        "Address.Zip",
					Type:        "text",
					Placeholder: "Zip",
					Value:       90210,
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := fields(tc.strct)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("fields() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestFields_InvalidVTypes(t *testing.T) {
	tests := []struct {
		notaStruct interface{}
	}{
		{"this is a string"},
		{123},
		{nil},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%t", tc.notaStruct), func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Errorf("fields(%v) did not panic", tc.notaStruct)
				}
			}()
			fields(tc.notaStruct)
		})
	}
}
