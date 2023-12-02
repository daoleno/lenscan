export type Publication = {
  publication_id: string
  contract_publication_id: string
  publication_type: string
  profile_id: string
  content_uri: string
  parent_publication_id: string
  root_publication_id: string
  is_hidden: boolean
  is_momoka: boolean
  momoka_proof: string
  app: string
  gardener_flagged: boolean
  transaction_executor: string
  tx_hash: string
  block_hash: string
  block_number: number
  log_index: number
  tx_index: number
  block_timestamp: Date
  source_timestamp: number
}
