import React, { FC } from "react";

export interface FormControlProps {
  label?: string
  info?: string
  children: React.ReactNode
}

export const FormControl: FC<FormControlProps> = ({ label, info, children }) => {
  return (
    <div className="form-control w-full">
      {label && (
        <label className="label font-bold opacity-75">
          <span className="label-text">{label}</span>
        </label>
      )}
      {children}
      {info && (
        <label className="label">
          <span className="label-text-alt">{info}</span>
        </label>
      )}
    </div>
  );
};
