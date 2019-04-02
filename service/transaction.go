package service

import (
	"encoding/json"
	"os"

	"github.com/newyu_transactions/models"
)

// Transaction is used to store the transaction details in a file
func Transaction(transactionDetails models.Transaction) error {

	// marshal details to JSON
	transactionDetailsJSON, err := json.Marshal(transactionDetails)
	if err != nil {
		return err
	}

	// append data into the file
	f, err := os.OpenFile("transaction.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write(transactionDetailsJSON); err != nil {
		return err
	}
	return f.Close()
}
