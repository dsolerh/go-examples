package main

import "common_patterns/construct"

func main() {
	construct.NewSettings(construct.SettingsOptions{
		Str:    "",
		Int:    0,
		Bool:   false,
		Map:    map[string]int{},
		Struct: construct.Foo{},
		Slice:  []*construct.Bar{},
		CallbackFn: func() error {
			return nil
		},
	})
}
