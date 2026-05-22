package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the KeywordGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordGroup{}

// KeywordGroup Keyword group. Represents a high level keyword targeting intent. e.g. the keyword group \"gift\" can target relevant search queries containing the word gift
type KeywordGroup struct {
	// Summary of impacts.
	ImpactSummary *string `json:"impactSummary,omitempty"`
	// Detailed Keyword group description.
	Description *string `json:"description,omitempty"`
	// Sample keywords that match the group.
	SampleKeywords []string `json:"sampleKeywords,omitempty"`
	// Unique Identifier for the keyword group. To be passed during targeting clause creation.
	Id string `json:"id"`
	// Keyword group text. Can be used for display purposes.
	Text string `json:"text"`
}

type _KeywordGroup KeywordGroup

// NewKeywordGroup instantiates a new KeywordGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordGroup(id string, text string) *KeywordGroup {
	this := KeywordGroup{}
	this.Id = id
	this.Text = text
	return &this
}

// NewKeywordGroupWithDefaults instantiates a new KeywordGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordGroupWithDefaults() *KeywordGroup {
	this := KeywordGroup{}
	return &this
}

// GetImpactSummary returns the ImpactSummary field value if set, zero value otherwise.
func (o *KeywordGroup) GetImpactSummary() string {
	if o == nil || IsNil(o.ImpactSummary) {
		var ret string
		return ret
	}
	return *o.ImpactSummary
}

// GetImpactSummaryOk returns a tuple with the ImpactSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordGroup) GetImpactSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.ImpactSummary) {
		return nil, false
	}
	return o.ImpactSummary, true
}

// HasImpactSummary returns a boolean if a field has been set.
func (o *KeywordGroup) HasImpactSummary() bool {
	if o != nil && !IsNil(o.ImpactSummary) {
		return true
	}

	return false
}

// SetImpactSummary gets a reference to the given string and assigns it to the ImpactSummary field.
func (o *KeywordGroup) SetImpactSummary(v string) {
	o.ImpactSummary = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KeywordGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KeywordGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KeywordGroup) SetDescription(v string) {
	o.Description = &v
}

// GetSampleKeywords returns the SampleKeywords field value if set, zero value otherwise.
func (o *KeywordGroup) GetSampleKeywords() []string {
	if o == nil || IsNil(o.SampleKeywords) {
		var ret []string
		return ret
	}
	return o.SampleKeywords
}

// GetSampleKeywordsOk returns a tuple with the SampleKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordGroup) GetSampleKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.SampleKeywords) {
		return nil, false
	}
	return o.SampleKeywords, true
}

// HasSampleKeywords returns a boolean if a field has been set.
func (o *KeywordGroup) HasSampleKeywords() bool {
	if o != nil && !IsNil(o.SampleKeywords) {
		return true
	}

	return false
}

// SetSampleKeywords gets a reference to the given []string and assigns it to the SampleKeywords field.
func (o *KeywordGroup) SetSampleKeywords(v []string) {
	o.SampleKeywords = v
}

// GetId returns the Id field value
func (o *KeywordGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KeywordGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KeywordGroup) SetId(v string) {
	o.Id = v
}

// GetText returns the Text field value
func (o *KeywordGroup) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *KeywordGroup) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *KeywordGroup) SetText(v string) {
	o.Text = v
}

func (o KeywordGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImpactSummary) {
		toSerialize["impactSummary"] = o.ImpactSummary
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SampleKeywords) {
		toSerialize["sampleKeywords"] = o.SampleKeywords
	}
	toSerialize["id"] = o.Id
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

type NullableKeywordGroup struct {
	value *KeywordGroup
	isSet bool
}

func (v NullableKeywordGroup) Get() *KeywordGroup {
	return v.value
}

func (v *NullableKeywordGroup) Set(val *KeywordGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordGroup(val *KeywordGroup) *NullableKeywordGroup {
	return &NullableKeywordGroup{value: val, isSet: true}
}

func (v NullableKeywordGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableKeywordGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
