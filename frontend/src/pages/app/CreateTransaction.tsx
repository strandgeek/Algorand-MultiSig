import { FC, useState } from "react";
import { AppLayout } from "../../layouts/AppLayout";
import { useMultiSigAccountQuery } from "../../client/queries";
import { TransactionType, TransactionTypeOptions } from "../../components/TransactionTypeOptions";
import { CreateTransferTransactionForm } from "../../components/CreateTransferTransactionForm";
import { useNavigate, useParams } from "react-router-dom";
import { useCreateTransactionMutation } from "../../client/mutations";

export interface CreateTransactionProps {}

export const CreateTransaction: FC<CreateTransactionProps> = () => {
  const params = useParams()
  const navigate = useNavigate()
  const [type, setType] = useState<TransactionType>('TRANSFER')
  const mutation = useCreateTransactionMutation()
  const { data: multiSigAccount } = useMultiSigAccountQuery(params.msaAddress)
  if (!multiSigAccount) {
    return null
  }
  const onRawTransactionCreated = async (rawTxnBase64: string) => {
    const res = await mutation.mutateAsync({
      multisig_account_address: multiSigAccount.address,
      raw_transaction_base_64: rawTxnBase64,
    })
    navigate(`/app/multisig-accounts/${multiSigAccount.address}/transactions/${res.txn_id}`);
  }
  return (
    <AppLayout>
      <div className="flex justify-center mt-12">
        <div className="max-w-xl w-full">
          <div className="card w-full bg-base-100 shadow-xl">
            <div className="card-body">
              <h2 className="card-title">Create Transaction</h2>
              <label className="mt-4">Transaction Type:</label>
              <TransactionTypeOptions type={type} setType={setType} />
              {type === 'TRANSFER' && (
                <CreateTransferTransactionForm
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
