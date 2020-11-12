package state

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("state", StateUser)
}

func StateUser(t *testing.T) {
	user := User{}
	user.WatchVideo()

	user.PurchaseVip()
	user.WatchVideo()

	user.Expire()
	user.WatchVideo()
}
