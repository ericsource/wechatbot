#### 接入
https://github.com/rany2/edge-tts

https://github.com/ericsource/tts-go

#### 安装
```bash
pip install edge-tts
go mod tidy
```

```bash
import "github.com/869413421/wechatbot/speech/edge_tts"
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