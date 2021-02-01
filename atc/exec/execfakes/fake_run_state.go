// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/exec"
	"github.com/concourse/concourse/atc/exec/build"
	"github.com/concourse/concourse/vars"
)

type FakeRunState struct {
	ArtifactRepositoryStub        func() *build.Repository
	artifactRepositoryMutex       sync.RWMutex
	artifactRepositoryArgsForCall []struct {
	}
	artifactRepositoryReturns struct {
		result1 *build.Repository
	}
	artifactRepositoryReturnsOnCall map[int]struct {
		result1 *build.Repository
	}
	IterateInterpolatedCredsStub        func(vars.TrackedVarsIterator)
	iterateInterpolatedCredsMutex       sync.RWMutex
	iterateInterpolatedCredsArgsForCall []struct {
		arg1 vars.TrackedVarsIterator
	}
	NewScopeStub        func() exec.RunState
	newScopeMutex       sync.RWMutex
	newScopeArgsForCall []struct {
	}
	newScopeReturns struct {
		result1 exec.RunState
	}
	newScopeReturnsOnCall map[int]struct {
		result1 exec.RunState
	}
	ParentStub        func() exec.RunState
	parentMutex       sync.RWMutex
	parentArgsForCall []struct {
	}
	parentReturns struct {
		result1 exec.RunState
	}
	parentReturnsOnCall map[int]struct {
		result1 exec.RunState
	}
	RedactionEnabledStub        func() bool
	redactionEnabledMutex       sync.RWMutex
	redactionEnabledArgsForCall []struct {
	}
	redactionEnabledReturns struct {
		result1 bool
	}
	redactionEnabledReturnsOnCall map[int]struct {
		result1 bool
	}
	ResultStub        func(atc.PlanID, interface{}) bool
	resultMutex       sync.RWMutex
	resultArgsForCall []struct {
		arg1 atc.PlanID
		arg2 interface{}
	}
	resultReturns struct {
		result1 bool
	}
	resultReturnsOnCall map[int]struct {
		result1 bool
	}
	RunStub        func(context.Context, atc.Plan) (bool, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 context.Context
		arg2 atc.Plan
	}
	runReturns struct {
		result1 bool
		result2 error
	}
	runReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	SetVarSourceConfigsStub        func(atc.VarSourceConfigs)
	setVarSourceConfigsMutex       sync.RWMutex
	setVarSourceConfigsArgsForCall []struct {
		arg1 atc.VarSourceConfigs
	}
	StoreResultStub        func(atc.PlanID, interface{})
	storeResultMutex       sync.RWMutex
	storeResultArgsForCall []struct {
		arg1 atc.PlanID
		arg2 interface{}
	}
	VariablesStub        func() *build.Variables
	variablesMutex       sync.RWMutex
	variablesArgsForCall []struct {
	}
	VarSourceConfigsStub        func() atc.VarSourceConfigs
	varSourceConfigsMutex       sync.RWMutex
	varSourceConfigsArgsForCall []struct {
	}
	varSourceConfigsReturns struct {
		result1 atc.VarSourceConfigs
	}
	varSourceConfigsReturnsOnCall map[int]struct {
		result1 atc.VarSourceConfigs
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRunState) ArtifactRepository() *build.Repository {
	fake.artifactRepositoryMutex.Lock()
	ret, specificReturn := fake.artifactRepositoryReturnsOnCall[len(fake.artifactRepositoryArgsForCall)]
	fake.artifactRepositoryArgsForCall = append(fake.artifactRepositoryArgsForCall, struct {
	}{})
	stub := fake.ArtifactRepositoryStub
	fakeReturns := fake.artifactRepositoryReturns
	fake.recordInvocation("ArtifactRepository", []interface{}{})
	fake.artifactRepositoryMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunState) ArtifactRepositoryCallCount() int {
	fake.artifactRepositoryMutex.RLock()
	defer fake.artifactRepositoryMutex.RUnlock()
	return len(fake.artifactRepositoryArgsForCall)
}

func (fake *FakeRunState) ArtifactRepositoryCalls(stub func() *build.Repository) {
	fake.artifactRepositoryMutex.Lock()
	defer fake.artifactRepositoryMutex.Unlock()
	fake.ArtifactRepositoryStub = stub
}

func (fake *FakeRunState) ArtifactRepositoryReturns(result1 *build.Repository) {
	fake.artifactRepositoryMutex.Lock()
	defer fake.artifactRepositoryMutex.Unlock()
	fake.ArtifactRepositoryStub = nil
	fake.artifactRepositoryReturns = struct {
		result1 *build.Repository
	}{result1}
}

func (fake *FakeRunState) ArtifactRepositoryReturnsOnCall(i int, result1 *build.Repository) {
	fake.artifactRepositoryMutex.Lock()
	defer fake.artifactRepositoryMutex.Unlock()
	fake.ArtifactRepositoryStub = nil
	if fake.artifactRepositoryReturnsOnCall == nil {
		fake.artifactRepositoryReturnsOnCall = make(map[int]struct {
			result1 *build.Repository
		})
	}
	fake.artifactRepositoryReturnsOnCall[i] = struct {
		result1 *build.Repository
	}{result1}
}

func (fake *FakeRunState) IterateInterpolatedCreds(arg1 vars.TrackedVarsIterator) {
	fake.iterateInterpolatedCredsMutex.Lock()
	fake.iterateInterpolatedCredsArgsForCall = append(fake.iterateInterpolatedCredsArgsForCall, struct {
		arg1 vars.TrackedVarsIterator
	}{arg1})
	stub := fake.IterateInterpolatedCredsStub
	fake.recordInvocation("IterateInterpolatedCreds", []interface{}{arg1})
	fake.iterateInterpolatedCredsMutex.Unlock()
	if stub != nil {
		fake.IterateInterpolatedCredsStub(arg1)
	}
}

func (fake *FakeRunState) IterateInterpolatedCredsCallCount() int {
	fake.iterateInterpolatedCredsMutex.RLock()
	defer fake.iterateInterpolatedCredsMutex.RUnlock()
	return len(fake.iterateInterpolatedCredsArgsForCall)
}

func (fake *FakeRunState) IterateInterpolatedCredsCalls(stub func(vars.TrackedVarsIterator)) {
	fake.iterateInterpolatedCredsMutex.Lock()
	defer fake.iterateInterpolatedCredsMutex.Unlock()
	fake.IterateInterpolatedCredsStub = stub
}

func (fake *FakeRunState) IterateInterpolatedCredsArgsForCall(i int) vars.TrackedVarsIterator {
	fake.iterateInterpolatedCredsMutex.RLock()
	defer fake.iterateInterpolatedCredsMutex.RUnlock()
	argsForCall := fake.iterateInterpolatedCredsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRunState) NewScope() exec.RunState {
	fake.newScopeMutex.Lock()
	ret, specificReturn := fake.newScopeReturnsOnCall[len(fake.newScopeArgsForCall)]
	fake.newScopeArgsForCall = append(fake.newScopeArgsForCall, struct {
	}{})
	stub := fake.NewScopeStub
	fakeReturns := fake.newScopeReturns
	fake.recordInvocation("NewScope", []interface{}{})
	fake.newScopeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunState) NewScopeCallCount() int {
	fake.newScopeMutex.RLock()
	defer fake.newScopeMutex.RUnlock()
	return len(fake.newScopeArgsForCall)
}

func (fake *FakeRunState) NewScopeCalls(stub func() exec.RunState) {
	fake.newScopeMutex.Lock()
	defer fake.newScopeMutex.Unlock()
	fake.NewScopeStub = stub
}

func (fake *FakeRunState) NewScopeReturns(result1 exec.RunState) {
	fake.newScopeMutex.Lock()
	defer fake.newScopeMutex.Unlock()
	fake.NewScopeStub = nil
	fake.newScopeReturns = struct {
		result1 exec.RunState
	}{result1}
}

func (fake *FakeRunState) NewScopeReturnsOnCall(i int, result1 exec.RunState) {
	fake.newScopeMutex.Lock()
	defer fake.newScopeMutex.Unlock()
	fake.NewScopeStub = nil
	if fake.newScopeReturnsOnCall == nil {
		fake.newScopeReturnsOnCall = make(map[int]struct {
			result1 exec.RunState
		})
	}
	fake.newScopeReturnsOnCall[i] = struct {
		result1 exec.RunState
	}{result1}
}

func (fake *FakeRunState) Parent() exec.RunState {
	fake.parentMutex.Lock()
	ret, specificReturn := fake.parentReturnsOnCall[len(fake.parentArgsForCall)]
	fake.parentArgsForCall = append(fake.parentArgsForCall, struct {
	}{})
	stub := fake.ParentStub
	fakeReturns := fake.parentReturns
	fake.recordInvocation("Parent", []interface{}{})
	fake.parentMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunState) ParentCallCount() int {
	fake.parentMutex.RLock()
	defer fake.parentMutex.RUnlock()
	return len(fake.parentArgsForCall)
}

func (fake *FakeRunState) ParentCalls(stub func() exec.RunState) {
	fake.parentMutex.Lock()
	defer fake.parentMutex.Unlock()
	fake.ParentStub = stub
}

func (fake *FakeRunState) ParentReturns(result1 exec.RunState) {
	fake.parentMutex.Lock()
	defer fake.parentMutex.Unlock()
	fake.ParentStub = nil
	fake.parentReturns = struct {
		result1 exec.RunState
	}{result1}
}

func (fake *FakeRunState) ParentReturnsOnCall(i int, result1 exec.RunState) {
	fake.parentMutex.Lock()
	defer fake.parentMutex.Unlock()
	fake.ParentStub = nil
	if fake.parentReturnsOnCall == nil {
		fake.parentReturnsOnCall = make(map[int]struct {
			result1 exec.RunState
		})
	}
	fake.parentReturnsOnCall[i] = struct {
		result1 exec.RunState
	}{result1}
}

func (fake *FakeRunState) RedactionEnabled() bool {
	fake.redactionEnabledMutex.Lock()
	ret, specificReturn := fake.redactionEnabledReturnsOnCall[len(fake.redactionEnabledArgsForCall)]
	fake.redactionEnabledArgsForCall = append(fake.redactionEnabledArgsForCall, struct {
	}{})
	stub := fake.RedactionEnabledStub
	fakeReturns := fake.redactionEnabledReturns
	fake.recordInvocation("RedactionEnabled", []interface{}{})
	fake.redactionEnabledMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunState) RedactionEnabledCallCount() int {
	fake.redactionEnabledMutex.RLock()
	defer fake.redactionEnabledMutex.RUnlock()
	return len(fake.redactionEnabledArgsForCall)
}

func (fake *FakeRunState) RedactionEnabledCalls(stub func() bool) {
	fake.redactionEnabledMutex.Lock()
	defer fake.redactionEnabledMutex.Unlock()
	fake.RedactionEnabledStub = stub
}

func (fake *FakeRunState) RedactionEnabledReturns(result1 bool) {
	fake.redactionEnabledMutex.Lock()
	defer fake.redactionEnabledMutex.Unlock()
	fake.RedactionEnabledStub = nil
	fake.redactionEnabledReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRunState) RedactionEnabledReturnsOnCall(i int, result1 bool) {
	fake.redactionEnabledMutex.Lock()
	defer fake.redactionEnabledMutex.Unlock()
	fake.RedactionEnabledStub = nil
	if fake.redactionEnabledReturnsOnCall == nil {
		fake.redactionEnabledReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.redactionEnabledReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRunState) Result(arg1 atc.PlanID, arg2 interface{}) bool {
	fake.resultMutex.Lock()
	ret, specificReturn := fake.resultReturnsOnCall[len(fake.resultArgsForCall)]
	fake.resultArgsForCall = append(fake.resultArgsForCall, struct {
		arg1 atc.PlanID
		arg2 interface{}
	}{arg1, arg2})
	stub := fake.ResultStub
	fakeReturns := fake.resultReturns
	fake.recordInvocation("Result", []interface{}{arg1, arg2})
	fake.resultMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunState) ResultCallCount() int {
	fake.resultMutex.RLock()
	defer fake.resultMutex.RUnlock()
	return len(fake.resultArgsForCall)
}

func (fake *FakeRunState) ResultCalls(stub func(atc.PlanID, interface{}) bool) {
	fake.resultMutex.Lock()
	defer fake.resultMutex.Unlock()
	fake.ResultStub = stub
}

func (fake *FakeRunState) ResultArgsForCall(i int) (atc.PlanID, interface{}) {
	fake.resultMutex.RLock()
	defer fake.resultMutex.RUnlock()
	argsForCall := fake.resultArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRunState) ResultReturns(result1 bool) {
	fake.resultMutex.Lock()
	defer fake.resultMutex.Unlock()
	fake.ResultStub = nil
	fake.resultReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRunState) ResultReturnsOnCall(i int, result1 bool) {
	fake.resultMutex.Lock()
	defer fake.resultMutex.Unlock()
	fake.ResultStub = nil
	if fake.resultReturnsOnCall == nil {
		fake.resultReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.resultReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRunState) Run(arg1 context.Context, arg2 atc.Plan) (bool, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 context.Context
		arg2 atc.Plan
	}{arg1, arg2})
	stub := fake.RunStub
	fakeReturns := fake.runReturns
	fake.recordInvocation("Run", []interface{}{arg1, arg2})
	fake.runMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRunState) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeRunState) RunCalls(stub func(context.Context, atc.Plan) (bool, error)) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeRunState) RunArgsForCall(i int) (context.Context, atc.Plan) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRunState) RunReturns(result1 bool, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRunState) RunReturnsOnCall(i int, result1 bool, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRunState) SetVarSourceConfigs(arg1 atc.VarSourceConfigs) {
	fake.setVarSourceConfigsMutex.Lock()
	fake.setVarSourceConfigsArgsForCall = append(fake.setVarSourceConfigsArgsForCall, struct {
		arg1 atc.VarSourceConfigs
	}{arg1})
	stub := fake.SetVarSourceConfigsStub
	fake.recordInvocation("SetVarSourceConfigs", []interface{}{arg1})
	fake.setVarSourceConfigsMutex.Unlock()
	if stub != nil {
		fake.SetVarSourceConfigsStub(arg1)
	}
}

func (fake *FakeRunState) SetVarSourceConfigsCallCount() int {
	fake.setVarSourceConfigsMutex.RLock()
	defer fake.setVarSourceConfigsMutex.RUnlock()
	return len(fake.setVarSourceConfigsArgsForCall)
}

func (fake *FakeRunState) SetVarSourceConfigsCalls(stub func(atc.VarSourceConfigs)) {
	fake.setVarSourceConfigsMutex.Lock()
	defer fake.setVarSourceConfigsMutex.Unlock()
	fake.SetVarSourceConfigsStub = stub
}

func (fake *FakeRunState) SetVarSourceConfigsArgsForCall(i int) atc.VarSourceConfigs {
	fake.setVarSourceConfigsMutex.RLock()
	defer fake.setVarSourceConfigsMutex.RUnlock()
	argsForCall := fake.setVarSourceConfigsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRunState) StoreResult(arg1 atc.PlanID, arg2 interface{}) {
	fake.storeResultMutex.Lock()
	fake.storeResultArgsForCall = append(fake.storeResultArgsForCall, struct {
		arg1 atc.PlanID
		arg2 interface{}
	}{arg1, arg2})
	stub := fake.StoreResultStub
	fake.recordInvocation("StoreResult", []interface{}{arg1, arg2})
	fake.storeResultMutex.Unlock()
	if stub != nil {
		fake.StoreResultStub(arg1, arg2)
	}
}

func (fake *FakeRunState) StoreResultCallCount() int {
	fake.storeResultMutex.RLock()
	defer fake.storeResultMutex.RUnlock()
	return len(fake.storeResultArgsForCall)
}

func (fake *FakeRunState) StoreResultCalls(stub func(atc.PlanID, interface{})) {
	fake.storeResultMutex.Lock()
	defer fake.storeResultMutex.Unlock()
	fake.StoreResultStub = stub
}

func (fake *FakeRunState) StoreResultArgsForCall(i int) (atc.PlanID, interface{}) {
	fake.storeResultMutex.RLock()
	defer fake.storeResultMutex.RUnlock()
	argsForCall := fake.storeResultArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRunState) Variables() *build.Variables {
	fake.variablesMutex.Lock()
	ret, specificReturn := fake.variablesReturnsOnCall[len(fake.variablesArgsForCall)]
	fake.variablesArgsForCall = append(fake.variablesArgsForCall, struct {
	}{})
	stub := fake.VarSourceConfigsStub
	fakeReturns := fake.varSourceConfigsReturns
	fake.recordInvocation("VarSourceConfigs", []interface{}{})
	fake.varSourceConfigsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunState) VarSourceConfigsCallCount() int {
	fake.varSourceConfigsMutex.RLock()
	defer fake.varSourceConfigsMutex.RUnlock()
	return len(fake.varSourceConfigsArgsForCall)
}

func (fake *FakeRunState) VarSourceConfigsCalls(stub func() atc.VarSourceConfigs) {
	fake.varSourceConfigsMutex.Lock()
	defer fake.varSourceConfigsMutex.Unlock()
	fake.VarSourceConfigsStub = stub
}

func (fake *FakeRunState) VarSourceConfigsReturns(result1 atc.VarSourceConfigs) {
	fake.varSourceConfigsMutex.Lock()
	defer fake.varSourceConfigsMutex.Unlock()
	fake.VarSourceConfigsStub = nil
	fake.varSourceConfigsReturns = struct {
		result1 atc.VarSourceConfigs
	}{result1}
}

func (fake *FakeRunState) VarSourceConfigsReturnsOnCall(i int, result1 atc.VarSourceConfigs) {
	fake.varSourceConfigsMutex.Lock()
	defer fake.varSourceConfigsMutex.Unlock()
	fake.VarSourceConfigsStub = nil
	if fake.varSourceConfigsReturnsOnCall == nil {
		fake.varSourceConfigsReturnsOnCall = make(map[int]struct {
			result1 atc.VarSourceConfigs
		})
	}
	fake.varSourceConfigsReturnsOnCall[i] = struct {
		result1 atc.VarSourceConfigs
	}{result1}
}

func (fake *FakeRunState) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.artifactRepositoryMutex.RLock()
	defer fake.artifactRepositoryMutex.RUnlock()
	fake.iterateInterpolatedCredsMutex.RLock()
	defer fake.iterateInterpolatedCredsMutex.RUnlock()
	fake.newScopeMutex.RLock()
	defer fake.newScopeMutex.RUnlock()
	fake.parentMutex.RLock()
	defer fake.parentMutex.RUnlock()
	fake.redactionEnabledMutex.RLock()
	defer fake.redactionEnabledMutex.RUnlock()
	fake.resultMutex.RLock()
	defer fake.resultMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.setVarSourceConfigsMutex.RLock()
	defer fake.setVarSourceConfigsMutex.RUnlock()
	fake.storeResultMutex.RLock()
	defer fake.storeResultMutex.RUnlock()
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRunState) recordInvocation(key string, args []interface{}) {
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

var _ exec.RunState = new(FakeRunState)
