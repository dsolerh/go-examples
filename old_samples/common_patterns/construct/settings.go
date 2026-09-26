package construct

import "cmp"

type UserToCounter = map[string]int

type Foo struct {
	Info string
	Code int
}

type Bar struct {
	Name      string
	Dificulty int
}

type Callback func() error

type Settings interface {
	GetStr() string
	GetInt() int
	GetBool() bool
	GetMap() UserToCounter
	GetStruct() Foo
	GetSlice() []*Bar
	GetCallbackFn() Callback
}

func NewSettings(options SettingsOptions) Settings {
	settings := &settings{
		_str:      cmp.Or(options.Str, "green"),
		_int:      cmp.Or(options.Int, 12),
		_bool:     cmp.Or(options.Bool, true),
		_map:      UserToCounter{},
		_slice:    []*Bar{},
		_callback: func() error { return nil },
	}

	if len(options.Map) != 0 {
		settings._map = options.Map
	}
	if len(options.Slice) != 0 {
		settings._slice = options.Slice
	}
	if options.CallbackFn != nil {
		settings._callback = options.CallbackFn
	}

	return settings
}

type SettingsOptions struct {
	Str        string
	Int        int
	Bool       bool
	Map        UserToCounter
	Struct     Foo
	Slice      []*Bar
	CallbackFn Callback
}

var _ Settings = (*settings)(nil)

type settings struct {
	_str      string
	_int      int
	_bool     bool
	_map      UserToCounter
	_struct   Foo
	_slice    []*Bar
	_callback Callback
}

// GetCallbackFn implements Settings.
func (s *settings) GetCallbackFn() Callback {
	return s._callback
}

// GetSlice implements Settings.
func (s *settings) GetSlice() []*Bar {
	return s._slice
}

// GetStruct implements Settings.
func (s *settings) GetStruct() Foo {
	return s._struct
}

// GetMap implements Settings.
func (s *settings) GetMap() map[string]int {
	return s._map
}

// GetStr implements Settings.
func (s *settings) GetStr() string {
	return s._str
}

// GetInt implements Settings.
func (s *settings) GetInt() int {
	return s._int
}

// GetBool implements Settings.
func (s *settings) GetBool() bool {
	return s._bool
}
