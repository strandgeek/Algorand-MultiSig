import { KeyIcon } from "@heroicons/react/outline";
import algosdk, { Transaction } from "algosdk";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { useMultiSigAccountQuery, useTransactionQuery } from "../../client/queries";
import { AddressInfoLabel } from "../../components/AddressInfoLabel";
import { AlgoAmountLabel } from "../../components/AlgoAmountLabel";
import { InfoList, InfoListItem } from "../../components/InfoList";
import { SignaturesList } from "../../components/SignaturesList";
import { StatusLabel } from "../../components/StatusLabel";
import { useSignTransaction } from "../../hooks/useSignTransaction";
import { AppLayout } from "../../layouts/AppLayout";
import { SignedTransaction } from "../../types/signedTransaction";
import { getEncodedAddress } from "../../utils/getEncodedAddress";

interface ViewTransactionProps {}

const ViewTransaction: React.FC<ViewTransactionProps> = () => {
  const params = useParams()
  const { data: txData } = useTransactionQuery(params.txId)
  const { data: multiSigAccount } = useMultiSigAccountQuery(params.msaAddress)
  const [transaction, setTransaction] = useState<Transaction>()
  const signTransaction = useSignTransaction({
    multiSigAccount,
    transaction: txData,
  })
  useEffect(() => {
    if (txData) {
      const transaction = algosdk.decodeUnsignedTransaction(Buffer.from(txData.raw_transaction, "base64"))
      console.log(transaction)
      setTransaction(transaction)
    }
  }, [txData])

  const txOverviewItems: InfoListItem[] = [
    {
      label: "TxID",
      value: params.txId,
    },
    {
      label: "Status",
      value: <StatusLabel status="PENDING" />,
    },
    {
      label: "Sender",
      value: (
        <AddressInfoLabel address={getEncodedAddress(transaction?.from.publicKey)} />
      ),
    },
    {
      label: "Receiver",
      value: (
        <AddressInfoLabel address={getEncodedAddress(transaction?.to.publicKey)} />
      ),
    },
    {
      label: "Amount",
      value: (
        <AlgoAmountLabel value={transaction?.amount || 0} />
      ),
    },
  ];
  const signedTransactions: SignedTransaction[] = [
    {
      id: 1,
      account: {
        id: 1,
        address: 'UXVPARFR5J7BI5RXVZPCO5OE4OWNXEOZ6ZCDO5VFSBNH2IVZAQVQ',
      },
      transaction: {
        id: 1,
        txn_id: '1234',
        raw_transaction: '',
      },
    },
    {
      id: 2,
      account: {
        id: 2,
        address: 'AAVPARFR5J7BI5RXVZPCO5OE4OWNXEOZ6ZCDO5VFSBNH2IVZAQVQ',
      },
      transaction: {
        id: 1,
        txn_id: '1234',
        raw_transaction: '',
      },
    }
  ]
  return (
    <AppLayout>
      <div className="mx-auto max-w-4xl mt-8">
        <div className="font-bold text-xl mb-4">Transaction Overview</div>
        <div className="card bg-base-100 mb-8">
          <InfoList items={txOverviewItems} />
        </div>
        <div className="flex items-center justify-between">
          <div className="font-bold text-xl mb-4">Signatures</div>
          <div className="flex items-center">
            <div className="text-sm mr-2">2 of 3</div>
            <progress
              className="progress progress-primary w-32"
              value="66"
              max="100"
            ></progress>
          </div>
        </div>
        <div className="card bg-base-100 p-2 px-4 mb-8">
          <SignaturesList  signedTransactions={signedTransactions} />
        </div>

        <button className="btn btn-lg btn-primary mt-4 btn-block" onClick={() => signTransaction()}>
          <KeyIcon className="h-6 w-6 mr-1" />
          Sign Transaction
        </button>
      </div>
    </AppLayout>
  );
};

export default ViewTransaction;
