// Code generated by mockery v2.43.2. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	bootstrap "github.com/absmach/magistrala/bootstrap"

	mock "github.com/stretchr/testify/mock"
)

// ConfigRepository is an autogenerated mock type for the ConfigRepository type
type ConfigRepository struct {
	mock.Mock
}

// ChangeState provides a mock function with given fields: ctx, domainID, id, state
func (_m *ConfigRepository) ChangeState(ctx context.Context, domainID string, id string, state bootstrap.State) error {
	ret := _m.Called(ctx, domainID, id, state)

	if len(ret) == 0 {
		panic("no return value specified for ChangeState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bootstrap.State) error); ok {
		r0 = rf(ctx, domainID, id, state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectThing provides a mock function with given fields: ctx, channelID, thingID
func (_m *ConfigRepository) ConnectThing(ctx context.Context, channelID string, thingID string) error {
	ret := _m.Called(ctx, channelID, thingID)

	if len(ret) == 0 {
		panic("no return value specified for ConnectThing")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, channelID, thingID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisconnectThing provides a mock function with given fields: ctx, channelID, thingID
func (_m *ConfigRepository) DisconnectThing(ctx context.Context, channelID string, thingID string) error {
	ret := _m.Called(ctx, channelID, thingID)

	if len(ret) == 0 {
		panic("no return value specified for DisconnectThing")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, channelID, thingID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListExisting provides a mock function with given fields: ctx, domainID, ids
func (_m *ConfigRepository) ListExisting(ctx context.Context, domainID string, ids []string) ([]bootstrap.Channel, error) {
	ret := _m.Called(ctx, domainID, ids)

	if len(ret) == 0 {
		panic("no return value specified for ListExisting")
	}

	var r0 []bootstrap.Channel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) ([]bootstrap.Channel, error)); ok {
		return rf(ctx, domainID, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) []bootstrap.Channel); ok {
		r0 = rf(ctx, domainID, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bootstrap.Channel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, domainID, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, domainID, id
func (_m *ConfigRepository) Remove(ctx context.Context, domainID string, id string) error {
	ret := _m.Called(ctx, domainID, id)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, domainID, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveChannel provides a mock function with given fields: ctx, id
func (_m *ConfigRepository) RemoveChannel(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for RemoveChannel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveThing provides a mock function with given fields: ctx, id
func (_m *ConfigRepository) RemoveThing(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for RemoveThing")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveAll provides a mock function with given fields: ctx, domainID, thingIDs, filter, offset, limit
func (_m *ConfigRepository) RetrieveAll(ctx context.Context, domainID string, thingIDs []string, filter bootstrap.Filter, offset uint64, limit uint64) bootstrap.ConfigsPage {
	ret := _m.Called(ctx, domainID, thingIDs, filter, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAll")
	}

	var r0 bootstrap.ConfigsPage
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, bootstrap.Filter, uint64, uint64) bootstrap.ConfigsPage); ok {
		r0 = rf(ctx, domainID, thingIDs, filter, offset, limit)
	} else {
		r0 = ret.Get(0).(bootstrap.ConfigsPage)
	}

	return r0
}

// RetrieveByExternalID provides a mock function with given fields: ctx, externalID
func (_m *ConfigRepository) RetrieveByExternalID(ctx context.Context, externalID string) (bootstrap.Config, error) {
	ret := _m.Called(ctx, externalID)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByExternalID")
	}

	var r0 bootstrap.Config
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bootstrap.Config, error)); ok {
		return rf(ctx, externalID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bootstrap.Config); ok {
		r0 = rf(ctx, externalID)
	} else {
		r0 = ret.Get(0).(bootstrap.Config)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, externalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveByID provides a mock function with given fields: ctx, domainID, id
func (_m *ConfigRepository) RetrieveByID(ctx context.Context, domainID string, id string) (bootstrap.Config, error) {
	ret := _m.Called(ctx, domainID, id)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByID")
	}

	var r0 bootstrap.Config
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bootstrap.Config, error)); ok {
		return rf(ctx, domainID, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bootstrap.Config); ok {
		r0 = rf(ctx, domainID, id)
	} else {
		r0 = ret.Get(0).(bootstrap.Config)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, domainID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, cfg, chsConnIDs
func (_m *ConfigRepository) Save(ctx context.Context, cfg bootstrap.Config, chsConnIDs []string) (string, error) {
	ret := _m.Called(ctx, cfg, chsConnIDs)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bootstrap.Config, []string) (string, error)); ok {
		return rf(ctx, cfg, chsConnIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bootstrap.Config, []string) string); ok {
		r0 = rf(ctx, cfg, chsConnIDs)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, bootstrap.Config, []string) error); ok {
		r1 = rf(ctx, cfg, chsConnIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, cfg
func (_m *ConfigRepository) Update(ctx context.Context, cfg bootstrap.Config) error {
	ret := _m.Called(ctx, cfg)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, bootstrap.Config) error); ok {
		r0 = rf(ctx, cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCert provides a mock function with given fields: ctx, domainID, thingID, clientCert, clientKey, caCert
func (_m *ConfigRepository) UpdateCert(ctx context.Context, domainID string, thingID string, clientCert string, clientKey string, caCert string) (bootstrap.Config, error) {
	ret := _m.Called(ctx, domainID, thingID, clientCert, clientKey, caCert)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCert")
	}

	var r0 bootstrap.Config
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) (bootstrap.Config, error)); ok {
		return rf(ctx, domainID, thingID, clientCert, clientKey, caCert)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) bootstrap.Config); ok {
		r0 = rf(ctx, domainID, thingID, clientCert, clientKey, caCert)
	} else {
		r0 = ret.Get(0).(bootstrap.Config)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, string) error); ok {
		r1 = rf(ctx, domainID, thingID, clientCert, clientKey, caCert)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateChannel provides a mock function with given fields: ctx, c
func (_m *ConfigRepository) UpdateChannel(ctx context.Context, c bootstrap.Channel) error {
	ret := _m.Called(ctx, c)

	if len(ret) == 0 {
		panic("no return value specified for UpdateChannel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, bootstrap.Channel) error); ok {
		r0 = rf(ctx, c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateConnections provides a mock function with given fields: ctx, domainID, id, channels, connections
func (_m *ConfigRepository) UpdateConnections(ctx context.Context, domainID string, id string, channels []bootstrap.Channel, connections []string) error {
	ret := _m.Called(ctx, domainID, id, channels, connections)

	if len(ret) == 0 {
		panic("no return value specified for UpdateConnections")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []bootstrap.Channel, []string) error); ok {
		r0 = rf(ctx, domainID, id, channels, connections)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewConfigRepository creates a new instance of ConfigRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfigRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConfigRepository {
	mock := &ConfigRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
