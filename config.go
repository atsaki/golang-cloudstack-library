package cloudstack

import (
	"reflect"
	"strings"
	"time"
)

var config struct {
	PollingInterval time.Duration
	Commands        map[string]Command
}

func init() {
	config.PollingInterval = 5
}

type Command struct {
	Name       string
	IsAsync    bool
	IsList     bool
	ObjectType string
}

func (cmd Command) Pointer() interface{} {
	refType := cmd.ReflectType()
	refTypeKind := reflect.ValueOf(refType).Kind()
	if refTypeKind != reflect.Ptr && refTypeKind != reflect.Slice {
		return nil
	}
	return reflect.New(refType).Interface()
}

func (cmd Command) ReflectType() reflect.Type {
	switch cmd.ObjectType {
	case "affinitygroup":
		if cmd.IsList {
			return reflect.TypeOf([]*AffinityGroup{})
		}
		return reflect.TypeOf(&AffinityGroup{})
	case "affinitygrouptype":
		if cmd.IsList {
			return reflect.TypeOf([]*AffinityGroupType{})
		}
		return reflect.TypeOf(&AffinityGroupType{})
	case "capacity":
		if cmd.IsList {
			return reflect.TypeOf([]*Capacity{})
		}
		return reflect.TypeOf(&Capacity{})
	case "diskoffering":
		if cmd.IsList {
			return reflect.TypeOf([]*DiskOffering{})
		}
		return reflect.TypeOf(&DiskOffering{})
	case "firewallrule":
		if cmd.IsList {
			return reflect.TypeOf([]*FirewallRule{})
		}
		return reflect.TypeOf(&FirewallRule{})
	case "ipaddress":
		if cmd.IsList {
			return reflect.TypeOf([]*PublicIpAddress{})
		}
		return reflect.TypeOf(&PublicIpAddress{})
	case "ipforwardingrule":
		if cmd.IsList {
			return reflect.TypeOf([]*IpForwardingRule{})
		}
		return reflect.TypeOf(&IpForwardingRule{})
	case "loadbalancer":
		if cmd.IsList {
			return reflect.TypeOf([]*LoadBalancerRule{})
		}
		return reflect.TypeOf(&LoadBalancerRule{})
	case "loadbalancerrule":
		if cmd.IsList {
			return reflect.TypeOf([]*LoadBalancerRule{})
		}
		return reflect.TypeOf(&LoadBalancerRule{})
	case "networkoffering":
		if cmd.IsList {
			return reflect.TypeOf([]*NetworkOffering{})
		}
		return reflect.TypeOf(&NetworkOffering{})
	case "networkservicecapability":
		if cmd.IsList {
			return reflect.TypeOf([]*NetworkServiceCapability{})
		}
		return reflect.TypeOf(&NetworkServiceCapability{})
	case "networkserviceprovider":
		if cmd.IsList {
			return reflect.TypeOf([]*NetworkServiceProvider{})
		}
		return reflect.TypeOf(&NetworkServiceProvider{})
	case "networkservice":
		if cmd.IsList {
			return reflect.TypeOf([]*NetworkService{})
		}
		return reflect.TypeOf(&NetworkService{})
	case "network":
		if cmd.IsList {
			return reflect.TypeOf([]*Network{})
		}
		return reflect.TypeOf(&Network{})
	case "nic":
		if cmd.IsList {
			return reflect.TypeOf([]*Nic{})
		}
		return reflect.TypeOf(&Nic{})
	case "portforwardingrule":
		if cmd.IsList {
			return reflect.TypeOf([]*PortForwardingRule{})
		}
		return reflect.TypeOf(&PortForwardingRule{})
	case "publicipaddress":
		if cmd.IsList {
			return reflect.TypeOf([]*PublicIpAddress{})
		}
		return reflect.TypeOf(&PublicIpAddress{})
	case "result":
		if cmd.IsList {
			return reflect.TypeOf([]*Result{})
		}
		return reflect.TypeOf(&Result{})
	case "securitygroupegress":
		if cmd.IsList {
			return reflect.TypeOf([]*SecurityGroupEgress{})
		}
		return reflect.TypeOf(&SecurityGroupEgress{})
	case "securitygroupingress":
		if cmd.IsList {
			return reflect.TypeOf([]*SecurityGroupIngress{})
		}
		return reflect.TypeOf(&SecurityGroupIngress{})
	case "securitygroup":
		if cmd.IsList {
			return reflect.TypeOf([]*SecurityGroup{})
		}
		return reflect.TypeOf(&SecurityGroup{})
	case "serviceoffering":
		if cmd.IsList {
			return reflect.TypeOf([]*ServiceOffering{})
		}
		return reflect.TypeOf(&ServiceOffering{})
	case "tag":
		if cmd.IsList {
			return reflect.TypeOf([]*Tag{})
		}
		return reflect.TypeOf(&Tag{})
	case "templatepermission":
		if cmd.IsList {
			return reflect.TypeOf([]*TemplatePermission{})
		}
		return reflect.TypeOf(&TemplatePermission{})
	case "template":
		if cmd.IsList {
			return reflect.TypeOf([]*Template{})
		}
		return reflect.TypeOf(&Template{})
	case "vmpassword":
		if cmd.IsList {
			return reflect.TypeOf([]*VMPassword{})
		}
		return reflect.TypeOf(&VMPassword{})
	case "virtualmachine":
		if cmd.IsList {
			return reflect.TypeOf([]*VirtualMachine{})
		}
		return reflect.TypeOf(&VirtualMachine{})
	case "volume":
		if cmd.IsList {
			return reflect.TypeOf([]*Volume{})
		}
		return reflect.TypeOf(&Volume{})
	case "zone":
		if cmd.IsList {
			return reflect.TypeOf([]*Zone{})
		}
		return reflect.TypeOf(&Zone{})
	default:
		return nil
	}
}

func getCommand(name string) *Command {
	switch strings.ToLower(name) {
	case "activateproject":
		return &Command{
			Name:       "activateProject",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "project",
		}
	case "addaccounttoproject":
		return &Command{
			Name:       "addAccountToProject",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "addbaremetaldhcp":
		return &Command{
			Name:       "addBaremetalDhcp",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "baremetaldhcp",
		}
	case "addbaremetalhost":
		return &Command{
			Name:       "addBaremetalHost",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "host",
		}
	case "addbaremetalpxekickstartserver":
		return &Command{
			Name:       "addBaremetalPxeKickStartServer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "baremetalpxekickstartserver",
		}
	case "addbaremetalpxepingserver":
		return &Command{
			Name:       "addBaremetalPxePingServer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "baremetalpxepingserver",
		}
	case "addbigswitchvnsdevice":
		return &Command{
			Name:       "addBigSwitchVnsDevice",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "bigswitchvnsdevice",
		}
	case "addciscoasa1000vresource":
		return &Command{
			Name:       "addCiscoAsa1000vResource",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "ciscoasa1000vresource",
		}
	case "addciscovnmcresource":
		return &Command{
			Name:       "addCiscoVnmcResource",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "ciscovnmcresource",
		}
	case "addcluster":
		return &Command{
			Name:       "addCluster",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "cluster",
		}
	case "addexternalfirewall":
		return &Command{
			Name:       "addExternalFirewall",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "externalfirewall",
		}
	case "addexternalloadbalancer":
		return &Command{
			Name:       "addExternalLoadBalancer",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "externalloadbalancer",
		}
	case "addf5loadbalancer":
		return &Command{
			Name:       "addF5LoadBalancer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "f5loadbalancer",
		}
	case "addguestos":
		return &Command{
			Name:       "addGuestOs",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "guestos",
		}
	case "addguestosmapping":
		return &Command{
			Name:       "addGuestOsMapping",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "guestosmapping",
		}
	case "addhost":
		return &Command{
			Name:       "addHost",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "host",
		}
	case "addimagestore":
		return &Command{
			Name:       "addImageStore",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "imagestore",
		}
	case "addiptonic":
		return &Command{
			Name:       "addIpToNic",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "iptonic",
		}
	case "addldapconfiguration":
		return &Command{
			Name:       "addLdapConfiguration",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "ldapconfiguration",
		}
	case "addnetscalerloadbalancer":
		return &Command{
			Name:       "addNetscalerLoadBalancer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "netscalerloadbalancer",
		}
	case "addnetworkdevice":
		return &Command{
			Name:       "addNetworkDevice",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "networkdevice",
		}
	case "addnetworkserviceprovider":
		return &Command{
			Name:       "addNetworkServiceProvider",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "networkserviceprovider",
		}
	case "addnictovirtualmachine":
		return &Command{
			Name:       "addNicToVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "addniciranvpdevice":
		return &Command{
			Name:       "addNiciraNvpDevice",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "niciranvpdevice",
		}
	case "addopendaylightcontroller":
		return &Command{
			Name:       "addOpenDaylightController",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "opendaylightcontroller",
		}
	case "addpaloaltofirewall":
		return &Command{
			Name:       "addPaloAltoFirewall",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "paloaltofirewall",
		}
	case "addregion":
		return &Command{
			Name:       "addRegion",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "region",
		}
	case "addresourcedetail":
		return &Command{
			Name:       "addResourceDetail",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "adds3":
		return &Command{
			Name:       "addS3",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "s3",
		}
	case "addsecondarystorage":
		return &Command{
			Name:       "addSecondaryStorage",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "secondarystorage",
		}
	case "addsrxfirewall":
		return &Command{
			Name:       "addSrxFirewall",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "srxfirewall",
		}
	case "addstratospheressp":
		return &Command{
			Name:       "addStratosphereSsp",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "stratospheressp",
		}
	case "addswift":
		return &Command{
			Name:       "addSwift",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "swift",
		}
	case "addtrafficmonitor":
		return &Command{
			Name:       "addTrafficMonitor",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "trafficmonitor",
		}
	case "addtraffictype":
		return &Command{
			Name:       "addTrafficType",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "traffictype",
		}
	case "adducsmanager":
		return &Command{
			Name:       "addUcsManager",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "ucsmanager",
		}
	case "addvmwaredc":
		return &Command{
			Name:       "addVmwareDc",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "vmwaredc",
		}
	case "addvpnuser":
		return &Command{
			Name:       "addVpnUser",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpnuser",
		}
	case "archivealerts":
		return &Command{
			Name:       "archiveAlerts",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "archiveevents":
		return &Command{
			Name:       "archiveEvents",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "assigncerttoloadbalancer":
		return &Command{
			Name:       "assignCertToLoadBalancer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "assigntogloballoadbalancerrule":
		return &Command{
			Name:       "assignToGlobalLoadBalancerRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "assigntoloadbalancerrule":
		return &Command{
			Name:       "assignToLoadBalancerRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "assignvirtualmachine":
		return &Command{
			Name:       "assignVirtualMachine",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "associateipaddress":
		return &Command{
			Name:       "associateIpAddress",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "ipaddress",
		}
	case "associateucsprofiletoblade":
		return &Command{
			Name:       "associateUcsProfileToBlade",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "ucsprofiletoblade",
		}
	case "attachiso":
		return &Command{
			Name:       "attachIso",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "attachvolume":
		return &Command{
			Name:       "attachVolume",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "volume",
		}
	case "authorizesecuritygroupegress":
		return &Command{
			Name:       "authorizeSecurityGroupEgress",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "securitygroupegress",
		}
	case "authorizesecuritygroupingress":
		return &Command{
			Name:       "authorizeSecurityGroupIngress",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "securitygroupingress",
		}
	case "cancelhostmaintenance":
		return &Command{
			Name:       "cancelHostMaintenance",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "host",
		}
	case "cancelstoragemaintenance":
		return &Command{
			Name:       "cancelStorageMaintenance",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "storagemaintenance",
		}
	case "changeserviceforrouter":
		return &Command{
			Name:       "changeServiceForRouter",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "router",
		}
	case "changeserviceforsystemvm":
		return &Command{
			Name:       "changeServiceForSystemVm",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "systemvm",
		}
	case "changeserviceforvirtualmachine":
		return &Command{
			Name:       "changeServiceForVirtualMachine",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "cleanvmreservations":
		return &Command{
			Name:       "cleanVMReservations",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "configuref5loadbalancer":
		return &Command{
			Name:       "configureF5LoadBalancer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "f5loadbalancer",
		}
	case "configureinternalloadbalancerelement":
		return &Command{
			Name:       "configureInternalLoadBalancerElement",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "internalloadbalancerelement",
		}
	case "configurenetscalerloadbalancer":
		return &Command{
			Name:       "configureNetscalerLoadBalancer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "netscalerloadbalancer",
		}
	case "configureovselement":
		return &Command{
			Name:       "configureOvsElement",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "ovselement",
		}
	case "configurepaloaltofirewall":
		return &Command{
			Name:       "configurePaloAltoFirewall",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "paloaltofirewall",
		}
	case "configuresrxfirewall":
		return &Command{
			Name:       "configureSrxFirewall",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "srxfirewall",
		}
	case "configurevirtualrouterelement":
		return &Command{
			Name:       "configureVirtualRouterElement",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualrouterelement",
		}
	case "copyiso":
		return &Command{
			Name:       "copyIso",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "iso",
		}
	case "copytemplate":
		return &Command{
			Name:       "copyTemplate",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "template",
		}
	case "createaccount":
		return &Command{
			Name:       "createAccount",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "account",
		}
	case "createaffinitygroup":
		return &Command{
			Name:       "createAffinityGroup",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "affinitygroup",
		}
	case "createautoscalepolicy":
		return &Command{
			Name:       "createAutoScalePolicy",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "autoscalepolicy",
		}
	case "createautoscalevmgroup":
		return &Command{
			Name:       "createAutoScaleVmGroup",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "autoscalevmgroup",
		}
	case "createautoscalevmprofile":
		return &Command{
			Name:       "createAutoScaleVmProfile",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "autoscalevmprofile",
		}
	case "createcondition":
		return &Command{
			Name:       "createCondition",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "condition",
		}
	case "createcounter":
		return &Command{
			Name:       "createCounter",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "counter",
		}
	case "creatediskoffering":
		return &Command{
			Name:       "createDiskOffering",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "diskoffering",
		}
	case "createdomain":
		return &Command{
			Name:       "createDomain",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "domain",
		}
	case "createegressfirewallrule":
		return &Command{
			Name:       "createEgressFirewallRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "firewallrule",
		}
	case "createfirewallrule":
		return &Command{
			Name:       "createFirewallRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "firewallrule",
		}
	case "creategloballoadbalancerrule":
		return &Command{
			Name:       "createGlobalLoadBalancerRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "globalloadbalancerrule",
		}
	case "createinstancegroup":
		return &Command{
			Name:       "createInstanceGroup",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "instancegroup",
		}
	case "createinternalloadbalancerelement":
		return &Command{
			Name:       "createInternalLoadBalancerElement",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "internalloadbalancerelement",
		}
	case "createipforwardingrule":
		return &Command{
			Name:       "createIpForwardingRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "ipforwardingrule",
		}
	case "createlbhealthcheckpolicy":
		return &Command{
			Name:       "createLBHealthCheckPolicy",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "lbhealthcheckpolicy",
		}
	case "createlbstickinesspolicy":
		return &Command{
			Name:       "createLBStickinessPolicy",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "lbstickinesspolicy",
		}
	case "createloadbalancer":
		return &Command{
			Name:       "createLoadBalancer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "loadbalancer",
		}
	case "createloadbalancerrule":
		return &Command{
			Name:       "createLoadBalancerRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "loadbalancer",
		}
	case "createnetwork":
		return &Command{
			Name:       "createNetwork",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "network",
		}
	case "createnetworkacl":
		return &Command{
			Name:       "createNetworkACL",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "networkacl",
		}
	case "createnetworkacllist":
		return &Command{
			Name:       "createNetworkACLList",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "networkacllist",
		}
	case "createnetworkoffering":
		return &Command{
			Name:       "createNetworkOffering",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "networkoffering",
		}
	case "createphysicalnetwork":
		return &Command{
			Name:       "createPhysicalNetwork",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "physicalnetwork",
		}
	case "createpod":
		return &Command{
			Name:       "createPod",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "pod",
		}
	case "createportforwardingrule":
		return &Command{
			Name:       "createPortForwardingRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "portforwardingrule",
		}
	case "createportableiprange":
		return &Command{
			Name:       "createPortableIpRange",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "portableiprange",
		}
	case "createprivategateway":
		return &Command{
			Name:       "createPrivateGateway",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "privategateway",
		}
	case "createproject":
		return &Command{
			Name:       "createProject",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "project",
		}
	case "createremoteaccessvpn":
		return &Command{
			Name:       "createRemoteAccessVpn",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "remoteaccessvpn",
		}
	case "createsshkeypair":
		return &Command{
			Name:       "createSSHKeyPair",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "keypair",
		}
	case "createsecondarystagingstore":
		return &Command{
			Name:       "createSecondaryStagingStore",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "secondarystagingstore",
		}
	case "createsecuritygroup":
		return &Command{
			Name:       "createSecurityGroup",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "securitygroup",
		}
	case "createserviceinstance":
		return &Command{
			Name:       "createServiceInstance",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "serviceinstance",
		}
	case "createserviceoffering":
		return &Command{
			Name:       "createServiceOffering",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "serviceoffering",
		}
	case "createsnapshot":
		return &Command{
			Name:       "createSnapshot",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "snapshot",
		}
	case "createsnapshotpolicy":
		return &Command{
			Name:       "createSnapshotPolicy",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "snapshotpolicy",
		}
	case "createstaticroute":
		return &Command{
			Name:       "createStaticRoute",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "staticroute",
		}
	case "createstoragenetworkiprange":
		return &Command{
			Name:       "createStorageNetworkIpRange",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "storagenetworkiprange",
		}
	case "createstoragepool":
		return &Command{
			Name:       "createStoragePool",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "storagepool",
		}
	case "createtags":
		return &Command{
			Name:       "createTags",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "createtemplate":
		return &Command{
			Name:       "createTemplate",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "template",
		}
	case "createuser":
		return &Command{
			Name:       "createUser",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "user",
		}
	case "createvmsnapshot":
		return &Command{
			Name:       "createVMSnapshot",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vmsnapshot",
		}
	case "createvpc":
		return &Command{
			Name:       "createVPC",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpc",
		}
	case "createvpcoffering":
		return &Command{
			Name:       "createVPCOffering",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpcoffering",
		}
	case "createvirtualrouterelement":
		return &Command{
			Name:       "createVirtualRouterElement",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualrouterelement",
		}
	case "createvlaniprange":
		return &Command{
			Name:       "createVlanIpRange",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "vlaniprange",
		}
	case "createvolume":
		return &Command{
			Name:       "createVolume",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "volume",
		}
	case "createvpnconnection":
		return &Command{
			Name:       "createVpnConnection",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpnconnection",
		}
	case "createvpncustomergateway":
		return &Command{
			Name:       "createVpnCustomerGateway",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpncustomergateway",
		}
	case "createvpngateway":
		return &Command{
			Name:       "createVpnGateway",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpngateway",
		}
	case "createzone":
		return &Command{
			Name:       "createZone",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "zone",
		}
	case "dedicatecluster":
		return &Command{
			Name:       "dedicateCluster",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "cluster",
		}
	case "dedicateguestvlanrange":
		return &Command{
			Name:       "dedicateGuestVlanRange",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "guestvlanrange",
		}
	case "dedicatehost":
		return &Command{
			Name:       "dedicateHost",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "host",
		}
	case "dedicatepod":
		return &Command{
			Name:       "dedicatePod",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "pod",
		}
	case "dedicatepubliciprange":
		return &Command{
			Name:       "dedicatePublicIpRange",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "publiciprange",
		}
	case "dedicatezone":
		return &Command{
			Name:       "dedicateZone",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "zone",
		}
	case "deleteaccount":
		return &Command{
			Name:       "deleteAccount",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteaccountfromproject":
		return &Command{
			Name:       "deleteAccountFromProject",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteaffinitygroup":
		return &Command{
			Name:       "deleteAffinityGroup",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletealerts":
		return &Command{
			Name:       "deleteAlerts",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteautoscalepolicy":
		return &Command{
			Name:       "deleteAutoScalePolicy",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteautoscalevmgroup":
		return &Command{
			Name:       "deleteAutoScaleVmGroup",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteautoscalevmprofile":
		return &Command{
			Name:       "deleteAutoScaleVmProfile",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletebigswitchvnsdevice":
		return &Command{
			Name:       "deleteBigSwitchVnsDevice",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteciscoasa1000vresource":
		return &Command{
			Name:       "deleteCiscoAsa1000vResource",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletecisconexusvsm":
		return &Command{
			Name:       "deleteCiscoNexusVSM",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteciscovnmcresource":
		return &Command{
			Name:       "deleteCiscoVnmcResource",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletecluster":
		return &Command{
			Name:       "deleteCluster",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletecondition":
		return &Command{
			Name:       "deleteCondition",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletecounter":
		return &Command{
			Name:       "deleteCounter",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletediskoffering":
		return &Command{
			Name:       "deleteDiskOffering",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletedomain":
		return &Command{
			Name:       "deleteDomain",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteegressfirewallrule":
		return &Command{
			Name:       "deleteEgressFirewallRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteevents":
		return &Command{
			Name:       "deleteEvents",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteexternalfirewall":
		return &Command{
			Name:       "deleteExternalFirewall",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteexternalloadbalancer":
		return &Command{
			Name:       "deleteExternalLoadBalancer",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletef5loadbalancer":
		return &Command{
			Name:       "deleteF5LoadBalancer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletefirewallrule":
		return &Command{
			Name:       "deleteFirewallRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletegloballoadbalancerrule":
		return &Command{
			Name:       "deleteGlobalLoadBalancerRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletehost":
		return &Command{
			Name:       "deleteHost",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteimagestore":
		return &Command{
			Name:       "deleteImageStore",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteinstancegroup":
		return &Command{
			Name:       "deleteInstanceGroup",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteipforwardingrule":
		return &Command{
			Name:       "deleteIpForwardingRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteiso":
		return &Command{
			Name:       "deleteIso",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletelbhealthcheckpolicy":
		return &Command{
			Name:       "deleteLBHealthCheckPolicy",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletelbstickinesspolicy":
		return &Command{
			Name:       "deleteLBStickinessPolicy",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteldapconfiguration":
		return &Command{
			Name:       "deleteLdapConfiguration",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "ldapconfiguration",
		}
	case "deleteloadbalancer":
		return &Command{
			Name:       "deleteLoadBalancer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteloadbalancerrule":
		return &Command{
			Name:       "deleteLoadBalancerRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletenetscalerloadbalancer":
		return &Command{
			Name:       "deleteNetscalerLoadBalancer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletenetwork":
		return &Command{
			Name:       "deleteNetwork",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletenetworkacl":
		return &Command{
			Name:       "deleteNetworkACL",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletenetworkacllist":
		return &Command{
			Name:       "deleteNetworkACLList",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletenetworkdevice":
		return &Command{
			Name:       "deleteNetworkDevice",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletenetworkoffering":
		return &Command{
			Name:       "deleteNetworkOffering",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletenetworkserviceprovider":
		return &Command{
			Name:       "deleteNetworkServiceProvider",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteniciranvpdevice":
		return &Command{
			Name:       "deleteNiciraNvpDevice",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteopendaylightcontroller":
		return &Command{
			Name:       "deleteOpenDaylightController",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "opendaylightcontroller",
		}
	case "deletepaloaltofirewall":
		return &Command{
			Name:       "deletePaloAltoFirewall",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletephysicalnetwork":
		return &Command{
			Name:       "deletePhysicalNetwork",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletepod":
		return &Command{
			Name:       "deletePod",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteportforwardingrule":
		return &Command{
			Name:       "deletePortForwardingRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteportableiprange":
		return &Command{
			Name:       "deletePortableIpRange",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteprivategateway":
		return &Command{
			Name:       "deletePrivateGateway",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteproject":
		return &Command{
			Name:       "deleteProject",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteprojectinvitation":
		return &Command{
			Name:       "deleteProjectInvitation",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteremoteaccessvpn":
		return &Command{
			Name:       "deleteRemoteAccessVpn",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletesshkeypair":
		return &Command{
			Name:       "deleteSSHKeyPair",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletesecondarystagingstore":
		return &Command{
			Name:       "deleteSecondaryStagingStore",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletesecuritygroup":
		return &Command{
			Name:       "deleteSecurityGroup",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteserviceoffering":
		return &Command{
			Name:       "deleteServiceOffering",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletesnapshot":
		return &Command{
			Name:       "deleteSnapshot",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletesnapshotpolicies":
		return &Command{
			Name:       "deleteSnapshotPolicies",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletesrxfirewall":
		return &Command{
			Name:       "deleteSrxFirewall",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletesslcert":
		return &Command{
			Name:       "deleteSslCert",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletestaticroute":
		return &Command{
			Name:       "deleteStaticRoute",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletestoragenetworkiprange":
		return &Command{
			Name:       "deleteStorageNetworkIpRange",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletestoragepool":
		return &Command{
			Name:       "deleteStoragePool",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletetags":
		return &Command{
			Name:       "deleteTags",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletetemplate":
		return &Command{
			Name:       "deleteTemplate",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletetrafficmonitor":
		return &Command{
			Name:       "deleteTrafficMonitor",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletetraffictype":
		return &Command{
			Name:       "deleteTrafficType",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deleteuser":
		return &Command{
			Name:       "deleteUser",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletevmsnapshot":
		return &Command{
			Name:       "deleteVMSnapshot",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletevpc":
		return &Command{
			Name:       "deleteVPC",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletevpcoffering":
		return &Command{
			Name:       "deleteVPCOffering",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletevlaniprange":
		return &Command{
			Name:       "deleteVlanIpRange",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletevolume":
		return &Command{
			Name:       "deleteVolume",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletevpnconnection":
		return &Command{
			Name:       "deleteVpnConnection",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletevpncustomergateway":
		return &Command{
			Name:       "deleteVpnCustomerGateway",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletevpngateway":
		return &Command{
			Name:       "deleteVpnGateway",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "deletezone":
		return &Command{
			Name:       "deleteZone",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "deployvirtualmachine":
		return &Command{
			Name:       "deployVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "destroyrouter":
		return &Command{
			Name:       "destroyRouter",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "router",
		}
	case "destroysystemvm":
		return &Command{
			Name:       "destroySystemVm",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "systemvm",
		}
	case "destroyvirtualmachine":
		return &Command{
			Name:       "destroyVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "detachiso":
		return &Command{
			Name:       "detachIso",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "detachvolume":
		return &Command{
			Name:       "detachVolume",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "volume",
		}
	case "disableaccount":
		return &Command{
			Name:       "disableAccount",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "account",
		}
	case "disableautoscalevmgroup":
		return &Command{
			Name:       "disableAutoScaleVmGroup",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "autoscalevmgroup",
		}
	case "disablecisconexusvsm":
		return &Command{
			Name:       "disableCiscoNexusVSM",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "cisconexusvsm",
		}
	case "disablestaticnat":
		return &Command{
			Name:       "disableStaticNat",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "disableuser":
		return &Command{
			Name:       "disableUser",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "user",
		}
	case "disassociateipaddress":
		return &Command{
			Name:       "disassociateIpAddress",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "enableaccount":
		return &Command{
			Name:       "enableAccount",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "account",
		}
	case "enableautoscalevmgroup":
		return &Command{
			Name:       "enableAutoScaleVmGroup",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "autoscalevmgroup",
		}
	case "enablecisconexusvsm":
		return &Command{
			Name:       "enableCiscoNexusVSM",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "cisconexusvsm",
		}
	case "enablestaticnat":
		return &Command{
			Name:       "enableStaticNat",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "enablestoragemaintenance":
		return &Command{
			Name:       "enableStorageMaintenance",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "storagemaintenance",
		}
	case "enableuser":
		return &Command{
			Name:       "enableUser",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "user",
		}
	case "expungevirtualmachine":
		return &Command{
			Name:       "expungeVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "extractiso":
		return &Command{
			Name:       "extractIso",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "iso",
		}
	case "extracttemplate":
		return &Command{
			Name:       "extractTemplate",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "template",
		}
	case "extractvolume":
		return &Command{
			Name:       "extractVolume",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "volume",
		}
	case "findhostsformigration":
		return &Command{
			Name:       "findHostsForMigration",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "host",
		}
	case "findstoragepoolsformigration":
		return &Command{
			Name:       "findStoragePoolsForMigration",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "storagepool",
		}
	case "generatealert":
		return &Command{
			Name:       "generateAlert",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "generateusagerecords":
		return &Command{
			Name:       "generateUsageRecords",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "getapilimit":
		return &Command{
			Name:       "getApiLimit",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "apilimit",
		}
	case "getcloudidentifier":
		return &Command{
			Name:       "getCloudIdentifier",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "cloudidentifier",
		}
	case "getuser":
		return &Command{
			Name:       "getUser",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "user",
		}
	case "getvmpassword":
		return &Command{
			Name:       "getVMPassword",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "vmpassword",
		}
	case "getvirtualmachineuserdata":
		return &Command{
			Name:       "getVirtualMachineUserData",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "virtualmachineuserdata",
		}
	case "importldapusers":
		return &Command{
			Name:       "importLdapUsers",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "ldapuser",
		}
	case "ldapcreateaccount":
		return &Command{
			Name:       "ldapCreateAccount",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "account",
		}
	case "listaccounts":
		return &Command{
			Name:       "listAccounts",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "account",
		}
	case "listaffinitygrouptypes":
		return &Command{
			Name:       "listAffinityGroupTypes",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "affinitygrouptype",
		}
	case "listaffinitygroups":
		return &Command{
			Name:       "listAffinityGroups",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "affinitygroup",
		}
	case "listalerts":
		return &Command{
			Name:       "listAlerts",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "alert",
		}
	case "listapis":
		return &Command{
			Name:       "listApis",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "api",
		}
	case "listasyncjobs":
		return &Command{
			Name:       "listAsyncJobs",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "asyncjob",
		}
	case "listautoscalepolicies":
		return &Command{
			Name:       "listAutoScalePolicies",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "autoscalepolicy",
		}
	case "listautoscalevmgroups":
		return &Command{
			Name:       "listAutoScaleVmGroups",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "autoscalevmgroup",
		}
	case "listautoscalevmprofiles":
		return &Command{
			Name:       "listAutoScaleVmProfiles",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "autoscalevmprofile",
		}
	case "listbaremetaldhcp":
		return &Command{
			Name:       "listBaremetalDhcp",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "baremetaldhcp",
		}
	case "listbaremetalpxeservers":
		return &Command{
			Name:       "listBaremetalPxeServers",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "baremetalpxeserver",
		}
	case "listbigswitchvnsdevices":
		return &Command{
			Name:       "listBigSwitchVnsDevices",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "bigswitchvnsdevice",
		}
	case "listcapabilities":
		return &Command{
			Name:       "listCapabilities",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "capability",
		}
	case "listcapacity":
		return &Command{
			Name:       "listCapacity",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "capacity",
		}
	case "listciscoasa1000vresources":
		return &Command{
			Name:       "listCiscoAsa1000vResources",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "ciscoasa1000vresource",
		}
	case "listcisconexusvsms":
		return &Command{
			Name:       "listCiscoNexusVSMs",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "cisconexusvsm",
		}
	case "listciscovnmcresources":
		return &Command{
			Name:       "listCiscoVnmcResources",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "ciscovnmcresource",
		}
	case "listclusters":
		return &Command{
			Name:       "listClusters",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "cluster",
		}
	case "listconditions":
		return &Command{
			Name:       "listConditions",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "condition",
		}
	case "listconfigurations":
		return &Command{
			Name:       "listConfigurations",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "configuration",
		}
	case "listcounters":
		return &Command{
			Name:       "listCounters",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "counter",
		}
	case "listdedicatedclusters":
		return &Command{
			Name:       "listDedicatedClusters",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "dedicatedcluster",
		}
	case "listdedicatedguestvlanranges":
		return &Command{
			Name:       "listDedicatedGuestVlanRanges",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "dedicatedguestvlanrange",
		}
	case "listdedicatedhosts":
		return &Command{
			Name:       "listDedicatedHosts",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "dedicatedhost",
		}
	case "listdedicatedpods":
		return &Command{
			Name:       "listDedicatedPods",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "dedicatedpod",
		}
	case "listdedicatedzones":
		return &Command{
			Name:       "listDedicatedZones",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "dedicatedzone",
		}
	case "listdeploymentplanners":
		return &Command{
			Name:       "listDeploymentPlanners",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "deploymentplanner",
		}
	case "listdiskofferings":
		return &Command{
			Name:       "listDiskOfferings",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "diskoffering",
		}
	case "listdomainchildren":
		return &Command{
			Name:       "listDomainChildren",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "domainchildren",
		}
	case "listdomains":
		return &Command{
			Name:       "listDomains",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "domain",
		}
	case "listegressfirewallrules":
		return &Command{
			Name:       "listEgressFirewallRules",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "firewallrule",
		}
	case "listeventtypes":
		return &Command{
			Name:       "listEventTypes",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "eventtype",
		}
	case "listevents":
		return &Command{
			Name:       "listEvents",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "event",
		}
	case "listexternalfirewalls":
		return &Command{
			Name:       "listExternalFirewalls",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "externalfirewall",
		}
	case "listexternalloadbalancers":
		return &Command{
			Name:       "listExternalLoadBalancers",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "externalloadbalancer",
		}
	case "listf5loadbalancernetworks":
		return &Command{
			Name:       "listF5LoadBalancerNetworks",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "network",
		}
	case "listf5loadbalancers":
		return &Command{
			Name:       "listF5LoadBalancers",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "f5loadbalancer",
		}
	case "listfirewallrules":
		return &Command{
			Name:       "listFirewallRules",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "firewallrule",
		}
	case "listgloballoadbalancerrules":
		return &Command{
			Name:       "listGlobalLoadBalancerRules",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "globalloadbalancerrule",
		}
	case "listguestosmapping":
		return &Command{
			Name:       "listGuestOsMapping",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "guestosmapping",
		}
	case "listhosts":
		return &Command{
			Name:       "listHosts",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "host",
		}
	case "listhypervisorcapabilities":
		return &Command{
			Name:       "listHypervisorCapabilities",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "hypervisorcapability",
		}
	case "listhypervisors":
		return &Command{
			Name:       "listHypervisors",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "hypervisor",
		}
	case "listimagestores":
		return &Command{
			Name:       "listImageStores",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "imagestore",
		}
	case "listinstancegroups":
		return &Command{
			Name:       "listInstanceGroups",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "instancegroup",
		}
	case "listinternalloadbalancerelements":
		return &Command{
			Name:       "listInternalLoadBalancerElements",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "internalloadbalancerelement",
		}
	case "listinternalloadbalancervms":
		return &Command{
			Name:       "listInternalLoadBalancerVMs",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "internalloadbalancervm",
		}
	case "listipforwardingrules":
		return &Command{
			Name:       "listIpForwardingRules",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "ipforwardingrule",
		}
	case "listisopermissions":
		return &Command{
			Name:       "listIsoPermissions",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "isopermission",
		}
	case "listisos":
		return &Command{
			Name:       "listIsos",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "iso",
		}
	case "listlbhealthcheckpolicies":
		return &Command{
			Name:       "listLBHealthCheckPolicies",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "lbhealthcheckpolicy",
		}
	case "listlbstickinesspolicies":
		return &Command{
			Name:       "listLBStickinessPolicies",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "lbstickinesspolicy",
		}
	case "listldapconfigurations":
		return &Command{
			Name:       "listLdapConfigurations",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "ldapconfiguration",
		}
	case "listldapusers":
		return &Command{
			Name:       "listLdapUsers",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "ldapuser",
		}
	case "listloadbalancerruleinstances":
		return &Command{
			Name:       "listLoadBalancerRuleInstances",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "loadbalancerruleinstance",
		}
	case "listloadbalancerrules":
		return &Command{
			Name:       "listLoadBalancerRules",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "loadbalancerrule",
		}
	case "listloadbalancers":
		return &Command{
			Name:       "listLoadBalancers",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "loadbalancer",
		}
	case "listnetscalerloadbalancernetworks":
		return &Command{
			Name:       "listNetscalerLoadBalancerNetworks",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "network",
		}
	case "listnetscalerloadbalancers":
		return &Command{
			Name:       "listNetscalerLoadBalancers",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "netscalerloadbalancer",
		}
	case "listnetworkacllists":
		return &Command{
			Name:       "listNetworkACLLists",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "networkacllist",
		}
	case "listnetworkacls":
		return &Command{
			Name:       "listNetworkACLs",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "networkacl",
		}
	case "listnetworkdevice":
		return &Command{
			Name:       "listNetworkDevice",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "networkdevice",
		}
	case "listnetworkisolationmethods":
		return &Command{
			Name:       "listNetworkIsolationMethods",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "networkisolationmethod",
		}
	case "listnetworkofferings":
		return &Command{
			Name:       "listNetworkOfferings",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "networkoffering",
		}
	case "listnetworkserviceproviders":
		return &Command{
			Name:       "listNetworkServiceProviders",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "networkserviceprovider",
		}
	case "listnetworks":
		return &Command{
			Name:       "listNetworks",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "network",
		}
	case "listniciranvpdevicenetworks":
		return &Command{
			Name:       "listNiciraNvpDeviceNetworks",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "network",
		}
	case "listniciranvpdevices":
		return &Command{
			Name:       "listNiciraNvpDevices",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "niciranvpdevice",
		}
	case "listnics":
		return &Command{
			Name:       "listNics",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "nic",
		}
	case "listopendaylightcontrollers":
		return &Command{
			Name:       "listOpenDaylightControllers",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "opendaylightcontroller",
		}
	case "listoscategories":
		return &Command{
			Name:       "listOsCategories",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "oscategory",
		}
	case "listostypes":
		return &Command{
			Name:       "listOsTypes",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "ostype",
		}
	case "listovselements":
		return &Command{
			Name:       "listOvsElements",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "ovselement",
		}
	case "listpaloaltofirewallnetworks":
		return &Command{
			Name:       "listPaloAltoFirewallNetworks",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "network",
		}
	case "listpaloaltofirewalls":
		return &Command{
			Name:       "listPaloAltoFirewalls",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "paloaltofirewall",
		}
	case "listphysicalnetworks":
		return &Command{
			Name:       "listPhysicalNetworks",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "physicalnetwork",
		}
	case "listpods":
		return &Command{
			Name:       "listPods",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "pod",
		}
	case "listportforwardingrules":
		return &Command{
			Name:       "listPortForwardingRules",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "portforwardingrule",
		}
	case "listportableipranges":
		return &Command{
			Name:       "listPortableIpRanges",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "portableiprange",
		}
	case "listprivategateways":
		return &Command{
			Name:       "listPrivateGateways",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "privategateway",
		}
	case "listprojectaccounts":
		return &Command{
			Name:       "listProjectAccounts",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "projectaccount",
		}
	case "listprojectinvitations":
		return &Command{
			Name:       "listProjectInvitations",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "projectinvitation",
		}
	case "listprojects":
		return &Command{
			Name:       "listProjects",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "project",
		}
	case "listpublicipaddresses":
		return &Command{
			Name:       "listPublicIpAddresses",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "publicipaddress",
		}
	case "listregions":
		return &Command{
			Name:       "listRegions",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "region",
		}
	case "listremoteaccessvpns":
		return &Command{
			Name:       "listRemoteAccessVpns",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "remoteaccessvpn",
		}
	case "listresourcedetails":
		return &Command{
			Name:       "listResourceDetails",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "resourcedetail",
		}
	case "listresourcelimits":
		return &Command{
			Name:       "listResourceLimits",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "resourcelimit",
		}
	case "listrouters":
		return &Command{
			Name:       "listRouters",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "router",
		}
	case "lists3s":
		return &Command{
			Name:       "listS3s",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "s3",
		}
	case "listsshkeypairs":
		return &Command{
			Name:       "listSSHKeyPairs",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "sshkeypair",
		}
	case "listsecondarystagingstores":
		return &Command{
			Name:       "listSecondaryStagingStores",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "secondarystagingstore",
		}
	case "listsecuritygroups":
		return &Command{
			Name:       "listSecurityGroups",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "securitygroup",
		}
	case "listserviceofferings":
		return &Command{
			Name:       "listServiceOfferings",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "serviceoffering",
		}
	case "listsnapshotpolicies":
		return &Command{
			Name:       "listSnapshotPolicies",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "snapshotpolicy",
		}
	case "listsnapshots":
		return &Command{
			Name:       "listSnapshots",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "snapshot",
		}
	case "listsrxfirewallnetworks":
		return &Command{
			Name:       "listSrxFirewallNetworks",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "network",
		}
	case "listsrxfirewalls":
		return &Command{
			Name:       "listSrxFirewalls",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "srxfirewall",
		}
	case "listsslcerts":
		return &Command{
			Name:       "listSslCerts",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "sslcert",
		}
	case "liststaticroutes":
		return &Command{
			Name:       "listStaticRoutes",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "staticroute",
		}
	case "liststoragenetworkiprange":
		return &Command{
			Name:       "listStorageNetworkIpRange",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "storagenetworkiprange",
		}
	case "liststoragepools":
		return &Command{
			Name:       "listStoragePools",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "storagepool",
		}
	case "liststorageproviders":
		return &Command{
			Name:       "listStorageProviders",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "storageprovider",
		}
	case "listsupportednetworkservices":
		return &Command{
			Name:       "listSupportedNetworkServices",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "networkservice",
		}
	case "listswifts":
		return &Command{
			Name:       "listSwifts",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "swift",
		}
	case "listsystemvms":
		return &Command{
			Name:       "listSystemVms",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "systemvm",
		}
	case "listtags":
		return &Command{
			Name:       "listTags",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "tag",
		}
	case "listtemplatepermissions":
		return &Command{
			Name:       "listTemplatePermissions",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "templatepermission",
		}
	case "listtemplates":
		return &Command{
			Name:       "listTemplates",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "template",
		}
	case "listtrafficmonitors":
		return &Command{
			Name:       "listTrafficMonitors",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "trafficmonitor",
		}
	case "listtraffictypeimplementors":
		return &Command{
			Name:       "listTrafficTypeImplementors",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "traffictypeimplementor",
		}
	case "listtraffictypes":
		return &Command{
			Name:       "listTrafficTypes",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "traffictype",
		}
	case "listucsblades":
		return &Command{
			Name:       "listUcsBlades",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "ucsblade",
		}
	case "listucsmanagers":
		return &Command{
			Name:       "listUcsManagers",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "ucsmanager",
		}
	case "listucsprofiles":
		return &Command{
			Name:       "listUcsProfiles",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "ucsprofile",
		}
	case "listusagerecords":
		return &Command{
			Name:       "listUsageRecords",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "usagerecord",
		}
	case "listusagetypes":
		return &Command{
			Name:       "listUsageTypes",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "usagetype",
		}
	case "listusers":
		return &Command{
			Name:       "listUsers",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "user",
		}
	case "listvmsnapshot":
		return &Command{
			Name:       "listVMSnapshot",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "vmsnapshot",
		}
	case "listvpcofferings":
		return &Command{
			Name:       "listVPCOfferings",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "vpcoffering",
		}
	case "listvpcs":
		return &Command{
			Name:       "listVPCs",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "vpc",
		}
	case "listvirtualmachines":
		return &Command{
			Name:       "listVirtualMachines",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "virtualmachine",
		}
	case "listvirtualrouterelements":
		return &Command{
			Name:       "listVirtualRouterElements",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "virtualrouterelement",
		}
	case "listvlanipranges":
		return &Command{
			Name:       "listVlanIpRanges",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "vlaniprange",
		}
	case "listvmwaredcs":
		return &Command{
			Name:       "listVmwareDcs",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "vmwaredc",
		}
	case "listvolumes":
		return &Command{
			Name:       "listVolumes",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "volume",
		}
	case "listvpnconnections":
		return &Command{
			Name:       "listVpnConnections",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "vpnconnection",
		}
	case "listvpncustomergateways":
		return &Command{
			Name:       "listVpnCustomerGateways",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "vpncustomergateway",
		}
	case "listvpngateways":
		return &Command{
			Name:       "listVpnGateways",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "vpngateway",
		}
	case "listvpnusers":
		return &Command{
			Name:       "listVpnUsers",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "vpnuser",
		}
	case "listzones":
		return &Command{
			Name:       "listZones",
			IsAsync:    false,
			IsList:     true,
			ObjectType: "zone",
		}
	case "lockaccount":
		return &Command{
			Name:       "lockAccount",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "account",
		}
	case "lockuser":
		return &Command{
			Name:       "lockUser",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "user",
		}
	case "markdefaultzoneforaccount":
		return &Command{
			Name:       "markDefaultZoneForAccount",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "defaultzoneforaccount",
		}
	case "migratesystemvm":
		return &Command{
			Name:       "migrateSystemVm",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "systemvm",
		}
	case "migratevirtualmachine":
		return &Command{
			Name:       "migrateVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "migratevirtualmachinewithvolume":
		return &Command{
			Name:       "migrateVirtualMachineWithVolume",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "migratevolume":
		return &Command{
			Name:       "migrateVolume",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "volume",
		}
	case "preparehostformaintenance":
		return &Command{
			Name:       "prepareHostForMaintenance",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "host",
		}
	case "preparetemplate":
		return &Command{
			Name:       "prepareTemplate",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "template",
		}
	case "queryasyncjobresult":
		return &Command{
			Name:       "queryAsyncJobResult",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "",
		}
	case "rebootrouter":
		return &Command{
			Name:       "rebootRouter",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "router",
		}
	case "rebootsystemvm":
		return &Command{
			Name:       "rebootSystemVm",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "systemvm",
		}
	case "rebootvirtualmachine":
		return &Command{
			Name:       "rebootVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "reconnecthost":
		return &Command{
			Name:       "reconnectHost",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "host",
		}
	case "recovervirtualmachine":
		return &Command{
			Name:       "recoverVirtualMachine",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "registeriso":
		return &Command{
			Name:       "registerIso",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "iso",
		}
	case "registersshkeypair":
		return &Command{
			Name:       "registerSSHKeyPair",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "sshkeypair",
		}
	case "registertemplate":
		return &Command{
			Name:       "registerTemplate",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "template",
		}
	case "registeruserkeys":
		return &Command{
			Name:       "registerUserKeys",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "userkeys",
		}
	case "releasededicatedcluster":
		return &Command{
			Name:       "releaseDedicatedCluster",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "releasededicatedguestvlanrange":
		return &Command{
			Name:       "releaseDedicatedGuestVlanRange",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "releasededicatedhost":
		return &Command{
			Name:       "releaseDedicatedHost",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "releasededicatedpod":
		return &Command{
			Name:       "releaseDedicatedPod",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "releasededicatedzone":
		return &Command{
			Name:       "releaseDedicatedZone",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "releasehostreservation":
		return &Command{
			Name:       "releaseHostReservation",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "releasepubliciprange":
		return &Command{
			Name:       "releasePublicIpRange",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "removecertfromloadbalancer":
		return &Command{
			Name:       "removeCertFromLoadBalancer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "removefromgloballoadbalancerrule":
		return &Command{
			Name:       "removeFromGlobalLoadBalancerRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "removefromloadbalancerrule":
		return &Command{
			Name:       "removeFromLoadBalancerRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "removeguestos":
		return &Command{
			Name:       "removeGuestOs",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "removeguestosmapping":
		return &Command{
			Name:       "removeGuestOsMapping",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "removeipfromnic":
		return &Command{
			Name:       "removeIpFromNic",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "removenicfromvirtualmachine":
		return &Command{
			Name:       "removeNicFromVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "removeregion":
		return &Command{
			Name:       "removeRegion",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "removeresourcedetail":
		return &Command{
			Name:       "removeResourceDetail",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "removevmwaredc":
		return &Command{
			Name:       "removeVmwareDc",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "removevpnuser":
		return &Command{
			Name:       "removeVpnUser",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "replacenetworkacllist":
		return &Command{
			Name:       "replaceNetworkACLList",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "resetapilimit":
		return &Command{
			Name:       "resetApiLimit",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "apilimit",
		}
	case "resetpasswordforvirtualmachine":
		return &Command{
			Name:       "resetPasswordForVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "resetsshkeyforvirtualmachine":
		return &Command{
			Name:       "resetSSHKeyForVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "resetvpnconnection":
		return &Command{
			Name:       "resetVpnConnection",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpnconnection",
		}
	case "resizevolume":
		return &Command{
			Name:       "resizeVolume",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "volume",
		}
	case "restartnetwork":
		return &Command{
			Name:       "restartNetwork",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "network",
		}
	case "restartvpc":
		return &Command{
			Name:       "restartVPC",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpc",
		}
	case "restorevirtualmachine":
		return &Command{
			Name:       "restoreVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "revertsnapshot":
		return &Command{
			Name:       "revertSnapshot",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "snapshot",
		}
	case "reverttovmsnapshot":
		return &Command{
			Name:       "revertToVMSnapshot",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "revokesecuritygroupegress":
		return &Command{
			Name:       "revokeSecurityGroupEgress",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "revokesecuritygroupingress":
		return &Command{
			Name:       "revokeSecurityGroupIngress",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "scalesystemvm":
		return &Command{
			Name:       "scaleSystemVm",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "systemvm",
		}
	case "scalevirtualmachine":
		return &Command{
			Name:       "scaleVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "startinternalloadbalancervm":
		return &Command{
			Name:       "startInternalLoadBalancerVM",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "internalloadbalancervm",
		}
	case "startrouter":
		return &Command{
			Name:       "startRouter",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "router",
		}
	case "startsystemvm":
		return &Command{
			Name:       "startSystemVm",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "systemvm",
		}
	case "startvirtualmachine":
		return &Command{
			Name:       "startVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "stopinternalloadbalancervm":
		return &Command{
			Name:       "stopInternalLoadBalancerVM",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "internalloadbalancervm",
		}
	case "stoprouter":
		return &Command{
			Name:       "stopRouter",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "router",
		}
	case "stopsystemvm":
		return &Command{
			Name:       "stopSystemVm",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "systemvm",
		}
	case "stopvirtualmachine":
		return &Command{
			Name:       "stopVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "suspendproject":
		return &Command{
			Name:       "suspendProject",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "project",
		}
	case "updateaccount":
		return &Command{
			Name:       "updateAccount",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "account",
		}
	case "updateautoscalepolicy":
		return &Command{
			Name:       "updateAutoScalePolicy",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "autoscalepolicy",
		}
	case "updateautoscalevmgroup":
		return &Command{
			Name:       "updateAutoScaleVmGroup",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "autoscalevmgroup",
		}
	case "updateautoscalevmprofile":
		return &Command{
			Name:       "updateAutoScaleVmProfile",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "autoscalevmprofile",
		}
	case "updatecloudtouseobjectstore":
		return &Command{
			Name:       "updateCloudToUseObjectStore",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "imagestore",
		}
	case "updatecluster":
		return &Command{
			Name:       "updateCluster",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "cluster",
		}
	case "updateconfiguration":
		return &Command{
			Name:       "updateConfiguration",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "configuration",
		}
	case "updatedefaultnicforvirtualmachine":
		return &Command{
			Name:       "updateDefaultNicForVirtualMachine",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "updatediskoffering":
		return &Command{
			Name:       "updateDiskOffering",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "diskoffering",
		}
	case "updatedomain":
		return &Command{
			Name:       "updateDomain",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "domain",
		}
	case "updateegressfirewallrule":
		return &Command{
			Name:       "updateEgressFirewallRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "firewallrule",
		}
	case "updatefirewallrule":
		return &Command{
			Name:       "updateFirewallRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "firewallrule",
		}
	case "updategloballoadbalancerrule":
		return &Command{
			Name:       "updateGlobalLoadBalancerRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "globalloadbalancerrule",
		}
	case "updateguestos":
		return &Command{
			Name:       "updateGuestOs",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "guestos",
		}
	case "updateguestosmapping":
		return &Command{
			Name:       "updateGuestOsMapping",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "guestosmapping",
		}
	case "updatehost":
		return &Command{
			Name:       "updateHost",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "host",
		}
	case "updatehostpassword":
		return &Command{
			Name:       "updateHostPassword",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "updatehypervisorcapabilities":
		return &Command{
			Name:       "updateHypervisorCapabilities",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "hypervisorcapabilities",
		}
	case "updateinstancegroup":
		return &Command{
			Name:       "updateInstanceGroup",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "instancegroup",
		}
	case "updateipaddress":
		return &Command{
			Name:       "updateIpAddress",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "ipaddress",
		}
	case "updateiso":
		return &Command{
			Name:       "updateIso",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "iso",
		}
	case "updateisopermissions":
		return &Command{
			Name:       "updateIsoPermissions",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "updatelbhealthcheckpolicy":
		return &Command{
			Name:       "updateLBHealthCheckPolicy",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "lbhealthcheckpolicy",
		}
	case "updatelbstickinesspolicy":
		return &Command{
			Name:       "updateLBStickinessPolicy",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "lbstickinesspolicy",
		}
	case "updateloadbalancer":
		return &Command{
			Name:       "updateLoadBalancer",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "loadbalancer",
		}
	case "updateloadbalancerrule":
		return &Command{
			Name:       "updateLoadBalancerRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "loadbalancer",
		}
	case "updatenetwork":
		return &Command{
			Name:       "updateNetwork",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "network",
		}
	case "updatenetworkaclitem":
		return &Command{
			Name:       "updateNetworkACLItem",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "networkaclitem",
		}
	case "updatenetworkacllist":
		return &Command{
			Name:       "updateNetworkACLList",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "updatenetworkoffering":
		return &Command{
			Name:       "updateNetworkOffering",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "networkoffering",
		}
	case "updatenetworkserviceprovider":
		return &Command{
			Name:       "updateNetworkServiceProvider",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "networkserviceprovider",
		}
	case "updatephysicalnetwork":
		return &Command{
			Name:       "updatePhysicalNetwork",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "physicalnetwork",
		}
	case "updatepod":
		return &Command{
			Name:       "updatePod",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "pod",
		}
	case "updateportforwardingrule":
		return &Command{
			Name:       "updatePortForwardingRule",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "portforwardingrule",
		}
	case "updateproject":
		return &Command{
			Name:       "updateProject",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "project",
		}
	case "updateprojectinvitation":
		return &Command{
			Name:       "updateProjectInvitation",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "result",
		}
	case "updateregion":
		return &Command{
			Name:       "updateRegion",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "region",
		}
	case "updateremoteaccessvpn":
		return &Command{
			Name:       "updateRemoteAccessVpn",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "remoteaccessvpn",
		}
	case "updateresourcecount":
		return &Command{
			Name:       "updateResourceCount",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "resourcecount",
		}
	case "updateresourcelimit":
		return &Command{
			Name:       "updateResourceLimit",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "resourcelimit",
		}
	case "updateserviceoffering":
		return &Command{
			Name:       "updateServiceOffering",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "serviceoffering",
		}
	case "updatestoragenetworkiprange":
		return &Command{
			Name:       "updateStorageNetworkIpRange",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "storagenetworkiprange",
		}
	case "updatestoragepool":
		return &Command{
			Name:       "updateStoragePool",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "storagepool",
		}
	case "updatetemplate":
		return &Command{
			Name:       "updateTemplate",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "template",
		}
	case "updatetemplatepermissions":
		return &Command{
			Name:       "updateTemplatePermissions",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "result",
		}
	case "updatetraffictype":
		return &Command{
			Name:       "updateTrafficType",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "traffictype",
		}
	case "updateuser":
		return &Command{
			Name:       "updateUser",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "user",
		}
	case "updatevmaffinitygroup":
		return &Command{
			Name:       "updateVMAffinityGroup",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "updatevpc":
		return &Command{
			Name:       "updateVPC",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpc",
		}
	case "updatevpcoffering":
		return &Command{
			Name:       "updateVPCOffering",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpcoffering",
		}
	case "updatevirtualmachine":
		return &Command{
			Name:       "updateVirtualMachine",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "virtualmachine",
		}
	case "updatevolume":
		return &Command{
			Name:       "updateVolume",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "volume",
		}
	case "updatevpnconnection":
		return &Command{
			Name:       "updateVpnConnection",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpnconnection",
		}
	case "updatevpncustomergateway":
		return &Command{
			Name:       "updateVpnCustomerGateway",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpncustomergateway",
		}
	case "updatevpngateway":
		return &Command{
			Name:       "updateVpnGateway",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "vpngateway",
		}
	case "updatezone":
		return &Command{
			Name:       "updateZone",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "zone",
		}
	case "upgraderoutertemplate":
		return &Command{
			Name:       "upgradeRouterTemplate",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "routertemplate",
		}
	case "uploadcustomcertificate":
		return &Command{
			Name:       "uploadCustomCertificate",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "customcertificate",
		}
	case "uploadsslcert":
		return &Command{
			Name:       "uploadSslCert",
			IsAsync:    false,
			IsList:     false,
			ObjectType: "sslcert",
		}
	case "uploadvolume":
		return &Command{
			Name:       "uploadVolume",
			IsAsync:    true,
			IsList:     false,
			ObjectType: "volume",
		}

	default:
		return nil
	}
}
