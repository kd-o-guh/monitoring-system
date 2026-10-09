CREATE TABLE metrics (
    id          BIGSERIAL PRIMARY KEY,
    hostname    TEXT NOT NULL,
    cpu         DOUBLE PRECISION NOT NULL CHECK (cpu BETWEEN 0 AND 100),
    memory      DOUBLE PRECISION NOT NULL CHECK (memory BETWEEN 0 AND 100),
    disk        DOUBLE PRECISION NOT NULL CHECK (disk BETWEEN 0 AND 100),
    collected_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_metrics_collected_at ON metrics (collected_at DESC, id DESC);
CREATE INDEX idx_metrics_hostname_collected_at
    ON metrics (hostname, collected_at DESC, id DESC);
