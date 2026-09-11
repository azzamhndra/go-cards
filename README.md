# Go Cards Project

This is a simple Go project I built while learning the fundamentals of the language.

The project simulates a basic deck of cards and was mainly used to practice how Go works with custom types, methods, slices, file handling, randomization, and testing.

## What I Learned

During this project, I learned and practiced:

- Creating custom types in Go
- Using slices and appending values
- Writing functions with return values
- Using method receivers
- Looping with `range`
- Splitting slices
- Converting custom types to strings
- Reading and writing files with the `os` package
- Handling errors
- Shuffling data using `math/rand`
- Using callback functions
- Writing basic tests with the `testing` package
- Using `Errorf` and `Fatalf` in tests
- Cleaning up temporary test files

## Features

The project currently supports:

- Creating a new deck of cards
- Printing all cards
- Shuffling the deck
- Dealing part of the deck
- Converting the deck into a string
- Saving the deck to a file
- Loading a deck from a file
- Basic automated tests

## Example

```go
cards := newDeck()

cards.shuffle()
cards.print()
```

A deck can also be saved and loaded from a file:

```go
cards.saveToFile("my_cards")

cards := newDeckFromFile("my_cards")
```

## Testing

Run the tests with:

```bash
go test
```

The current tests check that a new deck has the expected number of cards and that saving and loading a deck works correctly.

## Notes

This project is part of my process of learning Go from the basics.

The main goal is not to build a complete card game, but to understand core Go concepts before moving on to more advanced topics and backend development.
