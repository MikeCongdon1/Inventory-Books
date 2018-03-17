package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DBCon based on https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.2.html
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

func addOneBook(title string, auth string) {

	id1, ti1, qty1 := readTitleAuth(title, auth)

	fmt.Println("ID: " + strconv.Itoa(id1) + " Ti:" + ti1 + " qty: " + strconv.Itoa(qty1))

	if ti1 != "" {
		statement, err := DBCon.Prepare("UPDATE bookInv SET qty=?,modified=? where id=?")
		errcheck(err)
		res, err := statement.Exec(qty1+1, getTime(), id1)
		errcheck(err)
		fmt.Println(res)

	} else {
		statement, err := DBCon.Prepare("INSERT bookInv SET title=?,author=?,qty=?,modified=?")
		errcheck(err)

		res, err := statement.Exec(title, auth, 1, getTime())
		errcheck(err)
		fmt.Println(res)

	}

}
func concatTitleAuth(Ti string, Au string) {
	var QtyLoc int
	var ID1st int
	var IDList []int
	var Title1st string
	rows, err := DBCon.Query("SELECT * FROM bookInv where title=? && author=?", Ti, Au)
	errcheck(err)
	if rows != nil {
		for rows.Next() {
			var Id int
			var Title string
			var Author string
			var Qty int
			var Modified string

			err = rows.Scan(&Id, &Title, &Author, &Qty, &Modified)
			errcheck(err)

			fmt.Println(Id)
			fmt.Println(Title)
			fmt.Println(Author)
			fmt.Println(Modified)
			QtyLoc += Qty
			if ID1st == 0 {
				ID1st = Id
			}
			IDList = append(IDList, Id)
			fmt.Println(IDList)
			if Title == "" {
				Title1st = Title
			}
		}
		// TODO something is wrong with the Concat qty builder
		fmt.Println("UPDATE bookInv SET " + strconv.Itoa(QtyLoc) + " Ti " + Title1st)
		statement, err := DBCon.Prepare("UPDATE bookInv SET qty=?,modified=? where id=?")
		errcheck(err)
		_, err = statement.Exec(QtyLoc, getTime(), ID1st)
		errcheck(err)
		IDList = append(IDList[:0], IDList[0+1:]...) // deletes first
		for i := 0; i < len(IDList); i++ {
			statement, err := DBCon.Prepare("Delete from bookInv where id=?")
			errcheck(err)
			_, err = statement.Exec(IDList[i])
			errcheck(err)
		}

	} else {
		fmt.Println("Concat not needed")
	}

}
func getTime() (th string) {
	t := time.Now()
	th = t.Format("2006-01-02 15:04:05")
	return th
}
func readTitleAuth(Ti string, Au string) (Id int, Title string, Qty int) {
	concatTitleAuth(Ti, Au)
	/*rows, err := DBCon.Query("SELECT id,title, author,qty FROM bookInv where title=? && author=? limit 1", Ti, Au)
	fmt.Println(&rows)
	errcheck(err)
	for rows.Next() {
		var id int
		var title string
		var author string
		var qty int
		//var modified string

		err = rows.Scan(&id, &title, &author, &qty)
		errcheck(err)

	}*/
	rows, err := DBCon.Query("SELECT * FROM bookInv where title=? && author=? limit 1", Ti, Au)
	errcheck(err)
	for rows.Next() {
		var Id int
		var Title string
		var Author string
		var Qty int
		var Modified string

		err = rows.Scan(&Id, &Title, &Author, &Qty, &Modified)
		errcheck(err)

		fmt.Println(Id)
		fmt.Println(Title)
		fmt.Println(Author)
		fmt.Println(Modified)
		return Id, Title, Qty

	}
	return Id, Title, Qty

}

func addSome() {

	addOneBook("Huck Finn", "Mark Twain")
	/*
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
		}*/
	//readTitleAuth()
	/*
	   Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	   		checkErr(err)
	*/
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
