<div style="font-family: 'Fira Code','PingFang SC'">

# 后端API接口文档
`@Samshui - 05/11`

## User：用户模块
`Users表`

```
r.POST("/api/user/register", controller.Register)
r.POST("/api/user/login", controller.Login)
r.GET("/api/user/info", middleware.AuthMiddleware(), controller.Info)
r.POST("/api/user/changeTelephone", controller.ChangeTelephone)
r.POST("/api/user/telephoneIsExist", controller.TelephoneIsExisted)
```

### register：注册
> 新用户在注册后方可有权进入网站

`URL: /api/user/register`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| register | name, studentID, telephone, password | 用户注册 | POST |

### login：登录
> 用户在登录后方可进入网站

`URL: /api/user/login`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| login | telephone, password | 用户登录 | POST |

### info：获取用户信息
> 用户点击个人信息时获取对应用户信息

`URL: /api/user/info`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| info | 无（中间件获取报文Header并解析） | 获取用户信息 | GET |

### changeTelephone：修改用户联系方式
> 用户调用该方法进行联系方式的修改

`URL: /api/user/changeTelephone`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| changeTelephone | newTelephone, oldTelephone | 修改用户联系方式 | POST |

### telephoneIsExist：联系方式是否已被使用
> 查看联系方式是否被占用

`URL: /api/user/telephoneIsExist`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| telephoneIsExist | telephone | 查询联系方式是否被使用 | POST |

## Experiment：实验模块
`Experiments表`

```
r.POST("/api/experiment/add", controller.AddExperiment)
r.POST("/api/experiment/delete", controller.DeleteExperiment)
r.POST("/api/experiment/getByLabel", controller.GetExperimentByLabel)
r.POST("/api/experiment/all", controller.GetAllExperiments)
r.POST("/api/experiment/getByEID", controller.GetExperimentByEID)
```

### add：新增实验
> 新增实验

`URL: /api/experiment/add`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| add | EID, ELabel, EName, lab, EM, EN, EE, Site | 新增实验 | POST |

### delete：删除实验
> 删除实验

`URL: /api/experiment/delete`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| delete | EID | 删除实验 | POST |

### getByLabel：获取栏目实验
> 通过ELabel获取对应的栏目实验

`URL: /api/experiment/getByLabel`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| getByLabel | ELabel | 获取栏目实验 | POST |

### all：获取所有实验
> 获取所有的栏目实验

`URL: /api/experiment/all`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| all | 无 | 获取所有实验 | POST |

### getByEID：获取指定实验
> 获取指定实验

`URL: /api/experiment/getByEID`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| getByEID | EID | 获取指定实验 | POST |

## Record：预约记录模块
`Records表`

```
r.POST("/api/record/add", controller.AddRecord)
r.POST("/api/record/delete", controller.DeleteRecord)
r.POST("/api/record/getAll", controller.GetAllRecordByEID)
r.POST("/api/record/getRecordsSites", controller.GetAllSiteSelected)
r.POST("/api/record/getUserRecords", controller.GetAllRecordsByStudentID)
```

### add：增加新记录
> 增加新的记录

`URL: /api/record/add`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| add | EID, UID, date, time, site | 增加新预约记录 | POST |

### delete：删除记录
> 删除预约记录

`URL: /api/record/delete`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| delete | EID, UID, date, time, site | 删除预约记录 | POST |

### getAll：获取指定实验的全部预约记录
> 获取指定实验的全部预约记录

`URL: /api/record/getAll`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| getAll | EID | 获取指定实验的全部预约记录 | POST |

### getRecordsSites：获取当前实验所有已选的位置
> 获取当前实验所有已选的位置

`URL: /api/record/getRecordsSites`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| getRecordsSites | EID, date, time | 获取当前实验所有已选的位置 | POST |

### getUserRecords：获取当前用户所有预约记录
> 获取当前用户所有预约记录

`URL: /api/record/getUserRecords`

| 名称 | 参数 | 说明 | 方法 |
| :---: | :---: | :---: | :---: |
| getUserRecords | UID | 获取当前用户所有预约记录 | POST |
</div>