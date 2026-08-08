package main

import (
	"fmt"
	"maps"
)

func main() {

	//var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)

	// using a Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2,
	// }

	myMap := make(map[string]int)

	fmt.Println("Initial map:", myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println("Map after adding key1:", myMap)

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])

	myMap["code"] = 21
	fmt.Println("Map after updating code:", myMap)

	delete(myMap, "key1")
	fmt.Println("Map after deleting key1:", myMap)

	// check if a key exists in the map
	value, exists := myMap["key1"]
	if exists {
		fmt.Println("Key 'key1' exists with value:", value)
	} else {
		fmt.Println("Key 'key1' does not exist.")
	}

	// check if a key exists in the map
	value, exists = myMap["code"]
	if exists {
		fmt.Println("Key 'code' exists with value:", value)
	} else {
		fmt.Println("Key 'code' does not exist.")
	}

	// clear(myMap)
	// fmt.Println("Map after clearing:", myMap)

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11

	_, unknownvalue := myMap["key1"]
	// fmt.Println(value)
	fmt.Println("Is a value associated with key1:", unknownvalue)

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println("Map 2:", myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("myMap3 and myMap2 are equal.")
	} else {
		fmt.Println("myMap and myMap2 are not equal.")
	}

	for k, v := range myMap3 {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	var myMap4 map[string]string
	if myMap4 == nil {
		fmt.Println("myMap4 is nil.")
	} else {
		fmt.Println("myMap4 is not nil.")
	}

	val1 := myMap4["key"]
	fmt.Println(val1)

	// myMap4["key"] = "value" // This will cause a runtime panic because myMap4 is nil

	myMap4 = make(map[string]string)
	myMap4["key"] = "value"
	fmt.Println("myMap4 after adding a key-value pair:", myMap4)

	fmt.Println("myMap4 lenth is", len(myMap))

	myMap5 := make(map[string]map[string]string)

	myMap5["map1"] = myMap4
	fmt.Println("myMap5")

}
