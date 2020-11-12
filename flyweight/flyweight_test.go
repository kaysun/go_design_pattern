package flyweight

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("flyweight", Flyweight)
}

func Flyweight(t *testing.T) {
	board1 := ChessBoard{}
	board1.Init()
	board1.Move(1, 10, 20)

	board2 := ChessBoard{}
	board2.Init()
	board2.Move(2, 6, 6)
}
