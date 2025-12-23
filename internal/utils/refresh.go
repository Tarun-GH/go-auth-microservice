package utils

import (
	"sync"

	"github.com/google/uuid"
)

var refreshStore = struct {
	sync.Mutex
	data map[string]int
}{
	data: make(map[string]int),
}

func GenerateRefresh(userID int) string {
	refreshToken := uuid.NewString()

	refreshStore.Lock()
	refreshStore.data[refreshToken] = userID
	refreshStore.Unlock()

	return refreshToken
}

func GetUserIDFromRefreshToken(token string) (int, bool) {
	refreshStore.Lock()
	defer refreshStore.Unlock()

	userID, ok := refreshStore.data[token]
	return userID, ok
}
