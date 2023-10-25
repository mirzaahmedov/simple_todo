CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS todos (
    id uuid DEFAULT uuid_generate_v4 (),
    title text NOT NULL,
    content text NOT NULL,
    completed boolean DEFAULT FALSE,
    update_date timestamp,
    create_date timestamp DEFAULT now()
);

