package main

import (
	"fmt"
	"hello/categories"
)

func main() {
	cates := categories.GetCategories()
	for idx, cate := range cates {
		fmt.Println("Category", idx, ":", cate)
	}
}
