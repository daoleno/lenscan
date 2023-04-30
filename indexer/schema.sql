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
CREATE INDEX IF NOT EXISTS "Event_type_idx" ON "Event" ("type");


CREATE TABLE IF NOT EXISTS "LastBlock" (
    "id" INT PRIMARY KEY,
    "blockNumber" INT
);


-- for lens profile
CREATE OR REPLACE FUNCTION public.get_publications_summary(
    profile_id bigint,
    start_date timestamp without time zone,
    end_date timestamp without time zone
) RETURNS TABLE (date date, count bigint, level bigint) LANGUAGE plpgsql IMMUTABLE AS $function$
BEGIN
    RETURN QUERY SELECT
        DATE_TRUNC('day', to_timestamp("timestamp"))::date AS "date",
        COUNT(*) AS "count",
        CASE
            WHEN COUNT(*) > 4 THEN 4
            ELSE COUNT(*)
        END AS "level"
    FROM
        "Event"
    WHERE
        "type" IN ('PostCreated', 'CommentCreated', 'MirrorCreated')
        AND "data"->>'ProfileId' = profile_id::text
        AND to_timestamp("timestamp") >= start_date
        AND to_timestamp("timestamp") < end_date
    GROUP BY "date"
    ORDER BY "date";
END;
$function$