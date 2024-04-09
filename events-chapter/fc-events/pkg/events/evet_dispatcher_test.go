package events

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"
)

type EventDispatcherTestSuite struct {
	suite.Suite
	firstEvent         TestEvent
	secondEvent        TestEvent
	testEventHandler   TestEventHandler
	secondEventHandler TestEventHandler
	thirdEventHandler  TestEventHandler
	eventDispatcher    *EventDispatcher
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

func (teh *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()

	suite.testEventHandler = TestEventHandler{1}
	suite.secondEventHandler = TestEventHandler{2}
	suite.thirdEventHandler = TestEventHandler{3}

	suite.firstEvent = TestEvent{Name: "firstEvent", Payload: "testEventPayload"}
	suite.secondEvent = TestEvent{Name: "otherTestEvent", Payload: "otherTestEventPayload"}
}

type MockEventHandler struct {
	mock.Mock
}

func (meh *MockEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	meh.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_Success() {
	err := suite.eventDispatcher.Register(suite.firstEvent.GetName(), &suite.testEventHandler)

	//assert.Nil(suite.T(), err)
	suite.Nil(err)
	//assert.Equal(suite.T(), 1, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))

	err = suite.eventDispatcher.Register(suite.firstEvent.GetName(), &suite.secondEventHandler)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))

	// assert if the testEventHandler is already registered correctly
	assert.Equal(suite.T(), &suite.testEventHandler, suite.eventDispatcher.handlers[suite.firstEvent.GetName()][0])
	assert.Equal(suite.T(), &suite.secondEventHandler, suite.eventDispatcher.handlers[suite.firstEvent.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_ErrorAlreadyRegistered() {
	err := suite.eventDispatcher.Register(suite.firstEvent.GetName(), &suite.testEventHandler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))

	err = suite.eventDispatcher.Register(suite.firstEvent.GetName(), &suite.testEventHandler)
	assert.Equal(suite.T(), ErrHandlerAlreadyRegistered, err)
	assert.Equal(suite.T(), &suite.testEventHandler, suite.eventDispatcher.handlers[suite.firstEvent.GetName()][0])
	assert.Equal(suite.T(), 1, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear_Success() {
	// firstEvent 1 = firstEvent
	err := suite.eventDispatcher.Register(suite.firstEvent.GetName(), &suite.testEventHandler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))

	err = suite.eventDispatcher.Register(suite.firstEvent.GetName(), &suite.secondEventHandler)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))

	// firstEvent 2 = secondEvent
	err = suite.eventDispatcher.Register(suite.secondEvent.GetName(), &suite.thirdEventHandler)
	suite.Nil(err)

	err = suite.eventDispatcher.Clear()
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 0, len(suite.eventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has_Success() {
	// firstEvent 1 = firstEvent
	err := suite.eventDispatcher.Register(suite.firstEvent.GetName(), &suite.testEventHandler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))

	err = suite.eventDispatcher.Register(suite.firstEvent.GetName(), &suite.secondEventHandler)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))

	// verify if the testEventHandler and secondEventHandler is registered using Has
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.firstEvent.GetName(), &suite.testEventHandler))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.firstEvent.GetName(), &suite.secondEventHandler))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.firstEvent.GetName(), &suite.thirdEventHandler))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch_Success() {
	eh := &MockEventHandler{}
	eh.On("Handle", &suite.firstEvent)

	eh2 := &MockEventHandler{}
	eh2.On("Handle", &suite.firstEvent)

	err := suite.eventDispatcher.Register(suite.firstEvent.GetName(), eh)
	suite.Nil(err)
	err = suite.eventDispatcher.Register(suite.firstEvent.GetName(), eh2)
	suite.Nil(err)

	err = suite.eventDispatcher.Dispatch(&suite.firstEvent)
	suite.Nil(err)

	eh.AssertExpectations(suite.T())
	eh2.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove_Success() {
	err := suite.eventDispatcher.Register(suite.firstEvent.GetName(), &suite.testEventHandler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))

	err = suite.eventDispatcher.Register(suite.firstEvent.GetName(), &suite.secondEventHandler)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))

	err = suite.eventDispatcher.Remove(suite.firstEvent.GetName(), &suite.testEventHandler)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(suite.eventDispatcher.handlers[suite.firstEvent.GetName()]))
	assert.Equal(suite.T(), &suite.secondEventHandler, suite.eventDispatcher.handlers[suite.firstEvent.GetName()][0])
}
