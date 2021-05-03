package main

import (
	"github.com/stretchr/testify/mock"
)

// The mock store contains additonal methods for inspection
type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateTicket(ticket *Ticket) error {
	/*
		When this method is called, `m.Called` records the call, and also
		returns the result that we pass to it (which you will see in the
		handler tests)
	*/
	rets := m.Called(ticket)
	return rets.Error(0)
}

func (m *MockStore) GetTickets() ([]*Ticket, error) {
	rets := m.Called()
	/*
		Since `rets.Get()` is a generic method, that returns whatever we pass to it,
		we need to typecast it to the type we expect, which in this case is []*Bird
	*/
	return rets.Get(0).([]*Ticket), rets.Error(1)
}

func (m *MockStore) GetDeletedTickets() ([]*Ticket, error) {
	rets := m.Called()
	/*
		Since `rets.Get()` is a generic method, that returns whatever we pass to it,
		we need to typecast it to the type we expect, which in this case is []*Bird
	*/
	return rets.Get(0).([]*Ticket), rets.Error(1)
}

func (m *MockStore) DeleteTicket(ticket *Ticket) error {
	rets := m.Called(ticket)
	return rets.Error(0)
}

func (m *MockStore) InsertDeleteTicket(ticket *Ticket) error {
	rets := m.Called(ticket)
	return rets.Error(0)
}

func (m *MockStore) RecoverTicket(ticket *Ticket) error {
	rets := m.Called(ticket)
	return rets.Error(0)
}

func (m *MockStore) InsertRecoverTicket(ticket *Ticket) error {
	rets := m.Called(ticket)
	return rets.Error(0)
}

func (m *MockStore) UpdateTicket(ticket *Ticket, newId string) error {
	rets := m.Called(ticket, newId)
	return rets.Error(0)
}

func (m *MockStore) FilterTicket(filter string, searchfilter string) error {
	rets := m.Called(filter, searchfilter)
	return rets.Error(0)
}

func (m *MockStore) DeleteFilterTicket() error {
	rets := m.Called()
	return rets.Error(0)
}

func InitMockStore() *MockStore {
	/*
		Like the InitStore function we defined earlier, this function
		also initializes the store variable, but this time, it assigns
		a new MockStore instance to it, instead of an actual store
	*/
	s := new(MockStore)
	store = s
	return s
}
