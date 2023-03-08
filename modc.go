package modc

import (
	"runtime/debug"
)

const modName = "github.com/manelmontilla/modc"

// CallMeMaybe Hey, I just met you, and this is crazy.
func CallMeMaybe() string {
	v, _ := debug.ReadBuildInfo()
	version := "unknown"
	for _, m := range v.Deps {
		if m.Path == modName {
			version = m.Version
			break
		}
	}
	return version
}
