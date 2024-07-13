package handler

import (
	"fmt"
)

func MapHandler() {
	map1 := make(map[string]string)
	fmt.Println("empty map1", map1)
	map1["up"] = "vikas"
	map1["mp"] = "saurabh"
	map1["mh"] = "komal"
	map1["bh"] = "harsh"

	fmt.Println("map1", map1)
	// print value by key
	fmt.Println("key UP map1 value", map1["up"])

	map2 := map[string]string{"up": "luck", "bihar": "patna", "maharashtra": "mumbai"}
	fmt.Println("map2", map2)
	// iterate on map
	for key,val:=range map1{
		fmt.Println("map1 key:",key,"value:",val)
	}

	// checking if key is present
	if val,ok:=map1["up"]; ok {
		fmt.Println("map1 key is present value is",val)
	}else {
		fmt.Println("map1 key is not present ")
	}
}
