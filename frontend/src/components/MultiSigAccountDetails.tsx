import React, { FC } from "react";
import { MultiSigAccount } from "../types/multisigAccount";
import { getIdenticonSrc } from "../utils/getIdenticonSrc";
import { getShortAddress } from "../utils/getShortAddress";

export interface MultiSigAccountDetailsProps {
  multiSigAccount: MultiSigAccount
}

export const MultiSigAccountDetails: FC<
  MultiSigAccountDetailsProps
> = ({ multiSigAccount }) => {
  return (
    <div className="grid grid-cols-2  p-4">
      <div className="flex items-center">
        <div className="avatar">
          <div className="rounded-full w-12 h-12">
            <img
              src={getIdenticonSrc(
                multiSigAccount.address,
              )}
              width={160}
              height={160}
              alt={`MultiSig Address`}
            />
          </div>
        </div>
        <div className="ml-4">
          <div>
            {getShortAddress(
              multiSigAccount.address
            )}
          </div>
          <div className="opacity-70 text-sm">Threshold: {multiSigAccount.threshold}</div>
        </div>
      </div>
      <div className="flex items-center justify-end">
        <a className="text-primary flex items-center" href="">
          View on Explorer
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="h-4 w-4 ml-2"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            strokeWidth={2}
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
            />
          </svg>
        </a>
      </div>
    </div>
  );
};
