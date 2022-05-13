package main

import (
	. "antlr4-example/arrayinit"
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type SortToUnicodeString struct {
	BaseArrayInitListener
}

// EnterInit 将 { 翻译为 "
func (s *SortToUnicodeString) EnterInit(c *InitContext) {
	fmt.Print(`"`)
}

func (s *SortToUnicodeString) EnterValue(c *ValueContext) {
	if c.INT() == nil {
		return
	}
	value, _ := strconv.ParseInt(c.INT().GetText(), 16, 32)
	fmt.Printf("%#U", value)
}

// ExitInit 将 } 翻译为 "
func (s *SortToUnicodeString) ExitInit(c *InitContext) {
	fmt.Print(`"`)
}

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

	// 新建一个通用的、能够触发回调的语法分析树遍历器
	walker := antlr.NewParseTreeWalker()

	// 遍历语法分析过程中生成的语法分析树，触发回调
	walker.Walk(new(SortToUnicodeString), tree)
	fmt.Println()
}
