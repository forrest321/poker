package betting

// PotManagement handles the pot and side pots
type PotManagement interface {
	PotTotal() int
	HandleSidePots() error
	DistributeWinnings() error
}
