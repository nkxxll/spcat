package renderer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
)

var (
	KittyOptions = map[string]string{
		"a": "T",
		"f": "100",
	}
	prefix  = []byte{'\033', '_', 'G'}
	postfix = []byte{'\033', '\\'}
)

func optionsToBytes(options map[string]string) []byte {
	var b bytes.Buffer
	res := []string{}
	for k, v := range options {
		res = append(res, fmt.Sprintf("%s=%s", k, v))
	}
	b.WriteString(strings.Join(res, ","))
	b.WriteByte(';')
	return b.Bytes()
}

func generateChunk(b *bytes.Buffer, chunk []byte, options []byte) {
	b.Write(prefix)
	b.Write(options)
	b.Write(chunk)
	b.Write(postfix)
}

func GenerateKittyPic(payload []byte, options map[string]string) []byte {
	var b bytes.Buffer

	bufEnc := make([]byte, base64.StdEncoding.EncodedLen(len(payload)))
	base64.StdEncoding.Encode(bufEnc, payload)
	chunk := make([]byte, 4096)
	options["m"] = "1"
	ob := optionsToBytes(options)

	for len(bufEnc) >= 4096 {
		chunk, bufEnc = bufEnc[:4096], bufEnc[4096:]
		generateChunk(&b, chunk, ob)
	}

	options["m"] = "0"
	generateChunk(&b, bufEnc, optionsToBytes(options))

	return b.Bytes()
}
