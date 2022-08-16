import {
  CheckCircleIcon,
  ExternalLinkIcon,
  KeyIcon,
} from "@heroicons/react/outline";
import algosdk, { Transaction } from "algosdk";
import { AxiosError } from "axios";
import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { toast } from "react-toastify";
import {
  useMeQuery,
  useMultiSigAccountQuery,
  useTransactionQuery,
} from "../../client/queries";
import { AddressInfoLabel } from "../../components/AddressInfoLabel";
import { AlgoAmountLabel } from "../../components/AlgoAmountLabel";
import { InfoList, InfoListItem } from "../../components/InfoList";
import { LoadingSpinner } from "../../components/LoadingSpinner";
import { SignaturesList } from "../../components/SignaturesList";
import { StatusLabel } from "../../components/StatusLabel";
import { useSignTransaction } from "../../hooks/useSignTransaction";
import { AppLayout } from "../../layouts/AppLayout";
import { getTransactionUrl } from "../../utils/explorer";
import { getEncodedAddress } from "../../utils/getEncodedAddress";
import { getShortAddress } from "../../utils/getShortAddress";

interface ViewTransactionProps {}


const TX_TYPE_NAME: { [type: string]: string } = {
  pay: 'Payment',
  keyreg: 'Key Registration',
  acfg: 'Asset Configuration',
  axfer: 'Asset Transfer',
  afrz: 'Asset Freeze',
  appl: 'Application Transaction',
}

const ViewTransaction: React.FC<ViewTransactionProps> = () => {
  const params = useParams();
  const { data: me } = useMeQuery();
  const { data: txData, refetch, isLoading } = useTransactionQuery(params.txId);
  const { data: multiSigAccount } = useMultiSigAccountQuery(params.msaAddress);
  const [transaction, setTransaction] = useState<Transaction>();
  const signTransaction = useSignTransaction({
    multiSigAccount,
    transaction: txData,
  });
  useEffect(() => {
    if (txData) {
      const transaction = algosdk.decodeUnsignedTransaction(
        Buffer.from(txData.raw_transaction, "base64")
      );
      setTransaction(transaction);
    }
  }, [txData]);

  if (isLoading) {
    return <LoadingSpinner />
  }

  const txOverviewItems: InfoListItem[] = [
    {
      label: "TxID",
      value: (
        <>
          {txData?.status === "BROADCASTED" ? (
            <a
              className="text-primary flex items-center"
              href={getTransactionUrl(txData.txn_id)}
              target="_blank"
              rel="noreferrer"
            >
              {params.txId}
              <ExternalLinkIcon className="h-4 w-4 ml-2" />
            </a>
          ) : (
            <span>{params.txId}</span>
          )}
        </>
      ),
    },
    {
      label: "Status",
      value: <StatusLabel status={txData!.status} />,
    },
    {
      label: "Type",
      value: <span className="badge badge-outline">{transaction?.type && TX_TYPE_NAME[transaction?.type]}</span>,
    },
  ];

  const txDetailsItems: InfoListItem[] = []

  if (transaction?.type === 'pay') {
    txDetailsItems.push(
      {
        label: "Sender",
        value: (
          <Link to={`/app/multisig-accounts/${multiSigAccount?.address}`}>
            <AddressInfoLabel
              address={getEncodedAddress(transaction?.from?.publicKey)}
            />
          </Link>
        ),
      },
      {
        label: "Receiver",
        value: (
          <AddressInfoLabel
            address={getEncodedAddress(transaction?.to?.publicKey)}
          />
        ),
      },
      {
        label: "Amount",
        value: <AlgoAmountLabel value={transaction?.amount || 0} />,
      },
    )
  }

  if (transaction?.type === 'acfg') {
    txDetailsItems.push(
      {
        label: "Asset Name",
        value: <span>{transaction.assetName}</span>
      },
      {
        label: "Asset Unit Name",
        value: <span>{transaction.assetUnitName}</span>
      },
      {
        label: "Asset URL",
        value: <span>{transaction.assetURL}</span>,
      },
      {
        label: "Total",
        value: <span>{transaction.assetTotal}</span>,
      },
    )
  }

  const onSignClick = async () => {
    try {
      await signTransaction();
      refetch();
      toast.success("Transaction signed!");
    } catch (e: any) {
      const error = e as AxiosError;
      if (error.response?.status === 409) {
        toast.error("You already signed this transaction");
      } else {
        toast.error("Could not sign transaction");
      }
      console.error(`Could not sign transaction: ${e.message}`);
    }
  };
  const signaturesCount = txData?.signed_transactions?.length || 0;
  const requiredSignaturesTotal = multiSigAccount?.threshold;
  const alreadySigned = !!txData?.signed_transactions?.find(
    st => st.signer.address === me?.address
  );
  return (
    <AppLayout>
      <div className="mx-auto max-w-4xl mt-8">
        <div className="text-sm breadcrumbs mb-4">
          <ul>
            <li>
              <Link to="/app/multisig-accounts">MultiSig Accounts</Link>
            </li>
            <li>
              <Link to={`/app/multisig-accounts/${multiSigAccount?.address}`}>
                MultiSig Account ({getShortAddress(multiSigAccount?.address)})
              </Link>
            </li>
            <li>
              Transaction ({getShortAddress(txData?.txn_id)})
            </li>
          </ul>
        </div>
        <div className="font-bold text-xl mb-4">Transaction Overview</div>
        <div className="card bg-base-100 mb-8">
          <InfoList items={txOverviewItems} />
        </div>
        {txDetailsItems.length > 0 && (
          <>
            <div className="font-bold text-xl mb-4">Transaction Details</div>
            <div className="card bg-base-100 mb-8">
              <InfoList items={txDetailsItems} />
            </div>
          </>
        )}
        <div className="flex items-center justify-between">
          <div className="font-bold text-xl mb-4">Signatures</div>
          <div className="flex items-center">
            <div className="text-sm mr-2">
              {signaturesCount} of {requiredSignaturesTotal}
            </div>
            <progress
              className="progress progress-primary w-32"
              value={signaturesCount}
              max={requiredSignaturesTotal}
            ></progress>
          </div>
        </div>
        <div className="card bg-base-100 p-2 px-4 mb-8 overflow-visible">
          {signaturesCount > 0 ? (
            <SignaturesList
              signedTransactions={txData?.signed_transactions || []}
            />
          ) : (
            <div className="text-center p-8">
              <KeyIcon className="mx-auto h-12 w-12 text-gray-400" />
              <h3 className="mt-2 text-sm font-medium text-gray-900">
                No signatures
              </h3>
              <p className="mt-4 text-sm text-gray-500">
                No signers have signed this transaction yet
              </p>
            </div>
          )}
        </div>

        {alreadySigned ? (
          <div className="text-center text-green-500 flex items-center justify-center">
            <CheckCircleIcon className="w-5 h-5 mr-1" />
            You signed this transaction
          </div>
        ) : (
          <button
            className="btn btn-lg btn-primary mt-4 btn-block"
            onClick={onSignClick}
          >
            <KeyIcon className="h-6 w-6 mr-1" />
            Sign Transaction
          </button>
        )}
      </div>
    </AppLayout>
  );
};

export default ViewTransaction;
