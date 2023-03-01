# 项目来自
https://github.com/djun/wechatbot

# wechatbot
项目基于[openwechat](https://github.com/eatmoreapple/openwechat) 最新版本 开发
###目前实现了以下功能
 + 群聊@回复
 + 私聊回复
 + 自动通过回复
 + 回复心跳（但还是会出现微信退出，openwechat文档有说需要两微信相互发消息才不会退出）
 + 添加telegram接入
 
# 注册openai
chatGPT注册可以参考[这里](https://juejin.cn/post/7173447848292253704)

# 安装使用
```
# 获取项目
git clone https://github.com/ericsource/wechatbot.git

# 进入项目目录
cd wechatbot

# 复制配置文件
copy config.dev.json config.json

# 启动项目
go run main.go

启动前需替换config中的api_key

# teletgram 配置文件
"telegram_proxy": "sock5://127.0.0.1:7890",
"telegram_token": ""

# 启动telegram
go run telegram.go

```

# 测试代码
```
curl --socks5 127.0.0.1:1080 http://httpbin.org/ip
```

# 打包

```
./build.sh
```

# 运行
```
nohup ./restart.sh >/dev/null 2>log &
```

