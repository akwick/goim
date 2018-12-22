package logic

import (
	"context"
	"testing"

	"github.com/Terry-Mao/goim/api/comet/grpc"
	"github.com/issue9/assert"
)

func TestConnect(t *testing.T) {
	var (
		server    = "test_server"
		serverKey = "test_server_key"
		token     = []byte(`{"mid":1, "key":"test_server_key", "room_id":"test://test_room", "platform":"web", "accepts":[1000,1001,1002]}`)
		ol        = map[string]int32{"test://test_room": 100}
		c         = context.Background()
	)
	// connect
	mid, key, roomID, _, accepts, err := lg.Connect(c, server, serverKey, "", token)
	assert.Nil(t, err)
	assert.Equal(t, serverKey, key)
	assert.Equal(t, roomID, "test://test_room")
	assert.Equal(t, len(accepts), 3)
	t.Log(mid, key, roomID, accepts, err)
	// heartbeat
	err = lg.Heartbeat(c, mid, key, server)
	assert.Nil(t, err)
	// disconnect
	has, err := lg.Disconnect(c, mid, key, server)
	assert.Nil(t, err)
	assert.Equal(t, true, has)
	// renew
	online, err := lg.RenewOnline(c, server, ol)
	assert.Nil(t, err)
	assert.NotNil(t, online)
	// message
	err = lg.Receive(c, mid, &grpc.Proto{})
	assert.Nil(t, err)
}
