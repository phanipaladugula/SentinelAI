CREATE TABLE faces(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    embedding vector(512),
    created_at TIMESTAMP DEFAULT NOW() 
)