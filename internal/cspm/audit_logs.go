package cspm

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"net/url"
)

const (
	auditLogEndpoint = "/audit/redlock"
)

// ListAuditLogs Returns audit logs for events that took place on the Prisma Cloud platform
//
// https://prisma.pan.dev/api/cloud/cspm/audit-logs#operation/rl-audit-logs
func (c *CspmClient) ListAuditLogs(q *AuditLogQuery) ([]AuditLog, error) {
	var encoder = schema.NewEncoder()
	params := url.Values{}
	err := encoder.Encode(q, params)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to decode Audit Logs query provided: %v", err)}
	}
	var auditLogs []AuditLog
	err = c.getWithResponseInterface(auditLogEndpoint, params, &auditLogs)
	if err != nil {
		return nil, err
	}
	return auditLogs, nil
}

type AuditLogQuery struct {
	TimeType   string `schema:"timeType,omitempty"`
	TimeAmount string `schema:"timeAmount,omitempty"`
	TimeUnit   string `schema:"timeUnit,omitempty"`
}

type AuditLog struct {
	Action       string `json:"action,omitempty"`
	ActionType   string `json:"actionType,omitempty"`
	IpAddress    string `json:"ipAddress,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Result       string `json:"result,omitempty"`
	Timestamp    int    `json:"timestamp,omitempty"`
	User         string `json:"user,omitempty"`
}
