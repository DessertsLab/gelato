package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Rule map[string][]interface{}

func drill(r Rule) {
	for k, v := range r {
		fmt.Println(k, "|", v)
		fmt.Println("---------------------------")

		for _, t := range v {
			switch i := t.(type) {
			case string:
				fmt.Println("sssssssssssssss")
			case []string:
				fmt.Println(i)
			default:
				fmt.Println("others", i)
			}
		}

		fmt.Println("")

	}
}

func main() {
	r1 := Rule{}
	file, err := ioutil.ReadFile("../data/drillrules/r1.json")
	if err != nil {
		panic(err)
	}
	if json.Unmarshal([]byte(file), &r1); err != nil {
		panic(err)
	}
	drill(r1)
}
