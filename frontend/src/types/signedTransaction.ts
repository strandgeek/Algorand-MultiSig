import { Account } from "./account";
import { Transaction } from "./transaction";

export interface SignedTransaction {
  id: number;
  account: Account
  transaction: Transaction
}
