## list
>- help @list 
>- 插入元素有顺序，内部元素没有顺序

## LPUSH key element [element ...]
>- summary: Prepend one or multiple elements to a list
>- since: 1.0.0
>- 一个一个地插入列表中

## LRANGE key start stop
>- summary: Get a range of elements from a list
>- since: 1.0.0
>- 从左到右读取指定范围内数据


## LLEN key
>- summary: Get the length of a list
>- since: 1.0.0
>- 获取长度
  
## LINDEX key index
>- summary: Get an element from a list by its index
>- since: 1.0.0  
>- 根据下标获取列表元素

## LPOP key [count]
>- summary: Remove and get the first elements in a list
>- since: 1.0.0
>- 弹出第一个元素
  

## LPOS key element [RANK rank] [COUNT num-matches] [MAXLEN len]
>- summary: Return the index of matching elements on a list
>- since: 6.0.6 
>- 获取下标
 
## LPUSHX key element [element ...]
>- summary: Prepend an element to a list, only if the list exists
>- since: 2.2.0
>- 列表存在时头部追加元素
 
## LINSERT key BEFORE|AFTER pivot element
>- summary: Insert an element before or after another element in a list
>- since: 2.2.0 
>- 在指定元素前或者后插入元素
 
```
lpush city bj gz sh
lrange city 0 -1
llen city 
lindex city 1
lpop city 
lpop city 
lpop city 
lpush city bj gz sh 
lpush city gz sz 
lpos city gz 
lpos city gz count 2
lpushx city bj 
lpushx a as 
lrange city 0 -1
linsert city before bj first 
linsert city after gz newgz 

```
 
## LREM key count element
>- summary: Remove elements from a list
>- since: 1.0.0
>- 删除指定数量的元素
>- count >0 从左到右匹配删除count个数
>- count =0 移除所有匹配的值
>- count <0 从右到左删除匹配count个数

```
lpush mylist a b c d a d c b 
lrange mylist 0 -1 
lrem mylist 2 c 
lrange mylist 0 -1
lrem mylist -1 a 
lrange mylist 0 -1
lrem mylist 0 d 
lrange mylist 0 -1

```

## LSET key index element
>- summary: Set the value of an element in a list by its index
>- since: 1.0.0
>- 替换指定位置元素
```
lpush mylist2 1 2 3 
lset mylist2 1 a 
lrange mylist2 0 -1 
lset mylist2 2 b
lrange mylist2 0 -1
```

## LTRIM key start stop
>- summary: Trim a list to the specified range
>- since: 1.0.0 
>- 修剪列表，保留 start - stop 下标数据
```
lpush mylist3 one 
lpush mylist3 two 
lpush mylist3 three 
lpush mylist3 four 
lpush mylist3 five 
lrange mylist3 0 -1
ltrim mylist3 1 4
lrange mylist3 0 -1
```
 
 
## RPUSH key element [element ...]
>- summary: Append one or multiple elements to a list
>- since: 1.0.0
>- 从尾部追加元素

## RPUSHX key element [element ...]
>- summary: Append an element to a list, only if the list exists
>- since: 2.2.0 
>- 在已经存在的列表中，从尾部追加元素
 
## RPOP key [count]
>- summary: Remove and get the last elements in a list
>- since: 1.0.0
>- 从右边弹出一个元素

## RPOPLPUSH source destination
>- summary: Remove the last element in a list, prepend it to another list and return it
>- since: 1.2.0
  
## LMOVE source destination LEFT|RIGHT LEFT|RIGHT
>- summary: Pop an element from a list, push it to another list and return it
>- since: 6.2.0
>- 列表之间移动元素
```
rpush list1 x y z 
rpush list2 a b c 
lrange list1 0 -1 
lrange list2 0 -1

rpoplpush list1 list2 
lrange list1 0 -1 
lrange list2 0 -1

lmove list2 list1 left right
lrange list1 0 -1 
lrange list2 0 -1

```
  

## BLPOP key [key ...] timeout
>- summary: Remove and get the first element in a list, or block until one is available
>- since: 2.0.0

## BRPOP key [key ...] timeout
>- summary: Remove and get the last element in a list, or block until one is available
>- since: 2.0.0

## BRPOPLPUSH source destination timeout
>- summary: Pop an element from a list, push it to another list and return it; or block until one is available
>- since: 2.2.0

## BLMOVE source destination LEFT|RIGHT LEFT|RIGHT timeout
>- summary: Pop an element from a list, push it to another list and return it; or block until one is available
>- since: 6.2.0


  

  

