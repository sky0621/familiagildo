CREATE TABLE task (
    id BIGSERIAL PRIMARY KEY,
    node_id TEXT NOT NULL,
    content TEXT NOT NULL,
    reward TEXT,
    incentive TEXT
);