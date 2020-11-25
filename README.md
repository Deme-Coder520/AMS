# AMS 
### 基于beego框架的文章发布管理系统
**1、环境搭建：**

链接：https://blog.csdn.net/weixin_45634856/article/details/109490284

**2、主要用到的技术：**

golang(beego框架)、redis、mysql、jQuery、html、css

**3.部分主要知识点：**

QuerySeter：实现对表查询结果的一些操作，主要方法如下：

1）Count() (int64, error):统计表的行数

2）All(container interface{}, cols ...string) (int64, error)：获取所有数据

3）Limit(limit interface{}, args ...interface{}) QuerySeter：限制每一页显示的数量，给定初始显示位置

4）Filter(string, ...interface{}) QuerySeter：过滤器，相当于sql中的where，多表使用--filter(表名__字段,值)











### `regard--工作快乐，生活开心！`

