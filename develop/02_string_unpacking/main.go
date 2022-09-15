package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const maxDigit = 9

type Unpacker struct {
	sb strings.Builder
}

func NewUnpacker() *Unpacker {
	return &Unpacker{}
}

func (u *Unpacker) multiplySymbol(sym rune, count int) {
	for i := 0; i < count; i++ {
		u.sb.WriteRune(sym)
	}
}

func (u *Unpacker) Parse(src string) error {
	u.sb.Reset()

	var lastSym rune
	var count strings.Builder

	if len(src) == 0 {
		return errors.New("cannot parse empty string")
	}

	for i, sym := range src {
		if unicode.IsDigit(sym) && i == 0 {
			return errors.New("incorrect string")
		}

		if !unicode.IsDigit(sym) {
			val, err := strconv.Atoi(count.String())
			if err != nil && lastSym != 0 {
				u.sb.WriteRune(lastSym)
				lastSym = sym
			} else {
				u.multiplySymbol(lastSym, val)
				lastSym = sym
			}
			count.Reset()
		} else {
			count.WriteRune(sym)
		}
	}

	u.sb.WriteRune(lastSym)

	return nil
}

func (u *Unpacker) String() string {
	return u.sb.String()
}

func main() {
	unpacker := NewUnpacker()

	unpacker.Parse("a4bc2d5e")
	fmt.Println(unpacker)

	unpacker.Parse("45")
	fmt.Println(unpacker)
}
