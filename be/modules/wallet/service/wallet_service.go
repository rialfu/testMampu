package service

import (
	"context"
	"fmt"
	"log"
	"rialfu/wallet/database/entities"
	"rialfu/wallet/modules/wallet/dto"
	"rialfu/wallet/modules/wallet/repository"
	"rialfu/wallet/pkg/constants"
	"rialfu/wallet/pkg/helpers"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type WalletService interface {
	CheckBalance(ctx context.Context, userId string) (dto.BalanceResponse, error)
	WithdrawProcess(ctx context.Context, userId string, data dto.WithdrawRequest) (dto.WithdrawResponse, error)
	StoreBalance(ctx context.Context, data dto.StoreRequest) (dto.StoreResponse, error)
}

type walletService struct {
	walletRepository repository.WalletRepository
	wlr              repository.WalletLedgerRepository
	tr               repository.TransactionRepository
	mbr              repository.MasterBankRepositry
	wr               repository.WithdrawRepository
	dr               repository.DepositRepository
	db               *gorm.DB
	transactor       helpers.Transactor
}

func NewUserService(
	walletRepo repository.WalletRepository,
	tr repository.TransactionRepository,
	mbr repository.MasterBankRepositry,
	wlr repository.WalletLedgerRepository,
	wr repository.WithdrawRepository,
	dr repository.DepositRepository,
	db *gorm.DB,
) WalletService {
	transactor := helpers.NewGormTransactor(db)
	return &walletService{
		walletRepository: walletRepo,
		db:               db,
		tr:               tr,
		mbr:              mbr,
		transactor:       transactor,
		wlr:              wlr,
		wr:               wr,
		dr:               dr,
	}
}
func (s *walletService) CheckBalance(ctx context.Context, userId string) (dto.BalanceResponse, error) {
	data, isExist, err := s.walletRepository.GetByUserId(ctx, s.db, userId, false)
	if err != nil {
		return dto.BalanceResponse{}, err
	}
	if isExist == false {
		return dto.BalanceResponse{}, constants.ErrDataNotFound
	}

	return dto.BalanceResponse{ID: strconv.FormatUint(data.ID, 10), Balance: data.Balance.InexactFloat64(), LastUpdate: data.UpdatedAt.Format("2006-01-02 15:04:05")}, nil
}
func (s *walletService) WithdrawProcess(ctx context.Context, userId string, data dto.WithdrawRequest) (dto.WithdrawResponse, error) {
	var fallbackValue dto.WithdrawResponse
	var transactionData entities.Transaction
	err := s.transactor.WithTransaction(ctx, func(txCtx context.Context) error {
		dataWallet, isExist, err := s.walletRepository.GetByUserId(txCtx, s.db, userId, true)
		if err != nil {

			return err
		}
		if isExist == false {
			return constants.ErrDataNotFound
		}

		balance := decimal.NewFromFloat(data.Balance)
		newBalance := dataWallet.Balance.Sub(balance)
		if newBalance.IsNegative() {
			return dto.ErrInsufficientBalance
		}
		dataBank, isExist, err := s.mbr.GetById(txCtx, s.db, strconv.FormatUint(uint64(data.TargetBank), 10))
		if err != nil {
			return err
		}
		if isExist == false {
			return dto.ErrTragetBankNotFound
		}
		var noRef string
		dateTime := time.Now()
		for {
			noRef = s.generateRef(dateTime)
			_, isExist, err = s.tr.CheckReferenceNo(txCtx, s.db, noRef, false)
			if err != nil {
				return err
			}
			if isExist == false {
				break
			}
		}
		dataWallet.Balance = newBalance
		// err = s.walletRepository.UpdateBalance(txCtx, s.db, dataWallet)
		dataInsert := entities.Transaction{
			WalletID:        dataWallet.ID,
			UserID:          dataWallet.UserID,
			TransactionType: "withdraw",
			DateTrans:       dateTime,
			Status:          3,
			ReferenceNo:     noRef,
		}

		transactionData, err = s.tr.CreateTransaction(txCtx, s.db, dataInsert)
		if err != nil {
			return err
		}
		fallbackValue = dto.WithdrawResponse{
			ReferenceNo:   dataInsert.ReferenceNo,
			Fee:           decimal.Zero,
			Amount:        balance,
			TargetBank:    dataBank.BankName,
			TargetAccount: data.TargetAccount,
		}
		return nil

	})
	if transactionData.ReferenceNo != "" && err == nil {
		go s.transactor.WithTransaction(context.Background(), func(txCtx context.Context) error {
			dataTrans, isExist, err := s.tr.CheckReferenceNo(txCtx, s.db, transactionData.ReferenceNo, true)
			if err != nil {
				log.Fatalf("error running server: %v", err)
				return err
			}
			if isExist == false {
				return constants.ErrDataNotFound
			}
			if dataTrans.Status != 3 {
				return nil
			}
			dataWallet, isExist, err := s.walletRepository.GetByUserId(txCtx, s.db, userId, true)
			if err != nil {
				log.Fatalf("error running server: %v", err)
				return err
			}
			if isExist == false {
				return constants.ErrDataNotFound
			}
			balance := decimal.NewFromFloat(data.Balance)
			if dataWallet.Balance.LessThan(balance) {
				dataTrans.Status = 2
				s.tr.Update(txCtx, s.db, dataTrans)
				return nil
			}
			balanceBefore := dataWallet.Balance
			balanceAfter := balanceBefore.Sub(balance)
			ledger := entities.WalletLedger{
				WalletID:      dataTrans.WalletID,
				TransactionID: dataTrans.ID,
				Direction:     "debit",
				Amount:        balance,
				BalanceBefore: balanceBefore,
				BalanceAfter:  balanceAfter,
			}
			_, err = s.wlr.Create(txCtx, s.db, ledger)
			if err != nil {
				log.Fatalf("error running server: %v", err)
				dataTrans.Status = 2
				s.tr.Update(txCtx, s.db, dataTrans)
				return err
			}
			dataWallet.Balance = balanceAfter
			err = s.walletRepository.UpdateBalance(txCtx, s.db, dataWallet)
			if err != nil {
				log.Fatalf("error running server: %v", err)
				dataTrans.Status = 2
				s.tr.Update(txCtx, s.db, dataTrans)
				return err
			}
			_, err = s.wr.Create(txCtx, s.db, entities.Withdrawal{
				TransactionID: dataTrans.ID,
				Amount:        balance,
				Fee:           decimal.Zero,
				TargetBank:    data.TargetBank,
				TargetAccount: data.TargetAccount,
			})
			if err != nil {
				log.Fatalf("error running server: %v", err)
				dataTrans.Status = 2
				s.tr.Update(txCtx, s.db, dataTrans)
				return err
			}
			dataTrans.Status = 1
			err = s.tr.Update(txCtx, s.db, dataTrans)
			if err != nil {
				log.Fatalf("error running server: %v", err)
				dataTrans.Status = 2
				s.tr.Update(txCtx, s.db, dataTrans)
				return err
			}
			return nil

		})
	}

	return fallbackValue, err
}
func (s *walletService) StoreBalance(ctx context.Context, data dto.StoreRequest) (dto.StoreResponse, error) {
	var fallbackValue dto.StoreResponse
	err := s.transactor.WithTransaction(ctx, func(txCtx context.Context) error {

		dataWallet, isExist, err := s.walletRepository.GetById(txCtx, s.db, data.WalletID, true)
		if err != nil {
			return err
		}
		if isExist == false {
			return constants.ErrDataNotFound
		}

		balance := decimal.NewFromFloat(data.Amount)
		balanceBefore := dataWallet.Balance
		balanceAfter := balanceBefore.Add(balance)

		var noRef string
		dateTime := time.Now()
		for {
			noRef = s.generateRef(dateTime)
			_, isExist, err = s.tr.CheckReferenceNo(txCtx, s.db, noRef, false)
			if err != nil {
				return err
			}
			if isExist == false {
				break
			}
		}

		dataInsert := entities.Transaction{
			WalletID:        dataWallet.ID,
			UserID:          dataWallet.UserID,
			TransactionType: "deposit",
			DateTrans:       dateTime,
			Status:          3,
			ReferenceNo:     noRef,
		}

		transactionData, err := s.tr.CreateTransaction(txCtx, s.db, dataInsert)
		if err != nil {
			return err
		}
		ledger := entities.WalletLedger{
			WalletID:      transactionData.WalletID,
			TransactionID: transactionData.ID,
			Direction:     "credit",
			Amount:        balance,
			BalanceBefore: balanceBefore,
			BalanceAfter:  balanceAfter,
		}
		_, err = s.wlr.Create(txCtx, s.db, ledger)
		if err != nil {
			return err
		}
		depositInfo := entities.Deposit{
			TransactionID: transactionData.ID,
			Source:        data.PaymentType,
		}
		_, err = s.dr.Create(txCtx, s.db, depositInfo)
		if err != nil {
			return err
		}
		dataWallet.Balance = balanceAfter
		err = s.walletRepository.UpdateBalance(txCtx, s.db, dataWallet)
		if err != nil {
			return err
		}

		transactionData.Status = 1
		err = s.tr.Update(txCtx, s.db, transactionData)
		if err != nil {
			return err
		}
		fallbackValue = dto.StoreResponse{
			ReferenceNo: noRef,
			Fee:         decimal.Zero,
			Amount:      balance,
		}
		return nil

	})
	return fallbackValue, err
}

func (s *walletService) generateRef(dataTime time.Time) string {
	datePart := dataTime.Format("20060102") // 8 char
	randomPart := helpers.GenerateRandomString(17)
	ref := fmt.Sprintf("%s%s", datePart, randomPart)

	return ref
}
