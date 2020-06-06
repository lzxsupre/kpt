# ➿ kpt
大三下课设 - 智能门禁系统后端

## 🚀 API

API 根域名：http://xjj.pub:8000

数据库管理后台：http://xjj.pub:9999

API 返回数据均为以下格式

```json
{
    "s": 0,
    "m": "0",
    "d": null
}
```

- `s` - 状态码
  - 0 - 响应成功
  - 其他 - 错误码
- `m` - 状态消息
- `d` - 相应数据

### 💩 登录

#### 发送邮件验证码

- 请求示例

|   方法   |             URI              |      说明      | 需要 token |
| :------: | :--------------------------: | :------------: | :--------: |
| **POST** | /code?addr=1366723936@qq.com | 发送邮件验证码 |     否     |

#### 获取 token

- 请求示例

|   方法   |                   URI                    |      说明      | 需要 token |
| :------: | :--------------------------------------: | :------------: | :--------: |
| **POST** | /token?email=1366723936@qq.com&code=6371 | 获取登陆 token |     否     |

- 请求参数说明
  -  `code` - 邮箱验证码 

⚠️ 请求下面需要 token 的接口时，需要将该获取的 token 作为请求头部的 `access_token` 字段的值，否则接口会直接拦下该请求并返回状态 *403 Forbidden*

### 💩 用户

#### 获取学生信息

- 请求示例

|  方法   |               URI               |              说明              | 需要 token |
| :-----: | :-----------------------------: | :----------------------------: | :--------: |
| **GET** |    /auth/user?uid=2017213056    | 获取学号为2017213056的学生信息 |     是     |
| **GET** | /auth/user?name=谢金锦&status=1 |      获取谢金锦的学生信息      |     是     |
| **GET** |       /auth/user?status=2       |      获取所有管理员的信息      |     是     |
| **GET** |  /auth/user?class_id=08051703   |  获取08051703班的所有学生信息  |     是     |
| **GET** |            以此类推             |            以此类推            |     是     |

- 返回数据

```json
{
    "uid": "2017213056", // 学号
    "cid": "device_id_1", // RFID 卡号
    "class_id": "08051703", // 班级号
    "name": "谢金锦", // 姓名
    "email": "1366723936@qq.com", // 邮箱地址
    "status": 1, // 状态
    "ctime": 1591289287, // 创建时间
    "mtime": 1591289287  // 更新时间
},
```

- 返回字段说明
  - `status` 
    - 0 - 被 ban 用户
    - 1 - 一般用户
    - 2 - 管理员
    - 3 - 豪横管理员
    - 4 - 谢金锦

#### 添加一条学生信息

- 请求示例


|   方法   |    URL     |       说明       | 需要 token |
| :------: | :--------: | :--------------: | :--------: |
| **POST** | /auth/user | 请求JSON数据如下 |     是     |

- 请求数据示例

```json
{
    "uid": "2017213053", // 学号
    "cid": "device_id_3",  // RFID 卡号
    "class_id": "08051703",  // 班级号
    "name": "谢金锦", // 姓名
    "email": "1366723936@qq.com"  // 邮箱地址
}
```

⚠️ `status` 字段会默认设为 1

#### 更新学生信息

- 请求示例


|  方法   |    URL     |       说明       | 需要 token |
| :-----: | :--------: | :--------------: | :--------: |
| **PUT** | /auth/user | 请求JSON数据如下 |     是     |

- 请求数据示例

```json
{
    "uid": "2017213053", // 学号，必须字段！
    "email": "mivinci@qq.com"  // 要更新的邮箱地址
}
```

这样就会把学号为 `2017213056` 的学生的邮箱改为 `mivinci@qq.com`

⚠️ 若尝试使用该接口更改 `status` 字段，则更改无效。

#### 删除一条学生信息

- 请求示例

|    方法    |            URL            |               说明               | 需要 token |
| :--------: | :-----------------------: | :------------------------------: | :--------: |
| **DELETE** | /auth/user?uid=2017213056 | 软删除学号为2017213056学生的信息 |     是     |

⚠️ 该删除接口不会真正删除数据，只会将该学生的 `status` 字段更改为 `0` 

### 💩 打卡

#### 添加一条打卡记录

- 请求示例


|   方法   |    URL    |       说明       | 需要 token |
| :------: | :-------: | :--------------: | :--------: |
| **POST** | /pr/punch | 请求JSON数据如下 |     是     |

- 请求数据示例

```json
{
    "uid": "2017213056",
    "name": "谢金锦",
    "phone": "17783195535",
    "location": "重庆市南岸区重庆邮电大学",
    "is_temperature_ok": true,
    "did_meet_hubei": false,
    "has_symptom": false,
    "is_family_diagnosed": false,
    "did_meet_diagnoses": false,
    "is_family_suspected": false
}
```

#### 按学号获取打卡记录

- 请求示例

|  方法   |                 URL                  |                     说明                     | 需要 token |
| :-----: | :----------------------------------: | :------------------------------------------: | :--------: |
| **GET** |       /pr/punch?uid=2017213056       |        获取学号为2017213056的打卡记录        |     是     |
| **GET** | /pr/punch?uid=2017213056&name=谢金锦 | 获取学号为2017213056且名字为谢金锦的打卡记录 |     是     |
| **GET** |  /pr/punch?is_temperature_ok=false   |           获取体温不正常的打卡记录           |     是     |
| **GET** |               以此类推               |                   以此类推                   |     是     |

- 返回数据示例

```json
{
    "s": 0,
    "m": "0",
    "d": [
        {
            "id": 1,  // ID，该字段不重要
            "uid": "2017213053",  // 学号
            "name": "高寅",  // 姓名
            "phone": "17784450780",  // 手机号码
            "location": "重庆市南岸区重庆邮电大学",  // 当前所在地
            "is_temperature_ok": true,  // 体温是否正常
            "did_meet_hubei": false,  // 是否接触湖北人员
            "has_symptom": false,  // 有无症状
            "is_family_diagnosed": false,  // 本人或家人是否确诊
            "did_meet_diagnoses": false, // 是否接触确诊人员
            "is_family_suspected": false,  // 本人或家人是否为疑似病例
            "ctime": 1591190705  // 打卡时间戳
        }
    ]
}
```

#### 按时间段返回打卡记录

- 请求示例

|  方法   |                       URL                       |                 说明                  | 需要 token |
| :-----: | :---------------------------------------------: | :-----------------------------------: | :--------: |
| **GET** | /pr/punch/between?from=2020-05-30&to=2020-06-04 | 获取2020/5/30到2020/6/3的所有打卡记录 |     是     |

- 返回数据的结构与上一个 API 相同

#### 删除一条打卡记录

- 请求示例

|    方法    |      URL       |          说明           | 需要 token |
| :--------: | :------------: | :---------------------: | :--------: |
| **DELETE** | /pr/punch?id=1 | 永久删除id为1的打卡记录 |     是     |

  

### 💩 门禁

#### 添加一条门禁记录

- 请求方法和地址

|   方法   |   URL    |       说明       | 需要 token |
| :------: | :------: | :--------------: | :--------: |
| **POST** | /ac/scan | 请求JSON数据如下 |     是     |

- 请求数据示例

```json
{
	"uid": "2017213056", // 学号
  "cid": "32_bits_device_id", // RFID 卡号
	"tpt": 36.6  // 体温
}
```

#### 按学号获取门禁记录

- 请求方法和地址

|  方法   |           URL           |                说明                | 需要 token |
| :-----: | :---------------------: | :--------------------------------: | :--------: |
| **GET** | /ac/scan?uid=2017213056 | 获取学号为2017213056学生的门禁记录 |     是     |

- 返回数据示例

```json
{
    "s": 0,
    "m": "0",
    "d": [
        {
            "uid": "2017213056", // 学号
            "cid": "32_bits_device_id", // RFID 卡号
	          "tpt": 36.6,  // 体温
            "ctime": 1591190705 // 刷卡时间
        }
    ]
}
```

#### 按时间段返回门禁记录

- 请求方法和地址

|  方法   |                      URL                       |                 说明                  | 需要 token |
| :-----: | :--------------------------------------------: | :-----------------------------------: | :--------: |
| **GET** | /ac/scan/between?from=2020-05-30&to=2020-06-04 | 获取2020/5/30到2020/6/3的所有门禁记录 |     是     |

- 返回数据的结构与上一个 API 相同

#### 删除一条打卡记录

- 请求示例

|    方法    |      URL      |          说明           | 需要 token |
| :--------: | :-----------: | :---------------------: | :--------: |
| **DELETE** | /ac/scan?id=1 | 永久删除id为1的门禁记录 |     是     |

## 🌚 Attention

⚠️ 这些 API 都还没接用户认证，尽管塞库。