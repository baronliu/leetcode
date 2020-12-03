/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (41.70%)
 * Likes:    431
 * Dislikes: 0
 * Total Accepted:    36.2K
 * Total Submissions: 78.7K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 * 
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
 * 
 * 进阶:
 * 
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 * 
 * 示例:
 * 
 * LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
 * 
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得密钥 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得密钥 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 * 
 * 
 */

// @lc code=start
type ListNode struct {
    k, v int
    prev, next *ListNode
}

type LRUCache struct {
    len, capacity int
    head, tail *ListNode
    dict map[int]*ListNode
}


func Constructor(capacity int) LRUCache {
    //初始化容量
    c := LRUCache{capacity: capacity}
    //初始化字典
    c.dict = make(map[int]*ListNode, 0)
    return c
} 

func (this *LRUCache) moveToHead(n *ListNode) {
    //n.prev == nil的时候，当前节点就在首节点，无需变更
    if n.prev != nil {
        n.prev.next = n.next
        //当前节点如果是尾节点
        if n.next == nil {
            //替换尾节点
            this.tail = n.prev
        } else {
            //非尾节点时考虑下个节点的前置
            n.next.prev = n.prev
        }
        //建立当前头结点与替换节点的关联关系
        n.prev = nil
        n.next = this.head
        this.head.prev = n
        //替换头节点
        this.head = n
    }
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.dict[key]; !ok {
        //未能找到
        return -1
    } else {
        //可以找到：1,把节点挪至head; 2,返回这个节点的值
        this.moveToHead(node)
        return node.v
    }
}

func (this *LRUCache) Put(key int, value int)  {
    //判断当前是否有值
    curV := this.Get(key)

    //更新
    if curV != -1 {
        //只用替换下头结点的值即可
        this.head.v = value
        return
    }

    //新增:1.容量够->直接把节点添加到头部即可 2.容量不够->先添加再剔除
    n := &ListNode{k: key, v: value, next: this.head}
    if this.tail == nil {
        //尾为空的时候说明，这是添加的第一个元素，需要设置下尾
        this.tail = n
    } else {
        //尾不为空时，说明头也不为空，需要把当前的头的前置处理为新节点
        this.head.prev = n
    }
    //新增节点至头部
    this.head = n

    if this.len < this.capacity {
        this.len++
    } else {
        //删除尾节点
        delete(this.dict, this.tail.k)
        this.tail.prev.next = nil
        this.tail = this.tail.prev
    }

    this.dict[key] = n
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

