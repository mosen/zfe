package comms

// This file tracks the upstream zabbix/include/zbxjson.h for constants and other definitions needed to implement
// the zabbix protocol(s)

const (
	ZbxProtoValueGetActiveChecks     = "active checks"
	ZbxProtoValueProxyConfig         = "proxy config"
	ZbxProtoValueProxyHeartbeat      = "proxy heartbeat"
	ZbxProtoValueSenderData          = "sender data"
	ZbxProtoValueAgentData           = "agent data"
	ZbxProtoValueCommand             = "command"
	ZbxProtoValueJavaGatewayInternal = "java gateway internal"
	ZbxProtoValueJavaGatewayJmx      = "java gateway jmx"
	ZbxProtoValueGetQueue            = "queue.get"
	ZbxProtoValueGetStatus           = "status.get"
	ZbxProtoValueProxyData           = "proxy data"
	ZbxProtoValueProxyTasks          = "proxy tasks"
)
