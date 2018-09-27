# MonewayChallenge

A challenge consisting of building a simple microservices application

## Usage

Run the docker container in order to execute all the microservices at once.
`docker-compose up`

Open a cqlsh connection with the scylladb instance to create the keyspace
`docker exec -it monewaychallenge_scylla_1 cqlsh`

Type in this query to create the keyspace 'moneway' :
```
CREATE KEYSPACE IF NOT EXISTS moneway
WITH replication={'class': 'SimpleStrategy', 'replication_factor': 3};
```

Then, from another terminal it is possible to store a transaction for example :
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"amount":32,"currency":"dollars","notes":"","accountid":"paul_h","description":"buying some expensive bread","id":1}' \
  http://localhost:8080/createtransaction
      
```