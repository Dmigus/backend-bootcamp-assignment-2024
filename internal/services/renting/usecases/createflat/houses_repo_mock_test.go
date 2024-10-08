// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package createflat

//go:generate minimock -i backend-bootcamp-assignment-2024/internal/services/renting/usecases/createflat.HousesRepo -o houses_repo_mock_test.go -n HousesRepoMock -p createflat

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// HousesRepoMock implements HousesRepo
type HousesRepoMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcHouseUpdated          func(ctx context.Context, houseId int) (err error)
	inspectFuncHouseUpdated   func(ctx context.Context, houseId int)
	afterHouseUpdatedCounter  uint64
	beforeHouseUpdatedCounter uint64
	HouseUpdatedMock          mHousesRepoMockHouseUpdated
}

// NewHousesRepoMock returns a mock for HousesRepo
func NewHousesRepoMock(t minimock.Tester) *HousesRepoMock {
	m := &HousesRepoMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.HouseUpdatedMock = mHousesRepoMockHouseUpdated{mock: m}
	m.HouseUpdatedMock.callArgs = []*HousesRepoMockHouseUpdatedParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mHousesRepoMockHouseUpdated struct {
	mock               *HousesRepoMock
	defaultExpectation *HousesRepoMockHouseUpdatedExpectation
	expectations       []*HousesRepoMockHouseUpdatedExpectation

	callArgs []*HousesRepoMockHouseUpdatedParams
	mutex    sync.RWMutex
}

// HousesRepoMockHouseUpdatedExpectation specifies expectation struct of the HousesRepo.HouseUpdated
type HousesRepoMockHouseUpdatedExpectation struct {
	mock    *HousesRepoMock
	params  *HousesRepoMockHouseUpdatedParams
	results *HousesRepoMockHouseUpdatedResults
	Counter uint64
}

// HousesRepoMockHouseUpdatedParams contains parameters of the HousesRepo.HouseUpdated
type HousesRepoMockHouseUpdatedParams struct {
	ctx     context.Context
	houseId int
}

// HousesRepoMockHouseUpdatedResults contains results of the HousesRepo.HouseUpdated
type HousesRepoMockHouseUpdatedResults struct {
	err error
}

// Expect sets up expected params for HousesRepo.HouseUpdated
func (mmHouseUpdated *mHousesRepoMockHouseUpdated) Expect(ctx context.Context, houseId int) *mHousesRepoMockHouseUpdated {
	if mmHouseUpdated.mock.funcHouseUpdated != nil {
		mmHouseUpdated.mock.t.Fatalf("HousesRepoMock.HouseUpdated mock is already set by Set")
	}

	if mmHouseUpdated.defaultExpectation == nil {
		mmHouseUpdated.defaultExpectation = &HousesRepoMockHouseUpdatedExpectation{}
	}

	mmHouseUpdated.defaultExpectation.params = &HousesRepoMockHouseUpdatedParams{ctx, houseId}
	for _, e := range mmHouseUpdated.expectations {
		if minimock.Equal(e.params, mmHouseUpdated.defaultExpectation.params) {
			mmHouseUpdated.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmHouseUpdated.defaultExpectation.params)
		}
	}

	return mmHouseUpdated
}

// Inspect accepts an inspector function that has same arguments as the HousesRepo.HouseUpdated
func (mmHouseUpdated *mHousesRepoMockHouseUpdated) Inspect(f func(ctx context.Context, houseId int)) *mHousesRepoMockHouseUpdated {
	if mmHouseUpdated.mock.inspectFuncHouseUpdated != nil {
		mmHouseUpdated.mock.t.Fatalf("Inspect function is already set for HousesRepoMock.HouseUpdated")
	}

	mmHouseUpdated.mock.inspectFuncHouseUpdated = f

	return mmHouseUpdated
}

// Return sets up results that will be returned by HousesRepo.HouseUpdated
func (mmHouseUpdated *mHousesRepoMockHouseUpdated) Return(err error) *HousesRepoMock {
	if mmHouseUpdated.mock.funcHouseUpdated != nil {
		mmHouseUpdated.mock.t.Fatalf("HousesRepoMock.HouseUpdated mock is already set by Set")
	}

	if mmHouseUpdated.defaultExpectation == nil {
		mmHouseUpdated.defaultExpectation = &HousesRepoMockHouseUpdatedExpectation{mock: mmHouseUpdated.mock}
	}
	mmHouseUpdated.defaultExpectation.results = &HousesRepoMockHouseUpdatedResults{err}
	return mmHouseUpdated.mock
}

// Set uses given function f to mock the HousesRepo.HouseUpdated method
func (mmHouseUpdated *mHousesRepoMockHouseUpdated) Set(f func(ctx context.Context, houseId int) (err error)) *HousesRepoMock {
	if mmHouseUpdated.defaultExpectation != nil {
		mmHouseUpdated.mock.t.Fatalf("Default expectation is already set for the HousesRepo.HouseUpdated method")
	}

	if len(mmHouseUpdated.expectations) > 0 {
		mmHouseUpdated.mock.t.Fatalf("Some expectations are already set for the HousesRepo.HouseUpdated method")
	}

	mmHouseUpdated.mock.funcHouseUpdated = f
	return mmHouseUpdated.mock
}

// When sets expectation for the HousesRepo.HouseUpdated which will trigger the result defined by the following
// Then helper
func (mmHouseUpdated *mHousesRepoMockHouseUpdated) When(ctx context.Context, houseId int) *HousesRepoMockHouseUpdatedExpectation {
	if mmHouseUpdated.mock.funcHouseUpdated != nil {
		mmHouseUpdated.mock.t.Fatalf("HousesRepoMock.HouseUpdated mock is already set by Set")
	}

	expectation := &HousesRepoMockHouseUpdatedExpectation{
		mock:   mmHouseUpdated.mock,
		params: &HousesRepoMockHouseUpdatedParams{ctx, houseId},
	}
	mmHouseUpdated.expectations = append(mmHouseUpdated.expectations, expectation)
	return expectation
}

// Then sets up HousesRepo.HouseUpdated return parameters for the expectation previously defined by the When method
func (e *HousesRepoMockHouseUpdatedExpectation) Then(err error) *HousesRepoMock {
	e.results = &HousesRepoMockHouseUpdatedResults{err}
	return e.mock
}

// HouseUpdated implements HousesRepo
func (mmHouseUpdated *HousesRepoMock) HouseUpdated(ctx context.Context, houseId int) (err error) {
	mm_atomic.AddUint64(&mmHouseUpdated.beforeHouseUpdatedCounter, 1)
	defer mm_atomic.AddUint64(&mmHouseUpdated.afterHouseUpdatedCounter, 1)

	if mmHouseUpdated.inspectFuncHouseUpdated != nil {
		mmHouseUpdated.inspectFuncHouseUpdated(ctx, houseId)
	}

	mm_params := HousesRepoMockHouseUpdatedParams{ctx, houseId}

	// Record call args
	mmHouseUpdated.HouseUpdatedMock.mutex.Lock()
	mmHouseUpdated.HouseUpdatedMock.callArgs = append(mmHouseUpdated.HouseUpdatedMock.callArgs, &mm_params)
	mmHouseUpdated.HouseUpdatedMock.mutex.Unlock()

	for _, e := range mmHouseUpdated.HouseUpdatedMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmHouseUpdated.HouseUpdatedMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmHouseUpdated.HouseUpdatedMock.defaultExpectation.Counter, 1)
		mm_want := mmHouseUpdated.HouseUpdatedMock.defaultExpectation.params
		mm_got := HousesRepoMockHouseUpdatedParams{ctx, houseId}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmHouseUpdated.t.Errorf("HousesRepoMock.HouseUpdated got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmHouseUpdated.HouseUpdatedMock.defaultExpectation.results
		if mm_results == nil {
			mmHouseUpdated.t.Fatal("No results are set for the HousesRepoMock.HouseUpdated")
		}
		return (*mm_results).err
	}
	if mmHouseUpdated.funcHouseUpdated != nil {
		return mmHouseUpdated.funcHouseUpdated(ctx, houseId)
	}
	mmHouseUpdated.t.Fatalf("Unexpected call to HousesRepoMock.HouseUpdated. %v %v", ctx, houseId)
	return
}

// HouseUpdatedAfterCounter returns a count of finished HousesRepoMock.HouseUpdated invocations
func (mmHouseUpdated *HousesRepoMock) HouseUpdatedAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmHouseUpdated.afterHouseUpdatedCounter)
}

// HouseUpdatedBeforeCounter returns a count of HousesRepoMock.HouseUpdated invocations
func (mmHouseUpdated *HousesRepoMock) HouseUpdatedBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmHouseUpdated.beforeHouseUpdatedCounter)
}

// Calls returns a list of arguments used in each call to HousesRepoMock.HouseUpdated.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmHouseUpdated *mHousesRepoMockHouseUpdated) Calls() []*HousesRepoMockHouseUpdatedParams {
	mmHouseUpdated.mutex.RLock()

	argCopy := make([]*HousesRepoMockHouseUpdatedParams, len(mmHouseUpdated.callArgs))
	copy(argCopy, mmHouseUpdated.callArgs)

	mmHouseUpdated.mutex.RUnlock()

	return argCopy
}

// MinimockHouseUpdatedDone returns true if the count of the HouseUpdated invocations corresponds
// the number of defined expectations
func (m *HousesRepoMock) MinimockHouseUpdatedDone() bool {
	for _, e := range m.HouseUpdatedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.HouseUpdatedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterHouseUpdatedCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcHouseUpdated != nil && mm_atomic.LoadUint64(&m.afterHouseUpdatedCounter) < 1 {
		return false
	}
	return true
}

// MinimockHouseUpdatedInspect logs each unmet expectation
func (m *HousesRepoMock) MinimockHouseUpdatedInspect() {
	for _, e := range m.HouseUpdatedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to HousesRepoMock.HouseUpdated with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.HouseUpdatedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterHouseUpdatedCounter) < 1 {
		if m.HouseUpdatedMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to HousesRepoMock.HouseUpdated")
		} else {
			m.t.Errorf("Expected call to HousesRepoMock.HouseUpdated with params: %#v", *m.HouseUpdatedMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcHouseUpdated != nil && mm_atomic.LoadUint64(&m.afterHouseUpdatedCounter) < 1 {
		m.t.Error("Expected call to HousesRepoMock.HouseUpdated")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *HousesRepoMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockHouseUpdatedInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *HousesRepoMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *HousesRepoMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockHouseUpdatedDone()
}
