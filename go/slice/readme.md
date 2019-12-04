```cassandraql

    slice:切片
        特点:定义的时候无需指定长度
        定义的两种方式
            eg: s := []string{"a","b"}
            s2 := make([]string, 5, 5)
        
        切片指向的是一个底层数据 
        append会创建一个新的切片 
        slice是线程不安全的
        
        slice append后slice是变成了一个新的对象，但是其地址还是最初的那个slice的地址 
            
            func Addr() {
                s := []int{1, 2}
                for i := 0; i < 17; i++ {
                    fmt.Printf("addr %p,len %d ,cap %d\n", &s, len(s), cap(s))
                    s = append(s, i)
                }
            }
    

    slice分析过程:
        slice底层结构 
            //runtime/slice.go
            type slice struct {
            	array unsafe.Pointer //数组指针
            	len   int          //长度大小
            	cap   int         //容量大小
            }
            slice初始化过程:
                make([]type,len,cap)的过程中对调用makeslice(et *_type, len, cap int) 函数进行内存申请和初始化工作 
                   源码部分如下:
                        func makeslice(et *_type, len, cap int) unsafe.Pointer {
                            // math.MulUintptr(et.size, uintptr(cap))  (4*sys.PtrSize)=256
                            mem, overflow := math.MulUintptr(et.size, uintptr(cap))
                            if overflow || mem > maxAlloc || len < 0 || len > cap {
                                mem, overflow := math.MulUintptr(et.size, uintptr(len))
                                if overflow || mem > maxAlloc || len < 0 {
                                    panicmakeslicelen()
                                }
                                panicmakeslicecap()
                            }
                        
                            return mallocgc(mem, et, true)
                        }
                   注意mallocgc 函数的注释部分 (// Allocate an object of size bytes.
                                       // Small objects are allocated from the per-P cache's free lists.
                                       // Large objects (> 32 kB) are allocated straight from the heap.)
                  也就是说其实也是申请了内存(网上资料说make只是负责做初始化工作,new申请分配内存)
            slice扩容过程分为三种2种情况:
                
                growslice(et *_type, old slice, cap int) 
                a:
                   slice的长度小于1024，则还是按照2次方扩容
                b:
                   slice扩容在原来的cap大小的基础上按照1\4扩容   
                最后是slice赋值过程(memmove(p, old.array, lenmem)),在这个赋值过程中，slice的地址是不会发生变化，但是是一个新的slice
```