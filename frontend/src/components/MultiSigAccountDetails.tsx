import { ExternalLinkIcon } from "@heroicons/react/outline";
import React, { FC } from "react";
import { MultiSigAccount } from "../types/multisigAccount";
import { getAddressUrl } from "../utils/explorer";
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
        <a className="text-primary flex items-center" href={getAddressUrl(multiSigAccount.address)} target="_blank" rel="noreferrer">
          View on Explorer
          <ExternalLinkIcon className="h-4 w-4 ml-2" />
        </a>
      </div>
    </div>
  );
};
