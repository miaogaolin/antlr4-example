grammar Config;
import LexerRules;



stat: NL* sections | NL* attributes NL* sections;
// 键值对
attributes: attribute+;
attribute: ID {attributeCount++;} ':' .*? NL {attributeCount--;};

// numRow: 'Num' ':' cols NL; cols: (COL ',')* COL?;

// versionRow: 'Version' ':' version NL; version: DEC '.' DEC '.' DEC;

sections: section+;
section: condition template;

condition: ';' {conditionCount++;} expr? NL {conditionCount--;};
template: .+? NL NL+ | .+? NL* EOF;

expr: (arrayExpr | stringExpr | numExpr) (AND | OR) (
		arrayExpr
		| stringExpr
		| numExpr
	)
	| arrayExpr
	| stringExpr
	| numExpr
	| '(' expr ')';

arrayExpr: conditionKey (IN | NOT IN) array;
stringExpr: conditionKey ('=~' | '!~' | '==' | '!=') STRING;
numExpr: conditionKey ('==' | '>' | '<' | '!=') number;

conditionKey: '{' COL '}';

array:
	'[' STRING (',' STRING)* ']' // 字符串数组
	| '[' number (',' number)* ']'; // 数字数组
number: FLOAT | DEC;



