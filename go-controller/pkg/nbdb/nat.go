// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package nbdb

type (
	NATType = string
)

var (
	NATTypeDNAT        NATType = "dnat"
	NATTypeDNATAndSNAT NATType = "dnat_and_snat"
	NATTypeSNAT        NATType = "snat"
)

// NAT defines an object in NAT table
type NAT struct {
	UUID              string            `ovsdb:"_uuid"`
	AllowedExtIPs     *string           `ovsdb:"allowed_ext_ips"`
	ExemptedExtIPs    *string           `ovsdb:"exempted_ext_ips"`
	ExternalIDs       map[string]string `ovsdb:"external_ids"`
	ExternalIP        string            `ovsdb:"external_ip"`
	ExternalMAC       *string           `ovsdb:"external_mac"`
	ExternalPortRange string            `ovsdb:"external_port_range"`
	LogicalIP         string            `ovsdb:"logical_ip"`
	LogicalPort       *string           `ovsdb:"logical_port"`
	Options           map[string]string `ovsdb:"options"`
	Type              NATType           `ovsdb:"type"`
}
