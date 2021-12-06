package pjson

import (
	"encoding/json"
	"reflect"
	"strings"
)

type PJson struct {
	m map[string]interface{}
	l []interface{}

	isM bool
	err error
}

func Parse(json string) *PJson {
	return parse([]byte(json))
}

func ParseBytes(json []byte) *PJson {
	return parse(json)
}

func parse(js []byte) *PJson {
	var m = map[string]interface{}{}
	var l = []interface{}{}
	var e error
	isM := true

	err := json.Unmarshal(js, &m)
	if err != nil {
		err := json.Unmarshal(js, &l)
		if err != nil {
			e = err
			return nil
		}
		isM = false
	}

	return &PJson{
		l:   l,
		m:   m,
		err: e,
		isM: isM,
	}
}

func (p *PJson) Get(path string) ([]interface{}, error) {
	var result []interface{}

	pathLis := strings.Split(path, ".")
	if p.isM {
		p.core(p.m, pathLis, func(i interface{}, ex bool) {
			if ex {
				result = append(result, i)
			}
		})
	} else {
		p.core(p.l, pathLis, func(i interface{}, ex bool) {
			if ex {
				result = append(result, i)
			}
		})
	}

	return result, p.err
}

type packageFunc func(i interface{}, ex bool)

func (p *PJson) core(m interface{}, path []string, fn packageFunc) {
	if reflect.TypeOf(m).Kind() == reflect.Map {
		px, ok := m.(map[string]interface{})
		if !ok {
			fn(nil, false)
			return
		}

		if len(path) >= 2 {
			i, ex := px[path[0]]
			if !ex {
				fn(nil, false)
				return
			}

			if reflect.TypeOf(i).Kind() == reflect.Slice || reflect.TypeOf(i).Kind() == reflect.Map {
				p.core(i, path[1:], fn)
				return
			}

			fn(i, true)
			return
		} else {
			i, ex := px[path[0]]
			if !ex {
				fn(nil, false)
				return
			}
			fn(i, true)
			return
		}
	}

	if reflect.TypeOf(m).Kind() == reflect.Slice {
		p, ok := m.([]interface{})
		if !ok {
			fn(nil, false)
			return
		}

		for i := range p {
			fn(p[i], true)
		}
	}
}

func (p *PJson) Error() error {
	return p.err
}
