package digits

func IsDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func ToDigit(b byte) int64 {
	return int64(b - '0')
}

func IsSymbol(b byte) bool {
	return !IsDigit(b) && b != '.'
}
