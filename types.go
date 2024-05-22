package kvstoretypes

type (
	CommandPayload struct {
		I   string // id created by the auth
		M   byte   // Method (0x1, 0x2, 0x3, 0x4,)
		K   string // Key
		V   string // Value
		Ttl string // Time to live
	}
	ResponsePayload struct {
		I string // id created by the auth
		C string // Command
		V string // Value
	}
)
