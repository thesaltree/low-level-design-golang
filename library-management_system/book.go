package main

import (
	"sync"
)

type Book struct {
	ID            int
	BookItem      []BookItem
	Title         string
	Author        string
	PublishedYear string
	mu            sync.RWMutex
}

func NewBook(id int, title, author string, publishedYear string) *Book {
	bookCopies := &Book{ID: id, Title: title, Author: author, PublishedYear: publishedYear, BookItem: make([]BookItem, 0)}
	for i := 1; i <= 10; i++ {
		bookCopies.BookItem = append(bookCopies.BookItem, *NewBookItem(i, id))
	}
	return bookCopies
}

func (b *Book) IsBookAvailable() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, bookCopy := range b.BookItem {
		if bookCopy.Status == Available {

			return true
		}
	}
	return false
}

func (b *Book) BorrowBook() *BookItem {
	b.mu.Lock()
	defer b.mu.Unlock()
	for i := range b.BookItem {
		if b.BookItem[i].Status == Available {
			b.BookItem[i].BorrowBook()
			return &b.BookItem[i]
		}
	}
	return nil
}
