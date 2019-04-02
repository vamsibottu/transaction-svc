package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/newyu_transactions/models"
	"github.com/newyu_transactions/service"
)

// transaction is used to validate the inputs are return the response
func transaction(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	transactionDetails := models.Transaction{}

	if err := json.Unmarshal(body, &transactionDetails); err != nil {
		// logg is used to log error message into the console, that way we can check in ELK logs when deployed
		logg().Infof("error unmarshaling the input body: %s", err)
		// don't need to handle error message for marshal, as we are passing static value
		errResp, _ := json.Marshal(models.ErrorResponse{Message: "error unmarshaling the input body", Code: http.StatusBadRequest})

		http.Error(w, string(errResp), 400)
		return
	}

	if err := service.Transaction(transactionDetails); err != nil {
		// logg is used to log error message into the console, that way we can check in ELK logs when deployed
		logg().Infof("error storing transaction details: %s", err)
		// don't need to handle error message for marshal, as we are passing static value
		errResp, _ := json.Marshal(models.ErrorResponse{Message: "error storing transaction details", Code: http.StatusInternalServerError})

		http.Error(w, string(errResp), 400)
		return
	}
}
