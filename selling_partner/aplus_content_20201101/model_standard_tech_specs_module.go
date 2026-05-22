package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardTechSpecsModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardTechSpecsModule{}

// StandardTechSpecsModule The standard table of technical feature names and definitions.
type StandardTechSpecsModule struct {
	Headline *TextComponent `json:"headline,omitempty"`
	// The specification list.
	SpecificationList []StandardTextPairBlock `json:"specificationList"`
	// The number of tables you want present. Features are evenly divided between the tables.
	TableCount *int32 `json:"tableCount,omitempty"`
}

type _StandardTechSpecsModule StandardTechSpecsModule

// NewStandardTechSpecsModule instantiates a new StandardTechSpecsModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardTechSpecsModule(specificationList []StandardTextPairBlock) *StandardTechSpecsModule {
	this := StandardTechSpecsModule{}
	this.SpecificationList = specificationList
	return &this
}

// NewStandardTechSpecsModuleWithDefaults instantiates a new StandardTechSpecsModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardTechSpecsModuleWithDefaults() *StandardTechSpecsModule {
	this := StandardTechSpecsModule{}
	return &this
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *StandardTechSpecsModule) GetHeadline() TextComponent {
	if o == nil || IsNil(o.Headline) {
		var ret TextComponent
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardTechSpecsModule) GetHeadlineOk() (*TextComponent, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *StandardTechSpecsModule) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given TextComponent and assigns it to the Headline field.
func (o *StandardTechSpecsModule) SetHeadline(v TextComponent) {
	o.Headline = &v
}

// GetSpecificationList returns the SpecificationList field value
func (o *StandardTechSpecsModule) GetSpecificationList() []StandardTextPairBlock {
	if o == nil {
		var ret []StandardTextPairBlock
		return ret
	}

	return o.SpecificationList
}

// GetSpecificationListOk returns a tuple with the SpecificationList field value
// and a boolean to check if the value has been set.
func (o *StandardTechSpecsModule) GetSpecificationListOk() ([]StandardTextPairBlock, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpecificationList, true
}

// SetSpecificationList sets field value
func (o *StandardTechSpecsModule) SetSpecificationList(v []StandardTextPairBlock) {
	o.SpecificationList = v
}

// GetTableCount returns the TableCount field value if set, zero value otherwise.
func (o *StandardTechSpecsModule) GetTableCount() int32 {
	if o == nil || IsNil(o.TableCount) {
		var ret int32
		return ret
	}
	return *o.TableCount
}

// GetTableCountOk returns a tuple with the TableCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardTechSpecsModule) GetTableCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TableCount) {
		return nil, false
	}
	return o.TableCount, true
}

// HasTableCount returns a boolean if a field has been set.
func (o *StandardTechSpecsModule) HasTableCount() bool {
	if o != nil && !IsNil(o.TableCount) {
		return true
	}

	return false
}

// SetTableCount gets a reference to the given int32 and assigns it to the TableCount field.
func (o *StandardTechSpecsModule) SetTableCount(v int32) {
	o.TableCount = &v
}

func (o StandardTechSpecsModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	toSerialize["specificationList"] = o.SpecificationList
	if !IsNil(o.TableCount) {
		toSerialize["tableCount"] = o.TableCount
	}
	return toSerialize, nil
}

type NullableStandardTechSpecsModule struct {
	value *StandardTechSpecsModule
	isSet bool
}

func (v NullableStandardTechSpecsModule) Get() *StandardTechSpecsModule {
	return v.value
}

func (v *NullableStandardTechSpecsModule) Set(val *StandardTechSpecsModule) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardTechSpecsModule) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardTechSpecsModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardTechSpecsModule(val *StandardTechSpecsModule) *NullableStandardTechSpecsModule {
	return &NullableStandardTechSpecsModule{value: val, isSet: true}
}

func (v NullableStandardTechSpecsModule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardTechSpecsModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
