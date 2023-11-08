package cache

import (
	"testing"
	"time"
)

func TestCachOP(t *testing.T) {
	testData := []struct {
		key    string
		val    interface{}
		expire time.Duration
	}{
		{"fsfuhfu", 678, time.Second * 10},
		{"fsfedfaaswu", false, time.Second * 11},
		{"amifnn", true, time.Second * 12},
		{"kovmjai", map[string]interface{}{"a": 3, "b": false}, time.Second * 13},
		{"fsfuhfu", "aeaf", time.Second * 14},
	}
	c := NewMemCache()
	c.SetMaxMemory("10MB")
	for _, item := range testData {
		c.Set(item.key, item.val, item.expire)
		val, ok := c.Get(item.key)
		if !ok {
			t.Error("缓存取值失败")
		}
		if item.key != "kovmjai" && val != item.val {
			t.Error("缓存取值数据与预期不一致")
		}
		_, ok1 := val.(map[string]interface{})
		if item.key == "kovmjai" && !ok1 {
			t.Error("缓存取值数据与预期不一致")
		}
	}
	if int64(len(testData)) != c.Keys() {
		t.Error("缓存取值数据与预期不一致")
	}
	c.Del(testData[0].key)
	c.Del(testData[1].key)
	if int64(len(testData)) != c.Keys()+2 {
		t.Error("缓存取值数据与预期不一致")
	}

	time.Sleep(time.Second * 16)

	if c.Keys() != 0 {
		t.Error("过期缓存清空失败")
	}
}
