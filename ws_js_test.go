package websocket

func TestEcho(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c, resp, err := websocket.Dial(ctx, os.Getenv("WS_ECHO_SERVER_URL"), &websocket.DialOptions{
		Subprotocols: []string{"echo"},
	})
	assert.Success(t, err)
	defer c.Close(websocket.StatusInternalError, "")

	assertSubprotocol(t, c, "echo")
	assert.Equalf(t, &http.Response{}, resp, "unexpected http response")
	assertJSONEcho(t, ctx, c, 1024)
	assertEcho(t, ctx, c, websocket.MessageBinary, 1024)

	err = c.Close(websocket.StatusNormalClosure, "")
	assert.Success(t, err)
}
