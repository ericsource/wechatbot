#### 接入
https://github.com/ericsource/tts-go

#### 安装
```bash
pip install edge-tts
go mod tidy
```

```bash
import "github.com/ericsource/tts-go/src/edge_tts"
```

```bash
edge-tts --text "Hello, world!" --write-media hello.mp3

edge-tts --voice ar-EG-SalmaNeural --text "Hello, world!" --write-media hello_in_arabic.mp3

edge-tts --list-voices

edge-tts --volume=-50% --text "Hello, world!" --write-media hello_with_volume_halved.mp3
```

```json
edge-tts --list-voices
Name: Microsoft Server Speech Text to Speech Voice (af-ZA, AdriNeural)
ShortName: af-ZA-AdriNeural
Gender: Female
Locale: af-ZA

Name: Microsoft Server Speech Text to Speech Voice (am-ET, MekdesNeural)
ShortName: am-ET-MekdesNeural
Gender: Female
Locale: am-ET

edge-tts --voice ar-EG-SalmaNeural --text "مرحبا كيف حالك؟" --write-media hello_in_arabic.mp3 --write-subtitles hello_in_arabic.vtt
```

```go
package main

import (
	"context"
	"fmt"
	"github.com/ericsource/tts-go/src/azure_tts"
	//"github.com/spf13/pflag"
)

func usage() {
	fmt.Println("usage: edge-tts [-h] [-t TEXT] [-f FILE] [-v VOICE] [-l] [--rate RATE] [--volume VOLUME] [--words-in-cue WORDS_IN_CUE] [--write-media WRITE_MEDIA] [--write-subtitles WRITE_SUBTITLES] [--proxy PROXY]\n")
	fmt.Println("Microsoft Edge TTS\n")
	fmt.Println("options:")
	pflag.PrintDefaults()
}

func main() {

}
```