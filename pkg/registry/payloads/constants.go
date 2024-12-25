package payloads

const (
	DemoPayload PayloadType = "demo-payload"
	BasicCallback01 PayloadType = "basic-callback-01"
	MaliciousPayload01 PayloadType = "malicious-payload-01"
)

var GithubHeaders = map[string]string{
	"X-GitHub-Api-Version": "2022-11-28",
	"Accept": "application/vnd.github+json",
}