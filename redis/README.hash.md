## hash
>- help @hash

## HSET key field value [field value ...]
>- summary: Set the string value of a hash field
>- since: 2.0.0
>- 设置hash 

## HGET key field
>- summary: Get the value of a hash field
>- since: 2.0.0
>- 获取hash 字段关联的值

## HGETALL key
>- summary: Get all the fields and values in a hash
>- since: 2.0.0
>- 获取hash所有的字段和值

## HLEN key
>- summary: Get the number of fields in a hash
>- since: 2.0.0
>- 获取hash 长度
  
## HSTRLEN key field
>- summary: Get the length of the value of a hash field
>- since: 3.2.0  
>- 获取某个字段长度

## HEXISTS key field
>- summary: Determine if a hash field exists
>- since: 2.0.0
>- 判断字段是否存在

## HKEYS key
>- summary: Get all the fields in a hash
>- since: 2.0.0
>- 获取所有的字段
 
## HVALS key
>- summary: Get all the values in a hash
>- since: 2.0.0
>- 获取所有的值
 
## HMGET key field [field ...]
>- summary: Get the values of all the given hash fields
>- since: 2.0.0
>- 获取多个字段值

## HMSET key field value [field value ...]
>- summary: Set multiple hash fields to multiple values
>- since: 2.0.0
>- 设置多个字段值  
  
## HSETNX key field value
>- summary: Set the value of a hash field, only if the field does not exist
>- since: 2.0.0  
>- 设置某个字段值，不存在则设置
 
## HDEL key field [field ...]
>- summary: Delete one or more hash fields
>- since: 2.0.0 
>- 删除1个或者多个字段

## HINCRBY key field increment
>- summary: Increment the integer value of a hash field by the given number
>- since: 2.0.0
>- 数值计算

## HINCRBYFLOAT key field increment
>- summary: Increment the float value of a hash field by the given amount
>- since: 2.6.0  
>- 小数计算
  
## HRANDFIELD key [count [WITHVALUES]]
>- summary: Get one or multiple random fields from a hash
>- since: 6.2.0
>- 随机获取指定数量的字段

## HSCAN key cursor [MATCH pattern] [COUNT count]
>- summary: Incrementally iterate hash fields and associated values
>- since: 2.8.0  
>- 
  
```
hset student name jianzhu age 18 gender 1
hget student name
hgetall student
hlen student
hstrlen student name 
hexists student age 
hkeys student
hvals student
hmget student name gender
hmset student age 25 scores 86
hsetnx student rank 1
hsetnx student age 10
hdel student rank scores 
hincrby student age 3
hget student age 
hincrbyfloat student score 1.3
hget student score

hmset product p001 10 p002 5 p003 20 p004 14
hrandfield product 2 
hrandfield product 1 withvalues

```
  





  

  

  

  

 

  

  

  

  