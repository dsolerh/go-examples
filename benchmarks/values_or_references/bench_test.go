package valuesorreferences_test

import "testing"

var strResult string

type strStruct struct {
	phrase1 string
	phrase2 string
	phrase3 string
	phrase4 string
	phrase5 string
}

func (s strStruct) byValue() string {
	return s.phrase1 +
		s.phrase2 +
		s.phrase3 +
		s.phrase4 +
		s.phrase5
}

func (s strStruct) byReference() string {
	return s.phrase1 +
		s.phrase2 +
		s.phrase3 +
		s.phrase4 +
		s.phrase5
}

func Benchmark_strStruct(b *testing.B) {
	b.Run("passing the struct by value", func(b *testing.B) {
		var str string
		var phrases = strStruct{
			"superrrrrrr!",
			"kono marimo",
			"ore wa Monkey D. Luffy",
			"jiojojojojo",
			"kaisoku o ni ore wa naru",
		}
		for i := 0; i < b.N; i++ {
			str = phrases.byValue()
		}
		strResult = str
	})
	b.Run("passing the struct by reference", func(b *testing.B) {
		var str string
		var phrases = strStruct{
			"superrrrrrr!",
			"kono marimo",
			"ore wa Monkey D. Luffy",
			"jiojojojojo",
			"kaisoku o ni ore wa naru",
		}
		for i := 0; i < b.N; i++ {
			str = phrases.byReference()
		}
		strResult = str
	})

}
