## wapc-go 仓库用例
[wapc-go](https://github.com/wapc/wapc-go) 是 golang 实现的 wasm 运行时库。

wapc-go 中集成了 golang 的三种主要的 wasm 运行时库，本用例使用了 wazero 库作为运行时。
本用例的目标是编写一个 golang 代码 token.go , 并通过 tinygo 编译成 wasm 文件。
然后使用 wapc-go 中集成的 wazero 运行时加载 wasm 进行调用。

## 准备工作
### tinygo 安装
安装最新版本 tinygo: https://tinygo.org/getting-started/install

### build wasm
golang 版本最好 1.21+

build命令：`GOOS=wasi GOARCH=wasm  GO111MODULE=on tinygo build -o token.wasm -target=wasi  token.go`

其中 token.wasm 为编译好的 wasm 文件名， toke.go 为需要编译的 golang 代码

## 用例实现
token.go 中的代码通过 wapc-go 规定方式实现的一个简单的 token 合约代码。
并通过 tinygo 编译成 token.wasm

## 测试用例
在 main.go 中负责加载 token.wasm 并启动一个 web 服务用于测试 token.wasm 的代码执行。

在 example_test.go 中实现了用于访问 web 服务来调用 token.wasm 代码的测试用例。

在启动 main.go 之后，再去 example_test.go 中运行每一个测试用例用于调用 token 代码。



