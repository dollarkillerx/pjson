# pjson
easy parse json  (通过路径获取JSON中值  不需要struct) (Get the value in JSON by path, no need for struct)

## use

`github.com/dollarkillerx/pjson`

## demo
``` 
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
```
