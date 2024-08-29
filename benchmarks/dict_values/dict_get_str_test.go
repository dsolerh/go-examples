package dictvalues_test

import (
	"fmt"
	"testing"
)

type Dict = map[string]any

func Value[V any](d Dict, prop string) (V, error) {
	if str, ok := d[prop].(V); !ok {
		if _, exist := d[prop]; exist {
			return str, fmt.Errorf("prop %s is not present in the dict", prop)
		}
		return str, fmt.Errorf("prop %s is not of type %T", prop, str)
	} else {
		return str, nil
	}
}

var str string

func Benchmark_get_any_from_dict(b *testing.B) {
	var dict = Dict{
		"prop": "some very interesting string",
	}
	directCast := func(d Dict, prop string) (string, error) {
		str, ok := d[prop].(string)
		if ok {
			return str, nil
		}
		return "", fmt.Errorf("invalid string")
	}

	var _str string
	b.Run("using Value[string]", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_str, _ = Value[string](dict, "prop")
		}
		str = _str
	})
	b.Run("using direct cast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_str, _ = directCast(dict, "prop")
		}
		str = _str
	})
}

type IntLike interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func IntLikeValue[V IntLike](d Dict, prop string) (V, error) {
	val, exist := d[prop]
	if !exist {
		return V(0), fmt.Errorf("prop %s is not present in the dict", prop)
	}
	switch _val := val.(type) {
	case V:
		return _val, nil
	case int:
		v_val := V(_val)
		if int(v_val) != (_val) {
			return V(0), fmt.Errorf("prop %s cannot be safely converted from '%T' to '%T'", prop, _val, v_val)
		}
		return v_val, nil
	case int8:
		v_val := V(_val)
		if int8(v_val) != (_val) {
			return V(0), fmt.Errorf("prop %s cannot be safely converted from '%T' to '%T'", prop, _val, v_val)
		}
		return v_val, nil
	case int16:
		v_val := V(_val)
		if int16(v_val) != (_val) {
			return V(0), fmt.Errorf("prop %s cannot be safely converted from '%T' to '%T'", prop, _val, v_val)
		}
		return v_val, nil
	case int32:
		v_val := V(_val)
		if int32(v_val) != (_val) {
			return V(0), fmt.Errorf("prop %s cannot be safely converted from '%T' to '%T'", prop, _val, v_val)
		}
		return v_val, nil
	case int64:
		v_val := V(_val)
		if int64(v_val) != (_val) {
			return V(0), fmt.Errorf("prop %s cannot be safely converted from '%T' to '%T'", prop, _val, v_val)
		}
		return v_val, nil
	default:
		return V(0), fmt.Errorf("prop %s is not of type %T", prop, str)
	}
}

var _int int

func Benchmark_get_int_like_from_dict(b *testing.B) {
	var dict = Dict{
		"prop": int(23),
	}
	directCast := func(d Dict, prop string) (int, error) {
		num, ok := d[prop].(int)
		if ok {
			return num, nil
		}
		return 0, fmt.Errorf("invalid string")
	}

	var _num int
	var _num8 int8
	b.Run("using IntLikeValue[int]", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_num, _ = IntLikeValue[int](dict, "prop")
		}
		_int = _num
	})
	b.Run("using IntLikeValue[int8]", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_num8, _ = IntLikeValue[int8](dict, "prop")

		}
		_int = int(_num8)
	})
	b.Run("using direct cast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_num, _ = directCast(dict, "prop")
		}
		_int = _num
	})
}
