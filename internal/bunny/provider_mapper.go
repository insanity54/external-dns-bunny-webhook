package bunny

import (
	"sigs.k8s.io/external-dns/endpoint"
)

func recordToEndpoint(domain string, record *Record) *endpoint.Endpoint {
	var dnsName string
	if record.Name == "" {
		dnsName = domain
	} else {
		dnsName = record.Name + "." + domain
	}
	ep := endpoint.NewEndpointWithTTL(
		dnsName,
		record.Type.String(),
		endpoint.TTL(record.TTLSeconds),
		record.Value,
	)

	ps := providerSpecificOptionsFromRecord(record)
	ps.ApplyToEndpoint(ep)

	return ep
}
