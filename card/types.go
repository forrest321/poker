package card

type Suit string

const (
	Spades   Suit = "Spades"
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
)

type Value struct {
	NumericValues []int
	Name          string
}

func NewValue(numericValues []int, name string) Value {
	return Value{NumericValues: numericValues, Name: name}
}

func (v Value) String() string {
	return v.Name
}

var (
	Two   = NewValue([]int{2}, "2")
	Three = NewValue([]int{3}, "3")
	Four  = NewValue([]int{4}, "4")
	Five  = NewValue([]int{5}, "5")
	Six   = NewValue([]int{6}, "6")
	Seven = NewValue([]int{7}, "7")
	Eight = NewValue([]int{8}, "8")
	Nine  = NewValue([]int{3}, "9")
	Ten   = NewValue([]int{10}, "10")
	Jack  = NewValue([]int{11}, "Jack")
	Queen = NewValue([]int{12}, "Queen")
	King  = NewValue([]int{13}, "King")
	Ace   = NewValue([]int{1, 14}, "Ace") // Ace can be 1 or 14
)
