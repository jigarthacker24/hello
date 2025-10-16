package hello

import "testing"

func TestSay(t *testing.T) {
	got := Say("Jigar", true)
	want := "Hello, Jigar!!!"
	if got != want {
		t.Fatal("got:", got, " want:", want)
	}
	got = Say("", false)
	want = "Hello, world."
	if got != want {
		t.Fatal("got:", got, " want:", want)
	}
}

func TestSayWithPrefix(t *testing.T) {
	got := SayWithPrefix("Hey", "Jigar", true)
	want := "Hey, Jigar!!!"

	if got != want {
		t.Fatal("got: ", got, "want: ", want)
	}

	got = SayWithPrefix("", "", false)
	want = "Hello, world."

	if got != want {
		t.Fatal("got: ", got, "want: ", want)
	}
}
