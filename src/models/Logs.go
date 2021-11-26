package models

import (
	"fmt"
	"os"
)

const FILE string = "logs.txt"

type Logs struct {
	Metodo string
	URI    string
	Host   string
	SQL    string
}

func (log *Logs) SalvarLog() error {
	arquivo, err := os.OpenFile(FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	logs := fmt.Sprintf("%s || %s || %s || %s \n", log.Metodo, log.URI, log.Host, log.SQL)

	if _, err := arquivo.WriteString(logs); err != nil {
		return err
	}

	return nil
}

// var gerenciadorLog models.Logs
// gerenciadorLog.Metodo = r.Method
// gerenciadorLog.URI = r.RequestURI
// gerenciadorLog.Host = r.Host
// gerenciadorLog.SQL = ""
// _ = gerenciadorLog.SalvarLog()
