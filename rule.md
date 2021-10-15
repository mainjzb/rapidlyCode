<!-- TOC -->

- [规则](#规则)
- [长句翻译](#长句翻译)
- [等级](#等级)
- [区域](#区域)
- [地名 NPC](#地名-npc)
- [宠物](#宠物)
- [圣诞](#圣诞)
- [物品](#物品)
- [杂项+属性](#杂项属性)
- [语法](#语法)
- [万圣节](#万圣节)
- [装备](#装备)
- [职业](#职业)
- [倾向](#倾向)
- [冒险岛](#冒险岛)
- [属性](#属性)
- [装备_单](#装备_单)
- [temple](#temple)
- [BOSS](#boss)
- [单](#单)
- [单.](#单)

<!-- /TOC -->
&nbsp;(\S)== \1
’=='
‘=='

@import "./rule_stuff.md"
@import "./rule_cash.md"
@import "./rule_event.md"
@import "./rule_NPC.md"
@import "./rule_equipment.md"
@import "./rule_other.md"
@import "./rule_skill.md"

# 规则
	#  '#'和'-'开头是注释,会被急速代码忽视
	#  '=='左边是原文，右边是翻译
	#  '==' 左边支持正则
	#  首字母小写，自动匹配大写版本
	#  首字母小写的单词，自动忽略匹配上的单词首字母大小写。

	#  左边规则
	#  任何正则规则都能使用例如下
	#  \d 匹配一个数字
	#  ?  前面的内容可有可无

	#  右边规则
	#  \1 表示左边第一个被匹配的内容
	
	#  <code>content</code> content不会被翻译



#  长句翻译
	Grants durational stats==
	After your character reaches Lv.10, they will begin to earn two additional levels every time your character gains a level up to Lv.149.==10-148级之间，每升一级，则获得额外2级。
	Use this scroll to give Epic Potential to items that are ranked Rare or below==给B潜能或无潜能的装备进化为A潜能
	Zero characters must have completed up to Chapter 2==神之子必须完成第二章
	20 levels below and 20 levels above==20级以内
	2:00 AM PDT / 5:00 AM EDT / 11:00 AM CEST / 7:00 PM AEST==北京时间下午5点
	5:00 PM PDT / 8:00 PM EDT / 2:00 AM CEST / 10:00 AM AEST==北京时间早上8点
	12:59 AM PDT / 3:59 AM EDT / 9:59 AM CEST / 5:59 PM AEST==北京时间下午4点
	4:00 PM PST / 7:00 PM EST / 1:00 AM CET / 11:00 AM AEDT==北京时间早上8点
	3:59 PM PST / 6:59 PM EST / 12:59 AM CET / 10:59 AM AEDT==次日早上8点
	\(12\:00 AM UTC\)==(北京早上8点)
	All you have to do is log in to start receiving great items==你所要做的是登录游戏领取奖励！
	Use coupon to receive permanent item.==双击消耗栏卡卷获得物品。
	Coupon has 7-day duration.==卡卷保留7天。
	Receive 50% off when you do Potential appraisals or resets==装备潜能在鉴定或重置时花费的金币享受五折
	The chance to receive new monsters for your Monster Collection is doubled==怪物收藏在收藏新怪物时概率翻倍
	Potential appraisals or resets==装备潜能在鉴定或重置
	Receive 50% off when you do Ability resets==内在能力花费荣誉半价
	Can be received once per world==每个大区都可以收到一次
	(\d\d?)-day stat duration==属性持续\1天
	stat duration lasts for (\d\d?) days==属性持续\1天
	around your level==你的等级范围内
	You have until.*to purchase these items==你可以在活动时间内购买物品 
	Clear Pollo and Fritto once==完成一次保罗或普利托
	Defeat Elite Boss once==击败一次精英BOSS
	Participate in fighting the Inferno Wolf==参加一次地狱火焰狼
	Entry count limited to once daily. Clear count limited to once per week.==进入次数限每天一次，击杀次数限每周一次。
	Entry count limited to once daily. Clear count limited to once per month.==进入次数限每天一次，击杀次数限每月一次。
	Rewards can be claimed once==领取一次奖励
	as you burn your way along==在斗燃的路上
	Receive 2 stars when you successfully do Star Force 10 enhancements and lower==10星或以下强化成功时获得2星
	Receive 3x the amount of Surprise Missions==临时任务可完成次数提高3倍

# 等级 
	Lv. (\d?\d?\d) - (\d?\d?\d)==\1级 - \2级
	Lv. (\d?\d?\d) - (\d?\d?\d)==\1级 - \2级
	5th (J|j)ob (A|a)dvancement==5转
	Req. Lv: ([1-9]?\d\d)==\1级以上
	Lv. ([1-9]?\d\d) and above==\1级以上
	Lv. ([1-9]?\d?\d)==\1级
	Lv.(\d\d\d)==\1级
	(\d)(\w\w) (J|j)ob==\1转
	Lv.==等级

# 区域
	non-Reboot worlds only==仅限普通区
	non-Reboot worlds=="普通区"
	Non-Reboot Worlds=="普通区"
	
	Reboot \(EU\) or Reboot \(NA\) worlds==R区
	Reboot worlds? only==<span kdclassjsq="notranslate">仅限R区</span>
	Reboot Worlds?=="R区"
	Reboot worlds?=="R区"
	Reboot \(NA\)==<span kdclassjsq="notranslate">美服R区</span>
	REBOOT ONLY==仅限R区
	Reboot=="R区"
	All worlds==所有区
	\[All Worlds\]==
	\(EU\)==(欧服)
	\(NA\)==(美服)
	\bMaplers?\b==冒险家
	\bV.(\d\d\d)\b==\1版本
	Nexon Launcher==NX登录器
	NON-REBOOT ONLY==仅限普通区


# 地名 NPC

Arcane River Express Pass==奥术河任务快速完成界面
Evolving Room Coin Shop's Vendor-C2==进化研究所C2商店
Prankster Stormwing==雪地大作战
Dimensional Mirror==次元之境
Fairy Bros Daily Gifts==FB每日礼物
Party Point Merchant Coin Shop==组队积分商店
Party Point Merchant==组队积分商店
Partem Ruins==帕坦遗址
Immortal Gorgons==不朽的美杜莎
Tenebris==泰涅布里斯(789岛)
Vanishing Journey==消失的旅途(1岛)
Hungry Muto==饥饿的莫托(岛二每日)
Dream Defender==梦都防御战(岛三每日)
Arcane River==奥术河
Nameless Town==无名村
Yum Yum Island==真香岛
Chu Chu Island==啾啾岛(2岛)
Chew Chew Island==啾啾岛(2岛)
Lachelein==梦之都拉克兰(3岛)
Lachelein Clocktower==拉克兰钟楼(3岛)
Arcana==阿尔卡那(4岛)
Morass==莫拉斯(5岛)
Esfera==埃斯佩拉(6岛)
Sellas==塞拉斯(6.5岛)
Crimsonheart Castle==狮子王城堡
Monster Park REBORN Coupon==怪物公园优惠券
Monster Park==怪物公园
Ghost Park==鬼魂公园
Pollo/Fritto==保罗/普利托
Pollo and Fritto Entry Ticket==<img src="upload/attach/202006/2_TBX3AAR2V5DTWH6.png">保罗和普利托入场卷
Pollo and Fritto==保罗和普利托
Nett's Pyramid==奈特的金字塔
Nett’s Pyramid==奈特的金字塔
Dimension Invasion==次元入侵
Dimension Invade==次元入侵
Fox Valley==狐狸谷
Savage Terminal==野蛮终点站
Monster Collection==<span kdclassjsq="notranslate">怪物收藏</span>
Gold Beach==金海岸
Six Path Crossway==六岔路口
Temple of Time==时间神殿
Heroes of Maple==冒险英雄
Black Heaven==黑色天堂
Oz Tower==起源之塔
Tower of OZ Lobby==起源之塔大厅
Tower of OZ==起源之塔
Tower of Oz==起源之塔
Seed Tower==起源之塔
Ninja Castle==忍者城堡
Lith Harbor==明珠港
Tru's Info Shop==特鲁的情报商店（明珠港）
Tangyoon's Food==唐云的料理
Tang Yoon's Cooking Class==唐云的料理班
Tang Yoon==唐云
Alliance Supply Depot==同盟供应商
Maple Alliance==冒险同盟
Romeo and Juliet==<img src="upload/attach/202107/2_SSQKVQ443ZGEEFK.png"><span kdclassjsq="notranslate">罗密欧与朱丽叶</span>
Maple Union==联盟
Xerxes of Chryse==克里塞的薛西斯
Dimensional Crack==<img src="upload/attach/202107/2_NSX7UF3XBCB34J2.png"><span kdclassjsq="notranslate">玩具城组队任务</span>
Pirate Davy Jones==<img src="upload/attach/202107/2_2T7VN6T8SRYCJNZ.png"><span kdclassjsq="notranslate">海盗船组队任务</span>
Lord Pirate==<img src="upload/attach/202107/2_2T7VN6T8SRYCJNZ.png"><span kdclassjsq="notranslate">海盗船组队任务</span>
Dragon Rider==御龙魔
Kenta in Danger==<img src="upload/attach/202107/2_FBTWBZEWPCQ9U37.png"><span kdclassjsq="notranslate">坎特组队任务</span>
The Hoblin King==<img src="upload/attach/202107/2_Z7J9CSEZ7QWY44B.png"><span kdclassjsq="notranslate">侏儒帝王组队任务</span>
Resurrection of the Hob King==侏儒帝王的复活
Witch Shop==魔女商店
Twilight Perion==黄昏的勇士部落
Henesys Ruins==破坏的射手村
Herb Town==百草堂
Origin Sea==起源之海
Flag Race==公会跑旗赛
Grand Athenaeum==次元图书馆
Meso Market==金币市场
Pilot Irvin==飞行员艾文
Athena Pierce==皮尔斯
Lady Blair==布莱尔女士
Small Spirits?==小不点精灵
Tree Spirits?==树精灵
Rock Spirits?==石精灵
Marvel Machine==抽奖机
Labyrinth of Suffering==痛苦迷宫(8岛)
Limina\b==黎曼(9岛)
Deep Core==深渊核心
Corrupted Warriors?==腐败战士
Dark Ravine Portal==黑暗山谷光口
Dark Ravine portal==黑暗山谷光口
Dark Lords of Darkness==黑暗领主
Dark Ravine==黑暗山谷
Aura of Life==生命光环
Commerci Dungeon==航海副本
Commerci Merchant Union Voyage==组航副本
Bounty Hunter==赏金猎人
Fairy Academy==艾利涅妖精学院
Riena Strait==列娜海峡
Mushroom Shrine==蘑菇神社
Erda Spectrum==艾尔达的光谱(岛一组队)
Karuppa Town==卡尔帕镇
Cheong-woon Valley==青云谷
Twisted Aquarium==<span kdclassjsq="notranslate">镜像水下世界</span>
Showa Town==昭和镇
Dimensional Gate==次元门
Spirit Savior==拯救精灵(岛四每日)
Allian Base==外星人基地
New Leaf City==新叶城
Lulu Spinel==露露小姐姐
Reverse City==反转城市
Silent Crusade Shop==十字币商店
Crusader Coins==十字币
Hieizan Temple==比睿山寺
Kerning Jazz Bar==废弃都市爵士酒吧
Kerning City('s)?==废弃都市

Cernium==<span kdclassjsq="notranslate">Cernium</span>
Blackgate==黑门
Masteria==克拉奇亚
Crimsonwood==守护者的要塞
Anima==达恩维尔
Rien==里恩
Showa==昭和
Alishan==阿里山
Taotie==饕餮
Neinheart==南哈特
\bMikagami\b==粉色鲸鱼
\bNina\b==妮娜
\bKaruppa\b==卡尔帕
\bElodin\b==艾洛丁
\bGloom\b==黑暗之眼
\bHieizan\b==比睿山
Momijigaokan==枫叶山丘
Darknell==黑暗尼尔
Daknell==黑暗尼尔
Djunkel==黑暗尼尔
Ranmaru==森兰丸
Elinel==艾利涅
Henesys==射手村
Leafre==神木村
Escape\b==<img src="upload/attach/202107/2_QGZKEJYWJBYAYWG.png"><span kdclassjsq="notranslate">逃脱组队任务</span>
Kritias==克城
Alliance==同盟
Magatia==玛加提亚
Ariant==阿里安特
Tangyoon==唐云
Verdel==贝尔达
Edelstein==埃德尔斯坦
Orbis==天空之城
Ellinia==魔法密林
Ellinel==艾利涅
Grandis==格兰蒂斯
Ellin==艾利
Mu Lung Dojo==武陵道场
Mu Lung Points==道场积分
Mu Lung==武陵
Dojo==道场
Cassandra==卡珊德拉
Ludibrium==玩具城
Kerning==废弃都市
Tru\s==特鲁
Shinsoo==神兽
Glinda==格琳达
Moonbridge==月桥(7岛)
Aquarium==水下世界
\bHaven(?<!s)\b==避风港
Sleepywood==林中之城
Perion==勇士部落
Ardentmill==匠人街
Nautilus==诺特勒斯
Darkness==黑暗使者
\bZipangu\b==日本区域
\bWonky\b==妖精温莉
\bMoonbeam\b==阿琅

# 宠物

Expanded Auto Move Skill==扩展移动技能
Auto Move Skill==自动移动技能
Auto Feed and Movement Skill==自动喂食和移动技能
Expired Pickup Skill==额外拾取技能
Ignore Item Skill==忽略物品技能
Auto Buff Skill==自动BUFF技能
Auto HP Potion Skill==自动HP技能
Premium Pet Food==高级宠物食品
Permanent when applied==使用后永久学会
Pet Boost Package==宠物技能包
Silver Husky==银色哈士奇
Black Kitty Pet==黑色猫咪宠物
Lil Neiny Pet==迷你南哈特宠物
Jr\. Orchid==迷你奥尔卡
Jr\. Hilla==迷你希腊
Jr\. Von Leon==迷你狮子王
Harp Seal==小海豹
Von Soup==半半
Black Bean==黑色品克缤
Pet Snack==宠物点心
Dye Coupon==染色卡
Maple Tour==枫叶之旅(MT)
Magic Wagon==魔法马车



# 物品

Maplehood Watch==冒险幸运日
Super Megaphone==超级喇叭
Yummy Meat==美味的肉
Pendant Slot Coupon==<img src="upload/attach/202006/2_BAMZUX9Y7ZWZVKK.png"><span kdclassjsq="notranslate">项链扩展槽租用卡</span>
Infinite Revitalizer==<img src="upload/attach/202006/2_VH94NZCSUUBDEDW.png"><span kdclassjsq="notranslate">疲劳恢复药水</span>
Final Form Main Color Change Coupon==<img src="upload/attach/202006/2_ZBZTNE548R6TDVG.png"><span kdclassjsq="notranslate">狂龙战士主颜色变更券</span>
Final Form Sub Color Change Coupon==<img src="upload/attach/202006/2_YJXHG3CAGCQXBUU.png"><span kdclassjsq="notranslate">狂龙战士辅助颜色变更券</span>
Final Form Color Reset Coupon==<img src="upload/attach/202006/2_NT95VR8KX87J3XU.png"><span kdclassjsq="notranslate">狂龙战士颜色重置券</span>
Premium Water of Life==宠物复活药水
Alliance Medal==同盟币
Mastery Box==<img src="upload/attach/202106/2_RDUNNXWKW3SVA8U.png">能手册礼盒
Weapon Box==武器礼盒
Android Hearts?==机器人心脏
Inkwell Coin==财富币
Manji Coin==一千万财富币
Gaga Coin==一亿财富币
Chaos Circulator==<img src="upload/attach/202011/2_5ZHAEDQPXWCKZCZ.png">黑色内在还原器
Legendary Circulator==<img src="upload/attach/202106/2_VJZW7NFDN3XSUGC.png">绿色内在还原器
Miracle Circulators?==<img src="upload/attach/202011/2_VUPDS65J7F8GF3U.png">内在还原器
Perfect Potential Stamps?==完美潜能印章
Gold Potential Stamps?==<img src="upload/attach/202006/2_8XDTYUW2DTQDJ9V.png"><span kdclassjsq="notranslate">金色潜能印章</span>
Silver Potential Stamps?==<img src="upload/attach/202106/2_8YFUH473WS7N4U8.png"><span kdclassjsq="notranslate">银色潜能印章</span>
Carlton's Mustache Awakening Stamps?==<img src="upload/attach/202006/2_TZYSZP5ZUXZCCNG.png"><span kdclassjsq="notranslate">卡腾的胡子的印章</span>
Basic Potential Stamps?==<img src="upload/attach/202106/2_QW65RWN3KBGXEVV.png"><span kdclassjsq="notranslate">初级潜能印章</span>
Intermediate Potential Stamps?==<img src="upload/attach/202106/2_N6UKWCPTQU72TGQ.png"><span kdclassjsq="notranslate">中级潜能印章</span>
Advanced Potential Stamps?==<img src="upload/attach/202006/2_TZYSZP5ZUXZCCNG.png"><span kdclassjsq="notranslate">高级潜能印章</span>
Superior Potential Stamps?==<img src="upload/attach/202006/2_TZYSZP5ZUXZCCNG.png"><span kdclassjsq="notranslate">卓越潜能印章</span>
Special Potential Stamp==<img src="upload/attach/202006/2_TZYSZP5ZUXZCCNG.png"><span kdclassjsq="notranslate">特殊潜能印章</span>
Potential Items==潜能物品
Potential Stamp==潜能印章
Twisted Time==扭曲的时间
Intense Power Crystal==蓝水晶
Arachno Coins==蜘蛛币
Stone Cobweb Droplets?==蜘蛛币碎片石
Mastery Book Set==<img src="upload/attach/202106/2_RDUNNXWKW3SVA8U.png">能手册礼包
Mastery Book 30==<img src="upload/attach/202006/2_ZPRYVX6M2YFSH6V.png">能手册30
Mastery Book 20==<img src="upload/attach/202006/2_EZ57QFADU9JQVBY.png">能手册20
Mastery Books?==能手册
Maplehood Watch Mechanical Heart==枫叶M机器人心脏
Mega Character Burninator==<img src="upload/attach/202110/2_3Z89J69DRQZB9Z3.png">斗燃药水
Tera Character Burninator==泰斗燃药水
nodestones?==核心
nodes?==核心
Trait Growth Potion==<img src="upload/attach/202006/2_7JWQWGPPCBH225A.png">倾向成长药水
Selective Trait Boost Potion==<img src="upload/attach/202006/2_7JWQWGPPCBH225A.png"><span kdclassjsq="notranslate">倾向药水</span>
Trait Boost Potion==<img src="upload/attach/202006/2_7JWQWGPPCBH225A.png"><span kdclassjsq="notranslate">倾向药水</span>
Trait Increase Potion==<img src="upload/attach/202006/2_7JWQWGPPCBH225A.png"><span kdclassjsq="notranslate">倾向药水</span>
Circulator==内在还原器
2x EXP==双倍经验
2x Exp==双倍经验
2x Drop( Rate)? Coupon==<img src="upload/attach/202007/2_Z8SZ6E4MH3UAB5C.png">双倍爆率卡
2x Drop( Rate)?==双倍爆率
Power Elixirs?==<img src="upload/attach/202006/2_PG775SSSWRUH8EC.png">超级药水
Pollo and Fritto Entry Ticket==保罗/普利托入场券
Equip Tab (\d) Slot Coupon==背包装备栏\1格扩展卷
Use Tab (\d) Slot Coupon==背包消耗栏\1格扩展卷
Set-Up Tab (\d) Slot Coupon==背包点装栏\1格扩展卷
Etc Tab (\d) Slot Coupon==背包其他栏\1格扩展卷
Storage Room (\d)-Slot Coupon==仓库\1格扩展券
Selective 8 Slot Coupon==<img src="upload/attach/202006/2_YQAYC6GSBYKKQ5J.png"><span kdclassjsq="notranslate">可选择8格扩展卷</span>
Selective (\d) Slot Coupon==可选择\1格扩展卷
Selective (\d)-Slot Coupon==可选择\1格扩展卷
Selective (\d)-Slot==可选择\1格扩展
(\d) Slot Coupon==\1格扩展卷
Beauty Album Hair Slot Coupons?==头发储存槽扩展卡
Beauty Album Face Slot Coupons?==面部储存槽扩展卡
Innocence Scroll==<img src="upload/attach/202006/2_VQKHPU39PJBS4CX.png"><span kdclassjsq="notranslate">纯真卷轴</span>
Core Gemstone==核心
Job Advancement Coin==<img src="upload/attach/202006/2_D9NM3PWGU2C732W.png"><span kdclassjsq="notranslate">转职币</span>
SP Reset Scroll==<img src="upload/attach/202011/2_7EK579UAEGT9MNR.png"><span kdclassjsq="notranslate">SP重置卷</span>
AP Reset Scroll==<img src="upload/attach/202006/2_AYFY7H3FQNZR6CM.png"><span kdclassjsq="notranslate">AP重置卷</span>
Pet Coupon==宠物卡
Mysterious Monsterbloom==<img src="upload/attach/202006/2_NYBFB3MRKNMGJEB.png"><span kdclassjsq="notranslate">神秘怪物袋</span>
Onyx Apple==蓝彩苹果
Café Latte==咖啡牛奶
Extra Strength Hard Sanitizer==强力消毒剂
Extra-Strength Hand Sanitizer==强力消毒剂
Premium Cologne==高级香水
Snake Bone Soup==蛇骨汤
Extra Strength Hand Sanitizer==强力洗手液
Premium Carrot Juice==优质胡萝卜汁
Premium Hair Wax==优质发蜡
Cherry Chocolate Cupcake==巧克力蛋糕
White Candy Cupcake==白色糖果蛋糕
Strawberry Cupcake==草莓蛋糕
All Cure Potion==万能药水
Gingerbread Man Cookie==姜饼男人饼干
Gingerbread Woman Cookie==姜饼女人饼干
Hyper Teleport Rock==<img src="upload/attach/202006/2_Y4C5J5WNGGPV4U6.png"><span kdclassjsq="notranslate">传送石</span>
Teleport Rock==<img src="upload/attach/202006/2_Y4C5J5WNGGPV4U6.png"><span kdclassjsq="notranslate">传送石</span>
Extreme Breakthrough Growth Potion==极限突破成长药水
Extreme Breakthrough==极限突破
Scroll for Pet Equipment for ATT==宠物装备物理攻击卷轴
Scroll for Pet Equipment for M. ATT==宠物装备魔法攻击卷轴
Pet Equipment Stat Transfer Scroll==宠物装备属性转移卷轴
Meso Pouch==钱袋
Appraisal Magnifying Glass==鉴定放大镜




# Failiar 

Familiar Slot Expansion Coupon==<span kdclassjsq="XMtooltip">怪怪卡槽扩展券<span kdclassjsq="XMtooltiptext"><img src="https://i.loli.net/2020/05/24/sUJMjWSyn5pdZ6r.png"/></span></span>
Suspicious Fauxmiliar Coupon==<span kdclassjsq="XMtooltip">奇怪的怪怪卡券<span kdclassjsq="XMtooltiptext"><img src="https://i.loli.net/2020/05/24/JPWTIx5wC6ZLkm2.png"/></span></span>
Familiar system==怪怪卡系统
Familiar Essence==<img src="upload/attach/202108/2_3QT7X9VEN97PFDU.png"><span kdclassjsq="notranslate">怪怪精油</span>
Red Familiar Cards?==<img src="upload/attach/202108/2_UZNT2UJHC5573B5.png"><span kdclassjsq="notranslate">红色怪怪卡</span>
Unique Familiar Booster Pack==U级怪怪增强包
Familiar Booster Pack==<img src="upload/attach/202106/2_7N7Y5T6WQPZ8P3R.png">怪怪增强包
Suspicious Fauxmiliar Card==怪异万能卡
Familiar Breakthrough Card==怪怪突破卡
Unique Booster Pack==U级增强包
Booster Pack==增强包
Familiars?==怪怪卡

# 杂项+属性  

Theme Dungeon==主题副本
prerequisite quest==前置任务
Rock-Paper-Scissors==剪刀石头布
Double Miracle Time==DMT
8-bit==像素
Captivating Fragments?==浓姬碎片
Combo EXP orbs==连击经验球
Lunar New Year==农历新年
Genesis Force==创世之力
Bumper Crop Gift==丰收的礼物 
Daily Quest==每日任务
Cash Cover Eye Accessory==眼镜时装
Cash Cover Hat==帽子时装
after maintenance==维护后
star event notifier==星星
Star Force enhancements==装备星之力强化
Star Force Enhancing==装备星之力强化
Star Force==装备星之力强化
Tradeable within account==账号内可交易
Cash Shop==商城
Teleport Scroll==传送卷轴
Arboreal Blessings==神树的祝福
(\d)x Drop (Rate)? Coupon==\1倍爆率卡
(\d)x EXP Coupon==\1倍经验卡
Drop buff==爆率Buff
Wee Moon Bunny==小月兔
Respawn Token==原地复活符
Hyper Stat Points==超级属性点
Respawn Pass==无限*原地复活符
Epic Potential Scroll==A潜能卷
Legendary Potential Scrolls?==ss潜能卷轴
Epic Potential==A潜能
Buff Freezer==Buff保留符
Safety Charm==死亡保护符
Safety Gem==无限*死亡保护符
Weapon DEF/Magic DEF ==防御力
Bonus Stats==火花属性
Spell Traces? enhancements==痕迹强化
Spell Traces? Coupon==<img src="upload/attach/202103/2_9EN7RZ2DAQ9SFYY.png">咒语痕迹
Spell Traces?==<img src="upload/attach/202103/2_9EN7RZ2DAQ9SFYY.png">咒语痕迹
Mysterious Meso Pouch==随机金币袋
Hot Week==热力周
Storage Room==仓库
a Character Name Change Coupon==改名卡
OX Quiz==“对错问答”
OX quiz==“对错问答”
Soul Shards?==灵魂碎片
Maple World==冒险世界
Mirror World==镜像世界
True Arachnid Reflection==真*蜘蛛之影
Discovery Arena==探索竞技场
Discovery Desert==探索沙漠
Explorer-class character==冒险家职业
Grab a party==组队
Lab Server Coins==LAB币
Lab Server==LAB服务器
Lab world==LAB世界
boost potions==强化药水
Boost Potion==强化药水
Hard Will==困难级威尔
Hard Mode==困难模式
star reduction chance==降低等级率
star maintain chance==保留星级率
item destruction chance==装备破坏率
success chance==的成功率
Boss Monster Damage==BOSS伤害
Enemy DEF Ignored==无视防御
Monster DEF Ignored==无视防御
Ignored Monster Defense==无视防御
Ignored Monster DEF==无视防御
Ignored Monster Def==无视防御
Ignore DEF==无视防御
Lucky Item Scroll==幸运卷轴
Mage Suit==法师套服
HP Cost==HP消耗
Ark Points==Ark点
El Nath==冰封雪域
1 person is in the party==单人在队伍中
Advanced Bonus==高级
Intermediate Bonus==中级
Basic Bonus==基础
All Stats==全属性
World Leap system==转区系统
Monster Life Gem Coupon==<img src="upload/attach/202006/2_M82A6CGPK3TB76U.png"><span kdclassjsq="notranslate">怪物家园门票</span>
Monster Life Gem x7 Coupon==<img src="upload/attach/202006/2_M82A6CGPK3TB76U.png"><span kdclassjsq="notranslate">怪物家园门票x7</span>
Monster Life Gem x7==<img src="upload/attach/202006/2_M82A6CGPK3TB76U.png"><span kdclassjsq="notranslate">怪物家园门票x7</span>
Monster Life==<img src="upload/attach/202006/2_M82A6CGPK3TB76U.png"><span kdclassjsq="notranslate">怪物家园系统</span>
Trait Items==倾向物品
Raid Boss==BOSS
character slots==角色槽
EXP Coupons?==经验卡
Armor Supply Box==防具箱
Weapon Supply Box==武器箱
Potion Pot==药剂罐
Cash Inventory Transfer==商城转移
Hair Color Coupons?==染发卡
Skin Coupon==肤色卡
Royal Hair Coupon==皇家理发卡
Royal Face Coupon==皇家整容卡
Face Coupon==整容卡
Hair Style Coupon==理发卡
Hair Coupon==理发卡
Maple Points?==枫叶点
Captain Vaga==瓦加舰长
Double Rewards==双倍奖励
equipment covers==武器点状
\(M\)==(男)
\(F\)==(女)
Bonus Potential==附加潜能
inventory slots==背包空间
Bonus Potentials==附加潜能
Orange Mushroom==花蘑菇
Appreciation Gift Box(es)?==感谢礼物盒
Gift Box(es)?==礼物盒
Miu Miu the Merchant==缪缪商店卡
Maple Administrator==管理员
Maple Admin==管理员
Maple Relay Tier (\d) Box==冒险岛接力第\1礼盒
Daily Maple Relay==每日冒险接力
Maple Relay==冒险接力
Tier (\d)==\1级
Flash Jump==二段跳
Sunny Sunday Perks?==阳光周日
Sunny Sundays?==阳光周日
Transfer Hammer==托德之锤
Surprise Style Box==幸运箱
Cardcaptor Sakura==百变小樱
Sakura Kinomoto==小樱
Kero-chan==“小可”
Tomoyo Daidouji==“知世”
Act (\d)==第\1章
Sengoku High==战国高校
Bug Brawl==昆虫大作战
Inferno Wolf==地狱火焰狼
Flame Wolf==地狱火焰狼
Combo Kill orbs==连击经验球
Unique Equipped Item==唯一装备
(\d\d?)-day duration==保留\1天
(\d\d?) day duration==保留\1天
Maple Rewards UI==RP积分界面
Reward Points==RP积分
Maple Missions?==枫之使命
Sengoku Class Supplementary Pass==战国高校额外通行证
Inner Abilities==内在能力
1 hr==1小时
Divine Intelligence==神圣的智慧
Black Mage's==黑魔法师的
Black Mage==黑魔法师
(\d) minutes==\1分钟
tradeable within account==可仓库转移
Gold(en)? Hammer==<img src="upload/attach/202006/2_UVBTC26NK6TF2ZW.png">金锤子
party members?==队员
Infinite Pawsibilities==无限可能
Infinite Pawsibilites==无限可能
the Multikill counter==多连杀
Combo Kill counters?==连续击杀
Maple Guide stamps==冒险向导的盖章
Maple Guides?==冒险向导
advancement prerequisites?==前置任务
party quests?==组队任务
V Matrix==V矩阵
Double Jump==二段跳
Receive 50% off==享受五折
50% off==五折
Elite monsters?==精英怪
Elite Monsters?==精英怪
Elite Boss==精英BOSS
Elite boss==精英BOSS
Bounty Hunting==赏金猎人
Magnificent Soul==伟大的灵魂
Daily Box==每日盒子
party leader==队长
Lantern Erdas?==灯笼艾尔达
drop items==爆出物品
leveling==升级
Level Up==升级
Hard Boss==困难级BOSS
Ability Point==属性点
Skill Points?==技能点
Content Skills?==生活技能
Ability Skills?==能力技能
Goddess Statue Skills?==女神状态技能
Stat Skills?==状态技能
Support Skills?==生活技能
World Merged Party Quest==世界组队任务
Party Quest Expert==组队任务专家
Party Quest==组队任务
Party Points==组队积分
Bonus EXP==额外经验
Mysterious Arcane Symbol Box==<img src="upload/attach/202106/2_4GQN87X3EPV7BF9.png">神秘的奥术球礼盒
Arcane Symbol Selector Coupon==<img src="upload/attach/202006/2_Z254HG9DKYR9A3H.png">奥术球选择卡
Arcane Symbol Selection Coupon==<img src="upload/attach/202006/2_Z254HG9DKYR9A3H.png">奥术球选择卡
Labyrinth Arcane Symbol Box==<img src="upload/attach/202106/2_4GQN87X3EPV7BF9.png">奥术球礼盒
Arcane Symbol==<img src="upload/attach/202006/2_Z254HG9DKYR9A3H.png">奥术球
Selector Coupon==选择卡
Stage (\d)==\1阶
Combo Kill Orbs==连续击杀经验球
Combo Kills?==连续击杀
Battle Analysis==战斗分析
Battle Statistics UI==战斗统计界面
Arcane Power==奥术值
Relay Master Box(es)?==接力大师盒子
a party of (\d)-(\d)==\1-\2人的队伍
Boss Map Death==BOSS地图死亡
Etc tab==其他栏
Big Spender==大土豪
Hot Time==热力时间
Drop Rates?==爆率
Hot Days==令人激动的日子
Hard mode==困难模式
Fever Time==热力时间
Hot'n'Ready==速食
Surprise Missions==临时任务
Atmospheric Effect==特效
Pre-Creation==预创建
Hyper Stats==超级属性
Fairy Bros==FB每日
primary stat==<span kdclassjsq="notranslate">主属性</span>
master level==最高等级

# 语法
	to trade in for==换取
	dropped by (\w+)==通过\1掉落
	a special deal==优惠
    Hong Bao==红包
	logging out==退出登录




# 万圣节
	Zombuddy==僵尸
	Ghost Coins?==鬼币
	HALLOWEEN==万圣节
	Pumpkin Zombies?==南瓜僵尸
	Hallowbat==万圣节
	Hallowkitty==幽灵猫
	Advanced Hasty Hunting Box(es)?==高级冬季狩猎盒
	Superior Hasty Hunting Box(es)?==超级冬季狩猎盒
	Basic Hasty Hunting Box(es)?==初级冬季狩猎盒
	Hasty Hunting==冬季狩猎
	Frankenbalrogs==蝙蝠魔僵尸
	Ghoulbusters Coin Shop==鬼币兑换商店
	Ghoulbusters==僵尸猎杀
	Ghoulbusting==僵尸狩猎
	Worn Witch==破旧的女巫
	Zombie==僵尸



# 装备

New Gold Heart==黄金心脏
Eternal Flame Title==<img src="upload/attach/202106/2_3K85RT2PPUPTGUM.png"><span kdclassjsq="notranslate">永恒火焰称号</span>
Eternal Flame Ring==<span kdclassjsq="notranslate"><img src="upload/attach/202107/2_8Y9E6C25RURQWJU.png">永恒火焰戒指</span>
Eternal Flame==<span kdclassjsq="notranslate">永恒火焰</span>
Rudolph's Red Nose==鲁道夫的红鼻子
Maple Anniversary Equipments==冒险岛周年庆装备
Cubic Blade==<img src="upload/attach/202107/2_4XZTHM2QCBFFDSU.png"><span kdclassjsq="notranslate">魔方碎片</span>
Cubic Chaos Blade==<img src="upload/attach/202107/2_P78DFFU7Y8DD22S.png"><span kdclassjsq="notranslate">进阶魔方碎片</span>
Forever Unrelenting Black Flame==永恒的不灭黑色火焰
Forever Unrelenting Flame==<img src="upload/attach/202107/2_PJXKFPDPCEDKKEW.png">永恒的不灭火焰
Unrelenting Flame==<img src="upload/attach/202107/2_Q9NENZZA3UXZBMM.png">不灭之火焰
SS-rank==SS级
Pink Bean's Buddy==品克缤的伙伴
	

# 职业
	Demon Avenger==恶魔复仇者(白毛)
	Demon Slayer==红毛
	Beast Tamer==林之灵
	Angelic Buster==爆莉萌天使
	Dark Knight==黑骑
	dark Knight==黑骑士
	Wind Archer==风灵使者
	Thunder Breaker==奇袭者
	striker==奇袭者
	Dual Blade==双刀
	Dual blade==双刀
	Blaze Wizard==炎术士
	Night Walker==夜行者
	Night walker==夜行者
	Path finder==开拓者
	Pathfinder==开拓者
	Night Lord==标飞
	Night road==标飞	
	Wild Hunter==豹弩游侠
	Wild hunter==豹弩游侠
	Arch Mage==魔导师
	Bow Master==神射手(弓)
	Wind Archer==风灵使者
	Wind shooter==风灵使者
	Flame Wizard==炎术士
	Cannon Shooter==火炮
	Cannon shooter==火炮
	\(I/L\)==(冰雷)
	\(Ice, Lightning\)==(冰雷)
	Ice/Lightning==冰雷
	\(Fire Poison\)==(火毒)
	\(Fire, Poison\)==(火毒)
	Battle Mage==唤灵斗师
	Battle mage==唤灵斗师
	Dawn Warrior==魂骑士
	Seoul Master==魂骑士
	Capricious Sage==云游仙人
	Adele==阿黛尔



\bExplorer\b==冒险家
Resistance==反抗者
Bowmaster==神射手(弓)
Warrior==战士
Hayato==剑豪
Hayat==剑豪
Blaster==爆破手
Magician==法师
Mage==法师
Evan==龙神
Kanna==阴阳师
Lara\b==<span kdclassjsq="notranslate">拉拉</span>
Illium==圣晶使徒
Irium==圣晶使徒
Luminous\b==<span kdclassjsq="notranslate">夜光法师</span>
Kinesis==超能力者
Thief==飞侠
Xenon==尖兵
Zenon==尖兵
Pirate==海盗
Knights?==骑士
Bowman==弓箭手
Bowmen==弓箭手
\bArk\b==影魂异人
Aran==战神
Alan==战神
Heroes==英雄们
Hero==英雄
Bishop==主教
Shadower==刀飞
\bshadow\b==刀飞
Corsair==船长
captain==船长
Cannoneer==神炮王
Phantom==幻影
Phanton==幻影
phantom==幻影
Mihile==米哈尔
Mikhail==米哈尔
Michael==米哈尔
Shade==隐月
Hidden moon==隐月
Paladin==圣骑士
Archer==射手
Kaiser==狂龙战士
Cadena==魔链影士
Marksman==箭神(弩)
Crossbow master==箭神(弩)
Buccaneer==冲锋队长
Buccanner==冲锋队长
Viper==冲锋队长
Mercedes==双弩
Mechanic\b==机械师
Hoyoung==虎影
HOYOUNG==虎影


# 倾向
	\bAmbition\b==领袖气质
	\bEmpathy\b==感性
	\bInsight\b==洞察力
	\bWillpower\b==意志力
	\bDiligence\b==手技
	\bCharm\b==魅力

# 冒险岛
	MapleStoryM==冒险岛手机版
	MapleStory M==冒险岛手机版
	Maple M\s==冒险岛手机版
	MapleStory==冒险岛
	MAPLESTORY==冒险岛
	\bMaple\b==枫叶
	(\d\d)th Anniversary==\1周年
	Anniversary==周年

# 属性
Req Level==要求等级
Level==等级
Max MP==MP
MaxMP==MP
Max HP==HP
MaxHP==HP
Weapon Attack==攻击力
Ignore Defense==无视防御
Enemy Defense Ignored==无视防御
Defense==防御力
Attack Speed==攻速
Set Effects==套装属性
(\d) Set Effect==\1件套属性
(Grants\s)?Weapon ATT/Magic ATT==攻击力
Weapon ATT/magic ATT==攻击力
Magic Attack==魔法攻击力
Attack Power==物理攻击力
Weapon ATT==物理攻击力
Magic ATT==魔法攻击力
Normal Monster Damage==<span kdclassjsq="notranslate">普通怪物伤害</span>
Critical Rate==暴击命中率
damage when attacking BOSSes==BOSS伤害
damage when attacking monsters==怪物伤害
Boss Damage==BOSS伤害
BOSS Damage==BOSS伤害
Untradeable==不可交易
Grants stats STR/DEX/INT/LUK==全属性
(Grants)? STR/DEX/INT/LUK==全属性
attack speed IA==攻速内在

STR==力量
DEX==敏捷
LUK==运气
INT==智力
ATT==攻击力
DEF==防御力


# 装备_单
	Top==上衣
	\bBottom\b==裤子
	medals?==勋章
	Badge==徽章
	Totem==图腾
	Katana==武士刀
	Fan\s==扇子
	Scepter==驯兽魔法棒
	Spear==枪
	Polearm==矛
	Ellaha==机甲手枪
	Katara==副刀
	Cane==手杖(幻影)
	Dagger==匕首
	Hammer==单手锤
	Desperado==死亡使者(白毛)
	Axe==单手斧
	Saber==单手剑
	Wand==短杖
	Staff==长杖
	Guards==拳套
	Claw==拳套
	Pistol==手枪
	Siege Gun==手炮
	Hat==帽子
	Suit==套服
	\scapes?==披风
	Frozen==寒冰
	Absolab==埃苏莱布斯
	AbsoLab==埃苏莱布斯
	Bow==弓
	Crossbow==弩
	Chain==锁链
	\bGun\b==短枪
	Knuckle==指节
	armor==防具
	\btitle\b==称号
	\bTitle\b==称号
	Androids?==机器人
	Kodachi==小太刀
	
# temple
Empress's Gift==女皇的礼物

# BOSS
	Root Abyss Set Box==鲁比塔斯装备套装礼盒
	Root Abyss==鲁比塔斯
	Hard Lotus==困难级斯乌
	Hard Damien==困难级戴安米
	Hard Lotus's==困难级斯乌的
	Hard Will's==困难级威尔的
	Hard Will==困难级威尔
	Crimson Queen==血腥女皇
	Akechi Mitsuhide==明智光秀
	Von Bon==半半
	Normal Pink Bean==普通品克缤
	Pink Bean==品克缤
	Von Leon==狮子王
	Normal Cygnus==普通希纳斯女皇
	Easy Cygnus==简单希纳斯女皇
	Cygnus==希纳斯女皇
	Easy Zakum==简单扎昆
	Normal Zakum==普通扎昆
	Easy Horntail==简单暗黑龙王
	Normal Horntail==普通暗黑龙王
	Chaos Horntail==进阶暗黑龙王
	Easy Papulatus==简单闹钟
	Normal Papulatus==普通闹钟
	Papulatus Mark==闹钟脸饰
	Papulatus==闹钟
	Crimson Balrog==蝙蝠魔
	Verus Hillaroid==真神希腊机器人
	Verus Hilla==真神希拉
	Necro Orchid==死灵奥尔卡
	Necro Lotus==死灵斯乌
	Princess No==浓姬
	#######################
	Ursus==乌鲁斯
	Lotus==斯乌
	Magnus==麦格纳斯
	Will's==威尔的
	[^>]Will\b==威尔
	gollux==贝勒德
	Zakum==扎昆
	Horntail==暗黑龙王
	Hilla==希拉
	Pierre==皮埃尔
	Vellum==贝伦
	Arkarium==阿卡伊勒
	OMNI-CLN==钻机
	Damienroid==戴安米机器人
	Damien==戴安米
	Living Chains?==生存链
	Balrog==蝙蝠魔
	Arkarium==阿卡伊勒
	Yakuza==黑帮老大
	Gigatoad==忍者巨蛙

	
	Hard Lucid's==困难级路西德的
	Lucid's==路西德的
	Ebony Lucid==黑色路西德
	Powder Lucid==粉色路西德
	Rose Lucid==玫瑰路西德
	Lucid==路西德







# 单
Ability==内在能力
Erdas?==艾尔达
Tomoyo==“知世”
Syaoran==“小狼”
Syaoran Li==“小狼”
EXP==经验
Recipes?==配方
Julieta==朱莉埃塔
Skuas==斯库亚斯
Ossyria==神秘岛
Abrup==阿布鲁
Rings?==戒指
Guild==公会
Magnificent==伟大的
Chaos==困难级
CHAOS==困难级
Costume==套服
Legion==联盟
Bwuh==混乱
Eep==沮丧
Luv==喜欢
Kissy==吻
Yeti==白雪人
Pepe==蓝企鹅
\bMount\b==坐骑
World Gauge==世界储存槽
Determination Gauge==圣火储存槽
Spark of Determination==圣火火花
Determination==圣火
Desire==圣火
Guardian==守卫者
# 单.
coins?==币
Entry Coupons?==入场券
coupons?==券
channels?==频道
characters?==角色
maplers==冒险家们
\srings==戒指
equipment==装备
EVENTS?==活动
players?==玩家
reflection==反伤
Potions==药水
potions==药水
dorp==掉落
hardcore==极端
jobs?==职业
scrolls?==卷轴
boss==BOSS
sweetwater==漩涡
non-Transposed==非继承
transposed==继承
zero==神之子
Shoulder==护肩
Essence==精华
Fortuitous==幸运的
Swift==敏捷的
Clever==聪明的
Beefy==强壮的
Flashy==华丽的
Hearty==旺盛的
Potent\s==强效的
Radiant==闪耀的
Soul==灵魂
Items==物品
Afinas?==阿维纳斯
Bebe==贝贝
Chair==椅子
Player==玩家
Credit==信用值
Monad==莫奈得
Drop\b==爆率
drops\b==爆率
Megaphone==喇叭
Beret==蓓蕾帽
Epic (Rank)?== -A级(紫色)
potentials?==潜能
safeguard==强化防爆
gachapon==玩具
Tradeable==可交易
untradeable when equipped==装备后无法交易
untradeable==无法交易
Untradeble==无法交易
Untradable==无法交易
Untradeabe==无法交易
Cannot be traded==无法交易
Rewards?==奖励
Whiskers==胡子
Pollo==保罗
Fritto==普利托
hunting==猎杀
hunter==猎杀
hunt==猎杀 
permanent==永久
Orchid==奥尔卡
Togetherness==同伴

\bSpinel\b==绿水晶
\bequips?\b==装备套
\bbotters\b==开挂者
\bBotters\b==开挂者
\bevents?\b==活动
\bEvents?\b==活动
\bportals\b==光口
\bRune\b==符文
\brunes\b==符文
\bdungeon\b==副本
\bCape\b==披风
\bBox(es)?\b==盒子
\bRanks?\b==排名
\bSkills\b==技能
\bcharacters\b==玩家
\bdrop\b==掉落
\bEnhancement\b==增强
\bExpedition\b==远征
\bLil\b==迷你
\bLabyrinth\b==迷宫
\bNova\b==诺巴
\bNOVA\b==诺巴
\bTiara\b==头饰
\bclasss?\b==职业
\bclasses\b==职业
\bbans?\b==封号
\bBans?\b==封号
\bbanned\b==封号
\bHP\b==血量
binding skill==控制型技能
\bbind\b==控制
\bIED\b==无视防御
\bKeen\b==智慧的
\bReverse\b==翻转棋
\bAfterimage\b==残影
\binvincible\b==无敌状态
\b(m|M)esos?\b==金币
\b(\d)x\b==\1倍
\bFairgoer\b==街霸
\bGrants\b==获得
\bFlame\b==火花
\bandroid\b==机器人
Master\b==大师
\(Hard\)==(困难级)
\(Normal, Hard\)==(普通级，困难级)
\bAwakening\b==觉醒
\bAwake\b==觉醒
\bJusticar\b==司令

dungeons?==副本
Novice==新手
buffs==BUFF
cubes?==魔方

\bEarrings\b==耳环
\bGloves\b==手套
\bShoes\b==鞋子
\bCap\b==帽子
\bboots\b==鞋子
\bcloak\b==披风
\bbelt\b==腰带
\bBlunts?\b==钝器
\bTyrant\b==暴君
Emblem==纹章
\bOutfit\b==衣服
\bWeapon\b==武器

## 转录
kdclassjsq==class