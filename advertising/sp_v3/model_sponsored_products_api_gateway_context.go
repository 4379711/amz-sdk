package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsApiGatewayContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsApiGatewayContext{}

// SponsoredProductsApiGatewayContext struct for SponsoredProductsApiGatewayContext
type SponsoredProductsApiGatewayContext struct {
	MarketplaceId *string `json:"marketplaceId,omitempty"`
	ClientId      *string `json:"clientId,omitempty"`
	RequestId     *string `json:"requestId,omitempty"`
	EntityType    *string `json:"entityType,omitempty"`
	EntityId      *string `json:"entityId,omitempty"`
	ContentType   *string `json:"contentType,omitempty"`
	AdvertiserId  *string `json:"advertiserId,omitempty"`
}

// NewSponsoredProductsApiGatewayContext instantiates a new SponsoredProductsApiGatewayContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsApiGatewayContext() *SponsoredProductsApiGatewayContext {
	this := SponsoredProductsApiGatewayContext{}
	return &this
}

// NewSponsoredProductsApiGatewayContextWithDefaults instantiates a new SponsoredProductsApiGatewayContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsApiGatewayContextWithDefaults() *SponsoredProductsApiGatewayContext {
	this := SponsoredProductsApiGatewayContext{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *SponsoredProductsApiGatewayContext) GetMarketplaceId() string {
	if o == nil || IsNil(o.MarketplaceId) {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsApiGatewayContext) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceId) {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *SponsoredProductsApiGatewayContext) HasMarketplaceId() bool {
	if o != nil && !IsNil(o.MarketplaceId) {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *SponsoredProductsApiGatewayContext) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SponsoredProductsApiGatewayContext) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsApiGatewayContext) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SponsoredProductsApiGatewayContext) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SponsoredProductsApiGatewayContext) SetClientId(v string) {
	o.ClientId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *SponsoredProductsApiGatewayContext) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsApiGatewayContext) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *SponsoredProductsApiGatewayContext) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *SponsoredProductsApiGatewayContext) SetRequestId(v string) {
	o.RequestId = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *SponsoredProductsApiGatewayContext) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsApiGatewayContext) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *SponsoredProductsApiGatewayContext) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *SponsoredProductsApiGatewayContext) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SponsoredProductsApiGatewayContext) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsApiGatewayContext) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SponsoredProductsApiGatewayContext) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *SponsoredProductsApiGatewayContext) SetEntityId(v string) {
	o.EntityId = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *SponsoredProductsApiGatewayContext) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsApiGatewayContext) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *SponsoredProductsApiGatewayContext) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *SponsoredProductsApiGatewayContext) SetContentType(v string) {
	o.ContentType = &v
}

// GetAdvertiserId returns the AdvertiserId field value if set, zero value otherwise.
func (o *SponsoredProductsApiGatewayContext) GetAdvertiserId() string {
	if o == nil || IsNil(o.AdvertiserId) {
		var ret string
		return ret
	}
	return *o.AdvertiserId
}

// GetAdvertiserIdOk returns a tuple with the AdvertiserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsApiGatewayContext) GetAdvertiserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdvertiserId) {
		return nil, false
	}
	return o.AdvertiserId, true
}

// HasAdvertiserId returns a boolean if a field has been set.
func (o *SponsoredProductsApiGatewayContext) HasAdvertiserId() bool {
	if o != nil && !IsNil(o.AdvertiserId) {
		return true
	}

	return false
}

// SetAdvertiserId gets a reference to the given string and assigns it to the AdvertiserId field.
func (o *SponsoredProductsApiGatewayContext) SetAdvertiserId(v string) {
	o.AdvertiserId = &v
}

func (o SponsoredProductsApiGatewayContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceId) {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !IsNil(o.AdvertiserId) {
		toSerialize["advertiserId"] = o.AdvertiserId
	}
	return toSerialize, nil
}

type NullableSponsoredProductsApiGatewayContext struct {
	value *SponsoredProductsApiGatewayContext
	isSet bool
}

func (v NullableSponsoredProductsApiGatewayContext) Get() *SponsoredProductsApiGatewayContext {
	return v.value
}

func (v *NullableSponsoredProductsApiGatewayContext) Set(val *SponsoredProductsApiGatewayContext) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsApiGatewayContext) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsApiGatewayContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsApiGatewayContext(val *SponsoredProductsApiGatewayContext) *NullableSponsoredProductsApiGatewayContext {
	return &NullableSponsoredProductsApiGatewayContext{value: val, isSet: true}
}

func (v NullableSponsoredProductsApiGatewayContext) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsApiGatewayContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
