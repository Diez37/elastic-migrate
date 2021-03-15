package version

type Version struct {
	number      int32
	description string
}

func NewVersion(number int32, description string) *Version {
	return &Version{number: number, description: description}
}

func (version *Version) Number() int32 {
	return version.number
}

func (version *Version) Description() string {
	return version.description
}
