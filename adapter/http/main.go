package http

import (
	"github.com/JWRibas/financialplannergo/adapter/http/actuater"
	"github.com/JWRibas/financialplannergo/adapter/http/transaction"
	"log"
	"net/http"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreatATransaction)

	http.HandleFunc("/health", actuater.Health)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
