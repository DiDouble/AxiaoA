# 微信报警

### 企业微信需要用到的配置

0. 创建应用

   ![image-20200511105115015](E:\AllProject\src\AxiaoA\images\image-20200511105115015.png)

1. 维系报警需要用到的参数

   ![image-20200511104346250](E:\AllProject\src\AxiaoA\images\image-20200511104346250.png)

   2. ![image-20200511104613541](E:\AllProject\src\AxiaoA\images\image-20200511104613541.png)

3. ![image-20200511104813268](E:\AllProject\src\AxiaoA\images\image-20200511104813268.png)

4. ![image-20200511105013870](E:\AllProject\src\AxiaoA\images\image-20200511105013870.png)

下面是代码

```python
#!/usr/bin/python
#_*_coding:utf-8 _*_
import requests,sys,json
import urllib3
urllib3.disable_warnings()
reload(sys)
sys.setdefaultencoding('utf-8')
def GetTokenFromServer(Corpid,Secret):
    Url = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
    Data = {
        "corpid":Corpid,
        "corpsecret":Secret
    }
    r = requests.get(url=Url,params=Data,verify=False)
    print(r.json())
    if r.json()['errcode'] != 0:
        return False
    else:
        Token = r.json()['access_token']
        file = open('/tmp/zabbix_wechat_config.json', 'w')
        file.write(r.text)
        file.close()
        return Token
def SendMessage(User,Agentid,Subject,Content):
    try:
        file = open('/tmp/zabbix_wechat_config.json', 'r')
        Token = json.load(file)['access_token']
        file.close()
    except:
        Token = GetTokenFromServer(Corpid, Secret)
    n = 0
    Url = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s" % Token
    Data = {
        #"touser": User,
        #"totag": 2,
        "toparty": 1,
        "msgtype": "text",
        "agentid": xxxx,
        "text": {
            "content": Subject + '\n' + Content
        },
        "safe": "0"
    }
    r = requests.post(url=Url,data=json.dumps(Data),verify=False)
    while r.json()['errcode'] != 0 and n < 4:
        n+=1
        Token = GetTokenFromServer(Corpid, Secret)
        if Token:
            Url = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s" % Token
            r = requests.post(url=Url,data=json.dumps(Data),verify=False)
            print(r.json())
    return r.json()
if __name__ == '__main__':
    User = sys.argv[1]
    Subject = str(sys.argv[2])
    Content = str(sys.argv[3])
    Corpid = "xxxxxxxxxxxxx"
    Secret = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    #Tagid = "1"
    Agentid = "xxxxxx"
    Partyid = "2"
    Status = SendMessage(User,Agentid,Subject,Content)
    print Status
```

这里要说明一下,由于windows和linux默认的字符集是不一样的,所以需要调整一下,不然在测试的时候会报错,一直报一个/usr/bin/python^E,

修改方法

1. chmod +x  weixin.py

2. mv  weixin.py   /usr/lib/zabbix/alertscripts/

3. chown zabbix.zabbix weixin

4. vim weixin.py

5. : set ff 或者 :set fileformat

   1. 如果看到的是 fileformat=doc
      1. 执行 :set ff=unix或者:set fileformat=unix
   2. 如果出现的是 fileformat=unix
      1. 可以直接保存退出

6. 然后再执行  

   1. ```shell
      [root@zabbix /usr/lib/zabbix/alertscripts 09:59:08 37 ]# /usr/lib/zabbix/alertscripts/weixin.py name  "title test" "hello"
      {u'invaliduser': u'', u'errcode': 0, u'errmsg': u'ok. Warning: wrong json format. '}
      ```

      企业微信可以收到消息了.

## zabbix配置

![image-20200511110214848](E:\AllProject\src\AxiaoA\images\image-20200511110214848.png)

在管理页面点击用户点击创建用户,就可以创建一个叫微信的用户.

![image-20200511110317090](E:\AllProject\src\AxiaoA\images\image-20200511110317090.png)

修改用户数据.

![image-20200511110352953](E:\AllProject\src\AxiaoA\images\image-20200511110352953.png)

添加报警媒介.

![image-20200511114206112](E:\AllProject\src\AxiaoA\images\image-20200511114206112.png)

![image-20200511114256191](E:\AllProject\src\AxiaoA\images\image-20200511114256191.png)

![image-20200511114325383](E:\AllProject\src\AxiaoA\images\image-20200511114325383.png)

参数   ----这里的参数跟脚本中需要传三个参数对应.

```shell
{ALERT.SENDTO}
{ALERT.SUBJECT}
{ALERT.MESSAGE}
```

配置动作

![image-20200511174117685](E:\AllProject\src\AxiaoA\images\image-20200511174117685.png)

![image-20200511174154751](E:\AllProject\src\AxiaoA\images\image-20200511174154751.png)

消息内容我这里用的是汉语的,但是到了微信中发现就是16进制的.

