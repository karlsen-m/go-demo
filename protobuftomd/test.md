
# Interactluckyactivity-Service 接口文档


- ## [InteractluckyactivityPing](#Ping)      
- ## [InteractluckyactivityCreateActivity](#CreateActivity)      
- ## [InteractluckyactivityEditActivity](#EditActivity)      
- ## [InteractluckyactivityGetActivityList](#GetActivityList)      
- ## [InteractluckyactivityGetActivityDetails](#GetActivityDetails)      
- ## [InteractluckyactivityGetActivityStatus](#GetActivityStatus)      
- ## [InteractluckyactivityDelActivity](#DelActivity)      
- ## [InteractluckyactivityCreatePrizePool](#CreatePrizePool)      
- ## [InteractluckyactivityEditPrizePool](#EditPrizePool)      
- ## [InteractluckyactivityGetPrizePoolList](#GetPrizePoolList)      
- ## [InteractluckyactivityGetPrizePoolDetails](#GetPrizePoolDetails)      
- ## [InteractluckyactivityAdjustPrizeTotalNum](#AdjustPrizeTotalNum)      
- ## [InteractluckyactivityDelPrizePool](#DelPrizePool)      
- ## [InteractluckyactivityDelPrize](#DelPrize)      
- ## [InteractluckyactivityGetPrizeOrgList](#GetPrizeOrgList)      
- ## [InteractluckyactivityEditPrizeOrgNum](#EditPrizeOrgNum)      
- ## [InteractluckyactivityDeleteReceiveRecord](#DeleteReceiveRecord)      
- ## [InteractluckyactivityCreatePrizeWhitelist](#CreatePrizeWhitelist)      
- ## [InteractluckyactivityDelPrizeWhitelist](#DelPrizeWhitelist)      
- ## [InteractluckyactivityGetPrizeWhitelistList](#GetPrizeWhitelistList)      
- ## [InteractluckyactivityAdjustPrizeWhitelistNum](#AdjustPrizeWhitelistNum)      
- ## [InteractluckyactivityGetReceiveList](#GetReceiveList)      
- ## [InteractluckyactivityReplacement](#Replacement)      
- ## [InteractluckyactivityAllReplacement](#AllReplacement)      
- ## [InteractluckyactivityEditAppointmentOrg](#EditAppointmentOrg)      
- ## [InteractluckyactivityBatchEditAppointmentOrg](#BatchEditAppointmentOrg)      
- ## [InteractluckyactivityParticipateCodeList](#ParticipateCodeList)      
- ## [InteractluckyactivityDelParticipateCodeList](#DelParticipateCodeList)      
- ## [InteractluckyactivityMobileWhiteList](#MobileWhiteList)      
- ## [InteractluckyactivityDelMobileWhiteList](#DelMobileWhiteList)      
- ## [InteractluckyactivityRepeatMobileWhiteList](#RepeatMobileWhiteList)      
- ## [InteractluckyactivityRepeatMobileWhiteEdit](#RepeatMobileWhiteEdit)      
- ## [InteractluckyactivityDelRepeatMobileWhite](#DelRepeatMobileWhite)      
- ## [InteractluckyactivityAdminReceiveAddressEdit](#AdminReceiveAddressEdit)      
- ## [InteractluckyactivityActivityGiveBackPoint](#ActivityGiveBackPoint)      
- ## [InteractluckyactivityBatchCreateParticipateCode](#BatchCreateParticipateCode)      
- ## [InteractluckyactivityActivityListGetStaff](#ActivityListGetStaff)      
- ## [InteractluckyactivityPrizePoolListGetStaff](#PrizePoolListGetStaff)      
- ## [InteractluckyactivityMarketerCodesCreate](#MarketerCodesCreate)      
- ## [InteractluckyactivityMarketerCodesInfo](#MarketerCodesInfo)      
- ## [InteractluckyactivityReceiveListGetStaff](#ReceiveListGetStaff)      
- ## [InteractluckyactivityActivityInfoGetStaff](#ActivityInfoGetStaff)      
- ## [InteractluckyactivityReceiveVerifyStaff](#ReceiveVerifyStaff)      
- ## [InteractluckyactivityMarketerGetReceiveDetails](#MarketerGetReceiveDetails)      
- ## [InteractluckyactivityGetMobileCode](#GetMobileCode)      
- ## [InteractluckyactivityMarketerCreateUser](#MarketerCreateUser)      
- ## [InteractluckyactivityMarketerGetPrizeList](#MarketerGetPrizeList)      
- ## [InteractluckyactivityMarketerIssuePrize](#MarketerIssuePrize)      
- ## [InteractluckyactivityActivityInfoGet](#ActivityInfoGet)      
- ## [InteractluckyactivityActivityQualificationVerify](#ActivityQualificationVerify)      
- ## [InteractluckyactivityActivityPartakeCodeVerify](#ActivityPartakeCodeVerify)      
- ## [InteractluckyactivityActivityPartakeCodeVerifyInfo](#ActivityPartakeCodeVerifyInfo)      
- ## [InteractluckyactivityActivityQualificationVerifyPoint](#ActivityQualificationVerifyPoint)      
- ## [InteractluckyactivityActivityListGetUser](#ActivityListGetUser)      
- ## [InteractluckyactivityPrizeListGet](#PrizeListGet)      
- ## [InteractluckyactivityPrizeListGetByActIdAndUserId](#PrizeListGetByActIdAndUserId)      
- ## [InteractluckyactivityLuckDrawUser](#LuckDrawUser)      
- ## [InteractluckyactivityDirectReceivePrizeUser](#DirectReceivePrizeUser)      
- ## [InteractluckyactivityReceiveRecordListUser](#ReceiveRecordListUser)      
- ## [InteractluckyactivityReceiveMoneyUser](#ReceiveMoneyUser)      
- ## [InteractluckyactivityReceiveAddress](#ReceiveAddress)      
- ## [InteractluckyactivityReceiveChoiceOrg](#ReceiveChoiceOrg)      
- ## [InteractluckyactivityPlayUserList](#PlayUserList)      
- ## [InteractluckyactivityTakeFromOrgIdList](#TakeFromOrgIdList)      
- ## [InteractluckyactivityReceiveAddressEdit](#ReceiveAddressEdit)      
- ## [InteractluckyactivityReceiveInfoGetUser](#ReceiveInfoGetUser)      
- ## [InteractluckyactivityPrizeInfoGet](#PrizeInfoGet)      
- ## [InteractluckyactivityRechargeCallback](#RechargeCallback)      
- ## [InteractluckyactivityGetVerifyCardDetail](#GetVerifyCardDetail)      
- ## [InteractluckyactivityCreateVerifyCardOrder](#CreateVerifyCardOrder)      


<a name="Ping"></a>

## 前端method: InteractluckyactivityPing

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |

### 请求例子:
````

````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="CreateActivity"></a>

## 前端method: InteractluckyactivityCreateActivity

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|name|string|false||
|operator|[Operator](#Operator)|false||
|orgId|string|false||
|string|[//](#//)|false||
|timeRule|[TimeRule](#TimeRule)|false||
|actRule|string|false||
|takePartTotalNum|string|false||
|takePartDayNum|string|false||
|takePartUserNum|string|false||
|takePartDayUserNum|string|false||
|isMarketerExclusive|bool|false||
|userCodeNum|string|false||
|userDayCodeNum|string|false||
|takePartCodeNum|string|false||
|isShare|bool|false||
|shareTitle|string|false||
|shareImgId|string|false||
|status|string|false||
|addreVerifyType|string|false||
|addreIds|[]string|false||
|verifyType|string|false||
|verifyMemberLevel|[]string|false||
|verifyApiType|string|false||
|isApiMemberLevelMatchPrizePool|bool|false||
|thirdActId|string|false||
|verifyDesc|string|false||
|verifyNotPass|string|false||
|verifyPass|string|false||
|takePartIsBucklePoint|bool|false||
|pointType|string|false||
|pointNum|string|false||
|isMarketerSelectPrizePool|bool|false||
|recommendImgId|string|false||
|imgId|string|false||
|buttonImgId|string|false||
|bgColor|string|false||
|textColor|string|false||
|isShowUserTakePartNum|bool|false||
|userTakePartInitNum|string|false||
|isShowLuckyUser|bool|false||
|luckyUserNum|string|false||
|pushReceiveDataType|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|luckyBgImgId|string|false||
|isShowActivityGift|bool|false||
|eventId|string|false||
|isVerifyCard|bool|false||
|verifyCardCode|string|false||
|verifyCardNumber|string|false||
|verifyCardDesc|string|false||
|verifyCardPrice|string|false||
|isStaffPartakeCode|bool|false||
|isStaffRegistration|bool|false||
|staffRegisterConfig|[]string|false||
|isRegisterVerifyPualification|bool|false||
|isVerifyWhiteList|bool|false||
|whiteListIds|[]string|false||
|isVerifyWhiteListStaff|bool|false||
|isVerifyThirdApiStaff|bool|false||
|gameType|string|false||
|digiccyWalletId|string|false||
|digiccyActName|string|false||
|digiccyActNo|string|false||
|digiccySonActNo|string|false||
|isSendSms|bool|false||
|isDigiccy|bool|false||
|isOpenPrizePushData|bool|false||
|buttonMsg|string|false||
|isRepeatVerify|bool|false||
|birthdayVerify|string|false||
|birthdayVerifyStartMonth|string|false||
|birthdayVerifyEndMonth|string|false||
|birthdayVerifyBeforeDay|string|false||
|birthdayVerifyAfterDay|string|false||
|isGetApiTakePartNum|bool|false||

<details>
 <summary>
 <code>TimeRule </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|rule|string|false||
|startedAt|string|false||
|endedAt|string|false||
|ruleDay|string|false||
|ruleWeek|string|false||
|ruleTime|string|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"name":"name",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"orgId":"orgId",
"string":{},
"timeRule":{
"rule":"rule",
"startedAt":"startedAt",
"endedAt":"endedAt",
"ruleDay":"ruleDay",
"ruleWeek":"ruleWeek",
"ruleTime":"ruleTime"
},
"actRule":"actRule",
"takePartTotalNum":"takePartTotalNum",
"takePartDayNum":"takePartDayNum",
"takePartUserNum":"takePartUserNum",
"takePartDayUserNum":"takePartDayUserNum",
"isMarketerExclusive":true,
"userCodeNum":"userCodeNum",
"userDayCodeNum":"userDayCodeNum",
"takePartCodeNum":"takePartCodeNum",
"isShare":true,
"shareTitle":"shareTitle",
"shareImgId":"shareImgId",
"status":"status",
"addreVerifyType":"addreVerifyType",
"addreIds":["addreIds"],
"verifyType":"verifyType",
"verifyMemberLevel":["verifyMemberLevel"],
"verifyApiType":"verifyApiType",
"isApiMemberLevelMatchPrizePool":true,
"thirdActId":"thirdActId",
"verifyDesc":"verifyDesc",
"verifyNotPass":"verifyNotPass",
"verifyPass":"verifyPass",
"takePartIsBucklePoint":true,
"pointType":"pointType",
"pointNum":"pointNum",
"isMarketerSelectPrizePool":true,
"recommendImgId":"recommendImgId",
"imgId":"imgId",
"buttonImgId":"buttonImgId",
"bgColor":"bgColor",
"textColor":"textColor",
"isShowUserTakePartNum":true,
"userTakePartInitNum":"userTakePartInitNum",
"isShowLuckyUser":true,
"luckyUserNum":"luckyUserNum",
"pushReceiveDataType":"pushReceiveDataType",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"luckyBgImgId":"luckyBgImgId",
"isShowActivityGift":true,
"eventId":"eventId",
"isVerifyCard":true,
"verifyCardCode":"verifyCardCode",
"verifyCardNumber":"verifyCardNumber",
"verifyCardDesc":"verifyCardDesc",
"verifyCardPrice":"verifyCardPrice",
"isStaffPartakeCode":true,
"isStaffRegistration":true,
"staffRegisterConfig":["staffRegisterConfig"],
"isRegisterVerifyPualification":true,
"isVerifyWhiteList":true,
"whiteListIds":["whiteListIds"],
"isVerifyWhiteListStaff":true,
"isVerifyThirdApiStaff":true,
"gameType":"gameType",
"digiccyWalletId":"digiccyWalletId",
"digiccyActName":"digiccyActName",
"digiccyActNo":"digiccyActNo",
"digiccySonActNo":"digiccySonActNo",
"isSendSms":true,
"isDigiccy":true,
"isOpenPrizePushData":true,
"buttonMsg":"buttonMsg",
"isRepeatVerify":true,
"birthdayVerify":"birthdayVerify",
"birthdayVerifyStartMonth":"birthdayVerifyStartMonth",
"birthdayVerifyEndMonth":"birthdayVerifyEndMonth",
"birthdayVerifyBeforeDay":"birthdayVerifyBeforeDay",
"birthdayVerifyAfterDay":"birthdayVerifyAfterDay",
"isGetApiTakePartNum":true
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ActivityInfo](#ActivityInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ActivityInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|name|string|false||
|operator|[Operator](#Operator)|false||
|orgId|string|false||
|string|[//](#//)|false||
|timeRule|[TimeRule](#TimeRule)|false||
|actRule|string|false||
|takePartTotalNum|string|false||
|takePartDayNum|string|false||
|takePartUserNum|string|false||
|takePartDayUserNum|string|false||
|isMarketerExclusive|bool|false||
|userCodeNum|string|false||
|userDayCodeNum|string|false||
|takePartCodeNum|string|false||
|isShare|bool|false||
|shareTitle|string|false||
|shareImg|[ImageData](#ImageData)|false||
|status|string|false||
|addreVerifyType|string|false||
|addreData|[][Addre](#Addre)|false||
|verifyType|string|false||
|verifyMemberLevel|[]string|false||
|verifyApiType|string|false||
|isApiMemberLevelMatchPrizePool|bool|false||
|thirdActId|string|false||
|verifyDesc|string|false||
|verifyNotPass|string|false||
|verifyPass|string|false||
|takePartIsBucklePoint|bool|false||
|pointType|string|false||
|pointNum|string|false||
|isMarketerSelectPrizePool|bool|false||
|recommendImg|[ImageData](#ImageData)|false||
|img|[ImageData](#ImageData)|false||
|buttonImg|[ImageData](#ImageData)|false||
|bgColor|string|false||
|textColor|string|false||
|isShowUserTakePartNum|bool|false||
|userTakePartInitNum|string|false||
|isShowLuckyUser|bool|false||
|luckyUserNum|string|false||
|id|string|false||
|activityStatus|string|false||
|orgData|[OrgData](#OrgData)|false||
|createdAt|string|false||
|pushReceiveDataType|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|luckyBgImg|[ImageData](#ImageData)|false||
|isShowActivityGift|bool|false||
|eventId|string|false||
|isVerifyCard|bool|false||
|verifyCardCode|string|false||
|verifyCardNumber|string|false||
|verifyCardDesc|string|false||
|verifyCardPrice|string|false||
|isStaffPartakeCode|bool|false||
|isStaffRegistration|bool|false||
|staffRegisterConfig|[]string|false||
|isRegisterVerifyPualification|bool|false||
|isVerifyWhiteList|bool|false||
|whiteListList|[][WhiteListData](#WhiteListData)|false||
|isVerifyWhiteListStaff|bool|false||
|isVerifyThirdApiStaff|bool|false||
|gameType|string|false||
|digiccyWalletId|string|false||
|digiccyActName|string|false||
|digiccyActNo|string|false||
|digiccySonActNo|string|false||
|isSendSms|bool|false||
|isDigiccy|bool|false||
|isOpenPrizePushData|bool|false||
|buttonMsg|string|false||
|isRepeatVerify|bool|false||
|birthdayVerify|string|false||
|birthdayVerifyStartMonth|string|false||
|birthdayVerifyEndMonth|string|false||
|birthdayVerifyBeforeDay|string|false||
|birthdayVerifyAfterDay|string|false||
|isGetApiTakePartNum|bool|false||

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

<details>
 <summary>
 <code>WhiteListData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||

</details>

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

<details>
 <summary>
 <code>TimeRule </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|rule|string|false||
|startedAt|string|false||
|endedAt|string|false||
|ruleDay|string|false||
|ruleWeek|string|false||
|ruleTime|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

<details>
 <summary>
 <code>OrgData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|bool|[//](#//)|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"name":"name",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"orgId":"orgId",
"string":{},
"timeRule":{
"rule":"rule",
"startedAt":"startedAt",
"endedAt":"endedAt",
"ruleDay":"ruleDay",
"ruleWeek":"ruleWeek",
"ruleTime":"ruleTime"
},
"actRule":"actRule",
"takePartTotalNum":"takePartTotalNum",
"takePartDayNum":"takePartDayNum",
"takePartUserNum":"takePartUserNum",
"takePartDayUserNum":"takePartDayUserNum",
"isMarketerExclusive":true,
"userCodeNum":"userCodeNum",
"userDayCodeNum":"userDayCodeNum",
"takePartCodeNum":"takePartCodeNum",
"isShare":true,
"shareTitle":"shareTitle",
"shareImg":{
"id":"id",
"name":"name",
"url":"url"
},
"status":"status",
"addreVerifyType":"addreVerifyType",
"addreData":[
	{
"id":"id",
"fullName":"fullName"
}
],
"verifyType":"verifyType",
"verifyMemberLevel":["verifyMemberLevel"],
"verifyApiType":"verifyApiType",
"isApiMemberLevelMatchPrizePool":true,
"thirdActId":"thirdActId",
"verifyDesc":"verifyDesc",
"verifyNotPass":"verifyNotPass",
"verifyPass":"verifyPass",
"takePartIsBucklePoint":true,
"pointType":"pointType",
"pointNum":"pointNum",
"isMarketerSelectPrizePool":true,
"recommendImg":{
"id":"id",
"name":"name",
"url":"url"
},
"img":{
"id":"id",
"name":"name",
"url":"url"
},
"buttonImg":{
"id":"id",
"name":"name",
"url":"url"
},
"bgColor":"bgColor",
"textColor":"textColor",
"isShowUserTakePartNum":true,
"userTakePartInitNum":"userTakePartInitNum",
"isShowLuckyUser":true,
"luckyUserNum":"luckyUserNum",
"id":"id",
"activityStatus":"activityStatus",
"orgData":{
"id":"id",
"name":"name",
"bool":{}
},
"createdAt":"createdAt",
"pushReceiveDataType":"pushReceiveDataType",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"luckyBgImg":{
"id":"id",
"name":"name",
"url":"url"
},
"isShowActivityGift":true,
"eventId":"eventId",
"isVerifyCard":true,
"verifyCardCode":"verifyCardCode",
"verifyCardNumber":"verifyCardNumber",
"verifyCardDesc":"verifyCardDesc",
"verifyCardPrice":"verifyCardPrice",
"isStaffPartakeCode":true,
"isStaffRegistration":true,
"staffRegisterConfig":["staffRegisterConfig"],
"isRegisterVerifyPualification":true,
"isVerifyWhiteList":true,
"whiteListList":[
	{
"id":"id",
"name":"name"
}
],
"isVerifyWhiteListStaff":true,
"isVerifyThirdApiStaff":true,
"gameType":"gameType",
"digiccyWalletId":"digiccyWalletId",
"digiccyActName":"digiccyActName",
"digiccyActNo":"digiccyActNo",
"digiccySonActNo":"digiccySonActNo",
"isSendSms":true,
"isDigiccy":true,
"isOpenPrizePushData":true,
"buttonMsg":"buttonMsg",
"isRepeatVerify":true,
"birthdayVerify":"birthdayVerify",
"birthdayVerifyStartMonth":"birthdayVerifyStartMonth",
"birthdayVerifyEndMonth":"birthdayVerifyEndMonth",
"birthdayVerifyBeforeDay":"birthdayVerifyBeforeDay",
"birthdayVerifyAfterDay":"birthdayVerifyAfterDay",
"isGetApiTakePartNum":true
}
}
````



<a name="EditActivity"></a>

## 前端method: InteractluckyactivityEditActivity

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|name|string|false||
|operator|[Operator](#Operator)|false||
|orgId|string|false||
|startedAt|string|false||
|endedAt|string|false||
|timeRule|[TimeRule](#TimeRule)|false||
|actRule|string|false||
|takePartTotalNum|string|false||
|takePartDayNum|string|false||
|takePartUserNum|string|false||
|takePartDayUserNum|string|false||
|isMarketerExclusive|bool|false||
|userCodeNum|string|false||
|userDayCodeNum|string|false||
|takePartCodeNum|string|false||
|isShare|bool|false||
|shareTitle|string|false||
|shareImgId|string|false||
|status|string|false||
|addreVerifyType|string|false||
|addreIds|[]string|false||
|verifyType|string|false||
|verifyMemberLevel|[]string|false||
|verifyApiType|string|false||
|isApiMemberLevelMatchPrizePool|bool|false||
|thirdActId|string|false||
|verifyDesc|string|false||
|verifyNotPass|string|false||
|verifyPass|string|false||
|takePartIsBucklePoint|bool|false||
|pointType|string|false||
|pointNum|string|false||
|isMarketerSelectPrizePool|bool|false||
|recommendImgId|string|false||
|imgId|string|false||
|buttonImgId|string|false||
|bgColor|string|false||
|textColor|string|false||
|isShowUserTakePartNum|bool|false||
|userTakePartInitNum|string|false||
|isShowLuckyUser|bool|false||
|luckyUserNum|string|false||
|id|string|false||
|pushReceiveDataType|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|luckyBgImgId|string|false||
|isShowActivityGift|bool|false||
|eventId|string|false||
|isVerifyCard|bool|false||
|verifyCardCode|string|false||
|verifyCardNumber|string|false||
|verifyCardDesc|string|false||
|verifyCardPrice|string|false||
|isStaffPartakeCode|bool|false||
|isStaffRegistration|bool|false||
|staffRegisterConfig|[]string|false||
|isRegisterVerifyPualification|bool|false||
|isVerifyWhiteList|bool|false||
|whiteListIds|[]string|false||
|isVerifyWhiteListStaff|bool|false||
|isVerifyThirdApiStaff|bool|false||
|gameType|string|false||
|digiccyWalletId|string|false||
|digiccyActName|string|false||
|digiccyActNo|string|false||
|digiccySonActNo|string|false||
|isSendSms|bool|false||
|isDigiccy|bool|false||
|isOpenPrizePushData|bool|false||
|buttonMsg|string|false||
|isRepeatVerify|bool|false||
|birthdayVerify|string|false||
|birthdayVerifyStartMonth|string|false||
|birthdayVerifyEndMonth|string|false||
|birthdayVerifyBeforeDay|string|false||
|birthdayVerifyAfterDay|string|false||
|isGetApiTakePartNum|bool|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

<details>
 <summary>
 <code>TimeRule </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|rule|string|false||
|startedAt|string|false||
|endedAt|string|false||
|ruleDay|string|false||
|ruleWeek|string|false||
|ruleTime|string|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

### 请求例子:
````
{
"name":"name",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"orgId":"orgId",
"startedAt":"startedAt",
"endedAt":"endedAt",
"timeRule":{
"rule":"rule",
"startedAt":"startedAt",
"endedAt":"endedAt",
"ruleDay":"ruleDay",
"ruleWeek":"ruleWeek",
"ruleTime":"ruleTime"
},
"actRule":"actRule",
"takePartTotalNum":"takePartTotalNum",
"takePartDayNum":"takePartDayNum",
"takePartUserNum":"takePartUserNum",
"takePartDayUserNum":"takePartDayUserNum",
"isMarketerExclusive":true,
"userCodeNum":"userCodeNum",
"userDayCodeNum":"userDayCodeNum",
"takePartCodeNum":"takePartCodeNum",
"isShare":true,
"shareTitle":"shareTitle",
"shareImgId":"shareImgId",
"status":"status",
"addreVerifyType":"addreVerifyType",
"addreIds":["addreIds"],
"verifyType":"verifyType",
"verifyMemberLevel":["verifyMemberLevel"],
"verifyApiType":"verifyApiType",
"isApiMemberLevelMatchPrizePool":true,
"thirdActId":"thirdActId",
"verifyDesc":"verifyDesc",
"verifyNotPass":"verifyNotPass",
"verifyPass":"verifyPass",
"takePartIsBucklePoint":true,
"pointType":"pointType",
"pointNum":"pointNum",
"isMarketerSelectPrizePool":true,
"recommendImgId":"recommendImgId",
"imgId":"imgId",
"buttonImgId":"buttonImgId",
"bgColor":"bgColor",
"textColor":"textColor",
"isShowUserTakePartNum":true,
"userTakePartInitNum":"userTakePartInitNum",
"isShowLuckyUser":true,
"luckyUserNum":"luckyUserNum",
"id":"id",
"pushReceiveDataType":"pushReceiveDataType",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"luckyBgImgId":"luckyBgImgId",
"isShowActivityGift":true,
"eventId":"eventId",
"isVerifyCard":true,
"verifyCardCode":"verifyCardCode",
"verifyCardNumber":"verifyCardNumber",
"verifyCardDesc":"verifyCardDesc",
"verifyCardPrice":"verifyCardPrice",
"isStaffPartakeCode":true,
"isStaffRegistration":true,
"staffRegisterConfig":["staffRegisterConfig"],
"isRegisterVerifyPualification":true,
"isVerifyWhiteList":true,
"whiteListIds":["whiteListIds"],
"isVerifyWhiteListStaff":true,
"isVerifyThirdApiStaff":true,
"gameType":"gameType",
"digiccyWalletId":"digiccyWalletId",
"digiccyActName":"digiccyActName",
"digiccyActNo":"digiccyActNo",
"digiccySonActNo":"digiccySonActNo",
"isSendSms":true,
"isDigiccy":true,
"isOpenPrizePushData":true,
"buttonMsg":"buttonMsg",
"isRepeatVerify":true,
"birthdayVerify":"birthdayVerify",
"birthdayVerifyStartMonth":"birthdayVerifyStartMonth",
"birthdayVerifyEndMonth":"birthdayVerifyEndMonth",
"birthdayVerifyBeforeDay":"birthdayVerifyBeforeDay",
"birthdayVerifyAfterDay":"birthdayVerifyAfterDay",
"isGetApiTakePartNum":true
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ActivityInfo](#ActivityInfo)|false||

<details>
 <summary>
 <code>ActivityInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|name|string|false||
|operator|[Operator](#Operator)|false||
|orgId|string|false||
|string|[//](#//)|false||
|timeRule|[TimeRule](#TimeRule)|false||
|actRule|string|false||
|takePartTotalNum|string|false||
|takePartDayNum|string|false||
|takePartUserNum|string|false||
|takePartDayUserNum|string|false||
|isMarketerExclusive|bool|false||
|userCodeNum|string|false||
|userDayCodeNum|string|false||
|takePartCodeNum|string|false||
|isShare|bool|false||
|shareTitle|string|false||
|shareImg|[ImageData](#ImageData)|false||
|status|string|false||
|addreVerifyType|string|false||
|addreData|[][Addre](#Addre)|false||
|verifyType|string|false||
|verifyMemberLevel|[]string|false||
|verifyApiType|string|false||
|isApiMemberLevelMatchPrizePool|bool|false||
|thirdActId|string|false||
|verifyDesc|string|false||
|verifyNotPass|string|false||
|verifyPass|string|false||
|takePartIsBucklePoint|bool|false||
|pointType|string|false||
|pointNum|string|false||
|isMarketerSelectPrizePool|bool|false||
|recommendImg|[ImageData](#ImageData)|false||
|img|[ImageData](#ImageData)|false||
|buttonImg|[ImageData](#ImageData)|false||
|bgColor|string|false||
|textColor|string|false||
|isShowUserTakePartNum|bool|false||
|userTakePartInitNum|string|false||
|isShowLuckyUser|bool|false||
|luckyUserNum|string|false||
|id|string|false||
|activityStatus|string|false||
|orgData|[OrgData](#OrgData)|false||
|createdAt|string|false||
|pushReceiveDataType|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|luckyBgImg|[ImageData](#ImageData)|false||
|isShowActivityGift|bool|false||
|eventId|string|false||
|isVerifyCard|bool|false||
|verifyCardCode|string|false||
|verifyCardNumber|string|false||
|verifyCardDesc|string|false||
|verifyCardPrice|string|false||
|isStaffPartakeCode|bool|false||
|isStaffRegistration|bool|false||
|staffRegisterConfig|[]string|false||
|isRegisterVerifyPualification|bool|false||
|isVerifyWhiteList|bool|false||
|whiteListList|[][WhiteListData](#WhiteListData)|false||
|isVerifyWhiteListStaff|bool|false||
|isVerifyThirdApiStaff|bool|false||
|gameType|string|false||
|digiccyWalletId|string|false||
|digiccyActName|string|false||
|digiccyActNo|string|false||
|digiccySonActNo|string|false||
|isSendSms|bool|false||
|isDigiccy|bool|false||
|isOpenPrizePushData|bool|false||
|buttonMsg|string|false||
|isRepeatVerify|bool|false||
|birthdayVerify|string|false||
|birthdayVerifyStartMonth|string|false||
|birthdayVerifyEndMonth|string|false||
|birthdayVerifyBeforeDay|string|false||
|birthdayVerifyAfterDay|string|false||
|isGetApiTakePartNum|bool|false||

<details>
 <summary>
 <code>OrgData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|bool|[//](#//)|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

<details>
 <summary>
 <code>WhiteListData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||

</details>

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

<details>
 <summary>
 <code>TimeRule </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|rule|string|false||
|startedAt|string|false||
|endedAt|string|false||
|ruleDay|string|false||
|ruleWeek|string|false||
|ruleTime|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

</details>

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"name":"name",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"orgId":"orgId",
"string":{},
"timeRule":{
"rule":"rule",
"startedAt":"startedAt",
"endedAt":"endedAt",
"ruleDay":"ruleDay",
"ruleWeek":"ruleWeek",
"ruleTime":"ruleTime"
},
"actRule":"actRule",
"takePartTotalNum":"takePartTotalNum",
"takePartDayNum":"takePartDayNum",
"takePartUserNum":"takePartUserNum",
"takePartDayUserNum":"takePartDayUserNum",
"isMarketerExclusive":true,
"userCodeNum":"userCodeNum",
"userDayCodeNum":"userDayCodeNum",
"takePartCodeNum":"takePartCodeNum",
"isShare":true,
"shareTitle":"shareTitle",
"shareImg":{
"id":"id",
"name":"name",
"url":"url"
},
"status":"status",
"addreVerifyType":"addreVerifyType",
"addreData":[
	{
"id":"id",
"fullName":"fullName"
}
],
"verifyType":"verifyType",
"verifyMemberLevel":["verifyMemberLevel"],
"verifyApiType":"verifyApiType",
"isApiMemberLevelMatchPrizePool":true,
"thirdActId":"thirdActId",
"verifyDesc":"verifyDesc",
"verifyNotPass":"verifyNotPass",
"verifyPass":"verifyPass",
"takePartIsBucklePoint":true,
"pointType":"pointType",
"pointNum":"pointNum",
"isMarketerSelectPrizePool":true,
"recommendImg":{
"id":"id",
"name":"name",
"url":"url"
},
"img":{
"id":"id",
"name":"name",
"url":"url"
},
"buttonImg":{
"id":"id",
"name":"name",
"url":"url"
},
"bgColor":"bgColor",
"textColor":"textColor",
"isShowUserTakePartNum":true,
"userTakePartInitNum":"userTakePartInitNum",
"isShowLuckyUser":true,
"luckyUserNum":"luckyUserNum",
"id":"id",
"activityStatus":"activityStatus",
"orgData":{
"id":"id",
"name":"name",
"bool":{}
},
"createdAt":"createdAt",
"pushReceiveDataType":"pushReceiveDataType",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"luckyBgImg":{
"id":"id",
"name":"name",
"url":"url"
},
"isShowActivityGift":true,
"eventId":"eventId",
"isVerifyCard":true,
"verifyCardCode":"verifyCardCode",
"verifyCardNumber":"verifyCardNumber",
"verifyCardDesc":"verifyCardDesc",
"verifyCardPrice":"verifyCardPrice",
"isStaffPartakeCode":true,
"isStaffRegistration":true,
"staffRegisterConfig":["staffRegisterConfig"],
"isRegisterVerifyPualification":true,
"isVerifyWhiteList":true,
"whiteListList":[
	{
"id":"id",
"name":"name"
}
],
"isVerifyWhiteListStaff":true,
"isVerifyThirdApiStaff":true,
"gameType":"gameType",
"digiccyWalletId":"digiccyWalletId",
"digiccyActName":"digiccyActName",
"digiccyActNo":"digiccyActNo",
"digiccySonActNo":"digiccySonActNo",
"isSendSms":true,
"isDigiccy":true,
"isOpenPrizePushData":true,
"buttonMsg":"buttonMsg",
"isRepeatVerify":true,
"birthdayVerify":"birthdayVerify",
"birthdayVerifyStartMonth":"birthdayVerifyStartMonth",
"birthdayVerifyEndMonth":"birthdayVerifyEndMonth",
"birthdayVerifyBeforeDay":"birthdayVerifyBeforeDay",
"birthdayVerifyAfterDay":"birthdayVerifyAfterDay",
"isGetApiTakePartNum":true
}
}
````



<a name="GetActivityList"></a>

## 前端method: InteractluckyactivityGetActivityList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|operator|[Operator](#Operator)|false||
|page|string|false||
|pageSize|string|false||
|name|string|false||
|orgId|string|false||
|id|string|false||
|status|string|false||
|activityStatus|string|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"page":"page",
"pageSize":"pageSize",
"name":"name",
"orgId":"orgId",
"id":"id",
"status":"status",
"activityStatus":"activityStatus"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[GetActivityListData](#GetActivityListData)|false||

<details>
 <summary>
 <code>GetActivityListData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|activity|[][ActivityInfo](#ActivityInfo)|false||
|pageInfo|[PageInfo](#PageInfo)|false||

<details>
 <summary>
 <code>ActivityInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|name|string|false||
|operator|[Operator](#Operator)|false||
|orgId|string|false||
|string|[//](#//)|false||
|timeRule|[TimeRule](#TimeRule)|false||
|actRule|string|false||
|takePartTotalNum|string|false||
|takePartDayNum|string|false||
|takePartUserNum|string|false||
|takePartDayUserNum|string|false||
|isMarketerExclusive|bool|false||
|userCodeNum|string|false||
|userDayCodeNum|string|false||
|takePartCodeNum|string|false||
|isShare|bool|false||
|shareTitle|string|false||
|shareImg|[ImageData](#ImageData)|false||
|status|string|false||
|addreVerifyType|string|false||
|addreData|[][Addre](#Addre)|false||
|verifyType|string|false||
|verifyMemberLevel|[]string|false||
|verifyApiType|string|false||
|isApiMemberLevelMatchPrizePool|bool|false||
|thirdActId|string|false||
|verifyDesc|string|false||
|verifyNotPass|string|false||
|verifyPass|string|false||
|takePartIsBucklePoint|bool|false||
|pointType|string|false||
|pointNum|string|false||
|isMarketerSelectPrizePool|bool|false||
|recommendImg|[ImageData](#ImageData)|false||
|img|[ImageData](#ImageData)|false||
|buttonImg|[ImageData](#ImageData)|false||
|bgColor|string|false||
|textColor|string|false||
|isShowUserTakePartNum|bool|false||
|userTakePartInitNum|string|false||
|isShowLuckyUser|bool|false||
|luckyUserNum|string|false||
|id|string|false||
|activityStatus|string|false||
|orgData|[OrgData](#OrgData)|false||
|createdAt|string|false||
|pushReceiveDataType|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|luckyBgImg|[ImageData](#ImageData)|false||
|isShowActivityGift|bool|false||
|eventId|string|false||
|isVerifyCard|bool|false||
|verifyCardCode|string|false||
|verifyCardNumber|string|false||
|verifyCardDesc|string|false||
|verifyCardPrice|string|false||
|isStaffPartakeCode|bool|false||
|isStaffRegistration|bool|false||
|staffRegisterConfig|[]string|false||
|isRegisterVerifyPualification|bool|false||
|isVerifyWhiteList|bool|false||
|whiteListList|[][WhiteListData](#WhiteListData)|false||
|isVerifyWhiteListStaff|bool|false||
|isVerifyThirdApiStaff|bool|false||
|gameType|string|false||
|digiccyWalletId|string|false||
|digiccyActName|string|false||
|digiccyActNo|string|false||
|digiccySonActNo|string|false||
|isSendSms|bool|false||
|isDigiccy|bool|false||
|isOpenPrizePushData|bool|false||
|buttonMsg|string|false||
|isRepeatVerify|bool|false||
|birthdayVerify|string|false||
|birthdayVerifyStartMonth|string|false||
|birthdayVerifyEndMonth|string|false||
|birthdayVerifyBeforeDay|string|false||
|birthdayVerifyAfterDay|string|false||
|isGetApiTakePartNum|bool|false||

<details>
 <summary>
 <code>TimeRule </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|rule|string|false||
|startedAt|string|false||
|endedAt|string|false||
|ruleDay|string|false||
|ruleWeek|string|false||
|ruleTime|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

<details>
 <summary>
 <code>OrgData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|bool|[//](#//)|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

<details>
 <summary>
 <code>WhiteListData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||

</details>

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

</details>

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

</details>

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"activity":[
	{
"name":"name",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"orgId":"orgId",
"string":{},
"timeRule":{
"rule":"rule",
"startedAt":"startedAt",
"endedAt":"endedAt",
"ruleDay":"ruleDay",
"ruleWeek":"ruleWeek",
"ruleTime":"ruleTime"
},
"actRule":"actRule",
"takePartTotalNum":"takePartTotalNum",
"takePartDayNum":"takePartDayNum",
"takePartUserNum":"takePartUserNum",
"takePartDayUserNum":"takePartDayUserNum",
"isMarketerExclusive":true,
"userCodeNum":"userCodeNum",
"userDayCodeNum":"userDayCodeNum",
"takePartCodeNum":"takePartCodeNum",
"isShare":true,
"shareTitle":"shareTitle",
"shareImg":{
"id":"id",
"name":"name",
"url":"url"
},
"status":"status",
"addreVerifyType":"addreVerifyType",
"addreData":[
	{
"id":"id",
"fullName":"fullName"
}
],
"verifyType":"verifyType",
"verifyMemberLevel":["verifyMemberLevel"],
"verifyApiType":"verifyApiType",
"isApiMemberLevelMatchPrizePool":true,
"thirdActId":"thirdActId",
"verifyDesc":"verifyDesc",
"verifyNotPass":"verifyNotPass",
"verifyPass":"verifyPass",
"takePartIsBucklePoint":true,
"pointType":"pointType",
"pointNum":"pointNum",
"isMarketerSelectPrizePool":true,
"recommendImg":{
"id":"id",
"name":"name",
"url":"url"
},
"img":{
"id":"id",
"name":"name",
"url":"url"
},
"buttonImg":{
"id":"id",
"name":"name",
"url":"url"
},
"bgColor":"bgColor",
"textColor":"textColor",
"isShowUserTakePartNum":true,
"userTakePartInitNum":"userTakePartInitNum",
"isShowLuckyUser":true,
"luckyUserNum":"luckyUserNum",
"id":"id",
"activityStatus":"activityStatus",
"orgData":{
"id":"id",
"name":"name",
"bool":{}
},
"createdAt":"createdAt",
"pushReceiveDataType":"pushReceiveDataType",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"luckyBgImg":{
"id":"id",
"name":"name",
"url":"url"
},
"isShowActivityGift":true,
"eventId":"eventId",
"isVerifyCard":true,
"verifyCardCode":"verifyCardCode",
"verifyCardNumber":"verifyCardNumber",
"verifyCardDesc":"verifyCardDesc",
"verifyCardPrice":"verifyCardPrice",
"isStaffPartakeCode":true,
"isStaffRegistration":true,
"staffRegisterConfig":["staffRegisterConfig"],
"isRegisterVerifyPualification":true,
"isVerifyWhiteList":true,
"whiteListList":[
	{
"id":"id",
"name":"name"
}
],
"isVerifyWhiteListStaff":true,
"isVerifyThirdApiStaff":true,
"gameType":"gameType",
"digiccyWalletId":"digiccyWalletId",
"digiccyActName":"digiccyActName",
"digiccyActNo":"digiccyActNo",
"digiccySonActNo":"digiccySonActNo",
"isSendSms":true,
"isDigiccy":true,
"isOpenPrizePushData":true,
"buttonMsg":"buttonMsg",
"isRepeatVerify":true,
"birthdayVerify":"birthdayVerify",
"birthdayVerifyStartMonth":"birthdayVerifyStartMonth",
"birthdayVerifyEndMonth":"birthdayVerifyEndMonth",
"birthdayVerifyBeforeDay":"birthdayVerifyBeforeDay",
"birthdayVerifyAfterDay":"birthdayVerifyAfterDay",
"isGetApiTakePartNum":true
}
],
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
}
}
}
````



<a name="GetActivityDetails"></a>

## 前端method: InteractluckyactivityGetActivityDetails

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ActivityInfo](#ActivityInfo)|false||

<details>
 <summary>
 <code>ActivityInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|name|string|false||
|operator|[Operator](#Operator)|false||
|orgId|string|false||
|string|[//](#//)|false||
|timeRule|[TimeRule](#TimeRule)|false||
|actRule|string|false||
|takePartTotalNum|string|false||
|takePartDayNum|string|false||
|takePartUserNum|string|false||
|takePartDayUserNum|string|false||
|isMarketerExclusive|bool|false||
|userCodeNum|string|false||
|userDayCodeNum|string|false||
|takePartCodeNum|string|false||
|isShare|bool|false||
|shareTitle|string|false||
|shareImg|[ImageData](#ImageData)|false||
|status|string|false||
|addreVerifyType|string|false||
|addreData|[][Addre](#Addre)|false||
|verifyType|string|false||
|verifyMemberLevel|[]string|false||
|verifyApiType|string|false||
|isApiMemberLevelMatchPrizePool|bool|false||
|thirdActId|string|false||
|verifyDesc|string|false||
|verifyNotPass|string|false||
|verifyPass|string|false||
|takePartIsBucklePoint|bool|false||
|pointType|string|false||
|pointNum|string|false||
|isMarketerSelectPrizePool|bool|false||
|recommendImg|[ImageData](#ImageData)|false||
|img|[ImageData](#ImageData)|false||
|buttonImg|[ImageData](#ImageData)|false||
|bgColor|string|false||
|textColor|string|false||
|isShowUserTakePartNum|bool|false||
|userTakePartInitNum|string|false||
|isShowLuckyUser|bool|false||
|luckyUserNum|string|false||
|id|string|false||
|activityStatus|string|false||
|orgData|[OrgData](#OrgData)|false||
|createdAt|string|false||
|pushReceiveDataType|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|luckyBgImg|[ImageData](#ImageData)|false||
|isShowActivityGift|bool|false||
|eventId|string|false||
|isVerifyCard|bool|false||
|verifyCardCode|string|false||
|verifyCardNumber|string|false||
|verifyCardDesc|string|false||
|verifyCardPrice|string|false||
|isStaffPartakeCode|bool|false||
|isStaffRegistration|bool|false||
|staffRegisterConfig|[]string|false||
|isRegisterVerifyPualification|bool|false||
|isVerifyWhiteList|bool|false||
|whiteListList|[][WhiteListData](#WhiteListData)|false||
|isVerifyWhiteListStaff|bool|false||
|isVerifyThirdApiStaff|bool|false||
|gameType|string|false||
|digiccyWalletId|string|false||
|digiccyActName|string|false||
|digiccyActNo|string|false||
|digiccySonActNo|string|false||
|isSendSms|bool|false||
|isDigiccy|bool|false||
|isOpenPrizePushData|bool|false||
|buttonMsg|string|false||
|isRepeatVerify|bool|false||
|birthdayVerify|string|false||
|birthdayVerifyStartMonth|string|false||
|birthdayVerifyEndMonth|string|false||
|birthdayVerifyBeforeDay|string|false||
|birthdayVerifyAfterDay|string|false||
|isGetApiTakePartNum|bool|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

<details>
 <summary>
 <code>TimeRule </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|rule|string|false||
|startedAt|string|false||
|endedAt|string|false||
|ruleDay|string|false||
|ruleWeek|string|false||
|ruleTime|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

<details>
 <summary>
 <code>OrgData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|bool|[//](#//)|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

<details>
 <summary>
 <code>WhiteListData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||

</details>

</details>

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"name":"name",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"orgId":"orgId",
"string":{},
"timeRule":{
"rule":"rule",
"startedAt":"startedAt",
"endedAt":"endedAt",
"ruleDay":"ruleDay",
"ruleWeek":"ruleWeek",
"ruleTime":"ruleTime"
},
"actRule":"actRule",
"takePartTotalNum":"takePartTotalNum",
"takePartDayNum":"takePartDayNum",
"takePartUserNum":"takePartUserNum",
"takePartDayUserNum":"takePartDayUserNum",
"isMarketerExclusive":true,
"userCodeNum":"userCodeNum",
"userDayCodeNum":"userDayCodeNum",
"takePartCodeNum":"takePartCodeNum",
"isShare":true,
"shareTitle":"shareTitle",
"shareImg":{
"id":"id",
"name":"name",
"url":"url"
},
"status":"status",
"addreVerifyType":"addreVerifyType",
"addreData":[
	{
"id":"id",
"fullName":"fullName"
}
],
"verifyType":"verifyType",
"verifyMemberLevel":["verifyMemberLevel"],
"verifyApiType":"verifyApiType",
"isApiMemberLevelMatchPrizePool":true,
"thirdActId":"thirdActId",
"verifyDesc":"verifyDesc",
"verifyNotPass":"verifyNotPass",
"verifyPass":"verifyPass",
"takePartIsBucklePoint":true,
"pointType":"pointType",
"pointNum":"pointNum",
"isMarketerSelectPrizePool":true,
"recommendImg":{
"id":"id",
"name":"name",
"url":"url"
},
"img":{
"id":"id",
"name":"name",
"url":"url"
},
"buttonImg":{
"id":"id",
"name":"name",
"url":"url"
},
"bgColor":"bgColor",
"textColor":"textColor",
"isShowUserTakePartNum":true,
"userTakePartInitNum":"userTakePartInitNum",
"isShowLuckyUser":true,
"luckyUserNum":"luckyUserNum",
"id":"id",
"activityStatus":"activityStatus",
"orgData":{
"id":"id",
"name":"name",
"bool":{}
},
"createdAt":"createdAt",
"pushReceiveDataType":"pushReceiveDataType",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"luckyBgImg":{
"id":"id",
"name":"name",
"url":"url"
},
"isShowActivityGift":true,
"eventId":"eventId",
"isVerifyCard":true,
"verifyCardCode":"verifyCardCode",
"verifyCardNumber":"verifyCardNumber",
"verifyCardDesc":"verifyCardDesc",
"verifyCardPrice":"verifyCardPrice",
"isStaffPartakeCode":true,
"isStaffRegistration":true,
"staffRegisterConfig":["staffRegisterConfig"],
"isRegisterVerifyPualification":true,
"isVerifyWhiteList":true,
"whiteListList":[
	{
"id":"id",
"name":"name"
}
],
"isVerifyWhiteListStaff":true,
"isVerifyThirdApiStaff":true,
"gameType":"gameType",
"digiccyWalletId":"digiccyWalletId",
"digiccyActName":"digiccyActName",
"digiccyActNo":"digiccyActNo",
"digiccySonActNo":"digiccySonActNo",
"isSendSms":true,
"isDigiccy":true,
"isOpenPrizePushData":true,
"buttonMsg":"buttonMsg",
"isRepeatVerify":true,
"birthdayVerify":"birthdayVerify",
"birthdayVerifyStartMonth":"birthdayVerifyStartMonth",
"birthdayVerifyEndMonth":"birthdayVerifyEndMonth",
"birthdayVerifyBeforeDay":"birthdayVerifyBeforeDay",
"birthdayVerifyAfterDay":"birthdayVerifyAfterDay",
"isGetApiTakePartNum":true
}
}
````



<a name="GetActivityStatus"></a>

## 前端method: InteractluckyactivityGetActivityStatus

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|activityId|string|false||
|channelId|string|false||

### 请求例子:
````
{
"activityId":"activityId",
"channelId":"channelId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ActivityStatusData](#ActivityStatusData)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ActivityStatusData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|status|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"status":"status"
}
}
````



<a name="DelActivity"></a>

## 前端method: InteractluckyactivityDelActivity

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="CreatePrizePool"></a>

## 前端method: InteractluckyactivityCreatePrizePool

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|name|string|false||
|type|string|false||
|memberLevel|[]string|false||
|thirdPrizePoolId|string|false||
|isDefault|bool|false||
|specialInterval|string|false||
|prizeInfo|[][PrizeInfoReq](#PrizeInfoReq)|false||
|operator|[Operator](#Operator)|false||
|thirdOrgNumber|string|false||
|isSpecialPrize|bool|false||
|luckyTypeBgId|string|false||
|eventId|string|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

<details>
 <summary>
 <code>PrizeInfoReq </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|type|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|money|string|false||
|totalNum|string|false||
|isSpecial|bool|false||
|imgId|string|false||
|giftId|string|false||
|sort|string|false||
|desc|string|false||
|eId|string|false||
|limitAreaId|string|false||
|isSinceExtractOrgNum|bool|false||
|costPrice|string|false||
|pointNum|string|false||
|couponActId|string|false||
|isDelayIssue|bool|false||
|eventId|string|false||
|delayDay|string|false||
|delayIssueMsg|string|false||
|couponBagData|[][CouponBagData](#CouponBagData)|false||
|appointmentStartAt|string|false||
|appointmentEndAt|string|false||
|goodsId|string|false||
|delayDayNum|string|false||
|isAutomaticAppointment|bool|false||
|automaticAppointmentMsg|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|awardNum|string|false||
|awardUnit|string|false||
|jumpConfig|string|false||
|payType|string|false||
|detailDesc|string|false||
|status|string|false||
|isDefaultNetAppointment|bool|false||
|isLuckDrawOneFrequency|bool|false||
|isOpenSpecialRule|string|false||

<details>
 <summary>
 <code>CouponBagData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|stockId|string|false||
|num|string|false||
|couponActId|string|false||
|id|string|false||
|name|string|false||
|amount|string|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

</details>

### 请求例子:
````
{
"actId":"actId",
"name":"name",
"type":"type",
"memberLevel":["memberLevel"],
"thirdPrizePoolId":"thirdPrizePoolId",
"isDefault":true,
"specialInterval":"specialInterval",
"prizeInfo":[
	{
"id":"id",
"name":"name",
"type":"type",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"money":"money",
"totalNum":"totalNum",
"isSpecial":true,
"imgId":"imgId",
"giftId":"giftId",
"sort":"sort",
"desc":"desc",
"eId":"eId",
"limitAreaId":"limitAreaId",
"isSinceExtractOrgNum":true,
"costPrice":"costPrice",
"pointNum":"pointNum",
"couponActId":"couponActId",
"isDelayIssue":true,
"eventId":"eventId",
"delayDay":"delayDay",
"delayIssueMsg":"delayIssueMsg",
"couponBagData":[
	{
"stockId":"stockId",
"num":"num",
"couponActId":"couponActId",
"id":"id",
"name":"name",
"amount":"amount"
}
],
"appointmentStartAt":"appointmentStartAt",
"appointmentEndAt":"appointmentEndAt",
"goodsId":"goodsId",
"delayDayNum":"delayDayNum",
"isAutomaticAppointment":true,
"automaticAppointmentMsg":"automaticAppointmentMsg",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"awardNum":"awardNum",
"awardUnit":"awardUnit",
"jumpConfig":"jumpConfig",
"payType":"payType",
"detailDesc":"detailDesc",
"status":"status",
"isDefaultNetAppointment":true,
"isLuckDrawOneFrequency":true,
"isOpenSpecialRule":"isOpenSpecialRule"
}
],
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"thirdOrgNumber":"thirdOrgNumber",
"isSpecialPrize":true,
"luckyTypeBgId":"luckyTypeBgId",
"eventId":"eventId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizePoolInfo](#PrizePoolInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizePoolInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|name|string|false||
|type|string|false||
|memberLevel|[]string|false||
|thirdPrizePoolId|string|false||
|isDefault|bool|false||
|specialInterval|string|false||
|prizeInfo|[][PrizeInfo](#PrizeInfo)|false||
|thirdOrgNumber|string|false||
|isSpecialPrize|bool|false||
|luckyTypeBg|[ImageData](#ImageData)|false||
|eventId|string|false||

<details>
 <summary>
 <code>PrizeInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|name|string|false||
|type|string|false||
|money|string|false||
|totalNum|string|false||
|isSpecial|bool|false||
|imgId|string|false||
|giftId|string|false||
|sort|string|false||
|desc|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|imgData|[ImageData](#ImageData)|false||
|eId|string|false||
|count|string|false||
|limitArea|[Addre](#Addre)|false||
|limitProvince|string|false||
|limitCity|string|false||
|limitDistrict|string|false||
|isSinceExtractOrgNum|bool|false||
|prizeId|string|false||
|expendNum|string|false||
|costPrice|string|false||
|pointNum|string|false||
|couponActId|string|false||
|isDelayIssue|bool|false||
|eventId|string|false||
|delayDay|string|false||
|delayIssueMsg|string|false||
|couponBagData|[][CouponBagData](#CouponBagData)|false||
|residueNum|string|false||
|appointmentStartAt|string|false||
|appointmentEndAt|string|false||
|goodsId|string|false||
|delayDayNum|string|false||
|isAutomaticAppointment|bool|false||
|automaticAppointmentMsg|string|false||
|autoPreOrgId|string|false||
|dayCount|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|awardNum|string|false||
|awardUnit|string|false||
|jumpConfig|string|false||
|surplusNum|string|false||
|payType|string|false||
|detailDesc|string|false||
|status|string|false||
|isDefaultNetAppointment|bool|false||
|defaultPreOrgId|string|false||
|isLuckDrawOneFrequency|bool|false||
|isOpenSpecialRule|string|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

<details>
 <summary>
 <code>CouponBagData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|stockId|string|false||
|num|string|false||
|couponActId|string|false||
|id|string|false||
|name|string|false||
|amount|string|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"actId":"actId",
"name":"name",
"type":"type",
"memberLevel":["memberLevel"],
"thirdPrizePoolId":"thirdPrizePoolId",
"isDefault":true,
"specialInterval":"specialInterval",
"prizeInfo":[
	{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"name":"name",
"type":"type",
"money":"money",
"totalNum":"totalNum",
"isSpecial":true,
"imgId":"imgId",
"giftId":"giftId",
"sort":"sort",
"desc":"desc",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"eId":"eId",
"count":"count",
"limitArea":{
"id":"id",
"fullName":"fullName"
},
"limitProvince":"limitProvince",
"limitCity":"limitCity",
"limitDistrict":"limitDistrict",
"isSinceExtractOrgNum":true,
"prizeId":"prizeId",
"expendNum":"expendNum",
"costPrice":"costPrice",
"pointNum":"pointNum",
"couponActId":"couponActId",
"isDelayIssue":true,
"eventId":"eventId",
"delayDay":"delayDay",
"delayIssueMsg":"delayIssueMsg",
"couponBagData":[
	{
"stockId":"stockId",
"num":"num",
"couponActId":"couponActId",
"id":"id",
"name":"name",
"amount":"amount"
}
],
"residueNum":"residueNum",
"appointmentStartAt":"appointmentStartAt",
"appointmentEndAt":"appointmentEndAt",
"goodsId":"goodsId",
"delayDayNum":"delayDayNum",
"isAutomaticAppointment":true,
"automaticAppointmentMsg":"automaticAppointmentMsg",
"autoPreOrgId":"autoPreOrgId",
"dayCount":"dayCount",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"awardNum":"awardNum",
"awardUnit":"awardUnit",
"jumpConfig":"jumpConfig",
"surplusNum":"surplusNum",
"payType":"payType",
"detailDesc":"detailDesc",
"status":"status",
"isDefaultNetAppointment":true,
"defaultPreOrgId":"defaultPreOrgId",
"isLuckDrawOneFrequency":true,
"isOpenSpecialRule":"isOpenSpecialRule"
}
],
"thirdOrgNumber":"thirdOrgNumber",
"isSpecialPrize":true,
"luckyTypeBg":{
"id":"id",
"name":"name",
"url":"url"
},
"eventId":"eventId"
}
}
````



<a name="EditPrizePool"></a>

## 前端method: InteractluckyactivityEditPrizePool

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|name|string|false||
|type|string|false||
|memberLevel|[]string|false||
|thirdPrizePoolId|string|false||
|isDefault|bool|false||
|specialInterval|string|false||
|prizeInfo|[][PrizeInfoReq](#PrizeInfoReq)|false||
|operator|[Operator](#Operator)|false||
|thirdOrgNumber|string|false||
|isSpecialPrize|bool|false||
|luckyTypeBgId|string|false||
|eventId|string|false||

<details>
 <summary>
 <code>PrizeInfoReq </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|type|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|money|string|false||
|totalNum|string|false||
|isSpecial|bool|false||
|imgId|string|false||
|giftId|string|false||
|sort|string|false||
|desc|string|false||
|eId|string|false||
|limitAreaId|string|false||
|isSinceExtractOrgNum|bool|false||
|costPrice|string|false||
|pointNum|string|false||
|couponActId|string|false||
|isDelayIssue|bool|false||
|eventId|string|false||
|delayDay|string|false||
|delayIssueMsg|string|false||
|couponBagData|[][CouponBagData](#CouponBagData)|false||
|appointmentStartAt|string|false||
|appointmentEndAt|string|false||
|goodsId|string|false||
|delayDayNum|string|false||
|isAutomaticAppointment|bool|false||
|automaticAppointmentMsg|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|awardNum|string|false||
|awardUnit|string|false||
|jumpConfig|string|false||
|payType|string|false||
|detailDesc|string|false||
|status|string|false||
|isDefaultNetAppointment|bool|false||
|isLuckDrawOneFrequency|bool|false||
|isOpenSpecialRule|string|false||

<details>
 <summary>
 <code>CouponBagData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|stockId|string|false||
|num|string|false||
|couponActId|string|false||
|id|string|false||
|name|string|false||
|amount|string|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

</details>

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"actId":"actId",
"name":"name",
"type":"type",
"memberLevel":["memberLevel"],
"thirdPrizePoolId":"thirdPrizePoolId",
"isDefault":true,
"specialInterval":"specialInterval",
"prizeInfo":[
	{
"id":"id",
"name":"name",
"type":"type",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"money":"money",
"totalNum":"totalNum",
"isSpecial":true,
"imgId":"imgId",
"giftId":"giftId",
"sort":"sort",
"desc":"desc",
"eId":"eId",
"limitAreaId":"limitAreaId",
"isSinceExtractOrgNum":true,
"costPrice":"costPrice",
"pointNum":"pointNum",
"couponActId":"couponActId",
"isDelayIssue":true,
"eventId":"eventId",
"delayDay":"delayDay",
"delayIssueMsg":"delayIssueMsg",
"couponBagData":[
	{
"stockId":"stockId",
"num":"num",
"couponActId":"couponActId",
"id":"id",
"name":"name",
"amount":"amount"
}
],
"appointmentStartAt":"appointmentStartAt",
"appointmentEndAt":"appointmentEndAt",
"goodsId":"goodsId",
"delayDayNum":"delayDayNum",
"isAutomaticAppointment":true,
"automaticAppointmentMsg":"automaticAppointmentMsg",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"awardNum":"awardNum",
"awardUnit":"awardUnit",
"jumpConfig":"jumpConfig",
"payType":"payType",
"detailDesc":"detailDesc",
"status":"status",
"isDefaultNetAppointment":true,
"isLuckDrawOneFrequency":true,
"isOpenSpecialRule":"isOpenSpecialRule"
}
],
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"thirdOrgNumber":"thirdOrgNumber",
"isSpecialPrize":true,
"luckyTypeBgId":"luckyTypeBgId",
"eventId":"eventId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizePoolInfo](#PrizePoolInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizePoolInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|name|string|false||
|type|string|false||
|memberLevel|[]string|false||
|thirdPrizePoolId|string|false||
|isDefault|bool|false||
|specialInterval|string|false||
|prizeInfo|[][PrizeInfo](#PrizeInfo)|false||
|thirdOrgNumber|string|false||
|isSpecialPrize|bool|false||
|luckyTypeBg|[ImageData](#ImageData)|false||
|eventId|string|false||

<details>
 <summary>
 <code>PrizeInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|name|string|false||
|type|string|false||
|money|string|false||
|totalNum|string|false||
|isSpecial|bool|false||
|imgId|string|false||
|giftId|string|false||
|sort|string|false||
|desc|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|imgData|[ImageData](#ImageData)|false||
|eId|string|false||
|count|string|false||
|limitArea|[Addre](#Addre)|false||
|limitProvince|string|false||
|limitCity|string|false||
|limitDistrict|string|false||
|isSinceExtractOrgNum|bool|false||
|prizeId|string|false||
|expendNum|string|false||
|costPrice|string|false||
|pointNum|string|false||
|couponActId|string|false||
|isDelayIssue|bool|false||
|eventId|string|false||
|delayDay|string|false||
|delayIssueMsg|string|false||
|couponBagData|[][CouponBagData](#CouponBagData)|false||
|residueNum|string|false||
|appointmentStartAt|string|false||
|appointmentEndAt|string|false||
|goodsId|string|false||
|delayDayNum|string|false||
|isAutomaticAppointment|bool|false||
|automaticAppointmentMsg|string|false||
|autoPreOrgId|string|false||
|dayCount|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|awardNum|string|false||
|awardUnit|string|false||
|jumpConfig|string|false||
|surplusNum|string|false||
|payType|string|false||
|detailDesc|string|false||
|status|string|false||
|isDefaultNetAppointment|bool|false||
|defaultPreOrgId|string|false||
|isLuckDrawOneFrequency|bool|false||
|isOpenSpecialRule|string|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

<details>
 <summary>
 <code>CouponBagData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|stockId|string|false||
|num|string|false||
|couponActId|string|false||
|id|string|false||
|name|string|false||
|amount|string|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"actId":"actId",
"name":"name",
"type":"type",
"memberLevel":["memberLevel"],
"thirdPrizePoolId":"thirdPrizePoolId",
"isDefault":true,
"specialInterval":"specialInterval",
"prizeInfo":[
	{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"name":"name",
"type":"type",
"money":"money",
"totalNum":"totalNum",
"isSpecial":true,
"imgId":"imgId",
"giftId":"giftId",
"sort":"sort",
"desc":"desc",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"eId":"eId",
"count":"count",
"limitArea":{
"id":"id",
"fullName":"fullName"
},
"limitProvince":"limitProvince",
"limitCity":"limitCity",
"limitDistrict":"limitDistrict",
"isSinceExtractOrgNum":true,
"prizeId":"prizeId",
"expendNum":"expendNum",
"costPrice":"costPrice",
"pointNum":"pointNum",
"couponActId":"couponActId",
"isDelayIssue":true,
"eventId":"eventId",
"delayDay":"delayDay",
"delayIssueMsg":"delayIssueMsg",
"couponBagData":[
	{
"stockId":"stockId",
"num":"num",
"couponActId":"couponActId",
"id":"id",
"name":"name",
"amount":"amount"
}
],
"residueNum":"residueNum",
"appointmentStartAt":"appointmentStartAt",
"appointmentEndAt":"appointmentEndAt",
"goodsId":"goodsId",
"delayDayNum":"delayDayNum",
"isAutomaticAppointment":true,
"automaticAppointmentMsg":"automaticAppointmentMsg",
"autoPreOrgId":"autoPreOrgId",
"dayCount":"dayCount",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"awardNum":"awardNum",
"awardUnit":"awardUnit",
"jumpConfig":"jumpConfig",
"surplusNum":"surplusNum",
"payType":"payType",
"detailDesc":"detailDesc",
"status":"status",
"isDefaultNetAppointment":true,
"defaultPreOrgId":"defaultPreOrgId",
"isLuckDrawOneFrequency":true,
"isOpenSpecialRule":"isOpenSpecialRule"
}
],
"thirdOrgNumber":"thirdOrgNumber",
"isSpecialPrize":true,
"luckyTypeBg":{
"id":"id",
"name":"name",
"url":"url"
},
"eventId":"eventId"
}
}
````



<a name="GetPrizePoolList"></a>

## 前端method: InteractluckyactivityGetPrizePoolList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"actId":"actId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[GetPrizePoolListData](#GetPrizePoolListData)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>GetPrizePoolListData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|exportInfo|[ExportInfo](#ExportInfo)|false||
|poolInfo|[][PoolInfo](#PoolInfo)|false||

<details>
 <summary>
 <code>ExportInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|type|string|false||
|searchData|string|false||
|additionalData|string|false||

</details>

<details>
 <summary>
 <code>PoolInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|name|string|false||
|type|string|false||
|memberLevel|[]string|false||
|thirdPrizePoolId|string|false||
|isDefault|bool|false||
|specialInterval|string|false||
|prizeNum|string|false||
|createdAt|string|false||
|adminName|string|false||
|isVerify|bool|false||
|thirdOrgNumber|string|false||
|poolBgImageData|[ImageData](#ImageData)|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"exportInfo":{
"type":"type",
"searchData":"searchData",
"additionalData":"additionalData"
},
"poolInfo":[
	{
"id":"id",
"actId":"actId",
"name":"name",
"type":"type",
"memberLevel":["memberLevel"],
"thirdPrizePoolId":"thirdPrizePoolId",
"isDefault":true,
"specialInterval":"specialInterval",
"prizeNum":"prizeNum",
"createdAt":"createdAt",
"adminName":"adminName",
"isVerify":true,
"thirdOrgNumber":"thirdOrgNumber",
"poolBgImageData":{
"id":"id",
"name":"name",
"url":"url"
}
}
]
}
}
````



<a name="GetPrizePoolDetails"></a>

## 前端method: InteractluckyactivityGetPrizePoolDetails

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"actId":"actId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizePoolInfo](#PrizePoolInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizePoolInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|name|string|false||
|type|string|false||
|memberLevel|[]string|false||
|thirdPrizePoolId|string|false||
|isDefault|bool|false||
|specialInterval|string|false||
|prizeInfo|[][PrizeInfo](#PrizeInfo)|false||
|thirdOrgNumber|string|false||
|isSpecialPrize|bool|false||
|luckyTypeBg|[ImageData](#ImageData)|false||
|eventId|string|false||

<details>
 <summary>
 <code>PrizeInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|name|string|false||
|type|string|false||
|money|string|false||
|totalNum|string|false||
|isSpecial|bool|false||
|imgId|string|false||
|giftId|string|false||
|sort|string|false||
|desc|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|imgData|[ImageData](#ImageData)|false||
|eId|string|false||
|count|string|false||
|limitArea|[Addre](#Addre)|false||
|limitProvince|string|false||
|limitCity|string|false||
|limitDistrict|string|false||
|isSinceExtractOrgNum|bool|false||
|prizeId|string|false||
|expendNum|string|false||
|costPrice|string|false||
|pointNum|string|false||
|couponActId|string|false||
|isDelayIssue|bool|false||
|eventId|string|false||
|delayDay|string|false||
|delayIssueMsg|string|false||
|couponBagData|[][CouponBagData](#CouponBagData)|false||
|residueNum|string|false||
|appointmentStartAt|string|false||
|appointmentEndAt|string|false||
|goodsId|string|false||
|delayDayNum|string|false||
|isAutomaticAppointment|bool|false||
|automaticAppointmentMsg|string|false||
|autoPreOrgId|string|false||
|dayCount|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|awardNum|string|false||
|awardUnit|string|false||
|jumpConfig|string|false||
|surplusNum|string|false||
|payType|string|false||
|detailDesc|string|false||
|status|string|false||
|isDefaultNetAppointment|bool|false||
|defaultPreOrgId|string|false||
|isLuckDrawOneFrequency|bool|false||
|isOpenSpecialRule|string|false||

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

<details>
 <summary>
 <code>CouponBagData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|stockId|string|false||
|num|string|false||
|couponActId|string|false||
|id|string|false||
|name|string|false||
|amount|string|false||

</details>

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"actId":"actId",
"name":"name",
"type":"type",
"memberLevel":["memberLevel"],
"thirdPrizePoolId":"thirdPrizePoolId",
"isDefault":true,
"specialInterval":"specialInterval",
"prizeInfo":[
	{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"name":"name",
"type":"type",
"money":"money",
"totalNum":"totalNum",
"isSpecial":true,
"imgId":"imgId",
"giftId":"giftId",
"sort":"sort",
"desc":"desc",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"eId":"eId",
"count":"count",
"limitArea":{
"id":"id",
"fullName":"fullName"
},
"limitProvince":"limitProvince",
"limitCity":"limitCity",
"limitDistrict":"limitDistrict",
"isSinceExtractOrgNum":true,
"prizeId":"prizeId",
"expendNum":"expendNum",
"costPrice":"costPrice",
"pointNum":"pointNum",
"couponActId":"couponActId",
"isDelayIssue":true,
"eventId":"eventId",
"delayDay":"delayDay",
"delayIssueMsg":"delayIssueMsg",
"couponBagData":[
	{
"stockId":"stockId",
"num":"num",
"couponActId":"couponActId",
"id":"id",
"name":"name",
"amount":"amount"
}
],
"residueNum":"residueNum",
"appointmentStartAt":"appointmentStartAt",
"appointmentEndAt":"appointmentEndAt",
"goodsId":"goodsId",
"delayDayNum":"delayDayNum",
"isAutomaticAppointment":true,
"automaticAppointmentMsg":"automaticAppointmentMsg",
"autoPreOrgId":"autoPreOrgId",
"dayCount":"dayCount",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"awardNum":"awardNum",
"awardUnit":"awardUnit",
"jumpConfig":"jumpConfig",
"surplusNum":"surplusNum",
"payType":"payType",
"detailDesc":"detailDesc",
"status":"status",
"isDefaultNetAppointment":true,
"defaultPreOrgId":"defaultPreOrgId",
"isLuckDrawOneFrequency":true,
"isOpenSpecialRule":"isOpenSpecialRule"
}
],
"thirdOrgNumber":"thirdOrgNumber",
"isSpecialPrize":true,
"luckyTypeBg":{
"id":"id",
"name":"name",
"url":"url"
},
"eventId":"eventId"
}
}
````



<a name="AdjustPrizeTotalNum"></a>

## 前端method: InteractluckyactivityAdjustPrizeTotalNum

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|num|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"num":"num",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="DelPrizePool"></a>

## 前端method: InteractluckyactivityDelPrizePool

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"actId":"actId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="DelPrize"></a>

## 前端method: InteractluckyactivityDelPrize

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="GetPrizeOrgList"></a>

## 前端method: InteractluckyactivityGetPrizeOrgList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|prizeId|string|false||
|actId|string|false||
|orgName|string|false||
|page|string|false||
|pageSize|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"prizeId":"prizeId",
"actId":"actId",
"orgName":"orgName",
"page":"page",
"pageSize":"pageSize",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[GetPrizeOrgListDetails](#GetPrizeOrgListDetails)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>GetPrizeOrgListDetails </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|pageInfo|[PageInfo](#PageInfo)|false||
|importInfo|[ImportInfo](#ImportInfo)|false||
|prizeOrgList|[][PrizeOrgDetails](#PrizeOrgDetails)|false||

<details>
 <summary>
 <code>ImportInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|type|string|false||
|additionalData|string|false||

</details>

<details>
 <summary>
 <code>PrizeOrgDetails </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|orgId|string|false||
|orgName|string|false||
|orgNo|string|false||
|num|string|false||
|totalNum|string|false||
|createdAt|string|false||

</details>

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"importInfo":{
"type":"type",
"additionalData":"additionalData"
},
"prizeOrgList":[
	{
"id":"id",
"orgId":"orgId",
"orgName":"orgName",
"orgNo":"orgNo",
"num":"num",
"totalNum":"totalNum",
"createdAt":"createdAt"
}
]
}
}
````



<a name="EditPrizeOrgNum"></a>

## 前端method: InteractluckyactivityEditPrizeOrgNum

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|num|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"num":"num",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizeOrgDetails](#PrizeOrgDetails)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizeOrgDetails </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|orgId|string|false||
|orgName|string|false||
|orgNo|string|false||
|num|string|false||
|totalNum|string|false||
|createdAt|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"orgId":"orgId",
"orgName":"orgName",
"orgNo":"orgNo",
"num":"num",
"totalNum":"totalNum",
"createdAt":"createdAt"
}
}
````



<a name="DeleteReceiveRecord"></a>

## 前端method: InteractluckyactivityDeleteReceiveRecord

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="CreatePrizeWhitelist"></a>

## 前端method: InteractluckyactivityCreatePrizeWhitelist

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|operator|[Operator](#Operator)|false||
|actId|string|false||
|poolId|string|false||
|prizeId|string|false||
|mobile|string|false||
|totalNum|string|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"actId":"actId",
"poolId":"poolId",
"prizeId":"prizeId",
"mobile":"mobile",
"totalNum":"totalNum"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizeWhitelistData](#PrizeWhitelistData)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizeWhitelistData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|prizeId|string|false||
|mobile|string|false||
|totalNum|string|false||
|surplusNum|string|false||
|createdAt|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"prizeId":"prizeId",
"mobile":"mobile",
"totalNum":"totalNum",
"surplusNum":"surplusNum",
"createdAt":"createdAt"
}
}
````



<a name="DelPrizeWhitelist"></a>

## 前端method: InteractluckyactivityDelPrizeWhitelist

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|operator|[Operator](#Operator)|false||
|id|string|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"id":"id"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="GetPrizeWhitelistList"></a>

## 前端method: InteractluckyactivityGetPrizeWhitelistList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|prizeId|string|false||
|mobile|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"page":"page",
"pageSize":"pageSize",
"prizeId":"prizeId",
"mobile":"mobile",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizeWhitelistListData](#PrizeWhitelistListData)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizeWhitelistListData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|list|[][PrizeWhitelistData](#PrizeWhitelistData)|false||
|pageInfo|[PageInfo](#PageInfo)|false||
|importInfo|[ImportInfo](#ImportInfo)|false||

<details>
 <summary>
 <code>PrizeWhitelistData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|prizeId|string|false||
|mobile|string|false||
|totalNum|string|false||
|surplusNum|string|false||
|createdAt|string|false||

</details>

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

<details>
 <summary>
 <code>ImportInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|type|string|false||
|additionalData|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"list":[
	{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"prizeId":"prizeId",
"mobile":"mobile",
"totalNum":"totalNum",
"surplusNum":"surplusNum",
"createdAt":"createdAt"
}
],
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"importInfo":{
"type":"type",
"additionalData":"additionalData"
}
}
}
````



<a name="AdjustPrizeWhitelistNum"></a>

## 前端method: InteractluckyactivityAdjustPrizeWhitelistNum

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|operator|[Operator](#Operator)|false||
|num|string|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"num":"num"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizeWhitelistData](#PrizeWhitelistData)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizeWhitelistData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|prizeId|string|false||
|mobile|string|false||
|totalNum|string|false||
|surplusNum|string|false||
|createdAt|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"prizeId":"prizeId",
"mobile":"mobile",
"totalNum":"totalNum",
"surplusNum":"surplusNum",
"createdAt":"createdAt"
}
}
````



<a name="GetReceiveList"></a>

## 前端method: InteractluckyactivityGetReceiveList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|operator|[Operator](#Operator)|false||
|mobile|string|false||
|name|string|false||
|idCard|string|false||
|prizeName|string|false||
|prizeType|string|false||
|status|string|false||
|orgId|string|false||
|page|string|false||
|pageSize|string|false||
|actId|string|false||
|receiveStartAt|string|false||
|receiveEndAt|string|false||
|orgNo|string|false||
|guihuOrgId|string|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"mobile":"mobile",
"name":"name",
"idCard":"idCard",
"prizeName":"prizeName",
"prizeType":"prizeType",
"status":"status",
"orgId":"orgId",
"page":"page",
"pageSize":"pageSize",
"actId":"actId",
"receiveStartAt":"receiveStartAt",
"receiveEndAt":"receiveEndAt",
"orgNo":"orgNo",
"guihuOrgId":"guihuOrgId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[GetReceiveList](#GetReceiveList)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>GetReceiveList </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveList|[][ReceiveInfo](#ReceiveInfo)|false||
|importInfo|[ImportInfo](#ImportInfo)|false||
|pageInfo|[PageInfo](#PageInfo)|false||
|exportInfo|[ExportInfo](#ExportInfo)|false||
|deliveryExportInfo|[ExportInfo](#ExportInfo)|false||
|delDataImportInfo|[ImportInfo](#ImportInfo)|false||

<details>
 <summary>
 <code>ExportInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|type|string|false||
|searchData|string|false||
|additionalData|string|false||

</details>

<details>
 <summary>
 <code>ReceiveInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|orderNo|string|false||
|mobile|string|false||
|name|string|false||
|idCard|string|false||
|receivedAt|string|false||
|prizeName|string|false||
|prizeType|string|false||
|status|string|false||
|marketerName|string|false||
|orgData|[OrgData](#OrgData)|false||
|verifyStatus|string|false||
|receiveMobile|string|false||
|receiveAddress|string|false||
|receiveProvince|string|false||
|receiveCity|string|false||
|receiveDistrict|string|false||
|receiveName|string|false||
|orgId|string|false||
|orgName|string|false||
|appointmentAt|string|false||
|bankOrgNo|string|false||
|gdCrmOrgNo|string|false||
|guihuOrgName|string|false||
|errMsg|string|false||
|payType|string|false||
|payMoney|string|false||

<details>
 <summary>
 <code>OrgData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|bool|[//](#//)|false||

</details>

</details>

<details>
 <summary>
 <code>ImportInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|type|string|false||
|additionalData|string|false||

</details>

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"receiveList":[
	{
"id":"id",
"orderNo":"orderNo",
"mobile":"mobile",
"name":"name",
"idCard":"idCard",
"receivedAt":"receivedAt",
"prizeName":"prizeName",
"prizeType":"prizeType",
"status":"status",
"marketerName":"marketerName",
"orgData":{
"id":"id",
"name":"name",
"bool":{}
},
"verifyStatus":"verifyStatus",
"receiveMobile":"receiveMobile",
"receiveAddress":"receiveAddress",
"receiveProvince":"receiveProvince",
"receiveCity":"receiveCity",
"receiveDistrict":"receiveDistrict",
"receiveName":"receiveName",
"orgId":"orgId",
"orgName":"orgName",
"appointmentAt":"appointmentAt",
"bankOrgNo":"bankOrgNo",
"gdCrmOrgNo":"gdCrmOrgNo",
"guihuOrgName":"guihuOrgName",
"errMsg":"errMsg",
"payType":"payType",
"payMoney":"payMoney"
}
],
"importInfo":{
"type":"type",
"additionalData":"additionalData"
},
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"exportInfo":{
"type":"type",
"searchData":"searchData",
"additionalData":"additionalData"
},
"deliveryExportInfo":{
"type":"type",
"searchData":"searchData",
"additionalData":"additionalData"
},
"delDataImportInfo":{
"type":"type",
"additionalData":"additionalData"
}
}
}
````



<a name="Replacement"></a>

## 前端method: InteractluckyactivityReplacement

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"actId":"actId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="AllReplacement"></a>

## 前端method: InteractluckyactivityAllReplacement

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"actId":"actId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="EditAppointmentOrg"></a>

## 前端method: InteractluckyactivityEditAppointmentOrg

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveId|string|false||
|orgId|string|false||
|appointmentAt|string|false||

### 请求例子:
````
{
"receiveId":"receiveId",
"orgId":"orgId",
"appointmentAt":"appointmentAt"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="BatchEditAppointmentOrg"></a>

## 前端method: InteractluckyactivityBatchEditAppointmentOrg

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|operator|[Operator](#Operator)|false||
|appointmentAt|string|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"actId":"actId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"appointmentAt":"appointmentAt"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="ParticipateCodeList"></a>

## 前端method: InteractluckyactivityParticipateCodeList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|participateCode|string|false||
|page|string|false||
|pageSize|string|false||
|operator|[Operator](#Operator)|false||
|actId|string|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"participateCode":"participateCode",
"page":"page",
"pageSize":"pageSize",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"actId":"actId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ParticipateCodeList](#ParticipateCodeList)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ParticipateCodeList </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|participateCode|[][ParticipateCodeDetails](#ParticipateCodeDetails)|false||
|pageInfo|[PageInfo](#PageInfo)|false||
|importParticipateCode|[ImportInfo](#ImportInfo)|false||

<details>
 <summary>
 <code>ParticipateCodeDetails </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|code|string|false||
|operatorName|string|false||
|createAt|string|false||

</details>

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

<details>
 <summary>
 <code>ImportInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|type|string|false||
|additionalData|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"participateCode":[
	{
"id":"id",
"actId":"actId",
"code":"code",
"operatorName":"operatorName",
"createAt":"createAt"
}
],
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"importParticipateCode":{
"type":"type",
"additionalData":"additionalData"
}
}
}
````



<a name="DelParticipateCodeList"></a>

## 前端method: InteractluckyactivityDelParticipateCodeList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"actId":"actId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="MobileWhiteList"></a>

## 前端method: InteractluckyactivityMobileWhiteList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|mobile|string|false||
|page|string|false||
|pageSize|string|false||
|operator|[Operator](#Operator)|false||
|actId|string|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"mobile":"mobile",
"page":"page",
"pageSize":"pageSize",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"actId":"actId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[MobileWhiteList](#MobileWhiteList)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>MobileWhiteList </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|mobileWhiteList|[][MobileWhiteListDetails](#MobileWhiteListDetails)|false||
|pageInfo|[PageInfo](#PageInfo)|false||
|importMobileWhiteList|[ImportInfo](#ImportInfo)|false||
|exportMobileWhiteList|[ExportInfo](#ExportInfo)|false||

<details>
 <summary>
 <code>ImportInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|type|string|false||
|additionalData|string|false||

</details>

<details>
 <summary>
 <code>ExportInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|type|string|false||
|searchData|string|false||
|additionalData|string|false||

</details>

<details>
 <summary>
 <code>MobileWhiteListDetails </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|mobile|string|false||
|operatorName|string|false||
|createAt|string|false||

</details>

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"mobileWhiteList":[
	{
"id":"id",
"actId":"actId",
"mobile":"mobile",
"operatorName":"operatorName",
"createAt":"createAt"
}
],
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"importMobileWhiteList":{
"type":"type",
"additionalData":"additionalData"
},
"exportMobileWhiteList":{
"type":"type",
"searchData":"searchData",
"additionalData":"additionalData"
}
}
}
````



<a name="DelMobileWhiteList"></a>

## 前端method: InteractluckyactivityDelMobileWhiteList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"actId":"actId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="RepeatMobileWhiteList"></a>

## 前端method: InteractluckyactivityRepeatMobileWhiteList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|mobile|string|false||
|page|string|false||
|pageSize|string|false||
|operator|[Operator](#Operator)|false||
|actId|string|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"mobile":"mobile",
"page":"page",
"pageSize":"pageSize",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"actId":"actId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[RepeatMobileWhiteListData](#RepeatMobileWhiteListData)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>RepeatMobileWhiteListData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|list|[][RepeatMobileWhite](#RepeatMobileWhite)|false||
|pageInfo|[PageInfo](#PageInfo)|false||
|importRepeatMobileWhite|[ImportInfo](#ImportInfo)|false||
|exportRepeatMobileWhite|[ExportInfo](#ExportInfo)|false||

<details>
 <summary>
 <code>RepeatMobileWhite </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|mobile|string|false||
|totalNum|string|false||
|num|string|false||
|operatorName|string|false||
|createAt|string|false||
|merchantNum|string|false||

</details>

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

<details>
 <summary>
 <code>ImportInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|type|string|false||
|additionalData|string|false||

</details>

<details>
 <summary>
 <code>ExportInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|type|string|false||
|searchData|string|false||
|additionalData|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"list":[
	{
"id":"id",
"actId":"actId",
"mobile":"mobile",
"totalNum":"totalNum",
"num":"num",
"operatorName":"operatorName",
"createAt":"createAt",
"merchantNum":"merchantNum"
}
],
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"importRepeatMobileWhite":{
"type":"type",
"additionalData":"additionalData"
},
"exportRepeatMobileWhite":{
"type":"type",
"searchData":"searchData",
"additionalData":"additionalData"
}
}
}
````



<a name="RepeatMobileWhiteEdit"></a>

## 前端method: InteractluckyactivityRepeatMobileWhiteEdit

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|operator|[Operator](#Operator)|false||
|newMobile|string|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"actId":"actId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"newMobile":"newMobile"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="DelRepeatMobileWhite"></a>

## 前端method: InteractluckyactivityDelRepeatMobileWhite

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"id":"id",
"actId":"actId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="AdminReceiveAddressEdit"></a>

## 前端method: InteractluckyactivityAdminReceiveAddressEdit

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveId|string|false||
|mobile|string|false||
|address|string|false||
|province|string|false||
|city|string|false||
|district|string|false||
|name|string|false||
|addressId|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"receiveId":"receiveId",
"mobile":"mobile",
"address":"address",
"province":"province",
"city":"city",
"district":"district",
"name":"name",
"addressId":"addressId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="ActivityGiveBackPoint"></a>

## 前端method: InteractluckyactivityActivityGiveBackPoint

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"actId":"actId",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="BatchCreateParticipateCode"></a>

## 前端method: InteractluckyactivityBatchCreateParticipateCode

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|codeNum|string|false||
|operator|[Operator](#Operator)|false||

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"actId":"actId",
"codeNum":"codeNum",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="ActivityListGetStaff"></a>

## 前端method: InteractluckyactivityActivityListGetStaff

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|status|string|false||
|marketer|[Marketer](#Marketer)|false||

<details>
 <summary>
 <code>Marketer </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|name|string|false||
|mobile|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"page":"page",
"pageSize":"pageSize",
"status":"status",
"marketer":{
"isSuper":true,
"userId":"userId",
"name":"name",
"mobile":"mobile",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ActivityListGetStaffList](#ActivityListGetStaffList)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ActivityListGetStaffList </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|pageInfo|[PageInfo](#PageInfo)|false||
|list|[][ActivityArr](#ActivityArr)|false||

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

<details>
 <summary>
 <code>ActivityArr </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|channelId|string|false||
|no|string|false||
|name|string|false||
|orgId|string|false||
|startedAt|string|false||
|endAt|string|false||
|timeRule|[TimeRule](#TimeRule)|false||
|timeRuleDay|string|false||
|timeRuleWeek|string|false||
|timeRuleTime|string|false||
|takePartTotalNum|string|false||
|takePartDayNum|string|false||
|takePartUserNum|string|false||
|isMarketerExclusive|bool|false||
|userCodeNum|string|false||
|userDayCodeNum|string|false||
|takePartCodeNum|string|false||
|isShare|bool|false||
|shareImgId|string|false||
|status|string|false||
|addreVerifyType|string|false||
|addreIds|[]string|false||
|verifyType|string|false||
|operator|[Operator](#Operator)|false||
|addressData|[]string|false||
|ShareImgData|[ImageData](#ImageData)|false||
|verifyMemberLevel|string|false||
|verifyApiType|string|false||
|isApiMemberLevelMatchPrizePool|bool|false||
|thirdActId|string|false||
|verifyDesc|string|false||
|verifyNotPass|string|false||
|verifyPass|string|false||
|takePartIsBucklePoint|bool|false||
|pointType|string|false||
|pointNum|string|false||
|isMarketerSelectPrizePool|bool|false||
|recommendImgId|string|false||
|imgId|string|false||
|buttonImgId|string|false||
|bgColor|string|false||
|textColor|string|false||
|isShowUserTakePartNum|bool|false||
|userTakePartInitNum|string|false||
|isShowLuckyUser|bool|false||
|luckyUserNum|string|false||
|recommendImgData|[ImageData](#ImageData)|false||
|buttonImgData|[ImageData](#ImageData)|false||
|imgData|[ImageData](#ImageData)|false||
|prizePoolCount|string|false||
|actRule|string|false||
|shareTitle|string|false||
|luckyBgImg|[ImageData](#ImageData)|false||
|isShowActivityGift|bool|false||
|isVerifyCard|bool|false||
|verifyCardCode|string|false||
|verifyCardNumber|string|false||
|verifyCardDesc|string|false||
|verifyCardPrice|string|false||
|isStaffPartakeCode|bool|false||
|isStaffRegistration|bool|false||
|staffRegisterConfig|[]string|false||
|isRegisterVerifyPualification|bool|false||
|gameType|string|false||
|buttonMsg|string|false||
|activityStatus|string|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>TimeRule </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|rule|string|false||
|startedAt|string|false||
|endedAt|string|false||
|ruleDay|string|false||
|ruleWeek|string|false||
|ruleTime|string|false||

</details>

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"list":[
	{
"id":"id",
"channelId":"channelId",
"no":"no",
"name":"name",
"orgId":"orgId",
"startedAt":"startedAt",
"endAt":"endAt",
"timeRule":{
"rule":"rule",
"startedAt":"startedAt",
"endedAt":"endedAt",
"ruleDay":"ruleDay",
"ruleWeek":"ruleWeek",
"ruleTime":"ruleTime"
},
"timeRuleDay":"timeRuleDay",
"timeRuleWeek":"timeRuleWeek",
"timeRuleTime":"timeRuleTime",
"takePartTotalNum":"takePartTotalNum",
"takePartDayNum":"takePartDayNum",
"takePartUserNum":"takePartUserNum",
"isMarketerExclusive":true,
"userCodeNum":"userCodeNum",
"userDayCodeNum":"userDayCodeNum",
"takePartCodeNum":"takePartCodeNum",
"isShare":true,
"shareImgId":"shareImgId",
"status":"status",
"addreVerifyType":"addreVerifyType",
"addreIds":["addreIds"],
"verifyType":"verifyType",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"addressData":["addressData"],
"ShareImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"verifyMemberLevel":"verifyMemberLevel",
"verifyApiType":"verifyApiType",
"isApiMemberLevelMatchPrizePool":true,
"thirdActId":"thirdActId",
"verifyDesc":"verifyDesc",
"verifyNotPass":"verifyNotPass",
"verifyPass":"verifyPass",
"takePartIsBucklePoint":true,
"pointType":"pointType",
"pointNum":"pointNum",
"isMarketerSelectPrizePool":true,
"recommendImgId":"recommendImgId",
"imgId":"imgId",
"buttonImgId":"buttonImgId",
"bgColor":"bgColor",
"textColor":"textColor",
"isShowUserTakePartNum":true,
"userTakePartInitNum":"userTakePartInitNum",
"isShowLuckyUser":true,
"luckyUserNum":"luckyUserNum",
"recommendImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"buttonImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"prizePoolCount":"prizePoolCount",
"actRule":"actRule",
"shareTitle":"shareTitle",
"luckyBgImg":{
"id":"id",
"name":"name",
"url":"url"
},
"isShowActivityGift":true,
"isVerifyCard":true,
"verifyCardCode":"verifyCardCode",
"verifyCardNumber":"verifyCardNumber",
"verifyCardDesc":"verifyCardDesc",
"verifyCardPrice":"verifyCardPrice",
"isStaffPartakeCode":true,
"isStaffRegistration":true,
"staffRegisterConfig":["staffRegisterConfig"],
"isRegisterVerifyPualification":true,
"gameType":"gameType",
"buttonMsg":"buttonMsg",
"activityStatus":"activityStatus"
}
]
}
}
````



<a name="PrizePoolListGetStaff"></a>

## 前端method: InteractluckyactivityPrizePoolListGetStaff

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|marketer|[Marketer](#Marketer)|false||
|actId|string|false||

<details>
 <summary>
 <code>Marketer </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|name|string|false||
|mobile|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"marketer":{
"isSuper":true,
"userId":"userId",
"name":"name",
"mobile":"mobile",
"orgId":"orgId",
"channelId":"channelId"
},
"actId":"actId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizePoolList](#PrizePoolList)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizePoolList </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|list|[][PoolInfo](#PoolInfo)|false||

<details>
 <summary>
 <code>PoolInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|name|string|false||
|type|string|false||
|memberLevel|[]string|false||
|thirdPrizePoolId|string|false||
|isDefault|bool|false||
|specialInterval|string|false||
|prizeNum|string|false||
|createdAt|string|false||
|adminName|string|false||
|isVerify|bool|false||
|thirdOrgNumber|string|false||
|poolBgImageData|[ImageData](#ImageData)|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"list":[
	{
"id":"id",
"actId":"actId",
"name":"name",
"type":"type",
"memberLevel":["memberLevel"],
"thirdPrizePoolId":"thirdPrizePoolId",
"isDefault":true,
"specialInterval":"specialInterval",
"prizeNum":"prizeNum",
"createdAt":"createdAt",
"adminName":"adminName",
"isVerify":true,
"thirdOrgNumber":"thirdOrgNumber",
"poolBgImageData":{
"id":"id",
"name":"name",
"url":"url"
}
}
]
}
}
````



<a name="MarketerCodesCreate"></a>

## 前端method: InteractluckyactivityMarketerCodesCreate

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|orgId|string|false||
|prizePoolId|string|false||
|marketer|[Marketer](#Marketer)|false||

<details>
 <summary>
 <code>Marketer </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|name|string|false||
|mobile|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"actId":"actId",
"orgId":"orgId",
"prizePoolId":"prizePoolId",
"marketer":{
"isSuper":true,
"userId":"userId",
"name":"name",
"mobile":"mobile",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[MarketerCodeInfo](#MarketerCodeInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>MarketerCodeInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|no|string|false||
|actId|string|false||
|orgId|string|false||
|prizePoolId|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"no":"no",
"actId":"actId",
"orgId":"orgId",
"prizePoolId":"prizePoolId"
}
}
````



<a name="MarketerCodesInfo"></a>

## 前端method: InteractluckyactivityMarketerCodesInfo

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|no|string|false||

### 请求例子:
````
{
"no":"no"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[MarketerCodeInfo](#MarketerCodeInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>MarketerCodeInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|no|string|false||
|actId|string|false||
|orgId|string|false||
|prizePoolId|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"no":"no",
"actId":"actId",
"orgId":"orgId",
"prizePoolId":"prizePoolId"
}
}
````



<a name="ReceiveListGetStaff"></a>

## 前端method: InteractluckyactivityReceiveListGetStaff

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|keyword|string|false||
|page|string|false||
|pageSize|string|false||
|marketer|[Marketer](#Marketer)|false||

<details>
 <summary>
 <code>Marketer </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|name|string|false||
|mobile|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"actId":"actId",
"keyword":"keyword",
"page":"page",
"pageSize":"pageSize",
"marketer":{
"isSuper":true,
"userId":"userId",
"name":"name",
"mobile":"mobile",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ReceiveData](#ReceiveData)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ReceiveData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|pageInfo|[PageInfo](#PageInfo)|false||
|list|[][ReceiveArr](#ReceiveArr)|false||
|panel|[Panel](#Panel)|false||

<details>
 <summary>
 <code>ReceiveArr </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|userId|string|false||
|wxNickname|string|false||
|wxAvatar|string|false||
|wxMobile|string|false||
|createdAt|string|false||
|status|string|false||
|giftName|string|false||

</details>

<details>
 <summary>
 <code>Panel </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|useCount|string|false||

</details>

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"list":[
	{
"userId":"userId",
"wxNickname":"wxNickname",
"wxAvatar":"wxAvatar",
"wxMobile":"wxMobile",
"createdAt":"createdAt",
"status":"status",
"giftName":"giftName"
}
],
"panel":{
"useCount":"useCount"
}
}
}
````



<a name="ActivityInfoGetStaff"></a>

## 前端method: InteractluckyactivityActivityInfoGetStaff

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|marketer|[Marketer](#Marketer)|false||

<details>
 <summary>
 <code>Marketer </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|name|string|false||
|mobile|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"actId":"actId",
"marketer":{
"isSuper":true,
"userId":"userId",
"name":"name",
"mobile":"mobile",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ActivityArr](#ActivityArr)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ActivityArr </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|channelId|string|false||
|no|string|false||
|name|string|false||
|orgId|string|false||
|startedAt|string|false||
|endAt|string|false||
|timeRule|[TimeRule](#TimeRule)|false||
|timeRuleDay|string|false||
|timeRuleWeek|string|false||
|timeRuleTime|string|false||
|takePartTotalNum|string|false||
|takePartDayNum|string|false||
|takePartUserNum|string|false||
|isMarketerExclusive|bool|false||
|userCodeNum|string|false||
|userDayCodeNum|string|false||
|takePartCodeNum|string|false||
|isShare|bool|false||
|shareImgId|string|false||
|status|string|false||
|addreVerifyType|string|false||
|addreIds|[]string|false||
|verifyType|string|false||
|operator|[Operator](#Operator)|false||
|addressData|[]string|false||
|ShareImgData|[ImageData](#ImageData)|false||
|verifyMemberLevel|string|false||
|verifyApiType|string|false||
|isApiMemberLevelMatchPrizePool|bool|false||
|thirdActId|string|false||
|verifyDesc|string|false||
|verifyNotPass|string|false||
|verifyPass|string|false||
|takePartIsBucklePoint|bool|false||
|pointType|string|false||
|pointNum|string|false||
|isMarketerSelectPrizePool|bool|false||
|recommendImgId|string|false||
|imgId|string|false||
|buttonImgId|string|false||
|bgColor|string|false||
|textColor|string|false||
|isShowUserTakePartNum|bool|false||
|userTakePartInitNum|string|false||
|isShowLuckyUser|bool|false||
|luckyUserNum|string|false||
|recommendImgData|[ImageData](#ImageData)|false||
|buttonImgData|[ImageData](#ImageData)|false||
|imgData|[ImageData](#ImageData)|false||
|prizePoolCount|string|false||
|actRule|string|false||
|shareTitle|string|false||
|luckyBgImg|[ImageData](#ImageData)|false||
|isShowActivityGift|bool|false||
|isVerifyCard|bool|false||
|verifyCardCode|string|false||
|verifyCardNumber|string|false||
|verifyCardDesc|string|false||
|verifyCardPrice|string|false||
|isStaffPartakeCode|bool|false||
|isStaffRegistration|bool|false||
|staffRegisterConfig|[]string|false||
|isRegisterVerifyPualification|bool|false||
|gameType|string|false||
|buttonMsg|string|false||
|activityStatus|string|false||

<details>
 <summary>
 <code>TimeRule </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|rule|string|false||
|startedAt|string|false||
|endedAt|string|false||
|ruleDay|string|false||
|ruleWeek|string|false||
|ruleTime|string|false||

</details>

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"channelId":"channelId",
"no":"no",
"name":"name",
"orgId":"orgId",
"startedAt":"startedAt",
"endAt":"endAt",
"timeRule":{
"rule":"rule",
"startedAt":"startedAt",
"endedAt":"endedAt",
"ruleDay":"ruleDay",
"ruleWeek":"ruleWeek",
"ruleTime":"ruleTime"
},
"timeRuleDay":"timeRuleDay",
"timeRuleWeek":"timeRuleWeek",
"timeRuleTime":"timeRuleTime",
"takePartTotalNum":"takePartTotalNum",
"takePartDayNum":"takePartDayNum",
"takePartUserNum":"takePartUserNum",
"isMarketerExclusive":true,
"userCodeNum":"userCodeNum",
"userDayCodeNum":"userDayCodeNum",
"takePartCodeNum":"takePartCodeNum",
"isShare":true,
"shareImgId":"shareImgId",
"status":"status",
"addreVerifyType":"addreVerifyType",
"addreIds":["addreIds"],
"verifyType":"verifyType",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"addressData":["addressData"],
"ShareImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"verifyMemberLevel":"verifyMemberLevel",
"verifyApiType":"verifyApiType",
"isApiMemberLevelMatchPrizePool":true,
"thirdActId":"thirdActId",
"verifyDesc":"verifyDesc",
"verifyNotPass":"verifyNotPass",
"verifyPass":"verifyPass",
"takePartIsBucklePoint":true,
"pointType":"pointType",
"pointNum":"pointNum",
"isMarketerSelectPrizePool":true,
"recommendImgId":"recommendImgId",
"imgId":"imgId",
"buttonImgId":"buttonImgId",
"bgColor":"bgColor",
"textColor":"textColor",
"isShowUserTakePartNum":true,
"userTakePartInitNum":"userTakePartInitNum",
"isShowLuckyUser":true,
"luckyUserNum":"luckyUserNum",
"recommendImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"buttonImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"prizePoolCount":"prizePoolCount",
"actRule":"actRule",
"shareTitle":"shareTitle",
"luckyBgImg":{
"id":"id",
"name":"name",
"url":"url"
},
"isShowActivityGift":true,
"isVerifyCard":true,
"verifyCardCode":"verifyCardCode",
"verifyCardNumber":"verifyCardNumber",
"verifyCardDesc":"verifyCardDesc",
"verifyCardPrice":"verifyCardPrice",
"isStaffPartakeCode":true,
"isStaffRegistration":true,
"staffRegisterConfig":["staffRegisterConfig"],
"isRegisterVerifyPualification":true,
"gameType":"gameType",
"buttonMsg":"buttonMsg",
"activityStatus":"activityStatus"
}
}
````



<a name="ReceiveVerifyStaff"></a>

## 前端method: InteractluckyactivityReceiveVerifyStaff

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|marketerId|string|false||
|verifySn|string|false||
|receiveId|string|false||

### 请求例子:
````
{
"marketerId":"marketerId",
"verifySn":"verifySn",
"receiveId":"receiveId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="MarketerGetReceiveDetails"></a>

## 前端method: InteractluckyactivityMarketerGetReceiveDetails

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|no|string|false||
|marketer|[Marketer](#Marketer)|false||

<details>
 <summary>
 <code>Marketer </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|name|string|false||
|mobile|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"no":"no",
"marketer":{
"isSuper":true,
"userId":"userId",
"name":"name",
"mobile":"mobile",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ReceiveDetails](#ReceiveDetails)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ReceiveDetails </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|prizeType|string|false||
|prizeName|string|false||
|prizeImageData|[ImageData](#ImageData)|false||
|OrderNo|string|false||
|status|string|false||
|org|[OrgData](#OrgData)|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>OrgData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|bool|[//](#//)|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"prizeType":"prizeType",
"prizeName":"prizeName",
"prizeImageData":{
"id":"id",
"name":"name",
"url":"url"
},
"OrderNo":"OrderNo",
"status":"status",
"org":{
"id":"id",
"name":"name",
"bool":{}
}
}
}
````



<a name="GetMobileCode"></a>

## 前端method: InteractluckyactivityGetMobileCode

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|mobile|string|false||

### 请求例子:
````
{
"mobile":"mobile"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[GetMobileCodeResp_Data](#GetMobileCodeResp_Data)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>GetMobileCodeResp_Data </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|duration|int32|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"duration":1
}
}
````



<a name="MarketerCreateUser"></a>

## 前端method: InteractluckyactivityMarketerCreateUser

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|realName|string|false||
|mobile|string|false||
|idCode|string|false||
|verifyCode|string|false||
|channelId|string|false||
|marketer|[Marketer](#Marketer)|false||

<details>
 <summary>
 <code>Marketer </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|name|string|false||
|mobile|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"realName":"realName",
"mobile":"mobile",
"idCode":"idCode",
"verifyCode":"verifyCode",
"channelId":"channelId",
"marketer":{
"isSuper":true,
"userId":"userId",
"name":"name",
"mobile":"mobile",
"orgId":"orgId",
"channelId":"channelId"
}
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|userInfo|[UserInfo](#UserInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>UserInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|userId|string|false||
|mobile|string|false||
|idCode|string|false||
|realName|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"userInfo":{
"userId":"userId",
"mobile":"mobile",
"idCode":"idCode",
"realName":"realName"
}
}
````



<a name="MarketerGetPrizeList"></a>

## 前端method: InteractluckyactivityMarketerGetPrizeList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|poolId|string|false||
|prizeType|string|false||

### 请求例子:
````
{
"poolId":"poolId",
"prizeType":"prizeType"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizeData](#PrizeData)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizeData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|list|[][PrizeInfo](#PrizeInfo)|false||

<details>
 <summary>
 <code>PrizeInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|name|string|false||
|type|string|false||
|money|string|false||
|totalNum|string|false||
|isSpecial|bool|false||
|imgId|string|false||
|giftId|string|false||
|sort|string|false||
|desc|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|imgData|[ImageData](#ImageData)|false||
|eId|string|false||
|count|string|false||
|limitArea|[Addre](#Addre)|false||
|limitProvince|string|false||
|limitCity|string|false||
|limitDistrict|string|false||
|isSinceExtractOrgNum|bool|false||
|prizeId|string|false||
|expendNum|string|false||
|costPrice|string|false||
|pointNum|string|false||
|couponActId|string|false||
|isDelayIssue|bool|false||
|eventId|string|false||
|delayDay|string|false||
|delayIssueMsg|string|false||
|couponBagData|[][CouponBagData](#CouponBagData)|false||
|residueNum|string|false||
|appointmentStartAt|string|false||
|appointmentEndAt|string|false||
|goodsId|string|false||
|delayDayNum|string|false||
|isAutomaticAppointment|bool|false||
|automaticAppointmentMsg|string|false||
|autoPreOrgId|string|false||
|dayCount|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|awardNum|string|false||
|awardUnit|string|false||
|jumpConfig|string|false||
|surplusNum|string|false||
|payType|string|false||
|detailDesc|string|false||
|status|string|false||
|isDefaultNetAppointment|bool|false||
|defaultPreOrgId|string|false||
|isLuckDrawOneFrequency|bool|false||
|isOpenSpecialRule|string|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

<details>
 <summary>
 <code>CouponBagData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|stockId|string|false||
|num|string|false||
|couponActId|string|false||
|id|string|false||
|name|string|false||
|amount|string|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"list":[
	{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"name":"name",
"type":"type",
"money":"money",
"totalNum":"totalNum",
"isSpecial":true,
"imgId":"imgId",
"giftId":"giftId",
"sort":"sort",
"desc":"desc",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"eId":"eId",
"count":"count",
"limitArea":{
"id":"id",
"fullName":"fullName"
},
"limitProvince":"limitProvince",
"limitCity":"limitCity",
"limitDistrict":"limitDistrict",
"isSinceExtractOrgNum":true,
"prizeId":"prizeId",
"expendNum":"expendNum",
"costPrice":"costPrice",
"pointNum":"pointNum",
"couponActId":"couponActId",
"isDelayIssue":true,
"eventId":"eventId",
"delayDay":"delayDay",
"delayIssueMsg":"delayIssueMsg",
"couponBagData":[
	{
"stockId":"stockId",
"num":"num",
"couponActId":"couponActId",
"id":"id",
"name":"name",
"amount":"amount"
}
],
"residueNum":"residueNum",
"appointmentStartAt":"appointmentStartAt",
"appointmentEndAt":"appointmentEndAt",
"goodsId":"goodsId",
"delayDayNum":"delayDayNum",
"isAutomaticAppointment":true,
"automaticAppointmentMsg":"automaticAppointmentMsg",
"autoPreOrgId":"autoPreOrgId",
"dayCount":"dayCount",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"awardNum":"awardNum",
"awardUnit":"awardUnit",
"jumpConfig":"jumpConfig",
"surplusNum":"surplusNum",
"payType":"payType",
"detailDesc":"detailDesc",
"status":"status",
"isDefaultNetAppointment":true,
"defaultPreOrgId":"defaultPreOrgId",
"isLuckDrawOneFrequency":true,
"isOpenSpecialRule":"isOpenSpecialRule"
}
]
}
}
````



<a name="MarketerIssuePrize"></a>

## 前端method: InteractluckyactivityMarketerIssuePrize

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|mobile|string|false||
|identity|string|false||
|userId|string|false||
|longitude|string|false||
|latitude|string|false||
|actId|string|false||
|idCard|string|false||
|name|string|false||
|takePartCode|string|false||
|partakeCodeStaff|string|false||
|prizeId|string|false||
|marketer|[Marketer](#Marketer)|false||
|ticketFileId|string|false||

<details>
 <summary>
 <code>Marketer </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|name|string|false||
|mobile|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

### 请求例子:
````
{
"mobile":"mobile",
"identity":"identity",
"userId":"userId",
"longitude":"longitude",
"latitude":"latitude",
"actId":"actId",
"idCard":"idCard",
"name":"name",
"takePartCode":"takePartCode",
"partakeCodeStaff":"partakeCodeStaff",
"prizeId":"prizeId",
"marketer":{
"isSuper":true,
"userId":"userId",
"name":"name",
"mobile":"mobile",
"orgId":"orgId",
"channelId":"channelId"
},
"ticketFileId":"ticketFileId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="ActivityInfoGet"></a>

## 前端method: InteractluckyactivityActivityInfoGet

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||

### 请求例子:
````
{
"id":"id"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ActivityArr](#ActivityArr)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ActivityArr </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|channelId|string|false||
|no|string|false||
|name|string|false||
|orgId|string|false||
|startedAt|string|false||
|endAt|string|false||
|timeRule|[TimeRule](#TimeRule)|false||
|timeRuleDay|string|false||
|timeRuleWeek|string|false||
|timeRuleTime|string|false||
|takePartTotalNum|string|false||
|takePartDayNum|string|false||
|takePartUserNum|string|false||
|isMarketerExclusive|bool|false||
|userCodeNum|string|false||
|userDayCodeNum|string|false||
|takePartCodeNum|string|false||
|isShare|bool|false||
|shareImgId|string|false||
|status|string|false||
|addreVerifyType|string|false||
|addreIds|[]string|false||
|verifyType|string|false||
|operator|[Operator](#Operator)|false||
|addressData|[]string|false||
|ShareImgData|[ImageData](#ImageData)|false||
|verifyMemberLevel|string|false||
|verifyApiType|string|false||
|isApiMemberLevelMatchPrizePool|bool|false||
|thirdActId|string|false||
|verifyDesc|string|false||
|verifyNotPass|string|false||
|verifyPass|string|false||
|takePartIsBucklePoint|bool|false||
|pointType|string|false||
|pointNum|string|false||
|isMarketerSelectPrizePool|bool|false||
|recommendImgId|string|false||
|imgId|string|false||
|buttonImgId|string|false||
|bgColor|string|false||
|textColor|string|false||
|isShowUserTakePartNum|bool|false||
|userTakePartInitNum|string|false||
|isShowLuckyUser|bool|false||
|luckyUserNum|string|false||
|recommendImgData|[ImageData](#ImageData)|false||
|buttonImgData|[ImageData](#ImageData)|false||
|imgData|[ImageData](#ImageData)|false||
|prizePoolCount|string|false||
|actRule|string|false||
|shareTitle|string|false||
|luckyBgImg|[ImageData](#ImageData)|false||
|isShowActivityGift|bool|false||
|isVerifyCard|bool|false||
|verifyCardCode|string|false||
|verifyCardNumber|string|false||
|verifyCardDesc|string|false||
|verifyCardPrice|string|false||
|isStaffPartakeCode|bool|false||
|isStaffRegistration|bool|false||
|staffRegisterConfig|[]string|false||
|isRegisterVerifyPualification|bool|false||
|gameType|string|false||
|buttonMsg|string|false||
|activityStatus|string|false||

<details>
 <summary>
 <code>TimeRule </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|rule|string|false||
|startedAt|string|false||
|endedAt|string|false||
|ruleDay|string|false||
|ruleWeek|string|false||
|ruleTime|string|false||

</details>

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"channelId":"channelId",
"no":"no",
"name":"name",
"orgId":"orgId",
"startedAt":"startedAt",
"endAt":"endAt",
"timeRule":{
"rule":"rule",
"startedAt":"startedAt",
"endedAt":"endedAt",
"ruleDay":"ruleDay",
"ruleWeek":"ruleWeek",
"ruleTime":"ruleTime"
},
"timeRuleDay":"timeRuleDay",
"timeRuleWeek":"timeRuleWeek",
"timeRuleTime":"timeRuleTime",
"takePartTotalNum":"takePartTotalNum",
"takePartDayNum":"takePartDayNum",
"takePartUserNum":"takePartUserNum",
"isMarketerExclusive":true,
"userCodeNum":"userCodeNum",
"userDayCodeNum":"userDayCodeNum",
"takePartCodeNum":"takePartCodeNum",
"isShare":true,
"shareImgId":"shareImgId",
"status":"status",
"addreVerifyType":"addreVerifyType",
"addreIds":["addreIds"],
"verifyType":"verifyType",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"addressData":["addressData"],
"ShareImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"verifyMemberLevel":"verifyMemberLevel",
"verifyApiType":"verifyApiType",
"isApiMemberLevelMatchPrizePool":true,
"thirdActId":"thirdActId",
"verifyDesc":"verifyDesc",
"verifyNotPass":"verifyNotPass",
"verifyPass":"verifyPass",
"takePartIsBucklePoint":true,
"pointType":"pointType",
"pointNum":"pointNum",
"isMarketerSelectPrizePool":true,
"recommendImgId":"recommendImgId",
"imgId":"imgId",
"buttonImgId":"buttonImgId",
"bgColor":"bgColor",
"textColor":"textColor",
"isShowUserTakePartNum":true,
"userTakePartInitNum":"userTakePartInitNum",
"isShowLuckyUser":true,
"luckyUserNum":"luckyUserNum",
"recommendImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"buttonImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"prizePoolCount":"prizePoolCount",
"actRule":"actRule",
"shareTitle":"shareTitle",
"luckyBgImg":{
"id":"id",
"name":"name",
"url":"url"
},
"isShowActivityGift":true,
"isVerifyCard":true,
"verifyCardCode":"verifyCardCode",
"verifyCardNumber":"verifyCardNumber",
"verifyCardDesc":"verifyCardDesc",
"verifyCardPrice":"verifyCardPrice",
"isStaffPartakeCode":true,
"isStaffRegistration":true,
"staffRegisterConfig":["staffRegisterConfig"],
"isRegisterVerifyPualification":true,
"gameType":"gameType",
"buttonMsg":"buttonMsg",
"activityStatus":"activityStatus"
}
}
````



<a name="ActivityQualificationVerify"></a>

## 前端method: InteractluckyactivityActivityQualificationVerify

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|mobile|string|false||
|identity|string|false||
|userId|string|false||
|longitude|string|false||
|latitude|string|false||
|actId|string|false||
|idCard|string|false||
|marketerCodeId|string|false||
|name|string|false||
|takePartCode|string|false||
|partakeCodeStaff|string|false||
|isStaff|bool|false||
|prizeId|string|false||
|clientType|string|false||
|memberInfo|[MemberInfo](#MemberInfo)|false||
|idType|string|false||

<details>
 <summary>
 <code>MemberInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|mobile|string|false||
|realName|string|false||
|idCard|string|false||
|address|string|false||
|province|string|false||
|city|string|false||
|district|string|false||
|userId|string|false||
|cstmNo|string|false||

</details>

### 请求例子:
````
{
"mobile":"mobile",
"identity":"identity",
"userId":"userId",
"longitude":"longitude",
"latitude":"latitude",
"actId":"actId",
"idCard":"idCard",
"marketerCodeId":"marketerCodeId",
"name":"name",
"takePartCode":"takePartCode",
"partakeCodeStaff":"partakeCodeStaff",
"isStaff":true,
"prizeId":"prizeId",
"clientType":"clientType",
"memberInfo":{
"mobile":"mobile",
"realName":"realName",
"idCard":"idCard",
"address":"address",
"province":"province",
"city":"city",
"district":"district",
"userId":"userId",
"cstmNo":"cstmNo"
},
"idType":"idType"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PoolInfo](#PoolInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PoolInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|name|string|false||
|type|string|false||
|memberLevel|[]string|false||
|thirdPrizePoolId|string|false||
|isDefault|bool|false||
|specialInterval|string|false||
|prizeNum|string|false||
|createdAt|string|false||
|adminName|string|false||
|isVerify|bool|false||
|thirdOrgNumber|string|false||
|poolBgImageData|[ImageData](#ImageData)|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"actId":"actId",
"name":"name",
"type":"type",
"memberLevel":["memberLevel"],
"thirdPrizePoolId":"thirdPrizePoolId",
"isDefault":true,
"specialInterval":"specialInterval",
"prizeNum":"prizeNum",
"createdAt":"createdAt",
"adminName":"adminName",
"isVerify":true,
"thirdOrgNumber":"thirdOrgNumber",
"poolBgImageData":{
"id":"id",
"name":"name",
"url":"url"
}
}
}
````



<a name="ActivityPartakeCodeVerify"></a>

## 前端method: InteractluckyactivityActivityPartakeCodeVerify

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|code|string|false||
|userId|string|false||

### 请求例子:
````
{
"actId":"actId",
"code":"code",
"userId":"userId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="ActivityPartakeCodeVerifyInfo"></a>

## 前端method: InteractluckyactivityActivityPartakeCodeVerifyInfo

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|userId|string|false||

### 请求例子:
````
{
"actId":"actId",
"userId":"userId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ActivityPartakeCodeVerifyInfoData](#ActivityPartakeCodeVerifyInfoData)|false||

<details>
 <summary>
 <code>ActivityPartakeCodeVerifyInfoData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|status|string|false||
|partakeCodeStaff|string|false||

</details>

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"status":"status",
"partakeCodeStaff":"partakeCodeStaff"
}
}
````



<a name="ActivityQualificationVerifyPoint"></a>

## 前端method: InteractluckyactivityActivityQualificationVerifyPoint

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|userId|string|false||
|mobile|string|false||
|mid|string|false||

### 请求例子:
````
{
"actId":"actId",
"userId":"userId",
"mobile":"mobile",
"mid":"mid"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="ActivityListGetUser"></a>

## 前端method: InteractluckyactivityActivityListGetUser

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||

### 请求例子:
````
{
"page":"page",
"pageSize":"pageSize"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ActivityListGetStaffList](#ActivityListGetStaffList)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ActivityListGetStaffList </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|pageInfo|[PageInfo](#PageInfo)|false||
|list|[][ActivityArr](#ActivityArr)|false||

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

<details>
 <summary>
 <code>ActivityArr </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|channelId|string|false||
|no|string|false||
|name|string|false||
|orgId|string|false||
|startedAt|string|false||
|endAt|string|false||
|timeRule|[TimeRule](#TimeRule)|false||
|timeRuleDay|string|false||
|timeRuleWeek|string|false||
|timeRuleTime|string|false||
|takePartTotalNum|string|false||
|takePartDayNum|string|false||
|takePartUserNum|string|false||
|isMarketerExclusive|bool|false||
|userCodeNum|string|false||
|userDayCodeNum|string|false||
|takePartCodeNum|string|false||
|isShare|bool|false||
|shareImgId|string|false||
|status|string|false||
|addreVerifyType|string|false||
|addreIds|[]string|false||
|verifyType|string|false||
|operator|[Operator](#Operator)|false||
|addressData|[]string|false||
|ShareImgData|[ImageData](#ImageData)|false||
|verifyMemberLevel|string|false||
|verifyApiType|string|false||
|isApiMemberLevelMatchPrizePool|bool|false||
|thirdActId|string|false||
|verifyDesc|string|false||
|verifyNotPass|string|false||
|verifyPass|string|false||
|takePartIsBucklePoint|bool|false||
|pointType|string|false||
|pointNum|string|false||
|isMarketerSelectPrizePool|bool|false||
|recommendImgId|string|false||
|imgId|string|false||
|buttonImgId|string|false||
|bgColor|string|false||
|textColor|string|false||
|isShowUserTakePartNum|bool|false||
|userTakePartInitNum|string|false||
|isShowLuckyUser|bool|false||
|luckyUserNum|string|false||
|recommendImgData|[ImageData](#ImageData)|false||
|buttonImgData|[ImageData](#ImageData)|false||
|imgData|[ImageData](#ImageData)|false||
|prizePoolCount|string|false||
|actRule|string|false||
|shareTitle|string|false||
|luckyBgImg|[ImageData](#ImageData)|false||
|isShowActivityGift|bool|false||
|isVerifyCard|bool|false||
|verifyCardCode|string|false||
|verifyCardNumber|string|false||
|verifyCardDesc|string|false||
|verifyCardPrice|string|false||
|isStaffPartakeCode|bool|false||
|isStaffRegistration|bool|false||
|staffRegisterConfig|[]string|false||
|isRegisterVerifyPualification|bool|false||
|gameType|string|false||
|buttonMsg|string|false||
|activityStatus|string|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>TimeRule </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|rule|string|false||
|startedAt|string|false||
|endedAt|string|false||
|ruleDay|string|false||
|ruleWeek|string|false||
|ruleTime|string|false||

</details>

<details>
 <summary>
 <code>Operator </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|isSuper|bool|false||
|userId|string|false||
|mobile|string|false||
|name|string|false||
|orgId|string|false||
|channelId|string|false||

</details>

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"list":[
	{
"id":"id",
"channelId":"channelId",
"no":"no",
"name":"name",
"orgId":"orgId",
"startedAt":"startedAt",
"endAt":"endAt",
"timeRule":{
"rule":"rule",
"startedAt":"startedAt",
"endedAt":"endedAt",
"ruleDay":"ruleDay",
"ruleWeek":"ruleWeek",
"ruleTime":"ruleTime"
},
"timeRuleDay":"timeRuleDay",
"timeRuleWeek":"timeRuleWeek",
"timeRuleTime":"timeRuleTime",
"takePartTotalNum":"takePartTotalNum",
"takePartDayNum":"takePartDayNum",
"takePartUserNum":"takePartUserNum",
"isMarketerExclusive":true,
"userCodeNum":"userCodeNum",
"userDayCodeNum":"userDayCodeNum",
"takePartCodeNum":"takePartCodeNum",
"isShare":true,
"shareImgId":"shareImgId",
"status":"status",
"addreVerifyType":"addreVerifyType",
"addreIds":["addreIds"],
"verifyType":"verifyType",
"operator":{
"isSuper":true,
"userId":"userId",
"mobile":"mobile",
"name":"name",
"orgId":"orgId",
"channelId":"channelId"
},
"addressData":["addressData"],
"ShareImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"verifyMemberLevel":"verifyMemberLevel",
"verifyApiType":"verifyApiType",
"isApiMemberLevelMatchPrizePool":true,
"thirdActId":"thirdActId",
"verifyDesc":"verifyDesc",
"verifyNotPass":"verifyNotPass",
"verifyPass":"verifyPass",
"takePartIsBucklePoint":true,
"pointType":"pointType",
"pointNum":"pointNum",
"isMarketerSelectPrizePool":true,
"recommendImgId":"recommendImgId",
"imgId":"imgId",
"buttonImgId":"buttonImgId",
"bgColor":"bgColor",
"textColor":"textColor",
"isShowUserTakePartNum":true,
"userTakePartInitNum":"userTakePartInitNum",
"isShowLuckyUser":true,
"luckyUserNum":"luckyUserNum",
"recommendImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"buttonImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"prizePoolCount":"prizePoolCount",
"actRule":"actRule",
"shareTitle":"shareTitle",
"luckyBgImg":{
"id":"id",
"name":"name",
"url":"url"
},
"isShowActivityGift":true,
"isVerifyCard":true,
"verifyCardCode":"verifyCardCode",
"verifyCardNumber":"verifyCardNumber",
"verifyCardDesc":"verifyCardDesc",
"verifyCardPrice":"verifyCardPrice",
"isStaffPartakeCode":true,
"isStaffRegistration":true,
"staffRegisterConfig":["staffRegisterConfig"],
"isRegisterVerifyPualification":true,
"gameType":"gameType",
"buttonMsg":"buttonMsg",
"activityStatus":"activityStatus"
}
]
}
}
````



<a name="PrizeListGet"></a>

## 前端method: InteractluckyactivityPrizeListGet

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|page|string|false||
|pageSize|string|false||

### 请求例子:
````
{
"actId":"actId",
"page":"page",
"pageSize":"pageSize"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizeList](#PrizeList)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizeList </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|pageInfo|[PageInfo](#PageInfo)|false||
|list|[][PrizeInfo](#PrizeInfo)|false||
|count|string|false||
|prizeStockNum|string|false||
|poolInfo|[PoolInfo](#PoolInfo)|false||
|dayCount|string|false||

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

<details>
 <summary>
 <code>PrizeInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|name|string|false||
|type|string|false||
|money|string|false||
|totalNum|string|false||
|isSpecial|bool|false||
|imgId|string|false||
|giftId|string|false||
|sort|string|false||
|desc|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|imgData|[ImageData](#ImageData)|false||
|eId|string|false||
|count|string|false||
|limitArea|[Addre](#Addre)|false||
|limitProvince|string|false||
|limitCity|string|false||
|limitDistrict|string|false||
|isSinceExtractOrgNum|bool|false||
|prizeId|string|false||
|expendNum|string|false||
|costPrice|string|false||
|pointNum|string|false||
|couponActId|string|false||
|isDelayIssue|bool|false||
|eventId|string|false||
|delayDay|string|false||
|delayIssueMsg|string|false||
|couponBagData|[][CouponBagData](#CouponBagData)|false||
|residueNum|string|false||
|appointmentStartAt|string|false||
|appointmentEndAt|string|false||
|goodsId|string|false||
|delayDayNum|string|false||
|isAutomaticAppointment|bool|false||
|automaticAppointmentMsg|string|false||
|autoPreOrgId|string|false||
|dayCount|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|awardNum|string|false||
|awardUnit|string|false||
|jumpConfig|string|false||
|surplusNum|string|false||
|payType|string|false||
|detailDesc|string|false||
|status|string|false||
|isDefaultNetAppointment|bool|false||
|defaultPreOrgId|string|false||
|isLuckDrawOneFrequency|bool|false||
|isOpenSpecialRule|string|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

<details>
 <summary>
 <code>CouponBagData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|stockId|string|false||
|num|string|false||
|couponActId|string|false||
|id|string|false||
|name|string|false||
|amount|string|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

</details>

<details>
 <summary>
 <code>PoolInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|name|string|false||
|type|string|false||
|memberLevel|[]string|false||
|thirdPrizePoolId|string|false||
|isDefault|bool|false||
|specialInterval|string|false||
|prizeNum|string|false||
|createdAt|string|false||
|adminName|string|false||
|isVerify|bool|false||
|thirdOrgNumber|string|false||
|poolBgImageData|[ImageData](#ImageData)|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"list":[
	{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"name":"name",
"type":"type",
"money":"money",
"totalNum":"totalNum",
"isSpecial":true,
"imgId":"imgId",
"giftId":"giftId",
"sort":"sort",
"desc":"desc",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"eId":"eId",
"count":"count",
"limitArea":{
"id":"id",
"fullName":"fullName"
},
"limitProvince":"limitProvince",
"limitCity":"limitCity",
"limitDistrict":"limitDistrict",
"isSinceExtractOrgNum":true,
"prizeId":"prizeId",
"expendNum":"expendNum",
"costPrice":"costPrice",
"pointNum":"pointNum",
"couponActId":"couponActId",
"isDelayIssue":true,
"eventId":"eventId",
"delayDay":"delayDay",
"delayIssueMsg":"delayIssueMsg",
"couponBagData":[
	{
"stockId":"stockId",
"num":"num",
"couponActId":"couponActId",
"id":"id",
"name":"name",
"amount":"amount"
}
],
"residueNum":"residueNum",
"appointmentStartAt":"appointmentStartAt",
"appointmentEndAt":"appointmentEndAt",
"goodsId":"goodsId",
"delayDayNum":"delayDayNum",
"isAutomaticAppointment":true,
"automaticAppointmentMsg":"automaticAppointmentMsg",
"autoPreOrgId":"autoPreOrgId",
"dayCount":"dayCount",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"awardNum":"awardNum",
"awardUnit":"awardUnit",
"jumpConfig":"jumpConfig",
"surplusNum":"surplusNum",
"payType":"payType",
"detailDesc":"detailDesc",
"status":"status",
"isDefaultNetAppointment":true,
"defaultPreOrgId":"defaultPreOrgId",
"isLuckDrawOneFrequency":true,
"isOpenSpecialRule":"isOpenSpecialRule"
}
],
"count":"count",
"prizeStockNum":"prizeStockNum",
"poolInfo":{
"id":"id",
"actId":"actId",
"name":"name",
"type":"type",
"memberLevel":["memberLevel"],
"thirdPrizePoolId":"thirdPrizePoolId",
"isDefault":true,
"specialInterval":"specialInterval",
"prizeNum":"prizeNum",
"createdAt":"createdAt",
"adminName":"adminName",
"isVerify":true,
"thirdOrgNumber":"thirdOrgNumber",
"poolBgImageData":{
"id":"id",
"name":"name",
"url":"url"
}
},
"dayCount":"dayCount"
}
}
````



<a name="PrizeListGetByActIdAndUserId"></a>

## 前端method: InteractluckyactivityPrizeListGetByActIdAndUserId

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|userId|string|false||
|actId|string|false||

### 请求例子:
````
{
"userId":"userId",
"actId":"actId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizeList](#PrizeList)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizeList </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|pageInfo|[PageInfo](#PageInfo)|false||
|list|[][PrizeInfo](#PrizeInfo)|false||
|count|string|false||
|prizeStockNum|string|false||
|poolInfo|[PoolInfo](#PoolInfo)|false||
|dayCount|string|false||

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

<details>
 <summary>
 <code>PrizeInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|name|string|false||
|type|string|false||
|money|string|false||
|totalNum|string|false||
|isSpecial|bool|false||
|imgId|string|false||
|giftId|string|false||
|sort|string|false||
|desc|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|imgData|[ImageData](#ImageData)|false||
|eId|string|false||
|count|string|false||
|limitArea|[Addre](#Addre)|false||
|limitProvince|string|false||
|limitCity|string|false||
|limitDistrict|string|false||
|isSinceExtractOrgNum|bool|false||
|prizeId|string|false||
|expendNum|string|false||
|costPrice|string|false||
|pointNum|string|false||
|couponActId|string|false||
|isDelayIssue|bool|false||
|eventId|string|false||
|delayDay|string|false||
|delayIssueMsg|string|false||
|couponBagData|[][CouponBagData](#CouponBagData)|false||
|residueNum|string|false||
|appointmentStartAt|string|false||
|appointmentEndAt|string|false||
|goodsId|string|false||
|delayDayNum|string|false||
|isAutomaticAppointment|bool|false||
|automaticAppointmentMsg|string|false||
|autoPreOrgId|string|false||
|dayCount|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|awardNum|string|false||
|awardUnit|string|false||
|jumpConfig|string|false||
|surplusNum|string|false||
|payType|string|false||
|detailDesc|string|false||
|status|string|false||
|isDefaultNetAppointment|bool|false||
|defaultPreOrgId|string|false||
|isLuckDrawOneFrequency|bool|false||
|isOpenSpecialRule|string|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

<details>
 <summary>
 <code>CouponBagData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|stockId|string|false||
|num|string|false||
|couponActId|string|false||
|id|string|false||
|name|string|false||
|amount|string|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

</details>

<details>
 <summary>
 <code>PoolInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|name|string|false||
|type|string|false||
|memberLevel|[]string|false||
|thirdPrizePoolId|string|false||
|isDefault|bool|false||
|specialInterval|string|false||
|prizeNum|string|false||
|createdAt|string|false||
|adminName|string|false||
|isVerify|bool|false||
|thirdOrgNumber|string|false||
|poolBgImageData|[ImageData](#ImageData)|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"list":[
	{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"name":"name",
"type":"type",
"money":"money",
"totalNum":"totalNum",
"isSpecial":true,
"imgId":"imgId",
"giftId":"giftId",
"sort":"sort",
"desc":"desc",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"eId":"eId",
"count":"count",
"limitArea":{
"id":"id",
"fullName":"fullName"
},
"limitProvince":"limitProvince",
"limitCity":"limitCity",
"limitDistrict":"limitDistrict",
"isSinceExtractOrgNum":true,
"prizeId":"prizeId",
"expendNum":"expendNum",
"costPrice":"costPrice",
"pointNum":"pointNum",
"couponActId":"couponActId",
"isDelayIssue":true,
"eventId":"eventId",
"delayDay":"delayDay",
"delayIssueMsg":"delayIssueMsg",
"couponBagData":[
	{
"stockId":"stockId",
"num":"num",
"couponActId":"couponActId",
"id":"id",
"name":"name",
"amount":"amount"
}
],
"residueNum":"residueNum",
"appointmentStartAt":"appointmentStartAt",
"appointmentEndAt":"appointmentEndAt",
"goodsId":"goodsId",
"delayDayNum":"delayDayNum",
"isAutomaticAppointment":true,
"automaticAppointmentMsg":"automaticAppointmentMsg",
"autoPreOrgId":"autoPreOrgId",
"dayCount":"dayCount",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"awardNum":"awardNum",
"awardUnit":"awardUnit",
"jumpConfig":"jumpConfig",
"surplusNum":"surplusNum",
"payType":"payType",
"detailDesc":"detailDesc",
"status":"status",
"isDefaultNetAppointment":true,
"defaultPreOrgId":"defaultPreOrgId",
"isLuckDrawOneFrequency":true,
"isOpenSpecialRule":"isOpenSpecialRule"
}
],
"count":"count",
"prizeStockNum":"prizeStockNum",
"poolInfo":{
"id":"id",
"actId":"actId",
"name":"name",
"type":"type",
"memberLevel":["memberLevel"],
"thirdPrizePoolId":"thirdPrizePoolId",
"isDefault":true,
"specialInterval":"specialInterval",
"prizeNum":"prizeNum",
"createdAt":"createdAt",
"adminName":"adminName",
"isVerify":true,
"thirdOrgNumber":"thirdOrgNumber",
"poolBgImageData":{
"id":"id",
"name":"name",
"url":"url"
}
},
"dayCount":"dayCount"
}
}
````



<a name="LuckDrawUser"></a>

## 前端method: InteractluckyactivityLuckDrawUser

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|userId|string|false||
|actId|string|false||
|marketerCodeId|string|false||
|wxUser|[WxUser](#WxUser)|false||
|clientType|string|false||

<details>
 <summary>
 <code>WxUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|wxUnionId|string|false||
|wxMiniAppId|string|false||
|wxMiniOpenId|string|false||
|wxNickname|string|false||
|wxAvatar|string|false||
|wxMobileDim|string|false||
|realName|string|false||

</details>

### 请求例子:
````
{
"userId":"userId",
"actId":"actId",
"marketerCodeId":"marketerCodeId",
"wxUser":{
"wxUnionId":"wxUnionId",
"wxMiniAppId":"wxMiniAppId",
"wxMiniOpenId":"wxMiniOpenId",
"wxNickname":"wxNickname",
"wxAvatar":"wxAvatar",
"wxMobileDim":"wxMobileDim",
"realName":"realName"
},
"clientType":"clientType"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizeInfo](#PrizeInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizeInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|name|string|false||
|type|string|false||
|money|string|false||
|totalNum|string|false||
|isSpecial|bool|false||
|imgId|string|false||
|giftId|string|false||
|sort|string|false||
|desc|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|imgData|[ImageData](#ImageData)|false||
|eId|string|false||
|count|string|false||
|limitArea|[Addre](#Addre)|false||
|limitProvince|string|false||
|limitCity|string|false||
|limitDistrict|string|false||
|isSinceExtractOrgNum|bool|false||
|prizeId|string|false||
|expendNum|string|false||
|costPrice|string|false||
|pointNum|string|false||
|couponActId|string|false||
|isDelayIssue|bool|false||
|eventId|string|false||
|delayDay|string|false||
|delayIssueMsg|string|false||
|couponBagData|[][CouponBagData](#CouponBagData)|false||
|residueNum|string|false||
|appointmentStartAt|string|false||
|appointmentEndAt|string|false||
|goodsId|string|false||
|delayDayNum|string|false||
|isAutomaticAppointment|bool|false||
|automaticAppointmentMsg|string|false||
|autoPreOrgId|string|false||
|dayCount|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|awardNum|string|false||
|awardUnit|string|false||
|jumpConfig|string|false||
|surplusNum|string|false||
|payType|string|false||
|detailDesc|string|false||
|status|string|false||
|isDefaultNetAppointment|bool|false||
|defaultPreOrgId|string|false||
|isLuckDrawOneFrequency|bool|false||
|isOpenSpecialRule|string|false||

<details>
 <summary>
 <code>CouponBagData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|stockId|string|false||
|num|string|false||
|couponActId|string|false||
|id|string|false||
|name|string|false||
|amount|string|false||

</details>

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"name":"name",
"type":"type",
"money":"money",
"totalNum":"totalNum",
"isSpecial":true,
"imgId":"imgId",
"giftId":"giftId",
"sort":"sort",
"desc":"desc",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"eId":"eId",
"count":"count",
"limitArea":{
"id":"id",
"fullName":"fullName"
},
"limitProvince":"limitProvince",
"limitCity":"limitCity",
"limitDistrict":"limitDistrict",
"isSinceExtractOrgNum":true,
"prizeId":"prizeId",
"expendNum":"expendNum",
"costPrice":"costPrice",
"pointNum":"pointNum",
"couponActId":"couponActId",
"isDelayIssue":true,
"eventId":"eventId",
"delayDay":"delayDay",
"delayIssueMsg":"delayIssueMsg",
"couponBagData":[
	{
"stockId":"stockId",
"num":"num",
"couponActId":"couponActId",
"id":"id",
"name":"name",
"amount":"amount"
}
],
"residueNum":"residueNum",
"appointmentStartAt":"appointmentStartAt",
"appointmentEndAt":"appointmentEndAt",
"goodsId":"goodsId",
"delayDayNum":"delayDayNum",
"isAutomaticAppointment":true,
"automaticAppointmentMsg":"automaticAppointmentMsg",
"autoPreOrgId":"autoPreOrgId",
"dayCount":"dayCount",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"awardNum":"awardNum",
"awardUnit":"awardUnit",
"jumpConfig":"jumpConfig",
"surplusNum":"surplusNum",
"payType":"payType",
"detailDesc":"detailDesc",
"status":"status",
"isDefaultNetAppointment":true,
"defaultPreOrgId":"defaultPreOrgId",
"isLuckDrawOneFrequency":true,
"isOpenSpecialRule":"isOpenSpecialRule"
}
}
````



<a name="DirectReceivePrizeUser"></a>

## 前端method: InteractluckyactivityDirectReceivePrizeUser

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|prizeId|string|false||
|userId|string|false||
|marketerCodeId|string|false||
|wxUser|[WxUser](#WxUser)|false||
|isStaff|bool|false||
|ticketFileId|string|false||

<details>
 <summary>
 <code>WxUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|wxUnionId|string|false||
|wxMiniAppId|string|false||
|wxMiniOpenId|string|false||
|wxNickname|string|false||
|wxAvatar|string|false||
|wxMobileDim|string|false||
|realName|string|false||

</details>

### 请求例子:
````
{
"actId":"actId",
"prizeId":"prizeId",
"userId":"userId",
"marketerCodeId":"marketerCodeId",
"wxUser":{
"wxUnionId":"wxUnionId",
"wxMiniAppId":"wxMiniAppId",
"wxMiniOpenId":"wxMiniOpenId",
"wxNickname":"wxNickname",
"wxAvatar":"wxAvatar",
"wxMobileDim":"wxMobileDim",
"realName":"realName"
},
"isStaff":true,
"ticketFileId":"ticketFileId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizeInfo](#PrizeInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizeInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|name|string|false||
|type|string|false||
|money|string|false||
|totalNum|string|false||
|isSpecial|bool|false||
|imgId|string|false||
|giftId|string|false||
|sort|string|false||
|desc|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|imgData|[ImageData](#ImageData)|false||
|eId|string|false||
|count|string|false||
|limitArea|[Addre](#Addre)|false||
|limitProvince|string|false||
|limitCity|string|false||
|limitDistrict|string|false||
|isSinceExtractOrgNum|bool|false||
|prizeId|string|false||
|expendNum|string|false||
|costPrice|string|false||
|pointNum|string|false||
|couponActId|string|false||
|isDelayIssue|bool|false||
|eventId|string|false||
|delayDay|string|false||
|delayIssueMsg|string|false||
|couponBagData|[][CouponBagData](#CouponBagData)|false||
|residueNum|string|false||
|appointmentStartAt|string|false||
|appointmentEndAt|string|false||
|goodsId|string|false||
|delayDayNum|string|false||
|isAutomaticAppointment|bool|false||
|automaticAppointmentMsg|string|false||
|autoPreOrgId|string|false||
|dayCount|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|awardNum|string|false||
|awardUnit|string|false||
|jumpConfig|string|false||
|surplusNum|string|false||
|payType|string|false||
|detailDesc|string|false||
|status|string|false||
|isDefaultNetAppointment|bool|false||
|defaultPreOrgId|string|false||
|isLuckDrawOneFrequency|bool|false||
|isOpenSpecialRule|string|false||

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

<details>
 <summary>
 <code>CouponBagData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|stockId|string|false||
|num|string|false||
|couponActId|string|false||
|id|string|false||
|name|string|false||
|amount|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"name":"name",
"type":"type",
"money":"money",
"totalNum":"totalNum",
"isSpecial":true,
"imgId":"imgId",
"giftId":"giftId",
"sort":"sort",
"desc":"desc",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"eId":"eId",
"count":"count",
"limitArea":{
"id":"id",
"fullName":"fullName"
},
"limitProvince":"limitProvince",
"limitCity":"limitCity",
"limitDistrict":"limitDistrict",
"isSinceExtractOrgNum":true,
"prizeId":"prizeId",
"expendNum":"expendNum",
"costPrice":"costPrice",
"pointNum":"pointNum",
"couponActId":"couponActId",
"isDelayIssue":true,
"eventId":"eventId",
"delayDay":"delayDay",
"delayIssueMsg":"delayIssueMsg",
"couponBagData":[
	{
"stockId":"stockId",
"num":"num",
"couponActId":"couponActId",
"id":"id",
"name":"name",
"amount":"amount"
}
],
"residueNum":"residueNum",
"appointmentStartAt":"appointmentStartAt",
"appointmentEndAt":"appointmentEndAt",
"goodsId":"goodsId",
"delayDayNum":"delayDayNum",
"isAutomaticAppointment":true,
"automaticAppointmentMsg":"automaticAppointmentMsg",
"autoPreOrgId":"autoPreOrgId",
"dayCount":"dayCount",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"awardNum":"awardNum",
"awardUnit":"awardUnit",
"jumpConfig":"jumpConfig",
"surplusNum":"surplusNum",
"payType":"payType",
"detailDesc":"detailDesc",
"status":"status",
"isDefaultNetAppointment":true,
"defaultPreOrgId":"defaultPreOrgId",
"isLuckDrawOneFrequency":true,
"isOpenSpecialRule":"isOpenSpecialRule"
}
}
````



<a name="ReceiveRecordListUser"></a>

## 前端method: InteractluckyactivityReceiveRecordListUser

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|userId|string|false||
|page|string|false||
|pageSize|string|false||

### 请求例子:
````
{
"userId":"userId",
"page":"page",
"pageSize":"pageSize"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ReceiveRecordList](#ReceiveRecordList)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ReceiveRecordList </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|pageInfo|[PageInfo](#PageInfo)|false||
|list|[][ReceiveRecordInfo](#ReceiveRecordInfo)|false||

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

<details>
 <summary>
 <code>ReceiveRecordInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|marketerCodeId|string|false||
|userId|string|false||
|prizePoolId|string|false||
|prizeId|string|false||
|prizeName|string|false||
|prizeType|string|false||
|money|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|prizeImgId|string|false||
|prizeGiftId|string|false||
|prizeDesc|string|false||
|status|string|false||
|receivedAt|string|false||
|createdAt|string|false||
|orderNo|string|false||
|wxUser|[WxUser](#WxUser)|false||
|memberUser|[MemberUser](#MemberUser)|false||
|prizeImgData|[ImageData](#ImageData)|false||
|webUrl|string|false||
|distributionId|string|false||
|receiveUser|[ReceiveUser](#ReceiveUser)|false||
|receiveVerify|[ReceiveVerify](#ReceiveVerify)|false||
|logisticNumber|string|false||
|receiveStatus|string|false||
|webData|[Web](#Web)|false||
|actName|string|false||
|wxErrCode|string|false||
|wxErrMsg|string|false||
|couponActId|string|false||
|couponId|string|false||
|isDelayIssue|bool|false||
|delayIssueMsg|string|false||
|payMoney|string|false||
|payType|string|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>ReceiveUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveNameDim|string|false||
|receiveMobileDim|string|false||
|receiveAddress|string|false||
|receiveProvince|string|false||
|receiveCity|string|false||
|receiveDistrict|string|false||

</details>

<details>
 <summary>
 <code>ReceiveVerify </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveId|string|false||
|verifySn|string|false||
|verifyStatus|string|false||

</details>

<details>
 <summary>
 <code>Web </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|linkType|string|false||
|h5Link|string|false||

</details>

<details>
 <summary>
 <code>WxUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|wxUnionId|string|false||
|wxMiniAppId|string|false||
|wxMiniOpenId|string|false||
|wxNickname|string|false||
|wxAvatar|string|false||
|wxMobileDim|string|false||
|realName|string|false||

</details>

<details>
 <summary>
 <code>MemberUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|userMobileDim|string|false||
|userNameDim|string|false||
|userIdentityDim|string|false||
|userAddress|string|false||
|userProvince|string|false||
|userCity|string|false||
|userDistrict|string|false||
|userAddr|string|false||

</details>

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"list":[
	{
"id":"id",
"actId":"actId",
"marketerCodeId":"marketerCodeId",
"userId":"userId",
"prizePoolId":"prizePoolId",
"prizeId":"prizeId",
"prizeName":"prizeName",
"prizeType":"prizeType",
"money":"money",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"prizeImgId":"prizeImgId",
"prizeGiftId":"prizeGiftId",
"prizeDesc":"prizeDesc",
"status":"status",
"receivedAt":"receivedAt",
"createdAt":"createdAt",
"orderNo":"orderNo",
"wxUser":{
"wxUnionId":"wxUnionId",
"wxMiniAppId":"wxMiniAppId",
"wxMiniOpenId":"wxMiniOpenId",
"wxNickname":"wxNickname",
"wxAvatar":"wxAvatar",
"wxMobileDim":"wxMobileDim",
"realName":"realName"
},
"memberUser":{
"userMobileDim":"userMobileDim",
"userNameDim":"userNameDim",
"userIdentityDim":"userIdentityDim",
"userAddress":"userAddress",
"userProvince":"userProvince",
"userCity":"userCity",
"userDistrict":"userDistrict",
"userAddr":"userAddr"
},
"prizeImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"webUrl":"webUrl",
"distributionId":"distributionId",
"receiveUser":{
"receiveNameDim":"receiveNameDim",
"receiveMobileDim":"receiveMobileDim",
"receiveAddress":"receiveAddress",
"receiveProvince":"receiveProvince",
"receiveCity":"receiveCity",
"receiveDistrict":"receiveDistrict"
},
"receiveVerify":{
"receiveId":"receiveId",
"verifySn":"verifySn",
"verifyStatus":"verifyStatus"
},
"logisticNumber":"logisticNumber",
"receiveStatus":"receiveStatus",
"webData":{
"linkType":"linkType",
"h5Link":"h5Link"
},
"actName":"actName",
"wxErrCode":"wxErrCode",
"wxErrMsg":"wxErrMsg",
"couponActId":"couponActId",
"couponId":"couponId",
"isDelayIssue":true,
"delayIssueMsg":"delayIssueMsg",
"payMoney":"payMoney",
"payType":"payType"
}
]
}
}
````



<a name="ReceiveMoneyUser"></a>

## 前端method: InteractluckyactivityReceiveMoneyUser

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveId|string|false||
|userId|string|false||
|wxUser|[WxUser](#WxUser)|false||
|clientType|string|false||

<details>
 <summary>
 <code>WxUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|wxUnionId|string|false||
|wxMiniAppId|string|false||
|wxMiniOpenId|string|false||
|wxNickname|string|false||
|wxAvatar|string|false||
|wxMobileDim|string|false||
|realName|string|false||

</details>

### 请求例子:
````
{
"receiveId":"receiveId",
"userId":"userId",
"wxUser":{
"wxUnionId":"wxUnionId",
"wxMiniAppId":"wxMiniAppId",
"wxMiniOpenId":"wxMiniOpenId",
"wxNickname":"wxNickname",
"wxAvatar":"wxAvatar",
"wxMobileDim":"wxMobileDim",
"realName":"realName"
},
"clientType":"clientType"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ReceiveRecordInfo](#ReceiveRecordInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ReceiveRecordInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|marketerCodeId|string|false||
|userId|string|false||
|prizePoolId|string|false||
|prizeId|string|false||
|prizeName|string|false||
|prizeType|string|false||
|money|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|prizeImgId|string|false||
|prizeGiftId|string|false||
|prizeDesc|string|false||
|status|string|false||
|receivedAt|string|false||
|createdAt|string|false||
|orderNo|string|false||
|wxUser|[WxUser](#WxUser)|false||
|memberUser|[MemberUser](#MemberUser)|false||
|prizeImgData|[ImageData](#ImageData)|false||
|webUrl|string|false||
|distributionId|string|false||
|receiveUser|[ReceiveUser](#ReceiveUser)|false||
|receiveVerify|[ReceiveVerify](#ReceiveVerify)|false||
|logisticNumber|string|false||
|receiveStatus|string|false||
|webData|[Web](#Web)|false||
|actName|string|false||
|wxErrCode|string|false||
|wxErrMsg|string|false||
|couponActId|string|false||
|couponId|string|false||
|isDelayIssue|bool|false||
|delayIssueMsg|string|false||
|payMoney|string|false||
|payType|string|false||

<details>
 <summary>
 <code>WxUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|wxUnionId|string|false||
|wxMiniAppId|string|false||
|wxMiniOpenId|string|false||
|wxNickname|string|false||
|wxAvatar|string|false||
|wxMobileDim|string|false||
|realName|string|false||

</details>

<details>
 <summary>
 <code>MemberUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|userMobileDim|string|false||
|userNameDim|string|false||
|userIdentityDim|string|false||
|userAddress|string|false||
|userProvince|string|false||
|userCity|string|false||
|userDistrict|string|false||
|userAddr|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>ReceiveUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveNameDim|string|false||
|receiveMobileDim|string|false||
|receiveAddress|string|false||
|receiveProvince|string|false||
|receiveCity|string|false||
|receiveDistrict|string|false||

</details>

<details>
 <summary>
 <code>ReceiveVerify </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveId|string|false||
|verifySn|string|false||
|verifyStatus|string|false||

</details>

<details>
 <summary>
 <code>Web </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|linkType|string|false||
|h5Link|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"actId":"actId",
"marketerCodeId":"marketerCodeId",
"userId":"userId",
"prizePoolId":"prizePoolId",
"prizeId":"prizeId",
"prizeName":"prizeName",
"prizeType":"prizeType",
"money":"money",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"prizeImgId":"prizeImgId",
"prizeGiftId":"prizeGiftId",
"prizeDesc":"prizeDesc",
"status":"status",
"receivedAt":"receivedAt",
"createdAt":"createdAt",
"orderNo":"orderNo",
"wxUser":{
"wxUnionId":"wxUnionId",
"wxMiniAppId":"wxMiniAppId",
"wxMiniOpenId":"wxMiniOpenId",
"wxNickname":"wxNickname",
"wxAvatar":"wxAvatar",
"wxMobileDim":"wxMobileDim",
"realName":"realName"
},
"memberUser":{
"userMobileDim":"userMobileDim",
"userNameDim":"userNameDim",
"userIdentityDim":"userIdentityDim",
"userAddress":"userAddress",
"userProvince":"userProvince",
"userCity":"userCity",
"userDistrict":"userDistrict",
"userAddr":"userAddr"
},
"prizeImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"webUrl":"webUrl",
"distributionId":"distributionId",
"receiveUser":{
"receiveNameDim":"receiveNameDim",
"receiveMobileDim":"receiveMobileDim",
"receiveAddress":"receiveAddress",
"receiveProvince":"receiveProvince",
"receiveCity":"receiveCity",
"receiveDistrict":"receiveDistrict"
},
"receiveVerify":{
"receiveId":"receiveId",
"verifySn":"verifySn",
"verifyStatus":"verifyStatus"
},
"logisticNumber":"logisticNumber",
"receiveStatus":"receiveStatus",
"webData":{
"linkType":"linkType",
"h5Link":"h5Link"
},
"actName":"actName",
"wxErrCode":"wxErrCode",
"wxErrMsg":"wxErrMsg",
"couponActId":"couponActId",
"couponId":"couponId",
"isDelayIssue":true,
"delayIssueMsg":"delayIssueMsg",
"payMoney":"payMoney",
"payType":"payType"
}
}
````



<a name="ReceiveAddress"></a>

## 前端method: InteractluckyactivityReceiveAddress

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveId|string|false||
|userId|string|false||
|mobile|string|false||
|address|string|false||
|province|string|false||
|city|string|false||
|district|string|false||
|name|string|false||
|addressId|string|false||

### 请求例子:
````
{
"receiveId":"receiveId",
"userId":"userId",
"mobile":"mobile",
"address":"address",
"province":"province",
"city":"city",
"district":"district",
"name":"name",
"addressId":"addressId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="ReceiveChoiceOrg"></a>

## 前端method: InteractluckyactivityReceiveChoiceOrg

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveId|string|false||
|userId|string|false||
|orgId|string|false||
|appointmentAt|string|false||

### 请求例子:
````
{
"receiveId":"receiveId",
"userId":"userId",
"orgId":"orgId",
"appointmentAt":"appointmentAt"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ReceiveVerify](#ReceiveVerify)|false||

<details>
 <summary>
 <code>ReceiveVerify </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveId|string|false||
|verifySn|string|false||
|verifyStatus|string|false||

</details>

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"receiveId":"receiveId",
"verifySn":"verifySn",
"verifyStatus":"verifyStatus"
}
}
````



<a name="PlayUserList"></a>

## 前端method: InteractluckyactivityPlayUserList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||

### 请求例子:
````
{
"actId":"actId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ReceiveData](#ReceiveData)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ReceiveData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|pageInfo|[PageInfo](#PageInfo)|false||
|list|[][ReceiveArr](#ReceiveArr)|false||
|panel|[Panel](#Panel)|false||

<details>
 <summary>
 <code>ReceiveArr </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|userId|string|false||
|wxNickname|string|false||
|wxAvatar|string|false||
|wxMobile|string|false||
|createdAt|string|false||
|status|string|false||
|giftName|string|false||

</details>

<details>
 <summary>
 <code>Panel </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|useCount|string|false||

</details>

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"list":[
	{
"userId":"userId",
"wxNickname":"wxNickname",
"wxAvatar":"wxAvatar",
"wxMobile":"wxMobile",
"createdAt":"createdAt",
"status":"status",
"giftName":"giftName"
}
],
"panel":{
"useCount":"useCount"
}
}
}
````



<a name="TakeFromOrgIdList"></a>

## 前端method: InteractluckyactivityTakeFromOrgIdList

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|page|string|false||
|pageSize|string|false||
|longitude|string|false||
|latitude|string|false||
|name|string|false||
|userId|string|false||
|prizeId|string|false||
|orgId|string|false||

### 请求例子:
````
{
"actId":"actId",
"page":"page",
"pageSize":"pageSize",
"longitude":"longitude",
"latitude":"latitude",
"name":"name",
"userId":"userId",
"prizeId":"prizeId",
"orgId":"orgId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[TakeFromOrgIdList](#TakeFromOrgIdList)|false||

<details>
 <summary>
 <code>TakeFromOrgIdList </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|pageInfo|[PageInfo](#PageInfo)|false||
|list|[][TakeFromOrgIdData](#TakeFromOrgIdData)|false||

<details>
 <summary>
 <code>PageInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|page|string|false||
|pageSize|string|false||
|totalPage|string|false||
|total|string|false||

</details>

<details>
 <summary>
 <code>TakeFromOrgIdData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||
|address|string|false||
|distance|string|false||
|contactNumber|string|false||
|businessHours|string|false||
|icon|[ImageData](#ImageData)|false||

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

</details>

</details>

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"pageInfo":{
"page":"page",
"pageSize":"pageSize",
"totalPage":"totalPage",
"total":"total"
},
"list":[
	{
"id":"id",
"fullName":"fullName",
"address":"address",
"distance":"distance",
"contactNumber":"contactNumber",
"businessHours":"businessHours",
"icon":{
"id":"id",
"name":"name",
"url":"url"
}
}
]
}
}
````



<a name="ReceiveAddressEdit"></a>

## 前端method: InteractluckyactivityReceiveAddressEdit

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveId|string|false||
|userId|string|false||
|mobile|string|false||
|address|string|false||
|province|string|false||
|city|string|false||
|district|string|false||
|name|string|false||
|addressId|string|false||

### 请求例子:
````
{
"receiveId":"receiveId",
"userId":"userId",
"mobile":"mobile",
"address":"address",
"province":"province",
"city":"city",
"district":"district",
"name":"name",
"addressId":"addressId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="ReceiveInfoGetUser"></a>

## 前端method: InteractluckyactivityReceiveInfoGetUser

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveId|string|false||
|userId|string|false||

### 请求例子:
````
{
"receiveId":"receiveId",
"userId":"userId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[ReceiveRecordInfo](#ReceiveRecordInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>ReceiveRecordInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|marketerCodeId|string|false||
|userId|string|false||
|prizePoolId|string|false||
|prizeId|string|false||
|prizeName|string|false||
|prizeType|string|false||
|money|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|prizeImgId|string|false||
|prizeGiftId|string|false||
|prizeDesc|string|false||
|status|string|false||
|receivedAt|string|false||
|createdAt|string|false||
|orderNo|string|false||
|wxUser|[WxUser](#WxUser)|false||
|memberUser|[MemberUser](#MemberUser)|false||
|prizeImgData|[ImageData](#ImageData)|false||
|webUrl|string|false||
|distributionId|string|false||
|receiveUser|[ReceiveUser](#ReceiveUser)|false||
|receiveVerify|[ReceiveVerify](#ReceiveVerify)|false||
|logisticNumber|string|false||
|receiveStatus|string|false||
|webData|[Web](#Web)|false||
|actName|string|false||
|wxErrCode|string|false||
|wxErrMsg|string|false||
|couponActId|string|false||
|couponId|string|false||
|isDelayIssue|bool|false||
|delayIssueMsg|string|false||
|payMoney|string|false||
|payType|string|false||

<details>
 <summary>
 <code>WxUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|wxUnionId|string|false||
|wxMiniAppId|string|false||
|wxMiniOpenId|string|false||
|wxNickname|string|false||
|wxAvatar|string|false||
|wxMobileDim|string|false||
|realName|string|false||

</details>

<details>
 <summary>
 <code>MemberUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|userMobileDim|string|false||
|userNameDim|string|false||
|userIdentityDim|string|false||
|userAddress|string|false||
|userProvince|string|false||
|userCity|string|false||
|userDistrict|string|false||
|userAddr|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>ReceiveUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveNameDim|string|false||
|receiveMobileDim|string|false||
|receiveAddress|string|false||
|receiveProvince|string|false||
|receiveCity|string|false||
|receiveDistrict|string|false||

</details>

<details>
 <summary>
 <code>ReceiveVerify </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|receiveId|string|false||
|verifySn|string|false||
|verifyStatus|string|false||

</details>

<details>
 <summary>
 <code>Web </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|linkType|string|false||
|h5Link|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"actId":"actId",
"marketerCodeId":"marketerCodeId",
"userId":"userId",
"prizePoolId":"prizePoolId",
"prizeId":"prizeId",
"prizeName":"prizeName",
"prizeType":"prizeType",
"money":"money",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"prizeImgId":"prizeImgId",
"prizeGiftId":"prizeGiftId",
"prizeDesc":"prizeDesc",
"status":"status",
"receivedAt":"receivedAt",
"createdAt":"createdAt",
"orderNo":"orderNo",
"wxUser":{
"wxUnionId":"wxUnionId",
"wxMiniAppId":"wxMiniAppId",
"wxMiniOpenId":"wxMiniOpenId",
"wxNickname":"wxNickname",
"wxAvatar":"wxAvatar",
"wxMobileDim":"wxMobileDim",
"realName":"realName"
},
"memberUser":{
"userMobileDim":"userMobileDim",
"userNameDim":"userNameDim",
"userIdentityDim":"userIdentityDim",
"userAddress":"userAddress",
"userProvince":"userProvince",
"userCity":"userCity",
"userDistrict":"userDistrict",
"userAddr":"userAddr"
},
"prizeImgData":{
"id":"id",
"name":"name",
"url":"url"
},
"webUrl":"webUrl",
"distributionId":"distributionId",
"receiveUser":{
"receiveNameDim":"receiveNameDim",
"receiveMobileDim":"receiveMobileDim",
"receiveAddress":"receiveAddress",
"receiveProvince":"receiveProvince",
"receiveCity":"receiveCity",
"receiveDistrict":"receiveDistrict"
},
"receiveVerify":{
"receiveId":"receiveId",
"verifySn":"verifySn",
"verifyStatus":"verifyStatus"
},
"logisticNumber":"logisticNumber",
"receiveStatus":"receiveStatus",
"webData":{
"linkType":"linkType",
"h5Link":"h5Link"
},
"actName":"actName",
"wxErrCode":"wxErrCode",
"wxErrMsg":"wxErrMsg",
"couponActId":"couponActId",
"couponId":"couponId",
"isDelayIssue":true,
"delayIssueMsg":"delayIssueMsg",
"payMoney":"payMoney",
"payType":"payType"
}
}
````



<a name="PrizeInfoGet"></a>

## 前端method: InteractluckyactivityPrizeInfoGet

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||

### 请求例子:
````
{
"id":"id"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[PrizeInfo](#PrizeInfo)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>PrizeInfo </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|actId|string|false||
|poolId|string|false||
|name|string|false||
|type|string|false||
|money|string|false||
|totalNum|string|false||
|isSpecial|bool|false||
|imgId|string|false||
|giftId|string|false||
|sort|string|false||
|desc|string|false||
|stockId|string|false||
|mchId|string|false||
|equityId|string|false||
|equityType|string|false||
|imgData|[ImageData](#ImageData)|false||
|eId|string|false||
|count|string|false||
|limitArea|[Addre](#Addre)|false||
|limitProvince|string|false||
|limitCity|string|false||
|limitDistrict|string|false||
|isSinceExtractOrgNum|bool|false||
|prizeId|string|false||
|expendNum|string|false||
|costPrice|string|false||
|pointNum|string|false||
|couponActId|string|false||
|isDelayIssue|bool|false||
|eventId|string|false||
|delayDay|string|false||
|delayIssueMsg|string|false||
|couponBagData|[][CouponBagData](#CouponBagData)|false||
|residueNum|string|false||
|appointmentStartAt|string|false||
|appointmentEndAt|string|false||
|goodsId|string|false||
|delayDayNum|string|false||
|isAutomaticAppointment|bool|false||
|automaticAppointmentMsg|string|false||
|autoPreOrgId|string|false||
|dayCount|string|false||
|actEventData|[][ActEventData](#ActEventData)|false||
|awardNum|string|false||
|awardUnit|string|false||
|jumpConfig|string|false||
|surplusNum|string|false||
|payType|string|false||
|detailDesc|string|false||
|status|string|false||
|isDefaultNetAppointment|bool|false||
|defaultPreOrgId|string|false||
|isLuckDrawOneFrequency|bool|false||
|isOpenSpecialRule|string|false||

<details>
 <summary>
 <code>ActEventData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|eventType|string|false||
|eventId|string|false||

</details>

<details>
 <summary>
 <code>ImageData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|name|string|false||
|url|string|false||

</details>

<details>
 <summary>
 <code>Addre </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|id|string|false||
|fullName|string|false||

</details>

<details>
 <summary>
 <code>CouponBagData </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|stockId|string|false||
|num|string|false||
|couponActId|string|false||
|id|string|false||
|name|string|false||
|amount|string|false||

</details>

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"id":"id",
"actId":"actId",
"poolId":"poolId",
"name":"name",
"type":"type",
"money":"money",
"totalNum":"totalNum",
"isSpecial":true,
"imgId":"imgId",
"giftId":"giftId",
"sort":"sort",
"desc":"desc",
"stockId":"stockId",
"mchId":"mchId",
"equityId":"equityId",
"equityType":"equityType",
"imgData":{
"id":"id",
"name":"name",
"url":"url"
},
"eId":"eId",
"count":"count",
"limitArea":{
"id":"id",
"fullName":"fullName"
},
"limitProvince":"limitProvince",
"limitCity":"limitCity",
"limitDistrict":"limitDistrict",
"isSinceExtractOrgNum":true,
"prizeId":"prizeId",
"expendNum":"expendNum",
"costPrice":"costPrice",
"pointNum":"pointNum",
"couponActId":"couponActId",
"isDelayIssue":true,
"eventId":"eventId",
"delayDay":"delayDay",
"delayIssueMsg":"delayIssueMsg",
"couponBagData":[
	{
"stockId":"stockId",
"num":"num",
"couponActId":"couponActId",
"id":"id",
"name":"name",
"amount":"amount"
}
],
"residueNum":"residueNum",
"appointmentStartAt":"appointmentStartAt",
"appointmentEndAt":"appointmentEndAt",
"goodsId":"goodsId",
"delayDayNum":"delayDayNum",
"isAutomaticAppointment":true,
"automaticAppointmentMsg":"automaticAppointmentMsg",
"autoPreOrgId":"autoPreOrgId",
"dayCount":"dayCount",
"actEventData":[
	{
"eventType":"eventType",
"eventId":"eventId"
}
],
"awardNum":"awardNum",
"awardUnit":"awardUnit",
"jumpConfig":"jumpConfig",
"surplusNum":"surplusNum",
"payType":"payType",
"detailDesc":"detailDesc",
"status":"status",
"isDefaultNetAppointment":true,
"defaultPreOrgId":"defaultPreOrgId",
"isLuckDrawOneFrequency":true,
"isOpenSpecialRule":"isOpenSpecialRule"
}
}
````



<a name="RechargeCallback"></a>

## 前端method: InteractluckyactivityRechargeCallback

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|data|string|false||

### 请求例子:
````
{
"data":"data"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
}
}
````



<a name="GetVerifyCardDetail"></a>

## 前端method: InteractluckyactivityGetVerifyCardDetail

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|actId|string|false||
|userId|string|false||

### 请求例子:
````
{
"actId":"actId",
"userId":"userId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[VerifyCardStatus](#VerifyCardStatus)|false||

<details>
 <summary>
 <code>VerifyCardStatus </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|timeClockStatus|string|false||
|isVerifyCard|bool|false||

</details>

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"timeClockStatus":"timeClockStatus",
"isVerifyCard":true
}
}
````



<a name="CreateVerifyCardOrder"></a>

## 前端method: InteractluckyactivityCreateVerifyCardOrder

## 请求方式: **POST,GET**

### 请求参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|wxUser|[WxUser](#WxUser)|false||
|actId|string|false||
|userId|string|false||

<details>
 <summary>
 <code>WxUser </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|wxUnionId|string|false||
|wxMiniAppId|string|false||
|wxMiniOpenId|string|false||
|wxNickname|string|false||
|wxAvatar|string|false||
|wxMobileDim|string|false||
|realName|string|false||

</details>

### 请求例子:
````
{
"wxUser":{
"wxUnionId":"wxUnionId",
"wxMiniAppId":"wxMiniAppId",
"wxMiniOpenId":"wxMiniOpenId",
"wxNickname":"wxNickname",
"wxAvatar":"wxAvatar",
"wxMobileDim":"wxMobileDim",
"realName":"realName"
},
"actId":"actId",
"userId":"userId"
}
````
### 响应参数:


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|meta|[MetaRes](#MetaRes)|false||
|data|[CreateVerifyCardOrderResp_Data](#CreateVerifyCardOrderResp_Data)|false||

<details>
 <summary>
 <code>MetaRes </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|requestId|string|false||
|code|int32|false||
|msg|string|false||
|errMsg|string|false||

</details>

<details>
 <summary>
 <code>CreateVerifyCardOrderResp_Data </code>
  </summary>


|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |
|timestamp|string|false||
|nonceStr|string|false||
|package|string|false||
|signType|string|false||
|paySign|string|false||
|orderId|string|false||
|orderGoodsId|string|false||
|appId|string|false||

</details>

### 响应例子:
````
{
"meta":{
"requestId":"requestId",
"code":1,
"msg":"msg",
"errMsg":"errMsg"
},
"data":{
"timestamp":"timestamp",
"nonceStr":"nonceStr",
"package":"package",
"signType":"signType",
"paySign":"paySign",
"orderId":"orderId",
"orderGoodsId":"orderGoodsId",
"appId":"appId"
}
}
````
