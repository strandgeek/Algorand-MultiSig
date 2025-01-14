package broadcastsvc

import (
	"context"
	"fmt"
	"multisigdb-svc/client"
	"multisigdb-svc/model"

	b64 "encoding/base64"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var KnownTLSError = "Post \"https://testnet-algorand.api.purestake.io/ps2/v2/transactions\": net/http: TLS handshake timeout"

type BroadcastService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewBroadcastService(db *gorm.DB, logger *zap.Logger) *BroadcastService {
	return &BroadcastService{
		db:     db,
		logger: logger,
	}
}

func (s *BroadcastService) BroadcastAllSignedTxn() {
	var txns []model.Transaction
	if err := s.db.Where("status = ?", "READY").Find(&txns).Error; err != nil {
		s.logger.Error("Could not get ready transactions")
		return
	}
	if len(txns) == 0 {
		s.logger.Info("No Transaction with status ready found")
		return
	}
	s.logger.Info(fmt.Sprintf("Total %v transactions found with status ready found now broadcasting it to network", len(txns)))

	for _, txn := range txns {
		s.BroadcastTxn(&txn)
	}
}

func (s *BroadcastService) BroadcastTxn(txn *model.Transaction) error {
	s.updateTxnStatus(txn, "BROADCASTING")
	mergeTxns, txnId, err := s.mergeTransactions(txn)
	if err != nil {
		return err
	}
	algodClient := client.AlgoRandClient()
	_, err = algodClient.SendRawTransaction(mergeTxns).Do(context.Background())
	if err != nil {
		if err.Error() == KnownTLSError {
			s.logger.Error(fmt.Sprintf("Failed to send transaction %s with TLS error trying in next round", txnId))
			return err
		}
		s.updateTxnStatus(txn, "FAILED")
		s.logger.Error(fmt.Sprintf("Failed to send transaction %s with error message: %s", txnId, err))
		return err
	}
	go s.waitForConfirmation(txn, algodClient)
	return nil
}

func (s *BroadcastService) mergeTransactions(txn *model.Transaction) ([]byte, string, error) {
	var stxns []model.SignedTransaction
	if err := s.db.Where("transaction_id = ?", txn.Id).Find(&stxns).Error; err != nil {
		return nil, "", err
	}

	var mergedSignedTxns [][]byte
	for _, signedTxn := range stxns {
		decodedTxn, err := base64Decode(signedTxn.RawSignedTransaction)
		if err != nil {
			s.logger.Error("Error Found in Decoding the transaction with the error message ", zap.Error(err))
			return nil, "", err
		}
		mergedSignedTxns = append(mergedSignedTxns, decodedTxn)
	}
	txnId, signedTxns, err := crypto.MergeMultisigTransactions(mergedSignedTxns...)
	if err != nil {
		s.logger.Error("Error Found in Crypto Merge Multisig Transaction with the error message ", zap.Error(err))
		return nil, "", err
	}
	return signedTxns, txnId, nil
}

func (s *BroadcastService) waitForConfirmation(txn *model.Transaction, client *algod.Client) {
	status, err := client.Status().Do(context.Background())
	if err != nil {
		s.logger.Error(fmt.Sprintf("error getting algod status: %s\n", err))
		return
	}
	lastRound := status.LastRound
	for {
		pt, _, err := client.PendingTransactionInformation(txn.TxnId).Do(context.Background())
		if err != nil {
			s.logger.Error(fmt.Sprintf("error getting pending transaction: %s\n", err))
			s.updateTxnStatus(txn, "DECLINED")
			return
		}
		if pt.ConfirmedRound > 0 {
			s.updateTxnStatus(txn, "BROADCASTED")
			s.logger.Info(fmt.Sprintf("Transaction confirmed in round %d\n", pt.ConfirmedRound))
			break
		}
		s.logger.Info("Waiting for confirmation...")
		lastRound++
		status, err = client.StatusAfterBlock(lastRound).Do(context.Background())
	}
}

func (s *BroadcastService) updateTxnStatus(txn *model.Transaction, status string) error {
	txn.Status = status
	return s.db.Save(&txn).Error
}

func base64Decode(txn string) ([]byte, error) {
	sDec, err := b64.StdEncoding.DecodeString(txn)
	if err != nil {
		return nil, err
	}
	return []byte(sDec), nil
}
