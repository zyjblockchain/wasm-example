package main

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/wapc/wapc-guest-tinygo"
	"math/big"
)

/*
build to wasm
GOOS=wasi GOARCH=wasm  GO111MODULE=on tinygo build -o token.wasm -target=wasi  token.go
*/
var (
	Balances     = make(map[string]*big.Int)
	Name         = "Token_Name"
	Ticker       = "TOKEN"
	Denomination = 12
	TotalSupply  = big.NewInt(0)
)

func Info(payload []byte) ([]byte, error) {
	mmap := map[string]interface{}{
		"Name":         Name,
		"Ticker":       Ticker,
		"Denomination": Denomination,
		"TotalSupply":  TotalSupply,
	}

	return json.Marshal(mmap)
}

func Balance(payload []byte) ([]byte, error) {
	addr := string(payload)
	bal, ok := Balances[addr]
	if !ok {
		bal = big.NewInt(0)
	}
	return []byte(bal.String()), nil
}

func Transfer(payload []byte) ([]byte, error) {
	from := gjson.GetBytes(payload, "From").String()
	recipient := gjson.GetBytes(payload, "Recipient").String()
	qty := gjson.GetBytes(payload, "Quantity").String()

	qtyInt, ok := new(big.Int).SetString(qty, 10)
	if !ok {
		return nil, errors.New("qty incorrect")
	}
	if _, ok = Balances[from]; !ok {
		return nil, errors.New("from incorrect")
	}
	if Balances[from].Cmp(qtyInt) == -1 {
		return nil, errors.New("insufficient balance")
	}
	Balances[from] = new(big.Int).Sub(Balances[from], qtyInt)

	if _, ok = Balances[recipient]; !ok {
		Balances[recipient] = big.NewInt(0)
	}
	Balances[recipient] = new(big.Int).Add(Balances[recipient], qtyInt)
	return []byte("transfer success"), nil
}

func Mint(payload []byte) ([]byte, error) {
	recipient := gjson.GetBytes(payload, "Recipient").String()
	qty := gjson.GetBytes(payload, "Quantity").String()
	qtyInt, ok := new(big.Int).SetString(qty, 10)
	if !ok {
		return nil, errors.New("qty incorrect")
	}
	if _, ok = Balances[recipient]; !ok {
		Balances[recipient] = big.NewInt(0)
	}
	Balances[recipient] = new(big.Int).Add(Balances[recipient], qtyInt)
	TotalSupply = new(big.Int).Add(TotalSupply, qtyInt)
	return []byte("mint success"), nil
}

func main() {
	wapc.RegisterFunctions(wapc.Functions{
		"Info":     Info,
		"Balance":  Balance,
		"Transfer": Transfer,
		"Mint":     Mint,
	})
}
