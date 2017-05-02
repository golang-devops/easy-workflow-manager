package mailing

import (
	"net/mail"
)

type AddressesProvider interface {
	Addresses() []*mail.Address
}

func NewSimpleAddressesProvider(addresses ...*mail.Address) AddressesProvider {
	return &simpleAddressesProvider{
		addresses: addresses,
	}
}

type simpleAddressesProvider struct {
	addresses []*mail.Address
}

func (s *simpleAddressesProvider) Addresses() []*mail.Address {
	return s.addresses
}
