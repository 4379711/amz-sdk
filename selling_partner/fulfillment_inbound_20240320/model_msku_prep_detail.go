package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the MskuPrepDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MskuPrepDetail{}

// MskuPrepDetail An MSKU and its related prep details.
type MskuPrepDetail struct {
	AllOwnersConstraint  *AllOwnersConstraint `json:"allOwnersConstraint,omitempty"`
	LabelOwnerConstraint *OwnerConstraint     `json:"labelOwnerConstraint,omitempty"`
	// The merchant SKU, a merchant-supplied identifier for a specific SKU.
	Msku                string           `json:"msku"`
	PrepCategory        PrepCategory     `json:"prepCategory"`
	PrepOwnerConstraint *OwnerConstraint `json:"prepOwnerConstraint,omitempty"`
	// A list of preparation types associated with a preparation category.
	PrepTypes []PrepType `json:"prepTypes"`
}

type _MskuPrepDetail MskuPrepDetail

// NewMskuPrepDetail instantiates a new MskuPrepDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMskuPrepDetail(msku string, prepCategory PrepCategory, prepTypes []PrepType) *MskuPrepDetail {
	this := MskuPrepDetail{}
	this.Msku = msku
	this.PrepCategory = prepCategory
	this.PrepTypes = prepTypes
	return &this
}

// NewMskuPrepDetailWithDefaults instantiates a new MskuPrepDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMskuPrepDetailWithDefaults() *MskuPrepDetail {
	this := MskuPrepDetail{}
	return &this
}

// GetAllOwnersConstraint returns the AllOwnersConstraint field value if set, zero value otherwise.
func (o *MskuPrepDetail) GetAllOwnersConstraint() AllOwnersConstraint {
	if o == nil || IsNil(o.AllOwnersConstraint) {
		var ret AllOwnersConstraint
		return ret
	}
	return *o.AllOwnersConstraint
}

// GetAllOwnersConstraintOk returns a tuple with the AllOwnersConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MskuPrepDetail) GetAllOwnersConstraintOk() (*AllOwnersConstraint, bool) {
	if o == nil || IsNil(o.AllOwnersConstraint) {
		return nil, false
	}
	return o.AllOwnersConstraint, true
}

// HasAllOwnersConstraint returns a boolean if a field has been set.
func (o *MskuPrepDetail) HasAllOwnersConstraint() bool {
	if o != nil && !IsNil(o.AllOwnersConstraint) {
		return true
	}

	return false
}

// SetAllOwnersConstraint gets a reference to the given AllOwnersConstraint and assigns it to the AllOwnersConstraint field.
func (o *MskuPrepDetail) SetAllOwnersConstraint(v AllOwnersConstraint) {
	o.AllOwnersConstraint = &v
}

// GetLabelOwnerConstraint returns the LabelOwnerConstraint field value if set, zero value otherwise.
func (o *MskuPrepDetail) GetLabelOwnerConstraint() OwnerConstraint {
	if o == nil || IsNil(o.LabelOwnerConstraint) {
		var ret OwnerConstraint
		return ret
	}
	return *o.LabelOwnerConstraint
}

// GetLabelOwnerConstraintOk returns a tuple with the LabelOwnerConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MskuPrepDetail) GetLabelOwnerConstraintOk() (*OwnerConstraint, bool) {
	if o == nil || IsNil(o.LabelOwnerConstraint) {
		return nil, false
	}
	return o.LabelOwnerConstraint, true
}

// HasLabelOwnerConstraint returns a boolean if a field has been set.
func (o *MskuPrepDetail) HasLabelOwnerConstraint() bool {
	if o != nil && !IsNil(o.LabelOwnerConstraint) {
		return true
	}

	return false
}

// SetLabelOwnerConstraint gets a reference to the given OwnerConstraint and assigns it to the LabelOwnerConstraint field.
func (o *MskuPrepDetail) SetLabelOwnerConstraint(v OwnerConstraint) {
	o.LabelOwnerConstraint = &v
}

// GetMsku returns the Msku field value
func (o *MskuPrepDetail) GetMsku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msku
}

// GetMskuOk returns a tuple with the Msku field value
// and a boolean to check if the value has been set.
func (o *MskuPrepDetail) GetMskuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msku, true
}

// SetMsku sets field value
func (o *MskuPrepDetail) SetMsku(v string) {
	o.Msku = v
}

// GetPrepCategory returns the PrepCategory field value
func (o *MskuPrepDetail) GetPrepCategory() PrepCategory {
	if o == nil {
		var ret PrepCategory
		return ret
	}

	return o.PrepCategory
}

// GetPrepCategoryOk returns a tuple with the PrepCategory field value
// and a boolean to check if the value has been set.
func (o *MskuPrepDetail) GetPrepCategoryOk() (*PrepCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrepCategory, true
}

// SetPrepCategory sets field value
func (o *MskuPrepDetail) SetPrepCategory(v PrepCategory) {
	o.PrepCategory = v
}

// GetPrepOwnerConstraint returns the PrepOwnerConstraint field value if set, zero value otherwise.
func (o *MskuPrepDetail) GetPrepOwnerConstraint() OwnerConstraint {
	if o == nil || IsNil(o.PrepOwnerConstraint) {
		var ret OwnerConstraint
		return ret
	}
	return *o.PrepOwnerConstraint
}

// GetPrepOwnerConstraintOk returns a tuple with the PrepOwnerConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MskuPrepDetail) GetPrepOwnerConstraintOk() (*OwnerConstraint, bool) {
	if o == nil || IsNil(o.PrepOwnerConstraint) {
		return nil, false
	}
	return o.PrepOwnerConstraint, true
}

// HasPrepOwnerConstraint returns a boolean if a field has been set.
func (o *MskuPrepDetail) HasPrepOwnerConstraint() bool {
	if o != nil && !IsNil(o.PrepOwnerConstraint) {
		return true
	}

	return false
}

// SetPrepOwnerConstraint gets a reference to the given OwnerConstraint and assigns it to the PrepOwnerConstraint field.
func (o *MskuPrepDetail) SetPrepOwnerConstraint(v OwnerConstraint) {
	o.PrepOwnerConstraint = &v
}

// GetPrepTypes returns the PrepTypes field value
func (o *MskuPrepDetail) GetPrepTypes() []PrepType {
	if o == nil {
		var ret []PrepType
		return ret
	}

	return o.PrepTypes
}

// GetPrepTypesOk returns a tuple with the PrepTypes field value
// and a boolean to check if the value has been set.
func (o *MskuPrepDetail) GetPrepTypesOk() ([]PrepType, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrepTypes, true
}

// SetPrepTypes sets field value
func (o *MskuPrepDetail) SetPrepTypes(v []PrepType) {
	o.PrepTypes = v
}

func (o MskuPrepDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllOwnersConstraint) {
		toSerialize["allOwnersConstraint"] = o.AllOwnersConstraint
	}
	if !IsNil(o.LabelOwnerConstraint) {
		toSerialize["labelOwnerConstraint"] = o.LabelOwnerConstraint
	}
	toSerialize["msku"] = o.Msku
	toSerialize["prepCategory"] = o.PrepCategory
	if !IsNil(o.PrepOwnerConstraint) {
		toSerialize["prepOwnerConstraint"] = o.PrepOwnerConstraint
	}
	toSerialize["prepTypes"] = o.PrepTypes
	return toSerialize, nil
}

type NullableMskuPrepDetail struct {
	value *MskuPrepDetail
	isSet bool
}

func (v NullableMskuPrepDetail) Get() *MskuPrepDetail {
	return v.value
}

func (v *NullableMskuPrepDetail) Set(val *MskuPrepDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMskuPrepDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMskuPrepDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMskuPrepDetail(val *MskuPrepDetail) *NullableMskuPrepDetail {
	return &NullableMskuPrepDetail{value: val, isSet: true}
}

func (v NullableMskuPrepDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMskuPrepDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
