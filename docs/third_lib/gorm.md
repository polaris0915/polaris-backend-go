## Install

1. 安装gorm

```shell
go get -u gorm.io/gorm
```

2. 安装mysql驱动

```shell
go get -u gorm.io/driver/mysql
```

## 关于gorm的tag

以下是 GORM 支持的所有常用 `tag` 及其作用的详细列表，按类别分类整理：

------

### **1. 基础字段映射**

| Tag 名称         | 功能                                                         |
| ---------------- | ------------------------------------------------------------ |
| `column`         | 指定字段在数据库中的列名。                                   |
| `type`           | 指定数据库中字段的数据类型。                                 |
| `size`           | 设置字段的大小（长度），例如 `size:255` 会生成 `VARCHAR(255)`。 |
| `primaryKey`     | 将字段标记为主键。                                           |
| `autoIncrement`  | 将字段设置为自增列。                                         |
| `unique`         | 将字段值设置为唯一。                                         |
| `default`        | 指定字段的默认值。                                           |
| `not null`       | 指定字段为非空。                                             |
| `index`          | 为字段创建普通索引（可带索引名）。                           |
| `uniqueIndex`    | 为字段创建唯一索引（可带索引名）。                           |
| `comment`        | 为字段添加注释。                                             |
| `check`          | 设置字段的检查约束，例如 `check:age > 0`。                   |
| `embedded`       | 将结构体嵌入到当前表中。                                     |
| `embeddedPrefix` | 嵌套字段的前缀，例如 `embeddedPrefix:user_` 会将 `Address` 的 `City` 字段映射为 `user_city`。 |
| `<-`             | 限制字段的写权限，例如 `gorm:"<-:create"` 表示字段只能在创建时写入。 |
| `->`             | 限制字段的读权限，例如 `gorm:"->:false"` 表示字段无法读取。  |
| `serializer:...` | 指定字段的序列化方式，例如 `serializer:json` 将字段值序列化为 JSON 存储。 |

------

### **2. 时间字段**

| Tag 名称         | 功能                                                         |
| ---------------- | ------------------------------------------------------------ |
| `autoCreateTime` | 创建记录时自动填充字段为当前时间，支持 `int`, `int64`, `time.Time`。 |
| `autoUpdateTime` | 更新记录时自动填充字段为当前时间，支持 `int`, `int64`, `time.Time`。 |
| `precision`      | 指定时间字段的精度，例如 `precision:3` 对应 `DATETIME(3)`。  |

------

### **3. 软删除**

| Tag 名称     | 功能                                                         |
| ------------ | ------------------------------------------------------------ |
| `softDelete` | 将字段标记为软删除字段，支持 `time.Time` 或 GORM 提供的 `gorm.DeletedAt` 类型。 |

------

### **4. 外键和关联**

| Tag 名称           | 功能                                                         |
| ------------------ | ------------------------------------------------------------ |
| `foreignKey`       | 指定外键字段名称，例如 `gorm:"foreignKey:UserID"`。          |
| `references`       | 指定外键引用的字段名称，例如 `gorm:"references:ID"`。        |
| `constraint`       | 添加外键约束，例如 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`。 |
| `many2many`        | 指定多对多关联表的表名，例如 `gorm:"many2many:user_roles"`。 |
| `joinForeignKey`   | 多对多关联中的当前表外键，例如 `gorm:"joinForeignKey:UserID"`。 |
| `joinReferences`   | 多对多关联中的关联表外键，例如 `gorm:"joinReferences:RoleID"`。 |
| `polymorphic`      | 指定多态关联的类型，例如 `gorm:"polymorphic:Owner"`。        |
| `polymorphicValue` | 指定多态关联的值，例如 `gorm:"polymorphicValue:User"`。      |

------

### **5. 索引相关**

| Tag 名称      | 功能                                                         |
| ------------- | ------------------------------------------------------------ |
| `index`       | 创建普通索引，例如 `gorm:"index"`。                          |
| `uniqueIndex` | 创建唯一索引，例如 `gorm:"uniqueIndex"`。                    |
| `priority`    | 设置索引中字段的优先级，例如 `gorm:"uniqueIndex:idx_name,priority:1"`。 |

------

### **6. 忽略和高级控制**

| Tag 名称         | 功能                                                         |
| ---------------- | ------------------------------------------------------------ |
| `-`              | 忽略字段，不将字段映射到数据库中。                           |
| `serializer:...` | 指定字段的自定义序列化器，例如 `serializer:json` 或自定义的序列化器。 |

------

### **7. 自定义序列化器**

GORM 支持通过 `serializer` 为字段自定义序列化和反序列化逻辑，例如：

- `serializer:json`：将字段序列化为 JSON。
- `serializer:gob`：将字段序列化为 GOB。

