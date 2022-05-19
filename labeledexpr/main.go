package main

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// 在 antlr4权威指南一书中，按照 java 的做法是继承了
// BaseLabeledExprVisitor，该类型实现了默认的方法，只需要修改像定义的方法即可，
// 剩下的默认会帮你实现，但在 Go 语言中行不通。
// 必须要实现 LabeledExprVisitor 所有接口，这样在如下方法中，才可以走第一个分支，如果只实现了部分那就会走
// default 分支，这样会导致进入了默认 Visitor 实现，但默认的为实现 VisitChildren 方法。
// ```go
// func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
//	switch t := visitor.(type) {
//	case LabeledExprVisitor:
//		return t.VisitProg(s)
//
//	default:
//		return t.VisitChildren(s)
//	}
//}
// ```
// 新的想法：
// BaseLabeledExprVisitor 不会帮你走语法树，所有的步骤需要自己调用。
// 例如： Visit 和 VisitProg 必须写
type evalVisitor struct {
	BaseLabeledExprVisitor
	memory map[string]int
}

func NewEvalVisitor() *evalVisitor {
	visitor := new(evalVisitor)
	visitor.memory = make(map[string]int)
	return visitor
}

func (v *evalVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *evalVisitor) VisitProg(ctx *ProgContext) interface{} {
	for i := 0; i < ctx.GetChildCount(); i++ {
		v.Visit(ctx.Stat(i))
	}
	return 0
}

// Visit a parse tree produced by LabelExprParser#printExpr.
// expr NEWLINE
func (v *evalVisitor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	value := v.Visit(ctx.Expr()) // 计算 expr 子节点的值
	fmt.Println(value)           // 打印结果
	return value
}

// Visit a parse tree produced by LabelExprParser#assign.
// ID '=' expr NEWLINE
func (v *evalVisitor) VisitAssign(ctx *AssignContext) interface{} {
	id := ctx.ID().GetText()     // id 在 '=' 的左侧
	value := v.Visit(ctx.Expr()) // 计算右侧表达式的值
	v.memory[id] = value.(int)   // 将这个映射关系存储在计算器的 “内存” 中
	return value
}

// Visit a parse tree produced by LabelExprParser#parens.
// ‘（’ expr ')'
func (v *evalVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.Visit(ctx.Expr()) // 返回子表达式的值
}

// Visit a parse tree produced by LabelExprParser#MulDiv.
// expr op=('*'|'/') expr
func (v *evalVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	left := v.Visit(ctx.Expr(0))  // 计算左侧子表达式的值
	right := v.Visit(ctx.Expr(1)) // 计算右侧子表达式的值
	if ctx.GetOp().GetTokenType() == LabeledExprParserMUL {
		return left.(int) * right.(int)
	} else {
		return left.(int) / right.(int)
	}
}

// Visit a parse tree produced by LabelExprParser#AddSub.
func (v *evalVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	left := v.Visit(ctx.Expr(0))  // 计算左侧子表达式的值
	right := v.Visit(ctx.Expr(1)) // 计算右侧子表达式的值
	if ctx.GetOp().GetTokenType() == LabeledExprLexerADD {
		return left.(int) + right.(int)
	} else {
		return left.(int) - right.(int)
	}
}

// Visit a parse tree produced by LabelExprParser#id.
// ID
func (v *evalVisitor) VisitId(ctx *IdContext) interface{} {
	id := ctx.ID().GetText()
	if x, ok := v.memory[id]; ok {
		return x
	}
	return 0
}

// Visit a parse tree produced by LabelExprParser#int.
// INT
func (v *evalVisitor) VisitInt(ctx *IntContext) interface{} {
	i, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		return 0
	}
	return i
}

func main() {
	//input, _ := antlr.NewFileStream(os.Args[1])

	input := antlr.NewInputStream(`
193
a = 5
b = 6
a+b*2
(1+2)*3
`)
	lexer := NewLabeledExprLexer(input)

	// 新建一个此法符号的缓冲区，用于存储此法分析器将生成的词法符号
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)

	// 新建一个语法分析器，处理词法符号缓冲区中的内容
	parser := NewLabeledExprParser(tokens)

	tree := parser.Prog() // 针对 prog 规则，开始语法分析

	visitor := NewEvalVisitor()
	visitor.Visit(tree)
}
