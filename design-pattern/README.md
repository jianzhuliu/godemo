## 设计模式
- (快来，这里有23种设计模式的Go语言实现（一）)[https://bbs.huaweicloud.com/blogs/279505]
- (快来，这里有23种设计模式的Go语言实现（二）)[https://bbs.huaweicloud.com/blogs/280291]

## 创建型模式 (Creational Pattern)

#### 1、单例模式 (Singleton Pattern)
- 主要用于保证一个类仅有一个实例，并提供一个访问它的全局访问点
- go test -v ./msgpool/

#### 2、建造者模式 (Builder Pattern)
- 主要解决需要创建对象时需要传入多个参数，或者对初始化顺序有要求的场景
- go test -v ./msg/ 

#### 3、工厂模式 (Factory Method Pattern)
- 通过提供一个工厂对象或者工厂方法，为使用者隐藏了对象创建的细节
- go test -v ./event/

#### 4、抽象工厂模式 (Abstract Factory Pattern)
- 对工厂方法模式的优化，通过为工厂对象新增一个抽象层，让工厂对象遵循单一职责原则，也避免了霰弹式修改
- go test -v ./pipeline/
- ./plugin/

#### 5、原型模式 (Prototype Pattern)
- 让对象复制更加简单
- go test -v ./prototype/


## 结构型模式 (Structural Pattern)





## 行为型模式 (Behavioral Pattern)


## 全部测试
- go test ./...
- go test -v ./...

