package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the ContentUpdatePreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentUpdatePreview{}

// ContentUpdatePreview Preview of the changes that will be applied to the shipment.
type ContentUpdatePreview struct {
	// Identifier of a content update preview.
	ContentUpdatePreviewId string `json:"contentUpdatePreviewId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The time at which the content update expires. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format with pattern `yyyy-MM-ddTHH:mm:ss.sssZ`.
	Expiration           time.Time            `json:"expiration"`
	RequestedUpdates     RequestedUpdates     `json:"requestedUpdates"`
	TransportationOption TransportationOption `json:"transportationOption"`
}

type _ContentUpdatePreview ContentUpdatePreview

// NewContentUpdatePreview instantiates a new ContentUpdatePreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentUpdatePreview(contentUpdatePreviewId string, expiration time.Time, requestedUpdates RequestedUpdates, transportationOption TransportationOption) *ContentUpdatePreview {
	this := ContentUpdatePreview{}
	this.ContentUpdatePreviewId = contentUpdatePreviewId
	this.Expiration = expiration
	this.RequestedUpdates = requestedUpdates
	this.TransportationOption = transportationOption
	return &this
}

// NewContentUpdatePreviewWithDefaults instantiates a new ContentUpdatePreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentUpdatePreviewWithDefaults() *ContentUpdatePreview {
	this := ContentUpdatePreview{}
	return &this
}

// GetContentUpdatePreviewId returns the ContentUpdatePreviewId field value
func (o *ContentUpdatePreview) GetContentUpdatePreviewId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentUpdatePreviewId
}

// GetContentUpdatePreviewIdOk returns a tuple with the ContentUpdatePreviewId field value
// and a boolean to check if the value has been set.
func (o *ContentUpdatePreview) GetContentUpdatePreviewIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentUpdatePreviewId, true
}

// SetContentUpdatePreviewId sets field value
func (o *ContentUpdatePreview) SetContentUpdatePreviewId(v string) {
	o.ContentUpdatePreviewId = v
}

// GetExpiration returns the Expiration field value
func (o *ContentUpdatePreview) GetExpiration() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *ContentUpdatePreview) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *ContentUpdatePreview) SetExpiration(v time.Time) {
	o.Expiration = v
}

// GetRequestedUpdates returns the RequestedUpdates field value
func (o *ContentUpdatePreview) GetRequestedUpdates() RequestedUpdates {
	if o == nil {
		var ret RequestedUpdates
		return ret
	}

	return o.RequestedUpdates
}

// GetRequestedUpdatesOk returns a tuple with the RequestedUpdates field value
// and a boolean to check if the value has been set.
func (o *ContentUpdatePreview) GetRequestedUpdatesOk() (*RequestedUpdates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedUpdates, true
}

// SetRequestedUpdates sets field value
func (o *ContentUpdatePreview) SetRequestedUpdates(v RequestedUpdates) {
	o.RequestedUpdates = v
}

// GetTransportationOption returns the TransportationOption field value
func (o *ContentUpdatePreview) GetTransportationOption() TransportationOption {
	if o == nil {
		var ret TransportationOption
		return ret
	}

	return o.TransportationOption
}

// GetTransportationOptionOk returns a tuple with the TransportationOption field value
// and a boolean to check if the value has been set.
func (o *ContentUpdatePreview) GetTransportationOptionOk() (*TransportationOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportationOption, true
}

// SetTransportationOption sets field value
func (o *ContentUpdatePreview) SetTransportationOption(v TransportationOption) {
	o.TransportationOption = v
}

func (o ContentUpdatePreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contentUpdatePreviewId"] = o.ContentUpdatePreviewId
	toSerialize["expiration"] = o.Expiration
	toSerialize["requestedUpdates"] = o.RequestedUpdates
	toSerialize["transportationOption"] = o.TransportationOption
	return toSerialize, nil
}

type NullableContentUpdatePreview struct {
	value *ContentUpdatePreview
	isSet bool
}

func (v NullableContentUpdatePreview) Get() *ContentUpdatePreview {
	return v.value
}

func (v *NullableContentUpdatePreview) Set(val *ContentUpdatePreview) {
	v.value = val
	v.isSet = true
}

func (v NullableContentUpdatePreview) IsSet() bool {
	return v.isSet
}

func (v *NullableContentUpdatePreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentUpdatePreview(val *ContentUpdatePreview) *NullableContentUpdatePreview {
	return &NullableContentUpdatePreview{value: val, isSet: true}
}

func (v NullableContentUpdatePreview) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContentUpdatePreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
