package generate

import (
	"strconv"
)

var (
	seq int
)

func UniqueID() string {
	seq++
	return strconv.Itoa(seq)
}
