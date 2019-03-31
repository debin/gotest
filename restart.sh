#! /bin/bash
#默认进入的是登录用户的目录
#cd test/bee
# ps aux |grep hello | grep -v grep | awk '{print $2}' | xargs sudo kill -9

appname = "hello"

cd `dirname $0`
pidlist=`ps aux |grep $appname | grep -v grep | awk '{print $2}'`
if [ "$pidlist" = "" ]
then
    echo "no pid alive!"
else
for i in $pidlist
do
    echo "Kill the $1 process [ $i ]"
    sudo kill -9 $i
done
fi

sudo nohup ./`$appname` 2>&1 &
echo "$appname is start!"
echo `ps aux |grep $appname | grep -v grep`
echo
