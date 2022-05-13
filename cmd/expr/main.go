package main

import (
	. "antlr4-example/expr"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	// 新家一个 CharStream
	input := antlr.NewInputStream("a=1\n")

	lexer := NewExprLexer(input)

	// 新建一个此法符号的缓冲区，用于存储此法分析器将生成的词法符号
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)

	// 新建一个语法分析器，处理词法符号缓冲区中的内容
	parser := NewExprParser(tokens)

	tree := parser.Prog() // 针对 prog 规则，开始语法分析

	fmt.Println(tree.ToStringTree(nil, nil))
}
