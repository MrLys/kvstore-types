package kvstoretypes

type (
	CommandPayload struct {
		M   byte // Method (0x1, 0x2, 0x3, 0x4,)
		K   string
		V   string
		Ttl string
	}
	ResponsePayload struct {
		C string
		V string
	}
	MetaCommandPayload struct {
		A string         // auth string
		I string         // id string
		P CommandPayload // actual
	}
)
