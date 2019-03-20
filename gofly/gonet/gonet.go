package gonet

import (
	"sync"
)

type Topic string

type Net struct {
	nodesinitialized sync.WaitGroup
	nodesrunning     sync.WaitGroup

	subscribers          map[Topic][]chan interface{}
	universalsubscribers []chan MessageInfo
	nodes                []*Context
}

type Context struct {
	// initialzed is set to true when all nodes have called WaitForStart
	// and disables creating subscribers, while enabling publishing
	// This prevents race conditions on node startup
	Initialized bool

	net *Net

	Name string

	mainfunc func(*Context)

	// Initializing will be closed when the node finishes initializing
	initializing chan struct{}
	// Waiting will be closed when the node should start the main loop
	waiting chan struct{}
}

// WaitForStart blocks until all nodes are calling it.
// Every node must call this when it has finished initializing.
func (c *Context) WaitForStart() {
	close(c.initializing)
	c.Initialized = true
	<-c.waiting
}

func (n *Net) MakeNode(name string, node func(*Context)) {
	n.nodes = append(n.nodes, &Context{
		net:          n,
		Name:         name,
		mainfunc:     node,
		initializing: make(chan struct{}),
		waiting:      make(chan struct{}),
	})
}

// Run runs all of the previously provided functions,
// and blocks until they have all completed.
func (n *Net) Run() {
	n.nodesrunning.Add(len(n.nodes))

	// Initialize all nodes
	for _, context := range n.nodes {
		// Make a local variable pointing to the slice element
		// Prevents the pointer from changing under us
		context := context

		// start := time.Now()

		go func() {
			context.mainfunc(context)
			n.nodesrunning.Done()
		}()
		<-context.initializing

		// end := time.Now()
		// log.Printf("Node %#v initialized in %v",
		// 	context.Name,
		// 	end.Sub(start))
	}

	// Run nodes
	for _, c := range n.nodes {
		close(c.waiting)
	}

	// Let them run their due course
	n.nodesrunning.Wait()
}

func NewNet() *Net {
	n := new(Net)
	n.subscribers = make(map[Topic][]chan interface{})
	return n
}

func (c *Context) Publish(topic Topic, data interface{}) {
	n := c.net
	if !c.Initialized {
		panic("Publish called before all nodes initialized")
	}

	for _, c := range n.subscribers[topic] {
		c <- data
	}

	for _, c := range n.universalsubscribers {
		c <- MessageInfo{
			Topic: topic,
			Data:  data,
		}
	}
}

func (c *Context) Subscribe(topic Topic) chan interface{} {
	n := c.net
	if c.Initialized {
		panic("Subscriber created after initialization is complete")
	}
	ch := make(chan interface{}, 1)
	n.subscribers[topic] = append(n.subscribers[topic], ch)
	return ch
}
