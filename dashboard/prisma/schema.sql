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
CREATE INDEX IF NOT EXISTS "Event_data_ProfileId_idx" ON "Event" ((data -> 'ProfileId'));
CREATE INDEX IF NOT EXISTS "Event_type_idx" ON "Event" ("type");
CREATE INDEX IF NOT EXISTS "Event_timestamp_idx" ON "Event" ("timestamp");
CREATE INDEX idx_event_id_profile ON "Event" (id DESC, (data->'ProfileId'));


CREATE TABLE IF NOT EXISTS "MomokaTx" (
  "id" SERIAL PRIMARY KEY,
  "proofTxId" TEXT UNIQUE NOT NULL,
  "success" BOOLEAN,
  "dataAvailabilityId" TEXT UNIQUE NOT NULL,
  "publicationId" TEXT UNIQUE NOT NULL,
  "signature" TEXT UNIQUE NOT NULL,
  "timestampProofs" JSONB,
  "chainProofs" JSONB,
  "timestamp" BIGINT,
  "type" TEXT,
  "event" JSONB
);

CREATE INDEX IF NOT EXISTS "MomokaTx_proofTxId_idx" ON "MomokaTx" ("proofTxId");
CREATE INDEX IF NOT EXISTS "MomokaTx_event_ProfileId_idx" ON "MomokaTx" ((event -> 'ProfileId'));
CREATE INDEX IF NOT EXISTS "MomokaTx_type_idx" ON "MomokaTx" ("type");
CREATE INDEX idx_momokatx_id_profile ON "MomokaTx" (id DESC, (event->'ProfileId'));



CREATE TABLE IF NOT EXISTS "LastBlock" (
    "id" INT PRIMARY KEY,
    "blockNumber" INT
);

-- materialized view 
-- postgresql.conf
-- shared_preload_libraries = 'pg_cron'    # (change requires restart)
-- CREATE EXTENSION pg_cron
CREATE MATERIALIZED VIEW daily_transaction_count AS
SELECT type, TO_CHAR(to_timestamp("timestamp"), 'YYYY-MM-DD') AS date, COUNT(*) AS count
FROM "MomokaTx"
GROUP BY type, date
ORDER BY date DESC;

CREATE UNIQUE INDEX daily_transaction_count_pk ON daily_transaction_count (type, date);

CREATE OR REPLACE FUNCTION refresh_daily_transaction_count() RETURNS void AS $$
BEGIN
    REFRESH MATERIALIZED VIEW CONCURRENTLY daily_transaction_count;
END;
$$ LANGUAGE plpgsql;

SELECT cron.schedule_in_database('hourly-refresh-daily-transaction-count', '0 * * * *', 'SELECT refresh_daily_transaction_count()', 'lenscan');  -- every hour
-- 0 * * * * psql postgres://postgres:password@localhost:5432/lenscan -c "SELECT refresh_daily_transaction_count()"




-- for lens profile
CREATE MATERIALIZED VIEW publications_summary_by_date AS
SELECT
    DATE_TRUNC('day', to_timestamp("timestamp"))::date AS "mv_date",
    "data"->>'ProfileId' AS "profileId",
    COUNT(*) AS "mv_count",
    CASE
        WHEN COUNT(*) > 4 THEN 4
        ELSE COUNT(*)
    END AS "mv_level"
FROM
    "Event"
WHERE
    "type" IN ('PostCreated', 'CommentCreated', 'MirrorCreated')
GROUP BY "mv_date", "profileId"
ORDER BY "mv_date", "profileId";

CREATE UNIQUE INDEX publications_summary_by_date_pk ON publications_summary_by_date (mv_date, "profileId");

CREATE OR REPLACE FUNCTION refresh_publications_summary_by_date() RETURNS void AS $$
BEGIN
    REFRESH MATERIALIZED VIEW CONCURRENTLY publications_summary_by_date;
END;
$$ LANGUAGE plpgsql;

SELECT cron.schedule_in_database('hourly-refresh-publications-summary-by-date', '0 * * * *', 'SELECT refresh_publications_summary_by_date()', 'lenscan');  -- every hour
-- 0 * * * * psql postgres://postgres:password@localhost:5432/lenscan -c "SELECT refresh_publications_summary_by_date()"

CREATE OR REPLACE FUNCTION get_publications_summary(
    profile_id bigint,
    start_date timestamp without time zone,
    end_date timestamp without time zone
) RETURNS TABLE (date date, count bigint, level bigint) LANGUAGE plpgsql IMMUTABLE AS $function$
BEGIN
    RETURN QUERY 
    SELECT
        "mv_date" AS date,
        "mv_count" AS count,
        "mv_level" AS level
    FROM publications_summary_by_date WHERE "profileId"::bigint = profile_id AND "mv_date" >= start_date::date AND "mv_date" < end_date::date;
END;
$function$;
