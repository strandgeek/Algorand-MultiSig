import { FC } from "react";
import { AddressInfoLabel } from "./AddressInfoLabel";
import { Modal, ModalProps } from "./Modal";

export interface SelectAccountModalProps extends Pick<ModalProps, 'open' | 'setOpen'> {
  addresses: string[];
  onSelect: (address: string) => void;
}

export const SelectAccountModal: FC<SelectAccountModalProps> = ({ addresses, onSelect, open, setOpen }) => {
  return (
    <Modal open={open} setOpen={setOpen}>
      <div className="p-4">
        <div className="font-bold p-2 pb-4">
          Select the account to log in:
        </div>
        <ul className="menu menu-compact lg:menu-normal bg-base-100 w-full p-2 rounded-box">
          {addresses.map((addr: string) => (
            <li>
              <button onClick={() => onSelect(addr)}>
                <AddressInfoLabel
                  address={addr}
                  short
                />
              </button>
            </li>
          ))}
        </ul>
      </div>
    </Modal>
  );
};
