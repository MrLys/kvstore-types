package kvstoretypes

type (
	CommandPayload struct {
		I   string // id created by the auth
		M   byte   // Method (0x1, 0x2, 0x3, 0x4,)
		K   string
		V   string
		Ttl string
	}
	ResponsePayload struct {
		C string
		V string
	}
	AuthPayload struct {
		C string // client id
		S string // client secret
	}
)
