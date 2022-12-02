package main

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Section struct {
	Condition string
	Template  string
}

type Config struct {
	BaseConfigListener
	Attributes map[string]string
	Sections   []Section
}

func (c *Config) EnterAttribute(ctx *AttributeContext) {
	text := ctx.GetText()
	data := strings.Split(text, ":")
	if len(data) != 2 {
		return
	}
	if c.Attributes == nil {
		c.Attributes = make(map[string]string)
	}

	c.Attributes[data[0]] = strings.TrimSpace(data[1])
}

func (c *Config) EnterSection(ctx *SectionContext) {
	// con := ctx.Condition().GetText()

	start := ctx.Template().GetSourceInterval().Start
	stop := ctx.Template().GetSourceInterval().Stop
	tpl := ctx.GetStart().GetInputStream().GetText(start, stop)
	fmt.Println("-----------")
	fmt.Println(tpl)
	fmt.Println("----------")
}

func main() {
	// 新家一个 CharStream
	input := antlr.NewInputStream(`
;{col1}==1
	aaa
bbbb

`)

	lexer := NewConfigLexer(input)

	// 新建一个此法符号的缓冲区，用于存储此法分析器将生成的词法符号
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)

	// 新建一个语法分析器，处理词法符号缓冲区中的内容
	parser := NewConfigParser(tokens)

	tree := parser.Stat() // 针对 prog 规则，开始语法分析

	// 新建一个通用的、能够触发回调的语法分析树遍历器
	walker := antlr.NewParseTreeWalker()

	// 遍历语法分析过程中生成的语法分析树，触发回调
	cfg := new(Config)
	walker.Walk(cfg, tree)
}
