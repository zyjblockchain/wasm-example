package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gentleman.v2"
	"testing"
)

var url = "http://127.0.0.1:8088"

func Test_Info(t *testing.T) {
	name := "Info"
	payload := []byte("")
	result, err := submitHandle(name, payload, url)
	assert.NoError(t, err)

	t.Log(string(result))
}

func Test_Mint(t *testing.T) {
	name := "Mint"
	mmap := map[string]string{
		"Recipient": "AAAAAAAA",
		"Quantity":  "100000000000000",
	}
	payload, _ := json.Marshal(mmap)
	result, err := submitHandle(name, payload, url)
	assert.NoError(t, err)

	t.Log(string(result))
}

func Test_Transfer(t *testing.T) {
	name := "Transfer"
	mmap := map[string]string{
		"From":      "AAAAAAAA",
		"Recipient": "BBBBBBBB",
		"Quantity":  "30000000000000",
	}
	payload, _ := json.Marshal(mmap)
	result, err := submitHandle(name, payload, url)
	assert.NoError(t, err)

	t.Log(string(result))
}

func Test_Balance(t *testing.T) {
	name := "Balance"
	payload := []byte("AAAAAAAA")
	result, err := submitHandle(name, payload, url)
	assert.NoError(t, err)

	t.Log(string(result))
}

func submitHandle(name string, payload []byte, url string) ([]byte, error) {
	cli := gentleman.New().URL(url)
	req := cli.Request()
	req.AddPath(fmt.Sprintf("/%s", name))
	req.JSON(payload)
	resp, err := req.Send()
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
