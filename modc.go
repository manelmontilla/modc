package modc

import "runtime/debug"

// CallMeMaybe It's hard to look right at you baby, but here's my number
func CallMeMaybe() string {
	v, _ := debug.ReadBuildInfo()
	return v.Main.Version
}
