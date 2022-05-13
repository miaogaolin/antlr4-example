/** 用语法文件通常以 grammar 关键字开头
 *  这是一个名为 ArrayInit 的语法，它必须和文件名 ArrayInit.g4 相匹配
 */
grammar ArrayInit;

// 一条名为 init 的规则，它匹配一对花括号中的、逗号分割的 value
init: '{' value (',' value)* '}'; // 必须匹配至少一个 value

// 一个 value 可以是嵌套的花括号结构，也可以是简单的数据，即 INT 词法符号
value: init
    | INT
    ;

// 语法分析器的规则必须以小写字母开头，词法分析器的规则必须用大写字母开头
INT: [0-9]+; // 定义语法符号 INT，它由一个或多个数字组成
WS: [ \t\r\n]+ -> skip; // 定义此法规则 “空白符号”，丢弃之