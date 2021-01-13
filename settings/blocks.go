package settings

type Blocks struct {
	metadata            *bool
	read                *bool
	readOnlyAllowDelete *bool
	readOnly            *bool
	write               *bool
}

func NewBlocks() *Blocks {
	return &Blocks{}
}

func (blocks *Blocks) SetMetadata(metadata bool) *Blocks {
	blocks.metadata = &metadata

	return blocks
}

func (blocks *Blocks) SetRead(read bool) *Blocks {
	blocks.read = &read

	return blocks
}

func (blocks *Blocks) SetReadOnlyAllowDelete(readOnlyAllowDelete bool) *Blocks {
	blocks.readOnlyAllowDelete = &readOnlyAllowDelete

	return blocks
}

func (blocks *Blocks) SetReadOnly(readOnly bool) *Blocks {
	blocks.readOnly = &readOnly

	return blocks
}

func (blocks *Blocks) SetWrite(write bool) *Blocks {
	blocks.write = &write

	return blocks
}

func (blocks *Blocks) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if blocks.metadata != nil {
		source["metadata"] = *blocks.metadata
	}

	if blocks.read != nil {
		source["read"] = *blocks.read
	}

	if blocks.readOnlyAllowDelete != nil {
		source["read_only_allow_delete"] = *blocks.readOnlyAllowDelete
	}

	if blocks.readOnly != nil {
		source["read_only"] = *blocks.readOnly
	}

	if blocks.write != nil {
		source["write"] = *blocks.write
	}

	return source, nil
}
