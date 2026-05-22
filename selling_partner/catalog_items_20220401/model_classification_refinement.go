package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ClassificationRefinement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClassificationRefinement{}

// ClassificationRefinement A classification that you can use to refine your search.
type ClassificationRefinement struct {
	// The estimated number of results that would be returned if you refine your search by the specified `classificationId`.
	NumberOfResults int32 `json:"numberOfResults"`
	// Display name for the classification.
	DisplayName string `json:"displayName"`
	// The identifier of the classification that you can use to refine your search.
	ClassificationId string `json:"classificationId"`
}

type _ClassificationRefinement ClassificationRefinement

// NewClassificationRefinement instantiates a new ClassificationRefinement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassificationRefinement(numberOfResults int32, displayName string, classificationId string) *ClassificationRefinement {
	this := ClassificationRefinement{}
	this.NumberOfResults = numberOfResults
	this.DisplayName = displayName
	this.ClassificationId = classificationId
	return &this
}

// NewClassificationRefinementWithDefaults instantiates a new ClassificationRefinement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassificationRefinementWithDefaults() *ClassificationRefinement {
	this := ClassificationRefinement{}
	return &this
}

// GetNumberOfResults returns the NumberOfResults field value
func (o *ClassificationRefinement) GetNumberOfResults() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfResults
}

// GetNumberOfResultsOk returns a tuple with the NumberOfResults field value
// and a boolean to check if the value has been set.
func (o *ClassificationRefinement) GetNumberOfResultsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfResults, true
}

// SetNumberOfResults sets field value
func (o *ClassificationRefinement) SetNumberOfResults(v int32) {
	o.NumberOfResults = v
}

// GetDisplayName returns the DisplayName field value
func (o *ClassificationRefinement) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ClassificationRefinement) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ClassificationRefinement) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetClassificationId returns the ClassificationId field value
func (o *ClassificationRefinement) GetClassificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassificationId
}

// GetClassificationIdOk returns a tuple with the ClassificationId field value
// and a boolean to check if the value has been set.
func (o *ClassificationRefinement) GetClassificationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassificationId, true
}

// SetClassificationId sets field value
func (o *ClassificationRefinement) SetClassificationId(v string) {
	o.ClassificationId = v
}

func (o ClassificationRefinement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["numberOfResults"] = o.NumberOfResults
	toSerialize["displayName"] = o.DisplayName
	toSerialize["classificationId"] = o.ClassificationId
	return toSerialize, nil
}

type NullableClassificationRefinement struct {
	value *ClassificationRefinement
	isSet bool
}

func (v NullableClassificationRefinement) Get() *ClassificationRefinement {
	return v.value
}

func (v *NullableClassificationRefinement) Set(val *ClassificationRefinement) {
	v.value = val
	v.isSet = true
}

func (v NullableClassificationRefinement) IsSet() bool {
	return v.isSet
}

func (v *NullableClassificationRefinement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassificationRefinement(val *ClassificationRefinement) *NullableClassificationRefinement {
	return &NullableClassificationRefinement{value: val, isSet: true}
}

func (v NullableClassificationRefinement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableClassificationRefinement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
