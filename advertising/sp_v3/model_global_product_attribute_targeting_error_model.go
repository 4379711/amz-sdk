package sp_v3

import "github.com/bytedance/sonic"

// checks if the GlobalProductAttributeTargetingErrorModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalProductAttributeTargetingErrorModel{}

// GlobalProductAttributeTargetingErrorModel struct for GlobalProductAttributeTargetingErrorModel
type GlobalProductAttributeTargetingErrorModel struct {
	CountryCodes []string `json:"countryCodes,omitempty"`
	Code         string   `json:"code"`
	Message      *string  `json:"message,omitempty"`
}

type _GlobalProductAttributeTargetingErrorModel GlobalProductAttributeTargetingErrorModel

// NewGlobalProductAttributeTargetingErrorModel instantiates a new GlobalProductAttributeTargetingErrorModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalProductAttributeTargetingErrorModel(code string) *GlobalProductAttributeTargetingErrorModel {
	this := GlobalProductAttributeTargetingErrorModel{}
	this.Code = code
	return &this
}

// NewGlobalProductAttributeTargetingErrorModelWithDefaults instantiates a new GlobalProductAttributeTargetingErrorModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalProductAttributeTargetingErrorModelWithDefaults() *GlobalProductAttributeTargetingErrorModel {
	this := GlobalProductAttributeTargetingErrorModel{}
	return &this
}

// GetCountryCodes returns the CountryCodes field value if set, zero value otherwise.
func (o *GlobalProductAttributeTargetingErrorModel) GetCountryCodes() []string {
	if o == nil || IsNil(o.CountryCodes) {
		var ret []string
		return ret
	}
	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductAttributeTargetingErrorModel) GetCountryCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryCodes) {
		return nil, false
	}
	return o.CountryCodes, true
}

// HasCountryCodes returns a boolean if a field has been set.
func (o *GlobalProductAttributeTargetingErrorModel) HasCountryCodes() bool {
	if o != nil && !IsNil(o.CountryCodes) {
		return true
	}

	return false
}

// SetCountryCodes gets a reference to the given []string and assigns it to the CountryCodes field.
func (o *GlobalProductAttributeTargetingErrorModel) SetCountryCodes(v []string) {
	o.CountryCodes = v
}

// GetCode returns the Code field value
func (o *GlobalProductAttributeTargetingErrorModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GlobalProductAttributeTargetingErrorModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GlobalProductAttributeTargetingErrorModel) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GlobalProductAttributeTargetingErrorModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProductAttributeTargetingErrorModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GlobalProductAttributeTargetingErrorModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GlobalProductAttributeTargetingErrorModel) SetMessage(v string) {
	o.Message = &v
}

func (o GlobalProductAttributeTargetingErrorModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCodes) {
		toSerialize["countryCodes"] = o.CountryCodes
	}
	toSerialize["code"] = o.Code
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGlobalProductAttributeTargetingErrorModel struct {
	value *GlobalProductAttributeTargetingErrorModel
	isSet bool
}

func (v NullableGlobalProductAttributeTargetingErrorModel) Get() *GlobalProductAttributeTargetingErrorModel {
	return v.value
}

func (v *NullableGlobalProductAttributeTargetingErrorModel) Set(val *GlobalProductAttributeTargetingErrorModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalProductAttributeTargetingErrorModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalProductAttributeTargetingErrorModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalProductAttributeTargetingErrorModel(val *GlobalProductAttributeTargetingErrorModel) *NullableGlobalProductAttributeTargetingErrorModel {
	return &NullableGlobalProductAttributeTargetingErrorModel{value: val, isSet: true}
}

func (v NullableGlobalProductAttributeTargetingErrorModel) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalProductAttributeTargetingErrorModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
