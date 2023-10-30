CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    content VARCHAR(150) NOT NULL,
    description VARCHAR(200) NOT NULL,
    status BOOLEAN DEFAULT false,
    due timestamp NOT NULL,
    created_at timestamp NOT NULL
);