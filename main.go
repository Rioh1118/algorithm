package main

import (
	"fmt"

	"github.com/Rioh1118/algorithm/tree"
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
	tree.Insert("Wasm")
	tree.Delete("Mika")
	fmt.Println(tree.Get_min())
	fmt.Println(tree.Get_max())
	fmt.Println(tree)
}
