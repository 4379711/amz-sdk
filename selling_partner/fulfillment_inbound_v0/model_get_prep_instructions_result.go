package fulfillment_inbound_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetPrepInstructionsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPrepInstructionsResult{}

// GetPrepInstructionsResult Result for the get prep instructions operation
type GetPrepInstructionsResult struct {
	// A list of SKU labeling requirements and item preparation instructions.
	SKUPrepInstructionsList []SKUPrepInstructions `json:"SKUPrepInstructionsList,omitempty"`
	// A list of invalid SKU values and the reason they are invalid.
	InvalidSKUList []InvalidSKU `json:"InvalidSKUList,omitempty"`
	// A list of item preparation instructions.
	ASINPrepInstructionsList []ASINPrepInstructions `json:"ASINPrepInstructionsList,omitempty"`
	// A list of invalid ASIN values and the reasons they are invalid.
	InvalidASINList []InvalidASIN `json:"InvalidASINList,omitempty"`
}

// NewGetPrepInstructionsResult instantiates a new GetPrepInstructionsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPrepInstructionsResult() *GetPrepInstructionsResult {
	this := GetPrepInstructionsResult{}
	return &this
}

// NewGetPrepInstructionsResultWithDefaults instantiates a new GetPrepInstructionsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPrepInstructionsResultWithDefaults() *GetPrepInstructionsResult {
	this := GetPrepInstructionsResult{}
	return &this
}

// GetSKUPrepInstructionsList returns the SKUPrepInstructionsList field value if set, zero value otherwise.
func (o *GetPrepInstructionsResult) GetSKUPrepInstructionsList() []SKUPrepInstructions {
	if o == nil || IsNil(o.SKUPrepInstructionsList) {
		var ret []SKUPrepInstructions
		return ret
	}
	return o.SKUPrepInstructionsList
}

// GetSKUPrepInstructionsListOk returns a tuple with the SKUPrepInstructionsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPrepInstructionsResult) GetSKUPrepInstructionsListOk() ([]SKUPrepInstructions, bool) {
	if o == nil || IsNil(o.SKUPrepInstructionsList) {
		return nil, false
	}
	return o.SKUPrepInstructionsList, true
}

// HasSKUPrepInstructionsList returns a boolean if a field has been set.
func (o *GetPrepInstructionsResult) HasSKUPrepInstructionsList() bool {
	if o != nil && !IsNil(o.SKUPrepInstructionsList) {
		return true
	}

	return false
}

// SetSKUPrepInstructionsList gets a reference to the given []SKUPrepInstructions and assigns it to the SKUPrepInstructionsList field.
func (o *GetPrepInstructionsResult) SetSKUPrepInstructionsList(v []SKUPrepInstructions) {
	o.SKUPrepInstructionsList = v
}

// GetInvalidSKUList returns the InvalidSKUList field value if set, zero value otherwise.
func (o *GetPrepInstructionsResult) GetInvalidSKUList() []InvalidSKU {
	if o == nil || IsNil(o.InvalidSKUList) {
		var ret []InvalidSKU
		return ret
	}
	return o.InvalidSKUList
}

// GetInvalidSKUListOk returns a tuple with the InvalidSKUList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPrepInstructionsResult) GetInvalidSKUListOk() ([]InvalidSKU, bool) {
	if o == nil || IsNil(o.InvalidSKUList) {
		return nil, false
	}
	return o.InvalidSKUList, true
}

// HasInvalidSKUList returns a boolean if a field has been set.
func (o *GetPrepInstructionsResult) HasInvalidSKUList() bool {
	if o != nil && !IsNil(o.InvalidSKUList) {
		return true
	}

	return false
}

// SetInvalidSKUList gets a reference to the given []InvalidSKU and assigns it to the InvalidSKUList field.
func (o *GetPrepInstructionsResult) SetInvalidSKUList(v []InvalidSKU) {
	o.InvalidSKUList = v
}

// GetASINPrepInstructionsList returns the ASINPrepInstructionsList field value if set, zero value otherwise.
func (o *GetPrepInstructionsResult) GetASINPrepInstructionsList() []ASINPrepInstructions {
	if o == nil || IsNil(o.ASINPrepInstructionsList) {
		var ret []ASINPrepInstructions
		return ret
	}
	return o.ASINPrepInstructionsList
}

// GetASINPrepInstructionsListOk returns a tuple with the ASINPrepInstructionsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPrepInstructionsResult) GetASINPrepInstructionsListOk() ([]ASINPrepInstructions, bool) {
	if o == nil || IsNil(o.ASINPrepInstructionsList) {
		return nil, false
	}
	return o.ASINPrepInstructionsList, true
}

// HasASINPrepInstructionsList returns a boolean if a field has been set.
func (o *GetPrepInstructionsResult) HasASINPrepInstructionsList() bool {
	if o != nil && !IsNil(o.ASINPrepInstructionsList) {
		return true
	}

	return false
}

// SetASINPrepInstructionsList gets a reference to the given []ASINPrepInstructions and assigns it to the ASINPrepInstructionsList field.
func (o *GetPrepInstructionsResult) SetASINPrepInstructionsList(v []ASINPrepInstructions) {
	o.ASINPrepInstructionsList = v
}

// GetInvalidASINList returns the InvalidASINList field value if set, zero value otherwise.
func (o *GetPrepInstructionsResult) GetInvalidASINList() []InvalidASIN {
	if o == nil || IsNil(o.InvalidASINList) {
		var ret []InvalidASIN
		return ret
	}
	return o.InvalidASINList
}

// GetInvalidASINListOk returns a tuple with the InvalidASINList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPrepInstructionsResult) GetInvalidASINListOk() ([]InvalidASIN, bool) {
	if o == nil || IsNil(o.InvalidASINList) {
		return nil, false
	}
	return o.InvalidASINList, true
}

// HasInvalidASINList returns a boolean if a field has been set.
func (o *GetPrepInstructionsResult) HasInvalidASINList() bool {
	if o != nil && !IsNil(o.InvalidASINList) {
		return true
	}

	return false
}

// SetInvalidASINList gets a reference to the given []InvalidASIN and assigns it to the InvalidASINList field.
func (o *GetPrepInstructionsResult) SetInvalidASINList(v []InvalidASIN) {
	o.InvalidASINList = v
}

func (o GetPrepInstructionsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SKUPrepInstructionsList) {
		toSerialize["SKUPrepInstructionsList"] = o.SKUPrepInstructionsList
	}
	if !IsNil(o.InvalidSKUList) {
		toSerialize["InvalidSKUList"] = o.InvalidSKUList
	}
	if !IsNil(o.ASINPrepInstructionsList) {
		toSerialize["ASINPrepInstructionsList"] = o.ASINPrepInstructionsList
	}
	if !IsNil(o.InvalidASINList) {
		toSerialize["InvalidASINList"] = o.InvalidASINList
	}
	return toSerialize, nil
}

type NullableGetPrepInstructionsResult struct {
	value *GetPrepInstructionsResult
	isSet bool
}

func (v NullableGetPrepInstructionsResult) Get() *GetPrepInstructionsResult {
	return v.value
}

func (v *NullableGetPrepInstructionsResult) Set(val *GetPrepInstructionsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPrepInstructionsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPrepInstructionsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPrepInstructionsResult(val *GetPrepInstructionsResult) *NullableGetPrepInstructionsResult {
	return &NullableGetPrepInstructionsResult{value: val, isSet: true}
}

func (v NullableGetPrepInstructionsResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetPrepInstructionsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
