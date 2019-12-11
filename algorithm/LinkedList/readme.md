```cassandraql

    单向链表：
        特点: 方向是单向,顺序访问(默认从头部开始),使用指针构造的线性表 
    
        node结构体的设计
            type List struct {
                headNode *SingleNode
                count    int   //统计每次的链表个数
            }
        类似于redis中的简单动态字符串(sds)的设计
          struct __attribute__ ((__packed__)) sdshdr8 {
              uint8_t len; /* used */
              uint8_t alloc; /* excluding the header and null terminator */
              unsigned char flags; /* 3 lsb of type, 5 unused bits */
              char buf[];
          };  
        这样做的目的是获取链表的长度的时间复杂度是O(1)
        
    
    双向链表:
        由节点组成,每个数据节点中心有两个指针,分别指向后继和前置节点(上线和下线，类似传销)
        
        Head                  A                     B
            next    --->     next      --->       next
            pre     <---     pre       <---       pre
            
        Head                A                        B
            next    --->     next \                /next
            pre     <---     pre \  \           /    /pre
                                  \  \        /   /
                                   \  \  C   /  /
                                    \   next/ /
                                      \ pre  /
    https://images2017.cnblogs.com/blog/922236/201712/922236-20171220110004678-643077430.png
           
```

### 单链表操作
 >1.数组转链表  
 >2.链表成环   
 >3.链表环检测    