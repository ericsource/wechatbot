#!/bin/sh

while :

do

cd "/home/eric.cai/wechatbot/bin"

echo "Current DIR is " $PWD

stillRunning=$(ps -ef |grep "$PWD/wechatbot-386-linux" |grep -v "grep")

if [ "$stillRunning" ] ; then

echo "TWS service was already started by another way"

#echo "Kill it and then startup by this shell, other wise this shell will loop out this message annoyingly"
#kill -9 $pidof $PWD/wechatbot-386-linux

else

echo "TWS service was not started"

echo "Starting service ..."

nohup $PWD/wechatbot-386-linux >chat.log 2>&1 &

echo "TWS service was exited!"

fi

sleep 10

done