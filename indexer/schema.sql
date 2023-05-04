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
CREATE INDEX idx_event_id_profile ON "Event" (id, (data->'ProfileId'));




CREATE TABLE IF NOT EXISTS "LastBlock" (
    "id" INT PRIMARY KEY,
    "blockNumber" INT
);


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

REFRESH MATERIALIZED VIEW publications_summary_by_date;

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
