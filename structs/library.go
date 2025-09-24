package main

import "fmt"

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

func (b *Book) Borrow() bool {
	if !b.Available {
		return false
	}

	b.Available = false
	return true
}

func (b *Book) Return() bool {
	if b.Available {
		return false
	}

	b.Available = true
	return true
}

func (b *Book) String() string {
	status := "issued"

	if b.Available {
		status = "available"
	}

	return fmt.Sprintf("Title: %s | Author: %s | ISBN: %s | Status: %s", b.Title, b.Author, b.ISBN, status)
}

func main() {
	book := &Book{
		Title:     "War and Peace",
		Author:    "Lev Tolstoy",
		ISBN:      "3",
		Available: true,
	}

	fmt.Println(book)

	if book.Borrow() {
		fmt.Println("You took a book for yourself")
	} else {
		fmt.Println("The book has already been taken")
	}

	fmt.Println(book)

	if book.Return() {
		fmt.Println("You have successfully returned the book")
	} else {
		fmt.Println("The book is already in the library")
	}
	fmt.Println(book)
}
