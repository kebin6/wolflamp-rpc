package util

func IsRealPlayer(playerId uint64) bool {
	return playerId < 10000000000
}
