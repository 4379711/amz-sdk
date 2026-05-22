package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the GetDeliveryChallanDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeliveryChallanDocumentResponse{}

// GetDeliveryChallanDocumentResponse The `getDeliveryChallanDocumentResponse` response.
type GetDeliveryChallanDocumentResponse struct {
	DocumentDownload DocumentDownload `json:"documentDownload"`
}

type _GetDeliveryChallanDocumentResponse GetDeliveryChallanDocumentResponse

// NewGetDeliveryChallanDocumentResponse instantiates a new GetDeliveryChallanDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeliveryChallanDocumentResponse(documentDownload DocumentDownload) *GetDeliveryChallanDocumentResponse {
	this := GetDeliveryChallanDocumentResponse{}
	this.DocumentDownload = documentDownload
	return &this
}

// NewGetDeliveryChallanDocumentResponseWithDefaults instantiates a new GetDeliveryChallanDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeliveryChallanDocumentResponseWithDefaults() *GetDeliveryChallanDocumentResponse {
	this := GetDeliveryChallanDocumentResponse{}
	return &this
}

// GetDocumentDownload returns the DocumentDownload field value
func (o *GetDeliveryChallanDocumentResponse) GetDocumentDownload() DocumentDownload {
	if o == nil {
		var ret DocumentDownload
		return ret
	}

	return o.DocumentDownload
}

// GetDocumentDownloadOk returns a tuple with the DocumentDownload field value
// and a boolean to check if the value has been set.
func (o *GetDeliveryChallanDocumentResponse) GetDocumentDownloadOk() (*DocumentDownload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentDownload, true
}

// SetDocumentDownload sets field value
func (o *GetDeliveryChallanDocumentResponse) SetDocumentDownload(v DocumentDownload) {
	o.DocumentDownload = v
}

func (o GetDeliveryChallanDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["documentDownload"] = o.DocumentDownload
	return toSerialize, nil
}

type NullableGetDeliveryChallanDocumentResponse struct {
	value *GetDeliveryChallanDocumentResponse
	isSet bool
}

func (v NullableGetDeliveryChallanDocumentResponse) Get() *GetDeliveryChallanDocumentResponse {
	return v.value
}

func (v *NullableGetDeliveryChallanDocumentResponse) Set(val *GetDeliveryChallanDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeliveryChallanDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeliveryChallanDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeliveryChallanDocumentResponse(val *GetDeliveryChallanDocumentResponse) *NullableGetDeliveryChallanDocumentResponse {
	return &NullableGetDeliveryChallanDocumentResponse{value: val, isSet: true}
}

func (v NullableGetDeliveryChallanDocumentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetDeliveryChallanDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
