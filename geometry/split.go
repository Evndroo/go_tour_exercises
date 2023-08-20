package geometry

import (
	"fmt"
	"strings"
)

func ParseStringToWorldMap(s string) map[string]int {
	splited := strings.Fields(s)
	fmt.Println(splited)

	worldMap := make(map[string]int)

	for index, value := range splited {
		worldMap[value] = index + 1
	}

	return worldMap
}
