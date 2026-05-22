package finances_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the RelatedIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedIdentifier{}

// RelatedIdentifier Related business identifier of the transaction.
type RelatedIdentifier struct {
	// Enumerated set of related business identifier names.
	RelatedIdentifierName *string `json:"relatedIdentifierName,omitempty"`
	// Corresponding value of RelatedIdentifierName
	RelatedIdentifierValue *string `json:"relatedIdentifierValue,omitempty"`
}

// NewRelatedIdentifier instantiates a new RelatedIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedIdentifier() *RelatedIdentifier {
	this := RelatedIdentifier{}
	return &this
}

// NewRelatedIdentifierWithDefaults instantiates a new RelatedIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedIdentifierWithDefaults() *RelatedIdentifier {
	this := RelatedIdentifier{}
	return &this
}

// GetRelatedIdentifierName returns the RelatedIdentifierName field value if set, zero value otherwise.
func (o *RelatedIdentifier) GetRelatedIdentifierName() string {
	if o == nil || IsNil(o.RelatedIdentifierName) {
		var ret string
		return ret
	}
	return *o.RelatedIdentifierName
}

// GetRelatedIdentifierNameOk returns a tuple with the RelatedIdentifierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedIdentifier) GetRelatedIdentifierNameOk() (*string, bool) {
	if o == nil || IsNil(o.RelatedIdentifierName) {
		return nil, false
	}
	return o.RelatedIdentifierName, true
}

// HasRelatedIdentifierName returns a boolean if a field has been set.
func (o *RelatedIdentifier) HasRelatedIdentifierName() bool {
	if o != nil && !IsNil(o.RelatedIdentifierName) {
		return true
	}

	return false
}

// SetRelatedIdentifierName gets a reference to the given string and assigns it to the RelatedIdentifierName field.
func (o *RelatedIdentifier) SetRelatedIdentifierName(v string) {
	o.RelatedIdentifierName = &v
}

// GetRelatedIdentifierValue returns the RelatedIdentifierValue field value if set, zero value otherwise.
func (o *RelatedIdentifier) GetRelatedIdentifierValue() string {
	if o == nil || IsNil(o.RelatedIdentifierValue) {
		var ret string
		return ret
	}
	return *o.RelatedIdentifierValue
}

// GetRelatedIdentifierValueOk returns a tuple with the RelatedIdentifierValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedIdentifier) GetRelatedIdentifierValueOk() (*string, bool) {
	if o == nil || IsNil(o.RelatedIdentifierValue) {
		return nil, false
	}
	return o.RelatedIdentifierValue, true
}

// HasRelatedIdentifierValue returns a boolean if a field has been set.
func (o *RelatedIdentifier) HasRelatedIdentifierValue() bool {
	if o != nil && !IsNil(o.RelatedIdentifierValue) {
		return true
	}

	return false
}

// SetRelatedIdentifierValue gets a reference to the given string and assigns it to the RelatedIdentifierValue field.
func (o *RelatedIdentifier) SetRelatedIdentifierValue(v string) {
	o.RelatedIdentifierValue = &v
}

func (o RelatedIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RelatedIdentifierName) {
		toSerialize["relatedIdentifierName"] = o.RelatedIdentifierName
	}
	if !IsNil(o.RelatedIdentifierValue) {
		toSerialize["relatedIdentifierValue"] = o.RelatedIdentifierValue
	}
	return toSerialize, nil
}

type NullableRelatedIdentifier struct {
	value *RelatedIdentifier
	isSet bool
}

func (v NullableRelatedIdentifier) Get() *RelatedIdentifier {
	return v.value
}

func (v *NullableRelatedIdentifier) Set(val *RelatedIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedIdentifier(val *RelatedIdentifier) *NullableRelatedIdentifier {
	return &NullableRelatedIdentifier{value: val, isSet: true}
}

func (v NullableRelatedIdentifier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRelatedIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
