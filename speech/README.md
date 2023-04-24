#### 接入
https://github.com/rany2/edge-tts

#### 安装
```bash
pip install edge-tts
```

```bash
edge-tts --text "Hello, world!" --write-media hello.mp3

edge-tts --voice ar-EG-SalmaNeural --text "Hello, world!" --write-media hello_in_arabic.mp3

edge-tts --list-voices

edge-tts --volume=-50% --text "Hello, world!" --write-media hello_with_volume_halved.mp3
```