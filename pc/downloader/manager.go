package downloader

import (
	"bytes"
	"net/http"
	"sync"
	"time"
)

type TaskId int64
type LinkResolver func(fileId string) (link string, err error) // 由file id转link的回调

// 下载管理器
type Manager struct {
	CoroutineNumber       int   // 每个任务的协程数量
	SegmentSize           int64 // 分配片段时选取的长度
	WroteToDiskBufferSize int64 // 每多少字节写一次磁盘, 过大会导致内存占用太高, 过小会导致磁盘压力大
	LinkResolver          LinkResolver
	EventChan             chan *DownloadEvent
	HttpClient            *http.Client
	bufferPool            sync.Pool // 下载缓存池
	taskMap               map[TaskId]*Task
}

// 初始化
func (m *Manager) Init() error {
	m.EventChan = make(chan *DownloadEvent, 1)
	m.taskMap = map[TaskId]*Task{}
	m.bufferPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, m.WroteToDiskBufferSize))
		},
	}
	return nil
}

// 从磁盘文件恢复, 下载中的任务
func (*Manager) Resume(raw map[TaskId][]byte) error {
	return nil
}

// 任务快照, 用于保存当前进度
func (*Manager) Capture() (map[TaskId][]byte, error) {
	return nil, nil
}

// 新建任务
func (m *Manager) NewTask(fileId, savePath string,
	requestDecorator func(*http.Request) *http.Request) (id TaskId, err error) {
	id = TaskId(time.Now().UnixNano())
	task := &Task{
		Id:               id,
		fileId:           fileId,
		manager:          m,
		linkResolver:     m.LinkResolver,
		requestDecorator: requestDecorator,
		coroutineNumber:  m.CoroutineNumber,
		segmentSize:      m.SegmentSize,
		savePath:         savePath,
		httpClient:       m.HttpClient,
	}
	m.taskMap[id] = task
	err = task.start()
	return
}

// 暂停所有
func (*Manager) PauseAll() error {
	return nil
}

// 开始所有
func (*Manager) StartAll() error {
	return nil
}

// 暂停任务
func (*Manager) PauseTask(id TaskId) error {

	return nil
}

// 开始任务
func (*Manager) StartTask(id TaskId) error {
	return nil
}

// 取消任务
func (*Manager) CancelTask(id TaskId) error {
	return nil
}

// 获取一个缓存
func (m *Manager) getBuffer() *bytes.Buffer {
	return m.bufferPool.Get().(*bytes.Buffer)
}

// 释放一个缓存
func (m *Manager) releaseBuffer(b *bytes.Buffer) {
	b.Reset()
	m.bufferPool.Put(b)
}

type DownloadEvent struct {
	TaskId TaskId
	Event  string
	Data   interface{}
}

func (m *Manager) eventNotify(event *DownloadEvent) {
	// 取出未读消息
	select {
	case <-m.EventChan:
	default:
	}
	// 非阻塞写
	select {
	case m.EventChan <- event:
	default:
	}
}
