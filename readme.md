# old south

### 项目说明
OldSouth ，南方老头，即南极仙翁，寿星。
项目用来记录，分析身体各项指标，为了让自己有质量的多活几年，所以用了寿星做项目名。

项目分为几个部分。
第一部分是，通过不同输入端记录数据，上传到服务器。预计使用苹果的快捷指令+健康协同手表做，通过手表或其他健康设备能监测的身体数据直接通过自动化发送给服务端，血糖之类的第三方数据通过手动输入的方式发送给服务端。   
第二部分是服务端，提供输入接口，输出接口，数据处理分析等功能。   
第三部分是展示端，使用esp32作为展示端，显示数据。   


### 功能列表
* server 
- [ ] 记录
- [ ] 自动补全无结果数据，自动填充为上一次的数据
- [ ] 分析整理   
- [ ] ~~生成图片~~
- [ ] ~~发送图片给client~~

* esp32
- [ ] 实时接收图片
- [ ] 接收到新的图片后，刷新展示

### Meta数据代码说明
|代码|全拼|翻译|
|-|-|-|
|FBG|fasting blood glucose|空腹血糖|
|PBG|postprandial blood glucose|餐后血糖|
|W|weight|体重|
|BFR|Body fat rate|体脂率|
|E3|exercise energy expenditure|运动能量消耗|
|DoE|duration of exercise|运动时长|
|RHR|resting heart rate|静息心率|
|EHR|exercise heart rate|运动心率|
|ST|sleep time|睡眠时长|
|FR|food rate|食物比例，对比住院伙食|
|DP|diastolic (blood) pressure|舒张压|
|SP|systolic (blood) pressure|收缩压|

