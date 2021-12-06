package pjson

import (
	"fmt"
	"log"
	"testing"
)

func TestPJsonTestMapJson(t *testing.T) {
	r := `
	{
	"a": {
			"p" : {  "c" : "162"}
		}
	}
`

	json := Parse(r)
	get, err := json.Get("a.p.c")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(get)
}

func TestPJsonTestSliceJson(t *testing.T) {
	r := `
	{
	"a": {
			"p" : [{  "c" : "162"},{  "c" : "1622"}]
		}
	}
`

	json := Parse(r)
	get, err := json.Get("a.p.c")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(get)
}

func TestPJsonTestSliceJson2(t *testing.T) {
	r := `
	[{
	"a": {
			"p" : [{  "c" : "162"},{  "c" : "1622"}]
		}
	}]
`

	json := Parse(r)
	get, err := json.Get("a.p.c")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(get)
}
