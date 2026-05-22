package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BulkPortfolioOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkPortfolioOperationResponse{}

// BulkPortfolioOperationResponse struct for BulkPortfolioOperationResponse
type BulkPortfolioOperationResponse struct {
	Success []PortfolioSuccessResponseItem `json:"success,omitempty"`
	Error   []PortfolioFailureResponseItem `json:"error,omitempty"`
}

// NewBulkPortfolioOperationResponse instantiates a new BulkPortfolioOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkPortfolioOperationResponse() *BulkPortfolioOperationResponse {
	this := BulkPortfolioOperationResponse{}
	return &this
}

// NewBulkPortfolioOperationResponseWithDefaults instantiates a new BulkPortfolioOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkPortfolioOperationResponseWithDefaults() *BulkPortfolioOperationResponse {
	this := BulkPortfolioOperationResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BulkPortfolioOperationResponse) GetSuccess() []PortfolioSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []PortfolioSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkPortfolioOperationResponse) GetSuccessOk() ([]PortfolioSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BulkPortfolioOperationResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []PortfolioSuccessResponseItem and assigns it to the Success field.
func (o *BulkPortfolioOperationResponse) SetSuccess(v []PortfolioSuccessResponseItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BulkPortfolioOperationResponse) GetError() []PortfolioFailureResponseItem {
	if o == nil || IsNil(o.Error) {
		var ret []PortfolioFailureResponseItem
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkPortfolioOperationResponse) GetErrorOk() ([]PortfolioFailureResponseItem, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BulkPortfolioOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []PortfolioFailureResponseItem and assigns it to the Error field.
func (o *BulkPortfolioOperationResponse) SetError(v []PortfolioFailureResponseItem) {
	o.Error = v
}

func (o BulkPortfolioOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableBulkPortfolioOperationResponse struct {
	value *BulkPortfolioOperationResponse
	isSet bool
}

func (v NullableBulkPortfolioOperationResponse) Get() *BulkPortfolioOperationResponse {
	return v.value
}

func (v *NullableBulkPortfolioOperationResponse) Set(val *BulkPortfolioOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkPortfolioOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkPortfolioOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkPortfolioOperationResponse(val *BulkPortfolioOperationResponse) *NullableBulkPortfolioOperationResponse {
	return &NullableBulkPortfolioOperationResponse{value: val, isSet: true}
}

func (v NullableBulkPortfolioOperationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBulkPortfolioOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
