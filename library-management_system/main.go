package main

import "fmt"

func main() {
	library := GetLibraryInstance()

	book1 := NewBook(1, "Book 1", "John Doe", "2000")
	book2 := NewBook(2, "Book 2", "Jane Doe", "2005")
	book3 := NewBook(3, "Book 3", "John Doe", "2010")

	library.AddBook(book1)
	library.AddBook(book2)
	library.AddBook(book3)

	library.DisplayAvailableBooks()

	member1 := NewMember(1, "Craig Bob", "123-456-7890")
	member2 := NewMember(2, "Alice Johnson", "987-654-3210")

	library.AddMember(member1)
	library.AddMember(member2)

	user1borrow1, err := library.BorrowBookByMember(member1.ID, book1.ID)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	}
	user1borrow2, err := library.BorrowBookByMember(member1.ID, book2.ID)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	}
	user2borrow1, err := library.BorrowBookByMember(member2.ID, book1.ID)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	}

	member1.DisplayCurrentBorrowedBooks()
	member2.DisplayCurrentBorrowedBooks()

	library.DisplayAvailableBooks()

	library.ReturnBookByMember(member1.ID, user1borrow1.ID)
	library.ReturnBookByMember(member2.ID, user2borrow1.ID)
	library.ReturnBookByMember(member1.ID, user1borrow2.ID)

	member1.DisplayBorrowHistory()
	member2.DisplayBorrowHistory()

	library.DisplayAvailableBooks()
}
