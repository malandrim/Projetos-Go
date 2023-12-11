package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/malandrim/Projetos-Go/tree/main/DesafiosGoExpert/Multithreading/internal/entity"
)

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	chViaCep := make(chan string)
	chBrasilAPI := make(chan string)

	go func() {
		BuscaViaCepHandler(w, r, chViaCep)
	}()
	go func() {
		BuscaBrasilApiHandler(w, r, chBrasilAPI)
	}()

	select {
	case msg := <-chViaCep:
		fmt.Println("Dados originados pela API:", msg)
	case msg := <-chBrasilAPI:
		fmt.Println("Dados originados pela API:", msg)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout")
	}
}

func BuscaViaCepHandler(w http.ResponseWriter, r *http.Request, ch chan string) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, error := BuscaViaCep(cepParam, ch)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cep)
	t, _ := json.Marshal(cep)
	fmt.Printf(string(t) + "\n")
}
func BuscaViaCep(cep string, ch chan string) (*entity.ViaCEP, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	ch <- "VIA CEP"
	var c entity.ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}
func BuscaBrasilApiHandler(w http.ResponseWriter, r *http.Request, ch chan string) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, error := BuscaBrasilApi(cepParam, ch)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cep)
	t, _ := json.Marshal(cep)
	fmt.Printf(string(t) + "\n")
}
func BuscaBrasilApi(cep string, ch chan string) (*entity.BrasilApiCep, error) {
	resp, error := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	ch <- "Brasil API"
	var c entity.BrasilApiCep
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}
