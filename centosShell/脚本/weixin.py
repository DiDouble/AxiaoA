
#!/usr/bin/python2.7
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
        "toparty": 2,
        "msgtype": "text",
        "agentid": 1000004,
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

    Corpid = "wwaab3f7ceab21bf9e"
    Secret = "wouS10bkwBEDL2cHyXBdp-EuS9pRtcdCBi7kBbpZkCg"
    #Tagid = "1"
    Agentid = "1000004"
    Partyid = "2"

    Status = SendMessage(User,Agentid,Subject,Content)
    print Status
