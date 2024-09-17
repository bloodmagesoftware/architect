package version

import "fmt"

type Version struct {
	Major int `toml:"major"`
	Minor int `toml:"minor"`
}

var devel = Version{Major: -1, Minor: -1}

// Parse parses a Version struct from the string format v{major}.{minor}.{patch}
// Since the patch version is not used, it is not parsed.
func Parse(s string) (Version, error) {
	if s == "(devel)" {
		return devel, nil
	}
	var v Version
	_, err := fmt.Sscanf(s, "v%d.%d", &v.Major, &v.Minor)
	return v, err
}

// Compare returns 1 if v is newer than other, 0 if they are equal, and -1 if v is older than other.
func (v Version) Compare(other Version) int {
	if v.Major == -1 && v.Minor == -1 || other.Major == -1 && other.Minor == -1 {
		return 0
	}

	if v.Major > other.Major {
		return 1
	}
	if v.Major < other.Major {
		return -1
	}
	if v.Minor > other.Minor {
		return 1
	}
	if v.Minor < other.Minor {
		return -1
	}
	return 0
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d", v.Major, v.Minor)
}
