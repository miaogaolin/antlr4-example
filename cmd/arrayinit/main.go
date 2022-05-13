package main

import (
	. "antlr4-example/arrayinit"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input := antlr.NewInputStream("{1,{2,3},4}")
	// 新建一个词法分析器，处理输入
	lexer := NewArrayInitLexer(input)

	// 新建一个词法符号的缓冲区，用于存储词法分析器将生成的词法符号
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// 新建一个语法分析器，处理词法符号缓冲区中的内容
	parser := NewArrayInitParser(tokens)

	// 针对 init 规则，开始语法分析
	tree := parser.Init()
	// 用 LISP 风格打印生成的树
	fmt.Println(tree.ToStringTree(nil, nil))
}
