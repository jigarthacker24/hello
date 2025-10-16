package hello

import "testing"

func TestSay(t *testing.T) {
	got := Say("Jigar")
	want := "Hello, Jigar!"

	if got != want {
		t.Fatal("got:", got, " want:", want)
	}
}
