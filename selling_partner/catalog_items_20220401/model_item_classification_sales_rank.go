package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemClassificationSalesRank type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemClassificationSalesRank{}

// ItemClassificationSalesRank Sales rank of an Amazon catalog item.
type ItemClassificationSalesRank struct {
	// Identifier of the classification that is associated with the sales rank.
	ClassificationId string `json:"classificationId"`
	// Name of the sales rank.
	Title string `json:"title"`
	// Corresponding Amazon retail website URL for the sales category.
	Link *string `json:"link,omitempty"`
	// Sales rank.
	Rank int32 `json:"rank"`
}

type _ItemClassificationSalesRank ItemClassificationSalesRank

// NewItemClassificationSalesRank instantiates a new ItemClassificationSalesRank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemClassificationSalesRank(classificationId string, title string, rank int32) *ItemClassificationSalesRank {
	this := ItemClassificationSalesRank{}
	this.ClassificationId = classificationId
	this.Title = title
	this.Rank = rank
	return &this
}

// NewItemClassificationSalesRankWithDefaults instantiates a new ItemClassificationSalesRank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemClassificationSalesRankWithDefaults() *ItemClassificationSalesRank {
	this := ItemClassificationSalesRank{}
	return &this
}

// GetClassificationId returns the ClassificationId field value
func (o *ItemClassificationSalesRank) GetClassificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassificationId
}

// GetClassificationIdOk returns a tuple with the ClassificationId field value
// and a boolean to check if the value has been set.
func (o *ItemClassificationSalesRank) GetClassificationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassificationId, true
}

// SetClassificationId sets field value
func (o *ItemClassificationSalesRank) SetClassificationId(v string) {
	o.ClassificationId = v
}

// GetTitle returns the Title field value
func (o *ItemClassificationSalesRank) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ItemClassificationSalesRank) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ItemClassificationSalesRank) SetTitle(v string) {
	o.Title = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ItemClassificationSalesRank) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemClassificationSalesRank) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *ItemClassificationSalesRank) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *ItemClassificationSalesRank) SetLink(v string) {
	o.Link = &v
}

// GetRank returns the Rank field value
func (o *ItemClassificationSalesRank) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *ItemClassificationSalesRank) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *ItemClassificationSalesRank) SetRank(v int32) {
	o.Rank = v
}

func (o ItemClassificationSalesRank) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["classificationId"] = o.ClassificationId
	toSerialize["title"] = o.Title
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	toSerialize["rank"] = o.Rank
	return toSerialize, nil
}

type NullableItemClassificationSalesRank struct {
	value *ItemClassificationSalesRank
	isSet bool
}

func (v NullableItemClassificationSalesRank) Get() *ItemClassificationSalesRank {
	return v.value
}

func (v *NullableItemClassificationSalesRank) Set(val *ItemClassificationSalesRank) {
	v.value = val
	v.isSet = true
}

func (v NullableItemClassificationSalesRank) IsSet() bool {
	return v.isSet
}

func (v *NullableItemClassificationSalesRank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemClassificationSalesRank(val *ItemClassificationSalesRank) *NullableItemClassificationSalesRank {
	return &NullableItemClassificationSalesRank{value: val, isSet: true}
}

func (v NullableItemClassificationSalesRank) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemClassificationSalesRank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
