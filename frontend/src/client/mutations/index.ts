import { useMutation } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { client } from "..";
import { MultiSigAccount } from "../../types/multisigAccount";


interface CreateMutationInput {
  version: number;
  threshold: number;
  addresses: string[];
}

export const useCreateMultisigAccountMutation = () => useMutation<MultiSigAccount, AxiosError, CreateMutationInput>({
  mutationKey: ['create-multisig-account'],
  mutationFn: async (input) => {
    const { data } = await client.post('/multisig-accounts', input)
    return data
  }
})
