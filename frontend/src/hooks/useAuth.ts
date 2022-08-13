import algosdk from "algosdk";
import { useNavigate } from "react-router-dom"
import { client } from "../client";

const TESTNET_GENESIS_ID = 'testnet-v1.0';
const TESTNET_GENESIS_HASH = 'SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=';

type AuthFn = (address: string) => void

export const useAuth =  (): AuthFn => {
  const navigate = useNavigate()
  const auth = async (address: string) => {
    if (typeof window.AlgoSigner !== 'undefined') {
      const { AlgoSigner } = window
      const enc = new TextEncoder();
      const { data: { nonce } } = await client.post('/auth/nonce', {
        address,
      })
      const note = enc.encode(`Authentication. Nonce: ${nonce}`);
      const txn = algosdk.makePaymentTxnWithSuggestedParamsFromObject({
        from: address,
        to: address,
        amount: 0o00000,
        note: note,
        suggestedParams: {
            fee: 0o000,
            flatFee: true,
            firstRound: 10000,
            lastRound: 10200,
            genesisHash: TESTNET_GENESIS_HASH,
            genesisID: TESTNET_GENESIS_ID,
        },
    });
    const txnToSign = Buffer.from(algosdk.encodeUnsignedTransaction(txn)).toString("base64");
      const signResponse = await AlgoSigner.signTxn([
        {
          txn: txnToSign,
        },
      ]);
      const authData = {
        signed_tx_base64: signResponse[0].blob,
        pub_key: address,
      }
      const { data: { token } } = await client.post('/auth/complete', authData)
      localStorage.setItem('token', token)
      navigate('/app')
    } else {
      alert('AlgoSigner is not installed')
    };
  }
  return auth
}
