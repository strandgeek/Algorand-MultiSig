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
    // TODO: add loading spinner
    return <LoadingSpinner />
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
        <div className="overflow-x-auto w-full">
          <MultiSigAccountsTable multiSigAccounts={multiSigAccounts} />
        </div>
      </div>
    </AppLayout>
  );
};
