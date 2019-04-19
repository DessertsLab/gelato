package main

import (
	"fmt"

	"github.com/geo/geo"
)

func main() {
	res, _ := geo.GetGeo("南京环亚医疗美容门诊部有限公司")
	fmt.Println(res.getLat(), res.getLng())
}
