package biz

import (
	"context"
	v1 "dhb/app/app/api"
	"encoding/base64"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID                     int64
	Address                string
	LastBiw                uint64
	Password               string
	Undo                   int64
	AddressTwo             string
	PrivateKey             string
	AddressThree           string
	WordThree              string
	PrivateKeyThree        string
	Last                   uint64
	Amount                 uint64
	AmountBiw              uint64
	OutRate                int64
	Total                  uint64
	IsDelete               int64
	RecommendLevel         int64
	Out                    int64
	Vip                    int64
	VipAdmin               int64
	LockReward             int64
	CreatedAt              time.Time
	UpdatedAt              time.Time
	Lock                   int64
	AmountUsdt             float64
	MyTotalAmount          float64
	AmountUsdtGet          float64
	AmountRecommendUsdtGet float64
	RecommendUserReward    int64
	RecommendUser          int64
	AmountFour             float64
	AmountFourGet          float64
	RecommendUserH         int64
	One                    string
	Two                    string
	Three                  string
	Four                   string
	Five                   string
	Six                    string
	Seven                  string
	AmountSelf             uint64
}

type Total struct {
	ID    int64
	One   float64
	Two   float64
	Three float64
}

type Good struct {
	ID     int64
	Amount uint64
	Name   string
	One    string
	Two    string
	Three  string
}

type UserInfo struct {
	ID               int64
	UserId           int64
	Vip              int64
	UseVip           int64
	HistoryRecommend int64
	TeamCsdBalance   int64
}

type UserRecommend struct {
	ID            int64
	UserId        int64
	RecommendCode string
	Total         int64
	CreatedAt     time.Time
}

type UserRecommendArea struct {
	ID            int64
	RecommendCode string
	Num           int64
	CreatedAt     time.Time
}

type Trade struct {
	ID           int64
	UserId       int64
	AmountCsd    int64
	RelAmountCsd int64
	AmountHbs    int64
	RelAmountHbs int64
	Status       string
	CreatedAt    time.Time
}

type UserArea struct {
	ID         int64
	UserId     int64
	Amount     int64
	SelfAmount int64
	Level      int64
}

type UserCurrentMonthRecommend struct {
	ID              int64
	UserId          int64
	RecommendUserId int64
	Date            time.Time
}

type Config struct {
	ID      int64
	KeyName string
	Name    string
	Value   string
}

type UserBalance struct {
	ID                     int64
	UserId                 int64
	BalanceUsdt            int64
	BalanceUsdt2           int64
	BalanceDhb             int64
	RecommendTotal         int64
	AreaTotal              int64
	FourTotal              int64
	LocationTotal          int64
	BalanceC               int64
	RecommendTotalFloat    float64
	RecommendTotalFloatTwo float64
	AreaTotalFloat         float64
	AllFloat               float64
	AreaTotalFloatTwo      float64
	AreaTotalFloatThree    float64
	LocationTotalFloat     float64
	BalanceUsdtFloat       float64
	BalanceRawFloat        float64
	BalanceRawFloatNew     float64
	BalanceKsdtFloat       float64
}

type Stake struct {
	ID        int64
	UserId    int64
	Status    int64
	Day       int64
	Amount    float64
	Reward    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Withdraw struct {
	ID              int64
	UserId          int64
	Amount          int64
	AmountNew       float64
	RelAmountNew    float64
	RelAmount       int64
	BalanceRecordId int64
	Status          string
	Type            string
	Address         string
	CreatedAt       time.Time
}

type UserSortRecommendReward struct {
	UserId int64
	Total  int64
}

type UserUseCase struct {
	repo                          UserRepo
	urRepo                        UserRecommendRepo
	configRepo                    ConfigRepo
	uiRepo                        UserInfoRepo
	ubRepo                        UserBalanceRepo
	locationRepo                  LocationRepo
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo
	tx                            Transaction
	log                           *log.Helper
}

type LocationNew struct {
	ID                int64
	UserId            int64
	Num               int64
	Status            string
	Current           int64
	CurrentMax        int64
	CurrentMaxNew     int64
	StopLocationAgain int64
	OutRate           int64
	Count             int64
	StopCoin          int64
	Top               int64
	Usdt              int64
	Total             int64
	TotalTwo          int64
	TotalThree        int64
	Biw               int64
	TopNum            int64
	LastLevel         int64
	StopDate          time.Time
	CreatedAt         time.Time
}

type UserBalanceRecord struct {
	ID           int64
	UserId       int64
	Amount       int64
	CoinType     string
	Balance      int64
	Type         string
	AmountNew    float64
	AmountNewTwo float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type BalanceReward struct {
	ID        int64
	UserId    int64
	Status    int64
	Amount    int64
	SetDate   time.Time
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Reward struct {
	ID               int64
	UserId           int64
	Amount           int64
	AmountB          int64
	BalanceRecordId  int64
	Type             string
	TypeRecordId     int64
	Reason           string
	ReasonLocationId int64
	LocationType     string
	Address          string
	AmountNew        float64
	AmountNewTwo     float64
	CreatedAt        time.Time
}

type BuyRecord struct {
	ID          int64
	UserId      int64
	Status      int64
	Amount      float64
	AmountGet   float64
	LastUpdated int64
	One         string
	Two         string
	Three       string
	Four        int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Pagination struct {
	PageNum  int
	PageSize int
}

type ConfigRepo interface {
	GetConfigByKeys(ctx context.Context, keys ...string) ([]*Config, error)
	GetConfigs(ctx context.Context) ([]*Config, error)
	UpdateConfig(ctx context.Context, id int64, value string) (bool, error)
}

type UserBalanceRepo interface {
	GetTotal(ctx context.Context) (*Total, error)
	RecommendLocationRewardBiw(ctx context.Context, userId int64, rewardAmount int64, recommendNum int64, stop string, tmpMaxNew int64, feeRate int64) (int64, error)
	CreateUserBalance(ctx context.Context, u *User) (*UserBalance, error)
	CreateUserBalanceLock(ctx context.Context, u *User) (*UserBalance, error)
	LocationReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error)
	WithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error)
	RecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	SystemWithdrawReward(ctx context.Context, amount int64, locationId int64) error
	SystemReward(ctx context.Context, amount int64, locationId int64) error
	SystemFee(ctx context.Context, amount int64, locationId int64) error
	GetSystemYesterdayDailyReward(ctx context.Context) (*Reward, error)
	UserFee(ctx context.Context, userId int64, amount int64) (int64, error)
	RecommendWithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	NormalRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	NormalWithdrawRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	Deposit(ctx context.Context, userId int64, amount int64) (int64, error)
	DepositLast(ctx context.Context, userId int64, lastAmount int64, locationId int64) (int64, error)
	DepositDhb(ctx context.Context, userId int64, amount int64) (int64, error)
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
	GetRewardFourYes(ctx context.Context) (*Reward, error)
	GetUserRewardByUserId(ctx context.Context, userId int64) ([]*Reward, error)
	GetUserRewardByUserIdPage(ctx context.Context, b *Pagination, userId int64, reason string) ([]*Reward, error, int64)
	GetUserRewardDeposit(ctx context.Context, userId int64) ([]*Reward, error)
	GetLocationsToday(ctx context.Context) ([]*LocationNew, error)
	GetUserRewardByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserSortRecommendReward, error)
	GetUserRewards(ctx context.Context, b *Pagination, userId int64) ([]*Reward, error, int64)
	GetUserRewardsLastMonthFee(ctx context.Context) ([]*Reward, error)
	GetUserBalanceByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserBalance, error)
	GetUserBalanceUsdtTotal(ctx context.Context) (int64, error)
	GreateWithdraw(ctx context.Context, userId int64, relAmount float64, amount float64, coinType string, address string) (*Withdraw, error)
	WithdrawUsdt(ctx context.Context, userId int64, amount int64, tmpRecommendUserIdsInt []int64) error
	WithdrawUsdt2(ctx context.Context, userId int64, amount float64) error
	WithdrawISPAY(ctx context.Context, userId int64, amount float64) error
	ToAddressAmountUsdt(ctx context.Context, userId int64, toUserId int64, amount float64, address string) error
	ToAddressAmountKsdt(ctx context.Context, userId int64, toUserId int64, amount float64, address string) error
	ToAddressAmountRaw(ctx context.Context, userId int64, toUserId int64, amount float64, address string) error
	StakeAmount(ctx context.Context, userId int64, amount float64, day int64) error
	SetToday(ctx context.Context, userId int64, amount float64, day int64) error
	UnStakeAmount(ctx context.Context, userId int64, stakeId int64, amount float64) error
	Exchange(ctx context.Context, userId int64, amountUsdt, fee, amountRawSub float64) error
	GetSystemYesterdayLocationReward(ctx context.Context) ([]*UserBalanceRecord, error)
	GetRewardBuyYes(ctx context.Context) ([]*Reward, error)
	Exchange2(ctx context.Context, userId int64, amount int64, amountUsdtSubFee int64, amountUsdt int64, locationId int64) error
	WithdrawUsdt3(ctx context.Context, userId int64, amount int64) error
	TranUsdt(ctx context.Context, userId int64, toUserId int64, amount int64, tmpRecommendUserIdsInt []int64, tmpRecommendUserIdsInt2 []int64) error
	WithdrawDhb(ctx context.Context, userId int64, amount int64) error
	WithdrawC(ctx context.Context, userId int64, amount int64) error
	TranDhb(ctx context.Context, userId int64, toUserId int64, amount int64) error
	GetWithdrawByUserId(ctx context.Context, userId int64, coinType string, b *Pagination) ([]*Withdraw, int64, error)
	GetWithdrawByUserId2(ctx context.Context, userId int64) ([]*Withdraw, error)
	GetUserBalanceRecordByUserId(ctx context.Context, userId int64, typeCoin string, tran string) ([]*UserBalanceRecord, error)
	GetUserBalanceRecordsByUserId(ctx context.Context, userId int64) ([]*UserBalanceRecord, error)
	GetTradeByUserId(ctx context.Context, userId int64) ([]*Trade, error)
	GetWithdraws(ctx context.Context, b *Pagination, userId int64) ([]*Withdraw, error, int64)
	GetWithdrawPassOrRewarded(ctx context.Context) ([]*Withdraw, error)
	UpdateWithdraw(ctx context.Context, id int64, status string) (*Withdraw, error)
	GetWithdrawById(ctx context.Context, id int64) (*Withdraw, error)
	GetWithdrawNotDeal(ctx context.Context) ([]*Withdraw, error)
	GetUserBalanceRecordUserUsdtTotal(ctx context.Context, userId int64) (int64, error)
	GetUserBalanceRecordUsdtTotal(ctx context.Context) (int64, error)
	GetUserBalanceRecordUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotal(ctx context.Context) (int64, error)
	GetUserRewardUsdtTotal(ctx context.Context) (int64, error)
	GetSystemRewardUsdtTotal(ctx context.Context) (int64, error)
	UpdateWithdrawAmount(ctx context.Context, id int64, status string, amount int64) (*Withdraw, error)
	GetUserRewardRecommendSort(ctx context.Context) ([]*UserSortRecommendReward, error)
	GetUserRewardTodayTotalByUserId(ctx context.Context, userId int64) (*UserSortRecommendReward, error)
	GetGoods(ctx context.Context) ([]*Good, error)
	GetGoodsTwo(ctx context.Context) ([]*Good, error)
	GetGoodsOnline(ctx context.Context) ([]*Good, error)
	GetGoodsOnlineTwo(ctx context.Context) ([]*Good, error)
	GetGoodsOnlineThree(ctx context.Context) ([]*Good, error)
	GetGoodsThree(ctx context.Context) ([]*Good, error)
	GetAllUsersB(ctx context.Context) ([]*User, error)
	GetGoodsAll(ctx context.Context) ([]*Good, error)

	SetBalanceReward(ctx context.Context, userId int64, amount int64) error
	UpdateBalanceReward(ctx context.Context, userId int64, id int64, amount int64, status int64) error
	GetBalanceRewardByUserId(ctx context.Context, userId int64) ([]*BalanceReward, error)

	GetUserBalanceLock(ctx context.Context, userId int64) (*UserBalance, error)
	Trade(ctx context.Context, userId int64, amount int64, amountB int64, amountRel int64, amountBRel int64, tmpRecommendUserIdsInt []int64, amount2 int64) error
	GetStakeByUserId(ctx context.Context, userId int64) ([]*Stake, error)
}

type UserRecommendRepo interface {
	GetUserRecommends(ctx context.Context) ([]*UserRecommend, error)
	GetUserRecommendByUserId(ctx context.Context, userId int64) (*UserRecommend, error)
	GetUserRecommendsFour(ctx context.Context) ([]*UserRecommend, error)
	CreateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (*UserRecommend, error)
	UpdateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (bool, error)
	GetUserRecommendByCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCodeSum(ctx context.Context, code string) (int64, error)
	CreateUserRecommendArea(ctx context.Context, u *User, recommendUser *UserRecommend) (bool, error)
	DeleteOrOriginUserRecommendArea(ctx context.Context, code string, originCode string) (bool, error)
	GetUserRecommendLowArea(ctx context.Context, code string) ([]*UserRecommendArea, error)
	GetUserAreas(ctx context.Context, userIds []int64) ([]*UserArea, error)
	CreateUserArea(ctx context.Context, u *User) (bool, error)
	GetUserArea(ctx context.Context, userId int64) (*UserArea, error)
	UpdateUserRecommendTotal(ctx context.Context, userId int64, total int64) error
}

type UserCurrentMonthRecommendRepo interface {
	GetUserCurrentMonthRecommendByUserId(ctx context.Context, userId int64) ([]*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendGroupByUserId(ctx context.Context, b *Pagination, userId int64) ([]*UserCurrentMonthRecommend, error, int64)
	CreateUserCurrentMonthRecommend(ctx context.Context, u *UserCurrentMonthRecommend) (*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendCountByUserIds(ctx context.Context, userIds ...int64) (map[int64]int64, error)
	GetUserLastMonthRecommend(ctx context.Context) ([]int64, error)
}

type UserInfoRepo interface {
	GetBuyRecord(ctx context.Context, day int) ([]*BuyRecord, error)
	UpdateUserMyTotalAmountAdd(ctx context.Context, userId int64, amountUsdt, myTotal float64) error
	UpdateUserRewardRecommend2(ctx context.Context, id, userId int64, usdt, raw, usdtOrigin float64, amountOrigin float64, stop bool, address string) error
	UpdateUserRewardRecommendBrc(ctx context.Context, userId int64, raw float64, address string) error
	CreateEthUserRecordListByHash(ctx context.Context, r *EthUserRecord) (*EthUserRecord, error)
	UpdateUserNewTwoNewTwo(ctx context.Context, userId int64, amount uint64, amountRel, amountRelBrc, amountIspay float64, one, two, three string, four int64) error
	UpdateUserTwoIn(ctx context.Context, userId int64, amount uint64, amountRel, amountRelBrc float64, one, two, three string, four int64) error
	UpdateUserThreeIn(ctx context.Context, userId int64, amount uint64, amountRel, amountRelBrc float64, one, two, three string, four int64) error
	CreateUserInfo(ctx context.Context, u *User) (*UserInfo, error)
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
	UpdateUserInfo(ctx context.Context, u *UserInfo) (*UserInfo, error)
	GetUserInfoByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserInfo, error)
}

type UserRepo interface {
	GetAllUsersOrderAmountBiw(ctx context.Context) ([]*User, error)
	GetAllUsersRecommendOrder(ctx context.Context) ([]*User, error)
	GetEthUserRecordListByUserId(ctx context.Context, userId int64) ([]*EthUserRecord, error)
	InRecordNew(ctx context.Context, userId int64, address string, amount int64, coinType string) error
	UpdateUserNewTwoNew(ctx context.Context, userId int64, amount uint64, amountUsdt float64, last uint64) error
	UpdateUserMyTotalAmount(ctx context.Context, userId int64, amountUsdt float64) error
	UpdateUserMyRecommendTotalNum(ctx context.Context, userId int64, address string, rewardHb int64, tmpRewardU bool) error
	UpdateUserMyRecommendTotal(ctx context.Context, userId int64, amountUsdt float64) error
	UpdateUserAddressInfo(ctx context.Context, userId int64, one, two, three, four, five, six, seven string) error
	UpdateUserVip(ctx context.Context, userId int64, vip int64) error
	UpdateUserRewardRecommend2(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, level, i int64, address string) (int64, error)
	UpdateUserRewardRecommendNew(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, i int64, address string) (int64, error)
	UpdateUserMyTotalAmountSub(ctx context.Context, userId int64, amountUsdt float64) error
	UpdateUserRewardAreaTwo(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, level, i int64, address string) (int64, error)
	GetUserById(ctx context.Context, Id int64) (*User, error)
	GetBuyRecord(ctx context.Context, userId uint64, b *Pagination) ([]*BuyRecord, int64, error)
	GetBuyRecordDoing(userId uint64) ([]*BuyRecord, error)
	GetUserByAddresses(ctx context.Context, Addresses ...string) (map[string]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error)
	GetUserByUserIdsTwo(ctx context.Context, userIds []int64) (map[int64]*User, error)
	GetUsers(ctx context.Context, b *Pagination, address string) ([]*User, error, int64)
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserCount(ctx context.Context) (int64, error)
	GetUserCountToday(ctx context.Context) (int64, error)
}

func NewUserUseCase(repo UserRepo, tx Transaction, configRepo ConfigRepo, uiRepo UserInfoRepo, urRepo UserRecommendRepo, locationRepo LocationRepo, userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:                          repo,
		tx:                            tx,
		configRepo:                    configRepo,
		locationRepo:                  locationRepo,
		userCurrentMonthRecommendRepo: userCurrentMonthRecommendRepo,
		uiRepo:                        uiRepo,
		urRepo:                        urRepo,
		ubRepo:                        ubRepo,
		log:                           log.NewHelper(logger),
	}
}

func (uuc *UserUseCase) GetUserByAddress(ctx context.Context, Addresses ...string) (map[string]*User, error) {
	return uuc.repo.GetUserByAddresses(ctx, Addresses...)
}

func (uuc *UserUseCase) GetDhbConfig(ctx context.Context) ([]*Config, error) {
	return uuc.configRepo.GetConfigByKeys(ctx, "level1Dhb", "level2Dhb", "level3Dhb")
}

func (uuc *UserUseCase) GetExistUserByAddressOrCreate(ctx context.Context, u *User, req *v1.EthAuthorizeRequest) (*User, error, string) {
	var (
		user          *User
		rUser         *User
		recommendUser *UserRecommend
		err           error
		//decodeBytes   []byte
	)

	recommendUser = &UserRecommend{
		ID:            0,
		UserId:        0,
		RecommendCode: "",
		Total:         0,
		CreatedAt:     time.Time{},
	}

	user, err = uuc.repo.GetUserByAddress(ctx, u.Address) // 查询用户
	if nil == user || nil != err {
		code := req.SendBody.Code // 查询推荐码 abf00dd52c08a9213f225827bc3fb100 md5 dhbmachinefirst
		if "abf00dd52c08a9213f225827bc3fb100" != code {
			if 1 >= len(code) {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码1"), "无效的推荐码"
			}
			var (
				userRecommend *User
			)

			userRecommend, err = uuc.repo.GetUserByAddress(ctx, code)
			if nil == userRecommend || err != nil {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码1"), "无效的推荐码"
			}

			//if 0 >= userRecommend.Amount {
			//	if 0 >= userRecommend.OutRate {
			//		return nil, errors.New(500, "USER_ERROR", "推荐人未激活"), "推荐人未激活"
			//	}
			//}

			// 查询推荐人的相关信息
			recommendUser, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userRecommend.ID)
			if nil == recommendUser || err != nil {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码3"), "无效的推荐码3"
			}

			if 0 < recommendUser.UserId {
				rUser, err = uuc.repo.GetUserById(ctx, recommendUser.UserId)
				if nil == rUser || err != nil {
					return nil, errors.New(500, "USER_ERROR", "无效的推荐码3"), "无效的推荐码3"
				}
			}
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			user, err = uuc.repo.CreateUser(ctx, u) // 用户创建
			if err != nil {
				return err
			}

			_, err = uuc.urRepo.CreateUserRecommend(ctx, user, recommendUser) // 创建用户推荐信息
			if err != nil {
				return err
			}

			_, err = uuc.ubRepo.CreateUserBalance(ctx, user) // 创建余额信息
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			return nil, err, "错误"
		}
	}

	return user, nil, ""
}

func generateWord() string {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 定义总组数和每组的字符数
	numGroups := 12
	groupSize := 3

	// 生成随机字符串
	var result []string
	for i := 0; i < numGroups; i++ {
		result = append(result, randString(groupSize))
	}

	// 将字符串数组用逗号连接
	finalString := strings.Join(result, ",")
	return finalString
}

func randString(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (uuc *UserUseCase) UpdateUserRecommend(ctx context.Context, u *User, req *v1.RecommendUpdateRequest) (*v1.RecommendUpdateReply, error) {
	var (
		err                   error
		userId                int64
		recommendUser         *UserRecommend
		userRecommend         *UserRecommend
		locations             []*LocationNew
		myRecommendUser       *User
		myUserRecommendUserId int64
		Address               string
		decodeBytes           []byte
	)

	code := req.SendBody.Code // 查询推荐码 abf00dd52c08a9213f225827bc3fb100 md5 dhbmachinefirst
	if "abf00dd52c08a9213f225827bc3fb100" != code {
		decodeBytes, err = base64.StdEncoding.DecodeString(code)
		code = string(decodeBytes)
		if 1 >= len(code) {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码4")
		}
		if userId, err = strconv.ParseInt(code[1:], 10, 64); 0 >= userId || nil != err {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码5")
		}

		// 现有推荐人信息，判断推荐人是否改变
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, u.ID)
		if nil == userRecommend {
			return nil, err
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
			if 2 <= len(tmpRecommendUserIds) {
				myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
			myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
			if nil != err {
				return nil, err
			}
		}

		if nil == myRecommendUser {
			return &v1.RecommendUpdateReply{InviteUserAddress: ""}, nil
		}

		if myRecommendUser.ID == userId {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		if u.ID == userId {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		// 我的占位信息
		locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, u.ID)
		if nil != err {
			return nil, err
		}
		if nil != locations && 0 < len(locations) {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		// 查询推荐人的相关信息
		recommendUser, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userId)
		if err != nil {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码6")
		}

		// 推荐人信息
		myRecommendUser, err = uuc.repo.GetUserById(ctx, userId)
		if err != nil {
			return nil, err
		}

		// 更新
		_, err = uuc.urRepo.UpdateUserRecommend(ctx, u, recommendUser)
		if err != nil {
			return nil, err
		}
		Address = myRecommendUser.Address
	}

	return &v1.RecommendUpdateReply{InviteUserAddress: Address}, err
}

func (uuc *UserUseCase) UserInfo(ctx context.Context, user *User) (*v1.UserInfoReply, error) {

	var (
		err                   error
		myUser                *User
		userRecommend         *UserRecommend
		myUserRecommendUserId int64
		inviteUserAddress     string
		configs               []*Config
		userBalance           *UserBalance
		withdrawMin           float64
		withdrawRate          float64
		withdrawMinTwo        float64
		withdrawRateTwo       float64
		notice                string
		users                 []*User
		goods                 []*Good
		goodsTwo              []*Good
		goodsThree            []*Good
	)

	// 配置
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"withdraw_rate",
		"withdraw_amount_min",
		"withdraw_rate_two",
		"withdraw_amount_min_two",
		"notice",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_amount_min" == vConfig.KeyName {
				withdrawMin, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "withdraw_rate" == vConfig.KeyName {
				withdrawRate, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "withdraw_amount_min_two" == vConfig.KeyName {
				withdrawMinTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "withdraw_rate_two" == vConfig.KeyName {
				withdrawRateTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "notice" == vConfig.KeyName {
				notice = vConfig.Value
			}
		}
	}

	users, err = uuc.repo.GetAllUsers(ctx)
	if nil != err {
		return nil, err
	}

	usersMap := make(map[int64]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	if _, ok := usersMap[user.ID]; !ok {
		return nil, err
	}
	myUser = usersMap[user.ID]

	if 1 == myUser.IsDelete {
		return nil, nil
	}

	if 1 == myUser.Lock {
		return nil, nil
	}

	// 余额，收益总数
	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUser.ID)
	if nil != err {
		return nil, err
	}

	goods, err = uuc.ubRepo.GetGoods(ctx)
	if nil != err {
		return nil, err
	}
	resGoods := make([]*v1.UserInfoReply_List, 0)
	for _, v := range goods {
		resGoods = append(resGoods, &v1.UserInfoReply_List{
			Id:    v.ID,
			One:   v.Name,
			Two:   v.One,
			Three: v.Two,
			Four:  v.Amount,
			Five:  v.Three,
		})
	}

	goodsTwo, err = uuc.ubRepo.GetGoodsTwo(ctx)
	if nil != err {
		return nil, err
	}
	resGoodsTwo := make([]*v1.UserInfoReply_List, 0)
	for _, v := range goodsTwo {
		resGoodsTwo = append(resGoodsTwo, &v1.UserInfoReply_List{
			Id:    v.ID,
			One:   v.Name,
			Two:   v.One,
			Three: v.Two,
			Four:  v.Amount,
			Five:  v.Three,
		})
	}

	goodsThree, err = uuc.ubRepo.GetGoodsThree(ctx)
	if nil != err {
		return nil, err
	}
	resGoodsThree := make([]*v1.UserInfoReply_List, 0)
	for _, v := range goodsThree {
		resGoodsThree = append(resGoodsThree, &v1.UserInfoReply_List{
			Id:    v.ID,
			One:   v.Name,
			Two:   v.One,
			Three: v.Two,
			Four:  v.Amount,
			Five:  v.Three,
		})
	}

	//var (
	//	userYesExchange []*UserBalanceRecord
	//)
	//userYesExchange, err = uuc.ubRepo.GetSystemYesterdayLocationReward(ctx)
	//if nil != err {
	//	return nil, err
	//}
	//
	//tmpExchangeTotal := float64(0)
	//for _, v := range userYesExchange {
	//	tmpExchangeTotal += v.AmountNew + v.AmountNewTwo
	//}
	//
	//var (
	//	buyReward []*Reward
	//)
	//buyReward, err = uuc.ubRepo.GetRewardBuyYes(ctx)
	//if nil != err {
	//	return nil, err
	//}
	//
	//tmpBuy := float64(0)
	//for _, v := range buyReward {
	//	tmpBuy += v.AmountNew
	//}

	// 推荐
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, myUser.ID)
	if nil == userRecommend {
		return nil, err
	}

	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}
		if _, ok := usersMap[myUserRecommendUserId]; ok {
			inviteUserAddress = usersMap[myUserRecommendUserId].Address
		}
	}

	// 推荐人
	var (
		userRecommends []*UserRecommend
		myLowUser      map[int64][]*UserRecommend
	)

	myLowUser = make(map[int64][]*UserRecommend, 0)
	userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	if nil != err {
		return nil, err
	}

	for _, vUr := range userRecommends {
		// 我的直推
		var (
			myUserRecommendUserIdTmp int64
			tmpRecommendUserIds      []string
		)

		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserIdTmp, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserIdTmp {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserIdTmp]; !ok {
			myLowUser[myUserRecommendUserIdTmp] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserIdTmp] = append(myLowUser[myUserRecommendUserIdTmp], vUr)
	}

	// 获取业绩
	tmpAreaMax := uint64(0)
	tmpAreaMin := uint64(0)
	//tmpMaxId := int64(0)
	tmpRecommendNum := uint64(0)
	for _, vMyLowUser := range myLowUser[myUser.ID] {
		if _, ok := usersMap[vMyLowUser.UserId]; !ok {
			continue
		}

		if usersMap[vMyLowUser.UserId].Amount > 0 {
			tmpRecommendNum += 1
		}

		if tmpAreaMax < uint64(usersMap[vMyLowUser.UserId].MyTotalAmount)+usersMap[vMyLowUser.UserId].Amount {
			tmpAreaMax = uint64(usersMap[vMyLowUser.UserId].MyTotalAmount) + usersMap[vMyLowUser.UserId].Amount
			//tmpMaxId = vMyLowUser.ID
		}
	}

	if 0 < tmpAreaMax {
		if uint64(myUser.MyTotalAmount) > tmpAreaMax {
			tmpAreaMin = uint64(myUser.MyTotalAmount) - tmpAreaMax
		}
		//for _, vMyLowUser := range myLowUser[myUser.ID] {
		//	if _, ok := usersMap[vMyLowUser.UserId]; !ok {
		//		continue
		//	}
		//
		//	if tmpMaxId != vMyLowUser.ID {
		//		tmpAreaMin += uint64(usersMap[vMyLowUser.UserId].MyTotalAmount) + usersMap[vMyLowUser.UserId].Amount
		//	}
		//
		//}
	}

	tmpLevel := uint64(0)
	if 1500000 <= tmpAreaMin {
		tmpLevel = 5
	} else if 500000 <= tmpAreaMin {
		tmpLevel = 4
	} else if 150000 <= tmpAreaMin {
		tmpLevel = 3
	} else if 50000 <= tmpAreaMin {
		tmpLevel = 2
	} else if 10000 <= tmpAreaMin {
		tmpLevel = 1
	}

	if 0 < myUser.VipAdmin {
		tmpLevel = uint64(myUser.VipAdmin)
	}

	var (
		buyRecord    []*BuyRecord
		tmpBuy       float64
		tmpAmountGet float64
		tmpSubAmount float64
	)
	buyRecord, err = uuc.repo.GetBuyRecordDoing(uint64(myUser.ID))
	if nil != err {
		return nil, err
	}

	t := time.Date(2026, 2, 18, 14, 0, 0, 0, time.UTC)
	for _, vBuyRecord := range buyRecord {
		tmpBuy += vBuyRecord.Amount

		one := vBuyRecord.AmountGet
		num := 2.5
		if vBuyRecord.CreatedAt.After(t) {
			amount := uint64(vBuyRecord.Amount)
			if 4999 <= amount && 15001 > amount {
				num = 3
			} else if 29999 <= amount && 50001 > amount {
				num = 3.5
			} else if 99999 <= amount && 150001 > amount {
				num = 4
			}
		}

		if vBuyRecord.Amount*num <= vBuyRecord.AmountGet {
			one = vBuyRecord.Amount * num
		} else {
			tmpSubAmount += vBuyRecord.Amount*num - one
		}

		tmpAmountGet += one
	}

	one := myUser.One
	if "1" == one {
		one = ""
	}
	two := myUser.Two
	if "1" == two {
		two = ""
	}
	three := myUser.Three
	if "1" == three {
		three = ""
	}
	four := myUser.Four
	if "1" == four {
		four = ""
	}
	five := myUser.Five
	if "1" == five {
		five = ""
	}
	six := myUser.Six
	if "1" == six {
		six = ""
	}
	seven := myUser.Seven
	if "1" == seven {
		seven = ""
	}

	return &v1.UserInfoReply{
		Status:            "ok",
		Level:             tmpLevel,
		LocationNum:       tmpRecommendNum,
		Total:             fmt.Sprintf("%d", uint64(myUser.MyTotalAmount)),
		Max:               fmt.Sprintf("%d", tmpAreaMax),
		Min:               fmt.Sprintf("%d", tmpAreaMin),
		InviteUserAddress: inviteUserAddress,
		Buy:               fmt.Sprintf("%.2f", tmpBuy),
		AmountGetSub:      fmt.Sprintf("%.2f", tmpSubAmount),
		AmountGet:         fmt.Sprintf("%.2f", tmpAmountGet),
		OutNum:            uint64(myUser.OutRate),
		Location:          fmt.Sprintf("%.2f", userBalance.LocationTotalFloat),
		Recommend:         fmt.Sprintf("%.2f", userBalance.RecommendTotalFloat),
		RecommendTwo:      fmt.Sprintf("%.2f", userBalance.RecommendTotalFloatTwo),
		Team:              fmt.Sprintf("%.2f", userBalance.AreaTotalFloat),
		TeamTwo:           fmt.Sprintf("%.2f", userBalance.AreaTotalFloatTwo),
		All:               fmt.Sprintf("%.2f", userBalance.AllFloat),
		AmountUsdt:        fmt.Sprintf("%.2f", userBalance.BalanceUsdtFloat),
		Usdt:              fmt.Sprintf("%.2f", userBalance.BalanceUsdtFloat),
		WithdrawRate:      withdrawRate,
		WithdrawMin:       withdrawMin,
		Raw:               fmt.Sprintf("%.2f", userBalance.BalanceRawFloat),
		WithdrawRateTwo:   withdrawRateTwo,
		WithdrawMinTwo:    withdrawMinTwo,
		One:               one,
		Two:               two,
		Three:             three,
		Four:              four,
		Five:              five,
		Six:               six,
		Seven:             seven,
		Notice:            notice,
		Goods:             resGoods,
		GoodsTwo:          resGoodsTwo,
		GoodsThree:        resGoodsThree,
		RawNew:            fmt.Sprintf("%.2f", userBalance.BalanceRawFloatNew),
	}, nil
}

func (uuc *UserUseCase) UserRecommend(ctx context.Context, req *v1.RecommendListRequest) (*v1.RecommendListReply, error) {
	var (
		userRecommend   *UserRecommend
		myUserRecommend []*UserRecommend
		recommendTotal  uint64
		user            *User
		err             error
	)

	res := make([]*v1.RecommendListReply_List, 0)

	if 0 >= len(req.Address) {
		return &v1.RecommendListReply{
			Recommends: res,
		}, nil
	}

	user, _ = uuc.repo.GetUserByAddress(ctx, req.Address)
	if nil == user {
		return &v1.RecommendListReply{
			Recommends: res,
		}, nil
	}

	// 推荐
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend {
		return &v1.RecommendListReply{
			Recommends: res,
		}, nil
	}

	myUserRecommend, err = uuc.urRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatInt(user.ID, 10))
	if nil == myUserRecommend || nil != err {
		return &v1.RecommendListReply{
			Recommends: res,
		}, nil
	}

	//var (
	//	totalNum int64
	//)
	//totalNum, err = uuc.urRepo.GetUserRecommendLikeCodeSum(ctx, userRecommend.RecommendCode+"D"+strconv.FormatInt(user.ID, 10))
	//if nil != err {
	//	return nil, err
	//}

	tmpUserIds := make([]int64, 0)
	for _, vMyUserRecommend := range myUserRecommend {
		tmpUserIds = append(tmpUserIds, vMyUserRecommend.UserId)
	}
	if 0 >= len(tmpUserIds) {
		return &v1.RecommendListReply{
			Recommends: res,
		}, nil
	}

	var (
		usersMap map[int64]*User
	)

	usersMap, err = uuc.repo.GetUserByUserIdsTwo(ctx, tmpUserIds)
	if nil == usersMap || nil != err {
		return &v1.RecommendListReply{
			Recommends: res,
		}, nil
	}

	if 0 >= len(usersMap) {
		return &v1.RecommendListReply{
			Recommends: res,
		}, nil
	}

	for _, vMyUserRecommend := range myUserRecommend {
		if _, ok := usersMap[vMyUserRecommend.UserId]; !ok {
			continue
		}

		recommendTotal++
		res = append(res, &v1.RecommendListReply_List{
			Address:  usersMap[vMyUserRecommend.UserId].Address,
			Amount:   fmt.Sprintf("%d", uint64(usersMap[vMyUserRecommend.UserId].MyTotalAmount)+usersMap[vMyUserRecommend.UserId].Amount),
			MyAmount: fmt.Sprintf("%d", usersMap[vMyUserRecommend.UserId].Amount),
		})
	}

	return &v1.RecommendListReply{
		Recommends: res,
	}, nil
}

func (uuc *UserUseCase) UserArea(ctx context.Context, req *v1.UserAreaRequest, user *User) (*v1.UserAreaReply, error) {
	var (
		err             error
		locationId      = req.LocationId
		Locations       []*LocationNew
		LocationRunning *LocationNew
	)

	res := make([]*v1.UserAreaReply_List, 0)
	if 0 >= locationId {
		Locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, user.ID)
		if nil != err {
			return nil, err
		}
		for _, vLocations := range Locations {
			if "running" == vLocations.Status {
				LocationRunning = vLocations
				break
			}
		}

		if nil == LocationRunning {
			return &v1.UserAreaReply{Area: res}, nil
		}

		locationId = LocationRunning.ID
	}

	var (
		myLowLocations []*LocationNew
	)

	myLowLocations, err = uuc.locationRepo.GetLocationsByTop(ctx, locationId)
	if nil != err {
		return nil, err
	}

	for _, vMyLowLocations := range myLowLocations {
		var (
			userLow           *User
			tmpMyLowLocations []*LocationNew
		)

		userLow, err = uuc.repo.GetUserById(ctx, vMyLowLocations.UserId)
		if nil != err {
			continue
		}

		tmpMyLowLocations, err = uuc.locationRepo.GetLocationsByTop(ctx, vMyLowLocations.ID)
		if nil != err {
			return nil, err
		}

		var address1 string
		if 20 <= len(userLow.Address) {
			address1 = userLow.Address[:6] + "..." + userLow.Address[len(userLow.Address)-4:]
		}

		res = append(res, &v1.UserAreaReply_List{
			Address:    address1,
			LocationId: vMyLowLocations.ID,
			CountLow:   int64(len(tmpMyLowLocations)),
		})
	}

	return &v1.UserAreaReply{Area: res}, nil
}

func (uuc *UserUseCase) UserInfoOld(ctx context.Context, user *User) (*v1.UserInfoReply, error) {
	//var (
	//	myUser                  *User
	//	userInfo                *UserInfo
	//	locations               []*LocationNew
	//	myLastStopLocations     []*LocationNew
	//	userBalance             *UserBalance
	//	userRecommend           *UserRecommend
	//	userRecommends          []*UserRecommend
	//	userRewards             []*Reward
	//	userRewardTotal         int64
	//	userRewardWithdrawTotal int64
	//	encodeString            string
	//	recommendTeamNum        int64
	//	myCode                  string
	//	amount                  = "0"
	//	amount2                 = "0"
	//	configs                 []*Config
	//	myLastLocationCurrent   int64
	//	withdrawAmount          int64
	//	withdrawAmount2         int64
	//	myUserRecommendUserId   int64
	//	inviteUserAddress       string
	//	myRecommendUser         *User
	//	withdrawRate            int64
	//	myLocations             []*v1.UserInfoReply_List
	//	myLocations2            []*v1.UserInfoReply_List22
	//	allRewardList           []*v1.UserInfoReply_List9
	//	timeAgain               int64
	//	err                     error
	//)
	//
	//// 配置
	//configs, err = uuc.configRepo.GetConfigByKeys(ctx,
	//	"time_again",
	//)
	//if nil != configs {
	//	for _, vConfig := range configs {
	//		if "time_again" == vConfig.KeyName {
	//			timeAgain, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		}
	//	}
	//}
	//
	//myUser, err = uuc.repo.GetUserById(ctx, user.ID)
	//if nil != err {
	//	return nil, err
	//}
	//userInfo, err = uuc.uiRepo.GetUserInfoByUserId(ctx, myUser.ID)
	//if nil != err {
	//	return nil, err
	//}
	//locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, myUser.ID)
	//if nil != locations && 0 < len(locations) {
	//	for _, v := range locations {
	//		var tmp int64
	//		if v.Current <= v.CurrentMax {
	//			tmp = v.CurrentMax - v.Current
	//		}
	//		if "running" == v.Status {
	//			amount = fmt.Sprintf("%.4f", float64(tmp)/float64(100000))
	//		}
	//
	//		myLocations = append(myLocations, &v1.UserInfoReply_List{
	//			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//			Amount:    fmt.Sprintf("%.2f", float64(v.Usdt)/float64(100000)),
	//			Num:       v.Num,
	//			AmountMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
	//		})
	//	}
	//
	//}
	//
	//// 冻结
	//myLastStopLocations, err = uuc.locationRepo.GetMyStopLocationsLast(ctx, myUser.ID)
	//now := time.Now().UTC()
	//tmpNow := now.Add(8 * time.Hour)
	//if nil != myLastStopLocations {
	//	for _, vMyLastStopLocations := range myLastStopLocations {
	//		if tmpNow.Before(vMyLastStopLocations.StopDate.Add(time.Duration(timeAgain) * time.Minute)) {
	//			myLastLocationCurrent += vMyLastStopLocations.Current - vMyLastStopLocations.CurrentMax
	//		}
	//	}
	//}
	//
	//var (
	//	locations2 []*LocationNew
	//)
	//locations2, err = uuc.locationRepo.GetLocationsByUserId2(ctx, myUser.ID)
	//if nil != locations2 && 0 < len(locations2) {
	//	for _, v := range locations2 {
	//		var tmp int64
	//		if v.Current <= v.CurrentMax {
	//			tmp = v.CurrentMax - v.Current
	//		}
	//
	//		if "running" == v.Status {
	//			amount2 = fmt.Sprintf("%.4f", float64(tmp)/float64(100000))
	//		}
	//
	//		myLocations2 = append(myLocations2, &v1.UserInfoReply_List22{
	//			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//			Amount:    fmt.Sprintf("%.2f", float64(v.Usdt)/float64(100000)),
	//			AmountMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
	//		})
	//	}
	//
	//}
	//
	//userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUser.ID)
	//if nil != err {
	//	return nil, err
	//}
	//
	//userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, myUser.ID)
	//if nil == userRecommend {
	//	return nil, err
	//}
	//
	//myCode = "D" + strconv.FormatInt(myUser.ID, 10)
	//codeByte := []byte(myCode)
	//encodeString = base64.StdEncoding.EncodeToString(codeByte)
	//if "" != userRecommend.RecommendCode {
	//	tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
	//	if 2 <= len(tmpRecommendUserIds) {
	//		myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
	//	}
	//	myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
	//	if nil != err {
	//		return nil, err
	//	}
	//	inviteUserAddress = myRecommendUser.Address
	//	myCode = userRecommend.RecommendCode + myCode
	//}
	//
	//// 团队
	//var (
	//	teamUserIds        []int64
	//	teamUsers          map[int64]*User
	//	teamUserInfos      map[int64]*UserInfo
	//	teamUserAddresses  []*v1.UserInfoReply_List7
	//	recommendAddresses []*v1.UserInfoReply_List11
	//	teamLocations      map[int64][]*Location
	//	recommendUserIds   map[int64]int64
	//	userBalanceMap     map[int64]*UserBalance
	//)
	//recommendUserIds = make(map[int64]int64, 0)
	//userRecommends, err = uuc.urRepo.GetUserRecommendLikeCode(ctx, myCode)
	//if nil != userRecommends {
	//	for _, vUserRecommends := range userRecommends {
	//		if myCode == vUserRecommends.RecommendCode {
	//			recommendUserIds[vUserRecommends.UserId] = vUserRecommends.UserId
	//		}
	//		teamUserIds = append(teamUserIds, vUserRecommends.UserId)
	//	}
	//
	//	// 用户信息
	//	recommendTeamNum = int64(len(userRecommends))
	//	teamUsers, _ = uuc.repo.GetUserByUserIds(ctx, teamUserIds...)
	//	teamUserInfos, _ = uuc.uiRepo.GetUserInfoByUserIds(ctx, teamUserIds...)
	//	teamLocations, _ = uuc.locationRepo.GetLocationMapByIds(ctx, teamUserIds...)
	//	userBalanceMap, _ = uuc.ubRepo.GetUserBalanceByUserIds(ctx, teamUserIds...)
	//	if nil != teamUsers {
	//		for _, vTeamUsers := range teamUsers {
	//			var locationAmount int64
	//			if _, ok := teamUserInfos[vTeamUsers.ID]; !ok {
	//				continue
	//			}
	//
	//			if _, ok := userBalanceMap[vTeamUsers.ID]; !ok {
	//				continue
	//			}
	//
	//			if _, ok := teamLocations[vTeamUsers.ID]; ok {
	//
	//				for _, vTeamLocations := range teamLocations[vTeamUsers.ID] {
	//					locationAmount += vTeamLocations.Usdt
	//				}
	//			}
	//			if _, ok := recommendUserIds[vTeamUsers.ID]; ok {
	//				recommendAddresses = append(recommendAddresses, &v1.UserInfoReply_List11{
	//					Address: vTeamUsers.Address,
	//					Usdt:    fmt.Sprintf("%.2f", float64(locationAmount)/float64(100000)),
	//					Vip:     teamUserInfos[vTeamUsers.ID].Vip,
	//				})
	//			}
	//
	//			teamUserAddresses = append(teamUserAddresses, &v1.UserInfoReply_List7{
	//				Address: vTeamUsers.Address,
	//				Usdt:    fmt.Sprintf("%.2f", float64(locationAmount)/float64(100000)),
	//				Vip:     teamUserInfos[vTeamUsers.ID].Vip,
	//			})
	//		}
	//	}
	//}
	//
	//var (
	//	totalDeposit      int64
	//	userBalanceRecord []*UserBalanceRecord
	//	depositList       []*v1.UserInfoReply_List13
	//)
	//
	//userBalanceRecord, _ = uuc.ubRepo.GetUserBalanceRecordsByUserId(ctx, myUser.ID)
	//for _, vUserBalanceRecord := range userBalanceRecord {
	//	totalDeposit += vUserBalanceRecord.Amount
	//	depositList = append(depositList, &v1.UserInfoReply_List13{
	//		CreatedAt: vUserBalanceRecord.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//		Amount:    fmt.Sprintf("%.4f", float64(vUserBalanceRecord.Amount)/float64(100000)),
	//	})
	//}
	//
	//// 累计奖励
	//var (
	//	locationRewardList            []*v1.UserInfoReply_List2
	//	recommendRewardList           []*v1.UserInfoReply_List5
	//	locationTotal                 int64
	//	yesterdayLocationTotal        int64
	//	recommendRewardTotal          int64
	//	recommendRewardDhbTotal       int64
	//	yesterdayRecommendRewardTotal int64
	//)
	//
	//var startDate time.Time
	//var endDate time.Time
	//if 16 <= now.Hour() {
	//	startDate = now.AddDate(0, 0, -1)
	//	endDate = startDate.AddDate(0, 0, 1)
	//} else {
	//	startDate = now.AddDate(0, 0, -2)
	//	endDate = startDate.AddDate(0, 0, 1)
	//}
	//yesterdayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	//yesterdayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)
	//
	//fmt.Println(now, yesterdayStart, yesterdayEnd)
	//userRewards, err = uuc.ubRepo.GetUserRewardByUserId(ctx, myUser.ID)
	//if nil != userRewards {
	//	for _, vUserReward := range userRewards {
	//		if "location" == vUserReward.Reason {
	//			locationTotal += vUserReward.Amount
	//			if vUserReward.CreatedAt.Before(yesterdayEnd) && vUserReward.CreatedAt.After(yesterdayStart) {
	//				yesterdayLocationTotal += vUserReward.Amount
	//			}
	//			locationRewardList = append(locationRewardList, &v1.UserInfoReply_List2{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//
	//			userRewardTotal += vUserReward.Amount
	//			allRewardList = append(allRewardList, &v1.UserInfoReply_List9{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//		} else if "recommend" == vUserReward.Reason {
	//
	//			recommendRewardList = append(recommendRewardList, &v1.UserInfoReply_List5{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//
	//			recommendRewardTotal += vUserReward.Amount
	//			if vUserReward.CreatedAt.Before(yesterdayEnd) && vUserReward.CreatedAt.After(yesterdayStart) {
	//				yesterdayRecommendRewardTotal += vUserReward.Amount
	//			}
	//
	//			userRewardTotal += vUserReward.Amount
	//			allRewardList = append(allRewardList, &v1.UserInfoReply_List9{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//		} else if "reward_withdraw" == vUserReward.Reason {
	//			userRewardTotal += vUserReward.Amount
	//			userRewardWithdrawTotal += vUserReward.Amount
	//		} else if "recommend_token" == vUserReward.Reason {
	//			recommendRewardDhbTotal += vUserReward.Amount
	//		}
	//	}
	//}
	//
	//var (
	//	withdraws []*Withdraw
	//)
	//withdraws, err = uuc.ubRepo.GetWithdrawByUserId2(ctx, user.ID)
	//if nil != err {
	//	return nil, err
	//}
	//for _, v := range withdraws {
	//	if "usdt" == v.Type {
	//		withdrawAmount += v.RelAmount
	//	}
	//	if "usdt_2" == v.Type {
	//		withdrawAmount2 += v.RelAmount
	//	}
	//}
	//
	//return &v1.UserInfoReply{
	//	Address:                 myUser.Address,
	//	Level:                   userInfo.Vip,
	//	Amount:                  amount,
	//	Amount2:                 amount2,
	//	LocationList2:           myLocations2,
	//	AmountB:                 fmt.Sprintf("%.2f", float64(myLastLocationCurrent)/float64(100000)),
	//	DepositList:             depositList,
	//	BalanceUsdt:             fmt.Sprintf("%.2f", float64(userBalance.BalanceUsdt)/float64(100000)),
	//	BalanceUsdt2:            fmt.Sprintf("%.2f", float64(userBalance.BalanceUsdt2)/float64(100000)),
	//	BalanceDhb:              fmt.Sprintf("%.2f", float64(userBalance.BalanceDhb)/float64(100000)),
	//	InviteUrl:               encodeString,
	//	RecommendNum:            userInfo.HistoryRecommend,
	//	RecommendTeamNum:        recommendTeamNum,
	//	Total:                   fmt.Sprintf("%.2f", float64(userRewardTotal)/float64(100000)),
	//	RewardWithdraw:          fmt.Sprintf("%.2f", float64(userRewardWithdrawTotal)/float64(100000)),
	//	WithdrawAmount:          fmt.Sprintf("%.2f", float64(withdrawAmount)/float64(100000)),
	//	WithdrawAmount2:         fmt.Sprintf("%.2f", float64(withdrawAmount2)/float64(100000)),
	//	LocationTotal:           fmt.Sprintf("%.2f", float64(locationTotal)/float64(100000)),
	//	Account:                 "0xAfC39F3592A1024144D1ba6DC256397F4DbfbE2f",
	//	LocationList:            myLocations,
	//	TeamAddressList:         teamUserAddresses,
	//	AllRewardList:           allRewardList,
	//	InviteUserAddress:       inviteUserAddress,
	//	RecommendAddressList:    recommendAddresses,
	//	WithdrawRate:            withdrawRate,
	//	RecommendRewardTotal:    fmt.Sprintf("%.2f", float64(recommendRewardTotal)/float64(100000)),
	//	RecommendRewardDhbTotal: fmt.Sprintf("%.2f", float64(recommendRewardDhbTotal)/float64(100000)),
	//	TotalDeposit:            fmt.Sprintf("%.2f", float64(totalDeposit)/float64(100000)),
	//	WithdrawAll:             fmt.Sprintf("%.2f", float64(withdrawAmount+withdrawAmount2)/float64(100000)),
	//}, nil
	return nil, nil

}

func (uuc *UserUseCase) RewardList(ctx context.Context, req *v1.RewardListRequest, user *User) (*v1.RewardListReply, error) {
	res := make([]*v1.RewardListReply_List, 0)

	var (
		userRewards []*Reward
		count       int64
		reason      string
		err         error
	)

	if 1 == req.ReqType {
		reason = "buy"
	} else if 2 == req.ReqType {
		reason = "location"
	} else if 3 == req.ReqType {
		reason = "recommend"
	} else if 4 == req.ReqType {
		reason = "recommend_two"
	} else if 5 == req.ReqType {
		reason = "area"
	} else if 6 == req.ReqType {
		reason = "area_two"
	} else if 7 == req.ReqType {
		reason = "all"
	} else if 8 == req.ReqType {
		reason = "send"
	}

	userRewards, err, count = uuc.ubRepo.GetUserRewardByUserIdPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	}, user.ID, reason)
	if nil != err {
		return &v1.RewardListReply{
			Status: "ok",
			Count:  uint64(count),
			List:   res,
		}, err
	}

	for _, vUserReward := range userRewards {
		res = append(res, &v1.RewardListReply_List{
			CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.4f", vUserReward.AmountNew),
			Address:   vUserReward.Address,
			AmountTwo: fmt.Sprintf("%.4f", vUserReward.AmountNewTwo),
			Num:       uint64(vUserReward.TypeRecordId),
		})
	}

	return &v1.RewardListReply{
		Status: "ok",
		Count:  uint64(count),
		List:   res,
	}, nil
}

func (uuc *UserUseCase) RecommendRewardList(ctx context.Context, user *User) (*v1.RecommendRewardListReply, error) {

	res := &v1.RecommendRewardListReply{
		Rewards: make([]*v1.RecommendRewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) FeeRewardList(ctx context.Context, user *User) (*v1.FeeRewardListReply, error) {
	res := &v1.FeeRewardListReply{
		Rewards: make([]*v1.FeeRewardListReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) TranList(ctx context.Context, user *User, reqTypeCoin string, reqTran string) (*v1.TranListReply, error) {

	var (
		userBalanceRecord []*UserBalanceRecord
		typeCoin          = "usdt"
		tran              = "tran"
		err               error
	)

	res := &v1.TranListReply{
		Tran: make([]*v1.TranListReply_List, 0),
	}

	if "" != reqTypeCoin {
		typeCoin = reqTypeCoin
	}

	if "tran_to" == reqTran {
		tran = reqTran
	}

	userBalanceRecord, err = uuc.ubRepo.GetUserBalanceRecordByUserId(ctx, user.ID, typeCoin, tran)
	if nil != err {
		return res, err
	}

	for _, v := range userBalanceRecord {
		res.Tran = append(res.Tran, &v1.TranListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", float64(v.Amount)/float64(100000)),
		})
	}

	return res, nil
}

func (uuc *UserUseCase) SetTodayList(ctx context.Context, user *User) (*v1.SetTodayListReply, error) {
	res := make([]*v1.SetTodayListReply_List, 0)
	var (
		stake []*Stake
		err   error
	)

	stake, err = uuc.ubRepo.GetStakeByUserId(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	for _, vStake := range stake {
		res = append(res, &v1.SetTodayListReply_List{
			CreatedAt: vStake.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", vStake.Amount),
		})
	}

	return &v1.SetTodayListReply{
		Status: "ok",
		Count:  uint64(len(stake)),
		List:   res,
	}, nil
}

func (uuc *UserUseCase) WithdrawList(ctx context.Context, req *v1.WithdrawListRequest, user *User, reqTypeCoin string) (*v1.WithdrawListReply, error) {

	//res := make([]*v1.WithdrawListReply_List, 0)
	//res = append(res, &v1.WithdrawListReply_List{
	//	CreatedAt: "2006-01-02 15:04:05",
	//	Amount:    20,
	//}, &v1.WithdrawListReply_List{
	//	CreatedAt: "2006-01-03 15:04:05",
	//	Amount:    10,
	//})
	//return &v1.WithdrawListReply{
	//	Status: "ok",
	//	Count:  2,
	//	List:   nil,
	//}, nil

	var (
		withdraws []*Withdraw
		count     int64
		err       error
	)

	res := &v1.WithdrawListReply{
		Count: 0,
		List:  make([]*v1.WithdrawListReply_List, 0),
	}

	coinType := "USDT"
	if 2 == req.CoinType {
		coinType = "RAW"
	}

	withdraws, count, err = uuc.ubRepo.GetWithdrawByUserId(ctx, user.ID, coinType, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return res, err
	}

	for _, v := range withdraws {
		res.List = append(res.List, &v1.WithdrawListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", v.AmountNew),
		})
	}

	res.Count = uint64(count)
	return res, nil
}

func (uuc *UserUseCase) SetToday(ctx context.Context, req *v1.SetTodayRequest, user *User) (*v1.SetTodayReply, error) {
	var (
		stake []*Stake
		err   error
	)

	stake, err = uuc.ubRepo.GetStakeByUserId(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if 30 <= len(stake) {
		return &v1.SetTodayReply{
			Status: "签到已满30天",
		}, nil
	}

	now := time.Now().UTC()
	var lasUpdated time.Time
	if 16 <= now.Hour() {
		lasUpdated = now
	} else {
		lasUpdated = now.AddDate(0, 0, -1)
	}
	todayStart := time.Date(lasUpdated.Year(), lasUpdated.Month(), lasUpdated.Day(), 16, 0, 0, 0, time.UTC)

	//for _, vStake := range stake {
	//	if todayStart.Unix() <= vStake.Day {
	//		return &v1.SetTodayReply{
	//			Status: "今日已经签到",
	//		}, nil
	//	}
	//}

	amount := float64(1)
	if 4 == len(stake) {
		amount = float64(4)
	} else if 9 == len(stake) {
		amount = float64(6)
	} else if 19 == len(stake) {
		amount = float64(9)
	} else if 29 == len(stake) {
		amount = float64(16)
	}

	err = uuc.ubRepo.SetToday(ctx, user.ID, amount, todayStart.Unix())
	if nil != err {
		return &v1.SetTodayReply{
			Status: "签到错误",
		}, nil
	}

	return &v1.SetTodayReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) OrderList(ctx context.Context, req *v1.OrderListRequest, user *User) (*v1.OrderListReply, error) {
	res := make([]*v1.OrderListReply_List, 0)

	var (
		myUser    *User
		buyRecord []*BuyRecord
		goods     []*Good
		goodsMap  map[int64]*Good
		count     int64
		err       error
	)
	myUser, err = uuc.repo.GetUserById(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	buyRecord, count, err = uuc.repo.GetBuyRecord(ctx, uint64(myUser.ID), &Pagination{PageNum: int(req.Page), PageSize: 20})
	if nil != err {
		return &v1.OrderListReply{
			Status: "err",
			Count:  0,
			List:   res,
		}, nil
	}

	goods, err = uuc.ubRepo.GetGoodsAll(ctx)
	if nil != err {
		return nil, err
	}
	goodsMap = make(map[int64]*Good, 0)
	for _, v := range goods {
		goodsMap[v.ID] = v
	}

	num := 2.5
	numStr := "2.5"
	t := time.Date(2026, 2, 18, 14, 0, 0, 0, time.UTC)
	for _, vBuyRecord := range buyRecord {
		tmpAmountGet := vBuyRecord.AmountGet
		tmpAmountLast := float64(0)

		if vBuyRecord.CreatedAt.After(t) {
			amount := uint64(vBuyRecord.Amount)
			if 4999 <= amount && 15001 > amount {
				num = 3
				numStr = "3"
			} else if 29999 <= amount && 50001 > amount {
				num = 3.5
				numStr = "3.5"
			} else if 99999 <= amount && 150001 > amount {
				num = 4
				numStr = "4"
			}
		}

		if vBuyRecord.AmountGet >= vBuyRecord.Amount*num {
			tmpAmountGet = vBuyRecord.Amount * num
		} else {
			tmpAmountLast = vBuyRecord.Amount*num - vBuyRecord.AmountGet
		}

		oneTmp := ""
		if "1" != vBuyRecord.One {
			oneTmp = vBuyRecord.One
		}
		twoTmp := ""
		if "1" != vBuyRecord.One {
			twoTmp = vBuyRecord.Two
		}
		threeTmp := ""
		if "1" != vBuyRecord.One {
			threeTmp = vBuyRecord.Three
		}

		fourTmp := ""
		fiveTmp := ""
		sixTmp := ""
		if 0 != vBuyRecord.Four {
			if _, ok := goodsMap[vBuyRecord.Four]; ok {
				fourTmp = goodsMap[vBuyRecord.Four].Name
				fiveTmp = goodsMap[vBuyRecord.Four].One
				sixTmp = goodsMap[vBuyRecord.Four].Two
			}
		}

		res = append(res, &v1.OrderListReply_List{
			CreatedAt:  vBuyRecord.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:     fmt.Sprintf("%.2f", vBuyRecord.Amount),
			Status:     uint64(vBuyRecord.Status),
			AmountGet:  fmt.Sprintf("%.2f", tmpAmountGet),
			AmountLast: fmt.Sprintf("%.2f", tmpAmountLast),
			Num:        numStr,
			One:        oneTmp,
			Two:        twoTmp,
			Three:      threeTmp,
			Four:       fourTmp,
			Five:       fiveTmp,
			Six:        sixTmp,
		})
	}

	return &v1.OrderListReply{
		Status: "ok",
		Count:  uint64(count),
		List:   res,
	}, nil
}

func (uuc *UserUseCase) TradeList(ctx context.Context, user *User) (*v1.TradeListReply, error) {

	var (
		trades []*Trade
		err    error
	)

	res := &v1.TradeListReply{
		Trade: make([]*v1.TradeListReply_List, 0),
	}

	trades, err = uuc.ubRepo.GetTradeByUserId(ctx, user.ID)
	if nil != err {
		return res, err
	}

	for _, v := range trades {
		res.Trade = append(res.Trade, &v1.TradeListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			AmountCsd: fmt.Sprintf("%.2f", float64(v.AmountCsd)/float64(100000)),
			AmountHbs: fmt.Sprintf("%.2f", float64(v.AmountHbs)/float64(100000)),
			Status:    v.Status,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) UnStake(ctx context.Context, req *v1.UnStakeRequest, user *User) (*v1.UnStakeReply, error) {
	var (
		err    error
		stakes []*Stake
	)

	stakes, err = uuc.ubRepo.GetStakeByUserId(ctx, user.ID)
	if nil != err {
		return &v1.UnStakeReply{
			Status: "错误",
		}, nil
	}

	for _, v := range stakes {
		if 0 != v.Status {
			continue
		}

		if uint64(v.ID) != req.SendBody.Id {
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.ubRepo.UnStakeAmount(ctx, user.ID, v.ID, v.Amount) // 提现
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &v1.UnStakeReply{
				Status: "错误",
			}, nil
		}

		break
	}

	return &v1.UnStakeReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Stake(ctx context.Context, req *v1.StakeRequest, user *User) (*v1.StakeReply, error) {
	var (
		err         error
		userBalance *UserBalance
		day         = int64(10)
	)

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return &v1.StakeReply{
			Status: "错误",
		}, nil
	}

	amountFloat := float64(req.SendBody.Amount)
	if userBalance.BalanceRawFloat < amountFloat {
		return &v1.StakeReply{
			Status: "余额不足",
		}, nil
	}

	if 1 > amountFloat {
		return &v1.StakeReply{
			Status: "金额错误",
		}, nil
	}

	if 10 == req.SendBody.Day {
		day = 10
	} else if 30 == req.SendBody.Day {
		day = 30
	} else {
		return &v1.StakeReply{
			Status: "错误天数",
		}, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.ubRepo.StakeAmount(ctx, user.ID, amountFloat, day) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return &v1.StakeReply{
			Status: "错误",
		}, nil
	}

	return &v1.StakeReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) AmountTo(ctx context.Context, req *v1.AmountToRequest, user *User) (*v1.AmountToReply, error) {
	var (
		err         error
		userBalance *UserBalance
		toUser      *User
	)

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return &v1.AmountToReply{
			Status: "错误",
		}, nil
	}

	amountFloat := float64(req.SendBody.Amount)
	if 1 > amountFloat {
		return &v1.AmountToReply{
			Status: "最少1brc20|min 1 brc20",
		}, nil
	}

	if userBalance.BalanceRawFloatNew < amountFloat {
		return &v1.AmountToReply{
			Status: "brc20不足|brc20 not enough",
		}, nil
	}

	toUser, err = uuc.repo.GetUserByAddress(ctx, req.SendBody.Address)
	if nil == toUser || nil != err {
		return &v1.AmountToReply{
			Status: "不存在用户|not found address",
		}, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.ubRepo.ToAddressAmountKsdt(ctx, user.ID, toUser.ID, amountFloat, toUser.Address) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return &v1.AmountToReply{
			Status: "错误",
		}, nil
	}

	return &v1.AmountToReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Buy(ctx context.Context, req *v1.BuyRequest, user *User) (*v1.BuyReply, error) {
	t := time.Date(2026, 2, 18, 14, 0, 0, 0, time.UTC)

	// 推荐人
	var (
		err               error
		configs           []*Config
		recommend         float64
		uRate             float64
		bRate             float64
		sendRate          float64
		sendRecommendRate float64
		goodRate          float64
		priceBrc          float64
	)

	// 配置
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"u_rate",
		"b_rate",
		"recommend",
		"send_rate",
		"buy_recommend_rate",
		"good_rate",
		"price_brc",
	)
	if nil != err || nil == configs {
		return &v1.BuyReply{
			Status: "稍后重试|err wait",
		}, nil
	}

	for _, vConfig := range configs {
		if "recommend" == vConfig.KeyName {
			recommend, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "u_rate" == vConfig.KeyName {
			uRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "b_rate" == vConfig.KeyName {
			bRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "good_rate" == vConfig.KeyName {
			goodRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "price_brc" == vConfig.KeyName {
			priceBrc, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "send_rate" == vConfig.KeyName {
			sendRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_recommend_rate" == vConfig.KeyName {
			sendRecommendRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	var (
		goods    []*Good
		goodsMap map[int64]*Good
		goodsBuy *Good
	)
	goods, err = uuc.ubRepo.GetGoodsOnline(ctx)
	if nil != err {
		return &v1.BuyReply{
			Status: "稍后重试|err wait",
		}, nil
	}
	goodsMap = make(map[int64]*Good, 0)
	for _, v := range goods {
		goodsMap[v.ID] = v
	}

	if _, ok := goodsMap[req.SendBody.Id]; !ok {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}
	goodsBuy = goodsMap[req.SendBody.Id]

	amount := goodsBuy.Amount
	if 100 == amount {
		amount = 100
	} else if 300 == amount {
		amount = 300
	} else if 500 == amount {
		amount = 500
	} else if 1000 == amount {
		amount = 1000
	} else if 5000 == amount {
		amount = 5000
	} else if 10000 == amount {
		amount = 10000
	} else if 15000 == amount {
		amount = 15000
	} else if 30000 == amount {
		amount = 30000
	} else if 50000 == amount {
		amount = 50000
	} else if 100000 == amount {
		amount = 100000
	} else if 150000 == amount {
		amount = 150000
	} else {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}

	var (
		users    []*User
		usersMap map[int64]*User
	)
	users, err = uuc.ubRepo.GetAllUsersB(ctx)
	if nil == users {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}

	usersMap = make(map[int64]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	if _, ok := usersMap[user.ID]; !ok {
		fmt.Println("不存在用户")
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}
	user = usersMap[user.ID]

	var (
		amountRel    = float64(amount)
		amountRelBrc float64
		userBalance  *UserBalance
	)

	if 1 != req.SendBody.Full {
		amountRel = float64(amount) * goodRate
		amountRelBrc = (float64(amount) - float64(amount)*goodRate) / priceBrc
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil == userBalance || nil != err {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}

	if 1 < amountRelBrc && amountRelBrc > userBalance.BalanceRawFloatNew {
		return &v1.BuyReply{
			Status: "brc20余额不足|brc20 not enough",
		}, nil
	}

	if 1 < amountRel && amountRel > userBalance.BalanceUsdtFloat {
		return &v1.BuyReply{
			Status: "usdt余额不足|deposit usdt not enough",
		}, nil
	}

	var (
		userRecommends    []*UserRecommend
		userRecommendsMap map[int64]*UserRecommend
	)
	userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	if nil != err {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}

	myLowUser := make(map[int64][]*UserRecommend, 0)
	userRecommendsMap = make(map[int64]*UserRecommend, 0)
	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId  int64
			tmpRecommendUserIdsTmp []string
		)

		tmpRecommendUserIdsTmp = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIdsTmp) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIdsTmp[len(tmpRecommendUserIdsTmp)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	var (
		buyRecords []*BuyRecord
	)
	buyRecords, err = uuc.uiRepo.GetBuyRecord(ctx, 0)
	if nil != err {
		fmt.Println("认购数据查询错误")
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}

	userBuyRecords := make(map[int64][]*BuyRecord, 0)
	for _, v := range buyRecords {
		if _, ok := userBuyRecords[v.UserId]; !ok {
			userBuyRecords[v.UserId] = make([]*BuyRecord, 0)
		}

		userBuyRecords[v.UserId] = append(userBuyRecords[v.UserId], v)
	}

	// 推荐人
	var (
		userRecommend         *UserRecommend
		myUserRecommendUserId int64
		tmpRecommendUserIds   []string
	)
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil != err {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}
	}

	if 0 < myUserRecommendUserId {
		//if _, ok := usersMap[myUserRecommendUserId]; ok {
		//	if 0 >= usersMap[myUserRecommendUserId].Amount && 0 >= usersMap[myUserRecommendUserId].OutRate {
		//		return err
		//	}
		//}
	}

	one := ""
	two := ""
	three := ""
	if "1" != user.One {
		one += user.One
	}
	if "1" != user.Two {
		one += user.Two
	}
	if "1" != user.Three {
		one += user.Three
	}
	if "1" != user.Four {
		one += user.Four
	}
	if "1" != user.Five {
		one += user.Five
	}
	if "1" != user.Six {
		two = user.Six
	}
	if "1" != user.Seven {
		three = user.Seven
	}

	four := goodsBuy.ID
	// 入金
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.uiRepo.UpdateUserNewTwoNewTwo(ctx, user.ID, amount, amountRel, amountRelBrc, float64(amount)*sendRate, one, two, three, four)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "错误投资3", amount)
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}

	totalTmp := len(tmpRecommendUserIds) - 1
	for i := totalTmp; i >= 0; i-- {

		tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		if _, ok := usersMap[tmpUserId]; !ok {
			fmt.Println("buy遍历，信息缺失,user：", err, tmpUserId)
			continue
		}

		usersMap[tmpUserId].MyTotalAmount = usersMap[tmpUserId].MyTotalAmount + float64(amount)
	}

	for i := totalTmp; i >= 0; i-- {

		tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		if _, ok := usersMap[tmpUserId]; !ok {
			fmt.Println("buy遍历，信息缺失,user：", err, tmpUserId)
			continue
		}

		tmpMax := uint64(0)
		tmpAreaMin := uint64(0)
		for _, vV := range myLowUser[tmpUserId] {
			if _, ok2 := usersMap[vV.UserId]; ok2 {
				if tmpMax < uint64(usersMap[vV.UserId].MyTotalAmount)+usersMap[vV.UserId].AmountSelf {
					tmpMax = uint64(usersMap[vV.UserId].MyTotalAmount) + usersMap[vV.UserId].AmountSelf
				}
			}
		}

		tmpSend := float64(0)
		if 0 < tmpMax {
			if uint64(usersMap[tmpUserId].MyTotalAmount) > tmpMax {
				if 0 < usersMap[tmpUserId].Amount {
					tmpAreaMin = uint64(usersMap[tmpUserId].MyTotalAmount) - tmpMax
					if 1500000 <= tmpAreaMin && 2 == usersMap[tmpUserId].Last {
						tmpSend = 50000
					} else if 500000 <= tmpAreaMin && 1 == usersMap[tmpUserId].Last {
						tmpSend = 30000
					} else if 150000 <= tmpAreaMin && 0 == usersMap[tmpUserId].Last {
						tmpSend = 10000
					}
				}
			}
		}

		// 增加业绩
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.uiRepo.UpdateUserMyTotalAmountAdd(ctx, tmpUserId, float64(amount), tmpSend)
			if err != nil {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("遍历业绩：", err, tmpUserId, user)
			continue
		}

		if 1 == user.LockReward {
			continue
		}

		// 直推
		tmpRecommendUser := usersMap[tmpUserId]
		if i == totalTmp {
			// 入金
			if 0 < tmpRecommendUser.Amount {
				var (
					num = 2.5
				)

				amountRecommendTmp := float64(amount) * recommend
				if _, ok := userBuyRecords[tmpUserId]; !ok {
					continue
				}

				for _, vUserRecords := range userBuyRecords[tmpUserId] {

					if vUserRecords.CreatedAt.After(t) {
						amountB := uint64(vUserRecords.Amount)
						if 4999 <= amountB && 15001 > amountB {
							num = 3
						} else if 29999 <= amountB && 50001 > amountB {
							num = 3.5
						} else if 99999 <= amountB && 150001 > amountB {
							num = 4
						}
					}

					if vUserRecords.Amount*num <= vUserRecords.AmountGet {
						fmt.Println("错误的数据，已经最大却没停，all充值推荐", vUserRecords)
						continue
					}

					var (
						stopRecommend bool
					)
					tmpU := amountRecommendTmp
					if tmpU+vUserRecords.AmountGet >= vUserRecords.Amount*num {
						tmpU = math.Abs(vUserRecords.Amount*num - vUserRecords.AmountGet)
						vUserRecords.AmountGet = vUserRecords.Amount * num
						stopRecommend = true
					}

					amountRecommendTmp -= tmpU

					tmpURel := math.Round(tmpU*uRate*10000000) / 10000000
					tmpB := math.Round(tmpU*bRate*10000000) / 10000000

					fmt.Println("直推奖奖励：", amount, amountRecommendTmp, tmpU, tmpB, recommend, amountRecommendTmp, user, tmpRecommendUser)
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						err = uuc.uiRepo.UpdateUserRewardRecommend2(ctx, vUserRecords.ID, tmpUserId, tmpURel, tmpB, tmpU, vUserRecords.Amount, stopRecommend, user.Address)
						if err != nil {
							fmt.Println("错误分红直推：", err)
							return err
						}

						err = uuc.uiRepo.UpdateUserRewardRecommendBrc(ctx, tmpUserId, float64(amount)*sendRecommendRate, user.Address)
						if err != nil {
							fmt.Println("错误分红直推：", err)
							return err
						}

						return nil
					}); nil != err {
						fmt.Println("err reward recommend", err, amount, amountRecommendTmp, user, tmpRecommendUser)
					}

					if stopRecommend {
						//// 推荐人
						//var (
						//	userRecommendArea *UserRecommend
						//)
						//if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
						//	userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
						//} else {
						//	fmt.Println("错误分红业绩变更，信息缺失7：", err, amount, amountRecommendTmp, user, tmpRecommendUser)
						//	continue
						//}
						//
						//if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
						//	var tmpRecommendAreaUserIds []string
						//	tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")
						//
						//	for j := len(tmpRecommendAreaUserIds) - 1; j >= 0; j-- {
						//		if 0 >= len(tmpRecommendAreaUserIds[j]) {
						//			continue
						//		}
						//
						//		myUserRecommendAreaUserId, _ := strconv.ParseInt(tmpRecommendAreaUserIds[j], 10, 64) // 最后一位是直推人
						//		if 0 >= myUserRecommendAreaUserId {
						//			continue
						//		}
						//		if err = ruc.tx.ExecTx(ctx, func(ctx context.Context) error { //
						//			// 减掉业绩
						//			err = ruc.userInfoRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, vUserRecords.Amount)
						//			if err != nil {
						//				fmt.Println("错误分红社区：", err, amount, amountRecommendTmp, user, vUserRecords)
						//			}
						//
						//			return nil
						//		}); nil != err {
						//			fmt.Println("err reward recommend 2", err, amount, amountRecommendTmp, user, vUserRecords)
						//		}
						//	}
						//}

						if 0.000001 < amountRecommendTmp {
							continue
						}
					}

					break
				}
			}
		}
	}

	return &v1.BuyReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) BuyTwo(ctx context.Context, req *v1.BuyRequest, user *User) (*v1.BuyReply, error) {
	// 推荐人
	var (
		err      error
		configs  []*Config
		goodRate float64
		priceBrc float64
	)

	// 配置
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"good_two_rate",
		"price_brc",
	)
	if nil != err || nil == configs {
		return &v1.BuyReply{
			Status: "稍后重试|err wait",
		}, nil
	}

	for _, vConfig := range configs {

		if "good_two_rate" == vConfig.KeyName {
			goodRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "price_brc" == vConfig.KeyName {
			priceBrc, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	var (
		goods    []*Good
		goodsMap map[int64]*Good
		goodsBuy *Good
	)
	goods, err = uuc.ubRepo.GetGoodsOnlineTwo(ctx)
	if nil != err {
		return &v1.BuyReply{
			Status: "稍后重试|err wait",
		}, nil
	}
	goodsMap = make(map[int64]*Good, 0)
	for _, v := range goods {
		goodsMap[v.ID] = v
	}

	if _, ok := goodsMap[req.SendBody.Id]; !ok {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}
	goodsBuy = goodsMap[req.SendBody.Id]
	amount := goodsBuy.Amount

	var (
		amountRel    = float64(amount)
		amountRelBrc float64
		userBalance  *UserBalance
	)

	amountRel = float64(amount) * goodRate
	amountRelBrc = (float64(amount) - float64(amount)*goodRate) / priceBrc

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil == userBalance || nil != err {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}

	if 1 < amountRelBrc && amountRelBrc > userBalance.BalanceRawFloatNew {
		return &v1.BuyReply{
			Status: "brc20余额不足|brc20 not enough",
		}, nil
	}

	if 1 < amountRel && amountRel > userBalance.BalanceUsdtFloat {
		return &v1.BuyReply{
			Status: "usdt余额不足|deposit usdt not enough",
		}, nil
	}

	one := ""
	two := ""
	three := ""
	if "1" != user.One {
		one += user.One
	}
	if "1" != user.Two {
		one += user.Two
	}
	if "1" != user.Three {
		one += user.Three
	}
	if "1" != user.Four {
		one += user.Four
	}
	if "1" != user.Five {
		one += user.Five
	}
	if "1" != user.Six {
		two = user.Six
	}
	if "1" != user.Seven {
		three = user.Seven
	}

	four := goodsBuy.ID
	// 入金
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.uiRepo.UpdateUserTwoIn(ctx, user.ID, amount, amountRel, amountRelBrc, one, two, three, four)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "错误投资 two", amount)
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}

	// 推荐人
	var (
		userRecommend       *UserRecommend
		tmpRecommendUserIds []string
	)
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil != err {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	totalTmp := len(tmpRecommendUserIds) - 1
	for i := totalTmp; i >= 0; i-- {
		tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		// 增加业绩
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.uiRepo.UpdateUserMyTotalAmountAdd(ctx, tmpUserId, float64(amount), 0)
			if err != nil {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("遍历业绩：", err, tmpUserId, user)
			continue
		}
	}

	return &v1.BuyReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) BuyThree(ctx context.Context, req *v1.BuyRequest, user *User) (*v1.BuyReply, error) {
	// 推荐人
	var (
		err      error
		configs  []*Config
		goodRate float64
		priceBrc float64
	)

	// 配置
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"good_two_rate",
		"price_brc",
	)
	if nil != err || nil == configs {
		return &v1.BuyReply{
			Status: "稍后重试|err wait",
		}, nil
	}

	for _, vConfig := range configs {

		if "good_two_rate" == vConfig.KeyName {
			goodRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "price_brc" == vConfig.KeyName {
			priceBrc, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	var (
		goods    []*Good
		goodsMap map[int64]*Good
		goodsBuy *Good
	)
	goods, err = uuc.ubRepo.GetGoodsOnlineThree(ctx)
	if nil != err {
		return &v1.BuyReply{
			Status: "稍后重试|err wait",
		}, nil
	}
	goodsMap = make(map[int64]*Good, 0)
	for _, v := range goods {
		goodsMap[v.ID] = v
	}

	if _, ok := goodsMap[req.SendBody.Id]; !ok {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}
	goodsBuy = goodsMap[req.SendBody.Id]
	amount := goodsBuy.Amount

	var (
		amountRel    = float64(amount)
		amountRelBrc float64
		userBalance  *UserBalance
	)

	amountRel = float64(amount) * goodRate
	amountRelBrc = (float64(amount) - float64(amount)*goodRate) / priceBrc

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil == userBalance || nil != err {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}

	if 1 < amountRelBrc && amountRelBrc > userBalance.BalanceRawFloatNew {
		return &v1.BuyReply{
			Status: "brc20余额不足|brc20 not enough",
		}, nil
	}

	if 1 < amountRel && amountRel > userBalance.BalanceUsdtFloat {
		return &v1.BuyReply{
			Status: "usdt余额不足|deposit usdt not enough",
		}, nil
	}

	one := ""
	two := ""
	three := ""
	if "1" != user.One {
		one += user.One
	}
	if "1" != user.Two {
		one += user.Two
	}
	if "1" != user.Three {
		one += user.Three
	}
	if "1" != user.Four {
		one += user.Four
	}
	if "1" != user.Five {
		one += user.Five
	}
	if "1" != user.Six {
		two = user.Six
	}
	if "1" != user.Seven {
		three = user.Seven
	}

	four := goodsBuy.ID
	// 入金
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.uiRepo.UpdateUserThreeIn(ctx, user.ID, amount, amountRel, amountRelBrc, one, two, three, four)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "错误投资 two", amount)
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}

	// 推荐人
	var (
		userRecommend       *UserRecommend
		tmpRecommendUserIds []string
	)
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil != err {
		return &v1.BuyReply{
			Status: "参数错误 |err id",
		}, nil
	}
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	totalTmp := len(tmpRecommendUserIds) - 1
	for i := totalTmp; i >= 0; i-- {
		tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		// 增加业绩
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.uiRepo.UpdateUserMyTotalAmountAdd(ctx, tmpUserId, float64(amount), 0)
			if err != nil {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("遍历业绩：", err, tmpUserId, user)
			continue
		}
	}

	return &v1.BuyReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) EthUserRecordHandle(ctx context.Context, amount uint64, ethUserRecord ...*EthUserRecord) (bool, error, string) {

	var (
		err          error
		configs      []*Config
		recommendOne float64
		recommendTwo float64
		//b1           float64
	)

	// 配置
	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "recommend_one", "recommend_two")
	if nil != configs {
		for _, vConfig := range configs {
			if "recommend_one" == vConfig.KeyName {
				recommendOne, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "recommend_two" == vConfig.KeyName {
				recommendTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
		}
	}

	var (
		userRecommends    []*UserRecommend
		userRecommendsMap map[int64]*UserRecommend
	)
	userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	if nil != err {
		fmt.Println("今认购用户获取失败2")
		return false, err, "错误"
	}

	myLowUser := make(map[int64][]*UserRecommend, 0)
	userRecommendsMap = make(map[int64]*UserRecommend, 0)
	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	var (
		users    []*User
		usersMap map[int64]*User
	)
	users, err = uuc.repo.GetAllUsers(ctx)
	if nil == users {
		fmt.Println("认购用户获取失败")
		return false, nil, "错误"
	}

	usersMap = make(map[int64]*User, 0)

	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	stopUserIds := make(map[int64]bool, 0)
	for _, v := range ethUserRecord {
		if _, ok := usersMap[v.UserId]; !ok {
			fmt.Println("认购用户不存在", v)
			continue
		}

		last := uint64(0)
		newAmountUsdt := usersMap[v.UserId].AmountUsdt + float64(amount)
		if 100 <= newAmountUsdt && newAmountUsdt < 599 {
			last = 1
		} else if 600 <= newAmountUsdt && newAmountUsdt < 2999 {
			last = 2
		} else if 3000 <= newAmountUsdt && newAmountUsdt < 9999 {
			last = 3
		} else if 10000 <= newAmountUsdt && newAmountUsdt < 29999 {
			last = 4
		} else if 30000 <= newAmountUsdt {
			last = 5
		} else {
			fmt.Println("认购金额不足", v)
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.repo.UpdateUserNewTwoNew(ctx, v.UserId, amount, float64(amount), last)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "错误投资3", v)
			return false, err, "错误"
		}

		// 推荐人
		var (
			userRecommend       *UserRecommend
			tmpRecommendUserIds []string
		)
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, v.UserId)
		if nil != err {
			continue
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		}

		//tmpMapCurrentI := make(map[int]float64, 0)
		//tmpB1 := b1
		//for ii := 1; ii <= 18; ii++ {
		//	if ii > 1 {
		//		tmpMapCurrentI[ii] = tmpB1 / 2
		//	} else {
		//		tmpMapCurrentI[ii] = tmpB1
		//	}
		//	tmpB1 = tmpMapCurrentI[ii]
		//}

		tmpRecommendTotal := 1
		totalTmp := len(tmpRecommendUserIds) - 1
		//currentI := 1
		for i := totalTmp; i >= 0; i-- {
			//tmpCurrentI := currentI
			//currentI++

			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("buy遍历，信息缺失,user：", err, v, tmpUserId)
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if i == totalTmp {
				// 增加业绩
				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					err = uuc.repo.UpdateUserMyRecommendTotal(ctx, tmpUserId, float64(amount))
					if err != nil {
						return err
					}

					return nil
				}); nil != err {
					fmt.Println("遍历业绩：", err, v)
					continue
				}
			}

			// 增加业绩
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = uuc.repo.UpdateUserMyTotalAmount(ctx, tmpUserId, float64(amount))
				if err != nil {
					return err
				}

				return nil
			}); nil != err {
				fmt.Println("遍历业绩：", err, v)
				continue
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[tmpUserId]; ok {
				continue
			}

			// 入金
			if 0 < tmpRecommendUser.AmountUsdt {
				if 2 >= tmpRecommendTotal {

					var (
						num float64
					)

					if 1 == tmpRecommendUser.Last {
						num = 1.5
					} else if 2 == tmpRecommendUser.Last {
						num = 2
					} else if 3 == tmpRecommendUser.Last {
						num = 2.5
					} else if 4 == tmpRecommendUser.Last {
						num = 3
					} else if 5 == tmpRecommendUser.Last {
						num = 3.5
					} else {
						continue
					}

					amountRecommendTmp := float64(amount) * recommendOne
					if 2 == tmpRecommendTotal {
						amountRecommendTmp = float64(amount) * recommendTwo
					}

					var (
						stopRecommend bool
					)
					if amountRecommendTmp+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
						amountRecommendTmp = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
						stopRecommend = true
					}
					amountRecommendTmp = math.Round(amountRecommendTmp*10000000) / 10000000

					fmt.Println("直推奖奖励：", float64(amount), amountRecommendTmp, v, tmpRecommendUser)
					if amountRecommendTmp > 0 {
						if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
							var (
								code int64
							)
							code, err = uuc.repo.UpdateUserRewardRecommend2(ctx, tmpUserId, amountRecommendTmp, tmpRecommendUser.AmountUsdt, stopRecommend, int64(tmpRecommendTotal), int64(i), usersMap[v.UserId].Address)
							if code > 0 && err != nil {
								fmt.Println("错误分红特殊：", err)
							}

							return nil
						}); nil != err {
							fmt.Println("err reward area 2", err, v)
						}

						if stopRecommend {
							stopUserIds[tmpRecommendUser.ID] = true // 出局

							// 推荐人
							var (
								userRecommendArea *UserRecommend
							)
							if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
								userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
							} else {
								fmt.Println("错误分红业绩变更，信息缺失7：", err, v)
								continue
							}

							if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
								var tmpRecommendAreaUserIds []string
								tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

								for j := len(tmpRecommendAreaUserIds) - 1; j >= 0; j-- {
									if 0 >= len(tmpRecommendAreaUserIds[j]) {
										continue
									}

									myUserRecommendAreaUserId, _ := strconv.ParseInt(tmpRecommendAreaUserIds[j], 10, 64) // 最后一位是直推人
									if 0 >= myUserRecommendAreaUserId {
										continue
									}
									if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { //
										// 减掉业绩
										err = uuc.repo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
										if err != nil {
											fmt.Println("错误分红社区：", err, v)
										}

										return nil
									}); nil != err {
										fmt.Println("err reward area 2", err, v)
									}

									//// 级别降低
									//// 我的下级，更新vip
									//userIdsLowTmpTwo := make([]int64, 0)
									//for _, vTmpLow := range myLowUser[myUserRecommendAreaUserId] {
									//	userIdsLowTmpTwo = append(userIdsLowTmpTwo, vTmpLow.UserId)
									//}
									//if 0 < len(userIdsLowTmpTwo) {
									//	uuc.updateVip(ctx, tmpUserId, userIdsLowTmpTwo)
									//}
								}
							}
						}

					}

					tmpRecommendTotal++
				}

				//if currentI <= 18 {
				//	// 我的下级
				//	if _, ok := myLowUser[tmpUserId]; !ok {
				//		fmt.Println("错误分红帮扶，信息缺失3：", err, tmpUserId, v)
				//		continue
				//	}
				//
				//	if 0 >= len(myLowUser[tmpUserId]) {
				//		fmt.Println("错误分红帮扶，信息缺失3：", err, tmpUserId, v)
				//		continue
				//	}
				//
				//	// 条件
				//	if tmpCurrentI < 10 {
				//		if tmpCurrentI > len(myLowUser[tmpUserId]) {
				//			continue
				//		}
				//	} else if 10 > len(myLowUser[tmpUserId]) {
				//		continue
				//	}
				//
				//	if _, ok := tmpMapCurrentI[tmpCurrentI]; !ok {
				//		fmt.Println("错误分红帮扶，信息缺失35：", err, tmpUserId, v)
				//		continue
				//	}
				//	tmpRecommendAmount := float64(amount) * tmpMapCurrentI[tmpCurrentI]
				//
				//	var (
				//		stopRecommend   bool
				//		numRecommendTwo float64
				//	)
				//	if 1 == tmpRecommendUser.Last {
				//		numRecommendTwo = 1.5
				//	} else if 2 == tmpRecommendUser.Last {
				//		numRecommendTwo = 2
				//	} else if 3 == tmpRecommendUser.Last {
				//		numRecommendTwo = 2.5
				//	} else if 4 == tmpRecommendUser.Last {
				//		numRecommendTwo = 3
				//	} else if 5 == tmpRecommendUser.Last {
				//		numRecommendTwo = 3.5
				//	} else {
				//		continue
				//	}
				//
				//	if tmpRecommendAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*numRecommendTwo {
				//		tmpRecommendAmount = math.Abs(tmpRecommendUser.AmountUsdt*numRecommendTwo - tmpRecommendUser.AmountUsdtGet)
				//		stopRecommend = true
				//	}
				//
				//	// 分红
				//	tmpRecommendAmount = math.Round(tmpRecommendAmount*10000000) / 10000000
				//	if 0 >= tmpRecommendAmount {
				//		continue
				//	}
				//
				//	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				//		var (
				//			code int64
				//		)
				//
				//		code, err = uuc.repo.UpdateUserRewardRecommendNew(ctx, tmpRecommendUser.ID, tmpRecommendAmount, tmpRecommendUser.AmountUsdt, stopRecommend, int64(tmpCurrentI), usersMap[v.UserId].Address)
				//		if code > 0 && err != nil {
				//			fmt.Println("错误分红帮扶：", err, tmpRecommendUser)
				//		}
				//
				//		return nil
				//	}); nil != err {
				//		fmt.Println("err reward daily recommend 18", err, tmpRecommendUser)
				//	}
				//
				//	if stopRecommend {
				//		stopUserIds[tmpRecommendUser.ID] = true // 出局
				//
				//		// 推荐人
				//		var (
				//			userRecommendArea *UserRecommend
				//		)
				//		if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
				//			userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
				//		} else {
				//			fmt.Println("错误分红帮扶，信息缺失7：", err, tmpRecommendUser)
				//		}
				//
				//		if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
				//			var tmpRecommendAreaUserIds []string
				//			tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")
				//
				//			for j := len(tmpRecommendAreaUserIds) - 1; j >= 0; j-- {
				//				if 0 >= len(tmpRecommendAreaUserIds[j]) {
				//					continue
				//				}
				//
				//				myUserRecommendAreaUserId, _ := strconv.ParseInt(tmpRecommendAreaUserIds[j], 10, 64) // 最后一位是直推人
				//				if 0 >= myUserRecommendAreaUserId {
				//					continue
				//				}
				//
				//				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error {
				//					// 减掉业绩
				//					err = uuc.repo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
				//					if err != nil {
				//						fmt.Println("错误分红帮扶：", err, tmpRecommendUser)
				//					}
				//
				//					return nil
				//				}); nil != err {
				//					fmt.Println("err reward daily 业绩更新", err, v)
				//					continue
				//				}
				//			}
				//		}
				//	}
				//}
			}

			// 我的下级，更新vip
			userIdsLowTmp := make([]int64, 0)
			for _, vTmpLow := range myLowUser[tmpUserId] {
				userIdsLowTmp = append(userIdsLowTmp, vTmpLow.UserId)
			}
			if 0 < len(userIdsLowTmp) {
				uuc.updateVip(ctx, tmpUserId, userIdsLowTmp)
			}
		}
	}

	return true, nil, ""
}

func (uuc *UserUseCase) updateVip(ctx context.Context, tmpUserId int64, userIdsLowTmp []int64) {
	var (
		err error
	)

	// 下级的级别
	var (
		tmpLowUsers map[int64]*User
	)
	tmpLowUsers, err = uuc.repo.GetUserByUserIdsTwo(ctx, userIdsLowTmp)
	if err != nil {
		fmt.Println("update vip 遍历业绩2：", err)
		return
	}

	vip9 := 0
	vip8 := 0
	vip7 := 0
	vip6 := 0
	vip5 := 0
	vip4 := 0
	vip3 := 0
	vip2 := 0
	vip1 := 0
	// 获取业绩
	tmpAreaMax := float64(0)
	tmpAreaMin := float64(0)
	tmpMaxId := int64(0)
	for _, vMyLowUser := range tmpLowUsers {
		if 9 == vMyLowUser.Vip {
			vip9++
		} else if 8 == vMyLowUser.Vip {
			vip8++
		} else if 7 == vMyLowUser.Vip {
			vip7++
		} else if 6 == vMyLowUser.Vip {
			vip6++
		} else if 5 == vMyLowUser.Vip {
			vip5++
		} else if 4 == vMyLowUser.Vip {
			vip4++
		} else if 3 == vMyLowUser.Vip {
			vip3++
		} else if 2 == vMyLowUser.Vip {
			vip2++
		} else if 1 == vMyLowUser.Vip {
			vip1++
		}

		if tmpAreaMax < vMyLowUser.MyTotalAmount+vMyLowUser.AmountUsdt {
			tmpAreaMax = vMyLowUser.MyTotalAmount + vMyLowUser.AmountUsdt
			tmpMaxId = vMyLowUser.ID
		}
	}

	if 0 < tmpMaxId {
		for _, vMyLowUser := range tmpLowUsers {
			if tmpMaxId != vMyLowUser.ID {
				tmpAreaMin += vMyLowUser.MyTotalAmount + vMyLowUser.AmountUsdt
			}
		}
	}

	tmpVip := int64(0)
	if 10000000 <= tmpAreaMin && 2 <= vip9 {
		tmpVip = 10
	} else if 6000000 <= tmpAreaMin && 2 <= vip8+vip9 {
		tmpVip = 9
	} else if 3000000 <= tmpAreaMin && 2 <= vip7+vip8+vip9 {
		tmpVip = 8
	} else if 1300000 <= tmpAreaMin && 2 <= vip6+vip7+vip8+vip9 {
		tmpVip = 7
	} else if 600000 <= tmpAreaMin && 2 <= vip5+vip6+vip7+vip8+vip9 {
		tmpVip = 6
	} else if 300000 <= tmpAreaMin && 2 <= vip4+vip5+vip6+vip7+vip8+vip9 {
		tmpVip = 5
	} else if 100000 <= tmpAreaMin && 2 <= vip3+vip4+vip5+vip6+vip7+vip8+vip9 {
		tmpVip = 4
	} else if 50000 <= tmpAreaMin && 2 <= vip2+vip3+vip4+vip5+vip6+vip7+vip8+vip9 {
		tmpVip = 3
	} else if 20000 <= tmpAreaMin && 2 <= vip1+vip2+vip3+vip4+vip5+vip6+vip7+vip8+vip9 {
		tmpVip = 2
	} else if 5000 <= tmpAreaMin {
		tmpVip = 1
	}

	// 增加业绩
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.repo.UpdateUserVip(ctx, tmpUserId, tmpVip)
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println("update vip 遍历业绩 vip：", err)
		return
	}

	return
}

// Exchange Exchange.
func (uuc *UserUseCase) Exchange(ctx context.Context, req *v1.ExchangeRequest, user *User) (*v1.ExchangeReply, error) {
	var (
		//u           *User
		err         error
		userBalance *UserBalance
	)

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return &v1.ExchangeReply{
			Status: "错误",
		}, nil
	}

	amountFloat := float64(req.SendBody.Amount)

	if userBalance.BalanceUsdtFloat < amountFloat {
		return &v1.ExchangeReply{
			Status: "余额不足",
		}, nil
	}

	if 1 > amountFloat {
		return &v1.ExchangeReply{
			Status: "错误金额",
		}, nil
	}

	// 配置
	var (
		configs      []*Config
		exchangeRate float64
		bPrice       float64
	)
	configs, err = uuc.configRepo.GetConfigByKeys(ctx, "exchange_rate", "b_price")
	if nil != configs {
		for _, vConfig := range configs {
			if "exchange_rate" == vConfig.KeyName {
				exchangeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "b_price" == vConfig.KeyName {
				bPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
			}

		}
	}

	var (
		amountRaw       float64
		amountRawSubFee float64
	)

	amountRaw = amountFloat / bPrice
	fee := amountRaw * exchangeRate
	amountRawSubFee = amountRaw - fee
	if amountRawSubFee <= 0 {
		return &v1.ExchangeReply{
			Status: "fail price",
		}, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		err = uuc.ubRepo.Exchange(ctx, user.ID, amountFloat, fee, amountRawSubFee) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return &v1.ExchangeReply{
			Status: "错误",
		}, nil
	}

	return &v1.ExchangeReply{
		Status: "ok",
	}, nil

}

func (uuc *UserUseCase) SetInfo(ctx context.Context, req *v1.SetInfoRequest, user *User) (*v1.SetInfoReply, error) {
	var (
		err error
	)

	if 1 > len(req.SendBody.One) || 100 < len(req.SendBody.One) {
		return &v1.SetInfoReply{
			Status: "国家字符超限",
		}, nil
	}
	if 100 < len(req.SendBody.Two) {
		return &v1.SetInfoReply{
			Status: "省份字符超限",
		}, nil
	}
	if 100 < len(req.SendBody.Three) {
		return &v1.SetInfoReply{
			Status: "城市字符超限",
		}, nil
	}
	if 100 < len(req.SendBody.Four) {
		return &v1.SetInfoReply{
			Status: "地区字符超限",
		}, nil
	}
	if 1 > len(req.SendBody.Five) || 100 < len(req.SendBody.Five) {
		return &v1.SetInfoReply{
			Status: "详细字符数不符合标准",
		}, nil
	}
	if 1 > len(req.SendBody.Six) || 100 < len(req.SendBody.Six) {
		return &v1.SetInfoReply{
			Status: "收件人手机字符数不符合标准",
		}, nil
	}
	if 1 > len(req.SendBody.Seven) || 100 < len(req.SendBody.Seven) {
		return &v1.SetInfoReply{
			Status: "收件人字符数不符合标准",
		}, nil
	}

	err = uuc.repo.UpdateUserAddressInfo(ctx, user.ID, req.SendBody.One, "1", "1", "1", req.SendBody.Five, req.SendBody.Six, req.SendBody.Seven)
	if nil != err {
		return &v1.SetInfoReply{
			Status: "修改失败，稍后重试",
		}, nil
	}

	return &v1.SetInfoReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Withdraw(ctx context.Context, req *v1.WithdrawRequest, user *User) (*v1.WithdrawReply, error) {
	var (
		err         error
		userBalance *UserBalance
	)

	// 配置
	var (
		configs         []*Config
		withdrawMin     float64
		withdrawRate    float64
		withdrawMinTwo  float64
		withdrawRateTwo float64
		openWithdraw    uint64
	)
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"withdraw_rate",
		"withdraw_amount_min",
		"withdraw_rate_two",
		"withdraw_amount_min_two",
		"open_withdraw",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_amount_min" == vConfig.KeyName {
				withdrawMin, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "withdraw_rate" == vConfig.KeyName {
				withdrawRate, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "withdraw_amount_min_two" == vConfig.KeyName {
				withdrawMinTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "withdraw_rate_two" == vConfig.KeyName {
				withdrawRateTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "open_withdraw" == vConfig.KeyName {
				openWithdraw, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
		}
	}

	if 1 != openWithdraw {
		return &v1.WithdrawReply{
			Status: "暂未开放",
		}, nil
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return &v1.WithdrawReply{
			Status: "错误",
		}, nil
	}

	if 2 == req.SendBody.CoinType {
		return &v1.WithdrawReply{
			Status: "暂未开放",
		}, nil

		amountFloat := float64(req.SendBody.Amount)
		if userBalance.BalanceRawFloatNew < amountFloat {
			return &v1.WithdrawReply{
				Status: "可提余额不足",
			}, nil
		}

		if 1 > amountFloat {
			return &v1.WithdrawReply{
				Status: "错误金额",
			}, nil
		}

		if withdrawMinTwo > amountFloat {
			return &v1.WithdrawReply{
				Status: "fail min",
			}, nil
		}

		amountFloatSubFee := amountFloat - amountFloat*withdrawRateTwo
		if 0 >= amountFloat {
			return &v1.WithdrawReply{
				Status: "fail price",
			}, nil
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.ubRepo.WithdrawISPAY(ctx, user.ID, amountFloat) // 提现
			if nil != err {
				return err
			}
			_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, amountFloatSubFee, amountFloat, "RAW", user.Address)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &v1.WithdrawReply{
				Status: "错误",
			}, nil
		}
	} else {
		amountFloat := float64(req.SendBody.Amount)
		if userBalance.BalanceUsdtFloat < amountFloat {
			return &v1.WithdrawReply{
				Status: "可提余额不足",
			}, nil
		}

		if 1 > amountFloat {
			return &v1.WithdrawReply{
				Status: "错误金额",
			}, nil
		}

		if withdrawMin > amountFloat {
			return &v1.WithdrawReply{
				Status: "fail min",
			}, nil
		}

		amountFloatSubFee := amountFloat - amountFloat*withdrawRate
		if 0 >= amountFloat {
			return &v1.WithdrawReply{
				Status: "fail price",
			}, nil
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.ubRepo.WithdrawUsdt2(ctx, user.ID, amountFloat) // 提现
			if nil != err {
				return err
			}
			_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, amountFloatSubFee, amountFloat, "USDT", user.Address)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &v1.WithdrawReply{
				Status: "错误",
			}, nil
		}
	}

	return &v1.WithdrawReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Tran(ctx context.Context, req *v1.TranRequest, user *User, password string) (*v1.TranReply, error) {
	var (
		err         error
		userBalance *UserBalance
		toUser      *User
		u           *User
	)

	u, _ = uuc.repo.GetUserById(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if "" == u.Password || 6 > len(u.Password) {
		return nil, errors.New(500, "ERROR_TOKEN", "未设置密码，联系管理员")
	}

	if u.Password != user.Password {
		return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	}

	if password != u.Password {
		return nil, errors.New(500, "密码错误", "密码错误")
	}

	if "" == req.SendBody.Address {
		return &v1.TranReply{
			Status: "不存在地址",
		}, nil
	}

	toUser, err = uuc.repo.GetUserByAddress(ctx, req.SendBody.Address)
	if nil != err {
		return &v1.TranReply{
			Status: "不存在地址",
		}, nil
	}

	if user.ID == toUser.ID {
		return &v1.TranReply{
			Status: "不能给自己转账",
		}, nil
	}

	if "dhb" != req.SendBody.Type && "usdt" != req.SendBody.Type {
		return &v1.TranReply{
			Status: "fail",
		}, nil
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)

	if "dhb" == req.SendBody.Type {
		if userBalance.BalanceDhb < amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}

		if 10000000 > amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}
	}

	if "usdt" == req.SendBody.Type {
		if userBalance.BalanceUsdt < amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}

		if 1000000 > amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}
	}

	var (
		userRecommend  *UserRecommend
		userRecommend2 *UserRecommend
	)
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend {
		return &v1.TranReply{
			Status: "信息错误",
		}, nil
	}

	var (
		tmpRecommendUserIds          []string
		tmpRecommendUserIdsInt       []int64
		toUserTmpRecommendUserIds    []string
		toUserTmpRecommendUserIdsInt []int64
	)
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	if 1 < len(tmpRecommendUserIds) {
		lastKey := len(tmpRecommendUserIds) - 1
		if 1 <= lastKey {
			for i := 0; i <= lastKey; i++ {
				// 有占位信息，推荐人推荐人的上一代
				if lastKey-i <= 0 {
					break
				}

				tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
				tmpRecommendUserIdsInt = append(tmpRecommendUserIdsInt, tmpMyTopUserRecommendUserId)
			}
		}
	}

	userRecommend2, err = uuc.urRepo.GetUserRecommendByUserId(ctx, toUser.ID)
	if nil == userRecommend2 {
		return &v1.TranReply{
			Status: "信息错误",
		}, nil
	}
	if "" != userRecommend2.RecommendCode {
		toUserTmpRecommendUserIds = strings.Split(userRecommend2.RecommendCode, "D")
	}

	if 1 < len(toUserTmpRecommendUserIds) {
		lastKey2 := len(toUserTmpRecommendUserIds) - 1
		if 1 <= lastKey2 {
			for i := 0; i <= lastKey2; i++ {
				// 有占位信息，推荐人推荐人的上一代
				if lastKey2-i <= 0 {
					break
				}

				toUserTmpMyTopUserRecommendUserId, _ := strconv.ParseInt(toUserTmpRecommendUserIds[lastKey2-i], 10, 64) // 最后一位是直推人
				toUserTmpRecommendUserIdsInt = append(toUserTmpRecommendUserIdsInt, toUserTmpMyTopUserRecommendUserId)
			}
		}
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		if "usdt" == req.SendBody.Type {
			err = uuc.ubRepo.TranUsdt(ctx, user.ID, toUser.ID, amount, tmpRecommendUserIdsInt, toUserTmpRecommendUserIdsInt) // 提现
			if nil != err {
				return err
			}
		} else if "dhb" == req.SendBody.Type {
			err = uuc.ubRepo.TranDhb(ctx, user.ID, toUser.ID, amount) // 提现
			if nil != err {
				return err
			}
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.TranReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Trade(ctx context.Context, req *v1.WithdrawRequest, user *User, amount int64, amountB int64, amount2 int64, password string) (*v1.WithdrawReply, error) {
	var (
		u                   *User
		userBalance         *UserBalance
		userBalance2        *UserBalance
		configs             []*Config
		userRecommend       *UserRecommend
		withdrawRate        int64
		withdrawDestroyRate int64
		err                 error
	)

	u, _ = uuc.repo.GetUserById(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if "" == u.Password || 6 > len(u.Password) {
		return nil, errors.New(500, "ERROR_TOKEN", "未设置密码，联系管理员")
	}

	if u.Password != user.Password {
		return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	}

	if password != u.Password {
		return nil, errors.New(500, "密码错误", "密码错误")
	}

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "withdraw_rate",
		"withdraw_destroy_rate",
	)

	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_rate" == vConfig.KeyName {
				withdrawRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_destroy_rate" == vConfig.KeyName {
				withdrawDestroyRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	userBalance, err = uuc.ubRepo.GetUserBalanceLock(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	userBalance2, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if userBalance.BalanceUsdt < amount {
		return &v1.WithdrawReply{
			Status: "csd锁定部分的余额不足",
		}, nil
	}

	if userBalance2.BalanceDhb < amountB {
		return &v1.WithdrawReply{
			Status: "hbs锁定部分的余额不足",
		}, nil
	}

	// 推荐人
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend {
		return &v1.WithdrawReply{
			Status: "信息错误",
		}, nil
	}

	var (
		tmpRecommendUserIds    []string
		tmpRecommendUserIdsInt []int64
	)
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}
	lastKey := len(tmpRecommendUserIds) - 1
	if 1 <= lastKey {
		for i := 0; i <= lastKey; i++ {
			// 有占位信息，推荐人推荐人的上一代
			if lastKey-i <= 0 {
				break
			}

			tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
			tmpRecommendUserIdsInt = append(tmpRecommendUserIdsInt, tmpMyTopUserRecommendUserId)
		}
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		err = uuc.ubRepo.Trade(ctx, user.ID, amount, amountB, amount-amount/100*(withdrawRate+withdrawDestroyRate), amountB-amountB/100*(withdrawRate+withdrawDestroyRate), tmpRecommendUserIdsInt, amount2) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.WithdrawReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) SetBalanceReward(ctx context.Context, req *v1.SetBalanceRewardRequest, user *User) (*v1.SetBalanceRewardReply, error) {
	var (
		err         error
		userBalance *UserBalance
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)
	if 0 >= amount {
		return &v1.SetBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if userBalance.BalanceUsdt < amount {
		return &v1.SetBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		err = uuc.ubRepo.SetBalanceReward(ctx, user.ID, amount) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.SetBalanceRewardReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) DeleteBalanceReward(ctx context.Context, req *v1.DeleteBalanceRewardRequest, user *User) (*v1.DeleteBalanceRewardReply, error) {
	var (
		err            error
		balanceRewards []*BalanceReward
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)
	if 0 >= amount {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	balanceRewards, err = uuc.ubRepo.GetBalanceRewardByUserId(ctx, user.ID)
	if nil != err {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	var totalBalanceRewardAmount int64
	for _, vBalanceReward := range balanceRewards {
		totalBalanceRewardAmount += vBalanceReward.Amount
	}

	if totalBalanceRewardAmount < amount {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	for _, vBalanceReward := range balanceRewards {
		tmpAmount := int64(0)
		Status := int64(1)

		if amount-vBalanceReward.Amount < 0 {
			tmpAmount = amount
		} else {
			tmpAmount = vBalanceReward.Amount
			Status = 2
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.ubRepo.UpdateBalanceReward(ctx, user.ID, vBalanceReward.ID, tmpAmount, Status) // 提现
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return nil, err
		}
		amount -= tmpAmount

		if amount <= 0 {
			break
		}
	}

	return &v1.DeleteBalanceRewardReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) AdminRewardList(ctx context.Context, req *v1.AdminRewardListRequest) (*v1.AdminRewardListReply, error) {
	res := &v1.AdminRewardListReply{
		Rewards: make([]*v1.AdminRewardListReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListReply, error) {

	res := &v1.AdminUserListReply{
		Users: make([]*v1.AdminUserListReply_UserList, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error) {
	return uuc.repo.GetUserByUserIds(ctx, userIds...)
}

func (uuc *UserUseCase) GetUserByUserId(ctx context.Context, userId int64) (*User, error) {
	return uuc.repo.GetUserById(ctx, userId)
}

func (uuc *UserUseCase) AdminLocationList(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	res := &v1.AdminLocationListReply{
		Locations: make([]*v1.AdminLocationListReply_LocationList, 0),
	}
	return res, nil

}

func (uuc *UserUseCase) AdminRecommendList(ctx context.Context, req *v1.AdminUserRecommendRequest) (*v1.AdminUserRecommendReply, error) {
	res := &v1.AdminUserRecommendReply{
		Users: make([]*v1.AdminUserRecommendReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) AdminMonthRecommend(ctx context.Context, req *v1.AdminMonthRecommendRequest) (*v1.AdminMonthRecommendReply, error) {

	res := &v1.AdminMonthRecommendReply{
		Users: make([]*v1.AdminMonthRecommendReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfig(ctx context.Context, req *v1.AdminConfigRequest) (*v1.AdminConfigReply, error) {
	res := &v1.AdminConfigReply{
		Config: make([]*v1.AdminConfigReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) AdminConfigUpdate(ctx context.Context, req *v1.AdminConfigUpdateRequest) (*v1.AdminConfigUpdateReply, error) {
	res := &v1.AdminConfigUpdateReply{}
	return res, nil
}

func (uuc *UserUseCase) GetWithdrawPassOrRewardedList(ctx context.Context) ([]*Withdraw, error) {
	return uuc.ubRepo.GetWithdrawPassOrRewarded(ctx)
}

func (uuc *UserUseCase) UpdateWithdrawDoing(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "doing")
}

func (uuc *UserUseCase) UpdateWithdrawSuccess(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "success")
}

func (uuc *UserUseCase) AdminWithdrawList(ctx context.Context, req *v1.AdminWithdrawListRequest) (*v1.AdminWithdrawListReply, error) {
	res := &v1.AdminWithdrawListReply{
		Withdraw: make([]*v1.AdminWithdrawListReply_List, 0),
	}

	return res, nil

}

func (uuc *UserUseCase) AdminFee(ctx context.Context, req *v1.AdminFeeRequest) (*v1.AdminFeeReply, error) {
	return &v1.AdminFeeReply{}, nil
}

func (uuc *UserUseCase) AdminAll(ctx context.Context, req *v1.AdminAllRequest) (*v1.AdminAllReply, error) {

	return &v1.AdminAllReply{}, nil
}

func (uuc *UserUseCase) AdminWithdraw(ctx context.Context, req *v1.AdminWithdrawRequest) (*v1.AdminWithdrawReply, error) {
	return &v1.AdminWithdrawReply{}, nil
}
