CREATE IF NOT EXISTS TABLE job_table (
    ID BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES user_table(ID) ON DELETE CASCADE,
    target TEXT NOT NULL,
    config JSONB NOT NULL,
    status TEXT NOT NULL DEFAULT 'queued'
    CHECK (status IN ('queued', 'running', 'completed','failed','canceled')) ,
    error TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    started_at TIMESTAMPTZ,
    finished_at TIMESTAMPTZ,
)

CREATE INDEX idx_job_table ON job_table(user_id, created_at DESC);

CREATE IF NOT EXISTS job_results (
    ID BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    summary JSONB NOT NULL,
    timeseries JSONB NOT NULL,
    host_metrics JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)

DROP TABLE job_results
DROP TABLE job_table