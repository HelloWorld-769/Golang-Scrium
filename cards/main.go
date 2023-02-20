package main

func main() {
	cards := deck{"Ace of diamond", newCard()}
	cards = append(cards, "Six of spades")

	cards.print()
}
func newCard() string {
	return "Five of diamonds"
}
