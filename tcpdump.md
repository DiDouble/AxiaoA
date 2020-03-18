# tcpdump抓包工具详解

#### 第一类关键字抓包  包括host ,net, port 定义捕获数据包范围

```shell
#解惑主机包
tcpdump host  主机IP  
#截获192.168.0.0/24该网段的数据包
tcpdump net 192.168.0.0/24
#截获主机为IP1和IP2或者IP3的数据包
tcpdump -nn host IP1 and '(IP2 or IP3)'
#截获主机IP1和除了IP2通信的数据
tcpdump -nn host IP1 and! IP2
#截获端口80端口收到和发出的数据包
tcpdump host IP and prot 80

```

#### 第二种传输方向关键字

* src ,  #源

*  dst ,  # 目标

*  dst or src ,  # 目标或者源

*  dst and src ,  # 目标和源

  ```shell
  #截获源IP为IP1的并且目标是IP2的
  tcpdump -nn src IP1 and dst IP2
  #截获源ip为IP1的源端口为80 并且目标IP的IP2端口为22的
  tcpdump -nn src IP1 anf src port 80 and dst IP2 and dst port 22
  ```

  #### 第三种关键字查询

  主要包括

  * ffdi

  * ip

  * arp

  * rarp

  * tcp

  * udp

  * icmp

    ```shell
    #截获主机ip1 端口22 的tcp协议的数据包
    tcpdump -nn -v host ip1 and port22 and tcp
    #抓取网口eth0 上ip1 与除了ip2 以外的其他主机之间的icmp报文
    tcpdump -i eth0 -s 1400 -nn host ip1 and! ip2 and icmp -e
    
    ```

    