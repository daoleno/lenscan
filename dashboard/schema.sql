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


CREATE INDEX idx_poamh_currency ON public.publication_open_action_module_history (currency);
CREATE INDEX idx_poamh_history_id ON public.publication_open_action_module_history (history_id);
CREATE INDEX idx_poamh_publication_id ON public.publication_open_action_module_history (publication_id);
CREATE INDEX idx_poamr_acted_profile_id ON public.publication_open_action_module_acted_record (acted_profile_id);
CREATE INDEX idx_poamr_history_id ON public.publication_open_action_module_acted_record (history_id);
CREATE INDEX idx_poamr_block_timestamp ON public.publication_open_action_module_acted_record (block_timestamp DESC);
CREATE INDEX idx_poamr_transaction_executor ON public.publication_open_action_module_acted_record (transaction_executor);
-- Indexer tables

CREATE TABLE lens_profile_created_transaction (
    id SERIAL PRIMARY KEY,
    block_number BIGINT NOT NULL,
    block_timestamp TIMESTAMPTZ NOT NULL,
    transaction_hash TEXT NOT NULL,
    from_address TEXT NOT NULL,
    to_address TEXT NOT NULL,
    value NUMERIC NOT NULL,
    
    CONSTRAINT uniq_transaction_hash UNIQUE (transaction_hash)
);

CREATE INDEX idx_lens_profile_created_trans_to_address_value_timestamp ON lens_profile_created_transaction(to_address, value, block_timestamp DESC);
CREATE INDEX idx_lens_profile_created_trans_timestamp_to_address_value ON lens_profile_created_transaction(block_timestamp, to_address, value);
CREATE INDEX idx_lens_profile_created_trans_value_timestamp ON lens_profile_created_transaction(value, block_timestamp DESC);