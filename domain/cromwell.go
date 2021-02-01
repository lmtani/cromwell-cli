package domain

import (
	"time"
)

type WorkflowMetadataResponse struct {
	ID             string
	Status         string
	Submission     time.Time
	Start          time.Time
	End            time.Time
	Inputs         map[string]interface{}
	Outputs        map[string]interface{}
	Calls          map[string][]CallMetadata
	WorkflowName   string
	RootWorkflowID string
}

type FailureMessage struct {
	Failure   string
	Timestamp time.Time
}

type CallMetadata struct {
	Inputs            map[string]interface{}
	ExecutionStatus   string
	Backend           string
	BackendStatus     string
	Start             time.Time
	End               time.Time
	JobID             string
	Failures          FailureMessage
	Stdout            string
	Stderr            string
	Attempt           int
	ShardIndex        int
	Labels            Label
	MonitoringLog     string
	CommandLine       string
	DockerImageUsed   string
	SubWorkflowID     string
	RuntimeAttributes RuntimeAttributes
	CallCaching       CallCachingData
	ExecutionEvents   []ExecutionEvents
}

type ExecutionEvents struct {
	StartTime   time.Time
	Description string
	EndTime     time.Time
}

type RuntimeAttributes struct {
	BootDiskSizeGb string
	CPU            string
	Disks          string
	Docker         string
	Memory         string
}

type CallCachingData struct {
	Result string
	Hit    bool
}

type Label struct {
	CromwellWorkflowID string `json:"cromwell-workflow-id"`
	WdlTaskName        string `json:"wdl-task-name"`
}

type SubmitResponse struct {
	ID     string
	Status string
}

type WorkflowQueryResponse struct {
	Results           []QueryResponseWorkflow
	TotalResultsCount int
}

type QueryResponseWorkflow struct {
	ID                    string
	Name                  string
	Status                string
	Submission            string
	Start                 time.Time
	End                   time.Time
	MetadataArchiveStatus string
}

type CromwellUsecase interface {
	// Submit(ctx context.Context) SubmitResponse
	// Kill(ctx context.Context) SubmitResponse
	Query() (WorkflowQueryResponse, error)
}

type CromwellRepository interface {
	// Submit(ctx context.Context) SubmitResponse
	// Kill(ctx context.Context) SubmitResponse
	Query() (WorkflowQueryResponse, error)
}
