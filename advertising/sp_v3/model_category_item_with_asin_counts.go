package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CategoryItemWithAsinCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryItemWithAsinCounts{}

// CategoryItemWithAsinCounts struct for CategoryItemWithAsinCounts
type CategoryItemWithAsinCounts struct {
	// The path of the category, which contains the current category and all parent categories
	CategoryPath *string `json:"categoryPath,omitempty"`
	// The name of the category
	Name       *string       `json:"name,omitempty"`
	AsinCounts *IntegerRange `json:"asinCounts,omitempty"`
	// The category id of the parent node
	ParentCategoryId *string `json:"parentCategoryId,omitempty"`
	// The category id of the current node
	Id *string `json:"id,omitempty"`
}

// NewCategoryItemWithAsinCounts instantiates a new CategoryItemWithAsinCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryItemWithAsinCounts() *CategoryItemWithAsinCounts {
	this := CategoryItemWithAsinCounts{}
	return &this
}

// NewCategoryItemWithAsinCountsWithDefaults instantiates a new CategoryItemWithAsinCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryItemWithAsinCountsWithDefaults() *CategoryItemWithAsinCounts {
	this := CategoryItemWithAsinCounts{}
	return &this
}

// GetCategoryPath returns the CategoryPath field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCounts) GetCategoryPath() string {
	if o == nil || IsNil(o.CategoryPath) {
		var ret string
		return ret
	}
	return *o.CategoryPath
}

// GetCategoryPathOk returns a tuple with the CategoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCounts) GetCategoryPathOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryPath) {
		return nil, false
	}
	return o.CategoryPath, true
}

// HasCategoryPath returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCounts) HasCategoryPath() bool {
	if o != nil && !IsNil(o.CategoryPath) {
		return true
	}

	return false
}

// SetCategoryPath gets a reference to the given string and assigns it to the CategoryPath field.
func (o *CategoryItemWithAsinCounts) SetCategoryPath(v string) {
	o.CategoryPath = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCounts) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCounts) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCounts) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CategoryItemWithAsinCounts) SetName(v string) {
	o.Name = &v
}

// GetAsinCounts returns the AsinCounts field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCounts) GetAsinCounts() IntegerRange {
	if o == nil || IsNil(o.AsinCounts) {
		var ret IntegerRange
		return ret
	}
	return *o.AsinCounts
}

// GetAsinCountsOk returns a tuple with the AsinCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCounts) GetAsinCountsOk() (*IntegerRange, bool) {
	if o == nil || IsNil(o.AsinCounts) {
		return nil, false
	}
	return o.AsinCounts, true
}

// HasAsinCounts returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCounts) HasAsinCounts() bool {
	if o != nil && !IsNil(o.AsinCounts) {
		return true
	}

	return false
}

// SetAsinCounts gets a reference to the given IntegerRange and assigns it to the AsinCounts field.
func (o *CategoryItemWithAsinCounts) SetAsinCounts(v IntegerRange) {
	o.AsinCounts = &v
}

// GetParentCategoryId returns the ParentCategoryId field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCounts) GetParentCategoryId() string {
	if o == nil || IsNil(o.ParentCategoryId) {
		var ret string
		return ret
	}
	return *o.ParentCategoryId
}

// GetParentCategoryIdOk returns a tuple with the ParentCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCounts) GetParentCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCategoryId) {
		return nil, false
	}
	return o.ParentCategoryId, true
}

// HasParentCategoryId returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCounts) HasParentCategoryId() bool {
	if o != nil && !IsNil(o.ParentCategoryId) {
		return true
	}

	return false
}

// SetParentCategoryId gets a reference to the given string and assigns it to the ParentCategoryId field.
func (o *CategoryItemWithAsinCounts) SetParentCategoryId(v string) {
	o.ParentCategoryId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCounts) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCounts) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCounts) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CategoryItemWithAsinCounts) SetId(v string) {
	o.Id = &v
}

func (o CategoryItemWithAsinCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryPath) {
		toSerialize["categoryPath"] = o.CategoryPath
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AsinCounts) {
		toSerialize["asinCounts"] = o.AsinCounts
	}
	if !IsNil(o.ParentCategoryId) {
		toSerialize["parentCategoryId"] = o.ParentCategoryId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCategoryItemWithAsinCounts struct {
	value *CategoryItemWithAsinCounts
	isSet bool
}

func (v NullableCategoryItemWithAsinCounts) Get() *CategoryItemWithAsinCounts {
	return v.value
}

func (v *NullableCategoryItemWithAsinCounts) Set(val *CategoryItemWithAsinCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryItemWithAsinCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryItemWithAsinCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryItemWithAsinCounts(val *CategoryItemWithAsinCounts) *NullableCategoryItemWithAsinCounts {
	return &NullableCategoryItemWithAsinCounts{value: val, isSet: true}
}

func (v NullableCategoryItemWithAsinCounts) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCategoryItemWithAsinCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
