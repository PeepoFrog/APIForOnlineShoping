package helper

import (
	"fmt"
	"testing"
)

func TestGetOneUser(t *testing.T) {
	user := GetOneUser("632d91334c76193971356477")
	fmt.Println(user)
	if user == nil {
		fmt.Println("error")
	}
}
