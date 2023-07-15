CREATE TABLE IF NOT EXISTS servers (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL
);

INSERT INTO servers (id, name, type, status)
VALUES
    ('a5cb11fd-03cd-410d-94df-06fb328f6573', 'Atlas', 'small', 'stopped'),
    ('7bf55aeb-d0c1-4f29-a6e2-fe0a8854cb34', 'Cyprus', 'medium', 'starting'),
    ('7951302e-6690-46f8-aeb1-9488e11a483f', 'Orla', 'small', 'stopping'),
    ('b1677a4d-66b8-45af-a2a1-408fecdc5d0f', 'Nexus', 'large', 'running');