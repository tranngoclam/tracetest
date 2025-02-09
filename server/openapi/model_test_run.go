/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type TestRun struct {
	Id string `json:"id,omitempty"`

	TraceId string `json:"traceId,omitempty"`

	SpanId string `json:"spanId,omitempty"`

	// Test version used when running this test run
	TestVersion int32 `json:"testVersion,omitempty"`

	// Current execution state
	State string `json:"state,omitempty"`

	// Details of the cause for the last `FAILED` state
	LastErrorState string `json:"lastErrorState,omitempty"`

	// time in seconds it took for the test to complete, either success or fail. If the test is still running, it will show the time up to the time of the request
	ExecutionTime int32 `json:"executionTime,omitempty"`

	// time in milliseconds it took for the triggering transaction to complete, either success or fail. If the test is still running, it will show the time up to the time of the request
	TriggerTime int32 `json:"triggerTime,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	ServiceTriggeredAt time.Time `json:"serviceTriggeredAt,omitempty"`

	ServiceTriggerCompletedAt time.Time `json:"serviceTriggerCompletedAt,omitempty"`

	ObtainedTraceAt time.Time `json:"obtainedTraceAt,omitempty"`

	CompletedAt time.Time `json:"completedAt,omitempty"`

	Environment Environment `json:"environment,omitempty"`

	TriggerResult TriggerResult `json:"triggerResult,omitempty"`

	Trace Trace `json:"trace,omitempty"`

	Result AssertionResults `json:"result,omitempty"`

	Linter LinterResult `json:"linter,omitempty"`

	Outputs []TestRunOutputsInner `json:"outputs,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`

	TransactionId string `json:"transactionId,omitempty"`

	TransactionRunId string `json:"transactionRunId,omitempty"`
}

// AssertTestRunRequired checks if the required fields are not zero-ed
func AssertTestRunRequired(obj TestRun) error {
	if err := AssertEnvironmentRequired(obj.Environment); err != nil {
		return err
	}
	if err := AssertTriggerResultRequired(obj.TriggerResult); err != nil {
		return err
	}
	if err := AssertTraceRequired(obj.Trace); err != nil {
		return err
	}
	if err := AssertAssertionResultsRequired(obj.Result); err != nil {
		return err
	}
	if err := AssertLinterResultRequired(obj.Linter); err != nil {
		return err
	}
	for _, el := range obj.Outputs {
		if err := AssertTestRunOutputsInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseTestRunRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TestRun (e.g. [][]TestRun), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTestRunRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTestRun, ok := obj.(TestRun)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTestRunRequired(aTestRun)
	})
}
