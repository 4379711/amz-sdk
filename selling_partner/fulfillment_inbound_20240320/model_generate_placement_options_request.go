package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the GeneratePlacementOptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneratePlacementOptionsRequest{}

// GeneratePlacementOptionsRequest The `generatePlacementOptions` request.
type GeneratePlacementOptionsRequest struct {
	// Custom placement options you want to add to the plan. This is only used for the India (IN - A21TJRUUN4KGV) marketplace.
	CustomPlacement []CustomPlacementInput `json:"customPlacement,omitempty"`
}

// NewGeneratePlacementOptionsRequest instantiates a new GeneratePlacementOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneratePlacementOptionsRequest() *GeneratePlacementOptionsRequest {
	this := GeneratePlacementOptionsRequest{}
	return &this
}

// NewGeneratePlacementOptionsRequestWithDefaults instantiates a new GeneratePlacementOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneratePlacementOptionsRequestWithDefaults() *GeneratePlacementOptionsRequest {
	this := GeneratePlacementOptionsRequest{}
	return &this
}

// GetCustomPlacement returns the CustomPlacement field value if set, zero value otherwise.
func (o *GeneratePlacementOptionsRequest) GetCustomPlacement() []CustomPlacementInput {
	if o == nil || IsNil(o.CustomPlacement) {
		var ret []CustomPlacementInput
		return ret
	}
	return o.CustomPlacement
}

// GetCustomPlacementOk returns a tuple with the CustomPlacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePlacementOptionsRequest) GetCustomPlacementOk() ([]CustomPlacementInput, bool) {
	if o == nil || IsNil(o.CustomPlacement) {
		return nil, false
	}
	return o.CustomPlacement, true
}

// HasCustomPlacement returns a boolean if a field has been set.
func (o *GeneratePlacementOptionsRequest) HasCustomPlacement() bool {
	if o != nil && !IsNil(o.CustomPlacement) {
		return true
	}

	return false
}

// SetCustomPlacement gets a reference to the given []CustomPlacementInput and assigns it to the CustomPlacement field.
func (o *GeneratePlacementOptionsRequest) SetCustomPlacement(v []CustomPlacementInput) {
	o.CustomPlacement = v
}

func (o GeneratePlacementOptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomPlacement) {
		toSerialize["customPlacement"] = o.CustomPlacement
	}
	return toSerialize, nil
}

type NullableGeneratePlacementOptionsRequest struct {
	value *GeneratePlacementOptionsRequest
	isSet bool
}

func (v NullableGeneratePlacementOptionsRequest) Get() *GeneratePlacementOptionsRequest {
	return v.value
}

func (v *NullableGeneratePlacementOptionsRequest) Set(val *GeneratePlacementOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratePlacementOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratePlacementOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratePlacementOptionsRequest(val *GeneratePlacementOptionsRequest) *NullableGeneratePlacementOptionsRequest {
	return &NullableGeneratePlacementOptionsRequest{value: val, isSet: true}
}

func (v NullableGeneratePlacementOptionsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGeneratePlacementOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
