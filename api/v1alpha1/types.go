// Copyright 2021 Weaveworks or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MPL-2.0

package v1alpha1

import (
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// MicrovmSpec represents the specification for a microvm.
type MicrovmSpec struct {
	// VCPU specifies how many vcpu's the microvm will be allocated.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum:=1
	VCPU int64 `json:"vcpu"`

	// MemoryMb is the amount of memory in megabytes that the microvm will be allocated.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum:=1024
	MemoryMb int64 `json:"memoryMb"`

	// RootVolume specifies the volume to use for the root of the microvm.
	// +kubebuilder:validation:Required
	RootVolume Volume `json:"rootVolume"`

	// AdditionalVolumes specifies additional non-root volumes to attach to the microvm.
	// +optional
	AdditionalVolumes []Volume `json:"volumes,omitempty"`

	// Kernel specifies the kernel and its arguments to use.
	// +kubebuilder:validation:Required
	Kernel ContainerFileSource `json:"kernel"`

	// KernelCmdLine are the args to use for the kernel cmdline.
	// +kubebuilder:validation:MinLength:=5
	// +kubebuilder:default:=console=ttyS0 reboot=k panic=1 pci=off i8042.noaux i8042.nomux i8042.nopnp i8042.dumbkbd
	KernelCmdLine string `json:"kernelCmdline,omitempty"`

	// Initrd is an optional initial ramdisk to use.
	// +optional
	Initrd *ContainerFileSource `json:"initrd,omitempty"`

	// NetworkInterfaces specifies the network interfaces attached to the microvm.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems:=1
	NetworkInterfaces []NetworkInterface `json:"networkInterfaces"`
}

// MicrovmMachineTemplateResource describes the data needed to create a MicrovmMachine from a template.
type MicrovmMachineTemplateResource struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	ObjectMeta clusterv1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the specification of the machine.
	Spec MicrovmMachineSpec `json:"spec"`
}

// ContainerFileSource represents a file coming from a container.
type ContainerFileSource struct {
	// Image is the container image to use.
	// +kubebuilder:validation:Required
	Image string `json:"image"`
	// Filename is the name of the file in the container to use.
	// +optional
	Filename string `json:"filename,omitempty"`
}

// Volume represents a volume to be attached to a microvm.
type Volume struct {
	// ID is a unique identifier for this volume.
	// +kubebuilder:validation:Required
	ID string `json:"id"`

	// Image is the container image to use for the volume.
	// +kubebuilder:validation:Required
	Image string `json:"image"`
	// ReadOnly specifies that the volume is to be mounted readonly.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// MountPoint is the mount point of the volume in the machine.
	// +kubebuilder:default:=/
	MountPoint string `json:"mountPoint,omitempty"`
}

// IfaceType is a type representing the network interface types.
type IfaceType string

const (
	// IfaceTypeTap is a TAP network interface.
	IfaceTypeTap = "tap"
	// IfaceTypeMacvtap is a MACVTAP network interface.
	IfaceTypeMacvtap = "macvtap"
)

// NetworkInterface represents a network interface for the microvm.
type NetworkInterface struct {
	// GuestDeviceName is the name of the network interface to create in the microvm.
	// +kubebuilder:validation:Required
	GuestDeviceName string `json:"guestDeviceName"`
	// GuestMAC allows the specifying of a specific MAC address to use for the interface. If
	// not supplied a autogenerated MAC address will be used.
	// +optional
	GuestMAC string `json:"guestMac,omitempty"`
	// Type is the type of host network interface type to create to use by the guest.
	// +kubebuilder:validation:Enum=macvtap;tap
	Type IfaceType `json:"type"`
	// Address is an optional IP address to assign to this interface. If not supplied then DHCP will be used.
	// +optional
	Address string `json:"address,omitempty"`
}

// VMState is a type that represents the state of a microvm.
type VMState string

var (
	// VMStatePending indicates the microvm hasn't been started.
	VMStatePending = VMState("pending")
	// VMStateRunning indicates the microvm is running.
	VMStateRunning = VMState("running")
	// VMStateFailed indicates the microvm has failed.
	VMStateFailed = VMState("failed")
	// VMStateDeleted indicates the microvm has been deleted.
	VMStateDeleted = VMState("deleted")
	// VMStateUnknown indicates the microvm is in an state that is unknown/supported by CAPMVM.
	VMStateUnknown = VMState("unknown")
)
