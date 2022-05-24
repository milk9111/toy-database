package serialize

import (
	"encoding/gob"
	"os"
	"strings"
)

type Pager interface {
	Read(input interface{}) error
	Write(input interface{}) error
}

type pager struct {
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func OpenPager(filename string) (Pager, error) {
	if strings.TrimSpace(filename) == "" {
		return pager{}, nil
	}

	_, err := os.Stat(filename)

	var file *os.File
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(filename)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		file, err = os.OpenFile(filename, os.O_RDWR, 0644)
		if err != nil {
			return nil, err
		}
	}

	pager := pager{
		encoder: gob.NewEncoder(file),
		decoder: gob.NewDecoder(file),
	}

	return pager, nil
}

func (pager pager) Read(output interface{}) error {
	if pager.decoder == nil {
		return ErrSerializeResultNilDecoder
	}

	return pager.decoder.Decode(output)
}

func (pager pager) Write(input interface{}) error {
	if pager.encoder == nil {
		return ErrSerializeResultNilEncoder
	}

	return pager.encoder.Encode(input)
}
