package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// DBCon based on https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.2.html
// globals
var DBCon *sql.DB // Note the sql package provides the namespace
type Result sql.Result

func AddOneBook(title string, auth string) Result {
	id1, ti1, qty1 := ReadTitleAuth(title, auth)

	fmt.Println("ID: " + strconv.Itoa(id1) + " Ti:" + ti1 + " qty: " + strconv.Itoa(qty1))

	if ti1 != "" {
		statement, err := DBCon.Prepare("UPDATE bookInv SET qty=?,modified=? where id=?")
		errcheck(err)
		res, err := statement.Exec(qty1+1, getTime(), id1)
		errcheck(err)
		return res

	} else {
		statement, err := DBCon.Prepare("INSERT bookInv SET title=?,author=?,qty=?,modified=?")
		errcheck(err)

		res, err := statement.Exec(title, auth, 1, getTime())
		errcheck(err)
		return res

	}
}
func DeleteOneBook(title string, auth string) Result {
	//	comment
	fmt.Println("trying Add OneBook")
	id1, ti1, qty1 := ReadTitleAuth(title, auth)

	fmt.Println("ID: " + strconv.Itoa(id1) + " Ti:" + ti1 + " qty: " + strconv.Itoa(qty1))

	if ti1 != "" {
		// check that qty1 is greater than 1:
		// in the future if it's 1, we'll delete the entry instead of updating. for now, we'll just subtract
		statement, err := DBCon.Prepare("UPDATE bookInv SET qty=?,modified=? where id=?")
		errcheck(err)
		qty1--
		res, err := statement.Exec(qty1, getTime(), id1)
		errcheck(err)
		return res

	} else {

		statement, err := DBCon.Prepare("INSERT bookInv SET title=?,author=?,qty=?,modified=?")
		errcheck(err)
		res, err := statement.Exec(title, auth, -1, getTime())
		errcheck(err)
		fmt.Println("Book Did Not Exist in System!!!  ")
		return res

	}
}

func concatTitleAuth(Ti string, Au string) {
	var QtyLoc int
	var ID1st int
	var IDList []int
	var Title1st string
	//DBCon.Stats
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

func ReadTitleAuth(Ti string, Au string) (Id int, Title string, Qty int) {

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

// comment

func main() {
	OpenDB()

	addSome()

	defer DBCon.Close()

}

func errcheck(er error) {
	if er != nil {
		panic(er)
	}
}
