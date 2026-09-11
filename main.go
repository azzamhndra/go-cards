package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println("New deck:")
	cards.print()

	fmt.Println("\nDeck as string:")
	fmt.Println(cards.toString())

	cards.shuffle()

	fmt.Println("\nShuffled deck:")
	cards.print()

	hand, remainingDeck := cards.deal(5)

	fmt.Println("\nHand:")
	hand.print()

	fmt.Println("\nRemaining deck:")
	remainingDeck.print()

	filename := "my_cards"

	err := cards.saveToFile(filename)
	if err != nil {
		fmt.Println("Error saving deck:", err)
		return
	}

	fmt.Println("\nDeck saved to:", filename)

	loadedDeck := newDeckFromFile(filename)

	fmt.Println("\nLoaded deck:")
	loadedDeck.print()
}
