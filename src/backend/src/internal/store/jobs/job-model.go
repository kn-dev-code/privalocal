package jobs

import "time"

type JobStatus string

const (
    JobStatusQueued    JobStatus = "queued"
    JobStatusRunning   JobStatus = "running"
    JobStatusCompleted JobStatus = "completed"
    JobStatusFailed    JobStatus = "failed"
    JobStatusCanceled  JobStatus = "canceled"
)

type JobConfig struct {
    Method      string `json:"method"`
    DurationSec int    `json:"duration_sec"`
    RPS         int    `json:"rps,omitempty"`
    Concurrency int    `json:"concurrency,omitempty"`
}

type Job struct {
    ID         int64      `json:"id"          db:"id"`
    UserID     int64      `json:"-"           db:"user_id"`
    Target     string     `json:"target"      db:"target"`
    Config     JobConfig  `json:"config"      db:"config"`
    Status     JobStatus  `json:"status"      db:"status"`
    Error      *string    `json:"error"       db:"error"`
    CreatedAt  time.Time  `json:"created_at"  db:"created_at"`
    StartedAt  *time.Time `json:"started_at"  db:"started_at"`
    FinishedAt *time.Time `json:"finished_at" db:"finished_at"`
}

type LatencySummary struct {
    Requests  int     `json:"requests"`
    Errors    int     `json:"errors"`
    P50Ms     float64 `json:"p50_ms"`
    P95Ms     float64 `json:"p95_ms"`
    P99Ms     float64 `json:"p99_ms"`
    AvgTTFBMs float64 `json:"avg_ttfb_ms"`
    AvgDNSMs  float64 `json:"avg_dns_ms"`
    AvgTLSMs  float64 `json:"avg_tls_ms"`
}
type JobResult struct {
    JobID       int64          `json:"job_id"       db:"job_id"`
    Summary     LatencySummary `json:"summary"      db:"summary"`
    Timeseries  []SecondBucket `json:"timeseries"   db:"timeseries"`
    HostMetrics []HostSample   `json:"host_metrics" db:"host_metrics"`
    CreatedAt   time.Time      `json:"created_at"   db:"created_at"`
}

type SecondBucket struct {
	Second int `json:"second"`
	Requests int `json:"requests"`
	Errors int `json:"errors"`
	P50Ms float64 `json:"p50_ms"`
	P95Ms float64 `json:"p95_ms"`
}

type HostSample struct {
	Second int `json:"second"`
	CPUPct float64 `json:"cpu_pct"`
	MemPct float64 `json:"mem_pct"`
}