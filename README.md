# ➿ kpt
大三下课设 - 智能门禁系统后端

## 🚀 API

### 打卡

#### 添加一条打卡记录

- 请求方法和地址

  **POST**  http://xjj.pub:8000/pr/punch

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

- 返回数据示例

```json
{
    "s": 0,
    "m": "0",
    "d": null
}
```

"s" 为 0，则表示请求成功，下同。

#### 按学号获取打卡记录

- 请求方法和地址

  **GET**  http://xjj.pub:8000/pr/punch/2017213056

- 返回数据示例

```json
{
    "s": 0,
    "m": "0",
    "d": [
        {
            "id": 1,
            "uid": "2017213053",
            "name": "高寅",
            "phone": "17784450780",
            "location": "重庆市南岸区重庆邮电大学",
            "is_temperature_ok": true,
            "did_meet_hubei": false,
            "has_symptom": false,
            "is_family_diagnosed": false,
            "did_meet_diagnoses": false,
            "is_family_suspected": false,
            "ctime": 1591190705
        }
    ]
}
```

#### 按时间段返回打卡记录

- 请求方法和地址

  **GET**  http://xjj.pub:8000/pr/punch?from=2020-05-30&to=2020-06-04

- 返回数据的结构与上一个 API 相同

------

### 门禁

#### 添加一条门禁记录

- 请求方法和地址

  **POST**  http://xjj.pub:8000/ac/scan

- 请求数据示例

```json
{
	"uid": "2017213056",
  "cid": "32_bits_device_id",
	"tpt": 36.6
}
```

- 返回数据示例

```json
{
    "s": 0,
    "m": "0",
    "d": null
}
```

"s" 为 0，则表示请求成功，下同。

#### 按学号获取门禁记录

- 请求方法和地址

  **GET**  http://xjj.pub:8000/ac/scan/2017213056

- 返回数据示例

```json
{
    "s": 0,
    "m": "0",
    "d": [
        {
            "uid": "2017213056",
            "cid": "32_bits_device_id",
	          "tpt": 36.6
        }
    ]
}
```

#### 按时间段返回打卡记录

- 请求方法和地址

  **GET**  http://xjj.pub:8000/ac/scan?from=2020-05-30&to=2020-06-04

- 返回数据的结构与上一个 API 相同

## 🌚 Attention

这些 API 都还没接用户认证，尽管攻击。