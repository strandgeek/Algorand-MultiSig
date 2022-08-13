import { useMutation } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { client } from "..";
import { MultiSigAccount } from "../../types/multisigAccount";
import { SignedTransaction } from "../../types/signedTransaction";
import { Transaction } from "../../types/transaction";


interface CreateMultisigAccountMutationInput {
  version: number;
  threshold: number;
  addresses: string[];
}

interface CreateTransactionMutationInput {
  multisig_account_address: string;
  raw_transaction_base_64: string;
}

interface CreateSignedTransactionMutationInput {
  transaction_txn_id: string;
  raw_signed_transaction_base_64: string;
}

export const useCreateMultisigAccountMutation = () => useMutation<MultiSigAccount, AxiosError, CreateMultisigAccountMutationInput>({
  mutationKey: ['create-multisig-account'],
  mutationFn: async (input) => {
    const { data } = await client.post('/multisig-accounts', input)
    return data
  }
})

export const useCreateTransactionMutation = () => useMutation<Transaction, AxiosError, CreateTransactionMutationInput>({
  mutationKey: ['create-transaction'],
  mutationFn: async (input) => {
    const { data } = await client.post('/transactions', input)
    return data
  }
})

export const useCreateSignedTransactionMutation = () => useMutation<SignedTransaction, AxiosError, CreateSignedTransactionMutationInput>({
  mutationKey: ['create-signed-transaction'],
  mutationFn: async (input) => {
    const { data } = await client.post('/signed-transactions', input)
    return data
  }
})
