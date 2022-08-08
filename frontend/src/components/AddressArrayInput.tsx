import React, { FC } from "react";
import { useFormContext, useFieldArray } from "react-hook-form";
import { KeyIcon, PlusCircleIcon, XCircleIcon } from "@heroicons/react/outline";
import { SignerInput } from "./SignerInput";

export interface AddressArrayInputProps {}

export const AddressArrayInput: FC<AddressArrayInputProps> = props => {
  const { control, register, formState } = useFormContext();
  const { fields, append, prepend, remove, swap, move, insert } = useFieldArray(
    {
      control,
      name: "addresses",
    }
  );
  return (
    <div>
      {fields.map((field, index) => {
        const isValid = !(formState?.errors?.addresses && formState?.errors?.addresses[index])
        return (
          <div className="mt-2">
              <SignerInput
                key={field.id}
                error={!isValid}
                {...register(`addresses.${index}.value`)} 
                leftComp={
                  <span className="text-gray-500 sm:text-sm">
                    <KeyIcon className="w-6 h-6" color={isValid ? '#999999' : '#c0392b'} />
                  </span>
                }
                rightComp={
                  <button type="button" className="btn btn-link" onClick={() => remove(index)}>
                    <XCircleIcon className="w-6 h-6" color="#999999" />
                  </button>
                }
              />
          </div>
        )
      })}
      <button
        type="button"
        className="btn btn-link btn-primary mt-2"
        onClick={() => append("")}
      >
        <PlusCircleIcon className="w-4 h-4 mr-1" />
        Add Signer
      </button>
      {formState?.errors?.addresses && (
        <div className="mt-1 mb-4 text-error">
          Please enter valid Algorand addresses
        </div>
      )}
    </div>
  );
};
