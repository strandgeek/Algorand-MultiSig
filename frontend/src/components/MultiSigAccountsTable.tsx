import React, { FC } from "react";
import { MultiSigAccount } from "../types/multisigAccount";
import { getIdenticonSrc } from "../utils/getIdenticonSrc";
import { getShortAddress } from "../utils/getShortAddress";
import { AddressAvatarsStack } from "./AddressAvatarsStack";

export interface MultiSigAccountsTableRowProps {
  multiSigAccount: MultiSigAccount;
}

export const MultiSigAccountsTableRow: FC<MultiSigAccountsTableRowProps> = ({
  multiSigAccount: { address, accounts },
}) => {
  return (
    <tr className="cursor-pointer" onClick={() => null}>
      <td>
        <div className="flex items-center space-x-3">
          <div className="avatar">
            <div className="rounded-full w-12 h-12">
              <img
                src={getIdenticonSrc(address)}
                width={160}
                height={160}
                alt={`MultiSig Address`}
              />
            </div>
          </div>
          <div>
            <div className="text-sm opacity-70">{getShortAddress(address)}</div>
          </div>
        </div>
      </td>
      <td>
        <AddressAvatarsStack addresses={accounts.map(acc => acc.address)} />
      </td>
      <td>0 txns</td>
      <td className="text-right">
        <span className="font-normal opacity-70 text-sm">20h ago</span>
      </td>
    </tr>
  );
};

export interface MultiSigAccountsTableProps {
  multiSigAccounts: MultiSigAccount[];
}

export const MultiSigAccountsTable: FC<MultiSigAccountsTableProps> = ({ multiSigAccounts }) => {
  return (
    <table className="table w-full">
      <thead>
        <tr>
          <th>Name</th>
          <th>Signers</th>
          <th>Transactions</th>
          <th className="text-right">Last Update</th>
        </tr>
      </thead>
      <tbody>
        {multiSigAccounts.map(msa => <MultiSigAccountsTableRow key={msa.id} multiSigAccount={msa} />)}
      </tbody>
    </table>
  );
};
