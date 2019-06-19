package models

//请求定义描述
/*
	消息传递类型分三种
	1.Request：为客户端向服务器发送请求
	2.Response：客户端所有的请求，服务器端都会给返回一个Response
	3.Notice：服务器端对客户端主动发送的消息
*/

//常规接入动作
const (
	_                            = 1000000 + iota
	ActiveTypeAddtoQueueRequest  //开始排队
	ActiveTypeAddtoQueueResponse //排队反馈
	ActiveTypePreToBattleNotice  //准备进图战场(排队成功)
)

//角色操作动作
const (
	_                                 = 2000000 + iota
	ActiveTypeCharacterListRequest    //请求人物列表
	ActiveTypeCharacterListResponse   //人物列表
	ActiveTypeSelectCharacterRequest  //选择人物
	ActiveTypeSelectCharacterResponse //选择人物结果
	ActiveTypeBattleCountDownNotice   //战斗开始倒计时
	ActiveTypeBattleStartNotice       //战斗开始
	ActiveTypeBattleIntervalNotice    //战场定时广播
)

//战斗动作
const (
	_                         = 3000000 + iota
	ActiveTypeMoveRequest     //移动
	ActiveTypeMoveResponse    //移动结果
	ActiveTypeBuyItemRequest  //购买道具
	ActiveTypeBuyItemResponse //购买道具
	ActiveTypeAttachRequest   //攻击
	ActiveTypeAttachResponse  //攻击
	ActiveTypeAttachedNotice  //攻击通知
)
