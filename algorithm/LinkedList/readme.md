```cassandraql

    单向链表：
        特点: 方向是单向,顺序访问(默认从头部开始),使用指针构造的线性表 
    
    node结构体的设计
        type List struct {
        	headNode *SingleNode
        	count    int   //统计每次的链表个数
        }
    类似于redis中的简单动态字符串(sds)的设计
        struct sdshdr {
                int len; //buf中已占用的空间长度
                int  free;//buf中剩余可用的长度
            }    
    这样做的目的是获取链表的长度的时间复杂度是O(1)
    
    
```