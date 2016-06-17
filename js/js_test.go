package js

import "testing"

func TestWeaponizeScript(t *testing.T) {
	alert := "alert(1);"
	alertWeaponizedExpected := "eval(atob('YWxlcnQoMSk7'))"

	alertWeaponized, err := WeaponizeScript(alert)
	if err != nil {
		t.Errorf("recieved an error from the minifier: %v", err)
	}

	if alertWeaponized != alertWeaponizedExpected {
		t.Errorf("expected %s, got %s", alertWeaponizedExpected, alertWeaponized)
	}
}
