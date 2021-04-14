# Sensordata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Temperature** | **float32** |  | 
**Humidity** | **float32** |  | 

## Methods

### NewSensordata

`func NewSensordata(temperature float32, humidity float32, ) *Sensordata`

NewSensordata instantiates a new Sensordata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensordataWithDefaults

`func NewSensordataWithDefaults() *Sensordata`

NewSensordataWithDefaults instantiates a new Sensordata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemperature

`func (o *Sensordata) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *Sensordata) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *Sensordata) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.


### GetHumidity

`func (o *Sensordata) GetHumidity() float32`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *Sensordata) GetHumidityOk() (*float32, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *Sensordata) SetHumidity(v float32)`

SetHumidity sets Humidity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


