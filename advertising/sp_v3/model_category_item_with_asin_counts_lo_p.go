package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CategoryItemWithAsinCountsLoP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryItemWithAsinCountsLoP{}

// CategoryItemWithAsinCountsLoP struct for CategoryItemWithAsinCountsLoP
type CategoryItemWithAsinCountsLoP struct {
	// The path of the category, which contains the current category and all parent categories
	CategoryPath *string `json:"categoryPath,omitempty"`
	// The name of the category
	Name *string `json:"name,omitempty"`
	// The translated path of the category, which contains the current category and all parent categories.
	TranslatedCategoryPath *string       `json:"translatedCategoryPath,omitempty"`
	AsinCounts             *IntegerRange `json:"asinCounts,omitempty"`
	// The category id of the parent node
	ParentCategoryId *string `json:"parentCategoryId,omitempty"`
	// The category id of the current node
	Id *string `json:"id,omitempty"`
	// The translated name of the category.
	TranslatedName *string `json:"translatedName,omitempty"`
}

// NewCategoryItemWithAsinCountsLoP instantiates a new CategoryItemWithAsinCountsLoP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryItemWithAsinCountsLoP() *CategoryItemWithAsinCountsLoP {
	this := CategoryItemWithAsinCountsLoP{}
	return &this
}

// NewCategoryItemWithAsinCountsLoPWithDefaults instantiates a new CategoryItemWithAsinCountsLoP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryItemWithAsinCountsLoPWithDefaults() *CategoryItemWithAsinCountsLoP {
	this := CategoryItemWithAsinCountsLoP{}
	return &this
}

// GetCategoryPath returns the CategoryPath field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCountsLoP) GetCategoryPath() string {
	if o == nil || IsNil(o.CategoryPath) {
		var ret string
		return ret
	}
	return *o.CategoryPath
}

// GetCategoryPathOk returns a tuple with the CategoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCountsLoP) GetCategoryPathOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryPath) {
		return nil, false
	}
	return o.CategoryPath, true
}

// HasCategoryPath returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCountsLoP) HasCategoryPath() bool {
	if o != nil && !IsNil(o.CategoryPath) {
		return true
	}

	return false
}

// SetCategoryPath gets a reference to the given string and assigns it to the CategoryPath field.
func (o *CategoryItemWithAsinCountsLoP) SetCategoryPath(v string) {
	o.CategoryPath = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCountsLoP) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCountsLoP) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCountsLoP) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CategoryItemWithAsinCountsLoP) SetName(v string) {
	o.Name = &v
}

// GetTranslatedCategoryPath returns the TranslatedCategoryPath field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCountsLoP) GetTranslatedCategoryPath() string {
	if o == nil || IsNil(o.TranslatedCategoryPath) {
		var ret string
		return ret
	}
	return *o.TranslatedCategoryPath
}

// GetTranslatedCategoryPathOk returns a tuple with the TranslatedCategoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCountsLoP) GetTranslatedCategoryPathOk() (*string, bool) {
	if o == nil || IsNil(o.TranslatedCategoryPath) {
		return nil, false
	}
	return o.TranslatedCategoryPath, true
}

// HasTranslatedCategoryPath returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCountsLoP) HasTranslatedCategoryPath() bool {
	if o != nil && !IsNil(o.TranslatedCategoryPath) {
		return true
	}

	return false
}

// SetTranslatedCategoryPath gets a reference to the given string and assigns it to the TranslatedCategoryPath field.
func (o *CategoryItemWithAsinCountsLoP) SetTranslatedCategoryPath(v string) {
	o.TranslatedCategoryPath = &v
}

// GetAsinCounts returns the AsinCounts field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCountsLoP) GetAsinCounts() IntegerRange {
	if o == nil || IsNil(o.AsinCounts) {
		var ret IntegerRange
		return ret
	}
	return *o.AsinCounts
}

// GetAsinCountsOk returns a tuple with the AsinCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCountsLoP) GetAsinCountsOk() (*IntegerRange, bool) {
	if o == nil || IsNil(o.AsinCounts) {
		return nil, false
	}
	return o.AsinCounts, true
}

// HasAsinCounts returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCountsLoP) HasAsinCounts() bool {
	if o != nil && !IsNil(o.AsinCounts) {
		return true
	}

	return false
}

// SetAsinCounts gets a reference to the given IntegerRange and assigns it to the AsinCounts field.
func (o *CategoryItemWithAsinCountsLoP) SetAsinCounts(v IntegerRange) {
	o.AsinCounts = &v
}

// GetParentCategoryId returns the ParentCategoryId field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCountsLoP) GetParentCategoryId() string {
	if o == nil || IsNil(o.ParentCategoryId) {
		var ret string
		return ret
	}
	return *o.ParentCategoryId
}

// GetParentCategoryIdOk returns a tuple with the ParentCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCountsLoP) GetParentCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCategoryId) {
		return nil, false
	}
	return o.ParentCategoryId, true
}

// HasParentCategoryId returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCountsLoP) HasParentCategoryId() bool {
	if o != nil && !IsNil(o.ParentCategoryId) {
		return true
	}

	return false
}

// SetParentCategoryId gets a reference to the given string and assigns it to the ParentCategoryId field.
func (o *CategoryItemWithAsinCountsLoP) SetParentCategoryId(v string) {
	o.ParentCategoryId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCountsLoP) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCountsLoP) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCountsLoP) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CategoryItemWithAsinCountsLoP) SetId(v string) {
	o.Id = &v
}

// GetTranslatedName returns the TranslatedName field value if set, zero value otherwise.
func (o *CategoryItemWithAsinCountsLoP) GetTranslatedName() string {
	if o == nil || IsNil(o.TranslatedName) {
		var ret string
		return ret
	}
	return *o.TranslatedName
}

// GetTranslatedNameOk returns a tuple with the TranslatedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItemWithAsinCountsLoP) GetTranslatedNameOk() (*string, bool) {
	if o == nil || IsNil(o.TranslatedName) {
		return nil, false
	}
	return o.TranslatedName, true
}

// HasTranslatedName returns a boolean if a field has been set.
func (o *CategoryItemWithAsinCountsLoP) HasTranslatedName() bool {
	if o != nil && !IsNil(o.TranslatedName) {
		return true
	}

	return false
}

// SetTranslatedName gets a reference to the given string and assigns it to the TranslatedName field.
func (o *CategoryItemWithAsinCountsLoP) SetTranslatedName(v string) {
	o.TranslatedName = &v
}

func (o CategoryItemWithAsinCountsLoP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryPath) {
		toSerialize["categoryPath"] = o.CategoryPath
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TranslatedCategoryPath) {
		toSerialize["translatedCategoryPath"] = o.TranslatedCategoryPath
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
	if !IsNil(o.TranslatedName) {
		toSerialize["translatedName"] = o.TranslatedName
	}
	return toSerialize, nil
}

type NullableCategoryItemWithAsinCountsLoP struct {
	value *CategoryItemWithAsinCountsLoP
	isSet bool
}

func (v NullableCategoryItemWithAsinCountsLoP) Get() *CategoryItemWithAsinCountsLoP {
	return v.value
}

func (v *NullableCategoryItemWithAsinCountsLoP) Set(val *CategoryItemWithAsinCountsLoP) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryItemWithAsinCountsLoP) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryItemWithAsinCountsLoP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryItemWithAsinCountsLoP(val *CategoryItemWithAsinCountsLoP) *NullableCategoryItemWithAsinCountsLoP {
	return &NullableCategoryItemWithAsinCountsLoP{value: val, isSet: true}
}

func (v NullableCategoryItemWithAsinCountsLoP) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCategoryItemWithAsinCountsLoP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
