// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package createflat

//go:generate minimock -i backend-bootcamp-assignment-2024/internal/services/renting/usecases/createflat.FlatsRepo -o flats_repo_mock_test.go -n FlatsRepoMock -p createflat

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// FlatsRepoMock implements FlatsRepo
type FlatsRepoMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateFlat          func(ctx context.Context, flat Request) (fp1 *models.Flat, err error)
	inspectFuncCreateFlat   func(ctx context.Context, flat Request)
	afterCreateFlatCounter  uint64
	beforeCreateFlatCounter uint64
	CreateFlatMock          mFlatsRepoMockCreateFlat
}

// NewFlatsRepoMock returns a mock for FlatsRepo
func NewFlatsRepoMock(t minimock.Tester) *FlatsRepoMock {
	m := &FlatsRepoMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateFlatMock = mFlatsRepoMockCreateFlat{mock: m}
	m.CreateFlatMock.callArgs = []*FlatsRepoMockCreateFlatParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mFlatsRepoMockCreateFlat struct {
	mock               *FlatsRepoMock
	defaultExpectation *FlatsRepoMockCreateFlatExpectation
	expectations       []*FlatsRepoMockCreateFlatExpectation

	callArgs []*FlatsRepoMockCreateFlatParams
	mutex    sync.RWMutex
}

// FlatsRepoMockCreateFlatExpectation specifies expectation struct of the FlatsRepo.CreateFlat
type FlatsRepoMockCreateFlatExpectation struct {
	mock    *FlatsRepoMock
	params  *FlatsRepoMockCreateFlatParams
	results *FlatsRepoMockCreateFlatResults
	Counter uint64
}

// FlatsRepoMockCreateFlatParams contains parameters of the FlatsRepo.CreateFlat
type FlatsRepoMockCreateFlatParams struct {
	ctx  context.Context
	flat Request
}

// FlatsRepoMockCreateFlatResults contains results of the FlatsRepo.CreateFlat
type FlatsRepoMockCreateFlatResults struct {
	fp1 *models.Flat
	err error
}

// Expect sets up expected params for FlatsRepo.CreateFlat
func (mmCreateFlat *mFlatsRepoMockCreateFlat) Expect(ctx context.Context, flat Request) *mFlatsRepoMockCreateFlat {
	if mmCreateFlat.mock.funcCreateFlat != nil {
		mmCreateFlat.mock.t.Fatalf("FlatsRepoMock.CreateFlat mock is already set by Set")
	}

	if mmCreateFlat.defaultExpectation == nil {
		mmCreateFlat.defaultExpectation = &FlatsRepoMockCreateFlatExpectation{}
	}

	mmCreateFlat.defaultExpectation.params = &FlatsRepoMockCreateFlatParams{ctx, flat}
	for _, e := range mmCreateFlat.expectations {
		if minimock.Equal(e.params, mmCreateFlat.defaultExpectation.params) {
			mmCreateFlat.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateFlat.defaultExpectation.params)
		}
	}

	return mmCreateFlat
}

// Inspect accepts an inspector function that has same arguments as the FlatsRepo.CreateFlat
func (mmCreateFlat *mFlatsRepoMockCreateFlat) Inspect(f func(ctx context.Context, flat Request)) *mFlatsRepoMockCreateFlat {
	if mmCreateFlat.mock.inspectFuncCreateFlat != nil {
		mmCreateFlat.mock.t.Fatalf("Inspect function is already set for FlatsRepoMock.CreateFlat")
	}

	mmCreateFlat.mock.inspectFuncCreateFlat = f

	return mmCreateFlat
}

// Return sets up results that will be returned by FlatsRepo.CreateFlat
func (mmCreateFlat *mFlatsRepoMockCreateFlat) Return(fp1 *models.Flat, err error) *FlatsRepoMock {
	if mmCreateFlat.mock.funcCreateFlat != nil {
		mmCreateFlat.mock.t.Fatalf("FlatsRepoMock.CreateFlat mock is already set by Set")
	}

	if mmCreateFlat.defaultExpectation == nil {
		mmCreateFlat.defaultExpectation = &FlatsRepoMockCreateFlatExpectation{mock: mmCreateFlat.mock}
	}
	mmCreateFlat.defaultExpectation.results = &FlatsRepoMockCreateFlatResults{fp1, err}
	return mmCreateFlat.mock
}

// Set uses given function f to mock the FlatsRepo.CreateFlat method
func (mmCreateFlat *mFlatsRepoMockCreateFlat) Set(f func(ctx context.Context, flat Request) (fp1 *models.Flat, err error)) *FlatsRepoMock {
	if mmCreateFlat.defaultExpectation != nil {
		mmCreateFlat.mock.t.Fatalf("Default expectation is already set for the FlatsRepo.CreateFlat method")
	}

	if len(mmCreateFlat.expectations) > 0 {
		mmCreateFlat.mock.t.Fatalf("Some expectations are already set for the FlatsRepo.CreateFlat method")
	}

	mmCreateFlat.mock.funcCreateFlat = f
	return mmCreateFlat.mock
}

// When sets expectation for the FlatsRepo.CreateFlat which will trigger the result defined by the following
// Then helper
func (mmCreateFlat *mFlatsRepoMockCreateFlat) When(ctx context.Context, flat Request) *FlatsRepoMockCreateFlatExpectation {
	if mmCreateFlat.mock.funcCreateFlat != nil {
		mmCreateFlat.mock.t.Fatalf("FlatsRepoMock.CreateFlat mock is already set by Set")
	}

	expectation := &FlatsRepoMockCreateFlatExpectation{
		mock:   mmCreateFlat.mock,
		params: &FlatsRepoMockCreateFlatParams{ctx, flat},
	}
	mmCreateFlat.expectations = append(mmCreateFlat.expectations, expectation)
	return expectation
}

// Then sets up FlatsRepo.CreateFlat return parameters for the expectation previously defined by the When method
func (e *FlatsRepoMockCreateFlatExpectation) Then(fp1 *models.Flat, err error) *FlatsRepoMock {
	e.results = &FlatsRepoMockCreateFlatResults{fp1, err}
	return e.mock
}

// CreateFlat implements FlatsRepo
func (mmCreateFlat *FlatsRepoMock) CreateFlat(ctx context.Context, flat Request) (fp1 *models.Flat, err error) {
	mm_atomic.AddUint64(&mmCreateFlat.beforeCreateFlatCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateFlat.afterCreateFlatCounter, 1)

	if mmCreateFlat.inspectFuncCreateFlat != nil {
		mmCreateFlat.inspectFuncCreateFlat(ctx, flat)
	}

	mm_params := FlatsRepoMockCreateFlatParams{ctx, flat}

	// Record call args
	mmCreateFlat.CreateFlatMock.mutex.Lock()
	mmCreateFlat.CreateFlatMock.callArgs = append(mmCreateFlat.CreateFlatMock.callArgs, &mm_params)
	mmCreateFlat.CreateFlatMock.mutex.Unlock()

	for _, e := range mmCreateFlat.CreateFlatMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.fp1, e.results.err
		}
	}

	if mmCreateFlat.CreateFlatMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateFlat.CreateFlatMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateFlat.CreateFlatMock.defaultExpectation.params
		mm_got := FlatsRepoMockCreateFlatParams{ctx, flat}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateFlat.t.Errorf("FlatsRepoMock.CreateFlat got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateFlat.CreateFlatMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateFlat.t.Fatal("No results are set for the FlatsRepoMock.CreateFlat")
		}
		return (*mm_results).fp1, (*mm_results).err
	}
	if mmCreateFlat.funcCreateFlat != nil {
		return mmCreateFlat.funcCreateFlat(ctx, flat)
	}
	mmCreateFlat.t.Fatalf("Unexpected call to FlatsRepoMock.CreateFlat. %v %v", ctx, flat)
	return
}

// CreateFlatAfterCounter returns a count of finished FlatsRepoMock.CreateFlat invocations
func (mmCreateFlat *FlatsRepoMock) CreateFlatAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateFlat.afterCreateFlatCounter)
}

// CreateFlatBeforeCounter returns a count of FlatsRepoMock.CreateFlat invocations
func (mmCreateFlat *FlatsRepoMock) CreateFlatBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateFlat.beforeCreateFlatCounter)
}

// Calls returns a list of arguments used in each call to FlatsRepoMock.CreateFlat.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateFlat *mFlatsRepoMockCreateFlat) Calls() []*FlatsRepoMockCreateFlatParams {
	mmCreateFlat.mutex.RLock()

	argCopy := make([]*FlatsRepoMockCreateFlatParams, len(mmCreateFlat.callArgs))
	copy(argCopy, mmCreateFlat.callArgs)

	mmCreateFlat.mutex.RUnlock()

	return argCopy
}

// MinimockCreateFlatDone returns true if the count of the CreateFlat invocations corresponds
// the number of defined expectations
func (m *FlatsRepoMock) MinimockCreateFlatDone() bool {
	for _, e := range m.CreateFlatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateFlatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateFlatCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateFlat != nil && mm_atomic.LoadUint64(&m.afterCreateFlatCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateFlatInspect logs each unmet expectation
func (m *FlatsRepoMock) MinimockCreateFlatInspect() {
	for _, e := range m.CreateFlatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FlatsRepoMock.CreateFlat with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateFlatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateFlatCounter) < 1 {
		if m.CreateFlatMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FlatsRepoMock.CreateFlat")
		} else {
			m.t.Errorf("Expected call to FlatsRepoMock.CreateFlat with params: %#v", *m.CreateFlatMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateFlat != nil && mm_atomic.LoadUint64(&m.afterCreateFlatCounter) < 1 {
		m.t.Error("Expected call to FlatsRepoMock.CreateFlat")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *FlatsRepoMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateFlatInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *FlatsRepoMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *FlatsRepoMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateFlatDone()
}
