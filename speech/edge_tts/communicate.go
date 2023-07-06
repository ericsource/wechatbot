package edge_tts

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
	"unicode/utf8"
)

// https://github.com/rany2/edge-tts/blob/a1bacbe1fb93de1233e434e33b865f2ba13150db/src/edge_tts/communicate.py#L187

func getHeadersAndData(data interface{}) (map[string]string, []byte, error) {
	var dataBytes []byte
	switch t := data.(type) {
	case string:
		dataBytes = []byte(t)
	case []byte:
		dataBytes = t
	default:
		return nil, nil, fmt.Errorf("data must be string or []byte")
	}

	headers := make(map[string]string)
	lines := bytes.Split(dataBytes, []byte("\r\n"))
	for _, line := range lines {
		if colonIndex := bytes.IndexByte(line, ':'); colonIndex != -1 {
			key := string(line[:colonIndex])
			value := string(bytes.TrimSpace(line[colonIndex+1:]))
			headers[key] = value
		}
	}

	dataStartIndex := bytes.Index(dataBytes, []byte("\r\n\r\n")) + 4
	if dataStartIndex == -1 {
		return nil, nil, fmt.Errorf("invalid data format")
	}
	d := dataBytes[dataStartIndex:]

	return headers, d, nil
}

func removeIncompatibleCharacters(input interface{}) (string, error) {
	var str string
	switch t := input.(type) {
	case string:
		str = t
	case []byte:
		str = string(t)
	default:
		return "", fmt.Errorf("input must be string or []byte")
	}

	runes := []rune(str)
	for idx, r := range runes {
		if (0 <= r && r <= 8) || (11 <= r && r <= 12) || (14 <= r && r <= 31) {
			runes[idx] = ' '
		}
	}

	return string(runes), nil
}

func splitTextByByteLength(text interface{}, byteLength int) <-chan []byte {
	ch := make(chan []byte)

	go func() {
		defer close(ch)

		var str string
		switch t := text.(type) {
		case string:
			str = t
		case []byte:
			str = string(t)
		default:
			return
		}

		runes := []rune(str)
		start := 0
		for len(runes) > 0 {
			end := findSplitPoint(runes, byteLength)
			if end == -1 {
				break
			}

			substr := string(runes[start:end])
			ch <- []byte(substr)

			start = end
		}

		if start < len(runes) {
			substr := string(runes[start:])
			ch <- []byte(substr)
		}
	}()

	return ch
}

func findSplitPoint(runes []rune, byteLength int) int {
	var length int
	for i, r := range runes {
		length += utf8.RuneLen(r)
		if length > byteLength {
			return i
		}
	}
	return -1
}

func mkssml(text interface{}, voice, rate, volume string) string {
	var str string
	switch t := text.(type) {
	case string:
		str = t
	case []byte:
		str = string(t)
	default:
		return ""
	}

	ssml := fmt.Sprintf(
		`<speak version="1.0" xmlns="http://www.w3.org/2001/10/synthesis" xml:lang="en-US">
			<voice name="%s">
				<prosody pitch="+0Hz" rate="%s" volume="%s">%s</prosody>
			</voice>
		</speak>`,
		voice, rate, volume, str,
	)

	return ssml
}

func connectID() string {
	uuidObj := uuid.New()
	return strings.ReplaceAll(uuidObj.String(), "-", "")
}

func dateToString() string {
	// time.Now().UTC() 返回当前的时间对象，在全球范围内是协调世界时 (Coordinated Universal Time, UTC)。
	// time.Format() 方法用于格式化时间对象为指定的字符串格式。
	// "%a %b %d %Y %H:%M:%S GMT+0000 (Coordinated Universal Time)" 是 JavaScript 风格的日期字符串格式。
	return time.Now().UTC().Format("Mon Jan 02 2006 15:04:05 GMT-0700 (Coordinated Universal Time)")
}

func ssmlHeadersPlusData(requestID, timestamp, ssml string) string {
	headers := fmt.Sprintf("X-RequestId:%s\r\n", requestID)
	headers += "Content-Type:application/ssml+xml\r\n"
	headers += fmt.Sprintf("X-Timestamp:%sZ\r\n", timestamp) // This is not a mistake, Microsoft Edge bug.
	headers += "Path:ssml\r\n\r\n"

	return headers + ssml
}

func calcMaxMesgSize(voice, rate, volume string) int {
	websocketMaxSize := 1 << 16
	overheadPerMessage := len(ssmlHeadersPlusData(connectID(), dateToString(), mkssml("", voice, rate, volume))) + 50
	return websocketMaxSize - overheadPerMessage
}
