import React, { FC } from "react";
import { SignedTransaction } from "../types/signedTransaction";
import { getIdenticonSrc } from "../utils/getIdenticonSrc";
import { getShortAddress } from "../utils/getShortAddress";

export interface SignaturesListProps {
  signedTransactions: SignedTransaction[];
}

export const SignaturesList: FC<SignaturesListProps> = ({
  signedTransactions,
}) => {
  return (
    <div>
      {signedTransactions.map(signedTransaction => {
        const address = signedTransaction.signer.address
        return (
          <div
            key={signedTransaction.id}
            className="grid grid-cols-2 border-b border-b-base-300 last:border-none py-4"
          >
            <div className="flex items-center">
              <div className="avatar">
                <div className="rounded-full w-6 h-6">
                  <img
                    src={getIdenticonSrc(address)}
                    width={160}
                    height={160}
                    alt={`Address`}
                  />
                </div>
              </div>
              <div className="ml-4 text-sm">
                {getShortAddress(address)}
              </div>
            </div>
            <div className="flex items-center justify-end">
              <div className="text-green-500 flex items-center text-sm">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-4 w-4 mr-1"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                  strokeWidth={2}
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z"
                  />
                </svg>
                Signed
              </div>
            </div>
          </div>
        );
      })}
    </div>
  );
};
