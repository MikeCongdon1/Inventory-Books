package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//based on https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.2.html
// globals
var DBCon *sql.DB // Note the sql package provides the namespace

func create(name string) {

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
func addOne(title string, auth string) {
	ti1, au1, qty1 := readTitleAuth("Huck Finn", "Mark Twain")

}
func addSome() {
	statement, err := DBCon.Prepare("INSERT bookInv SET title=?,author=?,qty=?,modified=?")
	errcheck(err)

	t := time.Now()
	th := t.Format("2006-01-02 15:04:05")

	res, err := statement.Exec("Huck Finn", "Mark Twain", 1, th)
	errcheck(err)

	fmt.Println(res)
	rows, err := DBCon.Query("SELECT * FROM bookInv")
	errcheck(err)
	for rows.Next() {
		var uid int
		var title string
		var author string
		var qty int
		var modified string

		err = rows.Scan(&uid, &title, &author, &qty, &modified)
		errcheck(err)

		fmt.Println(uid)
		fmt.Println(title)
		fmt.Println(author)
		fmt.Println(modified)
	}
	//readTitleAuth()
	/*
	   Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	   		checkErr(err)
	*/
}

func readTitleAuth(Ti string, Au string) { //(title string, auth string, qty int) {
	rows, err := DBCon.Query("SELECT id,title, author,qty FROM bookInv where title=? && author=? limit 1", Ti, Au)
	errcheck(err)
	for rows.Next() {
		var id int
		var title string
		var author string
		var qty int
		//var modified string

		err = rows.Scan(&id, &title, &author, &qty)
		errcheck(err)

		fmt.Println("ID: " + strconv.Itoa(id) + " Ti:" + title + " au: " + " qty: " + strconv.Itoa(qty))
	}
}

func update() {

	/*
		id, err := res.LastInsertId()
		errcheck(err)

		// update

		statement, err = DBCon.Prepare("update bookInv set title=? where id=?")
		errcheck(err)

		res, err = statement.Exec("Tom Sawyer", id)
		errcheck(err)

		affect, err := res.RowsAffected()
		errcheck(err)

		fmt.Println(affect)

		rows, err := DBCon.Query("SELECT * FROM bookInv")
		errcheck(err)
	*/

}

func main() {
	var err error
	DBCon, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")

	errcheck(err)
	defer DBCon.Close()
	name := "books"

	fmt.Println(sql.Drivers())

	create(name)
	addSome()

}

func errcheck(er error) {
	if er != nil {
		panic(er)
	}
}
