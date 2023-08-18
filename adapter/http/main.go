package http

import (
	"github.com/JWRibas/financialplannergo/adapter/http/actuater"
	"github.com/JWRibas/financialplannergo/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreatATransaction)

	http.HandleFunc("/health", actuater.Health)

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))

}
