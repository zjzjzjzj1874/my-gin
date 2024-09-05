# HelperHTTPError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewHelperHTTPError

`func NewHelperHTTPError() *HelperHTTPError`

NewHelperHTTPError instantiates a new HelperHTTPError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelperHTTPErrorWithDefaults

`func NewHelperHTTPErrorWithDefaults() *HelperHTTPError`

NewHelperHTTPErrorWithDefaults instantiates a new HelperHTTPError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *HelperHTTPError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *HelperHTTPError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *HelperHTTPError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *HelperHTTPError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *HelperHTTPError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HelperHTTPError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HelperHTTPError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HelperHTTPError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


