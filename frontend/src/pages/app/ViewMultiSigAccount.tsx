import React, { FC } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { useMultiSigAccountQuery } from '../../client/queries'
import { MultiSigAccountDetails } from '../../components/MultiSigAccountDetails'
import { MultiSigAccountSignersList } from '../../components/MultiSigAccountSignersList'
import { AppLayout } from '../../layouts/AppLayout'
import { getIdenticonSrc } from '../../utils/getIdenticonSrc'
import { getShortAddress } from '../../utils/getShortAddress'

export interface ViewMultiSigAccountProps {
  
}

export const ViewMultiSigAccount: FC<ViewMultiSigAccountProps> = (props) => {
  const { msaAddress } = useParams()
  const { data: multiSigAccount } = useMultiSigAccountQuery(msaAddress)
  if (!multiSigAccount) {
    return null
  }
  return (
    <AppLayout>
      <div className="mx-auto max-w-4xl mt-8">
        <div className="font-bold text-xl mb-4">MultiSig Account Details</div>
        <div className="card bg-base-100 mb-8">
          <MultiSigAccountDetails  multiSigAccount={multiSigAccount} />
        </div>

        <div className="font-bold text-xl mb-4">Signers</div>
        <div className="card bg-base-100 p-2 px-4 mb-8">
          <MultiSigAccountSignersList multiSigAccount={multiSigAccount} />
        </div>
      </div>
    </AppLayout>
  )
}
