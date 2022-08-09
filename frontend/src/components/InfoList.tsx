import React, { FC } from "react";

export interface InfoListItem {
  label: string;
  value: React.ReactNode;
}

export interface InfoListProps {
  items: InfoListItem[];
}

export const InfoList: FC<InfoListProps> = ({ items }) => {
  return (
    <div>
      {items.map(i => (
        <div className="p-4 border-b border-b-base-300 flex items-center text-sm">
          <span className="text-base-content font-bold text-sm mr-4 w-24 inline-block">
            {i.label}:
          </span>
          {i.value}
        </div>
      ))}
    </div>
  );
};
