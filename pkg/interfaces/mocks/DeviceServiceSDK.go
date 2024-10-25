// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	bootstrapinterfaces "github.com/edgexfoundry/go-mod-bootstrap/v4/bootstrap/interfaces"
	dtos "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"

	echo "github.com/labstack/echo/v4"

	http "net/http"

	interfaces "github.com/edgexfoundry/device-sdk-go/v4/pkg/interfaces"

	logger "github.com/edgexfoundry/go-mod-core-contracts/v4/clients/logger"

	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	pkgmodels "github.com/edgexfoundry/device-sdk-go/v4/pkg/models"
)

// DeviceServiceSDK is an autogenerated mock type for the DeviceServiceSDK type
type DeviceServiceSDK struct {
	mock.Mock
}

// AddCustomRoute provides a mock function with given fields: route, authentication, handler, methods
func (_m *DeviceServiceSDK) AddCustomRoute(route string, authentication interfaces.Authentication, handler func(echo.Context) error, methods ...string) error {
	_va := make([]interface{}, len(methods))
	for _i := range methods {
		_va[_i] = methods[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, route, authentication, handler)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddCustomRoute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interfaces.Authentication, func(echo.Context) error, ...string) error); ok {
		r0 = rf(route, authentication, handler, methods...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddDevice provides a mock function with given fields: device
func (_m *DeviceServiceSDK) AddDevice(device models.Device) (string, error) {
	ret := _m.Called(device)

	if len(ret) == 0 {
		panic("no return value specified for AddDevice")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Device) (string, error)); ok {
		return rf(device)
	}
	if rf, ok := ret.Get(0).(func(models.Device) string); ok {
		r0 = rf(device)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(models.Device) error); ok {
		r1 = rf(device)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddDeviceAutoEvent provides a mock function with given fields: deviceName, event
func (_m *DeviceServiceSDK) AddDeviceAutoEvent(deviceName string, event models.AutoEvent) error {
	ret := _m.Called(deviceName, event)

	if len(ret) == 0 {
		panic("no return value specified for AddDeviceAutoEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, models.AutoEvent) error); ok {
		r0 = rf(deviceName, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddDeviceProfile provides a mock function with given fields: profile
func (_m *DeviceServiceSDK) AddDeviceProfile(profile models.DeviceProfile) (string, error) {
	ret := _m.Called(profile)

	if len(ret) == 0 {
		panic("no return value specified for AddDeviceProfile")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(models.DeviceProfile) (string, error)); ok {
		return rf(profile)
	}
	if rf, ok := ret.Get(0).(func(models.DeviceProfile) string); ok {
		r0 = rf(profile)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(models.DeviceProfile) error); ok {
		r1 = rf(profile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddProvisionWatcher provides a mock function with given fields: watcher
func (_m *DeviceServiceSDK) AddProvisionWatcher(watcher models.ProvisionWatcher) (string, error) {
	ret := _m.Called(watcher)

	if len(ret) == 0 {
		panic("no return value specified for AddProvisionWatcher")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(models.ProvisionWatcher) (string, error)); ok {
		return rf(watcher)
	}
	if rf, ok := ret.Get(0).(func(models.ProvisionWatcher) string); ok {
		r0 = rf(watcher)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(models.ProvisionWatcher) error); ok {
		r1 = rf(watcher)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddRoute provides a mock function with given fields: route, handler, methods
func (_m *DeviceServiceSDK) AddRoute(route string, handler func(http.ResponseWriter, *http.Request), methods ...string) error {
	_va := make([]interface{}, len(methods))
	for _i := range methods {
		_va[_i] = methods[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, route, handler)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddRoute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, func(http.ResponseWriter, *http.Request), ...string) error); ok {
		r0 = rf(route, handler, methods...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AsyncReadingsEnabled provides a mock function with given fields:
func (_m *DeviceServiceSDK) AsyncReadingsEnabled() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AsyncReadingsEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AsyncValuesChannel provides a mock function with given fields:
func (_m *DeviceServiceSDK) AsyncValuesChannel() chan *pkgmodels.AsyncValues {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AsyncValuesChannel")
	}

	var r0 chan *pkgmodels.AsyncValues
	if rf, ok := ret.Get(0).(func() chan *pkgmodels.AsyncValues); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *pkgmodels.AsyncValues)
		}
	}

	return r0
}

// DeviceCommand provides a mock function with given fields: deviceName, commandName
func (_m *DeviceServiceSDK) DeviceCommand(deviceName string, commandName string) (models.DeviceCommand, bool) {
	ret := _m.Called(deviceName, commandName)

	if len(ret) == 0 {
		panic("no return value specified for DeviceCommand")
	}

	var r0 models.DeviceCommand
	var r1 bool
	if rf, ok := ret.Get(0).(func(string, string) (models.DeviceCommand, bool)); ok {
		return rf(deviceName, commandName)
	}
	if rf, ok := ret.Get(0).(func(string, string) models.DeviceCommand); ok {
		r0 = rf(deviceName, commandName)
	} else {
		r0 = ret.Get(0).(models.DeviceCommand)
	}

	if rf, ok := ret.Get(1).(func(string, string) bool); ok {
		r1 = rf(deviceName, commandName)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// DeviceDiscoveryEnabled provides a mock function with given fields:
func (_m *DeviceServiceSDK) DeviceDiscoveryEnabled() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DeviceDiscoveryEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DeviceExistsForName provides a mock function with given fields: name
func (_m *DeviceServiceSDK) DeviceExistsForName(name string) bool {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for DeviceExistsForName")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DeviceProfiles provides a mock function with given fields:
func (_m *DeviceServiceSDK) DeviceProfiles() []models.DeviceProfile {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DeviceProfiles")
	}

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func() []models.DeviceProfile); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	return r0
}

// DeviceResource provides a mock function with given fields: deviceName, deviceResource
func (_m *DeviceServiceSDK) DeviceResource(deviceName string, deviceResource string) (models.DeviceResource, bool) {
	ret := _m.Called(deviceName, deviceResource)

	if len(ret) == 0 {
		panic("no return value specified for DeviceResource")
	}

	var r0 models.DeviceResource
	var r1 bool
	if rf, ok := ret.Get(0).(func(string, string) (models.DeviceResource, bool)); ok {
		return rf(deviceName, deviceResource)
	}
	if rf, ok := ret.Get(0).(func(string, string) models.DeviceResource); ok {
		r0 = rf(deviceName, deviceResource)
	} else {
		r0 = ret.Get(0).(models.DeviceResource)
	}

	if rf, ok := ret.Get(1).(func(string, string) bool); ok {
		r1 = rf(deviceName, deviceResource)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Devices provides a mock function with given fields:
func (_m *DeviceServiceSDK) Devices() []models.Device {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Devices")
	}

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func() []models.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	return r0
}

// DiscoveredDeviceChannel provides a mock function with given fields:
func (_m *DeviceServiceSDK) DiscoveredDeviceChannel() chan []pkgmodels.DiscoveredDevice {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DiscoveredDeviceChannel")
	}

	var r0 chan []pkgmodels.DiscoveredDevice
	if rf, ok := ret.Get(0).(func() chan []pkgmodels.DiscoveredDevice); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan []pkgmodels.DiscoveredDevice)
		}
	}

	return r0
}

// DriverConfigs provides a mock function with given fields:
func (_m *DeviceServiceSDK) DriverConfigs() map[string]string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DriverConfigs")
	}

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// GetDeviceByName provides a mock function with given fields: name
func (_m *DeviceServiceSDK) GetDeviceByName(name string) (models.Device, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetDeviceByName")
	}

	var r0 models.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Device, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) models.Device); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfileByName provides a mock function with given fields: name
func (_m *DeviceServiceSDK) GetProfileByName(name string) (models.DeviceProfile, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetProfileByName")
	}

	var r0 models.DeviceProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.DeviceProfile, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) models.DeviceProfile); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.DeviceProfile)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionWatcherByName provides a mock function with given fields: name
func (_m *DeviceServiceSDK) GetProvisionWatcherByName(name string) (models.ProvisionWatcher, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetProvisionWatcherByName")
	}

	var r0 models.ProvisionWatcher
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.ProvisionWatcher, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) models.ProvisionWatcher); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.ProvisionWatcher)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListenForCustomConfigChanges provides a mock function with given fields: configToWatch, sectionName, changedCallback
func (_m *DeviceServiceSDK) ListenForCustomConfigChanges(configToWatch interface{}, sectionName string, changedCallback func(interface{})) error {
	ret := _m.Called(configToWatch, sectionName, changedCallback)

	if len(ret) == 0 {
		panic("no return value specified for ListenForCustomConfigChanges")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, func(interface{})) error); ok {
		r0 = rf(configToWatch, sectionName, changedCallback)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadCustomConfig provides a mock function with given fields: customConfig, sectionName
func (_m *DeviceServiceSDK) LoadCustomConfig(customConfig interfaces.UpdatableConfig, sectionName string) error {
	ret := _m.Called(customConfig, sectionName)

	if len(ret) == 0 {
		panic("no return value specified for LoadCustomConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interfaces.UpdatableConfig, string) error); ok {
		r0 = rf(customConfig, sectionName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoggingClient provides a mock function with given fields:
func (_m *DeviceServiceSDK) LoggingClient() logger.LoggingClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LoggingClient")
	}

	var r0 logger.LoggingClient
	if rf, ok := ret.Get(0).(func() logger.LoggingClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.LoggingClient)
		}
	}

	return r0
}

// MetricsManager provides a mock function with given fields:
func (_m *DeviceServiceSDK) MetricsManager() bootstrapinterfaces.MetricsManager {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MetricsManager")
	}

	var r0 bootstrapinterfaces.MetricsManager
	if rf, ok := ret.Get(0).(func() bootstrapinterfaces.MetricsManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bootstrapinterfaces.MetricsManager)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *DeviceServiceSDK) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PatchDevice provides a mock function with given fields: updateDevice
func (_m *DeviceServiceSDK) PatchDevice(updateDevice dtos.UpdateDevice) error {
	ret := _m.Called(updateDevice)

	if len(ret) == 0 {
		panic("no return value specified for PatchDevice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(dtos.UpdateDevice) error); ok {
		r0 = rf(updateDevice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvisionWatchers provides a mock function with given fields:
func (_m *DeviceServiceSDK) ProvisionWatchers() []models.ProvisionWatcher {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ProvisionWatchers")
	}

	var r0 []models.ProvisionWatcher
	if rf, ok := ret.Get(0).(func() []models.ProvisionWatcher); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProvisionWatcher)
		}
	}

	return r0
}

// PublishDeviceDiscoveryProgressSystemEvent provides a mock function with given fields: progress, discoveredDeviceCount, message
func (_m *DeviceServiceSDK) PublishDeviceDiscoveryProgressSystemEvent(progress int, discoveredDeviceCount int, message string) {
	_m.Called(progress, discoveredDeviceCount, message)
}

// PublishGenericSystemEvent provides a mock function with given fields: eventType, action, details
func (_m *DeviceServiceSDK) PublishGenericSystemEvent(eventType string, action string, details interface{}) {
	_m.Called(eventType, action, details)
}

// PublishProfileScanProgressSystemEvent provides a mock function with given fields: reqId, progress, message
func (_m *DeviceServiceSDK) PublishProfileScanProgressSystemEvent(reqId string, progress int, message string) {
	_m.Called(reqId, progress, message)
}

// RemoveDeviceAutoEvent provides a mock function with given fields: deviceName, event
func (_m *DeviceServiceSDK) RemoveDeviceAutoEvent(deviceName string, event models.AutoEvent) error {
	ret := _m.Called(deviceName, event)

	if len(ret) == 0 {
		panic("no return value specified for RemoveDeviceAutoEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, models.AutoEvent) error); ok {
		r0 = rf(deviceName, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveDeviceByName provides a mock function with given fields: name
func (_m *DeviceServiceSDK) RemoveDeviceByName(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for RemoveDeviceByName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveDeviceProfileByName provides a mock function with given fields: name
func (_m *DeviceServiceSDK) RemoveDeviceProfileByName(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for RemoveDeviceProfileByName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveProvisionWatcher provides a mock function with given fields: name
func (_m *DeviceServiceSDK) RemoveProvisionWatcher(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for RemoveProvisionWatcher")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Run provides a mock function with given fields:
func (_m *DeviceServiceSDK) Run() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SecretProvider provides a mock function with given fields:
func (_m *DeviceServiceSDK) SecretProvider() bootstrapinterfaces.SecretProvider {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SecretProvider")
	}

	var r0 bootstrapinterfaces.SecretProvider
	if rf, ok := ret.Get(0).(func() bootstrapinterfaces.SecretProvider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bootstrapinterfaces.SecretProvider)
		}
	}

	return r0
}

// UpdateDevice provides a mock function with given fields: device
func (_m *DeviceServiceSDK) UpdateDevice(device models.Device) error {
	ret := _m.Called(device)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDevice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Device) error); ok {
		r0 = rf(device)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceOperatingState provides a mock function with given fields: name, state
func (_m *DeviceServiceSDK) UpdateDeviceOperatingState(name string, state models.OperatingState) error {
	ret := _m.Called(name, state)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDeviceOperatingState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, models.OperatingState) error); ok {
		r0 = rf(name, state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceProfile provides a mock function with given fields: profile
func (_m *DeviceServiceSDK) UpdateDeviceProfile(profile models.DeviceProfile) error {
	ret := _m.Called(profile)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDeviceProfile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.DeviceProfile) error); ok {
		r0 = rf(profile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProvisionWatcher provides a mock function with given fields: watcher
func (_m *DeviceServiceSDK) UpdateProvisionWatcher(watcher models.ProvisionWatcher) error {
	ret := _m.Called(watcher)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProvisionWatcher")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.ProvisionWatcher) error); ok {
		r0 = rf(watcher)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Version provides a mock function with given fields:
func (_m *DeviceServiceSDK) Version() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Version")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewDeviceServiceSDK creates a new instance of DeviceServiceSDK. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceServiceSDK(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceServiceSDK {
	mock := &DeviceServiceSDK{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
