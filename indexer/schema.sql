CREATE TABLE IF NOT EXISTS "Event" (
    "id" SERIAL PRIMARY KEY,
    "blockNumber" BIGINT,
    "txHash" VARCHAR(66),
    "txIndex" INT,
    "logIndex" INT,
    "removed" BOOLEAN,
    "timestamp" BIGINT,
    "type" TEXT,
    "data" JSONB
);

CREATE UNIQUE INDEX IF NOT EXISTS "Event_blockNumber_logIndex" ON "Event" ("blockNumber", "logIndex");
CREATE INDEX IF NOT EXISTS "Event_data_ProfileId_idx" ON "Event" ((data ->> 'ProfileId'));


CREATE TABLE IF NOT EXISTS "LastBlock" (
    "id" INT PRIMARY KEY,
    "blockNumber" INT
);