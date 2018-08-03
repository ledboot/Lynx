package v1

import (
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"github.com/pkg/errors"
	"fmt"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	isClose   bool
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {

	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}

	go conn.readlLoop()
	go conn.writeLoop()
	return conn, nil
}

func (conn *Connection) readMsg() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}

	return
}

func (conn *Connection) writeMsg(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}

	return
}

func (conn *Connection) close() {
	conn.wsConn.Close()
	conn.mutex.Lock()
	if !conn.isClose {
		close(conn.closeChan)
		conn.isClose = true
	}
	conn.mutex.Unlock()
}

func WsHandler(ctx *gin.Context) {
	var (
		wsConn *websocket.Conn
		conn   *Connection
		data   []byte
		err    error
	)

	if wsConn, err = upgrader.Upgrade(ctx.Writer, ctx.Request, nil); err != nil {
		goto ERR
	}

	if conn, err = InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		for {
			if err = conn.writeMsg([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if data, err = conn.readMsg(); err != nil {
			fmt.Println("readMsg get err!")
			goto ERR
		}
		if conn.writeMsg(data); err != nil {
			fmt.Println("writeMsg get err!")
			goto ERR
		}
	}

ERR:
	conn.close()
}

func (conn *Connection) readlLoop() {
	var (
		data []byte
		err  error
	)

	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			goto ERR
		}

	}
ERR:
	conn.close()
}

func (conn *Connection) writeLoop() {

	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}

		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.close()
}
