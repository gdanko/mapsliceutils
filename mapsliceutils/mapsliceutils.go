package mapsliceutils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
)

var keyExistsBool bool
var keyDeleteKeyFound bool
var keyGetResult *yaml.MapItem

func KeyExists(mapSlice *yaml.MapSlice, path string) (bool) {
	keyExistsBool = false
	keyExists(mapSlice, path, "", 0)
	return keyExistsBool
}

func keyExists(mapSlice *yaml.MapSlice, path string, pathItem string, level int) {
	pathBits := strings.Split(path, ".")
	pathLength := len(pathBits)
	if level > pathLength-1 { return }
	pathItem = pathBits[level]
	key := pathBits[len(pathBits)-1]
	for _, item := range *mapSlice {
		if keyExistsBool {return}
		if pathLength-1 == level && item.Key.(string) == key {
			keyExistsBool = true
			return
		}
		nestedSlice, mapIsNested := item.Value.(yaml.MapSlice)
		if mapIsNested { keyExists(&nestedSlice, path, pathItem, level+1) }
	}
}

func KeyDelete(mapSlice *yaml.MapSlice, path string) {
	keyDeleteKeyFound = false
	keyDelete(mapSlice, path, "", 0)
}

func keyDelete(mapSlice *yaml.MapSlice, path string, pathItem string, level int) {
	pathBits := strings.Split(path, ".")
	pathLength := len(pathBits)
	if level > pathLength-1 { return }
	pathItem = pathBits[level]
	key := pathBits[len(pathBits)-1]
	for i, item := range *mapSlice {
		fmt.Println(item.Key)
		if keyDeleteKeyFound {return}
		if pathLength-1 == level && item.Key.(string) == key {

			(*mapSlice) = (*mapSlice)[:i+copy((*mapSlice)[i:], (*mapSlice)[i+1:])]
			//(*mapSlice) = append((*mapSlice)[:i], (*mapSlice)[i+1:]...)
			//copy((*mapSlice)[i:], (*mapSlice)[i+1:]) // Copy last element to index i
			//(*mapSlice)[len((*mapSlice))-1] = yaml.MapItem{Key: "dummy", Value: "dummy"} // Erase last element (write zero value).
			//(*mapSlice) = (*mapSlice)[:len((*mapSlice))-1]   // Truncate slice.
			return
		}
		nestedSlice, mapIsNested := item.Value.(yaml.MapSlice)
		if mapIsNested { keyDelete(&nestedSlice, path, pathItem, level+1) }
	}	
}

func KeyGet(mapSlice *yaml.MapSlice, path string) (*yaml.MapItem) {
	keyGetResult = nil
	keyGet(mapSlice, path, "", 0)
	return keyGetResult
}

func keyGet(mapSlice *yaml.MapSlice, path string, pathItem string, level int) {
	pathBits := strings.Split(path, ".")
	pathLength := len(pathBits)
	if level > pathLength-1 { return }
	pathItem = pathBits[level]
	key := pathBits[len(pathBits)-1]
	for _, item := range *mapSlice {
		if keyGetResult != nil {return}
		if pathLength-1 == level && item.Key.(string) == key {
			keyGetResult = &item
			return
		}
		nestedSlice, mapIsNested := item.Value.(yaml.MapSlice)
		if mapIsNested { keyGet(&nestedSlice, path, pathItem, level+1) }
	}	
}

func Flatten(mapSlice *yaml.MapSlice, delimiter string) ([]string) {
	flattened := []string{}
	prefix := []string{}
	level := 0
	previous_level := 0
	if delimiter == "" { delimiter = "|" }
	recursiveFlatten(mapSlice, &flattened, prefix, delimiter, level, previous_level)
	return flattened
}

func recursiveFlatten(mapSlice *yaml.MapSlice, flattened *[]string, prefix []string, delimiter string, level int, previous_level int) {
	for _, item := range *mapSlice {
		itemKey := item.Key.(string)
		if level == 0 { prefix = []string{} }
		if previous_level == level && len(prefix) > 0 { prefix = prefix[:len(prefix)-1] }
		nestedSlice, isNested := item.Value.(yaml.MapSlice)
		if isNested {
			previous_level = level
			prefix = append(prefix, itemKey)
			recursiveFlatten(&nestedSlice, flattened, prefix, delimiter, level+1, previous_level)
		} else {
			tempArray := append(prefix, itemKey)
			*flattened = append(*flattened, strings.Join(tempArray, delimiter))
		}
	}
}

