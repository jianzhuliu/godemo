## string
>- help @string
>- 支持下标索引，从0开始，最后一位为 -1 

## SET key value [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|KEEPTTL] [NX|XX] [GET]
>- summary: Set the string value of a key
>- since: 1.0.0
>- 设置key-value,支持过期时间 ex 秒，px 毫秒, exat|pxat后接时间戳
>- 设置新值，并返回旧值，等同于 getset 
```
set str1 val1 
get str1 

set name jianzhu get 
set name bob xx get 
```

## GET key
>- summary: Get the value of a key
>- since: 1.0.0
>- 获取key值
```
set name jianzhu 
get name   
```

## MSET key value [key value ...]
>- summary: Set multiple keys to multiple values
>- since: 1.0.1
>- 设置多个key-value  M -> more
```
mset name jianzhu age 18
mget name age 
```

## MGET key [key ...]
>- summary: Get the values of all the given keys
>- since: 1.0.0
>- 获取多个key值
```
set name jianzhuu 
set age 18
mget name age 
```
  
## APPEND key value
>- summary: Append a value to a key
>- since: 2.0.0
>- 追加字符串
```
set name jianzhu 
append name " liu"
get name
```

## STRLEN key
>- summary: Get the length of the value stored in a key
>- since: 2.2.0
>- 获取value长度
```
set name jianzhu
get name
strlen name 
```

## INCR key
>- summary: Increment the integer value of a key by one
>- since: 1.0.0
>- 数值计算，加1
```
set age 18
incr age 
incr age 
get age 
```

## INCRBY key increment
>- summary: Increment the integer value of a key by the given amount
>- since: 1.0.0
>- 数值计算，增加指定数值
```
set age 15
incrby age 4
get age 
```

## INCRBYFLOAT key increment
>- summary: Increment the float value of a key by the given amount
>- since: 1.0.0
>- 小数计算，增加指定小数或者整数
```
set age 16
incrbyfloat age 0.3
incrbyfloat age 2
get age 
```
  
## DECR key
>- summary: Decrement the integer value of a key by one
>- since: 1.0.0
>- 数值减一
```
set age 19
decr age 
decr age 
get age 
```

## DECRBY key decrement
>- summary: Decrement the integer value of a key by the given number
>- since: 1.0.0
>- 数值计算，减去指定数值 
```
set age 25
decrby age 3
get age 
``` 

## MSETNX key value [key value ...]
>- summary: Set multiple keys to multiple values, only if none of the keys exist
>- since: 1.0.1
>- 只有设置的key都不存在时，才设置成功
>- msetnx mnxk1 mnxv1 n1 1 n2 2
>- msetnx mnxk1 mnxv1 
>- get mnxk1

## SETNX key value
>- summary: Set the value of a key, only if the key does not exist
>- since: 1.0.0
>- 只有对应的key不存在时，才设置成功
>- setnx n1 1
>- setnx nxk1 v1 

## GETSET key value
>- summary: Set the string value of a key and return its old value
>- since: 1.0.0
>- 设置新值，返回旧值
>- getset name jack
>- get name 
>- getset name bob 
>- get name 
  
## GETDEL key
>- summary: Get the value of a key and delete the key
>- since: 6.2.0
>- 获取value并删除key  
>- getdel nxk1 
>- get nxk1
 
## SETEX key seconds value
>- summary: Set the value and expiration of a key
>- since: 2.0.0  
>- 设置过期时间,单位秒
```
getdel age 
setex age 5 18
ttl age 
get age
```

## GETEX key [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|PERSIST]
>- summary: Get the value of a key and optionally set its expiration
>- since: 6.2.0 
>- 获取value值，并设置获取时间
```
getdel age 
setex age 30 23
ttl age 
get age 
getex age ex 50
ttl age
get age  
```
  
## PSETEX key milliseconds value
>- summary: Set the value and expiration in milliseconds of a key
>- since: 2.6.0
>- 设置一个值，过期时间为毫秒
```
getdel age 
psetex age 10000 25
ttl age 
get age 
```

## GETRANGE key start end
>- summary: Get a substring of the string stored at a key
>- since: 2.4.0
>- 获取子字符串，下标从0开始
```
set name jianzhu
getrange name 0 -1
getrange name 1 5

```

## SETRANGE key offset value
>- summary: Overwrite part of a string at key starting at the specified offset
>- since: 2.2.0
>- 从指定下标位置开始替换,可以替换多个字符
```
set name jianzhu 
get name 
setrange name 4 g
get name 
setrange name 2 kkk
get name 
```


## SETBIT key offset value
>- summary: Sets or clears the bit at offset in the string value stored at key
>- since: 2.2.0
>- 位操作，设置第几位对应的值，取 0或者1，从左到右，下标从0开始
```
setbit user:12 1 1
getbit user:12 0
getbit user:12 1
getbit user:12 2
bitcout user:12
```

## GETBIT key offset
>- summary: Returns the bit value at offset in the string value stored at key
>- since: 2.2.0
>- 位操作，获取第几位对应的值
```
setbit user:13 5 1
getbit user:13 5 
```
  
## BITCOUNT key [start end]
>- summary: Count set bits in a string
>- since: 2.6.0
>- 位操作，位计数
```
setbit user:3 3 1
setbit user:3 9 1
bitcount user:3
```

## BITOP operation destkey key [key ...]
>- summary: Perform bitwise operations between strings
>- since: 2.6.0
>- 位操作， 并保存到目标key中
>- operation操作类型，支持 
>- and 逻辑并
>- or 逻辑或
>- xor 逻辑异或
>- not 逻辑非
>- 除了not ,其它都可以接受一个或者多个key
```
set myk1 "\x04"    	--- 00000100
set myk2 "\x06"		--- 00000110
set myk3 "\x09"		--- 00001001

bitop and kand myk1 myk2   	--- 00000100
bitop or kor myk1 myk2 		--- 00000110
bitop xor kxor myk1 myk2 	--- 00000010
bitop not knot myk2 		--- 11111001


set k1 a --- 0110 0001
set k2 b --- 0110 0010
mget k1 k2 
bitop and kand k1 k2 
get kand 
bitop or kor k1 k2 
get kor  

```

## BITPOS key bit [start] [end]
>- summary: Find first bit set or clear in a string
>- since: 2.8.7
>- 返回字符串里面第一个被设置为1或者0的bit位, start|end 表示字节位置，1byte=8bit 
```
setbit mykey 3 1
setbit mykey 8 1
setbit mykey 9 1 
bitpos mykey 1 0
bitpos mykey 1 1
bitpos mykey 0 1

```

## BITFIELD key [GET type offset] [SET type offset value] [INCRBY type offset increment] [OVERFLOW WRAP|SAT|FAIL]
>- summary: Perform arbitrary bitfield integer operations on strings
>- since: 3.2.0


## STRALGO LCS algo-specific-argument [algo-specific-argument ...]
>- summary: Run algorithms (currently LCS) against strings
>- since: 6.0.0

