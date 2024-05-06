package kvstoretypes

type (
	CommandPayload struct {
		M   byte
		K   string
		V   string
		Ttl string
	}
	ResponsePayload struct {
		C string
		V string
	}
)
