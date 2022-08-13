import React, { FC, useState } from 'react'
import { toast } from 'react-toastify';
import { SelectAccountModal } from '../components/SelectAccountModal';
import { useAuth } from '../hooks/useAuth';
// import { useAuth } from '../hooks/useAuth';

export interface IndexProps {}

export const Index: FC<IndexProps> = () => {
  const [selectAccountModalOpen, setSelectAccountModalOpen] = useState(false)
  const [addresses, setAddresses] = useState<string[]>([])
  const auth = useAuth()
  const connectAccount = async () => {
    if (typeof window.AlgoSigner === 'undefined') {
      toast.error('AlgoSigner not installed')
    }
    const { AlgoSigner } = window
    await AlgoSigner.connect()
    const accounts = await AlgoSigner.accounts({
      ledger: 'TestNet'
    })
    setAddresses(accounts.map((acc: any) => acc.address))
    setSelectAccountModalOpen(true)
  }
  const onSelectAuthAccount = (address: string) => {
    auth(address)
  }
  return (
    <div>
      <SelectAccountModal
        addresses={addresses}
        open={selectAccountModalOpen}
        setOpen={() => setSelectAccountModalOpen(false)}
        onSelect={onSelectAuthAccount}
      />
      <button onClick={() => connectAccount()}>
        Authenticate
      </button>
    </div>
  )
}
