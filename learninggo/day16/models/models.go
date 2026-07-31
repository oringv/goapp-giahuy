package models

import "time"

type Book struct {
	Id         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Borrower struct {
	Id    string
	Name  string
	Email string
}

type BorrowRecord struct {
	BookId     string
	BorrowerId string
	BorrowDate time.Time
}
