package betting

// AdvancedBetting extends BasicBetting with advanced betting functionalities
type AdvancedBetting interface {
	BasicBetting
	Raise(playerID, amount int) error
	AllIn(playerID int) error
	MinimumRaise() int
}
