package group

import (
	"fmt"
	"testing"
)

func TestCreateDeviceGroup(t *testing.T) {
	res, err := CreateDeviceGroup("1111", "2222", "3333", "4444")
	fmt.Println(err)
	fmt.Println(res)
}
