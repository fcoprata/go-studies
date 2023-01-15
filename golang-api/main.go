package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Data struct {
	Name string `json:"name"`
	Age  int    `json::"age"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "<Método não permitido>", http.StatusMethodNotAllowed)
		fmt.Print("<Método não permitido>")
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "<Tipo de conteúdo não suportado>", http.StatusUnsupportedMediaType)
		return
	}

	data := Data{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "<Erro ao ler dados de entrada>", http.StatusBadRequest)
		return
	}

	json_data := fmt.Sprintf("{Nome:%v, Age:%v, Datetime:%v", data.Name, data.Age, time.Now().Format(time.RFC822))
	fmt.Printf(json_data)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<Dados recebidos com sucesso!>\n"))
}

func main() {
	fmt.Print("<GO API STARTING>\n")

	http.HandleFunc("/data", apiHandler)
	http.ListenAndServe(":8000", nil)
}
