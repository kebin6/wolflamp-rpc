package util

const PlayerMaxId = 10000000000

func IsRealPlayer(playerId uint64) bool {
	return playerId < PlayerMaxId
}
