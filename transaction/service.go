package transaction

import (
	"crowdfunding/campaign"
	"crowdfunding/payment"
	"errors"
)

type Service interface {
	GetTransactionsByCampaignID(input GetTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionsByCampaignID(input GetTransactionsInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)

	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Maaf anda bukan pemilik campaign")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transactions := Transaction{}
	transactions.Amount = input.Amount
	transactions.CampaignID = input.CampaignID
	transactions.UserID = input.User.ID
	transactions.Status = "pending"

	NewTransactions, err := s.repository.Save(transactions)

	if err != nil {
		return NewTransactions, err
	}

	paymentTransaction := payment.Transaction{
		ID:     NewTransactions.ID,
		Amount: NewTransactions.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)

	if err != nil {
		return NewTransactions, err
	}

	NewTransactions.PaymentURL = paymentURL

	NewTransactions, err = s.repository.Update(NewTransactions)

	if err != nil {
		return NewTransactions, err
	}

	return NewTransactions, nil
}
