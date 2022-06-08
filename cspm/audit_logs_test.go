package cspm

import (
	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/client"
	"net/http"
	"testing"
)

func TestCspmClient_ListAuditLogs(t *testing.T) {
	teardown := setup()
	defer teardown()
	c, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     server.URL,
		Schema:     "http",
		MaxRetries: 3,
	})
	assert.Nil(t, err)
	mux.HandleFunc(auditLogEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"timestamp":1654631968170,"user":"user1","ipAddress":"1.2.3.4","resourceName":"user1Resource","resourceType":"Login","action":"User1 (Cloud Provisioning Admin) logged in via access key.","result":"Successful"},{"timestamp":1654631968170,"user":"user2","ipAddress":"4.3.2.1","resourceName":"user2Resource","resourceType":"Secure - Policy","action":"User2(with role 'System Admin':'System Admin') updated the policy. Policy rule was updated","result":"Successful"},{"timestamp": 1654626935355,"user": "Prisma Cloud System Admin","ipAddress": "RedLock Internal IP", "resourceName": "Remediation", "resourceType": "Alerts", "action": "'Prisma Cloud System Admin' failed to auto-remediate alerts [[{\"resource_name\":datasec-asm2, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}, {\"resource_name\":aws-cloudtrail-logs-311657289423-35de3571, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}, {\"resource_name\":myversioningalerttestwar, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}, {\"resource_name\":newtestbucketprismas3, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}, {\"resource_name\":anothertestbuckerforar-prisma, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}, {\"resource_name\":buildkite-managedsecretsbucket-n520j4s8rmmv, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}, {\"resource_name\":buildkite-managedsecretsloggingbucket-1umu6v1wtxdc4, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}, {\"resource_name\":anotherawsbuckettest, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}, {\"resource_name\":aws-qs-bucket, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}, {\"resource_name\":aws-flowlog-bucket-test, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}, {\"resource_name\":datasec-asm, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}, {\"resource_name\":mybucket-asm, \"policy_id\":\"13e6067e-b4a1-42ea-970e-e079fb61d79f\"}]].","result": "Failed"}]`))
	})

	auditLogs, err := c.ListAuditLogs(&AuditLogQuery{
		TimeType:   "relative",
		TimeAmount: "24",
		TimeUnit:   "hours",
	})
	assert.Nil(t, err)
	assert.Equal(t, "1.2.3.4", auditLogs[0].IpAddress)
	assert.Equal(t, "user1Resource", auditLogs[0].ResourceName)
	assert.Equal(t, "Login", auditLogs[0].ResourceType)
	assert.Equal(t, "Successful", auditLogs[0].Result)
	assert.Equal(t, "user1", auditLogs[0].User)
	assert.Equal(t, "Alerts", auditLogs[2].ResourceType)
	assert.Equal(t, "Failed", auditLogs[2].Result)
	assert.Equal(t, "Remediation", auditLogs[2].ResourceName)
	assert.Equal(t, "RedLock Internal IP", auditLogs[2].IpAddress)
	assert.NotNil(t, auditLogs)
}

func TestCspmClient_ListAuditLogsFailure(t *testing.T) {
	teardown := setup()
	defer teardown()
	c, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     server.URL,
		Schema:     "http",
		MaxRetries: 3,
	})
	assert.Nil(t, err)
	mux.HandleFunc(auditLogEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`blah`))
	})
	auditLogs, err := c.ListAuditLogs(&AuditLogQuery{
		TimeType:   "relative",
		TimeAmount: "24",
		TimeUnit:   "hours",
	})
	assert.Nil(t, auditLogs)
	assert.Error(t, err)
}
