package main

type Status string

const (
	Borrowed  Status = "Borrowed"
	Available Status = "Available"
)

type BookItem struct {
	ID     int
	BookID int
	Status Status
}

func NewBookItem(id, bookID int) *BookItem {
	return &BookItem{ID: id, BookID: bookID, Status: Available}
}

func (bi *BookItem) BorrowBook() {
	bi.Status = Borrowed
}

func (bi *BookItem) ReturnBook() {
	bi.Status = Available
}
