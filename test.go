package main

import (
	"fmt"
	"github.com/kylelemons/godebug/pretty"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"mapsliceutils/mapsliceutils"
	"os"
)

type Dummy struct {
	Config *yaml.MapSlice
}

func NewDummy() *Dummy {
	var d Dummy
	return &d
}

func dummy1() {
	mapSlice := &yaml.MapSlice{}
	myYaml, _ := ioutil.ReadFile("/Users/gdanko/go/src/ct/test/nginx_base1.yml")
	if err := yaml.Unmarshal(myYaml, mapSlice); err != nil { panic(err) }
	exists := mapsliceutils.KeyExists(mapSlice, "http.main.default_type")
	fmt.Println(exists)
}

func dummy2() {
	mapSlice := &yaml.MapSlice{}
	myYaml, _ := ioutil.ReadFile("/Users/gdanko/go/src/mapsliceutils/dest.yml")
	if err := yaml.Unmarshal(myYaml, mapSlice); err != nil { panic(err) }
	//pretty.Print(mapSlice)
	//fmt.Println(999)
	mapsliceutils.KeyDelete(mapSlice, "http.main.server_blocks")
	pretty.Print(mapSlice)
}

func dummy3() {
	mapSlice := &yaml.MapSlice{}
	myYaml, _ := ioutil.ReadFile("/Users/gdanko/go/src/ct/test/nginx_base1.yml")
	if err := yaml.Unmarshal(myYaml, mapSlice); err != nil { panic(err) }
	foo := mapsliceutils.KeyGet(mapSlice, "http")

	pretty.Print(foo)
	fmt.Printf("foo is a %T\n", foo)
	fmt.Printf("foo.Value is a %T\n", foo.Value)
	fmt.Printf("foo.Value.(yaml.MapSlice) is a %T\n", foo.Value.(yaml.MapSlice))


	d := NewDummy()
	//a := &yaml.MapSlice{}
	slice := foo.Value.(yaml.MapSlice)

	fmt.Printf("slice is a %T\n", &slice)
	

	d.Config = &slice

	//fmt.Println((*foo.Value))
	//(*d.Config) = append((*d.Config), (*foo.Value.(yaml.MapSlice))[0])

	//(*d.Config) = append((*d.Config), *foo)
	//&d.Config = &foo.Value.(yaml.MapSlice)
	//pretty.Print(d)


}

func dummy4() {
	mapSlice := &yaml.MapSlice{}
	myYaml, _ := ioutil.ReadFile("/Users/gdanko/go/src/ct/test/nginx_base1.yml")
	if err := yaml.Unmarshal(myYaml, mapSlice); err != nil { panic(err) }
	flattened := mapsliceutils.Flatten(mapSlice, ".")
	pretty.Print(flattened)
}

func dummy5() {
	destFile := "/Users/gdanko/go/src/mapsliceutils/dest.yml"
	sourceFile := "/Users/gdanko/go/src/mapsliceutils/source.yml"
	destSlice := &yaml.MapSlice{}
	sourceSlice := &yaml.MapSlice{}
	destBytes, _ := ioutil.ReadFile(destFile)
	sourceBytes, _ := ioutil.ReadFile(sourceFile)
	if err := yaml.Unmarshal(destBytes, destSlice); err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(sourceBytes, sourceSlice); err != nil {
		panic(err)
	}
	mergedSlice := mapsliceutils.DeepMerge(destSlice, sourceSlice)
	pretty.Print(mergedSlice)
	
	
	


}



func main() {
	/*a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Remove the element at index i from a.
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	fmt.Println(a)

	a[len(a)-1] = ""   // Erase last element (write zero value).
	fmt.Println(a)
	
	a = a[:len(a)-1]   // Truncate slice.

	fmt.Println(a) // [A B E D]


	os.Exit(0)*/
	//dummy1() // KeyExists
	//dummy2() // KeyDelete
	dummy3() // KeyGet
	//dummy4() // Flatten
	//dummy5()
	os.Exit(0)
	fmt.Println(1)
}