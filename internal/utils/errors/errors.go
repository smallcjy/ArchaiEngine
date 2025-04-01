package errors

// gameServer错误码 2000
const (
	CodeSuccess = 200 //用这个表示成功

	//###	请求指令
	ErrCodeRequest    = 2000 // 请求错误
	ErrCodeNoCmd      = 2001 // 指令错误
	ErrCodeParam      = 2002 // 参数错误
	ErrCodeDataFormat = 2003 // 数据格式错误
	ErrCodePermission = 2006 // 权限不足
	ErrCodeUserAuth   = 2007 // 用户校验失败
	ErrCodeServer     = 2602 // 接口异常

	//###	房间/频道
	ErrCodeRoom          = 3000 //房间数据错误
	ErrCodeRoomExists    = 3001 //房间已存在
	ErrCodeRoomNonExists = 3002 //房间不存在
	ErrCodeNotJoinedRoom = 3003 //玩家不在房间内
	ErrCodeJoinedRoom    = 3004 //玩家已在房间内
	ErrCodeRoomMaxUser   = 3005 //房间人数已满
	ErrCodeRoomMinUser   = 3006 //房间人数不足

	//###	游戏
	ErrCodeGame          = 4000 // 游戏数据错误
	ErrCodeConf          = 4099 // 游戏配置错误
	ErrCodeDb            = 4100 // 游戏数据错误
	ErrCodeRedis		 = 4098 // 游戏数据错误
	ErrCodeNotJoinedGame = 4101 // 玩家未加入游戏
	ErrCodeGameNotStart  = 4102 // 游戏非进行中
	ErrCodeNotYouAction  = 4103 // 非你行动回合
	ErrCodeChooseStep    = 4104 // 现在是选牌库阶段
	ErrCodePokersNotDone = 4105 // 牌组未完成组合
	ErrCodeAlreadyInRoom = 4106 // 已在房间中
	ErrCodeConisLimit    = 4107 // 金币不足
	ErrCodeNotMateState  = 4108 // 非匹配状态
	ErrCodeNoAi          = 4109 // 非AI托管状态
	ErrCodeNotChooseStep = 4110 // 现在非选库阶段
	ErrCodeNotPushStep   = 4111 // 现在非出牌阶段
	ErrCodeIpErr         = 4112 // ip异常
	ErrCodeJoinedGame    = 4102 //已加入游戏
	ErrCodeJoinGameFail  = 4103 //加入游戏失败
	ErrCodeGameMaxPlayer = 4104 //游戏人数已满
	ErrCodeGameMinPlayer = 4105 //游戏人数不足
	ErrCodeNotAllReady   = 4106 //未全部准备
	ErrCodeSeatNo        = 4107 //座位被占用

	ErrCodeGameStatus       = 4200 //游戏状态错误
	ErrCodeGameStartFail    = 4201 //游戏开始失败（条件不足）
	ErrCodeGameStopFail     = 4202 //游戏结束失败（状态异常）
	ErrCodeGameWaiting      = 4203 //游戏未开始（状态）
	ErrCodeGameAlreadyStart = 4204 //游戏已开始（状态）
	ErrCodeGameIsOver       = 4205 //游戏已结束（状态）
	ErrCodePlayerRound      = 4206 //不是你的回合

	ErrCodePlayerDie       = 10000 //玩家已淘汰
	ErrCodeGetUserInfoFail = 10001 //获取用户信息失败
	ErrCodeKickSelf        = 10002 //不能踢出自己
	ErrCodeBeKick          = 10003 //您已被房主踢出房间
	ErrCodeUseItem         = 10004 //未开启道具模式
	ErrCodeNotFoundTarget  = 10005 //未找到目标
	ErrCodeRankQuery       = 10006 //获取排行榜失败
	ErrCodeSettingChange   = 10007 //管理员修改了门票金额，请重新加入
)

var (
	Success = newError(CodeSuccess, "")

	//###	请求指令
	ErrRequest    = newError(ErrCodeRequest, "请求错误")
	ErrUserAuth   = newError(ErrCodeUserAuth, "校验用户失败")
	ErrNoCmd      = newError(ErrCodeNoCmd, "指令错误")
	ErrParam      = newError(ErrCodeParam, "参数错误")
	ErrDataFormat = newError(ErrCodeDataFormat, "数据格式错误")
	ErrPermission = newError(ErrCodePermission, "权限不足")

	//###	房间/频道
	ErrRoom          = newError(ErrCodeRoom, "房间数据错误")
	ErrRoomExists    = newError(ErrCodeRoomExists, "房间已存在")
	ErrRoomNonExists = newError(ErrCodeRoomNonExists, "房间不存在")
	ErrNotJoinedRoom = newError(ErrCodeNotJoinedRoom, "玩家不在房间内")
	ErrJoinedRoom    = newError(ErrCodeJoinedRoom, "玩家已在房间内")
	ErrRoomMaxUser   = newError(ErrCodeRoomMaxUser, "房间人数已满")
	ErrRoomMinUser   = newError(ErrCodeRoomMinUser, "房间人数不足")

	//###	游戏
	ErrConf             = newError(ErrCodeConf, "游戏配置错误")
	ErrDb               = newError(ErrCodeDb, "数据库异常")
	ErrRedis            = newError(ErrCodeRedis, "Redis异常")
	ErrGame             = newError(ErrCodeGame, "游戏数据错误")
	ErrNotJoinedGame    = newError(ErrCodeNotJoinedGame, "玩家未加入游戏")
	ErrGameNotStart     = newError(ErrCodeGameNotStart, "游戏非进行中")
	ErrNotYouAction     = newError(ErrCodeNotYouAction, "非你行动回合")
	ErrChooseStep       = newError(ErrCodeChooseStep, "现在是选牌库阶段")
	ErrPokersNotDone    = newError(ErrCodePokersNotDone, "牌组未完成组合")
	ErrAlreadyInRoom    = newError(ErrCodeAlreadyInRoom, "已在房间中")
	ErrConisLimit       = newError(ErrCodeConisLimit, "金币不足")
	ErrNotMateState     = newError(ErrCodeNotMateState, "非匹配状态")
	ErrNoAi             = newError(ErrCodeNoAi, "非AI托管状态")
	ErrNotChooseStep    = newError(ErrCodeNotChooseStep, "现在非选库阶段")
	ErrNotPushStep      = newError(ErrCodeNotPushStep, "现在非出牌阶段")
	ErrIpErr            = newError(ErrCodeIpErr, "ip错误")
	ErrJoinedGame       = newError(ErrCodeJoinedGame, "已加入游戏")
	ErrJoinGameFail     = newError(ErrCodeJoinGameFail, "加入游戏失败")
	ErrGameMaxPlayer    = newError(ErrCodeGameMaxPlayer, "游戏人数已满")
	ErrGameMinPlayer    = newError(ErrCodeGameMinPlayer, "游戏人数不足")
	ErrNotAllReady      = newError(ErrCodeNotAllReady, "未全部准备")
	ErrGameStatus       = newError(ErrCodeGameStatus, "游戏状态错误")
	ErrGameStartFail    = newError(ErrCodeGameStartFail, "游戏开始失败")
	ErrGameStopFail     = newError(ErrCodeGameStopFail, "游戏结束失败")
	ErrGameWaiting      = newError(ErrCodeGameWaiting, "游戏未开始")
	ErrGameAlreadyStart = newError(ErrCodeGameAlreadyStart, "游戏已开始")
	ErrGameIsOver       = newError(ErrCodeGameIsOver, "游戏已结束")
	ErrPlayerRound      = newError(ErrCodePlayerRound, "不是你的回合")
	ErrSeatNo           = newError(ErrCodeSeatNo, "座位已被占用")
	ErrServer           = newError(ErrCodeServer, "接口异常")
	ErrPlayerDie        = newError(ErrCodePlayerDie, "玩家已淘汰")
	ErrGetUserInfoFail  = newError(ErrCodeGetUserInfoFail, "获取用户信息失败")
	ErrKickSelf         = newError(ErrCodeKickSelf, "不能踢出自己")
	ErrBeKick           = newError(ErrCodeBeKick, "您已被房主踢出房间")
	ErrUseItem          = newError(ErrCodeUseItem, "未开启道具模式")
	ErrNotFoundTarget   = newError(ErrCodeNotFoundTarget, "未找到目标")
	ErrRankQuery        = newError(ErrCodeRankQuery, "获取排行榜失败")
	ErrSettingChange    = newError(ErrCodeSettingChange, "管理员修改了门票金额，请重新加入")
)

func newError(code int32, detail string) *Error {
	return &Error{
		Code: code,
		Msg:  detail,
	}
}

type Error struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Error) Error() string {
	//b, _ := json.Marshal(e)
	return e.Msg
}
