CREATE TABLE alerts(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    camera_id UUID REFERENCES cameras(id),
    alert_type TEXT NOT NULL,
    severity TEXT DEFAULT 'INFO',
    message TEXT,
    is_read BOOLEAN DEFAULT FALSE,
    timestamp TIMESTAMP DEFAULT NOW()
)