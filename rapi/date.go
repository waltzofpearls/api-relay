package rapi

type APIDate struct {
	original []byte
	date *time.Date
}

func (ad *APIDate) Unmarshal(data []byte) error {
}

func (ad *APIDate) Marshal() ([]byte, error) {
}
