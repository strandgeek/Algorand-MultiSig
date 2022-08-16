import { Account } from "./account";

export interface SignedTransaction {
  id: number;
  signer: Account
  created_at: string
  updated_at: string
}
