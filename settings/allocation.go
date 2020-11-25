package settings

type Allocation struct {
    totalShardsPerNode *uint
    enable             *AllocationVal
    include            AllocationAttributes
    require            AllocationAttributes
    exclude            AllocationAttributes
}

type AllocationAttributes map[AllocationAttribute][]interface{}

func (attributes AllocationAttributes) Add(name string, values ...interface{}) AllocationAttributes {
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
