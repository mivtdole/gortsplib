package gortsplib

import (
	"fmt"
	"strconv"
	"strings"
)

// HeaderTransport is a Transport header.
type HeaderTransport struct {
	// protocol of the stream
	Protocol StreamProtocol

	// (optional) cast of the stream
	Cast *StreamCast

	// (optional) destination
	Destination *string

	// (optional) TTL
	TTL *uint

	// (optional) ports
	Ports *[2]int

	// (optional) client ports
	ClientPorts *[2]int

	// (optional) server ports
	ServerPorts *[2]int

	// (optional) interleaved frame ids
	InterleavedIds *[2]int

	// (optional) mode
	Mode *string
}

func parsePorts(val string) (*[2]int, error) {
	ports := strings.Split(val, "-")
	if len(ports) != 2 {
		return &[2]int{0, 0}, fmt.Errorf("invalid ports (%v)", val)
	}

	port1, err := strconv.ParseInt(ports[0], 10, 64)
	if err != nil {
		return &[2]int{0, 0}, fmt.Errorf("invalid ports (%v)", val)
	}

	port2, err := strconv.ParseInt(ports[1], 10, 64)
	if err != nil {
		return &[2]int{0, 0}, fmt.Errorf("invalid ports (%v)", val)
	}

	return &[2]int{int(port1), int(port2)}, nil
}

// ReadHeaderTransport parses a Transport header.
func ReadHeaderTransport(v HeaderValue) (*HeaderTransport, error) {
	if len(v) == 0 {
		return nil, fmt.Errorf("value not provided")
	}

	if len(v) > 1 {
		return nil, fmt.Errorf("value provided multiple times (%v)", v)
	}

	ht := &HeaderTransport{}

	parts := strings.Split(v[0], ";")
	if len(parts) == 0 {
		return nil, fmt.Errorf("invalid value (%v)", v)
	}

	switch parts[0] {
	case "RTP/AVP", "RTP/AVP/UDP":
		ht.Protocol = StreamProtocolUDP

	case "RTP/AVP/TCP":
		ht.Protocol = StreamProtocolTCP

	default:
		return nil, fmt.Errorf("invalid protocol (%v)", v)
	}
	parts = parts[1:]

	switch parts[0] {
	case "unicast":
		v := StreamUnicast
		ht.Cast = &v
		parts = parts[1:]

	case "multicast":
		v := StreamMulticast
		ht.Cast = &v
		parts = parts[1:]
	}

	for _, t := range parts {
		if strings.HasPrefix(t, "destination=") {
			v := t[len("destination="):]
			ht.Destination = &v

		} else if strings.HasPrefix(t, "ttl=") {
			v, err := strconv.ParseUint(t[len("ttl="):], 10, 64)
			if err != nil {
				return nil, err
			}
			vu := uint(v)
			ht.TTL = &vu

		} else if strings.HasPrefix(t, "port=") {
			ports, err := parsePorts(t[len("port="):])
			if err != nil {
				return nil, err
			}
			ht.Ports = ports

		} else if strings.HasPrefix(t, "client_port=") {
			ports, err := parsePorts(t[len("client_port="):])
			if err != nil {
				return nil, err
			}
			ht.ClientPorts = ports

		} else if strings.HasPrefix(t, "server_port=") {
			ports, err := parsePorts(t[len("server_port="):])
			if err != nil {
				return nil, err
			}
			ht.ServerPorts = ports

		} else if strings.HasPrefix(t, "interleaved=") {
			ports, err := parsePorts(t[len("interleaved="):])
			if err != nil {
				return nil, err
			}
			ht.InterleavedIds = ports

		} else if strings.HasPrefix(t, "mode=") {
			v := strings.ToLower(t[len("mode="):])
			v = strings.TrimPrefix(v, "\"")
			v = strings.TrimSuffix(v, "\"")
			ht.Mode = &v
		}

		// ignore non-standard keys
	}

	return ht, nil
}

// Write encodes a Transport header
func (ht *HeaderTransport) Write() HeaderValue {
	var vals []string

	if ht.Protocol == StreamProtocolUDP {
		vals = append(vals, "RTP/AVP")
	} else {
		vals = append(vals, "RTP/AVP/TCP")
	}

	if ht.Cast != nil {
		if *ht.Cast == StreamUnicast {
			vals = append(vals, "unicast")
		} else {
			vals = append(vals, "multicast")
		}
	}

	if ht.ClientPorts != nil {
		ports := *ht.ClientPorts
		vals = append(vals, "client_port="+strconv.FormatInt(int64(ports[0]), 10)+"-"+strconv.FormatInt(int64(ports[1]), 10))
	}

	if ht.ServerPorts != nil {
		ports := *ht.ServerPorts
		vals = append(vals, "server_port="+strconv.FormatInt(int64(ports[0]), 10)+"-"+strconv.FormatInt(int64(ports[1]), 10))
	}

	if ht.InterleavedIds != nil {
		ports := *ht.InterleavedIds
		vals = append(vals, "interleaved="+strconv.FormatInt(int64(ports[0]), 10)+"-"+strconv.FormatInt(int64(ports[1]), 10))
	}

	if ht.Mode != nil {
		vals = append(vals, "mode="+*ht.Mode)
	}

	return HeaderValue{strings.Join(vals, ";")}
}
