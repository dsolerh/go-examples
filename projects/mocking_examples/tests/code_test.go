package pkg_test

import (
	"mocking_example/mocks"
	"mocking_example/pkg"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUseInterface(t *testing.T) {
	i := mocks.NewInterface(t)
	expectedStr := "good"
	i.EXPECT().GetValue(1523, "string", time.Second).Return(expectedStr, nil)

	str, err := pkg.UseInterface(i)

	is := assert.New(t)
	is.Equal(str, "good", "should return %s but got %s", expectedStr, str)
	is.Nil(err, "should return no error")
}

func TestNotUseInterface(t *testing.T) {
	i := mocks.NewInterface(t)
	i.EXPECT().GetValue(0, "", time.Duration(0)).Return("", nil)
	fn := func(i pkg.Interface) int {
		i.GetValue(0, "", time.Duration(0))
		return 111
	}
	val := fn(i)

	is := assert.New(t)
	is.Equal(val, 111, "should return %d but got %d", val, 111)
}
