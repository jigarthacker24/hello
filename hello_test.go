package hello

import "testing"

func TestSay(t *testing.T) {
	got := Say("Jigar")
	want := "Hello, Jigar!"

	if got != want {
		t.Fatal("got:", got, " want:", want)
	}
}

func TestSayWithPrefix(t *testing.T) {
	got := SayWithPrefix("Hey", "Jigar")
	want := "Hey, Jigar!"

	if got != want {
		t.Fatal("got: ", got, "want: ", want)
	}

	got = SayWithPrefix("", "")
	want = "Hello, world!"

	if got != want {
		t.Fatal("got: ", got, "want: ", want)
	}
}
