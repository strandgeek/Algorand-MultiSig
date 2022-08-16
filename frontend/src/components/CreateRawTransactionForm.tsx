import algosdk from "algosdk";
import { FC } from "react";
import { FormProvider, useForm } from "react-hook-form";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import initiateAlgodClient from "../utils/algodClient";
import { FormControl } from "./FormControl";
import classNames from "classnames";
import { MultiSigAccount } from "../types/multisigAccount";
import { encode as encodeBytesToBase64 } from 'base64-arraybuffer'

export interface CreateRawTransactionFormProps {
  multiSigAccount: MultiSigAccount
  onRawTransactionCreated: (rawTxnBase64: string) => void
}

interface CreateRawTransactionFormData {
  rawTxnBase64: string;
}

const schema = yup
  .object({
    rawTxnBase64: yup.string().required(),
  })
  .required();

export const CreateRawTransactionForm: FC<
  CreateRawTransactionFormProps
> = ({ onRawTransactionCreated }) => {
  const methods = useForm<CreateRawTransactionFormData>({
    resolver: yupResolver(schema),
  });

  const onSubmit = async (data: CreateRawTransactionFormData) => {
    onRawTransactionCreated(data.rawTxnBase64)
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)}>
        <FormControl label="Raw Transaction" info="Paste the raw transaction in base64">
          <textarea
            placeholder=""
            className={classNames("input input-bordered w-full min-h-16", {
              "input-error": methods.formState.errors.rawTxnBase64,
            })}
            {...methods.register("rawTxnBase64")}
          />
        </FormControl>
        <button type="submit" className="btn btn-primary btn-block mt-8">
          Create Transaction
        </button>
      </form>
    </FormProvider>
  );
};
