package models

//请求定义描述
/*
	消息传递类型分三种
	1.Request：为客户端向服务器发送请求
	2.Response：客户端所有的请求，服务器端都会给返回一个Response
	3.Notice：服务器端对客户端主动发送的消息
*/

//用于标记当前客户端在什么类型场景中
const (
	_             = iota
	SceneInHall   //在大厅
	SceneInQueue  //在队列中
	SceneInBattle //在战斗中
	SceneInSettle //结算中
)

//系统级操作
const (
	_                            = 1000000 + iota
	ActiveTypeAddtoQueueRequest  //开始排队
	ActiveTypeAddtoQueueResponse //排队反馈
	ActiveTypePreToBattleNotice  //准备进图战场(排队成功)
	ActiveTypeChangeScene        //切换场景通知
)

//战场外操作
const (
	_                                 = 2000000 + iota
	ActiveTypeCharacterListRequest    //请求人物列表
	ActiveTypeCharacterListResponse   //人物列表
	ActiveTypeSelectCharacterRequest  //选择人物
	ActiveTypeSelectCharacterResponse //选择人物结果
	ActiveTypeBattleCountDownNotice   //战斗开始倒计时
	ActiveTypeBattleStartNotice       //战斗开始
)

//战场内操作
const (
	_                              = 3000000 + iota
	ActiveTypeBattleIntervalNotice //战场定时广播
	ActiveTypeBattleWinNotice      //战斗胜利通知
	ActiveTypeBattleLoseNotice     //战斗失败通知
	ActiveTypeBattleEndNotice      //战斗结束通知
	ActiveTypeMoveRequest          //移动
	ActiveTypeMoveResponse         //移动结果
	ActiveTypeBuyItemRequest       //购买道具
	ActiveTypeBuyItemResponse      //购买道具
	ActiveTypeAttachRequest        //攻击
	ActiveTypeAttachResponse       //攻击
	ActiveTypeAttachedNotice       //攻击通知
)
