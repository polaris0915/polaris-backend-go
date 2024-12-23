## User Table

```mysql
-- 用户表
CREATE TABLE IF NOT EXISTS user (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'id',
    identity VARCHAR(36) NOT NULL COMMENT '唯一ID',
    account VARCHAR(256) NOT NULL COMMENT '账号',
    password VARCHAR(512) NOT NULL COMMENT '密码',
		email VARCHAR(128) NOT NULL COMMENT "邮箱",
    name VARCHAR(256) NULL COMMENT '用户昵称',
    avatar BLOB NULL COMMENT '用户头像',
    profile VARCHAR(512) NULL COMMENT '用户简介',
    role VARCHAR(10) DEFAULT 'user' NOT NULL COMMENT '用户角色：user/admin/ban',
    created_at DATETIME NOT NULL COMMENT '创建时间',
    updated_at DATETIME NOT NULL COMMENT '更新时间',
    deleted_at DATETIME NULL COMMENT '删除时间',
    UNIQUE KEY unique_identity (identity), -- 确保 identity 唯一
		UNIQUE KEY unique_email (email), -- 确保 email 唯一
    INDEX idx_identity (identity)
) COMMENT = '用户' COLLATE = utf8mb4_unicode_ci;
```

