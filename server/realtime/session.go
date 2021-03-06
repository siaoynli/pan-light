package realtime

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type SessionId int64

type Session struct {
	Data interface{} // 业务数据存放

	id               SessionId // session 标识
	secret           string    // session 秘钥
	server           *Server
	conn             *websocket.Conn
	online           bool // 是否掉线了
	lock             sync.Mutex
	missMessage      []gson // 掉线错过的消息
	missMessageIndex int
	rooms            []*Room
}

func (s *Session) Rooms() []*Room {
	return s.rooms
}

func randomStr(lenght int) string {
	arr := make([]byte, lenght)
	src := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	for i := 0; i < lenght; i++ {
		arr[i] = byte(src[rand.Intn(len(src))])
	}
	return string(arr)
}

func newSession(conn *websocket.Conn, missMessageSize int, server *Server) *Session {
	s := &Session{
		id:          SessionId(time.Now().UnixNano()),
		secret:      randomStr(16),
		server:      server,
		conn:        conn,
		online:      true,
		missMessage: make([]gson, missMessageSize),
	}
	s.Emit("session.new", gson{
		"id":     strconv.FormatInt(int64(s.id), 10),
		"secret": s.secret,
	})
	return s
}

func (ss *Session) Id() SessionId {
	return ss.id
}

func (ss *Session) Emit(event string, data interface{}) {
	ss.write(gson{
		"type":    "event",
		"event":   event,
		"payload": data,
	})
}

func (ss *Session) write(data gson) (err error) {
	ss.lock.Lock()
	defer ss.lock.Unlock()
	if !ss.online {
		ss.missMessageIndex = (ss.missMessageIndex + 1) % len(ss.missMessage)
		ss.missMessage[ss.missMessageIndex] = gson{
			"time":    time.Now().Unix(),
			"message": data,
		}
		return errors.New("connection loss")
	}
	bin, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "json encode error")
	}
	if enc {
		err = websocket.Message.Send(ss.conn, encBin(bin))
	} else {
		_, err = ss.conn.Write(bin)
	}
	return
}

func (ss *Session) Read() (data gson, err error) {
	bin, err := receiveFullFrame(ss.conn)
	if err != nil {
		return
	}
	err = json.Unmarshal(bin, &data)
	return
}
func (ss *Session) InRoom(name string) bool {
	ss.lock.Lock()
	defer ss.lock.Unlock()
	for _, room := range ss.rooms {
		if room.name == name {
			return true
		}
	}
	return false
}

const enc = false
const key = "pan-light"

func xorBin(bin []byte) []byte {
	dst := make([]byte, len(bin))
	keyLen := len(key)
	for idx, b := range bin {
		dst[idx] = key[idx%keyLen] ^ b
	}
	return dst
}

func encBin(bin []byte) []byte {
	if !enc {
		return bin
	}
	buf := bytes.NewBuffer(make([]byte, 0, len(bin)))
	w := gzip.NewWriter(buf)
	w.Write(bin)
	w.Flush()
	return xorBin(buf.Bytes())
}
func dencBin(bin []byte) (dest []byte, err error) {
	if !enc {
		return bin, nil
	}
	buf := bytes.NewReader(xorBin(bin))
	r, err := gzip.NewReader(buf)
	if err != nil {
		return
	}
	dest, err = ioutil.ReadAll(r)
	return
}

// 接受完整帧
func receiveFullFrame(ws *websocket.Conn) ([]byte, error) {
	var data []byte
	for {
		var seg []byte
		fin, err := receiveFrame(websocket.Message, ws, &seg)
		if err != nil {
			return nil, err
		}
		data = append(data, seg...)
		if fin {
			break
		}
	}
	return dencBin(data)
}

// 接受帧
func receiveFrame(cd websocket.Codec, ws *websocket.Conn, v interface{}) (fin bool, err error) {
again:
	frame, err := ws.NewFrameReader()
	if frame.HeaderReader() != nil {
		bin := make([]byte, 1)
		frame.HeaderReader().Read(bin)
		fin = ((bin[0] >> 7) & 1) != 0
	}
	if err != nil {
		return
	}
	frame, err = ws.HandleFrame(frame)
	if err != nil {
		return
	}
	if frame == nil {
		goto again
	}

	payloadType := frame.PayloadType()

	data, err := ioutil.ReadAll(frame)
	if err != nil {
		return
	}
	return fin, cd.Unmarshal(data, payloadType, v)
}
