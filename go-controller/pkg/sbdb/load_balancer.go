// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package sbdb

type (
	LoadBalancerProtocol = string
)

var (
	LoadBalancerProtocolSCTP LoadBalancerProtocol = "sctp"
	LoadBalancerProtocolTCP  LoadBalancerProtocol = "tcp"
	LoadBalancerProtocolUDP  LoadBalancerProtocol = "udp"
)

// LoadBalancer defines an object in Load_Balancer table
type LoadBalancer struct {
	UUID        string                `ovsdb:"_uuid"`
	Datapaths   []string              `ovsdb:"datapaths"`
	ExternalIDs map[string]string     `ovsdb:"external_ids"`
	Name        string                `ovsdb:"name"`
	Options     map[string]string     `ovsdb:"options"`
	Protocol    *LoadBalancerProtocol `ovsdb:"protocol"`
	Vips        map[string]string     `ovsdb:"vips"`
}
