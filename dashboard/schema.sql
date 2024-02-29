-- create index for the following tables

CREATE INDEX idx_publication_metadata_language ON publication_metadata(language);
CREATE INDEX idx_publication_metadata_app ON publication_metadata(app);
CREATE INDEX idx_publication_metadata_timestamp_app ON publication_metadata(timestamp, app);


CREATE INDEX idx_profile_record_profile_id ON profile_record(profile_id);

CREATE INDEX idx_publication_record_publication_id ON publication_record(publication_id);
CREATE INDEX idx_publication_record_publication_type ON publication_record(publication_type);
CREATE INDEX idx_publication_record_block_timestamp ON publication_record(block_timestamp);
CREATE INDEX idx_publication_record_profile_id_block_timestamp ON publication_record(profile_id, block_timestamp DESC);

CREATE INDEX idx_publication_hashtag_hashtag ON publication_hashtag(hashtag);

CREATE INDEX idx_publication_reaction_type ON publication_reaction(type);

