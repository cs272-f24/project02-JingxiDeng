package main

import (
	"reflect"
	"testing"
)

func TestCleanHref(t *testing.T){
	tests := []struct{
		name string
		hostname string
		hrefs []string
		want []string
	}{
		{
			name: "General Case",
			hostname: "https://myusfca.usfca.edu/",
			hrefs: []string{"/", "/dashboard/", "/dashboard/compsci/", "/arts-sciences/", "https://ddsamuel.com/"},
			want: []string{"https://myusfca.usfca.edu/", "https://myusfca.usfca.edu/dashboard/", "https://myusfca.usfca.edu/dashboard/compsci/", "https://myusfca.usfca.edu/arts-sciences/", "INVALID HREF"},
		},
	}

	for _, test := range tests{
		actualHrefs := Clean(test.hostname, test.hrefs)

		if(!reflect.DeepEqual(test.want, actualHrefs)){
			t.Errorf("ERROR: %s\nExpected: %v\nActual:   %v\n", test.name, test.want, actualHrefs)
		}
	}
}
