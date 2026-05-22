package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingErrorObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingErrorObject{}

// SBForecastingErrorObject struct for SBForecastingErrorObject
type SBForecastingErrorObject struct {
	// Correlates the campaign to the campaign list index specified in the request. Zero-based.
	Index *int32 `json:"index,omitempty"`
	// The forecast error code.
	Code *string `json:"code,omitempty"`
	// The forecast error description.
	Description *string `json:"description,omitempty"`
}

// NewSBForecastingErrorObject instantiates a new SBForecastingErrorObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingErrorObject() *SBForecastingErrorObject {
	this := SBForecastingErrorObject{}
	return &this
}

// NewSBForecastingErrorObjectWithDefaults instantiates a new SBForecastingErrorObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingErrorObjectWithDefaults() *SBForecastingErrorObject {
	this := SBForecastingErrorObject{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SBForecastingErrorObject) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingErrorObject) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SBForecastingErrorObject) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *SBForecastingErrorObject) SetIndex(v int32) {
	o.Index = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SBForecastingErrorObject) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingErrorObject) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SBForecastingErrorObject) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SBForecastingErrorObject) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SBForecastingErrorObject) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingErrorObject) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SBForecastingErrorObject) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SBForecastingErrorObject) SetDescription(v string) {
	o.Description = &v
}

func (o SBForecastingErrorObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSBForecastingErrorObject struct {
	value *SBForecastingErrorObject
	isSet bool
}

func (v NullableSBForecastingErrorObject) Get() *SBForecastingErrorObject {
	return v.value
}

func (v *NullableSBForecastingErrorObject) Set(val *SBForecastingErrorObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingErrorObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingErrorObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingErrorObject(val *SBForecastingErrorObject) *NullableSBForecastingErrorObject {
	return &NullableSBForecastingErrorObject{value: val, isSet: true}
}

func (v NullableSBForecastingErrorObject) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingErrorObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
