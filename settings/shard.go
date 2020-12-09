package settings

const (
	CheckOnStartupFalse    CheckOnStartup = "false"
	CheckOnStartupTrue     CheckOnStartup = "true"
	CheckOnStartupChecksum CheckOnStartup = "checksum"
)

type CheckOnStartup string

type Shard struct {
	checkOnStartup CheckOnStartup
}

func NewShard(checkOnStartup CheckOnStartup) *Shard {
	return &Shard{checkOnStartup: checkOnStartup}
}

func (shard *Shard) Source() (interface{}, error) {
	return map[string]interface{}{
		"check_on_startup": shard.checkOnStartup,
	}, nil
}
