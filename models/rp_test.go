package models

import (
	"testing"
)

func TestGetDefaultRelyingParty(t *testing.T) {
	rp, err := GetDefaultRelyingParty()
	if err != nil {
		t.Fatal(err)
	}

	if rp.Id != "" && rp.Name != "" {
		t.Fatal("this object is not empty!")
	}
}
