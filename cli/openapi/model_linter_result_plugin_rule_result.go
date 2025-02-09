/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LinterResultPluginRuleResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinterResultPluginRuleResult{}

// LinterResultPluginRuleResult struct for LinterResultPluginRuleResult
type LinterResultPluginRuleResult struct {
	SpanId   *string  `json:"spanId,omitempty"`
	Errors   []string `json:"errors,omitempty"`
	Passed   *bool    `json:"passed,omitempty"`
	Severity *string  `json:"severity,omitempty"`
}

// NewLinterResultPluginRuleResult instantiates a new LinterResultPluginRuleResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinterResultPluginRuleResult() *LinterResultPluginRuleResult {
	this := LinterResultPluginRuleResult{}
	return &this
}

// NewLinterResultPluginRuleResultWithDefaults instantiates a new LinterResultPluginRuleResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinterResultPluginRuleResultWithDefaults() *LinterResultPluginRuleResult {
	this := LinterResultPluginRuleResult{}
	return &this
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *LinterResultPluginRuleResult) GetSpanId() string {
	if o == nil || isNil(o.SpanId) {
		var ret string
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResultPluginRuleResult) GetSpanIdOk() (*string, bool) {
	if o == nil || isNil(o.SpanId) {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *LinterResultPluginRuleResult) HasSpanId() bool {
	if o != nil && !isNil(o.SpanId) {
		return true
	}

	return false
}

// SetSpanId gets a reference to the given string and assigns it to the SpanId field.
func (o *LinterResultPluginRuleResult) SetSpanId(v string) {
	o.SpanId = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *LinterResultPluginRuleResult) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResultPluginRuleResult) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *LinterResultPluginRuleResult) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *LinterResultPluginRuleResult) SetErrors(v []string) {
	o.Errors = v
}

// GetPassed returns the Passed field value if set, zero value otherwise.
func (o *LinterResultPluginRuleResult) GetPassed() bool {
	if o == nil || isNil(o.Passed) {
		var ret bool
		return ret
	}
	return *o.Passed
}

// GetPassedOk returns a tuple with the Passed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResultPluginRuleResult) GetPassedOk() (*bool, bool) {
	if o == nil || isNil(o.Passed) {
		return nil, false
	}
	return o.Passed, true
}

// HasPassed returns a boolean if a field has been set.
func (o *LinterResultPluginRuleResult) HasPassed() bool {
	if o != nil && !isNil(o.Passed) {
		return true
	}

	return false
}

// SetPassed gets a reference to the given bool and assigns it to the Passed field.
func (o *LinterResultPluginRuleResult) SetPassed(v bool) {
	o.Passed = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *LinterResultPluginRuleResult) GetSeverity() string {
	if o == nil || isNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinterResultPluginRuleResult) GetSeverityOk() (*string, bool) {
	if o == nil || isNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *LinterResultPluginRuleResult) HasSeverity() bool {
	if o != nil && !isNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *LinterResultPluginRuleResult) SetSeverity(v string) {
	o.Severity = &v
}

func (o LinterResultPluginRuleResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinterResultPluginRuleResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SpanId) {
		toSerialize["spanId"] = o.SpanId
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.Passed) {
		toSerialize["passed"] = o.Passed
	}
	if !isNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	return toSerialize, nil
}

type NullableLinterResultPluginRuleResult struct {
	value *LinterResultPluginRuleResult
	isSet bool
}

func (v NullableLinterResultPluginRuleResult) Get() *LinterResultPluginRuleResult {
	return v.value
}

func (v *NullableLinterResultPluginRuleResult) Set(val *LinterResultPluginRuleResult) {
	v.value = val
	v.isSet = true
}

func (v NullableLinterResultPluginRuleResult) IsSet() bool {
	return v.isSet
}

func (v *NullableLinterResultPluginRuleResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinterResultPluginRuleResult(val *LinterResultPluginRuleResult) *NullableLinterResultPluginRuleResult {
	return &NullableLinterResultPluginRuleResult{value: val, isSet: true}
}

func (v NullableLinterResultPluginRuleResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinterResultPluginRuleResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
