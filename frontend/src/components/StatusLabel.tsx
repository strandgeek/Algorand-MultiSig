import classNames from "classnames";
import React, { FC } from "react";
import ClipLoader from "react-spinners/ClipLoader";
import { Transaction } from "../types/transaction";

export interface StatusLabelProps {
  status: Transaction['status'];
}

const STATUS_LABELS: {
  [key: string]: string;
} = {
  PENDING: "Pending Signatures",
  READY: "Ready",
  FAILED: "Failed",
  BROADCASTING: "Broadcasting...",
  BROADCASTED: "Broadcasted",
};

const STATUS_COLOR_CLASSNAME: {
  [key: string]: string;
} = {
  PENDING: "bg-yellow-500",
  READY: "bg-blue-500",
  BROADCASTING: "bg-blue-700",
  FAILED: "bg-red-500",
  BROADCASTED: "bg-green-500",
};

export const StatusLabel: FC<StatusLabelProps> = ({ status }) => {
  const dotClassName = classNames(
    "flex-shrink-0 w-2.5 h-2.5 rounded-full",
    STATUS_COLOR_CLASSNAME[status]
  );
  return (
    <div className="flex items-center space-x-2">
      {status === "BROADCASTING" ? (
        <ClipLoader size={14} className="ml-2" color="#999999" />
      ) : (
        <div className={dotClassName} aria-hidden="true" />
      )}
      <span>
        <span className="text-gray-500 font-normal">
          {STATUS_LABELS[status]}
        </span>
      </span>
    </div>
  );
};
