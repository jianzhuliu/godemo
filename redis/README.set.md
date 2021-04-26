## set 
>- help @set
>- 数据不重复

## SADD key member [member ...]
>- summary: Add one or more members to a set
>- since: 1.0.0
>- 添加元素到集合中

## SMEMBERS key
>- summary: Get all the members in a set
>- since: 1.0.0
>- 查看集合成员
  
## SISMEMBER key member
>- summary: Determine if a given value is a member of a set
>- since: 1.0.0
>- 判断元素是否存在于集合中

## SMISMEMBER key member [member ...]
>- summary: Returns the membership associated with the given elements for a set
>- since: 6.2.0
>- 判断多个元素，其中是否存在于集合中
  
## SRANDMEMBER key [count]
>- summary: Get one or multiple random members from a set
>- since: 1.0.0 
>- 随机获取指定数量的元素
>- count >0 获取数据的个数未 min(count, 元素个数)
>- count <0 获取固定数量元素，大小固定为 |count|,会存在重复数据

## SREM key member [member ...]
>- summary: Remove one or more members from a set
>- since: 1.0.0
>- 移除元素，支持多个
  
## SPOP key [count]
>- summary: Remove and return one or multiple random members from a set
>- since: 1.0.0  
>- 随机弹出一个元素并移除
  
## SCARD key
>- summary: Get the number of members in a set
>- since: 1.0.0  
>- 获取集合总个数

```
sadd myset a b c 
smembers myset 
sismember myset a 
sismember myset d 
sadd myset d e f g 
srandmember myset 2
srandmember myset 50
srandmember myset -5 
srandmember myset -20
srem myset a b 
spop myset 

```

## SMOVE source destination member
>- summary: Move a member from one set to another
>- since: 1.0.0
>- 从一个集合中转移元素到另一个集合中

## SDIFF key [key ...]
>- summary: Subtract multiple sets
>- since: 1.0.0
>- 差集

## SDIFFSTORE destination key [key ...]
>- summary: Subtract multiple sets and store the resulting set in a key
>- since: 1.0.0
>- 差集结果保存到指定集合中

## SINTER key [key ...]
>- summary: Intersect multiple sets
>- since: 1.0.0
>- 交集

## SINTERSTORE destination key [key ...]
>- summary: Intersect multiple sets and store the resulting set in a key
>- since: 1.0.0
>- 交集结果存放到指定集合中
 
## SUNION key [key ...]
>- summary: Add multiple sets
>- since: 1.0.0
>- 并集

## SUNIONSTORE destination key [key ...]
>- summary: Add multiple sets and store the resulting set in a key
>- since: 1.0.0
>- 并集结果保存到指定集合中

```
sadd myset1 a b c d 
sadd myset2 1 2 3 4
smove myset1 myset2 b 
smembers myset1 
smembers myset2

sadd myset2 a
sadd myset1 1 3
sdiff myset1 myset2 
sdiff myset2 myset1

sdiffstore diffset myset1 myset2 
smembers diffset

sinter myset1 myset2 
sinterstore interset myset1 myset2 
smembers interset

sunion myset1 myset2 
sunionstore unionset myset1 myset2 
smembers unionset

```

## SSCAN key cursor [MATCH pattern] [COUNT count]
>- summary: Incrementally iterate Set elements
>- since: 2.8.0 


  

  

  

  
  

  

  



