// Code generated by counterfeiter. DO NOT EDIT.
package stub

import (
	"sync"

	"go.bobheadxi.dev/zapx/zmetrics/metrics"
)

type StubGauge struct {
	SetStub        func(float64)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 float64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StubGauge) Set(arg1 float64) {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 float64
	}{arg1})
	fake.recordInvocation("Set", []interface{}{arg1})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		fake.SetStub(arg1)
	}
}

func (fake *StubGauge) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *StubGauge) SetCalls(stub func(float64)) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = stub
}

func (fake *StubGauge) SetArgsForCall(i int) float64 {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	argsForCall := fake.setArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StubGauge) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StubGauge) recordInvocation(key string, args []interface{}) {
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

var _ metrics.Gauge = new(StubGauge)
