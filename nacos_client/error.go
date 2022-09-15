package nacos_client

import "errors"

var (
	NoSelectorError      = errors.New("no such selector")
	NoAvailableNodeError = errors.New("no available node")
)
