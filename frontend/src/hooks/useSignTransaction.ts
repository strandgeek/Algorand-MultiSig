import { toast } from "react-toastify";
import { useMeQuery } from "../client/queries";
import { MultiSigAccount } from "../types/multisigAccount";
import { Transaction } from "../types/transaction";

type SignTransactionFn = () => void

interface UseSignTransactionOptions {
  multiSigAccount?: MultiSigAccount
  transaction?: Transaction
}

export const useSignTransaction = ({
    multiSigAccount,
    transaction,
  }: UseSignTransactionOptions): SignTransactionFn => {
    const { data: me } = useMeQuery()
    const sign = async () => {
      if (!me || !multiSigAccount || !transaction) {
        toast.error('Could not sign txn: Invalid MultiSig Account or Transaction')
        return null
      }
      const { AlgoSigner } = window
      await AlgoSigner.connect()
      const mparams = {
        version: multiSigAccount.version,
        threshold: multiSigAccount.threshold,
        addrs: multiSigAccount.accounts.map(acc => acc.address),
      }
      let signedTxs = await AlgoSigner.signTxn([
        {
          txn: transaction.raw_transaction,
          msig: mparams,
          signers: [me.address], // Use logged user as a signer
        },
      ]);
  
      const txID = signedTxs[0].txID;
      const signedTxn = signedTxs[0].blob;
      console.log({
        txID,
        signedTxn,
      })
      return null
    }
    return sign
}
