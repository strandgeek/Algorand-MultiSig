import { Account } from "./account";

export interface SignedTransaction {
  id: number;
  signer: Account
}
