import { SignedTransaction } from "./signedTransaction";

export interface Transaction {
  id: number;
  txn_id: string;
  raw_transaction: string;
  signed_transactions: SignedTransaction[]
  signed_transactions_count: number;
  status: "PENDING" | "READY" | "BROADCASTING" | "BROADCASTED" | "FAILED";
}
