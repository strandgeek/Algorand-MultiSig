import React, { FC } from "react";
import { getIdenticonSrc } from "../utils/getIdenticonSrc";

export interface AddressAvatarsStackProps {
  addresses: string[];
}

export const AddressAvatarsStack: FC<AddressAvatarsStackProps> = ({
  addresses,
}) => {
  return (
    <div className="avatar-group -space-x-5">
      {addresses.map((addr) => (
        <div className="avatar" key={addr}>
          <div className="w-8">
            <img
              src={getIdenticonSrc(addr)}
              width={160}
              height={160}
              alt={`Avatar Wallet ${addr}`}
            />
          </div>
        </div>
      ))}
    </div>
  );
};
