import { FC } from "react";
import { useMultiSigAccounts } from "../client/queries";
import { MultiSigAccountsTable } from "../components/MultiSigAccountsTable";
import { AppLayout } from "../layouts/AppLayout";

interface HomeProps {}


export const Home: FC<HomeProps> = () => {
  const { data: multiSigAccounts } = useMultiSigAccounts()
  if (!multiSigAccounts) {
    // TODO: add loading spinner
    return <div>Loading</div>
  }
  return (
    <AppLayout>
      <div className="mx-auto max-w-4xl mt-8">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-xl font-semibold">MultiSig Accounts</h3>
          <button className="btn btn-md btn-primary">Create</button>
        </div>
        <div className="overflow-x-auto w-full">
          <MultiSigAccountsTable multiSigAccounts={multiSigAccounts} />
        </div>
      </div>
    </AppLayout>
  );
};
