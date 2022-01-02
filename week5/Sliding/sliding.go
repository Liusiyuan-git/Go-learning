package Sliding

import (
	"sync"
	"time"
)

type Sliding struct {
	buckets        []*bucket     // 用于存放采集实例
	last           time.Time     // current 上一次采集的时间
	current        int           // 当前位置
	size           int           // 滑动窗口大小
	bucketDuration time.Duration // 滑动窗口滑动速率，默认是1s移动一次
	mu             sync.RWMutex  // 使用锁避免 data race
}

// bucket 用来存放采集数据
type bucket struct {
	val float64
}

func (b *bucket) Reset() {
	b.val = 0
}

type SlidingOption func(opt *Sliding)

//设置滑动窗口移动速率
func WithBucketDuration(t time.Duration) SlidingOption {
	return func(opt *Sliding) {
		opt.bucketDuration = t
	}
}

func NewSliding(opts ...SlidingOption) *Sliding {
	//初始化一个滑动窗口，默认大小是10
	sliding := &Sliding{
		size:           10,
		bucketDuration: time.Second,
		last:           time.Now(),
	}

	//采集options，自定义想要采集的指标
	for _, opt := range opts {
		opt(sliding)
	}

	//定义存放采集实例的slice，大小是sliding.size
	sliding.buckets = make([]*bucket, sliding.size)
	for i := range sliding.buckets {
		//初始化每个采集元素实例
		sliding.buckets[i] = &bucket{}
	}
	return sliding
}

// currentBucket 获取当前的 currentBucket
func (r *Sliding) currentBucket() *bucket {
	old := r.current
	// 计算滑动窗口移动了多少格
	s := int(time.Since(r.last) / r.bucketDuration)
	if s > 0 {
		//s > 0, 表明上次采集还是在上次，需要把last重置一下。
		r.last = time.Now()
	}

	// 计算新的 current
	r.current = (old + s) % r.size

	// s 大于 size，表明已经过了很久都没有采集了，需要关注的采集数据仅限于窗口的大小，所以需要重置一下
	if s > r.size {
		s = r.size
	}

	// 窗口每走一步，没有采集到数据的话，那默认采集到的是0，所以需要对之前的某些元素做清空
	for i := 1; i <= s; i++ {
		r.buckets[(old+i)%r.size].Reset()
	}
	return r.buckets[r.current]
}

// Add 添加一个采集数值
func (r *Sliding) Add(val float64) {
	if val == 0 {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.currentBucket().val += val
}

// Sum 数据统计
func (r *Sliding) Sum() float64 {
	var sum float64

	r.mu.RLock()
	defer r.mu.RUnlock()
	old := r.current
	s := int(time.Since(r.last) / r.bucketDuration)
	// 计算新的 current
	n := (old + s) % r.size

	// 求和
	for i := 0; i < r.size-s; i++ {
		sum += r.buckets[(n+i+1)%r.size].val
	}
	return sum
}
