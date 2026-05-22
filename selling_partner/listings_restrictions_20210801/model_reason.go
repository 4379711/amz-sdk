package listings_restrictions_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the Reason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reason{}

// Reason A reason for the restriction, including path forward links that may allow Selling Partners to remove the restriction, if available.
type Reason struct {
	// A message describing the reason for the restriction.
	Message string `json:"message"`
	// A code indicating why the listing is restricted.
	ReasonCode *string `json:"reasonCode,omitempty"`
	// A list of path forward links that may allow Selling Partners to remove the restriction.
	Links []Link `json:"links,omitempty"`
}

type _Reason Reason

// NewReason instantiates a new Reason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReason(message string) *Reason {
	this := Reason{}
	this.Message = message
	return &this
}

// NewReasonWithDefaults instantiates a new Reason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReasonWithDefaults() *Reason {
	this := Reason{}
	return &this
}

// GetMessage returns the Message field value
func (o *Reason) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Reason) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Reason) SetMessage(v string) {
	o.Message = v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *Reason) GetReasonCode() string {
	if o == nil || IsNil(o.ReasonCode) {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reason) GetReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *Reason) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *Reason) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Reason) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reason) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Reason) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *Reason) SetLinks(v []Link) {
	o.Links = v
}

func (o Reason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.ReasonCode) {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableReason struct {
	value *Reason
	isSet bool
}

func (v NullableReason) Get() *Reason {
	return v.value
}

func (v *NullableReason) Set(val *Reason) {
	v.value = val
	v.isSet = true
}

func (v NullableReason) IsSet() bool {
	return v.isSet
}

func (v *NullableReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReason(val *Reason) *NullableReason {
	return &NullableReason{value: val, isSet: true}
}

func (v NullableReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
