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

export interface CreateAssetTransactionFormProps {
  multiSigAccount: MultiSigAccount
  onRawTransactionCreated: (rawTxnBase64: string) => void
}

interface CreateAssetTransactionFormData {
  unitName: string;
  assetName: string;
  assetURL: string;
  supply: number;
}

const schema = yup
  .object({
    unitName: yup.string().min(2),
    assetName: yup.string().min(2),
    assetURL: yup.string().min(2),
    supply: yup.number().min(1).required(),
  })
  .required();

export const CreateAssetTransactionForm: FC<
  CreateAssetTransactionFormProps
> = ({ multiSigAccount, onRawTransactionCreated }) => {
  const methods = useForm<CreateAssetTransactionFormData>({
    resolver: yupResolver(schema),
  });

  const onSubmit = async (data: CreateAssetTransactionFormData) => {
    const c = await initiateAlgodClient();
    const params = await c.getTransactionParams().do();
    const creator = multiSigAccount.address;
    const defaultFrozen = false;
    let note = undefined;
    const manager = creator;
    const reserve = creator;
    const freeze = creator;
    const clawback = creator;
    let assetMetadataHash = undefined;
    const decimals = 0;
    const txn = algosdk.makeAssetCreateTxnWithSuggestedParams(
        creator,
        note,
        data.supply,
        decimals,
        defaultFrozen,
        manager,
        reserve,
        freeze,
        clawback,
        data.unitName,
        data.assetName,
        data.assetURL,
        assetMetadataHash,
        params,
    );
    const rawTxnBase64 = encodeBytesToBase64(txn.toByte())
    onRawTransactionCreated(rawTxnBase64)
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)}>
        <FormControl label="Name">
          <input
            type="text"
            placeholder=""
            className={classNames("input input-bordered w-full mb-2", {
              "input-error": methods.formState.errors.assetName,
            })}
            {...methods.register("assetName")}
          />
        </FormControl>
        <FormControl label="Unit Name">
          <input
            type="text"
            placeholder=""
            className={classNames("input input-bordered w-full mb-2", {
              "input-error": methods.formState.errors.unitName,
            })}
            {...methods.register("unitName")}
          />
        </FormControl>
        <FormControl label="Asset URL">
          <input
            type="text"
            placeholder=""
            className={classNames("input input-bordered w-full mb-2", {
              "input-error": methods.formState.errors.assetURL,
            })}
            {...methods.register("assetURL")}
          />
        </FormControl>
        <FormControl label="Supply">
          <input
            type="number"
            placeholder=""
            className={classNames("input input-bordered w-full mb-2", {
              "input-error": methods.formState.errors.supply,
            })}
            {...methods.register("supply")}
          />
        </FormControl>
        <button type="submit" className="btn btn-primary btn-block mt-8">
          Create Transaction
        </button>
      </form>
    </FormProvider>
  );
};
