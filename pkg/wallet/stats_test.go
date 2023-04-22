package wallet

import (
	"fmt"
	"testing"
)

func TestService_FindPaymentByID_nil(t *testing.T) {
	svc := Service{}
	account, err := svc.RegisterAccount("+992927030909")
	if err != nil {
		t.Errorf("ErrPhoneRegistered")
		return
	}

	err = svc.Deposit(account.ID, 10)
	if err != nil {
		switch err {
		case ErrAmountMustBePositive:
			fmt.Println("Сумма должна быть положительной")
		case ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")
		}
	}
	fmt.Println(account)

	_, err = svc.FindAccountByID(1)
	if err != nil {
		t.Errorf("ErrAccountNotFound: %v", err)
		return
	}
	payment := svc.payments[0]
	_, err = svc.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("ErrPaymentNotFound: %v", err)
		return
	}

}
