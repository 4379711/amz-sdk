package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ArchiveLocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArchiveLocationRequest{}

// ArchiveLocationRequest Request body for the Archive Locations API
type ArchiveLocationRequest struct {
	LocationExpressionIdFilter *LocationExpressionIdFilter `json:"locationExpressionIdFilter,omitempty"`
}

// NewArchiveLocationRequest instantiates a new ArchiveLocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveLocationRequest() *ArchiveLocationRequest {
	this := ArchiveLocationRequest{}
	return &this
}

// NewArchiveLocationRequestWithDefaults instantiates a new ArchiveLocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveLocationRequestWithDefaults() *ArchiveLocationRequest {
	this := ArchiveLocationRequest{}
	return &this
}

// GetLocationExpressionIdFilter returns the LocationExpressionIdFilter field value if set, zero value otherwise.
func (o *ArchiveLocationRequest) GetLocationExpressionIdFilter() LocationExpressionIdFilter {
	if o == nil || IsNil(o.LocationExpressionIdFilter) {
		var ret LocationExpressionIdFilter
		return ret
	}
	return *o.LocationExpressionIdFilter
}

// GetLocationExpressionIdFilterOk returns a tuple with the LocationExpressionIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveLocationRequest) GetLocationExpressionIdFilterOk() (*LocationExpressionIdFilter, bool) {
	if o == nil || IsNil(o.LocationExpressionIdFilter) {
		return nil, false
	}
	return o.LocationExpressionIdFilter, true
}

// HasLocationExpressionIdFilter returns a boolean if a field has been set.
func (o *ArchiveLocationRequest) HasLocationExpressionIdFilter() bool {
	if o != nil && !IsNil(o.LocationExpressionIdFilter) {
		return true
	}

	return false
}

// SetLocationExpressionIdFilter gets a reference to the given LocationExpressionIdFilter and assigns it to the LocationExpressionIdFilter field.
func (o *ArchiveLocationRequest) SetLocationExpressionIdFilter(v LocationExpressionIdFilter) {
	o.LocationExpressionIdFilter = &v
}

func (o ArchiveLocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocationExpressionIdFilter) {
		toSerialize["locationExpressionIdFilter"] = o.LocationExpressionIdFilter
	}
	return toSerialize, nil
}

type NullableArchiveLocationRequest struct {
	value *ArchiveLocationRequest
	isSet bool
}

func (v NullableArchiveLocationRequest) Get() *ArchiveLocationRequest {
	return v.value
}

func (v *NullableArchiveLocationRequest) Set(val *ArchiveLocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveLocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveLocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveLocationRequest(val *ArchiveLocationRequest) *NullableArchiveLocationRequest {
	return &NullableArchiveLocationRequest{value: val, isSet: true}
}

func (v NullableArchiveLocationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableArchiveLocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
