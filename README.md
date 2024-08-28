## wapc-go 仓库用例
[wapc-go](https://github.com/wapc/wapc-go) 是 golang 实现的 wasm 运行时库。

### 用例实现
token.go 中的代码通过 wapc-go 规定方式实现的一个简单的 token 合约代码。并通过 tinygo 编译成 token.wasm

### 测试用例
在 main.go 中负责加载 token.wasm 并启动一个 web 服务用于测试 token.wasm 的代码执行。
在 example_test.go 中实现了用于访问 web 服务来调用 token.wasm 代码的测试用例。
在启动 main.go 之后，再去 example_test.go 中运行每一个测试用例用于调用 token 代码。



