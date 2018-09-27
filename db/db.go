package db

import (
	"github.com/gocql/gocql"
	"log"
	"fmt"
)

var keyspaceCreationQuery = `
CREATE KEYSPACE IF NOT EXISTS moneway
WITH replication={'class': 'SimpleStrategy', 'replication_factor': 3}
`

var transactionBaseQuery = `
CREATE TABLE IF NOT EXISTS moneway.transactions (
    id int,
    accountid text,
    description text,
    amount int,
    currency text,
    notes text,
    createdat timeuuid,
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