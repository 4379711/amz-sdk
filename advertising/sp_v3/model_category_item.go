package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CategoryItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryItem{}

// CategoryItem struct for CategoryItem
type CategoryItem struct {
	// The category id of the parent node
	Parent *string `json:"parent,omitempty"`
	// The path of the category, which contains the current category and all parent categories
	Path *string `json:"path,omitempty"`
	// A flag which indicates if the current node may be targeted
	CanBeTargeted *bool `json:"canBeTargeted,omitempty"`
	// The name of the category
	Name *string `json:"name,omitempty"`
	// The category id of the current node
	Id *string `json:"id,omitempty"`
}

// NewCategoryItem instantiates a new CategoryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryItem() *CategoryItem {
	this := CategoryItem{}
	return &this
}

// NewCategoryItemWithDefaults instantiates a new CategoryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryItemWithDefaults() *CategoryItem {
	this := CategoryItem{}
	return &this
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *CategoryItem) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItem) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *CategoryItem) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *CategoryItem) SetParent(v string) {
	o.Parent = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CategoryItem) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItem) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CategoryItem) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CategoryItem) SetPath(v string) {
	o.Path = &v
}

// GetCanBeTargeted returns the CanBeTargeted field value if set, zero value otherwise.
func (o *CategoryItem) GetCanBeTargeted() bool {
	if o == nil || IsNil(o.CanBeTargeted) {
		var ret bool
		return ret
	}
	return *o.CanBeTargeted
}

// GetCanBeTargetedOk returns a tuple with the CanBeTargeted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItem) GetCanBeTargetedOk() (*bool, bool) {
	if o == nil || IsNil(o.CanBeTargeted) {
		return nil, false
	}
	return o.CanBeTargeted, true
}

// HasCanBeTargeted returns a boolean if a field has been set.
func (o *CategoryItem) HasCanBeTargeted() bool {
	if o != nil && !IsNil(o.CanBeTargeted) {
		return true
	}

	return false
}

// SetCanBeTargeted gets a reference to the given bool and assigns it to the CanBeTargeted field.
func (o *CategoryItem) SetCanBeTargeted(v bool) {
	o.CanBeTargeted = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CategoryItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CategoryItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CategoryItem) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CategoryItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CategoryItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CategoryItem) SetId(v string) {
	o.Id = &v
}

func (o CategoryItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.CanBeTargeted) {
		toSerialize["canBeTargeted"] = o.CanBeTargeted
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCategoryItem struct {
	value *CategoryItem
	isSet bool
}

func (v NullableCategoryItem) Get() *CategoryItem {
	return v.value
}

func (v *NullableCategoryItem) Set(val *CategoryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryItem(val *CategoryItem) *NullableCategoryItem {
	return &NullableCategoryItem{value: val, isSet: true}
}

func (v NullableCategoryItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCategoryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
