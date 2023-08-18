package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/JWRibas/financialplannergo/model/transaction"
	"github.com/JWRibas/financialplannergo/util"
	"io"
	"net/http"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:   "Salário",
			Amount:  1200.0,
			Type:    0,
			CreatAt: util.StringToDate("2023-05-01T15:40:00"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreatATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = io.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)

}
