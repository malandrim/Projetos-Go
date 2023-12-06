package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Cotacao struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type CotacaoBid struct {
	Bid string `json:"bid"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", CotacaoHandler)
	http.ListenAndServe(":8080", mux)
}

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctxConsulta := r.Context()
	log.Println("Request Iniciada...")
	ctxConsulta, cancelConsulta := context.WithTimeout(ctxConsulta, 200*time.Millisecond)
	defer cancelConsulta()

	parMoedas, varBid, err := BuscaCotacao(ctxConsulta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	select {
	case <-ctxConsulta.Done():
		log.Println("Request Cancelada pelo Cliente!")
	default:
		log.Println("Consulta realizada!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(varBid)

		//Chamando funcao para criar tabela caso nao exista (para ambiente dev)
		criarTabela()

		ctxSalva := context.Background()
		ctxSalva, cancelRegistro := context.WithTimeout(ctxSalva, 10*time.Millisecond)
		defer cancelRegistro()

		select {
		case <-ctxSalva.Done():
			log.Println("Excedido o limite de tempo de gravacao no banco de dados!")
		default:
			salvaCotacao(ctxSalva, parMoedas)
			log.Println("Gravacao no banco de dados realizada!")
		}
	}

	log.Println("Request Finalizada...")
}

func BuscaCotacao(ctx context.Context) (*Cotacao, *CotacaoBid, error) {
	resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		log.Println("erro busca cotacao")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("erro busca cotacao-2")
	}
	var c Cotacao
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Println("erro busca cotacao-3")
	}

	var cb CotacaoBid
	cb.Bid = c.USDBRL.Bid

	return &c, &cb, nil
}

func criarTabela() {
	db, err := sql.Open("sqlite3", "cotacoes.db")
	if err != nil {
		log.Println(err)
	}
	// criando a tabela tb_cotacoes na primeira execucao do codigo
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS tbcotacoes (id INTEGER PRIMARY KEY, code VARCHAR(64), codein VARCHAR(64), name	VARCHAR(64), high float(10,4), low float(10,4), varBid	float(10,4), pctChange	float(10,4), bid float(10,4), ask	float(10,4), timestamp	VARCHAR(64), create_date	VARCHAR(64))")
	if err != nil {
		log.Println("Erro ao criar tabela")
	}
	statement.Exec()
}

func salvaCotacao(ctx context.Context, c *Cotacao) {
	db, err := sql.Open("sqlite3", "cotacoes.db")
	if err != nil {
		log.Println("Falha ao abrir conexao com db")
	}

	stmt, err := db.Prepare("insert into tbcotacoes(code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Falha na criacao do statement ")
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, c.USDBRL.Code, c.USDBRL.Codein, c.USDBRL.Name, c.USDBRL.High, c.USDBRL.Low, c.USDBRL.VarBid, c.USDBRL.PctChange, c.USDBRL.Bid, c.USDBRL.Ask, c.USDBRL.Timestamp, c.USDBRL.CreateDate)

	if err != nil {
		log.Println("Falha ao executar a gravacao no bd!")
	}
}
