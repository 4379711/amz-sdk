package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductSet{}

// ProductSet struct for ProductSet
type ProductSet struct {
	ProofPoint *string `json:"proofPoint,omitempty"`
	Name       *string `json:"name,omitempty"`
}

// NewProductSet instantiates a new ProductSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductSet() *ProductSet {
	this := ProductSet{}
	return &this
}

// NewProductSetWithDefaults instantiates a new ProductSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductSetWithDefaults() *ProductSet {
	this := ProductSet{}
	return &this
}

// GetProofPoint returns the ProofPoint field value if set, zero value otherwise.
func (o *ProductSet) GetProofPoint() string {
	if o == nil || IsNil(o.ProofPoint) {
		var ret string
		return ret
	}
	return *o.ProofPoint
}

// GetProofPointOk returns a tuple with the ProofPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSet) GetProofPointOk() (*string, bool) {
	if o == nil || IsNil(o.ProofPoint) {
		return nil, false
	}
	return o.ProofPoint, true
}

// HasProofPoint returns a boolean if a field has been set.
func (o *ProductSet) HasProofPoint() bool {
	if o != nil && !IsNil(o.ProofPoint) {
		return true
	}

	return false
}

// SetProofPoint gets a reference to the given string and assigns it to the ProofPoint field.
func (o *ProductSet) SetProofPoint(v string) {
	o.ProofPoint = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductSet) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSet) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductSet) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductSet) SetName(v string) {
	o.Name = &v
}

func (o ProductSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProofPoint) {
		toSerialize["proofPoint"] = o.ProofPoint
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableProductSet struct {
	value *ProductSet
	isSet bool
}

func (v NullableProductSet) Get() *ProductSet {
	return v.value
}

func (v *NullableProductSet) Set(val *ProductSet) {
	v.value = val
	v.isSet = true
}

func (v NullableProductSet) IsSet() bool {
	return v.isSet
}

func (v *NullableProductSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductSet(val *ProductSet) *NullableProductSet {
	return &NullableProductSet{value: val, isSet: true}
}

func (v NullableProductSet) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
