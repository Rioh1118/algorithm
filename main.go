package main

import (
	"github.com/Rioh1118/algorithm/tree"
	"fmt"
)


func main() {
	tree := tree.NewNode[string]("Rio")
	tree.Insert("Mike")
	tree.Insert("John")
	tree.Insert("Doe")
	tree.Insert("Jane")
	tree.Insert("Mika")
	tree.Insert("Mikael")
	tree.Insert("Mikaela")

	fmt.Println(tree)
}
