package business

import (
	"reflect"
	"testing"

	"github.com/stori/internal/schema"
)

func TestBusiness_CalculateBalance(t *testing.T) {
	type args struct {
		transactions []schema.Transaction
	}
	tests := []struct {
		name    string
		args    args
		want    schema.Summary
		wantErr bool
	}{
		{
			name: "Success: Balance example",
			args: args{
				transactions: []schema.Transaction{
					{
						ID:          0,
						Date:        "2022/07/15",
						Transaction: 60.5,
					},
					{
						ID:          1,
						Date:        "2022/07/28",
						Transaction: -10.3,
					},
					{
						ID:          2,
						Date:        "2022/08/02",
						Transaction: -20.46,
					},
					{
						ID:          3,
						Date:        "2022/08/13",
						Transaction: 10,
					},
				},
			},
			want: schema.Summary{
				TotalBalance: 39.74,
				MonthTransactions: map[string]int64{
					"July":   2,
					"August": 2,
				},
				AverageDebitAmout:   -15.38,
				AverageCreditAmount: 35.25,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Business{}
			got, err := b.calculateBalance(tt.args.transactions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Business.CalculateBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Business.CalculateBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
