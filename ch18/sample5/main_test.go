package main

import "testing"

func TestLangTimeAdd1(t *testing.T) {
	if LangTimeAdd1(1, 2) != 3 {
		t.Error("result of 1+2 is not 3")
	}
}

func TestLangTimeAdd2(t *testing.T) {
	if LangTimeAdd2(1, 2) != 3 {
		t.Error("result of 1+2 is not 3")
	}
}

func TestLangTimeAdd3(t *testing.T) {
	if LangTimeAdd3(1, 2) != 3 {
		t.Error("result of 1+2 is not 3")
	}
}
