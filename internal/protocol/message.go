package protocol

type Message struct {
	Type       string `json:"type"`
	DeviceName string `json:"device_name,omitempty"`
	Text       string `json:"text,omitempty"`

	FileName string `json:"file_name,omitempty"`
	FileSize int64  `json:"file_size,omitempty"`
}
