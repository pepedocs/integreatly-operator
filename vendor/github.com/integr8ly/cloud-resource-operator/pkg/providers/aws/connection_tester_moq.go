// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package aws

import (
	"sync"
)

var (
	lockConnectionTesterMockTCPConnection sync.RWMutex
)

// Ensure, that ConnectionTesterMock does implement ConnectionTester.
// If this is not the case, regenerate this file with moq.
var _ ConnectionTester = &ConnectionTesterMock{}

// ConnectionTesterMock is a mock implementation of ConnectionTester.
//
//     func TestSomethingThatUsesConnectionTester(t *testing.T) {
//
//         // make and configure a mocked ConnectionTester
//         mockedConnectionTester := &ConnectionTesterMock{
//             TCPConnectionFunc: func(host string, port int) bool {
// 	               panic("mock out the TCPConnection method")
//             },
//         }
//
//         // use mockedConnectionTester in code that requires ConnectionTester
//         // and then make assertions.
//
//     }
type ConnectionTesterMock struct {
	// TCPConnectionFunc mocks the TCPConnection method.
	TCPConnectionFunc func(host string, port int) bool

	// calls tracks calls to the methods.
	calls struct {
		// TCPConnection holds details about calls to the TCPConnection method.
		TCPConnection []struct {
			// Host is the host argument value.
			Host string
			// Port is the port argument value.
			Port int
		}
	}
}

// TCPConnection calls TCPConnectionFunc.
func (mock *ConnectionTesterMock) TCPConnection(host string, port int) bool {
	if mock.TCPConnectionFunc == nil {
		panic("ConnectionTesterMock.TCPConnectionFunc: method is nil but ConnectionTester.TCPConnection was just called")
	}
	callInfo := struct {
		Host string
		Port int
	}{
		Host: host,
		Port: port,
	}
	lockConnectionTesterMockTCPConnection.Lock()
	mock.calls.TCPConnection = append(mock.calls.TCPConnection, callInfo)
	lockConnectionTesterMockTCPConnection.Unlock()
	return mock.TCPConnectionFunc(host, port)
}

// TCPConnectionCalls gets all the calls that were made to TCPConnection.
// Check the length with:
//     len(mockedConnectionTester.TCPConnectionCalls())
func (mock *ConnectionTesterMock) TCPConnectionCalls() []struct {
	Host string
	Port int
} {
	var calls []struct {
		Host string
		Port int
	}
	lockConnectionTesterMockTCPConnection.RLock()
	calls = mock.calls.TCPConnection
	lockConnectionTesterMockTCPConnection.RUnlock()
	return calls
}
