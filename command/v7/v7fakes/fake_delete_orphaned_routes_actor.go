// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeDeleteOrphanedRoutesActor struct {
	DeleteOrphanedRoutesStub        func(string) (v7action.Warnings, error)
	deleteOrphanedRoutesMutex       sync.RWMutex
	deleteOrphanedRoutesArgsForCall []struct {
		arg1 string
	}
	deleteOrphanedRoutesReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	deleteOrphanedRoutesReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteOrphanedRoutes(arg1 string) (v7action.Warnings, error) {
	fake.deleteOrphanedRoutesMutex.Lock()
	ret, specificReturn := fake.deleteOrphanedRoutesReturnsOnCall[len(fake.deleteOrphanedRoutesArgsForCall)]
	fake.deleteOrphanedRoutesArgsForCall = append(fake.deleteOrphanedRoutesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteOrphanedRoutes", []interface{}{arg1})
	fake.deleteOrphanedRoutesMutex.Unlock()
	if fake.DeleteOrphanedRoutesStub != nil {
		return fake.DeleteOrphanedRoutesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteOrphanedRoutesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteOrphanedRoutesCallCount() int {
	fake.deleteOrphanedRoutesMutex.RLock()
	defer fake.deleteOrphanedRoutesMutex.RUnlock()
	return len(fake.deleteOrphanedRoutesArgsForCall)
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteOrphanedRoutesCalls(stub func(string) (v7action.Warnings, error)) {
	fake.deleteOrphanedRoutesMutex.Lock()
	defer fake.deleteOrphanedRoutesMutex.Unlock()
	fake.DeleteOrphanedRoutesStub = stub
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteOrphanedRoutesArgsForCall(i int) string {
	fake.deleteOrphanedRoutesMutex.RLock()
	defer fake.deleteOrphanedRoutesMutex.RUnlock()
	argsForCall := fake.deleteOrphanedRoutesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteOrphanedRoutesReturns(result1 v7action.Warnings, result2 error) {
	fake.deleteOrphanedRoutesMutex.Lock()
	defer fake.deleteOrphanedRoutesMutex.Unlock()
	fake.DeleteOrphanedRoutesStub = nil
	fake.deleteOrphanedRoutesReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteOrphanedRoutesReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.deleteOrphanedRoutesMutex.Lock()
	defer fake.deleteOrphanedRoutesMutex.Unlock()
	fake.DeleteOrphanedRoutesStub = nil
	if fake.deleteOrphanedRoutesReturnsOnCall == nil {
		fake.deleteOrphanedRoutesReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.deleteOrphanedRoutesReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrphanedRoutesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteOrphanedRoutesMutex.RLock()
	defer fake.deleteOrphanedRoutesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteOrphanedRoutesActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.DeleteOrphanedRoutesActor = new(FakeDeleteOrphanedRoutesActor)
