package events

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	otherEvent      TestEvent
	handler         TestEventHandler
	otherHandler    TestEventHandler
	anotherHandler  TestEventHandler
	eventDispatcher *EventDispatcher
}

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (te *TestEvent) GetName() string {
	return te.Name
}

func (te *TestEvent) GetPayload() interface{} {
	return te.Payload
}

func (te *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (teh *TestEventHandler) Handle(event EventInterface) {
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()

	suite.handler = TestEventHandler{1}
	suite.otherHandler = TestEventHandler{2}
	suite.anotherHandler = TestEventHandler{3}

	suite.event = TestEvent{Name: "testEvent", Payload: "testEventPayload"}
	suite.otherEvent = TestEvent{Name: "otherTestEvent", Payload: "otherTestEventPayload"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_Success() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)

	//assert.Nil(suite.T(), err)
	suite.Nil(err)
	//assert.Equal(suite.T(), 1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.otherHandler)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// assert if the handler is already registered correctly
	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.otherHandler, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_ErrorAlreadyRegistered() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.Equal(suite.T(), ErrHandlerAlreadyRegistered, err)
	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), 1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}
