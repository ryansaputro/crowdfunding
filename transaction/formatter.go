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

	campaignsTransactionsFormatter := []CampaignTransactionsFormatter{}

	for _, transaction := range transactions {
		campaignsTransactionFormatter := FormatterCampaignTransaction(transaction)
		campaignsTransactionsFormatter = append(campaignsTransactionsFormatter, campaignsTransactionFormatter)
	}

	return campaignsTransactionsFormatter

}
