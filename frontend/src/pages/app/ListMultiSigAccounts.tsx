import { KeyIcon } from "@heroicons/react/outline";
import { FC } from "react";
import { Link } from "react-router-dom";
import { useMultiSigAccountsQuery } from "../../client/queries";
import { LoadingSpinner } from "../../components/LoadingSpinner";
import { MultiSigAccountsTable } from "../../components/MultiSigAccountsTable";
import { AppLayout } from "../../layouts/AppLayout";

interface ListMultiSigAccountsProps {}

export const ListMultiSigAccounts: FC<ListMultiSigAccountsProps> = () => {
  const { data: multiSigAccounts } = useMultiSigAccountsQuery();

  if (!multiSigAccounts) {
    return <LoadingSpinner />;
  }
  return (
    <AppLayout>
      <div className="mx-auto max-w-4xl mt-8">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-xl font-semibold">MultiSig Accounts</h3>
          <Link
            to="/app/multisig-accounts/create"
            className="btn btn-md btn-primary"
          >
            Create
          </Link>
        </div>
        {multiSigAccounts.length > 0 ? (
          <div className="overflow-x-auto w-full">
            <MultiSigAccountsTable multiSigAccounts={multiSigAccounts} />
          </div>
        ) : (
          <div className="relative block w-full border-2 border-gray-300 border-dashed rounded-lg p-12 text-center">
            <KeyIcon className="mx-auto h-12 w-12 text-gray-400" />
            <h3 className="mt-2 text-sm font-medium text-gray-900">
              No MultiSig Accounts
            </h3>
            <p className="mt-4 text-sm text-gray-500">
              No MultiSig Account have been created yet
            </p>
          </div>
        )}
      </div>
    </AppLayout>
  );
};
