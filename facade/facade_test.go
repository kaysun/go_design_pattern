package facade

import "testing"

func Test(t *testing.T) {
	t.Run("facade", RegisterUser)
}

func RegisterUser(t *testing.T) {
	facade := RegisterFacade{}
	facade.Register("18618193858")
}
