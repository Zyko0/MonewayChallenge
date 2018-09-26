package main

import(
	"net/http"
	"log"

	"google.golang.org/grpc"
	transaction "github.com/Zyko0/MonewayChallenge/transaction/pb"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"fmt"
)

type TransactionJSON struct {
	ID     int64 `json:"id"`
	AccountID string `json:"accountid"`
	Description string `json:"description"`
	Amount int64 `json:"amount"`
	Currency string `json:"currency"`
	Notes string `json:"notes"`
}

var transactionClient transaction.TransactionServiceClient

func balance(c *gin.Context) {

}

func credit(c *gin.Context) {

}

func debit(c *gin.Context) {

}

func storeTransaction(c *gin.Context) {
	transactionJson := TransactionJSON{}
	req := &transaction.TransactionRequest{}
	err := c.BindJSON(&transactionJson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not parse transaction data.")
	}
	storedMethod := "stored"
	if transactionJson.ID != -1 {
		storedMethod = "updated"
	}
	req.ID = transactionJson.ID
	req.AccountID = transactionJson.AccountID
	req.Amount = transactionJson.Amount
	req.Currency = transactionJson.Currency
	req.Description = transactionJson.Description
	req.Notes = transactionJson.Notes
	req.CreatedAt = ptypes.TimestampNow()
	res, err := transactionClient.StoreTransaction(c, req)
	if err == nil && res.Completed {
		c.JSON(http.StatusOK, fmt.Sprintf("Transaction correclty %s.", storedMethod))
	} else {
		c.JSON(http.StatusInternalServerError, "Could not process the transaction.")
	}
}

func main() {
	r := gin.Default()
	r.POST("/createtransaction", storeTransaction)
	r.POST("/updatetransaction", storeTransaction)
	r.GET("/balance", balance)
	r.POST("/credit", credit)
	r.POST("/debit", debit)
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	transactionClient = transaction.NewTransactionServiceClient(conn)
	r.Run(":8080")
}
