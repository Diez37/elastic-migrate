package settings

import (
	"reflect"
	"testing"
)

func TestBlocks_SetMetadata(t *testing.T) {
	boolSet := true
	boolChange := false

	type fields struct {
		metadata            *bool
		read                *bool
		readOnlyAllowDelete *bool
		readOnly            *bool
		write               *bool
	}
	type args struct {
		metadata bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Blocks
	}{
		{
			name: "set",
			args: args{metadata: true},
			want: &Blocks{metadata: &boolSet},
		},
		{
			name: "change",
			fields: fields{metadata: &boolSet},
			args: args{metadata: false},
			want: &Blocks{metadata: &boolChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blocks := &Blocks{
				metadata:            tt.fields.metadata,
				read:                tt.fields.read,
				readOnlyAllowDelete: tt.fields.readOnlyAllowDelete,
				readOnly:            tt.fields.readOnly,
				write:               tt.fields.write,
			}
			if got := blocks.SetMetadata(tt.args.metadata); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlocks_SetRead(t *testing.T) {
	boolSet := true
	boolChange := false

	type fields struct {
		metadata            *bool
		read                *bool
		readOnlyAllowDelete *bool
		readOnly            *bool
		write               *bool
	}
	type args struct {
		read bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Blocks
	}{
		{
			name: "set",
			args: args{read: true},
			want: &Blocks{read: &boolSet},
		},
		{
			name: "change",
			fields: fields{read: &boolSet},
			args: args{read: false},
			want: &Blocks{read: &boolChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blocks := &Blocks{
				metadata:            tt.fields.metadata,
				read:                tt.fields.read,
				readOnlyAllowDelete: tt.fields.readOnlyAllowDelete,
				readOnly:            tt.fields.readOnly,
				write:               tt.fields.write,
			}
			if got := blocks.SetRead(tt.args.read); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlocks_SetReadOnly(t *testing.T) {
	boolSet := true
	boolChange := false

	type fields struct {
		metadata            *bool
		read                *bool
		readOnlyAllowDelete *bool
		readOnly            *bool
		write               *bool
	}
	type args struct {
		readOnly bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Blocks
	}{
		{
			name: "set",
			args: args{readOnly: true},
			want: &Blocks{readOnly: &boolSet},
		},
		{
			name: "change",
			fields: fields{readOnly: &boolSet},
			args: args{readOnly: false},
			want: &Blocks{readOnly: &boolChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blocks := &Blocks{
				metadata:            tt.fields.metadata,
				read:                tt.fields.read,
				readOnlyAllowDelete: tt.fields.readOnlyAllowDelete,
				readOnly:            tt.fields.readOnly,
				write:               tt.fields.write,
			}
			if got := blocks.SetReadOnly(tt.args.readOnly); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetReadOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlocks_SetReadOnlyAllowDelete(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		metadata            *bool
		read                *bool
		readOnlyAllowDelete *bool
		readOnly            *bool
		write               *bool
	}
	type args struct {
		readOnlyAllowDelete bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Blocks
	}{
		{
			name: "set",
			args: args{readOnlyAllowDelete: true},
			want: &Blocks{readOnlyAllowDelete: &boolTrue},
		},
		{
			name: "change",
			fields: fields{readOnlyAllowDelete: &boolTrue},
			args: args{readOnlyAllowDelete: false},
			want: &Blocks{readOnlyAllowDelete: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blocks := &Blocks{
				metadata:            tt.fields.metadata,
				read:                tt.fields.read,
				readOnlyAllowDelete: tt.fields.readOnlyAllowDelete,
				readOnly:            tt.fields.readOnly,
				write:               tt.fields.write,
			}
			if got := blocks.SetReadOnlyAllowDelete(tt.args.readOnlyAllowDelete); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetReadOnlyAllowDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlocks_SetWrite(t *testing.T) {
	boolSet := true
	boolChange := false

	type fields struct {
		metadata            *bool
		read                *bool
		readOnlyAllowDelete *bool
		readOnly            *bool
		write               *bool
	}
	type args struct {
		write bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Blocks
	}{
		{
			name: "set",
			args: args{write: true},
			want: &Blocks{write: &boolSet},
		},
		{
			name: "change",
			fields: fields{write: &boolSet},
			args: args{write: false},
			want: &Blocks{write: &boolChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blocks := &Blocks{
				metadata:            tt.fields.metadata,
				read:                tt.fields.read,
				readOnlyAllowDelete: tt.fields.readOnlyAllowDelete,
				readOnly:            tt.fields.readOnly,
				write:               tt.fields.write,
			}
			if got := blocks.SetWrite(tt.args.write); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWrite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlocks_Source(t *testing.T) {
	boolTrue := true

	type fields struct {
		metadata            *bool
		read                *bool
		readOnlyAllowDelete *bool
		readOnly            *bool
		write               *bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "metadata",
			fields: fields{metadata: &boolTrue},
			want: map[string]interface{}{
				"metadata": true,
			},
		},
		{
			name: "read",
			fields: fields{read: &boolTrue},
			want: map[string]interface{}{
				"read": true,
			},
		},
		{
			name: "read_only_allow_delete",
			fields: fields{readOnlyAllowDelete: &boolTrue},
			want: map[string]interface{}{
				"read_only_allow_delete": true,
			},
		},
		{
			name: "read_only",
			fields: fields{readOnly: &boolTrue},
			want: map[string]interface{}{
				"read_only": true,
			},
		},
		{
			name: "write",
			fields: fields{write: &boolTrue},
			want: map[string]interface{}{
				"write": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blocks := &Blocks{
				metadata:            tt.fields.metadata,
				read:                tt.fields.read,
				readOnlyAllowDelete: tt.fields.readOnlyAllowDelete,
				readOnly:            tt.fields.readOnly,
				write:               tt.fields.write,
			}
			got, err := blocks.Source()
			if (err != nil) != tt.wantErr {
				t.Errorf("Source() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Source() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBlocks(t *testing.T) {
	tests := []struct {
		name string
		want *Blocks
	}{
		{
			want: &Blocks{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlocks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlocks() = %v, want %v", got, tt.want)
			}
		})
	}
}