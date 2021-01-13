package settings

import (
	"reflect"
	"testing"
)

func TestAllocationAttributes_HostIps(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name       string
		attributes AllocationAttributes
		args       args
		want       AllocationAttributes
	}{
		{
			name:       "set one",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"127.0.0.1"}},
			want:       AllocationAttributes{AllocationAttributeHostIp: []interface{}{"127.0.0.1"}},
		},
		{
			name:       "set two",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"127.0.0.1", "127.0.0.2"}},
			want:       AllocationAttributes{AllocationAttributeHostIp: []interface{}{"127.0.0.1", "127.0.0.2"}},
		},
		{
			name:       "add one",
			attributes: AllocationAttributes{AllocationAttributeHostIp: []interface{}{"127.0.0.1"}},
			args:       args{values: []interface{}{"127.0.0.2"}},
			want:       AllocationAttributes{AllocationAttributeHostIp: []interface{}{"127.0.0.1", "127.0.0.2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.attributes.HostIps(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HostIps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocationAttributes_Hosts(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name       string
		attributes AllocationAttributes
		args       args
		want       AllocationAttributes
	}{
		{
			name:       "set one",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"test.com"}},
			want:       AllocationAttributes{AllocationAttributeHost: []interface{}{"test.com"}},
		},
		{
			name:       "set two",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"test.com", "test2.com"}},
			want:       AllocationAttributes{AllocationAttributeHost: []interface{}{"test.com", "test2.com"}},
		},
		{
			name:       "add two",
			attributes: AllocationAttributes{AllocationAttributeHost: []interface{}{"test.com"}},
			args:       args{values: []interface{}{"test2.com"}},
			want:       AllocationAttributes{AllocationAttributeHost: []interface{}{"test.com", "test2.com"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.attributes.Hosts(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocationAttributes_Ids(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name       string
		attributes AllocationAttributes
		args       args
		want       AllocationAttributes
	}{
		{
			name:       "set one",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"id1"}},
			want:       AllocationAttributes{AllocationAttributeId: []interface{}{"id1"}},
		},
		{
			name:       "set two",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"id1", "id2"}},
			want:       AllocationAttributes{AllocationAttributeId: []interface{}{"id1", "id2"}},
		},
		{
			name:       "add one",
			attributes: AllocationAttributes{AllocationAttributeId: []interface{}{"id1"}},
			args:       args{values: []interface{}{"id2"}},
			want:       AllocationAttributes{AllocationAttributeId: []interface{}{"id1", "id2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.attributes.Ids(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ids() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocationAttributes_Ips(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name       string
		attributes AllocationAttributes
		args       args
		want       AllocationAttributes
	}{
		{
			name:       "set one",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"127.0.0.1"}},
			want:       AllocationAttributes{AllocationAttributeIp: []interface{}{"127.0.0.1"}},
		},
		{
			name:       "set two",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"127.0.0.1", "127.0.0.2"}},
			want:       AllocationAttributes{AllocationAttributeIp: []interface{}{"127.0.0.1", "127.0.0.2"}},
		},
		{
			name:       "add one",
			attributes: AllocationAttributes{AllocationAttributeIp: []interface{}{"127.0.0.1"}},
			args:       args{values: []interface{}{"127.0.0.2"}},
			want:       AllocationAttributes{AllocationAttributeIp: []interface{}{"127.0.0.1", "127.0.0.2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.attributes.Ips(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ips() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocationAttributes_Names(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name       string
		attributes AllocationAttributes
		args       args
		want       AllocationAttributes
	}{
		{
			name:       "set one",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"test"}},
			want:       AllocationAttributes{AllocationAttributeName: []interface{}{"test"}},
		},
		{
			name:       "set two",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"test", "test2"}},
			want:       AllocationAttributes{AllocationAttributeName: []interface{}{"test", "test2"}},
		},
		{
			name:       "add one",
			attributes: AllocationAttributes{AllocationAttributeName: []interface{}{"test"}},
			args:       args{values: []interface{}{"test2"}},
			want:       AllocationAttributes{AllocationAttributeName: []interface{}{"test", "test2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.attributes.Names(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocationAttributes_PublishIps(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name       string
		attributes AllocationAttributes
		args       args
		want       AllocationAttributes
	}{
		{
			name:       "set one",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"127.0.0.1"}},
			want:       AllocationAttributes{AllocationAttributePublishIp: []interface{}{"127.0.0.1"}},
		},
		{
			name:       "set two",
			attributes: AllocationAttributes{},
			args:       args{values: []interface{}{"127.0.0.1", "127.0.0.2"}},
			want:       AllocationAttributes{AllocationAttributePublishIp: []interface{}{"127.0.0.1", "127.0.0.2"}},
		},
		{
			name:       "add one",
			attributes: AllocationAttributes{AllocationAttributePublishIp: []interface{}{"127.0.0.1"}},
			args:       args{values: []interface{}{"127.0.0.2"}},
			want:       AllocationAttributes{AllocationAttributePublishIp: []interface{}{"127.0.0.1", "127.0.0.2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.attributes.PublishIps(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishIps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocationAttributes_Set(t *testing.T) {
	type args struct {
		name   string
		values []interface{}
	}
	tests := []struct {
		name       string
		attributes AllocationAttributes
		args       args
		want       AllocationAttributes
	}{
		{
			name:       "set one",
			attributes: AllocationAttributes{},
			args:       args{name: "test", values: []interface{}{"127.0.0.1"}},
			want:       AllocationAttributes{"test": []interface{}{"127.0.0.1"}},
		},
		{
			name:       "set two",
			attributes: AllocationAttributes{},
			args:       args{name: "test", values: []interface{}{"127.0.0.1", "127.0.0.2"}},
			want:       AllocationAttributes{"test": []interface{}{"127.0.0.1", "127.0.0.2"}},
		},
		{
			name:       "add one",
			attributes: AllocationAttributes{"test": []interface{}{"127.0.0.1"}},
			args:       args{name: "test", values: []interface{}{"127.0.0.2"}},
			want:       AllocationAttributes{"test": []interface{}{"127.0.0.1", "127.0.0.2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.attributes.Set(tt.args.name, tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocation_SetEnable(t *testing.T) {
	allocationAll := AllocationAll
	allocationPrimaries := AllocationPrimaries
	allocationNewPrimaries := AllocationNewPrimaries
	allocationNone := AllocationNone

	type fields struct {
		totalShardsPerNode *uint
		enable             *AllocationVal
		include            AllocationAttributes
		require            AllocationAttributes
		exclude            AllocationAttributes
	}
	type args struct {
		enable AllocationVal
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Allocation
	}{
		{
			name:   "allocation_all",
			fields: fields{},
			args:   args{enable: AllocationAll},
			want:   &Allocation{enable: &allocationAll},
		},
		{
			name:   "allocation_new_primaries",
			fields: fields{},
			args:   args{enable: AllocationNewPrimaries},
			want:   &Allocation{enable: &allocationNewPrimaries},
		},
		{
			name:   "allocation_primaries",
			fields: fields{},
			args:   args{enable: AllocationPrimaries},
			want:   &Allocation{enable: &allocationPrimaries},
		},
		{
			name:   "allocation_none",
			fields: fields{},
			args:   args{enable: AllocationNone},
			want:   &Allocation{enable: &allocationNone},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allocation := &Allocation{
				totalShardsPerNode: tt.fields.totalShardsPerNode,
				enable:             tt.fields.enable,
				include:            tt.fields.include,
				require:            tt.fields.require,
				exclude:            tt.fields.exclude,
			}
			if got := allocation.SetEnable(tt.args.enable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetEnable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocation_SetExclude(t *testing.T) {
	type fields struct {
		totalShardsPerNode *uint
		enable             *AllocationVal
		include            AllocationAttributes
		require            AllocationAttributes
		exclude            AllocationAttributes
	}
	type args struct {
		exclude AllocationAttributes
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Allocation
	}{
		{
			args: args{exclude: AllocationAttributes{}},
			want: &Allocation{exclude: AllocationAttributes{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allocation := &Allocation{
				totalShardsPerNode: tt.fields.totalShardsPerNode,
				enable:             tt.fields.enable,
				include:            tt.fields.include,
				require:            tt.fields.require,
				exclude:            tt.fields.exclude,
			}
			if got := allocation.SetExclude(tt.args.exclude); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExclude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocation_SetInclude(t *testing.T) {
	type fields struct {
		totalShardsPerNode *uint
		enable             *AllocationVal
		include            AllocationAttributes
		require            AllocationAttributes
		exclude            AllocationAttributes
	}
	type args struct {
		include AllocationAttributes
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Allocation
	}{
		{
			args: args{include: AllocationAttributes{}},
			want: &Allocation{include: AllocationAttributes{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allocation := &Allocation{
				totalShardsPerNode: tt.fields.totalShardsPerNode,
				enable:             tt.fields.enable,
				include:            tt.fields.include,
				require:            tt.fields.require,
				exclude:            tt.fields.exclude,
			}
			if got := allocation.SetInclude(tt.args.include); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetInclude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocation_SetRequire(t *testing.T) {
	type fields struct {
		totalShardsPerNode *uint
		enable             *AllocationVal
		include            AllocationAttributes
		require            AllocationAttributes
		exclude            AllocationAttributes
	}
	type args struct {
		require AllocationAttributes
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Allocation
	}{
		{
			args: args{require: AllocationAttributes{}},
			want: &Allocation{require: AllocationAttributes{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allocation := &Allocation{
				totalShardsPerNode: tt.fields.totalShardsPerNode,
				enable:             tt.fields.enable,
				include:            tt.fields.include,
				require:            tt.fields.require,
				exclude:            tt.fields.exclude,
			}
			if got := allocation.SetRequire(tt.args.require); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRequire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocation_SetTotalShardsPerNode(t *testing.T) {
	totalShardsPerNodeSet := uint(1)
	totalShardsPerNodeChange := uint(2)

	type fields struct {
		totalShardsPerNode *uint
		enable             *AllocationVal
		include            AllocationAttributes
		require            AllocationAttributes
		exclude            AllocationAttributes
	}
	type args struct {
		totalShardsPerNode uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Allocation
	}{
		{
			name: "set",
			args: args{totalShardsPerNode: totalShardsPerNodeSet},
			want: &Allocation{totalShardsPerNode: &totalShardsPerNodeSet},
		},
		{
			name:   "change",
			fields: fields{totalShardsPerNode: &totalShardsPerNodeSet},
			args:   args{totalShardsPerNode: totalShardsPerNodeChange},
			want:   &Allocation{totalShardsPerNode: &totalShardsPerNodeChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allocation := &Allocation{
				totalShardsPerNode: tt.fields.totalShardsPerNode,
				enable:             tt.fields.enable,
				include:            tt.fields.include,
				require:            tt.fields.require,
				exclude:            tt.fields.exclude,
			}
			if got := allocation.SetTotalShardsPerNode(tt.args.totalShardsPerNode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTotalShardsPerNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocation_Source(t *testing.T) {
	totalShardsPerNode := uint(1)
	allocationAll := AllocationAll
	allocationPrimaries := AllocationPrimaries
	allocationNewPrimaries := AllocationNewPrimaries
	allocationNone := AllocationNone
	allocationAttributes := AllocationAttributes{"test": []interface{}{"test"}}

	type fields struct {
		totalShardsPerNode *uint
		enable             *AllocationVal
		include            AllocationAttributes
		require            AllocationAttributes
		exclude            AllocationAttributes
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name:   "total_shards_per_node",
			fields: fields{totalShardsPerNode: &totalShardsPerNode},
			want: map[string]interface{}{
				"total_shards_per_node": totalShardsPerNode,
			},
			wantErr: false,
		},
		{
			name:   "enable_allocationAll",
			fields: fields{enable: &allocationAll},
			want: map[string]interface{}{
				"enable": allocationAll,
			},
			wantErr: false,
		},
		{
			name:   "enable_allocationPrimaries",
			fields: fields{enable: &allocationPrimaries},
			want: map[string]interface{}{
				"enable": allocationPrimaries,
			},
			wantErr: false,
		},
		{
			name:   "enable_allocationNewPrimaries",
			fields: fields{enable: &allocationNewPrimaries},
			want: map[string]interface{}{
				"enable": allocationNewPrimaries,
			},
			wantErr: false,
		},
		{
			name:   "enable_allocationNone",
			fields: fields{enable: &allocationNone},
			want: map[string]interface{}{
				"enable": allocationNone,
			},
			wantErr: false,
		},
		{
			name:   "include",
			fields: fields{include: allocationAttributes},
			want: map[string]interface{}{
				"include": allocationAttributes,
			},
			wantErr: false,
		},
		{
			name:   "require",
			fields: fields{require: allocationAttributes},
			want: map[string]interface{}{
				"require": allocationAttributes,
			},
			wantErr: false,
		},
		{
			name:   "exclude",
			fields: fields{exclude: allocationAttributes},
			want: map[string]interface{}{
				"exclude": allocationAttributes,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allocation := &Allocation{
				totalShardsPerNode: tt.fields.totalShardsPerNode,
				enable:             tt.fields.enable,
				include:            tt.fields.include,
				require:            tt.fields.require,
				exclude:            tt.fields.exclude,
			}
			got, err := allocation.Source()
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

func TestNewAllocation(t *testing.T) {
	tests := []struct {
		name string
		want *Allocation
	}{
		{
			want: &Allocation{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAllocation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAllocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAllocationAttributes(t *testing.T) {
	tests := []struct {
		name string
		want AllocationAttributes
	}{
		{
			want: AllocationAttributes{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAllocationAttributes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAllocationAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}
