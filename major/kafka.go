package major

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
)

var (
	client sarama.AsyncProducer
)

func asyncKafkaProducer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true //必须有这个选项
	// config.Producer.Timeout = 5 * time.Second
	client, err := sarama.NewAsyncProducer(strings.Split("192.168.27.124:9092", ","), config)
	defer client.Close()
	if err != nil {
		return
	}
	//必须有这个匿名函数内容
	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		success := p.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					fmt.Println("发送失败: ", err)
				}
			case <-success:
			}
		}
	}(client)
}

func SendKafkaMsg(topic string, msg []byte) {
	client.Input() <- &sarama.ProducerMessage{Topic: topic, Value: sarama.ByteEncoder(msg)}
}
