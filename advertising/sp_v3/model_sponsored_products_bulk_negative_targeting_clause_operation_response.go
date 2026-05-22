package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBulkNegativeTargetingClauseOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBulkNegativeTargetingClauseOperationResponse{}

// SponsoredProductsBulkNegativeTargetingClauseOperationResponse struct for SponsoredProductsBulkNegativeTargetingClauseOperationResponse
type SponsoredProductsBulkNegativeTargetingClauseOperationResponse struct {
	Success []SponsoredProductsNegativeTargetingClauseSuccessResponseItem `json:"success,omitempty"`
	Error   []SponsoredProductsNegativeTargetingClauseFailureResponseItem `json:"error,omitempty"`
}

// NewSponsoredProductsBulkNegativeTargetingClauseOperationResponse instantiates a new SponsoredProductsBulkNegativeTargetingClauseOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBulkNegativeTargetingClauseOperationResponse() *SponsoredProductsBulkNegativeTargetingClauseOperationResponse {
	this := SponsoredProductsBulkNegativeTargetingClauseOperationResponse{}
	return &this
}

// NewSponsoredProductsBulkNegativeTargetingClauseOperationResponseWithDefaults instantiates a new SponsoredProductsBulkNegativeTargetingClauseOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBulkNegativeTargetingClauseOperationResponseWithDefaults() *SponsoredProductsBulkNegativeTargetingClauseOperationResponse {
	this := SponsoredProductsBulkNegativeTargetingClauseOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsBulkNegativeTargetingClauseOperationResponse) GetSuccess() []SponsoredProductsNegativeTargetingClauseSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsNegativeTargetingClauseSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkNegativeTargetingClauseOperationResponse) GetSuccessOk() ([]SponsoredProductsNegativeTargetingClauseSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsBulkNegativeTargetingClauseOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsNegativeTargetingClauseSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsBulkNegativeTargetingClauseOperationResponse) SetSuccess(v []SponsoredProductsNegativeTargetingClauseSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SponsoredProductsBulkNegativeTargetingClauseOperationResponse) GetError() []SponsoredProductsNegativeTargetingClauseFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []SponsoredProductsNegativeTargetingClauseFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBulkNegativeTargetingClauseOperationResponse) GetErrorOk() ([]SponsoredProductsNegativeTargetingClauseFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SponsoredProductsBulkNegativeTargetingClauseOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []SponsoredProductsNegativeTargetingClauseFailureResponseItem and assigns it to the Error field.
func (o *SponsoredProductsBulkNegativeTargetingClauseOperationResponse) SetError(v []SponsoredProductsNegativeTargetingClauseFailureResponseItem) {
	o.Error = v
}

func (o SponsoredProductsBulkNegativeTargetingClauseOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBulkNegativeTargetingClauseOperationResponse struct {
	value *SponsoredProductsBulkNegativeTargetingClauseOperationResponse
	isSet bool
}

func (v NullableSponsoredProductsBulkNegativeTargetingClauseOperationResponse) Get() *SponsoredProductsBulkNegativeTargetingClauseOperationResponse {
	return v.value
}

func (v *NullableSponsoredProductsBulkNegativeTargetingClauseOperationResponse) Set(val *SponsoredProductsBulkNegativeTargetingClauseOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBulkNegativeTargetingClauseOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBulkNegativeTargetingClauseOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBulkNegativeTargetingClauseOperationResponse(val *SponsoredProductsBulkNegativeTargetingClauseOperationResponse) *NullableSponsoredProductsBulkNegativeTargetingClauseOperationResponse {
	return &NullableSponsoredProductsBulkNegativeTargetingClauseOperationResponse{value: val, isSet: true}
}

func (v NullableSponsoredProductsBulkNegativeTargetingClauseOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBulkNegativeTargetingClauseOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
