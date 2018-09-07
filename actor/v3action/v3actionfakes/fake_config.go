// Code generated by counterfeiter. DO NOT EDIT.
package v3actionfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/cli/actor/v3action"
)

type FakeConfig struct {
	AccessTokenStub        func() string
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct{}
	accessTokenReturns     struct {
		result1 string
	}
	accessTokenReturnsOnCall map[int]struct {
		result1 string
	}
	DialTimeoutStub        func() time.Duration
	dialTimeoutMutex       sync.RWMutex
	dialTimeoutArgsForCall []struct{}
	dialTimeoutReturns     struct {
		result1 time.Duration
	}
	dialTimeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	PollingIntervalStub        func() time.Duration
	pollingIntervalMutex       sync.RWMutex
	pollingIntervalArgsForCall []struct{}
	pollingIntervalReturns     struct {
		result1 time.Duration
	}
	pollingIntervalReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	SSHOAuthClientStub        func() string
	sSHOAuthClientMutex       sync.RWMutex
	sSHOAuthClientArgsForCall []struct{}
	sSHOAuthClientReturns     struct {
		result1 string
	}
	sSHOAuthClientReturnsOnCall map[int]struct {
		result1 string
	}
	StartupTimeoutStub        func() time.Duration
	startupTimeoutMutex       sync.RWMutex
	startupTimeoutArgsForCall []struct{}
	startupTimeoutReturns     struct {
		result1 time.Duration
	}
	startupTimeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	StagingTimeoutStub        func() time.Duration
	stagingTimeoutMutex       sync.RWMutex
	stagingTimeoutArgsForCall []struct{}
	stagingTimeoutReturns     struct {
		result1 time.Duration
	}
	stagingTimeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) AccessToken() string {
	fake.accessTokenMutex.Lock()
	ret, specificReturn := fake.accessTokenReturnsOnCall[len(fake.accessTokenArgsForCall)]
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct{}{})
	fake.recordInvocation("AccessToken", []interface{}{})
	fake.accessTokenMutex.Unlock()
	if fake.AccessTokenStub != nil {
		return fake.AccessTokenStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.accessTokenReturns.result1
}

func (fake *FakeConfig) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakeConfig) AccessTokenReturns(result1 string) {
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) AccessTokenReturnsOnCall(i int, result1 string) {
	fake.AccessTokenStub = nil
	if fake.accessTokenReturnsOnCall == nil {
		fake.accessTokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.accessTokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) DialTimeout() time.Duration {
	fake.dialTimeoutMutex.Lock()
	ret, specificReturn := fake.dialTimeoutReturnsOnCall[len(fake.dialTimeoutArgsForCall)]
	fake.dialTimeoutArgsForCall = append(fake.dialTimeoutArgsForCall, struct{}{})
	fake.recordInvocation("DialTimeout", []interface{}{})
	fake.dialTimeoutMutex.Unlock()
	if fake.DialTimeoutStub != nil {
		return fake.DialTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.dialTimeoutReturns.result1
}

func (fake *FakeConfig) DialTimeoutCallCount() int {
	fake.dialTimeoutMutex.RLock()
	defer fake.dialTimeoutMutex.RUnlock()
	return len(fake.dialTimeoutArgsForCall)
}

func (fake *FakeConfig) DialTimeoutReturns(result1 time.Duration) {
	fake.DialTimeoutStub = nil
	fake.dialTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) DialTimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.DialTimeoutStub = nil
	if fake.dialTimeoutReturnsOnCall == nil {
		fake.dialTimeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.dialTimeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) PollingInterval() time.Duration {
	fake.pollingIntervalMutex.Lock()
	ret, specificReturn := fake.pollingIntervalReturnsOnCall[len(fake.pollingIntervalArgsForCall)]
	fake.pollingIntervalArgsForCall = append(fake.pollingIntervalArgsForCall, struct{}{})
	fake.recordInvocation("PollingInterval", []interface{}{})
	fake.pollingIntervalMutex.Unlock()
	if fake.PollingIntervalStub != nil {
		return fake.PollingIntervalStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pollingIntervalReturns.result1
}

func (fake *FakeConfig) PollingIntervalCallCount() int {
	fake.pollingIntervalMutex.RLock()
	defer fake.pollingIntervalMutex.RUnlock()
	return len(fake.pollingIntervalArgsForCall)
}

func (fake *FakeConfig) PollingIntervalReturns(result1 time.Duration) {
	fake.PollingIntervalStub = nil
	fake.pollingIntervalReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) PollingIntervalReturnsOnCall(i int, result1 time.Duration) {
	fake.PollingIntervalStub = nil
	if fake.pollingIntervalReturnsOnCall == nil {
		fake.pollingIntervalReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.pollingIntervalReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) SSHOAuthClient() string {
	fake.sSHOAuthClientMutex.Lock()
	ret, specificReturn := fake.sSHOAuthClientReturnsOnCall[len(fake.sSHOAuthClientArgsForCall)]
	fake.sSHOAuthClientArgsForCall = append(fake.sSHOAuthClientArgsForCall, struct{}{})
	fake.recordInvocation("SSHOAuthClient", []interface{}{})
	fake.sSHOAuthClientMutex.Unlock()
	if fake.SSHOAuthClientStub != nil {
		return fake.SSHOAuthClientStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sSHOAuthClientReturns.result1
}

func (fake *FakeConfig) SSHOAuthClientCallCount() int {
	fake.sSHOAuthClientMutex.RLock()
	defer fake.sSHOAuthClientMutex.RUnlock()
	return len(fake.sSHOAuthClientArgsForCall)
}

func (fake *FakeConfig) SSHOAuthClientReturns(result1 string) {
	fake.SSHOAuthClientStub = nil
	fake.sSHOAuthClientReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) SSHOAuthClientReturnsOnCall(i int, result1 string) {
	fake.SSHOAuthClientStub = nil
	if fake.sSHOAuthClientReturnsOnCall == nil {
		fake.sSHOAuthClientReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.sSHOAuthClientReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) StartupTimeout() time.Duration {
	fake.startupTimeoutMutex.Lock()
	ret, specificReturn := fake.startupTimeoutReturnsOnCall[len(fake.startupTimeoutArgsForCall)]
	fake.startupTimeoutArgsForCall = append(fake.startupTimeoutArgsForCall, struct{}{})
	fake.recordInvocation("StartupTimeout", []interface{}{})
	fake.startupTimeoutMutex.Unlock()
	if fake.StartupTimeoutStub != nil {
		return fake.StartupTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startupTimeoutReturns.result1
}

func (fake *FakeConfig) StartupTimeoutCallCount() int {
	fake.startupTimeoutMutex.RLock()
	defer fake.startupTimeoutMutex.RUnlock()
	return len(fake.startupTimeoutArgsForCall)
}

func (fake *FakeConfig) StartupTimeoutReturns(result1 time.Duration) {
	fake.StartupTimeoutStub = nil
	fake.startupTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) StartupTimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.StartupTimeoutStub = nil
	if fake.startupTimeoutReturnsOnCall == nil {
		fake.startupTimeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.startupTimeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) StagingTimeout() time.Duration {
	fake.stagingTimeoutMutex.Lock()
	ret, specificReturn := fake.stagingTimeoutReturnsOnCall[len(fake.stagingTimeoutArgsForCall)]
	fake.stagingTimeoutArgsForCall = append(fake.stagingTimeoutArgsForCall, struct{}{})
	fake.recordInvocation("StagingTimeout", []interface{}{})
	fake.stagingTimeoutMutex.Unlock()
	if fake.StagingTimeoutStub != nil {
		return fake.StagingTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stagingTimeoutReturns.result1
}

func (fake *FakeConfig) StagingTimeoutCallCount() int {
	fake.stagingTimeoutMutex.RLock()
	defer fake.stagingTimeoutMutex.RUnlock()
	return len(fake.stagingTimeoutArgsForCall)
}

func (fake *FakeConfig) StagingTimeoutReturns(result1 time.Duration) {
	fake.StagingTimeoutStub = nil
	fake.stagingTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) StagingTimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.StagingTimeoutStub = nil
	if fake.stagingTimeoutReturnsOnCall == nil {
		fake.stagingTimeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.stagingTimeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	fake.dialTimeoutMutex.RLock()
	defer fake.dialTimeoutMutex.RUnlock()
	fake.pollingIntervalMutex.RLock()
	defer fake.pollingIntervalMutex.RUnlock()
	fake.sSHOAuthClientMutex.RLock()
	defer fake.sSHOAuthClientMutex.RUnlock()
	fake.startupTimeoutMutex.RLock()
	defer fake.startupTimeoutMutex.RUnlock()
	fake.stagingTimeoutMutex.RLock()
	defer fake.stagingTimeoutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v3action.Config = new(FakeConfig)
