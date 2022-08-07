import React, { FC } from "react";
import { useMeQuery } from "../client/queries";
import { MultiSigAccount } from "../types/multisigAccount";
import { getIdenticonSrc } from "../utils/getIdenticonSrc";
import { getShortAddress } from "../utils/getShortAddress";

export interface MultiSigAccountSignersListProps {
  multiSigAccount: MultiSigAccount
}

export const MultiSigAccountSignersList: FC<
  MultiSigAccountSignersListProps
> = ({ multiSigAccount }) => {
  const { data: me } = useMeQuery()
  return (
    <>
      {multiSigAccount.accounts.map((acc, idx) => (
        <div
          key={acc.id}
          className="grid grid-cols-2 border-b border-b-base-300 last:border-none py-4"
        >
          <div className="flex items-center">
            <div className="avatar">
              <div className="rounded-full w-6 h-6">
                <img
                  src={getIdenticonSrc(
                    acc.address
                  )}
                  width={160}
                  height={160}
                  alt={`MultiSig Address`}
                />
              </div>
            </div>
            <div className="ml-4 text-sm">
              {getShortAddress(
                acc.address
              )}
            </div>
          </div>
          <div className="flex items-center justify-end">
            {acc.address ===  me?.address && <span>You</span>}
          </div>
        </div>
      ))}
    </>
  );
};
