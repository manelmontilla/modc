package modc

import "runtime/debug"

// CallMeMaybe Hey, I just met you, and this is crazy
func CallMeMaybe() string {
	v, _ := debug.ReadBuildInfo()
	return v.Main.Version
}
