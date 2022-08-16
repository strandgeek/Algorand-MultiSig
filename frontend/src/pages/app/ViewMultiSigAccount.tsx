import { InboxIcon } from "@heroicons/react/outline";
import React, { FC } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import {
  useMultiSigAccountQuery,
  useMultiSigAccountTransactionsQuery,
} from "../../client/queries";
import { LoadingSpinner } from "../../components/LoadingSpinner";
import { MultiSigAccountDetails } from "../../components/MultiSigAccountDetails";
import { MultiSigAccountSignersList } from "../../components/MultiSigAccountSignersList";
import { TransactionsTable } from "../../components/TransactionsTable";
import { AppLayout } from "../../layouts/AppLayout";
import { getShortAddress } from "../../utils/getShortAddress";

export interface ViewMultiSigAccountProps {}

export const ViewMultiSigAccount: FC<ViewMultiSigAccountProps> = props => {
  const navigate = useNavigate();
  const { msaAddress } = useParams();
  const { data: multiSigAccount } = useMultiSigAccountQuery(msaAddress);
  const { data: transactions } =
    useMultiSigAccountTransactionsQuery(msaAddress);
  if (!multiSigAccount) {
    return <LoadingSpinner />;
  }
  return (
    <AppLayout>
      <div className="mx-auto max-w-4xl mt-8">
        <div className="text-sm breadcrumbs mb-4">
          <ul>
            <li>
              <Link to="/app/multisig-accounts">MultiSig Accounts</Link>
            </li>
            <li>
              MultiSig Account ({getShortAddress(multiSigAccount.address)})
            </li>
          </ul>
        </div>
        <div className="font-bold text-xl mb-4">MultiSig Account Details</div>
        <div className="card bg-base-100 mb-8">
          <MultiSigAccountDetails multiSigAccount={multiSigAccount} />
        </div>

        <div className="font-bold text-xl mb-4">Signers</div>
        <div className="card bg-base-100 p-2 px-4 mb-8">
          <MultiSigAccountSignersList multiSigAccount={multiSigAccount} />
        </div>

        <div className="flex justify-between items-center mb-4">
          <h3 className="text-xl font-semibold">Transactions</h3>
          <button
            className="btn btn-primary"
            onClick={() =>
              navigate(
                `/app/multisig-accounts/${msaAddress}/transactions/create`
              )
            }
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-4 w-4 mr-2"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              strokeWidth={2}
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M12 4v16m8-8H4"
              />
            </svg>
            Create Transaction
          </button>
        </div>

        <div className="overflow-x-auto w-full">
          {transactions && transactions.length > 0 && (
            <TransactionsTable
              transactions={transactions}
              multiSigAccount={multiSigAccount}
            />
          )}
          {transactions && transactions.length === 0 && (
            <div className="relative block w-full border-2 border-gray-300 border-dashed rounded-lg p-12 text-center">
              <InboxIcon className="mx-auto h-12 w-12 text-gray-400" />
              <h3 className="mt-2 text-sm font-medium text-gray-900">
                No transactions
              </h3>
              <p className="mt-4 text-sm text-gray-500">
                No transaction have been created yet
              </p>
            </div>
          )}
        </div>
      </div>
    </AppLayout>
  );
};
