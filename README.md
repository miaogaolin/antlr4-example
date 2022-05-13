# antlr4 命令
根据 g4 文件生成 go 语言，listener 模式。
```shell
antlr4 -Dlanguage=Go MyGrammar.g4
```
visitor模式，命令如下：
```shell
antlr4 -Dlanguage=Go -no-listener -visitor MyGrammar.g4
```
其它参数：
* -o 设置输出
* -package 设置包名，默认 parser