package wallet

import "testing"

func TestService_FindAccountById_NotFound(t *testing.T) {
	svc := &Service{}
	svc.RegisterAccount("+998998029829")
	fakeID := int64(777)
	account, err := svc.FindAccountByID(fakeID)
	if account != nil && err != ErrAccountNotFound {
		t.Error("Account Found!")
	}
}

func TestService_FindAccountById_Found(t *testing.T) {
	svc := &Service{}
	svc.RegisterAccount("+998998029829")
	fakeID := int64(1)
	account, err := svc.FindAccountByID(fakeID)
	if account == nil && err == ErrAccountNotFound {
		t.Error("Account Not Found!")
	}
}