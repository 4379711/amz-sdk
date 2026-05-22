package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the AsinMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsinMetadata{}

// AsinMetadata The A+ Content ASIN with additional metadata for content management. If you don't include the `includedDataSet` parameter in a call to the `listContentDocumentAsinRelations` operation, the related ASINs are returned without metadata.
type AsinMetadata struct {
	// The Amazon Standard Identification Number (ASIN).
	Asin string `json:"asin"`
	// The set of ASIN badges.
	BadgeSet []AsinBadge `json:"badgeSet,omitempty"`
	// The Amazon Standard Identification Number (ASIN).
	Parent *string `json:"parent,omitempty"`
	// The title for the ASIN in the Amazon catalog.
	Title *string `json:"title,omitempty"`
	// The default image for the ASIN in the Amazon catalog.
	ImageUrl *string `json:"imageUrl,omitempty"`
	// A set of content reference keys.
	ContentReferenceKeySet []string `json:"contentReferenceKeySet,omitempty"`
}

type _AsinMetadata AsinMetadata

// NewAsinMetadata instantiates a new AsinMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsinMetadata(asin string) *AsinMetadata {
	this := AsinMetadata{}
	this.Asin = asin
	return &this
}

// NewAsinMetadataWithDefaults instantiates a new AsinMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsinMetadataWithDefaults() *AsinMetadata {
	this := AsinMetadata{}
	return &this
}

// GetAsin returns the Asin field value
func (o *AsinMetadata) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *AsinMetadata) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *AsinMetadata) SetAsin(v string) {
	o.Asin = v
}

// GetBadgeSet returns the BadgeSet field value if set, zero value otherwise.
func (o *AsinMetadata) GetBadgeSet() []AsinBadge {
	if o == nil || IsNil(o.BadgeSet) {
		var ret []AsinBadge
		return ret
	}
	return o.BadgeSet
}

// GetBadgeSetOk returns a tuple with the BadgeSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinMetadata) GetBadgeSetOk() ([]AsinBadge, bool) {
	if o == nil || IsNil(o.BadgeSet) {
		return nil, false
	}
	return o.BadgeSet, true
}

// HasBadgeSet returns a boolean if a field has been set.
func (o *AsinMetadata) HasBadgeSet() bool {
	if o != nil && !IsNil(o.BadgeSet) {
		return true
	}

	return false
}

// SetBadgeSet gets a reference to the given []AsinBadge and assigns it to the BadgeSet field.
func (o *AsinMetadata) SetBadgeSet(v []AsinBadge) {
	o.BadgeSet = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *AsinMetadata) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinMetadata) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *AsinMetadata) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *AsinMetadata) SetParent(v string) {
	o.Parent = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AsinMetadata) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinMetadata) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AsinMetadata) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AsinMetadata) SetTitle(v string) {
	o.Title = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *AsinMetadata) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinMetadata) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *AsinMetadata) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *AsinMetadata) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetContentReferenceKeySet returns the ContentReferenceKeySet field value if set, zero value otherwise.
func (o *AsinMetadata) GetContentReferenceKeySet() []string {
	if o == nil || IsNil(o.ContentReferenceKeySet) {
		var ret []string
		return ret
	}
	return o.ContentReferenceKeySet
}

// GetContentReferenceKeySetOk returns a tuple with the ContentReferenceKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinMetadata) GetContentReferenceKeySetOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentReferenceKeySet) {
		return nil, false
	}
	return o.ContentReferenceKeySet, true
}

// HasContentReferenceKeySet returns a boolean if a field has been set.
func (o *AsinMetadata) HasContentReferenceKeySet() bool {
	if o != nil && !IsNil(o.ContentReferenceKeySet) {
		return true
	}

	return false
}

// SetContentReferenceKeySet gets a reference to the given []string and assigns it to the ContentReferenceKeySet field.
func (o *AsinMetadata) SetContentReferenceKeySet(v []string) {
	o.ContentReferenceKeySet = v
}

func (o AsinMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asin"] = o.Asin
	if !IsNil(o.BadgeSet) {
		toSerialize["badgeSet"] = o.BadgeSet
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.ContentReferenceKeySet) {
		toSerialize["contentReferenceKeySet"] = o.ContentReferenceKeySet
	}
	return toSerialize, nil
}

type NullableAsinMetadata struct {
	value *AsinMetadata
	isSet bool
}

func (v NullableAsinMetadata) Get() *AsinMetadata {
	return v.value
}

func (v *NullableAsinMetadata) Set(val *AsinMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAsinMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAsinMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsinMetadata(val *AsinMetadata) *NullableAsinMetadata {
	return &NullableAsinMetadata{value: val, isSet: true}
}

func (v NullableAsinMetadata) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsinMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
