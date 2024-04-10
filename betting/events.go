package betting

// Events defines optional event handling in the betting process
type Events interface {
	OnBetPlaced(playerID, amount int)
	OnRoundStarted()
	OnRoundEnded()
	OnPlayerFolded(playerID int)
}
