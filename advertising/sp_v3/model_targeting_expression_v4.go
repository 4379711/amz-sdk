package sp_v3

import "github.com/bytedance/sonic"

// checks if the TargetingExpressionV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingExpressionV4{}

// TargetingExpressionV4 The targeting expression. The `type` property specifies the targeting option. Use `CLOSE_MATCH` to match your auto targeting ads closely to the specified value. Use `LOOSE_MATCH` to match your auto targeting ads broadly to the specified value. Use `SUBSTITUTES` to display your auto targeting ads along with substitutable products. Use `COMPLEMENTS` to display your auto targeting ads along with affiliated products. Use `KEYWORD_BROAD_MATCH` to broadly match your keyword targeting ads with search queries. Use `KEYWORD_EXACT_MATCH` to exactly match your keyword targeting ads with search queries. Use `KEYWORD_PHRASE_MATCH` to match your keyword targeting ads with search phrases. your keyword targeting ads with search phrases. Use `PAT_ASIN` to match your product attribute targeting ads with product ASIN. Use `PAT_CATEGORY` to match your product attribute targeting ads with product category. Use `PAT_CATEGORY_REFINEMENT` to match your product attribute targeting ads with product attribute, including brand, price, rating, prime shipping eligible, age range and genre. Use `KEYWORD_GROUP` to match your keyword targeting ads with keyword group.
type TargetingExpressionV4 struct {
	Type string `json:"type"`
	// The targeting expression value.
	Value *string `json:"value,omitempty"`
}

type _TargetingExpressionV4 TargetingExpressionV4

// NewTargetingExpressionV4 instantiates a new TargetingExpressionV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingExpressionV4(type_ string) *TargetingExpressionV4 {
	this := TargetingExpressionV4{}
	this.Type = type_
	return &this
}

// NewTargetingExpressionV4WithDefaults instantiates a new TargetingExpressionV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingExpressionV4WithDefaults() *TargetingExpressionV4 {
	this := TargetingExpressionV4{}
	return &this
}

// GetType returns the Type field value
func (o *TargetingExpressionV4) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TargetingExpressionV4) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TargetingExpressionV4) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TargetingExpressionV4) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingExpressionV4) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TargetingExpressionV4) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TargetingExpressionV4) SetValue(v string) {
	o.Value = &v
}

func (o TargetingExpressionV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTargetingExpressionV4 struct {
	value *TargetingExpressionV4
	isSet bool
}

func (v NullableTargetingExpressionV4) Get() *TargetingExpressionV4 {
	return v.value
}

func (v *NullableTargetingExpressionV4) Set(val *TargetingExpressionV4) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingExpressionV4) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingExpressionV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingExpressionV4(val *TargetingExpressionV4) *NullableTargetingExpressionV4 {
	return &NullableTargetingExpressionV4{value: val, isSet: true}
}

func (v NullableTargetingExpressionV4) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingExpressionV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
