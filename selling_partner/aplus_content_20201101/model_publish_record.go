package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the PublishRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublishRecord{}

// PublishRecord The full context for an A+ Content publishing event.
type PublishRecord struct {
	// The marketplace ID is the globally unique identifier of a marketplace. To find the ID for your marketplace, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	MarketplaceId string `json:"marketplaceId"`
	// The IETF language tag, which supports the primary language subtag and one secondary language subtag. The secondary language subtag is usually a regional designation. This doesn't support subtags other than the primary and secondary subtags. **Pattern:** ^[a-z]{2,}-[A-Z0-9]{2,}$
	Locale string `json:"locale"`
	// The Amazon Standard Identification Number (ASIN).
	Asin        string      `json:"asin"`
	ContentType ContentType `json:"contentType"`
	// The A+ Content document subtype. This represents a special-purpose type of an A+ Content document. Not every A+ Content document type has a subtype, and subtypes can change at any time.
	ContentSubType *string `json:"contentSubType,omitempty"`
	// A unique reference key for the A+ Content document. A content reference key cannot form a permalink and might change in the future. A content reference key is not guaranteed to match any A+ content identifier.
	ContentReferenceKey string `json:"contentReferenceKey"`
}

type _PublishRecord PublishRecord

// NewPublishRecord instantiates a new PublishRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishRecord(marketplaceId string, locale string, asin string, contentType ContentType, contentReferenceKey string) *PublishRecord {
	this := PublishRecord{}
	this.MarketplaceId = marketplaceId
	this.Locale = locale
	this.Asin = asin
	this.ContentType = contentType
	this.ContentReferenceKey = contentReferenceKey
	return &this
}

// NewPublishRecordWithDefaults instantiates a new PublishRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishRecordWithDefaults() *PublishRecord {
	this := PublishRecord{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *PublishRecord) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *PublishRecord) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *PublishRecord) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetLocale returns the Locale field value
func (o *PublishRecord) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *PublishRecord) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *PublishRecord) SetLocale(v string) {
	o.Locale = v
}

// GetAsin returns the Asin field value
func (o *PublishRecord) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *PublishRecord) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *PublishRecord) SetAsin(v string) {
	o.Asin = v
}

// GetContentType returns the ContentType field value
func (o *PublishRecord) GetContentType() ContentType {
	if o == nil {
		var ret ContentType
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *PublishRecord) GetContentTypeOk() (*ContentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *PublishRecord) SetContentType(v ContentType) {
	o.ContentType = v
}

// GetContentSubType returns the ContentSubType field value if set, zero value otherwise.
func (o *PublishRecord) GetContentSubType() string {
	if o == nil || IsNil(o.ContentSubType) {
		var ret string
		return ret
	}
	return *o.ContentSubType
}

// GetContentSubTypeOk returns a tuple with the ContentSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishRecord) GetContentSubTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentSubType) {
		return nil, false
	}
	return o.ContentSubType, true
}

// HasContentSubType returns a boolean if a field has been set.
func (o *PublishRecord) HasContentSubType() bool {
	if o != nil && !IsNil(o.ContentSubType) {
		return true
	}

	return false
}

// SetContentSubType gets a reference to the given string and assigns it to the ContentSubType field.
func (o *PublishRecord) SetContentSubType(v string) {
	o.ContentSubType = &v
}

// GetContentReferenceKey returns the ContentReferenceKey field value
func (o *PublishRecord) GetContentReferenceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentReferenceKey
}

// GetContentReferenceKeyOk returns a tuple with the ContentReferenceKey field value
// and a boolean to check if the value has been set.
func (o *PublishRecord) GetContentReferenceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentReferenceKey, true
}

// SetContentReferenceKey sets field value
func (o *PublishRecord) SetContentReferenceKey(v string) {
	o.ContentReferenceKey = v
}

func (o PublishRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["locale"] = o.Locale
	toSerialize["asin"] = o.Asin
	toSerialize["contentType"] = o.ContentType
	if !IsNil(o.ContentSubType) {
		toSerialize["contentSubType"] = o.ContentSubType
	}
	toSerialize["contentReferenceKey"] = o.ContentReferenceKey
	return toSerialize, nil
}

type NullablePublishRecord struct {
	value *PublishRecord
	isSet bool
}

func (v NullablePublishRecord) Get() *PublishRecord {
	return v.value
}

func (v *NullablePublishRecord) Set(val *PublishRecord) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishRecord) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishRecord(val *PublishRecord) *NullablePublishRecord {
	return &NullablePublishRecord{value: val, isSet: true}
}

func (v NullablePublishRecord) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePublishRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
