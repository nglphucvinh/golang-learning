package wallet

import (
	"errors"
	"fmt"
)

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return errors.New("incoming code not match")
	}
	fmt.Println("Security code verified!")
	return nil
}
