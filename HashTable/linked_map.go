package HashTable

//哈希表简单实现，拉链法解决哈希冲突
const (
	expandFactor = 0.75
)

//哈希表的数据结构
type HashMap struct {
	m   []*KeyPairs //hash表的桶
	cap int
	len int
}
type KeyPairs struct { //键值对
	key   string
	value interface{}
	next  *KeyPairs
}

//初始化一个hashmap
func NewHashMap(cap int) *HashMap {
	if cap < 16 {
		cap = 16
	}
	hashMap := new(HashMap)
	hashMap.m = make([]*KeyPairs, cap, cap)
	return hashMap
}

func (h *HashMap) Index(key string) int {
	return BKDRHash(key, h.cap) //计算键的哈希值
}

//写入一个键值对
func (h *HashMap) Put(key string, value interface{}) {
	index := h.Index(key)
	element := h.m[index]
	if element == nil { //该下标没有值，直接写入
		h.m[index] = &KeyPairs{key, value, nil}
	} else { //有值，找到最后一个节点
		for element.next != nil {
			if element.key == key { //如果是相同的键就进行修改操作
				element.value = value
				return
			}
			element = element.next
		}
		element.next = &KeyPairs{key, value, nil}
	}
	h.len++
	if (float64(h.len) / float64(h.cap)) > expandFactor { //需要扩容，把切片容量变为2倍
		newH := NewHashMap(2 * h.cap)
		for _, pairs := range h.m {
			for pairs != nil {
				newH.Put(pairs.key, pairs.value)
			}
		}
		h.cap = newH.cap
		h.m = newH.m
	}
}

//获取值
func (h *HashMap) Get(key string) interface{} {
	index := h.Index(key)
	pairs := h.m[index]
	if pairs.key == key {
		return pairs.value
	}
	for pairs.next != nil {
		if pairs.key == key {
			return pairs.value
		}
		pairs = pairs.next
	}
	return nil
}

//哈希函数
func BKDRHash(str string, cap int) int {
	seed := int(131) // 31 131 1313 13131 131313 etc..
	hash := int(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + int(str[i])
	}
	return hash % cap
}

//测试
//func main() {
//	h := NewHashMap(32)
//	h.Put("aa","bb")
//	res := h.Get("aa")
//	fmt.Println(res)
//}
