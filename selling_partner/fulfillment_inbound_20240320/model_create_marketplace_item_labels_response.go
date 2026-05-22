package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateMarketplaceItemLabelsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMarketplaceItemLabelsResponse{}

// CreateMarketplaceItemLabelsResponse The `createMarketplaceItemLabels` response.
type CreateMarketplaceItemLabelsResponse struct {
	// Resources to download the requested document.
	DocumentDownloads []DocumentDownload `json:"documentDownloads"`
}

type _CreateMarketplaceItemLabelsResponse CreateMarketplaceItemLabelsResponse

// NewCreateMarketplaceItemLabelsResponse instantiates a new CreateMarketplaceItemLabelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMarketplaceItemLabelsResponse(documentDownloads []DocumentDownload) *CreateMarketplaceItemLabelsResponse {
	this := CreateMarketplaceItemLabelsResponse{}
	this.DocumentDownloads = documentDownloads
	return &this
}

// NewCreateMarketplaceItemLabelsResponseWithDefaults instantiates a new CreateMarketplaceItemLabelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMarketplaceItemLabelsResponseWithDefaults() *CreateMarketplaceItemLabelsResponse {
	this := CreateMarketplaceItemLabelsResponse{}
	return &this
}

// GetDocumentDownloads returns the DocumentDownloads field value
func (o *CreateMarketplaceItemLabelsResponse) GetDocumentDownloads() []DocumentDownload {
	if o == nil {
		var ret []DocumentDownload
		return ret
	}

	return o.DocumentDownloads
}

// GetDocumentDownloadsOk returns a tuple with the DocumentDownloads field value
// and a boolean to check if the value has been set.
func (o *CreateMarketplaceItemLabelsResponse) GetDocumentDownloadsOk() ([]DocumentDownload, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentDownloads, true
}

// SetDocumentDownloads sets field value
func (o *CreateMarketplaceItemLabelsResponse) SetDocumentDownloads(v []DocumentDownload) {
	o.DocumentDownloads = v
}

func (o CreateMarketplaceItemLabelsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["documentDownloads"] = o.DocumentDownloads
	return toSerialize, nil
}

type NullableCreateMarketplaceItemLabelsResponse struct {
	value *CreateMarketplaceItemLabelsResponse
	isSet bool
}

func (v NullableCreateMarketplaceItemLabelsResponse) Get() *CreateMarketplaceItemLabelsResponse {
	return v.value
}

func (v *NullableCreateMarketplaceItemLabelsResponse) Set(val *CreateMarketplaceItemLabelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMarketplaceItemLabelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMarketplaceItemLabelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMarketplaceItemLabelsResponse(val *CreateMarketplaceItemLabelsResponse) *NullableCreateMarketplaceItemLabelsResponse {
	return &NullableCreateMarketplaceItemLabelsResponse{value: val, isSet: true}
}

func (v NullableCreateMarketplaceItemLabelsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateMarketplaceItemLabelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
