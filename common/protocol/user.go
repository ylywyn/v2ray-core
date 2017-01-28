package protocol

import (
	"v2ray.com/core/common/errors"
)

var (
	ErrUserMissing        = errors.New("User is not specified.")
	ErrAccountMissing     = errors.New("Account is not specified.")
	ErrNonMessageType     = errors.New("Not a protobuf message.")
	ErrUnknownAccountType = errors.New("Unknown account type.")
)

func (v *User) GetTypedAccount() (Account, error) {
	if v.GetAccount() == nil {
		return nil, ErrAccountMissing
	}

	rawAccount, err := v.Account.GetInstance()
	if err != nil {
		return nil, err
	}
	if asAccount, ok := rawAccount.(AsAccount); ok {
		return asAccount.AsAccount()
	}
	if account, ok := rawAccount.(Account); ok {
		return account, nil
	}
	return nil, errors.New("Unknown account type: ", v.Account.Type)
}

func (v *User) GetSettings() UserSettings {
	settings := UserSettings{
		PayloadReadTimeout: 120,
	}
	if v.Level > 0 {
		settings.PayloadReadTimeout = 0
	}
	return settings
}

type UserSettings struct {
	PayloadReadTimeout uint32
}
