package settings

type Write struct {
	waitForActiveShards uint32
}

func NewWrite(waitForActiveShards uint32) *Write {
	return &Write{waitForActiveShards: waitForActiveShards}
}

func (write *Write) Source() (interface{}, error) {
	return map[string]interface{}{
		"wait_for_active_shards": write.waitForActiveShards,
	}, nil
}
