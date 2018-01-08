// Code generated by counterfeiter. DO NOT EDIT.
package cmdfakes

import (
	"context"
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/perm/cmd"
)

type FakeQueryProbe struct {
	CleanupStub        func(context.Context, lager.Logger, string) error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}
	cleanupReturns struct {
		result1 error
	}
	cleanupReturnsOnCall map[int]struct {
		result1 error
	}
	SetupStub        func(context.Context, lager.Logger, string) error
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}
	setupReturns struct {
		result1 error
	}
	setupReturnsOnCall map[int]struct {
		result1 error
	}
	RunStub        func(context.Context, lager.Logger, string) (bool, []time.Duration, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}
	runReturns struct {
		result1 bool
		result2 []time.Duration
		result3 error
	}
	runReturnsOnCall map[int]struct {
		result1 bool
		result2 []time.Duration
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeQueryProbe) Cleanup(arg1 context.Context, arg2 lager.Logger, arg3 string) error {
	fake.cleanupMutex.Lock()
	ret, specificReturn := fake.cleanupReturnsOnCall[len(fake.cleanupArgsForCall)]
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Cleanup", []interface{}{arg1, arg2, arg3})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanupReturns.result1
}

func (fake *FakeQueryProbe) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeQueryProbe) CleanupArgsForCall(i int) (context.Context, lager.Logger, string) {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return fake.cleanupArgsForCall[i].arg1, fake.cleanupArgsForCall[i].arg2, fake.cleanupArgsForCall[i].arg3
}

func (fake *FakeQueryProbe) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeQueryProbe) CleanupReturnsOnCall(i int, result1 error) {
	fake.CleanupStub = nil
	if fake.cleanupReturnsOnCall == nil {
		fake.cleanupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeQueryProbe) Setup(arg1 context.Context, arg2 lager.Logger, arg3 string) error {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Setup", []interface{}{arg1, arg2, arg3})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setupReturns.result1
}

func (fake *FakeQueryProbe) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *FakeQueryProbe) SetupArgsForCall(i int) (context.Context, lager.Logger, string) {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return fake.setupArgsForCall[i].arg1, fake.setupArgsForCall[i].arg2, fake.setupArgsForCall[i].arg3
}

func (fake *FakeQueryProbe) SetupReturns(result1 error) {
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeQueryProbe) SetupReturnsOnCall(i int, result1 error) {
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeQueryProbe) Run(arg1 context.Context, arg2 lager.Logger, arg3 string) (bool, []time.Duration, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Run", []interface{}{arg1, arg2, arg3})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.runReturns.result1, fake.runReturns.result2, fake.runReturns.result3
}

func (fake *FakeQueryProbe) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeQueryProbe) RunArgsForCall(i int) (context.Context, lager.Logger, string) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1, fake.runArgsForCall[i].arg2, fake.runArgsForCall[i].arg3
}

func (fake *FakeQueryProbe) RunReturns(result1 bool, result2 []time.Duration, result3 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 bool
		result2 []time.Duration
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeQueryProbe) RunReturnsOnCall(i int, result1 bool, result2 []time.Duration, result3 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 []time.Duration
			result3 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 bool
		result2 []time.Duration
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeQueryProbe) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeQueryProbe) recordInvocation(key string, args []interface{}) {
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

var _ cmd.QueryProbe = new(FakeQueryProbe)
