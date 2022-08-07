import { useQuery } from "@tanstack/react-query";
import { client } from "..";
import { Account } from "../../types/account";
import { MultiSigAccount } from "../../types/multisigAccount";

export const useMeQuery = () => useQuery<Account>(['me'], async () => {
  const { data } = await client.get('/auth/me')
  return data
})

export const useMultiSigAccountsQuery = () => useQuery<MultiSigAccount[]>(['multisig-accounts'], async () => {
  const { data } = await client.get('/multisig-accounts')
  return data
})

export const useMultiSigAccountQuery = (address?: string) => useQuery<MultiSigAccount>(['multisig-accounts', address], async () => {
  const { data } = await client.get(`/multisig-accounts/${address}`)
  return data
}, {
  enabled: !!address,
})

