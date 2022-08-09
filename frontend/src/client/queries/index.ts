import { useQuery } from "@tanstack/react-query";
import { client } from "..";
import { Account } from "../../types/account";
import { MultiSigAccount } from "../../types/multisigAccount";
import { Transaction } from "../../types/transaction";

export const useMeQuery = () => useQuery<Account>(['me'], async () => {
  const { data } = await client.get('/auth/me')
  return data
})

export const useMultiSigAccountsQuery = () => useQuery<MultiSigAccount[]>(['multisig-accounts'], async () => {
  const { data } = await client.get('/multisig-accounts')
  return data
})

export const useMultiSigAccountQuery = (address?: string) => useQuery<MultiSigAccount>(['multisig-account', address], async () => {
  const { data } = await client.get(`/multisig-accounts/${address}`)
  return data
}, {
  enabled: !!address,
})

export const useMultiSigAccountTransactionsQuery = (
  msAddress?: string
) => useQuery<Transaction[]>(['multisig-account-transactions', msAddress], async () => {
  const { data } = await client.get(`/multisig-accounts/${msAddress}/transactions`)
  return data
}, {
  enabled: !!msAddress,
})

export const useTransactionQuery = (
  txId?: string
) => useQuery<Transaction>(['transaction', txId], async () => {
  const { data } = await client.get(`/transactions/${txId}`)
  return data
}, {
  enabled: !!txId,
  refetchInterval: 1000,
})

