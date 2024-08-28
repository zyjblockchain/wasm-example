package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wapc/wapc-go"
	"github.com/wapc/wapc-go/engines/wazero"
	"os"
)

func main() {
	ctx := context.Background()
	wasm, err := os.ReadFile("token/token.wasm")
	if err != nil {
		panic(err)
	}
	engine := wazero.Engine()
	module, err := engine.New(ctx, wapc.NoOpHostCallHandler, wasm, &wapc.ModuleConfig{
		Logger: wapc.PrintlnLogger,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	if err != nil {
		panic(err)
	}
	defer module.Close(ctx)
	instance, err := module.Instantiate(ctx)
	if err != nil {
		panic(err)
	}
	defer instance.Close(ctx)

	go runApi(instance, ctx)

	select {}
}

func runApi(instance wapc.Instance, ctx context.Context) {
	r := gin.Default()
	r.GET("/:handleName", func(c *gin.Context) {
		handleName := c.Param("handleName")
		payload, err := c.GetRawData()
		if err != nil {
			c.JSON(400, err)
			return
		}
		result, err := instance.Invoke(ctx, handleName, payload)
		if err != nil {
			c.JSON(400, err)
			return
		}
		// fmt.Println(string(result))
		c.String(200, string(result))
	})

	r.Run(":8088")
}
