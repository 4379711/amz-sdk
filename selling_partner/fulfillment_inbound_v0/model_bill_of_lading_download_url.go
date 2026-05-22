package fulfillment_inbound_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the BillOfLadingDownloadURL type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillOfLadingDownloadURL{}

// BillOfLadingDownloadURL Download URL for the bill of lading.
type BillOfLadingDownloadURL struct {
	// URL to download the bill of lading for the package. Note: The URL will only be valid for 15 seconds
	DownloadURL *string `json:"DownloadURL,omitempty"`
}

// NewBillOfLadingDownloadURL instantiates a new BillOfLadingDownloadURL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillOfLadingDownloadURL() *BillOfLadingDownloadURL {
	this := BillOfLadingDownloadURL{}
	return &this
}

// NewBillOfLadingDownloadURLWithDefaults instantiates a new BillOfLadingDownloadURL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillOfLadingDownloadURLWithDefaults() *BillOfLadingDownloadURL {
	this := BillOfLadingDownloadURL{}
	return &this
}

// GetDownloadURL returns the DownloadURL field value if set, zero value otherwise.
func (o *BillOfLadingDownloadURL) GetDownloadURL() string {
	if o == nil || IsNil(o.DownloadURL) {
		var ret string
		return ret
	}
	return *o.DownloadURL
}

// GetDownloadURLOk returns a tuple with the DownloadURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillOfLadingDownloadURL) GetDownloadURLOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadURL) {
		return nil, false
	}
	return o.DownloadURL, true
}

// HasDownloadURL returns a boolean if a field has been set.
func (o *BillOfLadingDownloadURL) HasDownloadURL() bool {
	if o != nil && !IsNil(o.DownloadURL) {
		return true
	}

	return false
}

// SetDownloadURL gets a reference to the given string and assigns it to the DownloadURL field.
func (o *BillOfLadingDownloadURL) SetDownloadURL(v string) {
	o.DownloadURL = &v
}

func (o BillOfLadingDownloadURL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownloadURL) {
		toSerialize["DownloadURL"] = o.DownloadURL
	}
	return toSerialize, nil
}

type NullableBillOfLadingDownloadURL struct {
	value *BillOfLadingDownloadURL
	isSet bool
}

func (v NullableBillOfLadingDownloadURL) Get() *BillOfLadingDownloadURL {
	return v.value
}

func (v *NullableBillOfLadingDownloadURL) Set(val *BillOfLadingDownloadURL) {
	v.value = val
	v.isSet = true
}

func (v NullableBillOfLadingDownloadURL) IsSet() bool {
	return v.isSet
}

func (v *NullableBillOfLadingDownloadURL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillOfLadingDownloadURL(val *BillOfLadingDownloadURL) *NullableBillOfLadingDownloadURL {
	return &NullableBillOfLadingDownloadURL{value: val, isSet: true}
}

func (v NullableBillOfLadingDownloadURL) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBillOfLadingDownloadURL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
