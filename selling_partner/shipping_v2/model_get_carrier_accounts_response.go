package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the GetCarrierAccountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCarrierAccountsResponse{}

// GetCarrierAccountsResponse The Response  for the GetCarrierAccountsResponse operation.
type GetCarrierAccountsResponse struct {
	// A list of ActiveAccount
	ActiveAccounts []ActiveAccount `json:"activeAccounts"`
}

type _GetCarrierAccountsResponse GetCarrierAccountsResponse

// NewGetCarrierAccountsResponse instantiates a new GetCarrierAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCarrierAccountsResponse(activeAccounts []ActiveAccount) *GetCarrierAccountsResponse {
	this := GetCarrierAccountsResponse{}
	this.ActiveAccounts = activeAccounts
	return &this
}

// NewGetCarrierAccountsResponseWithDefaults instantiates a new GetCarrierAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCarrierAccountsResponseWithDefaults() *GetCarrierAccountsResponse {
	this := GetCarrierAccountsResponse{}
	return &this
}

// GetActiveAccounts returns the ActiveAccounts field value
func (o *GetCarrierAccountsResponse) GetActiveAccounts() []ActiveAccount {
	if o == nil {
		var ret []ActiveAccount
		return ret
	}

	return o.ActiveAccounts
}

// GetActiveAccountsOk returns a tuple with the ActiveAccounts field value
// and a boolean to check if the value has been set.
func (o *GetCarrierAccountsResponse) GetActiveAccountsOk() ([]ActiveAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveAccounts, true
}

// SetActiveAccounts sets field value
func (o *GetCarrierAccountsResponse) SetActiveAccounts(v []ActiveAccount) {
	o.ActiveAccounts = v
}

func (o GetCarrierAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activeAccounts"] = o.ActiveAccounts
	return toSerialize, nil
}

type NullableGetCarrierAccountsResponse struct {
	value *GetCarrierAccountsResponse
	isSet bool
}

func (v NullableGetCarrierAccountsResponse) Get() *GetCarrierAccountsResponse {
	return v.value
}

func (v *NullableGetCarrierAccountsResponse) Set(val *GetCarrierAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCarrierAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCarrierAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCarrierAccountsResponse(val *GetCarrierAccountsResponse) *NullableGetCarrierAccountsResponse {
	return &NullableGetCarrierAccountsResponse{value: val, isSet: true}
}

func (v NullableGetCarrierAccountsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetCarrierAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
