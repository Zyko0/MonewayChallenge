package db

import (
	"github.com/gocql/gocql"
	"log"
	"fmt"
	"github.com/Zyko0/MonewayChallenge/transaction/pb"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx"
)

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

func StoreTransaction(req *pb.TransactionRequest) {
	stmt, names := qb.Insert("moneway.transactions").
		Columns("id", "accountid", "createdat", "description", "amount", "currency", "notes").
		ToCql()

	err := gocqlx.Query(session.Query(stmt), names).BindStruct(req).ExecRelease()
	if err != nil {
		log.Fatal(err)
	}
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
	fmt.Println("CONNECTED")
	session.Query(transactionBaseQuery).Iter()
	session.Query(balanceBaseQuery).Iter()
}