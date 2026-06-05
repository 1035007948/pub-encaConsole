# 环境噪声投诉采样证据归档端到端协同控制台

## 项目简介

本项目是一个环境噪声投诉采样证据归档的端到端协同控制台系统，面向环保投诉处理员、采样人员、企业整改负责人，围绕投诉单、采样点位、噪声读数、时段分类、证据附件、整改措施完成安排采样点位、归档噪声证据、追踪整改措施和生成复测结论。

## 技术栈

### 前端
- React 18 + TypeScript
- Vite (构建工具)
- Mantine (UI组件库)
- Zustand (状态管理)
- TanStack Table (表格组件)
- Zod (数据校验)
- React Router (路由)

### 后端
- Go 1.21+
- Gin (Web框架)
- PostgreSQL (数据库)
- GORM (ORM)

## 快速开始

### 1. 启动PostgreSQL数据库

```bash
```

等待数据库启动完成（约10秒）。

### 2. 后端设置

```bash
cd backend

# 安装Go依赖
go mod download

# 运行数据库迁移和种子数据
go run src/cmd/init/main.go

# 启动后端服务
go run src/cmd/server/main.go
```

后端服务将在 http://localhost:8080 启动。

### 3. 前端设置

```bash
cd frontend

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

前端服务将在 http://localhost:5173 启动。

## API接口说明

### 核心实体CRUD

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/complaints | 获取投诉单列表 |
| POST | /api/complaints | 创建投诉单 |
| GET | /api/complaints/:id | 获取投诉单详情 |
| PUT | /api/complaints/:id | 更新投诉单 |
| DELETE | /api/complaints/:id | 删除投诉单 |
| GET | /api/sampling-points | 获取采样点位列表 |
| POST | /api/sampling-points | 创建采样点位 |
| GET | /api/noise-readings | 获取噪声读数列表 |
| POST | /api/noise-readings | 创建噪声读数 |
| GET | /api/time-periods | 获取时段分类列表 |
| POST | /api/time-periods | 创建时段分类 |
| GET | /api/evidence-attachments | 获取证据附件列表 |
| POST | /api/evidence-attachments | 上传证据附件 |
| GET | /api/rectification-measures | 获取整改措施列表 |
| POST | /api/rectification-measures | 创建整改措施 |
| GET | /api/retest-records | 获取复测记录列表 |
| POST | /api/retest-records | 创建复测记录 |

### 状态流转

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/complaints/:id/transition | 投诉单状态流转 |
| POST | /api/sampling-points/:id/transition | 采样点位状态流转 |

### 批量操作

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/time-periods/batch-import | 时段分类批量导入预检 |
| POST | /api/complaints/batch-action | 投诉单批量操作 |

### 异常分诊

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/anomalies | 获取异常事件列表 |
| POST | /api/anomalies | 创建异常事件 |
| PUT | /api/anomalies/:id/resolve | 解决异常事件 |

### 领域计算

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/calculate/priority | 计算处理优先级 |
| POST | /api/calculate/completeness | 计算证据完整度 |
| POST | /api/calculate/compliance | 校验采样时段合规性 |

### 统计聚合

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/statistics/dashboard | 获取统计驾驶舱数据 |
| GET | /api/statistics/completeness | 获取证据完整度统计 |
| GET | /api/statistics/rectification-rate | 获取整改闭环率统计 |
| GET | /api/statistics/retest-pass-rate | 获取复测通过率统计 |

### 归档快照

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/archive/snapshot | 创建归档快照 |
| GET | /api/archive/snapshots | 获取归档快照列表 |
| GET | /api/archive/export | 导出归档数据 |

### 规则配置

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/rules | 获取规则配置列表 |
| POST | /api/rules | 创建规则配置 |
| PUT | /api/rules/:id | 更新规则配置 |

### 审计日志

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/audit-logs | 获取审计日志列表 |

### Seed和健康检查

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/health | 健康检查 |
| POST | /api/seed/reset | 重置种子数据 |
| GET | /api/seed/browse | 浏览种子数据 |

## 测试

### 后端测试

```bash
cd backend
go test ./tests/... -v
```

### 前端测试

```bash
cd frontend
npm run test
```

## 核心功能模块

1. **投诉单创建与流转** - 创建投诉单并进行状态流转管理
2. **安排采样点位前端工作台** - 工作台式界面安排采样点位
3. **归档噪声证据后端计算与校验** - 后端计算等效声级并校验证据
4. **采样时段不合规异常分诊** - 自动识别不合规时段并分诊
5. **噪声读数详情、证据和历史** - 查看噪声读数详情和历史记录
6. **证据完整度统计驾驶舱** - 可视化统计证据完整度
7. **时段分类批量导入和复核** - 批量导入时段分类并复核
8. **规则配置、命中解释和审计日志** - 配置规则并查看审计日志
9. **归档快照与导出** - 创建归档快照并导出数据
10. **Seed数据浏览和接口健康检查** - 浏览种子数据并检查接口健康

## 核心数据实体

- 投诉单 (Complaint)
- 采样点位 (SamplingPoint)
- 噪声读数 (NoiseReading)
- 时段分类 (TimePeriod)
- 证据附件 (EvidenceAttachment)
- 整改措施 (RectificationMeasure)
- 复测记录 (RetestRecord)
- 状态流转记录 (StatusTransition)
- 规则配置 (RuleConfig)
- 异常事件 (AnomalyEvent)

## 领域业务规则

1. 根据投诉等级、采样时段、读数偏差和证据完整度计算处理优先级
2. 采样时段不合规时进入异常队列，记录触发字段、阈值、处理人和处理时限
3. 安排采样点位前校验投诉单、采样点位和噪声读数一致性，缺少关键字段只允许保存草稿
4. 证据完整度、整改闭环率和复测通过率按日/批次/责任角色聚合
5. 状态流转限制下一步动作，驳回必须填写原因

## 常见问题排查

### 数据库连接失败

确保PostgreSQL已启动：
```bash
```

如果未启动，运行：
```bash
```

### 前端无法连接后端

检查后端是否在 http://localhost:8080 运行。前端API配置在 `frontend/src/api/config.ts`。

### 端口冲突

- 前端默认端口：5173
- 后端默认端口：8080
- PostgreSQL默认端口：5432

可在对应配置文件中修改端口。
