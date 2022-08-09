import React, { FC } from "react";

export interface AlgoAmountLabelProps {
  value: number | bigint;
}

const formatter = new Intl.NumberFormat("en-US", {
  style: "decimal",
  maximumFractionDigits: 4,
  minimumFractionDigits: 4,
});

export const AlgoAmountLabel: FC<AlgoAmountLabelProps> = ({ value }) => {

  const algo = Number(value) * 10**-6
  return (
    <div className="flex items-center">
      <div className="mr-2">
        <img src={"/algo.png"} width={14} height={14} alt="Algorand Icon" />
      </div>
      {formatter.format(algo)}
    </div>
  );
};
