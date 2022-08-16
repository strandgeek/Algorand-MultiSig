import moment from "moment";
import React, { FC } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { MultiSigAccount } from "../types/multisigAccount";
import { Transaction } from "../types/transaction";
import { getShortAddress } from "../utils/getShortAddress";
import { StatusLabel } from "./StatusLabel";

export interface TransactionsTableRowProps {
  transaction: Transaction
  multiSigAccount: MultiSigAccount
}

export const TransactionsTableRow: FC<TransactionsTableRowProps> = ({
  transaction,
  multiSigAccount,
}) => {
  const params = useParams()
  const navigate = useNavigate();
  const lastUpdate = moment(transaction.updated_at).fromNow();
  const {
    txn_id,
  } = transaction
  return (
    <tr
      className="cursor-pointer"
      onClick={() => navigate(`/app/multisig-accounts/${params.msaAddress}/transactions/${txn_id}`)}
    >
      <td>
        <div className="flex items-center space-x-3">
          <div>
            <div className="text-sm opacity-70">
              {getShortAddress(txn_id)}
            </div>
          </div>
        </div>
      </td>
      <td className="text-center">{transaction.signed_transactions_count} of {multiSigAccount.threshold}</td>
      <td className="text-center">
        <StatusLabel status={transaction.status} />
      </td>
      <td className="text-center">
        <span className="font-normal opacity-70 text-sm">{lastUpdate}</span>
      </td>
    </tr>
  );
};

export interface TransactionsTableProps {
  transactions: Transaction[]
  multiSigAccount: MultiSigAccount
}

export const TransactionsTable: FC<TransactionsTableProps> = ({ transactions, multiSigAccount }) => {
  return (
    <table className="table w-full">
      <thead>
        <tr>
          <th>TxID</th>
          <th className="text-center">Signatures</th>
          <th className="text-left">Status</th>
          <th className="text-center">Last Update</th>
        </tr>
      </thead>
      <tbody>
        {transactions.map(t => <TransactionsTableRow key={t.id} transaction={t} multiSigAccount={multiSigAccount} />)}
      </tbody>
    </table>
  );
};
