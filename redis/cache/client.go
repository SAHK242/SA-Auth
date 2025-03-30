package cache

//
//import (
//	"context"
//	"fmt"
//	"reflect"
//	"sync"
//	"time"
//
//	"github.com/eko/gocache/lib/v4/codec"
//	"github.com/goccy/go-json"
//	"github.com/redis/go-redis/v9"
//
//)
//
//type cacheEntry struct {
//	ctx context.Context
//	key string
//	val []byte
//	ttl time.Duration
//}
//
//type MyCache[T any] interface {
//	Get(ctx context.Context, key string, returnObj T) (T, error)
//	GetWithTTL(ctx context.Context, key string, returnObj T) (T, int64, error)
//	Set(ctx context.Context, key string, value T)
//	SetWithTTL(ctx context.Context, key string, value T, ttl time.Duration)
//	SetSync(ctx context.Context, key string, value T) error
//	SetWithTTLSync(ctx context.Context, key string, value T, ttl time.Duration) error
//	Delete(ctx context.Context, key string) error
//	GetStats() *codec.Stats
//	Close()
//}
//
//type myCache[T any] struct {
//	client        *redis.Client
//	loadFunc      func(ctx context.Context, key string) (T, error)
//	keyEncodeFunc func(key string) string
//	keyDecodeFunc func(key string) string
//	errFunc       func(err error)
//	ttl           time.Duration
//	setChannel    chan *cacheEntry
//	setWg         *sync.WaitGroup
//	errChannel    chan Error
//	errWg         *sync.WaitGroup
//	statsChannel  chan *cacheStatsItem
//	statsWg       *sync.WaitGroup
//	statsEnabled  bool
//	stats         *Stats
//}
//
//func NewCache[T any](redis *redis.Client, opts ...Option[T]) MyCache[T] {
//	o := eval(opts)
//
//	loadableCache := &myCache[T]{
//		client:        redis,
//		loadFunc:      o.loadFunc,
//		keyEncodeFunc: o.keyEncodeFunc,
//		keyDecodeFunc: o.keyDecodeFunc,
//		ttl:           o.ttl,
//		setChannel:    make(chan *cacheEntry, 100),
//		setWg:         &sync.WaitGroup{},
//		errChannel:    make(chan Error, 100),
//		errWg:         &sync.WaitGroup{},
//	}
//
//	if o.statsEnabled {
//		loadableCache.statsChannel = make(chan *cacheStatsItem, 100)
//		loadableCache.statsWg = &sync.WaitGroup{}
//		loadableCache.statsEnabled = true
//		loadableCache.stats = newStats()
//	}
//
//	loadableCache.setWg.Add(1)
//	go loadableCache.startCachePutBackHandler()
//
//	loadableCache.errWg.Add(1)
//	go loadableCache.startCacheErrorHandler()
//
//	if loadableCache.statsEnabled {
//		loadableCache.statsWg.Add(1)
//		go loadableCache.startCacheStatsHandler()
//	}
//
//	return loadableCache
//}
//
//func (m *myCache[T]) Get(ctx context.Context, key string, returnObj T) (T, error) {
//	var err error
//	var raw string
//	var object T
//
//	key = m.keyEncodeFunc(key)
//
//	raw, err = m.client.Get(ctx, key).Result()
//
//	if err == nil {
//		if !reflectutil.IsPrimitive(returnObj) {
//			if e := m.unmarshall(raw, &returnObj); e != nil {
//				m.errChannel <- NewGetError(e)
//				return object, e
//			}
//
//			m.errChannel <- NewGetError(nil)
//
//			return returnObj, err
//		} else {
//			o, _ := reflectutil.FromString(raw, reflect.TypeOf(returnObj))
//
//			m.errChannel <- NewGetError(nil)
//
//			return o.(T), err
//		}
//
//	} else {
//		m.errChannel <- NewGetError(err)
//
//		if m.loadFunc != nil {
//			object, err = m.loadFunc(ctx, m.keyDecodeFunc(key))
//		}
//
//		if err != nil {
//			return object, err
//		} else {
//			go m.Set(ctx, m.keyDecodeFunc(key), object)
//			return object, err
//		}
//	}
//}
//
//func (m *myCache[T]) GetWithTTL(ctx context.Context, key string, returnObj T) (T, int64, error) {
//	var err error
//	var raw string
//	var object T
//	var ttl time.Duration
//
//	key = m.keyEncodeFunc(key)
//
//	raw, err = m.client.Get(ctx, key).Result()
//
//	if err == nil {
//		if !reflectutil.IsPrimitive(returnObj) {
//			if e := m.unmarshall(raw, &returnObj); e != nil {
//				m.errChannel <- NewGetError(e)
//				return object, 0, e
//			}
//
//			ttl, err = m.client.TTL(ctx, key).Result()
//			if err != nil {
//				m.errChannel <- NewGetError(err)
//				return returnObj, 0, err
//			}
//
//			m.errChannel <- NewGetError(nil)
//			return returnObj, int64(ttl.Seconds()), err
//		} else {
//			ttl, err = m.client.TTL(ctx, key).Result()
//			if err != nil {
//				m.errChannel <- NewGetError(err)
//				return returnObj, 0, err
//			}
//
//			o, _ := reflectutil.FromString(raw, reflect.TypeOf(returnObj))
//
//			m.errChannel <- NewGetError(nil)
//
//			return o.(T), int64(ttl.Seconds()), err
//		}
//	} else {
//		m.errChannel <- NewGetError(err)
//
//		if m.loadFunc != nil {
//			object, err = m.loadFunc(ctx, m.keyDecodeFunc(key))
//		}
//
//		if err != nil {
//			return object, 0, err
//		} else {
//			// Put the object back in the cache
//			go m.Set(ctx, m.keyDecodeFunc(key), object)
//			return object, int64(m.ttl.Seconds()), err
//		}
//	}
//}
//
//func (m *myCache[T]) Set(ctx context.Context, key string, value T) {
//	m.SetWithTTL(ctx, key, value, m.ttl)
//}
//
//func (m *myCache[T]) SetWithTTL(ctx context.Context, key string, value T, ttl time.Duration) {
//	t := reflect.TypeOf(value)
//
//	kind := t.Kind()
//
//	if kind == reflect.Ptr && reflect.ValueOf(value).IsNil() {
//		m.errChannel <- NewSetError(fmt.Errorf("value cannot be nil"))
//	} else if reflect.ValueOf(value).IsZero() {
//		m.errChannel <- NewSetError(fmt.Errorf("value cannot be zero"))
//	}
//
//	key = m.keyEncodeFunc(key)
//
//	if kind != reflect.Ptr && kind != reflect.Struct && kind != reflect.Slice && kind != reflect.Map {
//		m.setChannel <- &cacheEntry{
//			ctx: ctx,
//			key: key,
//			val: []byte(fmt.Sprintf("%v", value)),
//			ttl: ttl,
//		}
//
//		return
//	}
//
//	bytes, err := m.marshall(value)
//
//	if err != nil {
//		m.errChannel <- NewSetError(err)
//	} else {
//		m.setChannel <- &cacheEntry{
//			ctx: ctx,
//			key: key,
//			val: bytes,
//			ttl: ttl,
//		}
//	}
//}
//
//func (m *myCache[T]) SetSync(ctx context.Context, key string, value T) error {
//	return m.SetWithTTLSync(ctx, key, value, m.ttl)
//}
//
//func (m *myCache[T]) SetWithTTLSync(ctx context.Context, key string, value T, ttl time.Duration) error {
//	t := reflect.TypeOf(value)
//
//	kind := t.Kind()
//
//	if kind == reflect.Ptr && reflect.ValueOf(value).IsNil() {
//		return fmt.Errorf("value cannot be nil")
//	} else if reflect.ValueOf(value).IsZero() {
//		return fmt.Errorf("value cannot be zero")
//	}
//
//	key = m.keyEncodeFunc(key)
//
//	if kind != reflect.Ptr && kind != reflect.Struct && kind != reflect.Slice && kind != reflect.Map {
//		item := &cacheEntry{
//			ctx: ctx,
//			key: key,
//			val: []byte(fmt.Sprintf("%v", value)),
//			ttl: ttl,
//		}
//
//		return m.client.Set(item.ctx, item.key, item.val, item.ttl).Err()
//	}
//
//	bytes, err := m.marshall(value)
//
//	if err != nil {
//		return err
//	} else {
//		item := &cacheEntry{
//			ctx: ctx,
//			key: key,
//			val: bytes,
//			ttl: ttl,
//		}
//
//		return m.client.Set(item.ctx, item.key, item.val, item.ttl).Err()
//	}
//}
//
//func (m *myCache[T]) Delete(ctx context.Context, key string) error {
//	key = m.keyEncodeFunc(key)
//	return m.client.Del(ctx, key).Err()
//}
//
//func (m *myCache[T]) GetStats() *codec.Stats {
//	if m.statsEnabled {
//		return m.stats.get()
//	}
//
//	return nil
//}
//
//func (m *myCache[T]) Close() {
//	close(m.setChannel)
//	close(m.errChannel)
//
//	m.setWg.Wait()
//	m.errWg.Wait()
//
//	if m.statsEnabled {
//		close(m.statsChannel)
//		m.statsWg.Wait()
//	}
//}
//
//func (m *myCache[T]) marshall(value T) ([]byte, error) {
//	bytes, err := json.Marshal(&value)
//
//	return bytes, err
//}
//
//func (m *myCache[T]) unmarshall(raw any, returnObj any) error {
//	var err error
//
//	switch v := raw.(type) {
//	case []byte:
//		err = json.Unmarshal(v, returnObj)
//	case string:
//		err = json.Unmarshal([]byte(v), returnObj)
//	default:
//		err = fmt.Errorf("cannot unmarshall type %T from cache", v)
//	}
//
//	return err
//}
//
//func (m *myCache[T]) startCachePutBackHandler() {
//	defer m.setWg.Done()
//
//	for item := range m.setChannel {
//		err := m.client.Set(item.ctx, item.key, item.val, item.ttl).Err()
//
//		m.errChannel <- NewSetError(err)
//	}
//}
//
//func (m *myCache[T]) startCacheErrorHandler() {
//	defer m.errWg.Done()
//
//	for item := range m.errChannel {
//		if item.Err != nil && m.errFunc != nil {
//			m.errFunc(item.Err)
//		}
//
//		if m.statsEnabled {
//			m.statsChannel <- &cacheStatsItem{
//				op:      item.Op,
//				success: item.Err == nil,
//			}
//		}
//	}
//}
//
//func (m *myCache[T]) startCacheStatsHandler() {
//	defer m.statsWg.Done()
//
//	for item := range m.statsChannel {
//		m.stats.updateStats(item.op, item.success)
//	}
//}
