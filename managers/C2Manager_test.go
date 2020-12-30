package managers

import (
	"fmt"
	"testing"
)

func TestC2Manager_hashPassword(t *testing.T) {
	var pw = "tester123"
	var sm C2Manager
	suc, hpw := sm.hashPassword(pw)
	fmt.Println("hashed pw: ", hpw)
	valid := sm.ValidatePassword(pw, hpw)
	if !suc || hpw == "" || !valid {
		t.Fail()
	}
}
