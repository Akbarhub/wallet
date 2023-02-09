package wallet

import (
	"github.com/Akbarhub/wallet/pkg/types"
)

type Service struct{
	nextAccountID	int64
	accounts		[]types.Account
	payments		[]types.Payment
}

func RegisterAccount(s *Service, phone types.Phone) {
	for _, account := range s.accounts {
		if account.Phone == phone{
			return
		}
	}
	s.nextAccountID++
	s.accounts = append(s.accounts, types.Account{
		ID:			s.nextAccountID,
		Phone:		phone,
		Balance:	0,
	})
}