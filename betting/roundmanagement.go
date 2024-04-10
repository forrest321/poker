package betting

// RoundManagement handles the start and end of betting rounds
type RoundManagement interface {
	StartRound()
	EndRound()
}
