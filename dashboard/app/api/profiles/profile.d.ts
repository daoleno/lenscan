export type Profile = {
  profile_id: string
  owned_by: string
  is_burnt: boolean
  tx_hash: string
  block_hash: string
  block_number: number
  log_index: number
  tx_index: number
  block_timestamp: Date
  source_timestamp: bigint
}
