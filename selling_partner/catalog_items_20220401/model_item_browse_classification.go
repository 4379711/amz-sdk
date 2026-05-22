package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemBrowseClassification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemBrowseClassification{}

// ItemBrowseClassification Classification (browse node) for an Amazon catalog item.
type ItemBrowseClassification struct {
	// Display name for the classification.
	DisplayName string `json:"displayName"`
	// Identifier of the classification.
	ClassificationId string                    `json:"classificationId"`
	Parent           *ItemBrowseClassification `json:"parent,omitempty"`
}

type _ItemBrowseClassification ItemBrowseClassification

// NewItemBrowseClassification instantiates a new ItemBrowseClassification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemBrowseClassification(displayName string, classificationId string) *ItemBrowseClassification {
	this := ItemBrowseClassification{}
	this.DisplayName = displayName
	this.ClassificationId = classificationId
	return &this
}

// NewItemBrowseClassificationWithDefaults instantiates a new ItemBrowseClassification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemBrowseClassificationWithDefaults() *ItemBrowseClassification {
	this := ItemBrowseClassification{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *ItemBrowseClassification) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ItemBrowseClassification) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ItemBrowseClassification) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetClassificationId returns the ClassificationId field value
func (o *ItemBrowseClassification) GetClassificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassificationId
}

// GetClassificationIdOk returns a tuple with the ClassificationId field value
// and a boolean to check if the value has been set.
func (o *ItemBrowseClassification) GetClassificationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassificationId, true
}

// SetClassificationId sets field value
func (o *ItemBrowseClassification) SetClassificationId(v string) {
	o.ClassificationId = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ItemBrowseClassification) GetParent() ItemBrowseClassification {
	if o == nil || IsNil(o.Parent) {
		var ret ItemBrowseClassification
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemBrowseClassification) GetParentOk() (*ItemBrowseClassification, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ItemBrowseClassification) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given ItemBrowseClassification and assigns it to the Parent field.
func (o *ItemBrowseClassification) SetParent(v ItemBrowseClassification) {
	o.Parent = &v
}

func (o ItemBrowseClassification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["classificationId"] = o.ClassificationId
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	return toSerialize, nil
}

type NullableItemBrowseClassification struct {
	value *ItemBrowseClassification
	isSet bool
}

func (v NullableItemBrowseClassification) Get() *ItemBrowseClassification {
	return v.value
}

func (v *NullableItemBrowseClassification) Set(val *ItemBrowseClassification) {
	v.value = val
	v.isSet = true
}

func (v NullableItemBrowseClassification) IsSet() bool {
	return v.isSet
}

func (v *NullableItemBrowseClassification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemBrowseClassification(val *ItemBrowseClassification) *NullableItemBrowseClassification {
	return &NullableItemBrowseClassification{value: val, isSet: true}
}

func (v NullableItemBrowseClassification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemBrowseClassification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
