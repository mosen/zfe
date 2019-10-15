package hosts

import (
	"database/sql"
	zpg "github.com/mosen/zfe/pkg/postgresql"
)

// See: zabbix/frontends/php/include/defines.inc.php
const (
	ZbxFlagDiscoveryNormal    = 0
	ZbxFlagDiscoveryRule      = 1
	ZbxFlagDiscoveryPrototype = 2
	ZbxFlagDiscoveryCreated   = 4
)

// See: zabbix/frontends/php/include/defines.inc.php:310
const (
	HostStatusMonitored    = 0
	HostStatusNotMonitored = 1
	HostStatusTemplate     = 3
	HostStatusProxyActive  = 5
	HostStatusProxyPassive = 6
)

type Host struct {
	HostId       int                 `json:"hostid"`
	Host         string              `json:"host"`
	Name         string              `json:"name"`
	TemplateId   sql.NullInt64       `json:"-"`
	Description  string              `json:"description"`
	ProxyHostId  sql.NullInt64       `json:"proxy_hostid" db:"proxy_hostid"`
	ProxyAddress string              `json:"proxy_address" db:"proxy_address"`
	Status       int                 `json:"status"`
	DisableUntil zpg.ZabbixTimestamp `json:"disable_until" db:"disable_until"`
	Error        string              `json:"error"`
	Available    int                 `json:"available"`
	ErrorsFrom   zpg.ZabbixTimestamp `json:"errors_from" db:"errors_from"`
	LastAccess   zpg.ZabbixTimestamp `json:"lastaccess" db:"lastaccess"`
	Flags        int                 `json:"flags"`
	AutoCompress int                 `json:"auto_compress" db:"auto_compress"`

	// Maintenance
	MaintenanceId     sql.NullInt64       `json:"maintenanceid" db:"maintenanceid"`
	MaintenanceStatus int                 `json:"maintenance_status" db:"maintenance_status"`
	MaintenanceType   int                 `json:"maintenance_type" db:"maintenance_type"`
	MaintenanceFrom   zpg.ZabbixTimestamp `json:"maintenance_from" db:"maintenance_from"`

	// SNMP
	SnmpErrorsFrom   zpg.ZabbixTimestamp `json:"snmp_errors_from" db:"snmp_errors_from"`
	SnmpError        string              `json:"snmp_error" db:"snmp_error"`
	SnmpDisableUntil zpg.ZabbixTimestamp `json:"snmp_disable_until" db:"snmp_disable_until"`
	SnmpAvailable    int                 `json:"snmp_available" db:"snmp_available"`

	// IPMI Connection Parameters / Status
	IpmiAuthtype     int                 `json:"ipmi_authtype" db:"ipmi_authtype"`
	IpmiPrivilege    int                 `json:"ipmi_privilege" db:"ipmi_privilege"`
	IpmiUsername     string              `json:"ipmi_username" db:"ipmi_username"`
	IpmiPassword     string              `json:"ipmi_password" db:"ipmi_password"`
	IpmiDisableUntil zpg.ZabbixTimestamp `json:"ipmi_disable_until" db:"ipmi_disable_until"`
	IpmiAvailable    int                 `json:"ipmi_available" db:"ipmi_available"`
	IpmiError        string              `json:"ipmi_error" db:"ipmi_error"`
	IpmiErrorsFrom   zpg.ZabbixTimestamp `json:"ipmi_errors_from" db:"ipmi_errors_from"`

	// JMX Connection Parameters / Status
	JMXDisableUntil zpg.ZabbixTimestamp `json:"jmx_disable_until" db:"jmx_disable_until"`
	JMXAvailable    int                 `json:"jmx_available" db:"jmx_available"`
	JMXErrorsFrom   zpg.ZabbixTimestamp `json:"jmx_errors_from" db:"jmx_errors_from"`
	JMXError        string              `json:"jmx_error" db:"jmx_error"`

	// TLS Connection Parameters / Status
	TLSConnect     int    `json:"tls_connect" db:"tls_connect"`
	TLSAccept      int    `json:"tls_accept" db:"tls_accept"`
	TLSIssuer      string `json:"tls_issuer" db:"tls_issuer"`
	TLSSubject     string `json:"tls_subject" db:"tls_subject"`
	TLSPSKIdentity string `json:"tls_psk_identity" db:"tls_psk_identity"`
	TLSPSK         string `json:"tls_psk" db:"tls_psk"`
}

type Template Host

type HostInterface struct {
	InterfaceId int               `json:"interfaceid"`
	DNS         string            `json:"dns"`
	HostId      int               `json:"hostid"`
	IP          string            `json:"ip"`
	Main        zpg.ZabbixBoolean `json:"main"`
	Port        string            `json:"port"`
	Type        string            `json:"type"`
	UseIP       zpg.ZabbixBoolean `json:"useip"`
	Bulk        zpg.ZabbixBoolean `json:"bulk"`
}
