CREATE TABLE IF NOT EXISTS access_logs (
    id SERIAL PRIMARY KEY,
    request_time TIMESTAMP,
    remote_addr TEXT,
    request_method TEXT,
    request_uri TEXT,
    response_status INT,
    response_time FLOAT
);
