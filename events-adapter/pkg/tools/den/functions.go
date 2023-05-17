package den

import (
	"bytes"
	"events-adapter/pkg/errs"
	"github.com/sirupsen/logrus"
	"github.com/ugorji/go/codec"
)

// DecodeJson десериализует json в переданную структуру с помощью codec
func DecodeJson(to any, from []byte) *errs.Error {
	const source = "DecodeJson"

	handle := new(codec.JsonHandle)
	decoder := codec.NewDecoderBytes(from, handle)

	err := decoder.Decode(to)
	if err != nil {
		return errs.NewError(logrus.WarnLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryFunc, nil)
	}

	return nil
}

// EncodeJson кодирует информацию в json
func EncodeJson(from any) (*bytes.Buffer, *errs.Error) {
	const source = "EncodeJson"

	var to []byte

	handle := new(codec.JsonHandle)
	encoder := codec.NewEncoderBytes(&to, handle)

	err := encoder.Encode(from)
	if err != nil {
		return nil, errs.NewError(logrus.WarnLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryFunc, nil)
	}

	buf := new(bytes.Buffer)
	buf.WriteString(string(to))

	return buf, nil
}
