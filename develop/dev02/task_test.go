package main

import "testing"

func TestUnpacker(t *testing.T) {
	t.Run("a4bc2d5e", func(t *testing.T) {
		unpacker := NewUnpacker()

		err := unpacker.Parse("a4bc2d5e")
		if err != nil {
			t.Error(err)
		}

		got := unpacker.String()
		want := "aaaabccddddde"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("abcd", func(t *testing.T) {
		unpacker := NewUnpacker()

		err := unpacker.Parse("abcd")
		if err != nil {
			t.Error(err)
		}

		got := unpacker.String()
		want := "abcd"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("45", func(t *testing.T) {
		unpacker := NewUnpacker()

		err := unpacker.Parse("45")
		if err == nil {
			t.Error("should invoke error")
		}

		got := unpacker.String()
		want := ""

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("", func(t *testing.T) {
		unpacker := NewUnpacker()

		err := unpacker.Parse("")
		if err == nil {
			t.Error("should invoke error")
		}

		got := unpacker.String()
		want := ""

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
