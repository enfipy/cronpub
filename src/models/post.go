package models

import (
	"bytes"
	"encoding/gob"
)

type Post struct {
	FileType       FileType
	TelegramFileID string
	FileLink       string
}

func (post *Post) EncodeBinary() ([]byte, error) {
	bytesBuffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(bytesBuffer)
	err := encoder.Encode(post)
	if err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

func (post *Post) DecodeBinary(by []byte) error {
	bytesBuffer := new(bytes.Buffer)
	bytesBuffer.Write(by)
	decoder := gob.NewDecoder(bytesBuffer)
	return decoder.Decode(post)
}
