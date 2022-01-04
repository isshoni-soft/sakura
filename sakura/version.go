package sakura

import "strconv"

type Version struct {
	Major int
	Minor int
	Patch int
	Snapshot bool
}

func (v *Version) GetVersion() (result string) {
	result = strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch)

	if v.Snapshot {
		result = result + "-SNAPSHOT"
	}

	return
}
