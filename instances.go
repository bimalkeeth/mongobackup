package main

import "time"

type Reservations struct {
	Instances     []Instance
	OwnerId       string
	ReservationId string
}

type Instance struct {
	AmiLaunchIndex                   int
	Architecture                     string
	BlockDeviceMappings              []BlockDeviceMappings
	CapacityReservationSpecification CapacityReservationSpecification
	ClientToken                      string
	CpuOptions                       CpuOptions
	EbsOptimized                     bool
	EnaSupport                       bool
	Hypervisor                       string
	IamInstanceProfile               IamInstanceProfile
	ImageId                          string
	InstanceId                       string
	InstanceType                     string
	KeyName                          string
	LaunchTime                       string
	Monitoring                       Monitoring
	NetworkInterfaces                []NetworkInterface
	Placement                        Placement
	PrivateDnsName                   string
	PrivateIpAddress                 string
	PublicDnsName                    string
	PublicIpAddress                  string
	RootDeviceName                   string
	RootDeviceType                   string
	SecurityGroups                   []SecurityGroup
	SourceDestCheck                  string
	State                            State
	StateTransitionReason            string
	SubnetId                         string
	Tags                             Tag
	VirtualizationType               string
	VpcId                            string
}

type BlockDeviceMappings struct {
	DeviceName string
	Ebs        Ebs
}
type Ebs struct {
	AttachTime          time.Time
	DeleteOnTermination bool
	Status              string
	VolumeId            string
}
type CapacityReservationSpecification struct {
	CapacityReservationPreference string
}

type CpuOptions struct {
	CoreCount      int
	ThreadsPerCore int
}

type IamInstanceProfile struct {
	Arn string
	Id  string
}
type Monitoring struct {
	State string
}
type NetworkInterface struct {
	Association        Association
	Attachment         Attachment
	Description        string
	Groups             []Group
	MacAddress         string
	NetworkInterfaceId string
	OwnerId            string
	PrivateDnsName     string
	PrivateIpAddress   string
	PrivateIpAddresses []PrivateIpAddress
	SourceDestCheck    string
	Status             string
	SubnetId           string
	VpcId              string
}
type Association struct {
	IpOwnerId     string
	PublicDnsName string
	PublicIp      string
}

type Attachment struct {
	AttachTime          time.Time
	AttachmentId        string
	DeleteOnTermination bool
	DeviceIndex         int
	Status              string
}

type Group struct {
	GroupId   string
	GroupName string
}

type PrivateIpAddress struct {
	Association      Association
	Primary          bool
	PrivateDnsName   string
	PrivateIpAddress string
}

type Placement struct {
	AvailabilityZone string
	GroupName        string
	Tenancy          string
}

type SecurityGroup struct {
	GroupId   string
	GroupName string
}

type State struct {
	Code int
	Name string
}

type Tag struct {
	Key   string
	Value string
}
