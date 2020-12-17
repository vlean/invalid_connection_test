
# 文档整理
- https://www.cnblogs.com/ivictor/p/5979731.html

interactive_timeout针对交互式连接，wait_timeout针对非交互式连接

- https://blog.csdn.net/idwtwt/article/details/91345364
- https://blog.csdn.net/dghpgyss/article/details/86480837

mysql防止空闲连接过多，超过了mysql_connection配置之后，会拒绝新连接，
自动关闭空闲连接超过wait_timeout、使用中连接超过interactive_timeout参数的连接

database/sql库连接池的max_lifetime大于wait_timeout，有可能拿到失效连接，
建议设置wait_timeout/2


- https://github.com/go-sql-driver/mysql/pull/934

go-sql-driver/mysql 1.5版本增加失效连接 读写前检查



# 测试方案

## mysql配置

mysql 5.7
max_connections = 5
wait_timeout = 5


## 测试方案

仅修改go-sql-driver/mysql版本和SetConnMaxLifetime值

起5个go协程执行mysql查询，之后sleep 10s，循环执行该操作，观察报错信息


 | go-sql-driver/mysql  | SetConnMaxLifetime | 结果|
 | -- | -- | -- |
 | 1.4.1 | 20s  | 稳定出现invalid connection |
 | 1.4.1| 2.5s | 未出现invalid connection; maxLifetimeClosed持续增加 |
 | 1.5.0 | 20s | 出现报错closing bad idle connection: EOF，但是可以正常查询 |

- 并发原因

- 触发wait_timeout