package wallet

import "testing"

func TestService_FindAccountByID_NotFound(t *testing.T) {
	svc := &Service{}
	svc.RegisterAccount("+998998029829")
	fakeID := int64(777)
	account, err := svc.FindAccountByID(fakeID)
	if account != nil && err != ErrAccountNotFound {
		t.Error("Account Found!")
	}
}

func TestService_FindAccountByID_Found(t *testing.T) {
	svc := &Service{}
	svc.RegisterAccount("+998998029829")
	id := int64(1)
	account, err := svc.FindAccountByID(id)
	if account == nil && err == ErrAccountNotFound {
		t.Error("Account Not Found!")
	}
}

func TestService_Reject_Success(t *testing.T) {
	svc := &Service{}
	svc.RegisterAccount("+998998029829")
	svc.Deposit(1, 10_000)
	payment, err := svc.Pay(1, 5_000, "mobile")
	if err != nil {
		t.Error(err)
	}

	err2 := svc.Reject(payment.ID)
	if err2 != nil {
		t.Error("Fail")
	}

}

func TestService_Reject_Fail(t *testing.T) {
	svc := &Service{}
	svc.RegisterAccount("+998998029829")
	svc.Deposit(1, 10_000)
	payment, err := svc.Pay(1, 5_000, "mobile")
	if err != nil {
		t.Error(err)
	}
	fakeID := payment.ID + "test"
	err2 := svc.Reject(fakeID)
	if err2 == nil {
		t.Error("Fail")
	}

}