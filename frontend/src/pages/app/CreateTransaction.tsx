import { FC, useState } from "react";
import { AppLayout } from "../../layouts/AppLayout";
import { useMultiSigAccountQuery } from "../../client/queries";
import {
  TransactionType,
  TransactionTypeOptions,
} from "../../components/TransactionTypeOptions";
import { CreateTransferTransactionForm } from "../../components/CreateTransferTransactionForm";
import { Link, useNavigate, useParams } from "react-router-dom";
import { useCreateTransactionMutation } from "../../client/mutations";
import { CreateAssetTransactionForm } from "../../components/CreateAssetTransactionForm";
import { CreateRawTransactionForm } from "../../components/CreateRawTransactionForm";
import { toast } from "react-toastify";
import { getShortAddress } from "../../utils/getShortAddress";

export interface CreateTransactionProps {}

export const CreateTransaction: FC<CreateTransactionProps> = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [type, setType] = useState<TransactionType>("TRANSFER");
  const mutation = useCreateTransactionMutation();
  const { data: multiSigAccount } = useMultiSigAccountQuery(params.msaAddress);
  if (!multiSigAccount) {
    return null;
  }
  const onRawTransactionCreated = async (rawTxnBase64: string) => {
    try {
      const res = await mutation.mutateAsync({
        multisig_account_address: multiSigAccount.address,
        raw_transaction_base_64: rawTxnBase64,
      });
      navigate(
        `/app/multisig-accounts/${multiSigAccount.address}/transactions/${res.txn_id}`
      );
    } catch (error) {
      toast.error("Could not create transaction. Please check the fields");
    }
  };
  return (
    <AppLayout>
      <div className="flex justify-center mt-12">
        <div className="max-w-xl w-full">
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
              <li>Create Transaction</li>
            </ul>
          </div>
          <div className="card w-full bg-base-100 shadow-xl">
            <div className="card-body">
              <h2 className="card-title">Create Transaction</h2>
              <label className="mt-4">Transaction Type:</label>
              <TransactionTypeOptions type={type} setType={setType} />
              {type === "TRANSFER" && (
                <CreateTransferTransactionForm
                  multiSigAccount={multiSigAccount}
                  onRawTransactionCreated={onRawTransactionCreated}
                />
              )}
              {type === "ASSET" && (
                <CreateAssetTransactionForm
                  multiSigAccount={multiSigAccount}
                  onRawTransactionCreated={onRawTransactionCreated}
                />
              )}
              {type === "RAW" && (
                <CreateRawTransactionForm
                  multiSigAccount={multiSigAccount}
                  onRawTransactionCreated={onRawTransactionCreated}
                />
              )}
            </div>
          </div>
        </div>
      </div>
    </AppLayout>
  );
};
