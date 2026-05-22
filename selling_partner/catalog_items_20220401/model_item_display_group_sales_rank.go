package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemDisplayGroupSalesRank type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemDisplayGroupSalesRank{}

// ItemDisplayGroupSalesRank Sales rank of an Amazon catalog item, grouped by website display group.
type ItemDisplayGroupSalesRank struct {
	// Name of the website display group that is associated with the sales rank
	WebsiteDisplayGroup string `json:"websiteDisplayGroup"`
	// Name of the sales rank.
	Title string `json:"title"`
	// Corresponding Amazon retail website URL for the sales rank.
	Link *string `json:"link,omitempty"`
	// Sales rank.
	Rank int32 `json:"rank"`
}

type _ItemDisplayGroupSalesRank ItemDisplayGroupSalesRank

// NewItemDisplayGroupSalesRank instantiates a new ItemDisplayGroupSalesRank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemDisplayGroupSalesRank(websiteDisplayGroup string, title string, rank int32) *ItemDisplayGroupSalesRank {
	this := ItemDisplayGroupSalesRank{}
	this.WebsiteDisplayGroup = websiteDisplayGroup
	this.Title = title
	this.Rank = rank
	return &this
}

// NewItemDisplayGroupSalesRankWithDefaults instantiates a new ItemDisplayGroupSalesRank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemDisplayGroupSalesRankWithDefaults() *ItemDisplayGroupSalesRank {
	this := ItemDisplayGroupSalesRank{}
	return &this
}

// GetWebsiteDisplayGroup returns the WebsiteDisplayGroup field value
func (o *ItemDisplayGroupSalesRank) GetWebsiteDisplayGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebsiteDisplayGroup
}

// GetWebsiteDisplayGroupOk returns a tuple with the WebsiteDisplayGroup field value
// and a boolean to check if the value has been set.
func (o *ItemDisplayGroupSalesRank) GetWebsiteDisplayGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebsiteDisplayGroup, true
}

// SetWebsiteDisplayGroup sets field value
func (o *ItemDisplayGroupSalesRank) SetWebsiteDisplayGroup(v string) {
	o.WebsiteDisplayGroup = v
}

// GetTitle returns the Title field value
func (o *ItemDisplayGroupSalesRank) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ItemDisplayGroupSalesRank) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ItemDisplayGroupSalesRank) SetTitle(v string) {
	o.Title = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ItemDisplayGroupSalesRank) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDisplayGroupSalesRank) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *ItemDisplayGroupSalesRank) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *ItemDisplayGroupSalesRank) SetLink(v string) {
	o.Link = &v
}

// GetRank returns the Rank field value
func (o *ItemDisplayGroupSalesRank) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *ItemDisplayGroupSalesRank) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *ItemDisplayGroupSalesRank) SetRank(v int32) {
	o.Rank = v
}

func (o ItemDisplayGroupSalesRank) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["websiteDisplayGroup"] = o.WebsiteDisplayGroup
	toSerialize["title"] = o.Title
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	toSerialize["rank"] = o.Rank
	return toSerialize, nil
}

type NullableItemDisplayGroupSalesRank struct {
	value *ItemDisplayGroupSalesRank
	isSet bool
}

func (v NullableItemDisplayGroupSalesRank) Get() *ItemDisplayGroupSalesRank {
	return v.value
}

func (v *NullableItemDisplayGroupSalesRank) Set(val *ItemDisplayGroupSalesRank) {
	v.value = val
	v.isSet = true
}

func (v NullableItemDisplayGroupSalesRank) IsSet() bool {
	return v.isSet
}

func (v *NullableItemDisplayGroupSalesRank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemDisplayGroupSalesRank(val *ItemDisplayGroupSalesRank) *NullableItemDisplayGroupSalesRank {
	return &NullableItemDisplayGroupSalesRank{value: val, isSet: true}
}

func (v NullableItemDisplayGroupSalesRank) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemDisplayGroupSalesRank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
