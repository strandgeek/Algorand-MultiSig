import { useQuery } from "@tanstack/react-query";
import { client } from "..";
import { Account } from "../../types/account";
import { MultiSigAccount } from "../../types/multisigAccount";

export const useMultiSigAccountsQuery = () => useQuery<MultiSigAccount[]>(['multisig-accounts'], async () => {
  const { data } = await client.get('/multisig-accounts')
  return data
})

export const useMeQuery = () => useQuery<Account>(['me'], async () => {
  const { data } = await client.get('/auth/me')
  return data
})

