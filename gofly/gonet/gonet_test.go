package gonet_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/HALtheWise/tello-experiments/gofly/gonet"
)

func TestPubSub(t *testing.T) {
	const topic gonet.Topic = "topicname"
	net := gonet.NewNet()

	net.MakeNode("Talker", func(n *gonet.Context) {
		n.WaitForStart()

		n.Publish(topic, "This Is A Message")
	})

	net.MakeNode("Listener", func(c *gonet.Context) {
		ch := c.Subscribe(topic)
		c.WaitForStart()

		select {
		case message := <-ch:
			t.Logf("Got message: %+v\n", message)

		case <-time.After(1 * time.Second):
			t.Error("Message not recieved before timeout")
			return
		}
	})

	net.Run()
}

func TestUniversalSubs(t *testing.T) {
	net := gonet.NewNet()

	net.MakeNode("Publisher A", func(c *gonet.Context) {
		c.WaitForStart()
		c.Publish("topic1", "data1")
		c.Publish("topic2", 2)
	})

	golden := []gonet.MessageInfo{
		gonet.MessageInfo{Topic: "topic1", Data: "data1"},
		gonet.MessageInfo{Topic: "topic2", Data: 2},
	}

	net.MakeNode("Logger", func(c *gonet.Context) {
		ch := c.SubscribeAll()

		var recieved []gonet.MessageInfo

		c.WaitForStart()

	loop:
		for len(recieved) < len(golden) {
			select {
			case msg := <-ch:
				recieved = append(recieved, msg)

			case <-time.After(1 * time.Second):
				t.Error("Message not recieved before timeout")
				break loop
			}
		}

		if !reflect.DeepEqual(recieved, golden) {
			t.Errorf("Incorrect data recieved, expected\n%#v got\n%#v", golden, recieved)
		}
	})

	net.Run()
}
