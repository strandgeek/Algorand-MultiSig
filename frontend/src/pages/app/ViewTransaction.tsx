import { KeyIcon } from "@heroicons/react/outline";
import { AddressInfoLabel } from "../../components/AddressInfoLabel";
import { AlgoAmountLabel } from "../../components/AlgoAmountLabel";
import { InfoList, InfoListItem } from "../../components/InfoList";
import { SignaturesList } from "../../components/SignaturesList";
import { StatusLabel } from "../../components/StatusLabel";
import { AppLayout } from "../../layouts/AppLayout";
import { SignedTransaction } from "../../types/signedTransaction";

interface ViewTransactionProps {}

const ViewTransaction: React.FC<ViewTransactionProps> = () => {
  const txOverviewItems: InfoListItem[] = [
    {
      label: "TxID",
      value: "1234",
    },
    {
      label: "Status",
      value: <StatusLabel status="BROADCASTED" />,
    },
    {
      label: "Sender",
      value: (
        <AddressInfoLabel address="UXVPARFR5J7BI5RXVZPCO5OE4OWNXEOZ6ZCDO5VFSBNH2IVZAQVQ" />
      ),
    },
    {
      label: "Receiver",
      value: (
        <AddressInfoLabel address="UXVPARFR5J7BI5RXVZPCO5OE4OWNXEOZ6ZCDO5VFSBNH2IVZAQVQ" />
      ),
    },
    {
      label: "Amount",
      value: (
        <AlgoAmountLabel value={5000000} />
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

        <button className="btn btn-lg btn-primary mt-4 btn-block">
          <KeyIcon className="h-6 w-6 mr-1" />
          Sign Transaction
        </button>
      </div>
    </AppLayout>
  );
};

export default ViewTransaction;
