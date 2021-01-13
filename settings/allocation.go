package settings

const (
    AllocationAttributeName      AllocationAttribute = "_name"
    AllocationAttributeHostIp    AllocationAttribute = "_host_ip"
    AllocationAttributePublishIp AllocationAttribute = "_publish_ip"
    AllocationAttributeIp        AllocationAttribute = "_ip"
    AllocationAttributeHost      AllocationAttribute = "_host"
    AllocationAttributeId        AllocationAttribute = "_id"

    AllocationAll          AllocationVal = "all"
    AllocationPrimaries    AllocationVal = "primaries"
    AllocationNewPrimaries AllocationVal = "new_primaries"
    AllocationNone         AllocationVal = "none"
)

type AllocationAttribute string
type AllocationVal string

type AllocationAttributes map[AllocationAttribute][]interface{}

func NewAllocationAttributes() AllocationAttributes {
    return AllocationAttributes{}
}

func (attributes AllocationAttributes) Set(name string, values ...interface{}) AllocationAttributes {
    attributeName := AllocationAttribute(name)

    attributes[attributeName] = append(attributes[attributeName], values...)

    return attributes
}

func (attributes AllocationAttributes) Names(values ...interface{}) AllocationAttributes {
    attributes[AllocationAttributeName] = append(attributes[AllocationAttributeName], values...)

    return attributes
}

func (attributes AllocationAttributes) HostIps(values ...interface{}) AllocationAttributes {
    attributes[AllocationAttributeHostIp] = append(attributes[AllocationAttributeHostIp], values...)

    return attributes
}

func (attributes AllocationAttributes) PublishIps(values ...interface{}) AllocationAttributes {
    attributes[AllocationAttributePublishIp] = append(attributes[AllocationAttributePublishIp], values...)

    return attributes
}

func (attributes AllocationAttributes) Ips(values ...interface{}) AllocationAttributes {
    attributes[AllocationAttributeIp] = append(attributes[AllocationAttributeIp], values...)

    return attributes
}

func (attributes AllocationAttributes) Hosts(values ...interface{}) AllocationAttributes {
    attributes[AllocationAttributeHost] = append(attributes[AllocationAttributeHost], values...)

    return attributes
}

func (attributes AllocationAttributes) Ids(values ...interface{}) AllocationAttributes {
    attributes[AllocationAttributeId] = append(attributes[AllocationAttributeId], values...)

    return attributes
}

type Allocation struct {
    totalShardsPerNode *uint
    enable             *AllocationVal
    include            AllocationAttributes
    require            AllocationAttributes
    exclude            AllocationAttributes
}

func NewAllocation() *Allocation {
    return &Allocation{}
}

func (allocation *Allocation) SetTotalShardsPerNode(totalShardsPerNode uint) *Allocation {
    allocation.totalShardsPerNode = &totalShardsPerNode

    return allocation
}

func (allocation *Allocation) SetEnable(enable AllocationVal) *Allocation {
    allocation.enable = &enable

    return allocation
}

func (allocation *Allocation) SetInclude(include AllocationAttributes) *Allocation {
    allocation.include = include

    return allocation
}

func (allocation *Allocation) SetRequire(require AllocationAttributes) *Allocation {
    allocation.require = require

    return allocation
}

func (allocation *Allocation) SetExclude(exclude AllocationAttributes) *Allocation {
    allocation.exclude = exclude

    return allocation
}

func (allocation *Allocation) Source() (interface{}, error) {
    source := map[string]interface{}{}

    if allocation.totalShardsPerNode != nil {
        source["total_shards_per_node"] = *allocation.totalShardsPerNode
    }

    if allocation.enable != nil {
        source["enable"] = *allocation.enable
    }

    if len(allocation.include) > 0 {
        source["include"] = allocation.include
    }

    if len(allocation.require) > 0 {
        source["require"] = allocation.require
    }

    if len(allocation.exclude) > 0 {
        source["exclude"] = allocation.exclude
    }

    return source, nil
}
