package betting

// Validation provides methods for validating betting actions
type Validation interface {
	IsValidBet(playerID, amount int) bool
	CanCheck(playerID int) bool
	CanCall(playerID int) bool
	CanRaise(playerID int) bool
}
