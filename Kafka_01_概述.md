# kafka

#### 定义

kafka是一个分布式的基于发布、订阅模式的消息队列（Message Queue），主要用于大数据实时处理

![avatar](images\image-20200417153406955.png)

![avatar](images\image-20200417153854519.png)

![avatar](images\image-20200417153957657.png)

![avatar](images\image-20200417154140402.png)

![avatar](images\image-20200417154311823.png)

### kafka 架构

![avatar](images\image-20200417155956129.png)

1. producer : 消息生产者,就是想kafka broker发消息的客户端

2. consumer: 消息的消费者,向 kfaka broker去消息的客户端

3. consumerGroup (CG) : 消费者组,有多个consumer组成.