package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteAllSPTargetsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteAllSPTargetsResponseContent{}

// SponsoredProductsDeleteAllSPTargetsResponseContent struct for SponsoredProductsDeleteAllSPTargetsResponseContent
type SponsoredProductsDeleteAllSPTargetsResponseContent struct {
	Success []SponsoredProductsAllTargetsSuccessResponseItem `json:"success"`
	Error   []SponsoredProductsAllTargetsFailureResponseItem `json:"error,omitempty"`
}

type _SponsoredProductsDeleteAllSPTargetsResponseContent SponsoredProductsDeleteAllSPTargetsResponseContent

// NewSponsoredProductsDeleteAllSPTargetsResponseContent instantiates a new SponsoredProductsDeleteAllSPTargetsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteAllSPTargetsResponseContent(success []SponsoredProductsAllTargetsSuccessResponseItem) *SponsoredProductsDeleteAllSPTargetsResponseContent {
	this := SponsoredProductsDeleteAllSPTargetsResponseContent{}
	this.Success = success
	return &this
}

// NewSponsoredProductsDeleteAllSPTargetsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteAllSPTargetsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteAllSPTargetsResponseContentWithDefaults() *SponsoredProductsDeleteAllSPTargetsResponseContent {
	this := SponsoredProductsDeleteAllSPTargetsResponseContent{}
	return &this
}

// GetSuccess returns the Success field value
func (o *SponsoredProductsDeleteAllSPTargetsResponseContent) GetSuccess() []SponsoredProductsAllTargetsSuccessResponseItem {
	if o == nil {
		var ret []SponsoredProductsAllTargetsSuccessResponseItem
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteAllSPTargetsResponseContent) GetSuccessOk() ([]SponsoredProductsAllTargetsSuccessResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Success, true
}

// SetSuccess sets field value
func (o *SponsoredProductsDeleteAllSPTargetsResponseContent) SetSuccess(v []SponsoredProductsAllTargetsSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsDeleteAllSPTargetsResponseContent) GetError() []SponsoredProductsAllTargetsFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsAllTargetsFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteAllSPTargetsResponseContent) GetErrorOk() ([]SponsoredProductsAllTargetsFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsDeleteAllSPTargetsResponseContent) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsAllTargetsFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsDeleteAllSPTargetsResponseContent) SetError(v []SponsoredProductsAllTargetsFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsDeleteAllSPTargetsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteAllSPTargetsResponseContent struct {
	value *SponsoredProductsDeleteAllSPTargetsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteAllSPTargetsResponseContent) Get() *SponsoredProductsDeleteAllSPTargetsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteAllSPTargetsResponseContent) Set(val *SponsoredProductsDeleteAllSPTargetsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteAllSPTargetsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteAllSPTargetsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteAllSPTargetsResponseContent(val *SponsoredProductsDeleteAllSPTargetsResponseContent) *NullableSponsoredProductsDeleteAllSPTargetsResponseContent {
	return &NullableSponsoredProductsDeleteAllSPTargetsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteAllSPTargetsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteAllSPTargetsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
