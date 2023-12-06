package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type CotacaoDolar struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	salvaArquivo(body)
	defer res.Body.Close()
}

func salvaArquivo(body []byte) {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	var c CotacaoDolar
	err = json.Unmarshal(body, &c)
	valor := c.Bid
	_, err = f.WriteString(fmt.Sprintf("Dólar:{%v}", valor))
	fmt.Printf("Arquivo criado com sucesso... \n")
	f.Close()
}
