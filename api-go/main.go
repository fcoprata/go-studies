package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Print("API GO STARTING")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Obter o horário atual
		currentTime := time.Now()

		// Ajustar o horário para o fuso horário de Brasília
		brasiliaTime := currentTime.In(time.FixedZone("Brasília", -3*60*60))

		// Retornar o horário no formato "2006-01-02 15:04:05"
		fmt.Fprint(w, brasiliaTime.Format("2006-01-02 15:04:05"))
	})

	http.ListenAndServe(":8000", nil)
}
