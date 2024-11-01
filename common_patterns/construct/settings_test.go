package construct

import (
	"reflect"
	"testing"
)

func TestNewSettings(t *testing.T) {
	type args struct {
		options SettingsOptions
	}
	tests := []struct {
		name string
		args args
		want Settings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSettings(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}
