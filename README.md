# go_book



### 不理解：
1. 觉得所有抛出的文案，都应该在handler给出，其余地方应该使用同一的标识，不然国际化很难做。

2. 数据库的初始化应该全局 维护地址池， 为什么一定要先全部需要的调用都出初始化，好处是限制层级调用方式，坏处是很难用。

3. 涉及到业务的文案，内部应该调用全局标识，只在handler层面返回对应的文案，便于后续整理 用于国际化

4. session的写入和获取 为什么不能放service