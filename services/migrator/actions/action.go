package actions

import "github.com/olivere/elastic/v7"

const (
	Add     ActionType = "add"
	Delete  ActionType = "delete"
	Update  ActionType = "update"
	Reindex ActionType = "reindex"

	Fields   Object = "fields"
	Index    Object = "index"
	Template Object = "template"
	Alias    Object = "alias"
)

type ActionType string
type Object string

type Action interface {
	PerformRequestOptions() elastic.PerformRequestOptions
}
