package modc

import "runtime/debug"

// CallMeMaybe It's hard to look right at you.
func CallMeMaybe() string {
	v, _ := debug.ReadBuildInfo()
	return v.Main.Version
}
