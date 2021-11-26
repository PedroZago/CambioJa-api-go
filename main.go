package main

import (
	"api/src/configs"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	configs.CarregarBD()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", configs.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configs.Porta), r))
}
