package pubsub

import (
	// go packages
	"fmt"
	"strings"
	// 3party packages
	zmq "github.com/alecthomas/gozmq"
)

// pub-sub message
type Message struct {
	Type    string
	Channel string
	Data    string
}

// ZMQ Client Subscribe to a Channel
func (client *ZMQClient) Subscribe(channels ...interface{}) error {
	for _, channel := range channels {
		client.sub.SetSockOptString(zmq.SUBSCRIBE, channel.(string))
	}
	return nil
}

// ZMQ client Unsubscribe from a channel
func (client *ZMQClient) Unsubscribe(channels ...interface{}) error {
	for _, channel := range channels {
		client.sub.SetSockOptString(zmq.UNSUBSCRIBE, channel.(string))
	}
	return nil
}

// ZMQ client Publish a Message
func (client *ZMQClient) Publish(channel, message string) {
	client.pub.Send([]byte(channel+" "+message), 0)
}

// ZMQ client Receive a Message
func (client *ZMQClient) Receive() Message {
	message, _ := client.sub.Recv(0)
	parts := strings.SplitN(string(message), " ", 2)
	return Message{Type: "message", Channel: parts[0], Data: parts[1]}
}

// Client interface : Subscribe, Unsubscribe, Publish, Receive
type Client interface {
	Subscribe(channels ...interface{}) (err error)
	Unsubscribe(channels ...interface{}) (err error)
	Publish(channel string, message string)
	Receive() (message Message)
}

// ZMQ Client
type ZMQClient struct {
	pub *zmq.Socket
	sub *zmq.Socket
}

// NewZMQClient to use it
func NewZMQClient(host string) *ZMQClient {
	context, _ := zmq.NewContext()
	pub, _ := context.NewSocket(zmq.PUSH)
	pub.Connect(fmt.Sprintf("tcp://%s:%d", host, 8080))
	sub, _ := context.NewSocket(zmq.SUB)
	sub.Connect(fmt.Sprintf("tcp://%s:%d", host, 8081))
	return &ZMQClient{pub, sub}
}
