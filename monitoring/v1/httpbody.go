package monitoring

func (h *HttpBody) UnmarshalJSON(data []byte) error {

	h.Data = string(data)

	return nil
}
