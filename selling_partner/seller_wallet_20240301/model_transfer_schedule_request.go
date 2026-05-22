package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
)

// checks if the TransferScheduleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferScheduleRequest{}

// TransferScheduleRequest Request body to initiate a scheduled transfer from a Seller Wallet bank account to another customer-defined bank account.
type TransferScheduleRequest struct {
	// The unique identifier of the source Amazon Seller Wallet bank account from which money is debited.
	SourceAccountId string `json:"sourceAccountId"`
	// The three-letter currency code of the source payment method country, in ISO 4217 format.
	SourceCurrencyCode string `json:"sourceCurrencyCode"`
	// The unique identifier of the destination bank account where the money is deposited.
	DestinationAccountId             string                       `json:"destinationAccountId"`
	DestinationTransactionInstrument TransactionInstrumentDetails `json:"destinationTransactionInstrument"`
	TransactionType                  TransactionType              `json:"transactionType"`
	TransferScheduleInformation      TransferScheduleInformation  `json:"transferScheduleInformation"`
	PaymentPreference                PaymentPreference            `json:"paymentPreference"`
	TransferScheduleStatus           *TransferScheduleStatus      `json:"transferScheduleStatus,omitempty"`
}

type _TransferScheduleRequest TransferScheduleRequest

// NewTransferScheduleRequest instantiates a new TransferScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferScheduleRequest(sourceAccountId string, sourceCurrencyCode string, destinationAccountId string, destinationTransactionInstrument TransactionInstrumentDetails, transactionType TransactionType, transferScheduleInformation TransferScheduleInformation, paymentPreference PaymentPreference) *TransferScheduleRequest {
	this := TransferScheduleRequest{}
	this.SourceAccountId = sourceAccountId
	this.SourceCurrencyCode = sourceCurrencyCode
	this.DestinationAccountId = destinationAccountId
	this.DestinationTransactionInstrument = destinationTransactionInstrument
	this.TransactionType = transactionType
	this.TransferScheduleInformation = transferScheduleInformation
	this.PaymentPreference = paymentPreference
	return &this
}

// NewTransferScheduleRequestWithDefaults instantiates a new TransferScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferScheduleRequestWithDefaults() *TransferScheduleRequest {
	this := TransferScheduleRequest{}
	return &this
}

// GetSourceAccountId returns the SourceAccountId field value
func (o *TransferScheduleRequest) GetSourceAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAccountId
}

// GetSourceAccountIdOk returns a tuple with the SourceAccountId field value
// and a boolean to check if the value has been set.
func (o *TransferScheduleRequest) GetSourceAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAccountId, true
}

// SetSourceAccountId sets field value
func (o *TransferScheduleRequest) SetSourceAccountId(v string) {
	o.SourceAccountId = v
}

// GetSourceCurrencyCode returns the SourceCurrencyCode field value
func (o *TransferScheduleRequest) GetSourceCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceCurrencyCode
}

// GetSourceCurrencyCodeOk returns a tuple with the SourceCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *TransferScheduleRequest) GetSourceCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceCurrencyCode, true
}

// SetSourceCurrencyCode sets field value
func (o *TransferScheduleRequest) SetSourceCurrencyCode(v string) {
	o.SourceCurrencyCode = v
}

// GetDestinationAccountId returns the DestinationAccountId field value
func (o *TransferScheduleRequest) GetDestinationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationAccountId
}

// GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field value
// and a boolean to check if the value has been set.
func (o *TransferScheduleRequest) GetDestinationAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationAccountId, true
}

// SetDestinationAccountId sets field value
func (o *TransferScheduleRequest) SetDestinationAccountId(v string) {
	o.DestinationAccountId = v
}

// GetDestinationTransactionInstrument returns the DestinationTransactionInstrument field value
func (o *TransferScheduleRequest) GetDestinationTransactionInstrument() TransactionInstrumentDetails {
	if o == nil {
		var ret TransactionInstrumentDetails
		return ret
	}

	return o.DestinationTransactionInstrument
}

// GetDestinationTransactionInstrumentOk returns a tuple with the DestinationTransactionInstrument field value
// and a boolean to check if the value has been set.
func (o *TransferScheduleRequest) GetDestinationTransactionInstrumentOk() (*TransactionInstrumentDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationTransactionInstrument, true
}

// SetDestinationTransactionInstrument sets field value
func (o *TransferScheduleRequest) SetDestinationTransactionInstrument(v TransactionInstrumentDetails) {
	o.DestinationTransactionInstrument = v
}

// GetTransactionType returns the TransactionType field value
func (o *TransferScheduleRequest) GetTransactionType() TransactionType {
	if o == nil {
		var ret TransactionType
		return ret
	}

	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value
// and a boolean to check if the value has been set.
func (o *TransferScheduleRequest) GetTransactionTypeOk() (*TransactionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionType, true
}

// SetTransactionType sets field value
func (o *TransferScheduleRequest) SetTransactionType(v TransactionType) {
	o.TransactionType = v
}

// GetTransferScheduleInformation returns the TransferScheduleInformation field value
func (o *TransferScheduleRequest) GetTransferScheduleInformation() TransferScheduleInformation {
	if o == nil {
		var ret TransferScheduleInformation
		return ret
	}

	return o.TransferScheduleInformation
}

// GetTransferScheduleInformationOk returns a tuple with the TransferScheduleInformation field value
// and a boolean to check if the value has been set.
func (o *TransferScheduleRequest) GetTransferScheduleInformationOk() (*TransferScheduleInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferScheduleInformation, true
}

// SetTransferScheduleInformation sets field value
func (o *TransferScheduleRequest) SetTransferScheduleInformation(v TransferScheduleInformation) {
	o.TransferScheduleInformation = v
}

// GetPaymentPreference returns the PaymentPreference field value
func (o *TransferScheduleRequest) GetPaymentPreference() PaymentPreference {
	if o == nil {
		var ret PaymentPreference
		return ret
	}

	return o.PaymentPreference
}

// GetPaymentPreferenceOk returns a tuple with the PaymentPreference field value
// and a boolean to check if the value has been set.
func (o *TransferScheduleRequest) GetPaymentPreferenceOk() (*PaymentPreference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentPreference, true
}

// SetPaymentPreference sets field value
func (o *TransferScheduleRequest) SetPaymentPreference(v PaymentPreference) {
	o.PaymentPreference = v
}

// GetTransferScheduleStatus returns the TransferScheduleStatus field value if set, zero value otherwise.
func (o *TransferScheduleRequest) GetTransferScheduleStatus() TransferScheduleStatus {
	if o == nil || IsNil(o.TransferScheduleStatus) {
		var ret TransferScheduleStatus
		return ret
	}
	return *o.TransferScheduleStatus
}

// GetTransferScheduleStatusOk returns a tuple with the TransferScheduleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferScheduleRequest) GetTransferScheduleStatusOk() (*TransferScheduleStatus, bool) {
	if o == nil || IsNil(o.TransferScheduleStatus) {
		return nil, false
	}
	return o.TransferScheduleStatus, true
}

// HasTransferScheduleStatus returns a boolean if a field has been set.
func (o *TransferScheduleRequest) HasTransferScheduleStatus() bool {
	if o != nil && !IsNil(o.TransferScheduleStatus) {
		return true
	}

	return false
}

// SetTransferScheduleStatus gets a reference to the given TransferScheduleStatus and assigns it to the TransferScheduleStatus field.
func (o *TransferScheduleRequest) SetTransferScheduleStatus(v TransferScheduleStatus) {
	o.TransferScheduleStatus = &v
}

func (o TransferScheduleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceAccountId"] = o.SourceAccountId
	toSerialize["sourceCurrencyCode"] = o.SourceCurrencyCode
	toSerialize["destinationAccountId"] = o.DestinationAccountId
	toSerialize["destinationTransactionInstrument"] = o.DestinationTransactionInstrument
	toSerialize["transactionType"] = o.TransactionType
	toSerialize["transferScheduleInformation"] = o.TransferScheduleInformation
	toSerialize["paymentPreference"] = o.PaymentPreference
	if !IsNil(o.TransferScheduleStatus) {
		toSerialize["transferScheduleStatus"] = o.TransferScheduleStatus
	}
	return toSerialize, nil
}

type NullableTransferScheduleRequest struct {
	value *TransferScheduleRequest
	isSet bool
}

func (v NullableTransferScheduleRequest) Get() *TransferScheduleRequest {
	return v.value
}

func (v *NullableTransferScheduleRequest) Set(val *TransferScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferScheduleRequest(val *TransferScheduleRequest) *NullableTransferScheduleRequest {
	return &NullableTransferScheduleRequest{value: val, isSet: true}
}

func (v NullableTransferScheduleRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransferScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
