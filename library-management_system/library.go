package main

import (
	"fmt"
	"sync"
)

var (
	libraryInstance *Library
	once            sync.Once
)

type Library struct {
	books   map[int]*Book
	members map[int]*Member
}

func GetLibraryInstance() *Library {
	once.Do(func() {
		libraryInstance = &Library{books: make(map[int]*Book), members: make(map[int]*Member)}
	})
	return libraryInstance
}

func (l *Library) AddBook(book *Book) {
	l.books[book.ID] = book
	fmt.Printf("Book %d has been added\n", book.ID)
}

func (l *Library) RemoveBook(id int) {
	delete(l.books, id)
	fmt.Printf("Book %d has been removed\n", id)
}

func (l *Library) AddMember(member *Member) {
	l.members[member.ID] = member
	fmt.Printf("Member %d has been added\n", member.ID)
}

func (l *Library) RemoveMember(id int) {
	delete(l.members, id)
	fmt.Printf("Member %d has been removed\n", id)
}

func (l *Library) BorrowBookByMember(memberID int, bookID int) (*BookItem, error) {
	if l.members[memberID] == nil || l.books[bookID] == nil {
		return nil, fmt.Errorf("Member or book not found")
	}

	member := l.members[memberID]

	if member.IsQuotaFull() {
		return nil, fmt.Errorf("Member %d has reached their borrowing quota\n", memberID)
	}

	book := l.books[bookID]

	if !book.IsBookAvailable() {
		return nil, fmt.Errorf("Book %d is not available\n", bookID)
	}

	borrowedBook := book.BorrowBook()

	member.AddBorrowedBook(borrowedBook)
	fmt.Printf("Member %d has borrowed book %d with item id %d\n", memberID, bookID, borrowedBook.ID)
	return borrowedBook, nil
}

func (l *Library) ReturnBookByMember(memberID int, bookItemID int) {
	member := l.members[memberID]

	for _, book := range member.CurrentBorrowed {
		if book.ID == bookItemID {
			book.ReturnBook()
			member.RemoveBorrowedBook(book)
			fmt.Printf("Member %d has returned book %d with item id %d\n", memberID, book.BookID, bookItemID)
			return
		}
	}
}

func (l *Library) DisplayAvailableBooks() {
	fmt.Println("Available Books:")
	for id, book := range l.books {
		if book.IsBookAvailable() {
			fmt.Printf("Book %d: %s by %s (published in %s)\n", id, book.Title, book.Author, book.PublishedYear)
		}
	}
}
