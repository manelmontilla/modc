package modc

import (
	"runtime/debug"
	"testing"
)

func TestCallMeMaybe(t *testing.T) {
	v, _ := debug.ReadBuildInfo()
	tests := []struct {
		name string
		want string
	}{
		{
			name: "golden path",
			want: v.Main.Version,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CallMeMaybe(); got != tt.want {
				t.Errorf("CallMeMaybe() = %v, want %v", got, tt.want)
			}
		})
	}
}
