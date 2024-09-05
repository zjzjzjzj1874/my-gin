
## 企业账号人员分配
设计一个支持多级组织架构的用户管理系统需要考虑数据模型、权限管理和用户交互等多个方面。以下是一个详细的设计思路：

1. 数据模型设计

1.1 部门（Department）
   ID: 唯一标识符
   名称: 部门名称
   父部门ID: 指向上级部门（支持多级）
   描述: 可选的部门描述
   创建时间: 时间戳
   sql
```mysql
CREATE TABLE departments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    parent_id INT,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES departments(id)
);
```

1.2 用户（User）
ID: 唯一标识符
用户名: 唯一的用户名
邮箱: 唯一的邮箱
手机号: 可选
密码: 存储加密后的密码
部门ID: 关联的部门
创建时间: 时间戳
```sql
CREATE TABLE users (
   id INT PRIMARY KEY AUTO_INCREMENT,
   username VARCHAR(255) NOT NULL UNIQUE,
   email VARCHAR(255) NOT NULL UNIQUE,
   phone VARCHAR(15),
   password VARCHAR(255) NOT NULL,
   department_id INT,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   FOREIGN KEY (department_id) REFERENCES departments(id)
);
```
1.3 权限组（Permission Group）
ID: 唯一标识符
名称: 权限组名称
描述: 可选的权限组描述
创建时间: 时间戳
```sql
CREATE TABLE permission_groups (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

1.4 用户权限关系（User Permissions）
用户ID: 关联用户
权限组ID: 关联权限组
```sql
CREATE TABLE user_permissions (
    user_id INT,
    permission_group_id INT,
    PRIMARY KEY (user_id, permission_group_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (permission_group_id) REFERENCES permission_groups(id)
);
```

2. 权限管理设计
   多级部门管理: 通过 parent_id 字段实现多级组织架构，支持无限层级的部门划分。
   权限组管理: 权限组可以跨部门，允许灵活的权限分配。
   用户分配权限: 用户可以被分配到多个权限组，方便管理。
3. 用户交互设计
   创建部门: 提供一个界面，允许管理员创建新部门，选择上级部门。
   创建用户: 提供一个界面，允许管理员输入用户信息，并选择其部门和权限组。
   权限管理: 提供界面管理权限组，允许将用户添加到权限组。
4. 示例流程
   创建部门: 管理员在系统中创建部门，选择上级部门，形成层级结构。
   创建权限组: 管理员创建权限组，定义各个权限。
   创建用户: 管理员创建用户，分配相应的部门和权限组。
   权限验证: 系统在用户登录后，根据其权限组进行权限检查，确定其能访问的功能。