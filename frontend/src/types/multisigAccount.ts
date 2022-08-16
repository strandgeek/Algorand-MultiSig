import { Account } from "./account"

export interface MultiSigAccount {
  id: number
  version: number
  threshold: number
  accounts: Account[]
  address: string
  created_at: string
  updated_at: string
}
