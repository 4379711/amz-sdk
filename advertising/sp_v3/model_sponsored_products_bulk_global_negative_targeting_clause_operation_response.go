package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse{}

// SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse struct for SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse
type SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse struct {
	Success []SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse instantiates a new SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse() *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse {
	this := SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponseWithDefaults instantiates a new SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponseWithDefaults() *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse {
	this := SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) GetSuccess() []SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) GetSuccessOk() ([]SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) SetSuccess(v []SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) GetError() []SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) GetErrorOk() ([]SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) SetError(v []SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse struct {
	value *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) Get() *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) Set(val *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse(val *SponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) *NullableSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse {
	return &NullableSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkGlobalNegativeTargetingClauseOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
