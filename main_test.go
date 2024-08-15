package main

import "testing"

func Test_sayHello(t *testing.T) {
	name := "bob"
	want := "hello bob"
	if got := SayHello(name); got != want {
		t.Errorf("hello()= %q, want %q", got, want)
	}
}
