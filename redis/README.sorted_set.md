## sorted_set
>- help @sorted_set
>- 有序不重复
>- score 分数
>- member 元素
>- rank 排名


## ZADD key [NX|XX] [GT|LT] [CH] [INCR] score member [score member ...]
>- summary: Add one or more members to a sorted set, or update its score if it already exists
>- since: 1.2.0
>- 有序集合中添加元素

## ZRANGE key min max [BYSCORE|BYLEX] [REV] [LIMIT offset count] [WITHSCORES]
>- summary: Return a range of members in a sorted set
>- since: 1.2.0
>- 获取有序集合指定范围的数据

## ZRANGESTORE dst src min max [BYSCORE|BYLEX] [REV] [LIMIT offset count]
>- summary: Store a range of members from sorted set into another key
>- since: 6.2.0
>- 存储有序集合指定范围的数据到指定集合中
 
## ZREVRANGE key start stop [WITHSCORES]
>- summary: Return a range of members in a sorted set, by index, with scores ordered from high to low
>- since: 1.2.0
>- 反序获取有序集合指定范围的数据
  
## ZCARD key
>- summary: Get the number of members in a sorted set
>- since: 1.2.0 
>- 获取有序集合的成员数量
 
## ZCOUNT key min max
>- summary: Count the members in a sorted set with scores within the given values
>- since: 2.0.0
>- 根据分数范围获取对应的成员
 
## ZINCRBY key increment member
>- summary: Increment the score of a member in a sorted set
>- since: 1.2.0
>- 增加成员分数
  
## ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
>- summary: Return a range of members in a sorted set, by score
>- since: 1.0.5
>- 根据分数排序，指定分数范围获取成员

## ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]
>- summary: Return a range of members in a sorted set, by score, with scores ordered from high to low
>- since: 2.2.0
>- 反序按分数排序范围搜索
  
## ZRANDMEMBER key [count [WITHSCORES]]
>- summary: Get one or multiple random elements from a sorted set
>- since: 6.2.0
>- 随机获取指定数量的成员
>- count >0 获取数据的个数未 min(count, 元素个数)
>- count <0 获取固定数量元素，大小固定为 |count|,会存在重复数据

## ZSCORE key member
>- summary: Get the score associated with the given member in a sorted set
>- since: 1.2.0
>- 获取成员的分数
  
## ZMSCORE key member [member ...]
>- summary: Get the score associated with the given members in a sorted set
>- since: 6.2.0 
>- 获取有序集合指定成员的分数

## ZPOPMAX key [count]
>- summary: Remove and return members with the highest scores in a sorted set
>- since: 5.0.0
>- 弹出分数最大的指定个数成员

## ZPOPMIN key [count]
>- summary: Remove and return members with the lowest scores in a sorted set
>- since: 5.0.0
>- 弹出分数最小的指定个数成员
  
## ZRANK key member
>- summary: Determine the index of a member in a sorted set
>- since: 2.0.0
>- 获取某个成员的排名，按分数

## ZREVRANK key member
>- summary: Determine the index of a member in a sorted set, with scores ordered from high to low
>- since: 2.0.0
>- 获取某个成员的反序排名，按分数
  
```
zadd myzset 1 one 2 two 3 three 4 four 5 five
zrange myzset 0 -1 withscores
zrange myzset 2 4 byscore limit 0 3 withscores
zrange myzset - + bylex limit 0 3 
zcard myzset 
zcount myzset 2 4
zincrby myzset 3 three 
zrange myzset 0 -1 withscores
zrangebyscore myzset (2 4 withscores
zrangebyscore myzset 2 (4 withscores
zrandmember myzset 3 withscores
zmscore myzset one three 
zpopmax myzset 
zpopmin myzset 
zrank myzset five 


```

## ZLEXCOUNT key min max
>- summary: Count the number of members in a sorted set between a given lexicographical range
>- since: 2.8.9  
>- 根据字典排序范围统计成员数
>- min max 成员名称需要加\[ 表示包含 (表示不包含, -和+表示最小值，最大值

## ZRANGEBYLEX key min max [LIMIT offset count]
>- summary: Return a range of members in a sorted set, by lexicographical range
>- since: 2.8.9
>- 按字典排序范围搜索
 
## ZREVRANGEBYLEX key max min [LIMIT offset count]
>- summary: Return a range of members in a sorted set, by lexicographical range, ordered from higher to lower strings.
>- since: 2.8.9 
>- 反序按字典排序范围搜索
  
```
zadd zs1 1 a 2 b 3 c 4 d 5 e 6 f
zlexcount zs1 - +
zlexcount zs1 [b [d
zlexcount zs1 (b [d
zlexcount zs1 (b (d
zrangebylex zs1 [b [d
zrangebylex zs1 (b [d

```

## BZPOPMAX key [key ...] timeout
>- summary: Remove and return the member with the highest score from one or more sorted sets, or block until one is available
>- since: 5.0.0
>- 弹出阻塞,可指定过期时间

## BZPOPMIN key [key ...] timeout
>- summary: Remove and return the member with the lowest score from one or more sorted sets, or block until one is available
>- since: 5.0.0
>- 弹出阻塞,可指定过期时间

## ZDIFF numkeys key [key ...] [WITHSCORES]
>- summary: Subtract multiple sorted sets
>- since: 6.2.0
>- 差集

## ZDIFFSTORE destination numkeys key [key ...]
>- summary: Subtract multiple sorted sets and store the resulting sorted set in a new key
>- since: 6.2.0
>- 差集结果保存到指定集合中

  
## ZINTER numkeys key [key ...] [WEIGHTS weight] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
>- summary: Intersect multiple sorted sets
>- since: 6.2.0
>- 交集

## ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight] [AGGREGATE SUM|MIN|MAX]
>- summary: Intersect multiple sorted sets and store the resulting sorted set in a new key
>- since: 2.0.0
>- 交集结果存放到指定集合中

## ZUNION numkeys key [key ...] [WEIGHTS weight] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
>- summary: Add multiple sorted sets
>- since: 6.2.0
>- 并集

## ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight] [AGGREGATE SUM|MIN|MAX]
>- summary: Add multiple sorted sets and store the resulting sorted set in a new key
>- since: 2.0.0  
>- 并集结果保存到指定集合中
  
```
zadd zset1 1 a 2 b 3 c 4 d 5 e 6 f
zadd zset2 10 c 11 b 12 g 13 h 
zdiff 2 zset1 zset2 withscores 

zinter 2 zset1 zset2 withscores
zinter 2 zset1 zset2 aggregate min withscores
zinter 2 zset1 zset2 aggregate max withscores

zinterstore storezinter 2 zset1 zset2
zrange storezinter 0 -1 withscores

zunion 2 zset1 zset2 withscores
zunion 2 zset1 zset2 aggregate min withscores
zunion 2 zset1 zset2 aggregate max withscores

zunion storeunion 2 zset1 zset2 aggregate max
zrange storeunion 0 -1 withscores
```

  

## ZREM key member [member ...]
>- summary: Remove one or more members from a sorted set
>- since: 1.2.0
>- 删除一个或者多个成员

## ZREMRANGEBYLEX key min max
>- summary: Remove all members in a sorted set between the given lexicographical range
>- since: 2.8.9
>- 根据字典排序，删除指定范围的成员

## ZREMRANGEBYRANK key start stop
>- summary: Remove all members in a sorted set within the given indexes
>- since: 2.0.0
>- 根据排名，按下标范围删除成员

## ZREMRANGEBYSCORE key min max
>- summary: Remove all members in a sorted set within the given scores
>- since: 1.2.0
>- 根据分数排名，按分数范围删除成员

```
zrem zset1 b 
zremrangebylex zset1 (c [d
zremrangebyscore zset1 3 5
```

## ZSCAN key cursor [MATCH pattern] [COUNT count]
>- summary: Incrementally iterate sorted sets elements and associated scores
>- since: 2.8.0

  

  

