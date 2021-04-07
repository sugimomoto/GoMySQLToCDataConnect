package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Connect to CData Connect by MySQL Wire protocol")
	db, err := sql.Open("mysql", "USER_NAME:PASSWORD@tcp(TENANT_NAME.cdatacloud.net:3306)/Salesforce1?tls=true")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT Id,Name from Account limit 10")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var account Account
		rows.Scan(&account.Id, &account.Name)

		fmt.Printf("%s \n", account)
	}
}

type Account struct {
	Id   string
	Name string
}
