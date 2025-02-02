import React, { FC } from "react";
import moment from 'moment'
import { useNavigate } from "react-router-dom";
import { MultiSigAccount } from "../types/multisigAccount";
import { getIdenticonSrc } from "../utils/getIdenticonSrc";
import { getShortAddress } from "../utils/getShortAddress";
import { AddressAvatarsStack } from "./AddressAvatarsStack";

export interface MultiSigAccountsTableRowProps {
  multiSigAccount: MultiSigAccount;
}

export const MultiSigAccountsTableRow: FC<MultiSigAccountsTableRowProps> = ({
  multiSigAccount: { address, accounts, updated_at },
}) => {
  const navigate = useNavigate()
  const lastUpdate = moment(updated_at).fromNow();
  return (
    <tr className="cursor-pointer" onClick={() => navigate(`/app/multisig-accounts/${address}`)}>
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
      <td className="text-right">
        <span className="font-normal opacity-70 text-sm">{lastUpdate}</span>
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
          <th className="text-right">Last Update</th>
        </tr>
      </thead>
      <tbody>
        {multiSigAccounts.map(msa => <MultiSigAccountsTableRow key={msa.id} multiSigAccount={msa} />)}
      </tbody>
    </table>
  );
};
