import algosdk from "algosdk";
import React, { FC } from "react";
import { FormProvider, useForm } from "react-hook-form";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import initiateAlgodClient from "../utils/algodClient";
import { FormControl } from "./FormControl";
import { SignerInput } from "./SignerInput";
import { KeyIcon } from "@heroicons/react/outline";
import classNames from "classnames";
import { MultiSigAccount } from "../types/multisigAccount";
import { encode as encodeBytesToBase64 } from 'base64-arraybuffer'

export interface CreateTransferTransactionFormProps {
  multiSigAccount: MultiSigAccount
  onRawTransactionCreated: (rawTxnBase64: string) => void
}

interface CreateTransferTransactionFormData {
  to: string;
  amount: number;
}

const schema = yup
  .object({
    to: yup.string().length(58).required(),
    amount: yup.number().required(),
  })
  .required();

export const CreateTransferTransactionForm: FC<
  CreateTransferTransactionFormProps
> = ({ multiSigAccount, onRawTransactionCreated }) => {
  const methods = useForm<CreateTransferTransactionFormData>({
    resolver: yupResolver(schema),
  });

  const onSubmit = async (data: CreateTransferTransactionFormData) => {
    const c = await initiateAlgodClient();
    const params = await c.getTransactionParams().do();
    const txn = algosdk.makePaymentTxnWithSuggestedParams(
      multiSigAccount.address,
      data.to,
      data.amount,
      undefined,
      undefined,
      params
    );
    const rawTxnBase64 = encodeBytesToBase64(txn.toByte())
    onRawTransactionCreated(rawTxnBase64)
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)}>
        <FormControl label="To Address">
          <input
            type="text"
            placeholder=""
            className={classNames("input input-bordered w-full", {
              "input-error": methods.formState.errors.amount,
            })}
            {...methods.register("to")}
          />
        </FormControl>
        <div className="mt-4">
          <FormControl label="Amount">
            <input
              type="number"
              placeholder=""
              className={classNames("input input-bordered w-full", {
                "input-error": methods.formState.errors.amount,
              })}
              {...methods.register("amount")}
            />
          </FormControl>
        </div>
        <button type="submit" className="btn btn-primary btn-block mt-8">
          Create Transaction
        </button>
      </form>
    </FormProvider>
  );
};