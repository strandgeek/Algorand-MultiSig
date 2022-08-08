import {
  CodeIcon,
  CollectionIcon,
  PaperAirplaneIcon,
} from "@heroicons/react/outline";
import classNames from "classnames";
import React, { FC } from "react";

export type TransactionType = 'TRANSFER' | 'ASSET' | 'RAW'

export interface TransactionTypeOptionItemProps {
  icon: React.ReactNode;
  title: string;
  active?: boolean;
  onClick: () => void;
}

export const TransactionTypeOptionItem: FC<TransactionTypeOptionItemProps> = ({
  active,
  icon,
  title,
  onClick,
}) => {
  const groupClassName = classNames(
    "group w-full p-8 inline-block border-2 border-base-content rounded-md",
    {
      "hover:border-primary hover:opacity-30 opacity-20": !active,
      "border-primary": active,
    }
  );
  const iconClassName = classNames("h-8 w-8 group-hover:text-primary", {
    "text-primary": active,
  });
  const titleClassName = classNames(
    "w-full text-center mt-2 font-bold text-sm",
    {
      "group-hover:text-primary": !active,
      "text-primary": active,
    }
  );
  return (
    <button type="button" className="my-4 w-full" onClick={onClick}>
      <div>
        <div className={groupClassName}>
          <div className="flex justify-center">
            <div className={iconClassName}>{icon}</div>
          </div>
          <div className={titleClassName}>{title}</div>
        </div>
      </div>
    </button>
  );
};

export interface TransactionTypeOptionsProps {
  type: TransactionType,
  setType: React.Dispatch<React.SetStateAction<TransactionType>>
}

export const TransactionTypeOptions: FC<
  TransactionTypeOptionsProps
> = ({ type, setType }) => {
  return (
    <div className="flex space-x-2">
      <TransactionTypeOptionItem
        icon={<PaperAirplaneIcon className="rotate-90" />}
        title="Transfer"
        active={type === 'TRANSFER'}
        onClick={() => setType('TRANSFER')}
      />
      <TransactionTypeOptionItem
        icon={<CollectionIcon />}
        title="Create Asset"
        active={type === 'ASSET'}
        onClick={() => setType('ASSET')}
      />
      <TransactionTypeOptionItem
        icon={<CodeIcon />}
        title="Raw"
        active={type === 'RAW'}
        onClick={() => setType('RAW')}
      />
    </div>
  );
};
