package transaction

import (
	"time"
)

type CampaignTransactionsFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionsFormatter struct {
	ID        int                                   `json:"id"`
	Amount    int                                   `json:"amount"`
	Status    string                                `json:"status"`
	CreatedAt time.Time                             `json:"created_at"`
	Campaign  UserTransactionsListCampaignFormatter `json:"campaign"`
}

type UserTransactionsListCampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatterCampaignTransaction(transaction Transaction) CampaignTransactionsFormatter {
	formatter := CampaignTransactionsFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return formatter
}

func FormatterCampaignTransactions(transactions []Transaction) []CampaignTransactionsFormatter {

	if len(transactions) == 0 {
		return []CampaignTransactionsFormatter{}
	}

	campaignsTransactionsFormatter := []CampaignTransactionsFormatter{}

	for _, transaction := range transactions {
		campaignsTransactionFormatter := FormatterCampaignTransaction(transaction)
		campaignsTransactionsFormatter = append(campaignsTransactionsFormatter, campaignsTransactionFormatter)
	}

	return campaignsTransactionsFormatter

}

func FormatterUserTransaction(transaction Transaction) UserTransactionsFormatter {
	formatter := UserTransactionsFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := UserTransactionsListCampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name

	campaignFormatter.ImageURL = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName

	}

	formatter.Campaign = campaignFormatter

	return formatter

}

func FormatterUserTransactions(transactions []Transaction) []UserTransactionsFormatter {

	if len(transactions) == 0 {
		return []UserTransactionsFormatter{}
	}
	userTransactionsFormatter := []UserTransactionsFormatter{}

	for _, transaction := range transactions {
		userTransactionFormatter := FormatterUserTransaction(transaction)
		userTransactionsFormatter = append(userTransactionsFormatter, userTransactionFormatter)
	}

	return userTransactionsFormatter

}

func FormatterTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:         transaction.ID,
		CampaignID: transaction.CampaignID,
		UserID:     transaction.UserID,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		Code:       transaction.Code,
		PaymentURL: transaction.PaymentURL,
	}

	return formatter
}
