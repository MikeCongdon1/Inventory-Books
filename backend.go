package main

import (
	"database/sql"
	"fmt"
	"time"
)

func Create(name string) {

	_, err := DBCon.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	errcheck(err)
	DBCon.Close()

	DBCon, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/"+name)
	errcheck(err)

	_, err = DBCon.Exec("USE " + name)
	errcheck(err)

	_, err = DBCon.Exec("CREATE TABLE IF NOT EXISTS bookInv ( id MEDIUMINT NOT NULL AUTO_INCREMENT, title varchar(32), author varchar(32), qty MEDIUMINT, modified varchar(32), Primary Key (id) )")
	errcheck(err)

}

func OpenDB() {
	// comment

	var err error
	DBCon, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	errcheck(err)
	name := "books"
	Create(name)
	fmt.Println(sql.Drivers())
}

func getTime() (th string) {
	t := time.Now()
	th = t.Format("2006-01-02 15:04:05")
	return th
}
