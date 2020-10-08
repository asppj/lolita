package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Printf("1 ctx pointer:%+v\n", ctx)
	fmt.Printf("2 ctx pointer:%+v\n", &ctx)
	find(ctx)
	finder(&ctx)
	finder(ctx)
	finder2(ctx)
	finder2(&ctx)
}

func find(ctx context.Context) {
	fmt.Printf("1 ctx :%+v\n", ctx)
	fmt.Printf("2 ctx pointer:%+v\n", &ctx)
}
func finder(ctx interface{}) {
	fmt.Printf("1er ctx :%+v\n", ctx)
	fmt.Printf("2er ctx pointer:%+v\n", &ctx)
}
func finder2(ctx interface{}) {
	if c, ok := ctx.(context.Context); ok {
		fmt.Printf("1er c :%+v\n", c)
		fmt.Printf("2er c pointer:%+v\n", &c)
		return
	}
	fmt.Println("err")
}
