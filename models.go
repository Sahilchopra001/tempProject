package db

import "time"

// Library represents the Library table in the database
type Library struct {
	ID   int
	Name string
}

// User represents the Users table in the database
type User struct {
	ID            int
	Name          string
	Email         string
	ContactNumber string
	Role          string
	LibID         int
}

// BookInventory represents the Bookinventory table in the database
type BookInventory struct {
	ISBN            string
	LibID           int
	Title           string
	Authors         string
	Publisher       string
	Version         string
	TotalCopies     int
	AvailableCopies int
}

// RequestEvent represents the RequestEvents table in the database
type RequestEvent struct {
	ReqID        int
	BookID       string
	ReaderID     int
	RequestDate  time.Time
	ApprovalDate time.Time
	ApproverID   int
	RequestType  string
}

// IssueRegistry represents the IssueRegistry table in the database
type IssueRegistry struct {
	IssueID            int
	ISBN               string
	ReaderID           int
	IssueApproverID    int
	IssueStatus        string
	IssueDate          time.Time
	ExpectedReturnDate time.Time
	ReturnDate         time.Time
	ReturnApproverID   int
}
