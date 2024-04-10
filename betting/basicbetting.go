package betting

// BasicBetting defines the basic actions of betting
type BasicBetting interface {
	PlaceBet(playerID, amount int) error
	Call(playerID int) error
	Fold(playerID int) error
	Check(playerID int) error
	CurrentBet() int
}
