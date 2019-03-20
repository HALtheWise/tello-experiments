package gonet_test

import (
	"testing"
	"time"

	"github.com/HALtheWise/tello-experiments/gofly/gonet"
)

func TestGonet(t *testing.T) {
	const topic gonet.Topic = "topicname"
	net := gonet.NewNet()

	net.MakeNode("Talker", func(n *gonet.Context) {
		n.WaitForStart()

		n.Publish(topic, "This Is A Message")
	})

	net.MakeNode("Listener", func(n *gonet.Context) {
		c := n.Subscribe(topic)
		n.WaitForStart()

		select {
		case message := <-c:
			t.Logf("Got message: %+v\n", message)

		case <-time.After(1 * time.Second):
			t.Errorf("Message not recieved before timeout")

		}
	})

	net.Run()
}
