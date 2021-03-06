// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakePackagesActor struct {
	GetApplicationPackagesStub        func(string, string) ([]v7action.Package, v7action.Warnings, error)
	getApplicationPackagesMutex       sync.RWMutex
	getApplicationPackagesArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationPackagesReturns struct {
		result1 []v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	getApplicationPackagesReturnsOnCall map[int]struct {
		result1 []v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePackagesActor) GetApplicationPackages(arg1 string, arg2 string) ([]v7action.Package, v7action.Warnings, error) {
	fake.getApplicationPackagesMutex.Lock()
	ret, specificReturn := fake.getApplicationPackagesReturnsOnCall[len(fake.getApplicationPackagesArgsForCall)]
	fake.getApplicationPackagesArgsForCall = append(fake.getApplicationPackagesArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationPackages", []interface{}{arg1, arg2})
	fake.getApplicationPackagesMutex.Unlock()
	if fake.GetApplicationPackagesStub != nil {
		return fake.GetApplicationPackagesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationPackagesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakePackagesActor) GetApplicationPackagesCallCount() int {
	fake.getApplicationPackagesMutex.RLock()
	defer fake.getApplicationPackagesMutex.RUnlock()
	return len(fake.getApplicationPackagesArgsForCall)
}

func (fake *FakePackagesActor) GetApplicationPackagesCalls(stub func(string, string) ([]v7action.Package, v7action.Warnings, error)) {
	fake.getApplicationPackagesMutex.Lock()
	defer fake.getApplicationPackagesMutex.Unlock()
	fake.GetApplicationPackagesStub = stub
}

func (fake *FakePackagesActor) GetApplicationPackagesArgsForCall(i int) (string, string) {
	fake.getApplicationPackagesMutex.RLock()
	defer fake.getApplicationPackagesMutex.RUnlock()
	argsForCall := fake.getApplicationPackagesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePackagesActor) GetApplicationPackagesReturns(result1 []v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.getApplicationPackagesMutex.Lock()
	defer fake.getApplicationPackagesMutex.Unlock()
	fake.GetApplicationPackagesStub = nil
	fake.getApplicationPackagesReturns = struct {
		result1 []v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePackagesActor) GetApplicationPackagesReturnsOnCall(i int, result1 []v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.getApplicationPackagesMutex.Lock()
	defer fake.getApplicationPackagesMutex.Unlock()
	fake.GetApplicationPackagesStub = nil
	if fake.getApplicationPackagesReturnsOnCall == nil {
		fake.getApplicationPackagesReturnsOnCall = make(map[int]struct {
			result1 []v7action.Package
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationPackagesReturnsOnCall[i] = struct {
		result1 []v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePackagesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationPackagesMutex.RLock()
	defer fake.getApplicationPackagesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePackagesActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.PackagesActor = new(FakePackagesActor)
