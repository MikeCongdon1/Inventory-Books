package main

import (
	"fmt"
)

func addSome() {

	stat := AddOneBook("Huck Finn", "Mark Twain")
	fmt.Println(stat)
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
