package main

import (
	"fmt"
	"maps"
)

func main() {

	// Map decleration :- var mapVariable map[keyType]valueType
	// Also can use make :- mapVariable := make(map[keyType]valueType)
	// Also using map literal:- mapVariable = map[keyType]valueType{key1: value1, key2: value2}

	myMap := make(map[string]int)
	fmt.Println(myMap) // outputs map[]

	myMap["key1"] = 9
	myMap["key2"] = 20
	fmt.Println(myMap) // outputs map[key1:9 key2:20]

	fmt.Println(myMap["key1"]) // accessing an existant key
	fmt.Println(myMap["key3"]) // zero value of valueType is returned so here 0 is returned

	delete(myMap, "key1") // deleting key1 from myMap
	fmt.Println(myMap)

	delete(myMap, "key6") // nothing happens if non-existent key is passed in the map to be deleted
	fmt.Println(myMap)

	myMap["key3"] = 3
	myMap["key4"] = 4
	fmt.Println(myMap) // Before clearing
	clear(myMap)
	fmt.Println(myMap) // After clearing

	myMap["key1"] = 3
	myMap["key2"] = 4

	// dive depper into accessing a value from the map
	value, ok := myMap["key1"]
	// it returnes another optional value that can be used to check the existence of value associated with the key
	fmt.Println("Value is :", value)
	fmt.Println(ok)

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	// equaliy check for maps
	if maps.Equal(myMap, myMap2) {
		fmt.Println("Mymap and myMap2 are equal")
	} else {
		fmt.Println("Mymap and myMap2 aren't equal")
	}

	// iterating over the map with for loop
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	var myMap4 map[string]string
	if myMap4 == nil {
		fmt.Println("Map is initilized to nil value")
	}

	val := myMap4["key"]
	fmt.Println(val) // zero value of map i.e ""

	// won't work as already initilized to nil value
	// myMap4["key"] = "Value"
	// fmt.Println(myMap4)

	// Solution
	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	fmt.Println(myMap4)

	// can get length of map with len func --> gives number of key
	fmt.Println("myMap4 length is:", len(myMap4))

	// we have nested map concept as well
	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)

}
