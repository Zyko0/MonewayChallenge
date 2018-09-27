package db

import (
	"github.com/gocql/gocql"
//	"github.com/scylladb/gocqlx"
//	"github.com/scylladb/gocqlx/qb"
	"log"
	"fmt"
)

var balanceBaseQuery = `
CREATE TABLE IF NOT EXISTS gocqlx_test.migrate_table (
    testint int,
    testuuid timeuuid,
    PRIMARY KEY(testint, testuuid)
)
`
var transactionBaseQuery = `
CREATE TABLE IF NOT EXISTS gocqlx_test.migrate_table (
    testint int,
    testuuid timeuuid,
    PRIMARY KEY(testint, testuuid)
)
`

func Init() {
	cluster := gocql.NewCluster("172.19.0.1","172.19.0.2","172.19.0.3","172.19.0.4")
	cluster.DisableInitialHostLookup = true
	cluster.ProtoVersion = 4
	cluster.CQLVersion = "5.0.1"
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CONNECTED")
	b := session.NewBatch(gocql.LoggedBatch)
	b.Query(balanceBaseQuery)
	b.Query(transactionBaseQuery)
	/*_, _, err = gocqlx.CompileNamedQuery([]byte(balanceBaseQuery))
	if err != nil {
		log.Fatal(err)
	}
	_, _, err = gocqlx.CompileNamedQuery([]byte(transactionBaseQuery))
	if err != nil {
		log.Fatal(err)
	}*/
}