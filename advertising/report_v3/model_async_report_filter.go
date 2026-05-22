package report_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AsyncReportFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncReportFilter{}

// AsyncReportFilter struct for AsyncReportFilter
type AsyncReportFilter struct {
	// The field name of the filter
	Field *string `json:"field,omitempty"`
	// The values to be filtered by
	Values []string `json:"values,omitempty"`
}

// NewAsyncReportFilter instantiates a new AsyncReportFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncReportFilter() *AsyncReportFilter {
	this := AsyncReportFilter{}
	return &this
}

// NewAsyncReportFilterWithDefaults instantiates a new AsyncReportFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncReportFilterWithDefaults() *AsyncReportFilter {
	this := AsyncReportFilter{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *AsyncReportFilter) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncReportFilter) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *AsyncReportFilter) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *AsyncReportFilter) SetField(v string) {
	o.Field = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *AsyncReportFilter) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncReportFilter) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *AsyncReportFilter) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *AsyncReportFilter) SetValues(v []string) {
	o.Values = v
}

func (o AsyncReportFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableAsyncReportFilter struct {
	value *AsyncReportFilter
	isSet bool
}

func (v NullableAsyncReportFilter) Get() *AsyncReportFilter {
	return v.value
}

func (v *NullableAsyncReportFilter) Set(val *AsyncReportFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncReportFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncReportFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncReportFilter(val *AsyncReportFilter) *NullableAsyncReportFilter {
	return &NullableAsyncReportFilter{value: val, isSet: true}
}

func (v NullableAsyncReportFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsyncReportFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
