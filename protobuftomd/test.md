


<a name="PingReq_Sub_Test_User"></a>

### PingReq_Sub_Test_User
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|nameUser|string|true|用户名|


<a name="PingReq_Sub_Test"></a>

### PingReq_Sub_Test
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|nameTest|string|false|测试名|
|user|[PingReq_Sub_Test_User](#PingReq_Sub_Test_User)|true|用户|


<a name="PingReq_Sub_Tests"></a>

### PingReq_Sub_Tests
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|nameTest|string|false||


<a name="PingReq_Sub"></a>

### PingReq_Sub
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|age|string|true|年龄|
|test|[PingReq_Sub_Test](#PingReq_Sub_Test)|false||
|tests|[PingReq_Sub_Tests](#PingReq_Sub_Tests)|false||


<a name="PingReq"></a>

### PingReq
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|name|string|false|用户名|
|sub|[PingReq_Sub](#PingReq_Sub)|false||


<a name="User"></a>

### User
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|nameUser|string|false||


<a name="PingRes_Sub"></a>

### PingRes_Sub
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|age|string|false||


<a name="PingRes"></a>

### PingRes
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|name|string|false||
|sub|[PingRes_Sub](#PingRes_Sub)|false||


<a name="Operator"></a>

### Operator
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|[]string|false||
|channelId|string|false||


<a name="GetGiftbagReq"></a>

### GetGiftbagReq
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|name|string|false||
|operator|[Operator](#Operator)|false||


<a name="GetGiftbagRes"></a>

### GetGiftbagRes
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||


<a name="MetaRes"></a>

### MetaRes
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||