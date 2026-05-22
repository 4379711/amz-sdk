package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the TransactionInitiationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionInitiationRequest{}

// TransactionInitiationRequest Request body to initiate a transaction from a Seller Wallet bank account to another customer-defined bank account.
type TransactionInitiationRequest struct {
	// The unique identifier of the source Amazon Seller Wallet bank account from which the money is debited.
	SourceAccountId string `json:"sourceAccountId"`
	// The unique identifier of the destination bank account where the money is deposited.
	DestinationAccountId *string `json:"destinationAccountId,omitempty"`
	// A description of the transaction.
	Description                      string                       `json:"description"`
	DestinationTransactionInstrument TransactionInstrumentDetails `json:"destinationTransactionInstrument"`
	DestinationAccountHolderAddress  *AccountHolderAddress        `json:"destinationAccountHolderAddress,omitempty"`
	SourceAmount                     Currency                     `json:"sourceAmount"`
	TransferRateDetails              *TransferRatePreview         `json:"transferRateDetails,omitempty"`
	// The time at which the transaction was initiated in [ISO 8601 date time format](https://developer-docs.amazon.com/sp-api/docs/iso-8601).
	RequestTime time.Time `json:"requestTime"`
}

type _TransactionInitiationRequest TransactionInitiationRequest

// NewTransactionInitiationRequest instantiates a new TransactionInitiationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionInitiationRequest(sourceAccountId string, description string, destinationTransactionInstrument TransactionInstrumentDetails, sourceAmount Currency, requestTime time.Time) *TransactionInitiationRequest {
	this := TransactionInitiationRequest{}
	this.SourceAccountId = sourceAccountId
	this.Description = description
	this.DestinationTransactionInstrument = destinationTransactionInstrument
	this.SourceAmount = sourceAmount
	this.RequestTime = requestTime
	return &this
}

// NewTransactionInitiationRequestWithDefaults instantiates a new TransactionInitiationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionInitiationRequestWithDefaults() *TransactionInitiationRequest {
	this := TransactionInitiationRequest{}
	return &this
}

// GetSourceAccountId returns the SourceAccountId field value
func (o *TransactionInitiationRequest) GetSourceAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAccountId
}

// GetSourceAccountIdOk returns a tuple with the SourceAccountId field value
// and a boolean to check if the value has been set.
func (o *TransactionInitiationRequest) GetSourceAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAccountId, true
}

// SetSourceAccountId sets field value
func (o *TransactionInitiationRequest) SetSourceAccountId(v string) {
	o.SourceAccountId = v
}

// GetDestinationAccountId returns the DestinationAccountId field value if set, zero value otherwise.
func (o *TransactionInitiationRequest) GetDestinationAccountId() string {
	if o == nil || IsNil(o.DestinationAccountId) {
		var ret string
		return ret
	}
	return *o.DestinationAccountId
}

// GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInitiationRequest) GetDestinationAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationAccountId) {
		return nil, false
	}
	return o.DestinationAccountId, true
}

// HasDestinationAccountId returns a boolean if a field has been set.
func (o *TransactionInitiationRequest) HasDestinationAccountId() bool {
	if o != nil && !IsNil(o.DestinationAccountId) {
		return true
	}

	return false
}

// SetDestinationAccountId gets a reference to the given string and assigns it to the DestinationAccountId field.
func (o *TransactionInitiationRequest) SetDestinationAccountId(v string) {
	o.DestinationAccountId = &v
}

// GetDescription returns the Description field value
func (o *TransactionInitiationRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransactionInitiationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransactionInitiationRequest) SetDescription(v string) {
	o.Description = v
}

// GetDestinationTransactionInstrument returns the DestinationTransactionInstrument field value
func (o *TransactionInitiationRequest) GetDestinationTransactionInstrument() TransactionInstrumentDetails {
	if o == nil {
		var ret TransactionInstrumentDetails
		return ret
	}

	return o.DestinationTransactionInstrument
}

// GetDestinationTransactionInstrumentOk returns a tuple with the DestinationTransactionInstrument field value
// and a boolean to check if the value has been set.
func (o *TransactionInitiationRequest) GetDestinationTransactionInstrumentOk() (*TransactionInstrumentDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationTransactionInstrument, true
}

// SetDestinationTransactionInstrument sets field value
func (o *TransactionInitiationRequest) SetDestinationTransactionInstrument(v TransactionInstrumentDetails) {
	o.DestinationTransactionInstrument = v
}

// GetDestinationAccountHolderAddress returns the DestinationAccountHolderAddress field value if set, zero value otherwise.
func (o *TransactionInitiationRequest) GetDestinationAccountHolderAddress() AccountHolderAddress {
	if o == nil || IsNil(o.DestinationAccountHolderAddress) {
		var ret AccountHolderAddress
		return ret
	}
	return *o.DestinationAccountHolderAddress
}

// GetDestinationAccountHolderAddressOk returns a tuple with the DestinationAccountHolderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInitiationRequest) GetDestinationAccountHolderAddressOk() (*AccountHolderAddress, bool) {
	if o == nil || IsNil(o.DestinationAccountHolderAddress) {
		return nil, false
	}
	return o.DestinationAccountHolderAddress, true
}

// HasDestinationAccountHolderAddress returns a boolean if a field has been set.
func (o *TransactionInitiationRequest) HasDestinationAccountHolderAddress() bool {
	if o != nil && !IsNil(o.DestinationAccountHolderAddress) {
		return true
	}

	return false
}

// SetDestinationAccountHolderAddress gets a reference to the given AccountHolderAddress and assigns it to the DestinationAccountHolderAddress field.
func (o *TransactionInitiationRequest) SetDestinationAccountHolderAddress(v AccountHolderAddress) {
	o.DestinationAccountHolderAddress = &v
}

// GetSourceAmount returns the SourceAmount field value
func (o *TransactionInitiationRequest) GetSourceAmount() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.SourceAmount
}

// GetSourceAmountOk returns a tuple with the SourceAmount field value
// and a boolean to check if the value has been set.
func (o *TransactionInitiationRequest) GetSourceAmountOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAmount, true
}

// SetSourceAmount sets field value
func (o *TransactionInitiationRequest) SetSourceAmount(v Currency) {
	o.SourceAmount = v
}

// GetTransferRateDetails returns the TransferRateDetails field value if set, zero value otherwise.
func (o *TransactionInitiationRequest) GetTransferRateDetails() TransferRatePreview {
	if o == nil || IsNil(o.TransferRateDetails) {
		var ret TransferRatePreview
		return ret
	}
	return *o.TransferRateDetails
}

// GetTransferRateDetailsOk returns a tuple with the TransferRateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInitiationRequest) GetTransferRateDetailsOk() (*TransferRatePreview, bool) {
	if o == nil || IsNil(o.TransferRateDetails) {
		return nil, false
	}
	return o.TransferRateDetails, true
}

// HasTransferRateDetails returns a boolean if a field has been set.
func (o *TransactionInitiationRequest) HasTransferRateDetails() bool {
	if o != nil && !IsNil(o.TransferRateDetails) {
		return true
	}

	return false
}

// SetTransferRateDetails gets a reference to the given TransferRatePreview and assigns it to the TransferRateDetails field.
func (o *TransactionInitiationRequest) SetTransferRateDetails(v TransferRatePreview) {
	o.TransferRateDetails = &v
}

// GetRequestTime returns the RequestTime field value
func (o *TransactionInitiationRequest) GetRequestTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RequestTime
}

// GetRequestTimeOk returns a tuple with the RequestTime field value
// and a boolean to check if the value has been set.
func (o *TransactionInitiationRequest) GetRequestTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestTime, true
}

// SetRequestTime sets field value
func (o *TransactionInitiationRequest) SetRequestTime(v time.Time) {
	o.RequestTime = v
}

func (o TransactionInitiationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceAccountId"] = o.SourceAccountId
	if !IsNil(o.DestinationAccountId) {
		toSerialize["destinationAccountId"] = o.DestinationAccountId
	}
	toSerialize["description"] = o.Description
	toSerialize["destinationTransactionInstrument"] = o.DestinationTransactionInstrument
	if !IsNil(o.DestinationAccountHolderAddress) {
		toSerialize["destinationAccountHolderAddress"] = o.DestinationAccountHolderAddress
	}
	toSerialize["sourceAmount"] = o.SourceAmount
	if !IsNil(o.TransferRateDetails) {
		toSerialize["transferRateDetails"] = o.TransferRateDetails
	}
	toSerialize["requestTime"] = o.RequestTime
	return toSerialize, nil
}

type NullableTransactionInitiationRequest struct {
	value *TransactionInitiationRequest
	isSet bool
}

func (v NullableTransactionInitiationRequest) Get() *TransactionInitiationRequest {
	return v.value
}

func (v *NullableTransactionInitiationRequest) Set(val *TransactionInitiationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInitiationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInitiationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInitiationRequest(val *TransactionInitiationRequest) *NullableTransactionInitiationRequest {
	return &NullableTransactionInitiationRequest{value: val, isSet: true}
}

func (v NullableTransactionInitiationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransactionInitiationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
