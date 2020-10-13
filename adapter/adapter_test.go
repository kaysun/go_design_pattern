// package adapter 适配器，以手机充电为例，适配器将220v电压转为5v电压
package adapter

import "testing"

func Test(t *testing.T) {
	t.Run("adaptor: ", ConvertVolts)
}

func ConvertVolts(t *testing.T) {
	var adaptee Adaptee
	adaptee = &Volts220{}

	var target Target
	target = &Adapter{adaptee}
	target.CovertTo5V()
}