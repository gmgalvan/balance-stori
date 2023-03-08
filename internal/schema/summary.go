package schema

type Summary struct {
	TotalBalance        float64
	MonthTransactions   map[string]int64
	AverageDebitAmout   float64
	AverageCreditAmount float64
}
