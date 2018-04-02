package main

import (
	"fmt"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddOneBook(t *testing.T) {
	type args struct {
		title string
		auth  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing Huck Finn",
			args: args{
				title: "Huck Finn",
				auth:  "Mark Twain",
			},
		},
		{
			name: "testing Toom Sawyer",
			args: args{
				title: "Tom Sawyer",
				auth:  "Mark Twain",
			},
		},

		// TODO: Add test cases.
	}
	fmt.Println("trying Add OneBook")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenDB()

			AddOneBook(tt.args.title, tt.args.auth)
		})
	}
}

func TestReadTitleAuth(t *testing.T) {
	type ar struct {
		Ti string
		Au string
	}
	//var tt args
	//tt.Au = "Mark Twain"
	//tt.Ti = "Tom Sawyer"
	tests := []struct {
		name      string
		args      ar
		wantID    int
		wantTitle string
	}{
		{name: "Hik Fin",
			args:      ar{Ti: "Tom Sawyer", Au: "Mark Twain"},
			wantID:    2,
			wantTitle: "Tom Sawyer"},
		{name: "ddda Fin",
			args:      ar{Ti: "Huck Finn", Au: "Mark Twain"},
			wantID:    1,
			wantTitle: "Huck Finn"},
	}
	OpenDB()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, gotTitle, gotQty := ReadTitleAuth(tt.args.Ti, tt.args.Au)
			if gotID != tt.wantID {
				t.Errorf("ReadTitleAuth() gotID = %v, want %v", gotID, tt.wantID)
			}
			if gotTitle != tt.wantTitle {
				t.Errorf("ReadTitleAuth() gotTitle = %v, want %v", gotTitle, tt.wantTitle)
			}
			fmt.Println("title: " + gotTitle + ", qty: " + strconv.Itoa(gotQty))
		})
	}
}

func TestDeleteOneBook(t *testing.T) {
	type args struct {
		title string
		auth  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing Huck Finn",
			args: args{
				title: "Huck Finn",
				auth:  "Mark Twain",
			},
		},
		{
			name: "testing Toom Sawyer",
			args: args{
				title: "Tom Sawyer",
				auth:  "Mark Twain",
			},
		},

		// TODO: Add test cases.
	}
	fmt.Println("trying Delete OneBook")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenDB()

			res := DeleteOneBook(tt.args.title, tt.args.auth)
			fmt.Println(res)
		})
	}
}
