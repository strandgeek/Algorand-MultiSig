import React, { FC } from "react";
import { FormControl } from "../../components/FormControl";
import { AddressArrayInput } from "../../components/AddressArrayInput";
import { AppLayout } from "../../layouts/AppLayout";
import { FormProvider, useForm } from "react-hook-form";
import { useMeQuery } from "../../client/queries";
import { SignerInput } from "../../components/SignerInput";
import { KeyIcon } from "@heroicons/react/outline";
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from "yup";
import classNames from "classnames";
import { useCreateMultisigAccountMutation } from "../../client/mutations";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";

export interface CreateMultiSigAccountProps {}

interface CreateMultiSigFormData {
  addresses: {
    value: string;
  }[],
  threshold: number
}

const schema = yup.object({
  addresses: yup.array().of(yup.object().shape({
    value: yup.string().length(58).required(),
  })).required().min(1),
  threshold: yup.number().required(),
}).required();


export const CreateMultiSigAccount: FC<CreateMultiSigAccountProps> = props => {
  const navigate = useNavigate()
  const { data: me } = useMeQuery();
  const mutation = useCreateMultisigAccountMutation()
  const methods = useForm<CreateMultiSigFormData>({
    resolver: yupResolver(schema),
  });
  const onSubmit = async (data: CreateMultiSigFormData) => {
    try {      
      const multisigAccount = await mutation.mutateAsync({
        version: 1,
        addresses: [
          me!.address,
          ...data.addresses.map(addr => addr.value),
        ],
        threshold: data.threshold,
      })
      toast.success('MultiSig Account Created!');
      console.log(multisigAccount)
      navigate(`/app/multisig-accounts/${multisigAccount?.address}`)
    } catch (error) {
      toast.error('Failed to create MultiSig Account.');
    }
  }
  return (
    <AppLayout>
      <div className="flex justify-center mt-12">
        <div className="max-w-xl w-full">
          <div className="card w-full bg-base-100 shadow-xl">
            <div className="card-body">
              <h2 className="card-title">Create MultiSig Account</h2>
              <FormProvider {...methods}>
                <form onSubmit={methods.handleSubmit(onSubmit)}>
                  <FormControl label="Accounts / Signers">
                    <SignerInput
                      value={me?.address}
                      disabled
                      leftComp={
                        <span className="text-gray-500 sm:text-sm">
                          <KeyIcon className="w-6 h-6" />
                        </span>
                      }
                      rightComp={
                        <div
                          className="px-2 mr-2 text-xs badge badge-primary flex items-center"
                        >
                          You
                        </div>
                      }
                    />
                    <AddressArrayInput />
                  </FormControl>
                  <FormControl
                    label="Threshold"
                    info="How many signatures is necessary to make a transaction?"
                  >
                    <input
                      type="number"
                      placeholder=""
                      className={
                        classNames(
                          'input input-bordered w-full',
                          {
                            'input-error': methods.formState.errors.threshold,
                          }
                        )
                      }
                      {...methods.register('threshold')}
                    />
                  </FormControl>
                  <button type="submit" className="btn btn-primary btn-block mt-4">
                    Create MultiSig Account
                  </button>
                </form>
              </FormProvider>
            </div>
          </div>
        </div>
      </div>
    </AppLayout>
  );
};
