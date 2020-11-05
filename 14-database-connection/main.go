package main

import "fmt"

type Connector interface {
	Connect()
}

type SQLConnection struct {
	conString string
}

func (s SQLConnection) Connect() {
	fmt.Println("SQL " + s.conString)
}

type OracleConnection struct {
	conString string
}

func (s OracleConnection) Connect() {
	fmt.Println("Oracle " + s.conString)
}

type DBConnection struct {
	db Connector
}

func (c DBConnection) DBConnect() {
	c.db.Connect()
}

func main() {

	sqlConnection := SQLConnection{conString: "Connection is connected"}
	con1 := DBConnection{db: sqlConnection}
	con1.DBConnect()

	orcConnection := OracleConnection{conString: "Connection is Connected"}
	con2 := DBConnection{db: orcConnection}
	con2.DBConnect()

}
