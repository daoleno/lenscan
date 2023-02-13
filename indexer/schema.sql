
CREATE TABLE IF NOT EXISTS "Event" (
    "id" SERIAL PRIMARY KEY,
    "blockNumber" BIGINT,
    "txHash" VARCHAR(66),
    "txIndex" INT,
    "logIndex" INT,
    "removed" BOOLEAN,
    "event" TEXT
);


CREATE TABLE IF NOT EXISTS "BaseInitialized" (
    "event_id" INT REFERENCES "Event"("id"),
    "name" TEXT,
    "symbol" TEXT,
    "timestamp" BIGINT
);


CREATE TABLE IF NOT EXISTS "StateSet" (
    "event_id" INT REFERENCES "Event"("id"),
    "caller" VARCHAR(42),
    "prevState" TEXT,
    "newState" TEXT,
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "GovernanceSet" (
    "event_id" INT REFERENCES "Event"("id"),
    "caller" VARCHAR(42),
    "prevGovernance" VARCHAR(42),
    "newGovernance" VARCHAR(42),
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "EmergencyAdminSet" (
    "event_id" INT REFERENCES "Event"("id"),
    "caller" VARCHAR(42),
    "oldEmergencyAdmin" VARCHAR(42),
    "newEmergencyAdmin" VARCHAR(42),
    "timestamp" BIGINT
);


CREATE TABLE IF NOT EXISTS "ProfileCreatorWhitelisted" (
    "event_id" INT REFERENCES "Event"("id"),
    "profileCreator" VARCHAR(42),
    "whitelisted" BOOLEAN,
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "FollowModuleWhitelisted" (
    "event_id" int references "Event"("id"),
    "followModule" varchar(42),
    "whitelisted" boolean,
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "ReferenceModuleWhitelisted" (
    "event_id" int references "Event"("id"),
    "referenceModule" varchar(42),
    "whitelisted" boolean,
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "CollectModuleWhitelisted" (
    "event_id" int references "Event"("id"),
    "collectModule" varchar(42),
    "whitelisted" boolean,
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "ProfileCreated" (
    "event_id" int references "Event"("id"),
    "profileId" bigint,
    "creator" varchar(42),
    "to" varchar(42),
    "handle" text,
    "imageURI" text,
    "followModule" varchar(42),
    "followModuleReturnData" bytea,
    "followNFTURI" text,
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "DefaultProfileSet" (
    "event_id" int references "Event"("id"),
    "wallet" varchar(42),
    "profileId" bigint,
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "DispatcherSet" (
    "event_id" int references "Event"("id"),
    "profileId" bigint,
    "dispatcher" varchar(42),
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "ProfileImageURISet" (
    "event_id" int references "Event"("id"),
    "profileId" bigint,
    "imageURI" text,
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "FollowNFTURISet" (
    "event_id" int references "Event"("id"),
    "profileId" bigint,
    "followNFTURI" text,
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "FollowModuleSet" (
    "event_id" int references "Event"("id"),
    "profileId" bigint,
    "followModule" varchar(42),
    "followModuleReturnData" bytea,
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "PostCreated" (
    "event_id" int references "Event"("id"),
    "profileId" bigint,
    "pubId" bigint,
    "contentURI" text,
    "collectModule" varchar(42),
    "collectModuleReturnData" bytea,
    "referenceModule" varchar(42),
    "referenceModuleReturnData" bytea,
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "CommentCreated" (
    "event_id" int references "Event"("id"),
    "profileId" bigint,
    "pubId" bigint,
    "contentURI" text,
    "profileIdPointed" bigint,
    "pubIdPointed" bigint,
    "referenceModuleData" bytea,
    "collectModule" varchar(42),
    "collectModuleReturnData" bytea,
    "referenceModule" varchar(42),
    "referenceModuleReturnData" bytea,
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "MirrorCreated" (
    "event_id" int references "Event"("id"),
    "profileId" bigint,
    "pubId" bigint,
    "profileIdPointed" bigint,
    "pubIdPointed" bigint,
    "referenceModuleData" bytea,
    "referenceModule" varchar(42),
    "referenceModuleReturnData" bytea,
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "FollowNFTDeployed" (
    "event_id" int references "Event"("id"),
    "profileId" bigint,
    "followNFT" varchar(42),
    "timestamp" bigint
);

CREATE TABLE IF NOT EXISTS "CollectNFTDeployed" (
    "event_id" INT REFERENCES "Event"("id"),
    "profileId" BIGINT,
    "pubId" BIGINT,
    "collectNFT" VARCHAR(42),
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "Collected" (
    "event_id" INT REFERENCES "Event"("id"),
    "collector" VARCHAR(42),
    "profileId" BIGINT,
    "pubId" BIGINT,
    "rootProfileId" BIGINT,
    "rootPubId" BIGINT,
    "collectModuleData" BYTEA,
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "Followed" (
    "event_id" INT REFERENCES "Event"("id"),
    "follower" VARCHAR(42),
    "profileIds" BYTEA,
    "followModuleDatas" BYTEA,
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS FollowNFTTransferred (
    "event_id" INT REFERENCES "Event"("id"),
    "profileId" BIGINT,
    "followNFTId" BIGINT,
    "from" VARCHAR(42),
    "to" VARCHAR(42),
    "timestamp" BIGINT
);


CREATE TABLE IF NOT EXISTS "CollectNFTTransferred" (
    "event_id" INT REFERENCES "Event"("id"),
    "profileId" BIGINT,
    "pubId" BIGINT,
    "collectNFTId" BIGINT,
    "from" VARCHAR(42),
    "to" VARCHAR(42),
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "FollowNFTInitialized" (
    "event_id" INT REFERENCES "Event"("id"),
    "profileId" BIGINT,
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "FollowNFTDelegatedPowerChanged" (
    "event_id" INT REFERENCES "Event"("id"),
    "delegate" VARCHAR(42),
    "newPower" BIGINT,
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "CollectNFTInitialized" (
    "event_id" INT REFERENCES "Event"("id"),
    "profileId" BIGINT,
    "pubId" BIGINT,
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "ModuleGlobalsGovernanceSet" (
    "event_id" INT REFERENCES "Event"("id"),
    "prevGovernance" VARCHAR(42),
    "newGovernance" VARCHAR(42),
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "ModuleGlobalsTreasurySet" (
    "event_id" INT REFERENCES "Event"("id"),
    "prevTreasury" VARCHAR(42),
    "newTreasury" VARCHAR(42),
    "timestamp" BIGINT
);


CREATE TABLE IF NOT EXISTS "ModuleGlobalsTreasuryFeeSet" (
    "event_id" INT REFERENCES "Event"("id"),
    "prevTreasuryFee" INT,
    "newTreasuryFee" INT,
    "timestamp" BIGINT
);


CREATE TABLE IF NOT EXISTS "ModuleGlobalsCurrencyWhitelisted" (
    "event_id" INT REFERENCES "Event"("id"),
    "currency" VARCHAR(42),
    "prevWhitelisted" BOOLEAN,
    "whitelisted" BOOLEAN,
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "FeeModuleBaseConstructed" (
    "event_id" INT REFERENCES "Event"("id"),
    "moduleGlobals" VARCHAR(42),
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "ModuleBaseConstructed" (
    "event_id" INT REFERENCES "Event"("id"),
    "hub" VARCHAR(42),
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "FollowsApproved" (
    "event_id" INT REFERENCES "Event"("id"),
    "owner" VARCHAR(42),
    "profileId" BIGINT,
    "addresses" VARCHAR(42)[],
    "approved" BOOLEAN[],
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "FollowsToggled" (
    "event_id" INT REFERENCES "Event"("id"),
    "owner" VARCHAR(42),
    "profileIds" BIGINT[],
    "enabled" BOOLEAN[],
    "timestamp" BIGINT
);

CREATE TABLE IF NOT EXISTS "ProfileMetadataSet" (
    "event_id" INT REFERENCES "Event"("id"),
    "profileId" BIGINT,
    "metadata" TEXT,
    "timestamp" BIGINT
);