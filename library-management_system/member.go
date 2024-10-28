package main

import "fmt"

type Member struct {
	ID              int
	Name            string
	ContactInfo     string
	CurrentBorrowed []*BookItem
	BorrowHistory   []*BookItem
}

func NewMember(id int, name, contactInfo string) *Member {
	return &Member{ID: id, Name: name, ContactInfo: contactInfo, CurrentBorrowed: make([]*BookItem, 0), BorrowHistory: make([]*BookItem, 0)}
}

func (m *Member) IsQuotaFull() bool {
	return len(m.CurrentBorrowed) >= 3
}

func (m *Member) AddBorrowedBook(bookItem *BookItem) {
	m.CurrentBorrowed = append(m.CurrentBorrowed, bookItem)
}

func (m *Member) RemoveBorrowedBook(bookItem *BookItem) {
	for i, bi := range m.CurrentBorrowed {
		if bi.ID == bookItem.ID {
			m.CurrentBorrowed = append(m.CurrentBorrowed[:i], m.CurrentBorrowed[i+1:]...)
			break
		}
	}
	m.BorrowHistory = append(m.BorrowHistory, bookItem)
	fmt.Printf("Book %d has been added to borrow history for Member %s\n", bookItem.ID, m.Name)
}

func (m *Member) DisplayCurrentBorrowedBooks() {
	fmt.Printf("Current borrowed books for Member %s:\n", m.Name)
	for _, bi := range m.CurrentBorrowed {
		fmt.Printf(" - Book %d (Item ID: %d)\n", bi.BookID, bi.ID)
	}
}

func (m *Member) DisplayBorrowHistory() {
	fmt.Printf("Borrow history for Member %s:\n", m.Name)
	for _, bi := range m.BorrowHistory {
		fmt.Printf(" - Book %d (Item ID: %d)\n", bi.BookID, bi.ID)
	}
}
