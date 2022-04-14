package utils

import (
	"testing"
)

func TestCleanFileName_Empty(t *testing.T) {
	_, err := CleanFileName("")
	if err == nil {
		t.Fatal()
	}
}

func TestCleanFileName_Normal(t *testing.T) {
	f, err := CleanFileName("test")
	if err != nil {
		t.Fatal()
	}
	if f != "test.desktop" {
		t.Fatal()
	}
}

func TestCleanFileName_Specials(t *testing.T) {
	f, err := CleanFileName("test-!Name@ehehe 33 ]]¼` ` ` ¼€")
	if err != nil {
		t.Fatal()
	}
	if f != "test-name-ehehe-33.desktop" {
		t.Fatal()
	}
}
