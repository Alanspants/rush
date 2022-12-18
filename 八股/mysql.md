### 1. 数据库三范式

1. 列的原子性，每一列都是不可分割的原子数据项。如下图contact中包含地址和电话，不符合要求。

   ![img](https://img-blog.csdnimg.cn/ade3e540541e4e3c81d1d89a95e74555.png)

2. 要求不存在部分依赖，即非主属性完全依赖于主属性。非主键列不能只依赖于主键的一部分。如下图kc_name依赖与kc_id就能查出来，不符合要求。

   ![img](https://img-blog.csdnimg.cn/6ca1e5249baa4512bc4cff970613b9e6.png)

3. 表中不能有循环依赖，非主属性直接依赖于主属性，不能是间接依赖。比如这个表中，sex_code依赖于id, 而sex_desc依赖于sex_code，这样sex_desc间接依赖于id，这样是不符合第三范式的。

   ![img](https://img-blog.csdnimg.cn/375ec0ad491a425d972bb2bad834be46.png)

### 2. 主键

表示唯一的列。例如学生表中学生ID是唯一的，学生ID就是主键。

### 3. 外键

外键用于连接两个表，是别的表的主键。成绩表中的学号，是学生表的主键，是成绩表的外键。

### 4. 索引

索引是为了加速对表中数据行的检索而创建的一种分散的存储结构。

索引是针对表而创立的，他是由数据页面以外的索引页面组成，每个索引页面中的行都会含有逻辑指针，以便加速物理数据。

索引主要是用来优化查询速率。

1. 创建索引方法

```mysql
CREATE INDEX index_name ON table_name (column_name);
```

2. 单列索引：基于单一字段创建

```mysql
CREATE INDEX index_name ON table_name (column_name);
```

3. 唯一索引：不仅用来提升查询性能，还能保证数据完整性。唯一索引不允许向表中插入任何重复值。

```mysql
CREATE UNIQUE INDEX index_name on table_name (column_name);
```

4. 聚簇索引：聚簇索引在表中两个或更多的列的基础上建立。创建单列还是聚簇索引，要看每次查询中，哪些列在作为过滤条件的where语句中最常出现。

```mysql
CREATE INDEX index_name on table_name (column_name_1, column_name_2)
```

