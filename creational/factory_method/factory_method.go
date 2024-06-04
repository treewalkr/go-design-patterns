package factorymethod

import (
	"errors"
	"fmt"
)

// DeviceSpecs represents the specifications of a mobile device
type DeviceSpecs struct {
	RAM      string
	Storage  string
	CPU      string
	Platform string
}

// MobileDevice defines the interface for mobile devices
type MobileDevice interface {
	GetPlatform() string
	GetSpecs() DeviceSpecs
	Update(newVersion string) error
	InstallApp(appName string) error
}

// BaseDevice contains common functionality for all devices
type BaseDevice struct {
	specs DeviceSpecs
}

// AndroidDevice is a concrete implementation of MobileDevice for Android
type AndroidDevice struct {
	BaseDevice
	googleServices bool
}

// NewAndroidDevice creates a new AndroidDevice with specified specs
func NewAndroidDevice(specs DeviceSpecs) *AndroidDevice {
	if specs.Platform == "" {
		specs.Platform = "Android 13.0"
	}
	return &AndroidDevice{
		BaseDevice:     BaseDevice{specs: specs},
		googleServices: true,
	}
}

func (d *AndroidDevice) GetPlatform() string {
	return d.specs.Platform
}

func (d *AndroidDevice) GetSpecs() DeviceSpecs {
	return d.specs
}

func (d *AndroidDevice) Update(newVersion string) error {
	if !d.googleServices {
		return errors.New("cannot update: Google Services not available")
	}
	d.specs.Platform = fmt.Sprintf("Android %s", newVersion)
	return nil
}

func (d *AndroidDevice) InstallApp(appName string) error {
	if !d.googleServices {
		return errors.New("cannot install app: Google Play Store not available")
	}
	fmt.Printf("Installing %s from Google Play Store\n", appName)
	return nil
}

// IosDevice is a concrete implementation of MobileDevice for iOS
type IosDevice struct {
	BaseDevice
	appleID string
}

// NewIosDevice creates a new IosDevice with specified specs
func NewIosDevice(specs DeviceSpecs) *IosDevice {
	if specs.Platform == "" {
		specs.Platform = "iOS 16.0"
	}
	return &IosDevice{
		BaseDevice: BaseDevice{specs: specs},
		appleID:    "",
	}
}

func (d *IosDevice) GetPlatform() string {
	return d.specs.Platform
}

func (d *IosDevice) GetSpecs() DeviceSpecs {
	return d.specs
}

func (d *IosDevice) Update(newVersion string) error {
	if d.appleID == "" {
		return errors.New("cannot update: Apple ID not configured")
	}
	d.specs.Platform = fmt.Sprintf("iOS %s", newVersion)
	return nil
}

func (d *IosDevice) InstallApp(appName string) error {
	if d.appleID == "" {
		return errors.New("cannot install app: Apple ID not configured")
	}
	fmt.Printf("Installing %s from App Store\n", appName)
	return nil
}

// DeviceType represents the type of mobile device
type DeviceType string

const (
	Android DeviceType = "Android"
	IOS     DeviceType = "iOS"
)

// DeviceFactory is responsible for creating mobile devices
type DeviceFactory struct {
	defaultSpecs DeviceSpecs
}

// NewDeviceFactory creates a new DeviceFactory with default specs
func NewDeviceFactory(defaultSpecs DeviceSpecs) *DeviceFactory {
	return &DeviceFactory{defaultSpecs: defaultSpecs}
}

// CreateDevice is the factory method for creating mobile devices
func (f *DeviceFactory) CreateDevice(deviceType DeviceType, specs *DeviceSpecs) (MobileDevice, error) {
	finalSpecs := f.defaultSpecs
	if specs != nil {
		finalSpecs = *specs
	}

	switch deviceType {
	case Android:
		return NewAndroidDevice(finalSpecs), nil
	case IOS:
		return NewIosDevice(finalSpecs), nil
	default:
		return nil, fmt.Errorf("unsupported device type: %s", deviceType)
	}
}

// Example demonstrates the usage of the improved Factory Method pattern
func Example() {
	// Create a factory with default specs
	defaultSpecs := DeviceSpecs{
		RAM:     "8GB",
		Storage: "128GB",
		CPU:     "Octa-core",
	}
	factory := NewDeviceFactory(defaultSpecs)

	// Create an Android device with default specs
	androidDevice, err := factory.CreateDevice(Android, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create an iOS device with custom specs
	customSpecs := DeviceSpecs{
		RAM:      "6GB",
		Storage:  "256GB",
		CPU:      "A15 Bionic",
		Platform: "iOS 17.0",
	}
	iosDevice, err := factory.CreateDevice(IOS, &customSpecs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Demonstrate usage
	fmt.Printf("Android Device Specs: %+v\n", androidDevice.GetSpecs())
	fmt.Printf("iOS Device Specs: %+v\n", iosDevice.GetSpecs())

	// Try updating and installing apps
	androidDevice.Update("14.0")
	androidDevice.InstallApp("WhatsApp")

	iosDevice.Update("17.1")
	iosDevice.InstallApp("Instagram")
}
