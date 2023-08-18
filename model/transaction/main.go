package transaction

import "time"

type Transaction struct {
	Title   string    `json:"title"`
	Amount  float32   `json:"amount"`
	Type    int       `json:"type"`
	CreatAt time.Time `json:"creatAt"`
}

type Transactions []Transaction
