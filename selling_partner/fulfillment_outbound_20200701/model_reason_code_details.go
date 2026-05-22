package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the ReasonCodeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReasonCodeDetails{}

// ReasonCodeDetails A return reason code, a description, and an optional description translation.
type ReasonCodeDetails struct {
	// A code that indicates a valid return reason.
	ReturnReasonCode string `json:"returnReasonCode"`
	// A human readable description of the return reason code.
	Description string `json:"description"`
	// A translation of the description. The translation is in the language specified in the Language request parameter.
	TranslatedDescription *string `json:"translatedDescription,omitempty"`
}

type _ReasonCodeDetails ReasonCodeDetails

// NewReasonCodeDetails instantiates a new ReasonCodeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReasonCodeDetails(returnReasonCode string, description string) *ReasonCodeDetails {
	this := ReasonCodeDetails{}
	this.ReturnReasonCode = returnReasonCode
	this.Description = description
	return &this
}

// NewReasonCodeDetailsWithDefaults instantiates a new ReasonCodeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReasonCodeDetailsWithDefaults() *ReasonCodeDetails {
	this := ReasonCodeDetails{}
	return &this
}

// GetReturnReasonCode returns the ReturnReasonCode field value
func (o *ReasonCodeDetails) GetReturnReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnReasonCode
}

// GetReturnReasonCodeOk returns a tuple with the ReturnReasonCode field value
// and a boolean to check if the value has been set.
func (o *ReasonCodeDetails) GetReturnReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnReasonCode, true
}

// SetReturnReasonCode sets field value
func (o *ReasonCodeDetails) SetReturnReasonCode(v string) {
	o.ReturnReasonCode = v
}

// GetDescription returns the Description field value
func (o *ReasonCodeDetails) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ReasonCodeDetails) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ReasonCodeDetails) SetDescription(v string) {
	o.Description = v
}

// GetTranslatedDescription returns the TranslatedDescription field value if set, zero value otherwise.
func (o *ReasonCodeDetails) GetTranslatedDescription() string {
	if o == nil || IsNil(o.TranslatedDescription) {
		var ret string
		return ret
	}
	return *o.TranslatedDescription
}

// GetTranslatedDescriptionOk returns a tuple with the TranslatedDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonCodeDetails) GetTranslatedDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TranslatedDescription) {
		return nil, false
	}
	return o.TranslatedDescription, true
}

// HasTranslatedDescription returns a boolean if a field has been set.
func (o *ReasonCodeDetails) HasTranslatedDescription() bool {
	if o != nil && !IsNil(o.TranslatedDescription) {
		return true
	}

	return false
}

// SetTranslatedDescription gets a reference to the given string and assigns it to the TranslatedDescription field.
func (o *ReasonCodeDetails) SetTranslatedDescription(v string) {
	o.TranslatedDescription = &v
}

func (o ReasonCodeDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["returnReasonCode"] = o.ReturnReasonCode
	toSerialize["description"] = o.Description
	if !IsNil(o.TranslatedDescription) {
		toSerialize["translatedDescription"] = o.TranslatedDescription
	}
	return toSerialize, nil
}

type NullableReasonCodeDetails struct {
	value *ReasonCodeDetails
	isSet bool
}

func (v NullableReasonCodeDetails) Get() *ReasonCodeDetails {
	return v.value
}

func (v *NullableReasonCodeDetails) Set(val *ReasonCodeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableReasonCodeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableReasonCodeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReasonCodeDetails(val *ReasonCodeDetails) *NullableReasonCodeDetails {
	return &NullableReasonCodeDetails{value: val, isSet: true}
}

func (v NullableReasonCodeDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReasonCodeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
