package hostgroups

import (
	zpg "github.com/mosen/zfe/pkg/postgresql"
)

type Hostgroup struct {
	GroupId  int               `json:"groupid"`
	Name     string            `json:"name"`
	Flags    int               `json:"flags"`
	Internal zpg.ZabbixBoolean `json:"internal"`
}
