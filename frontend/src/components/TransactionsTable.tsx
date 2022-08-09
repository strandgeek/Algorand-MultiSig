import React, { FC } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Transaction } from "../types/transaction";
import { getShortAddress } from "../utils/getShortAddress";

export interface TransactionsTableRowProps {
  transaction: Transaction
}

export const TransactionsTableRow: FC<TransactionsTableRowProps> = ({
  transaction,
}) => {
  const params = useParams()
  const navigate = useNavigate();
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
      <td className="text-center">0 of 3</td>
      <td className="text-center">
        <div className="flex justify-center items-center space-x-3 lg:pl-2">
          <div
            className="flex-shrink-0 w-2.5 h-2.5 rounded-full bg-yellow-500"
            aria-hidden="true"
          />
          <a href="#" className="truncate hover:text-gray-600">
            <span>
              <span className="text-gray-500 font-normal">
                Waiting Signatures
              </span>
            </span>
          </a>
        </div>
      </td>
      <td className="text-right">
        <span className="font-normal opacity-70 text-sm">20h ago</span>
      </td>
    </tr>
  );
};

export interface TransactionsTableProps {
  transactions: Transaction[]
}

export const TransactionsTable: FC<TransactionsTableProps> = ({ transactions }) => {
  return (
    <table className="table w-full">
      <thead>
        <tr>
          <th>TxID</th>
          <th className="text-center">Signatures</th>
          <th className="text-center">Status</th>
          <th className="text-right">Last Update</th>
        </tr>
      </thead>
      <tbody>
        {transactions.map(t => <TransactionsTableRow key={t.id} transaction={t} />)}
      </tbody>
    </table>
  );
};
