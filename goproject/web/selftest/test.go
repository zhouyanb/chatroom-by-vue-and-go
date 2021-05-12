package main

import "fmt"

func main() {
	 student:= make(map[string]string)
	student["saber"] = "123"
	student["archer"] = "345"
	student["ooc"] = "11"
	for key := range student {
		fmt.Println(key)
	}
}
