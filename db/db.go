package db

import (
	"github.com/gocql/gocql"
	"log"
	"github.com/Zyko0/MonewayChallenge/transaction/pb"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx"
	"github.com/golang/protobuf/ptypes/timestamp"
	"math/rand"
)

type TransactionDB struct {
	id int64 `db:"id"`
	accountID string `db:"accountid"`
	createdAt timestamp.Timestamp `db:"createdat"`
	description string `db:"description"`
	amount int64 `db:"amount"`
	currency string `db:"currency"`
	notes string `db:"notes"`
}

var session gocql.Session

var keyspaceCreationQuery = `
CREATE KEYSPACE IF NOT EXISTS moneway
WITH replication={'class': 'SimpleStrategy', 'replication_factor': 3}
`

var transactionBaseQuery = `
CREATE TABLE IF NOT EXISTS moneway.transactions (
    id int,
    accountid text,
    createdat timeuuid,
    description text,
    amount int,
    currency text,
    notes text,
    PRIMARY KEY(id)
)
`
var balanceBaseQuery = `
CREATE TABLE IF NOT EXISTS moneway.balances (
    accountid text,
    amount int,
    currency text,
    PRIMARY KEY(accountid)
)
`

func StoreTransaction(req *pb.TransactionRequest) (*pb.TransactionRequest, error){
	if req.ID == -1 {
		req.ID = int64(rand.Intn(150000)) // Should be gocql.RandomUUID() but since i didnt set up any pb with strings as ids..
	}
	obj := &TransactionDB{
		id:req.ID,
		accountID:req.AccountID,
		createdAt:*req.CreatedAt,
		description:req.Description,
		amount:req.Amount,
		currency:req.Currency,
		notes:req.Notes,
	}
	stmt, names := qb.Insert("moneway.transactions").
		Columns("id", "accountid", "createdat", "description", "amount", "currency", "notes").
		ToCql()

	err := gocqlx.Query(session.Query(stmt), names).BindStruct(obj).ExecRelease()
	if err != nil {
		log.Fatal(err)
	}
	return req, nil
}

func Init() {
	cluster := gocql.NewCluster("172.19.0.1","172.19.0.2","172.19.0.3","172.19.0.4")
	cluster.Keyspace = "system"
	cluster.DisableInitialHostLookup = true
	cluster.ProtoVersion = 4
	cluster.CQLVersion = "5.0.1"
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	session.Query(transactionBaseQuery).Iter()
	session.Query(balanceBaseQuery).Iter()
}