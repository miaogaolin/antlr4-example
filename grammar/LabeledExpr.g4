grammar LabeledExpr;
import CommonLexerRules;
// 起始规则，语法分析的七点
prog:   stat+;

stat:   expr NEWLINE            # printExpr
    |   ID '=' expr NEWLINE     # assign
    |   NEWLINE                 # blank
    ;
expr:   expr op=(MUL|DIV) expr     # MulDiv
    |   expr op=(ADD|SUB) expr     # AddSub
    |   INT                     # int
    |   ID                      # id
    |   '(' expr ')'            # parens
    ;

