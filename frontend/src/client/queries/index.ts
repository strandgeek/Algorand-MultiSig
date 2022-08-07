import { useQuery } from "@tanstack/react-query";
import { client } from "..";
import { MultiSigAccount } from "../../types/multisigAccount";

export const useMultiSigAccounts = () => useQuery<MultiSigAccount[]>(['multisig-accounts'], async () => {
  const { data } = await client.get('/multisig-accounts')
  return data
})
