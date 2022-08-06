import { FC } from "react";
import { MultiSigAccountsTable } from "../components/MultiSigAccountsTable";
import { AppLayout } from "../layouts/AppLayout";
import { MultiSigAccount } from "../types/multisigAccount";

interface HomeProps {}


const MULTISIG_ACCOUNTS_EXAMPLE: MultiSigAccount[] = [
  {
    id: 1,
    version: 1,
    threshold: 2,
    address: 'N465KSFSNT3JA5G45TBBQJDJ7LQOM2STEATQYXJZRAOEHHGNV6DT4AUF2M',
    accounts: [
      {
        id: 1,
        address: 'DNKRBWSS3GF57WGRR7OMG6DUS2EHPMYLJ364LOCJG6C5FBZKX4DMKVKS22',
      },
      {
        id: 2,
        address: '3ZWWK2AFVQFW2G7ZTXWRXUPGFQP6AZHXI6SHHD3664OCMWCK3MSUEXQ333',
      },
    ]
  }
]


export const Home: FC<HomeProps> = () => {
  // TODO: Fetch data from API
  const multiSigAccounts = MULTISIG_ACCOUNTS_EXAMPLE
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
