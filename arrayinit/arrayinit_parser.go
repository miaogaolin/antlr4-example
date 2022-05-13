// Code generated from ArrayInit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ArrayInit

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ArrayInitParser struct {
	*antlr.BaseParser
}

var arrayinitParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func arrayinitParserInit() {
	staticData := &arrayinitParserStaticData
	staticData.literalNames = []string{
		"", "'{'", "','", "'}'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "INT", "WS",
	}
	staticData.ruleNames = []string{
		"init", "value",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 20, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 9, 8,
		0, 10, 0, 12, 0, 12, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 18, 8, 1, 1, 1,
		0, 0, 2, 0, 2, 0, 0, 19, 0, 4, 1, 0, 0, 0, 2, 17, 1, 0, 0, 0, 4, 5, 5,
		1, 0, 0, 5, 10, 3, 2, 1, 0, 6, 7, 5, 2, 0, 0, 7, 9, 3, 2, 1, 0, 8, 6, 1,
		0, 0, 0, 9, 12, 1, 0, 0, 0, 10, 8, 1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11,
		13, 1, 0, 0, 0, 12, 10, 1, 0, 0, 0, 13, 14, 5, 3, 0, 0, 14, 1, 1, 0, 0,
		0, 15, 18, 3, 0, 0, 0, 16, 18, 5, 4, 0, 0, 17, 15, 1, 0, 0, 0, 17, 16,
		1, 0, 0, 0, 18, 3, 1, 0, 0, 0, 2, 10, 17,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ArrayInitParserInit initializes any static state used to implement ArrayInitParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewArrayInitParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ArrayInitParserInit() {
	staticData := &arrayinitParserStaticData
	staticData.once.Do(arrayinitParserInit)
}

// NewArrayInitParser produces a new parser instance for the optional input antlr.TokenStream.
func NewArrayInitParser(input antlr.TokenStream) *ArrayInitParser {
	ArrayInitParserInit()
	this := new(ArrayInitParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &arrayinitParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ArrayInit.g4"

	return this
}

// ArrayInitParser tokens.
const (
	ArrayInitParserEOF  = antlr.TokenEOF
	ArrayInitParserT__0 = 1
	ArrayInitParserT__1 = 2
	ArrayInitParserT__2 = 3
	ArrayInitParserINT  = 4
	ArrayInitParserWS   = 5
)

// ArrayInitParser rules.
const (
	ArrayInitParserRULE_init  = 0
	ArrayInitParserRULE_value = 1
)

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArrayInitParserRULE_init
	return p
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArrayInitParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *InitContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *ArrayInitParser) Init() (localctx IInitContext) {
	this := p
	_ = this

	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ArrayInitParserRULE_init)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Match(ArrayInitParserT__0)
	}
	{
		p.SetState(5)
		p.Value()
	}
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ArrayInitParserT__1 {
		{
			p.SetState(6)
			p.Match(ArrayInitParserT__1)
		}
		{
			p.SetState(7)
			p.Value()
		}

		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(13)
		p.Match(ArrayInitParserT__2)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArrayInitParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArrayInitParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Init() IInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitContext)
}

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(ArrayInitParserINT, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayInitListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *ArrayInitParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ArrayInitParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ArrayInitParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(15)
			p.Init()
		}

	case ArrayInitParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(16)
			p.Match(ArrayInitParserINT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
