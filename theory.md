### 数据生成理论
#### 伪随机数生成法
* https://www.jianshu.com/p/c009ad062964

* 平方取中法 
* 线性同余法

#### 生成高斯分布
* The Box–Muller transform 把一对均匀分布随机数映射到一对标准正态分布随机数
* https://www.cnblogs.com/mightycode/p/8370616.html 近似算法



#### 生成任意分布
* https://www.cnblogs.com/skiwnchiwns/p/10345428.html

* Inverse Transform Method 如果我们可以给出概率分布的累积分布函数（下文简称CDF）及其逆函数的解析表达式，则可以非常简单便捷的生成指定分布随机数。
* Acceptance-Rejection Method ARM的适用范围比ITM要大，只要给出概率密度函数（下文简称PDF）的解析表达式即可，而大多数常用分布的PDF是可以查到的。ARM本质上是一种模拟方法，而非直接数学方法。它每次生成新的随机数后，通过另一个随机数来保证其被接受概率服从指定的PDF。虽然ARM从效率上不如ITM，但是其适应性更广，在无法得到CDF的逆函数时，ARM是不错的选择。
* 当目标分布可以用其它分布经过四则运算表示时，可以使用组合算法生成对应随机数。

### 其他方法
* rejection sampling
* inverse cdf
* box-muller
* ziggurat
