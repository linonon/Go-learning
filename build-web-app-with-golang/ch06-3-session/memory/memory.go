package memory

import (
	"container/list"
	"sync"
	"time"

	"github.com/astaxie/session"
)

var pder = &Provider{list: list.New()}

type SessionStore struct {
	sid          string                      // session ID
	timeAccessed time.Time                   // Last time the session was accessed
	value        map[interface{}]interface{} // session value
}

func (ss *SessionStore) Set(key, value interface{}) error {
	ss.value[key] = value
	pder.SessionUpdate(ss.sid)
	return nil
}

func (ss *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(ss.sid)
	if v, ok := ss.value[key]; ok {
		return v
	} else {
		return nil
	}
}

func (ss *SessionStore) Delete(key interface{}) error {
	delete(ss.value, key)
	pder.SessionUpdate(ss.sid)
	return nil
}

func (ss *SessionStore) SessionID() string {
	return ss.sid
}

type Provider struct {
	lock     sync.Mutex               // for lock
	sessions map[string]*list.Element // store in memory
	list     *list.List               // Gc
}

func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{})
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
}

func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

func init() {
	pder.sessions = make(map[string]*list.Element)
	session.Register("memory", pder)
}