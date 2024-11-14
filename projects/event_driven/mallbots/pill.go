package main

//go:generate go run golang.org/x/tools/cmd/stringer -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Ibuprofen
)
