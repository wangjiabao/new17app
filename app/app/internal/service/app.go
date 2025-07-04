package service

import (
	"context"
	"crypto/ecdsa"
	"crypto/md5"
	v1 "dhb/app/app/api"
	"dhb/app/app/internal/biz"
	"dhb/app/app/internal/conf"
	"dhb/app/app/internal/pkg/middleware/auth"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// AppService service.
type AppService struct {
	v1.UnimplementedAppServer

	uuc *biz.UserUseCase
	ruc *biz.RecordUseCase
	log *log.Helper
	ca  *conf.Auth
}

// NewAppService new a service.
func NewAppService(uuc *biz.UserUseCase, ruc *biz.RecordUseCase, logger log.Logger, ca *conf.Auth) *AppService {
	return &AppService{uuc: uuc, ruc: ruc, log: log.NewHelper(logger), ca: ca}
}

var lockEthAuthorize sync.Mutex

// EthAuthorize ethAuthorize.
func (a *AppService) EthAuthorize(ctx context.Context, req *v1.EthAuthorizeRequest) (*v1.EthAuthorizeReply, error) {
	lockEthAuthorize.Lock()
	defer lockEthAuthorize.Unlock()
	userAddress := req.SendBody.Address // 以太坊账户
	if "" == userAddress || 20 > len(userAddress) ||
		strings.EqualFold("0x000000000000000000000000000000000000dead", userAddress) {
		return &v1.EthAuthorizeReply{
			Token:  "",
			Status: "账户地址参数错误",
		}, nil
	}

	//// 验证
	//var (
	//	res     bool
	//	address string
	//	err     error
	//)
	//
	//res, address, err = verifySig2(req.SendBody.Sign, req.SendBody.PublicKey, "login")
	//if !res || nil != err || 0 >= len(address) || userAddress != address {
	//	return &v1.EthAuthorizeReply{
	//		Token:  "",
	//		Status: "地址签名错误",
	//	}, nil
	//}

	// 验证
	var (
		res bool
		err error
	)
	res, err = addressCheck(userAddress)
	if nil != err {
		return &v1.EthAuthorizeReply{
			Token:  "",
			Status: "地址验证失败",
		}, nil
	}
	if !res {
		return &v1.EthAuthorizeReply{
			Token:  "",
			Status: "地址格式错误",
		}, nil
	}

	if 10 >= len(req.SendBody.Sign) {
		return &v1.EthAuthorizeReply{
			Token:  "",
			Status: "签名错误",
		}, nil
	}

	var (
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(userAddress))
	if !res || addressFromSign != userAddress {
		return &v1.EthAuthorizeReply{
			Token:  "",
			Status: "地址签名错误",
		}, nil
	}

	//if "" == req.SendBody.Password || 6 > len(req.SendBody.Password) {
	//	return nil, errors.New(500, "AUTHORIZE_ERROR", "账户密码必须大于6位")
	//}

	// TODO 验证签名
	//password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	password := ""
	// 根据地址查询用户，不存在时则创建
	user, err, msg := a.uuc.GetExistUserByAddressOrCreate(ctx, &biz.User{
		Address:  userAddress,
		Password: password,
	}, req)
	if err != nil {
		return &v1.EthAuthorizeReply{
			Token:  "",
			Status: msg,
		}, nil
	}

	if 1 == user.IsDelete {
		return &v1.EthAuthorizeReply{
			Token:  "",
			Status: "用户已删除",
		}, nil
	}

	//if 1 == user.Lock {
	//	return &v1.EthAuthorizeReply{
	//		Token:  "",
	//		Status: "用户已锁定",
	//	}, nil
	//}

	claims := auth.CustomClaims{
		UserId:   user.ID,
		UserType: "user",
		Password: password,
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),                      // 签名的生效时间
			ExpiresAt: jwt2.NewNumericDate(time.Now().Add(100 * time.Hour)), // 2天过期
			Issuer:    "user",
		},
	}
	token, err := auth.CreateToken(claims, a.ca.JwtKey)
	if err != nil {
		return &v1.EthAuthorizeReply{
			Token:  token,
			Status: "生成token失败",
		}, nil
	}

	userInfoRsp := v1.EthAuthorizeReply{
		Token:  token,
		Status: "ok",
	}
	return &userInfoRsp, nil
}

// Deposit deposit.
func (a *AppService) Deposit(ctx context.Context, req *v1.DepositRequest) (*v1.DepositReply, error) {
	return &v1.DepositReply{}, nil
}

// UserArea UserArea.
func (a *AppService) UserArea(ctx context.Context, req *v1.UserAreaRequest) (*v1.UserAreaReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.UserArea(ctx, req, &biz.User{
		ID: userId,
	})
}

// UserInfo userInfo.
func (a *AppService) UserInfo(ctx context.Context, req *v1.UserInfoRequest) (*v1.UserInfoReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.UserInfo(ctx, &biz.User{
		ID: userId,
	})
}

// RecommendUpdate recommendUpdate.
func (a *AppService) RecommendUpdate(ctx context.Context, req *v1.RecommendUpdateRequest) (*v1.RecommendUpdateReply, error) {
	return &v1.RecommendUpdateReply{InviteUserAddress: ""}, nil
	//// 在上下文 context 中取出 claims 对象
	//var userId int64
	//if claims, ok := jwt.FromContext(ctx); ok {
	//	c := claims.(jwt2.MapClaims)
	//	if c["UserId"] == nil {
	//		return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
	//	}
	//	userId = int64(c["UserId"].(float64))
	//}
	//
	//var (
	//	user *biz.User
	//	err  error
	//)
	//user, err = a.uuc.GetUserByUserId(ctx, userId)
	//if nil != err {
	//	return nil, err
	//}
	//
	//if 1 == user.IsDelete {
	//	return nil, errors.New(500, "AUTHORIZE_ERROR", "用户已删除")
	//}
	//
	////var (
	////	address string
	////	res     bool
	////)
	////
	////res, address, err = verifySig2(req.SendBody.Sign, req.SendBody.PublicKey, "login")
	////if !res || nil != err || 0 >= len(address) || user.Address != address {
	////	return nil, errors.New(500, "AUTHORIZE_ERROR", "地址签名错误")
	////}
	//
	//var (
	//	res             bool
	//	addressFromSign string
	//)
	//res, addressFromSign = verifySig(req.SendBody.Sign, []byte(user.Address))
	//if !res || addressFromSign != user.Address {
	//	return nil, errors.New(500, "AUTHORIZE_ERROR", "地址签名错误")
	//}
	//return a.uuc.UpdateUserRecommend(ctx, &biz.User{
	//	ID: userId,
	//}, req)
}

// RewardList rewardList.
func (a *AppService) RewardList(ctx context.Context, req *v1.RewardListRequest) (*v1.RewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return &v1.RewardListReply{
				Status: "无效TOKEN",
			}, nil
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.RewardList(ctx, req, &biz.User{
		ID: userId,
	})
}

func (a *AppService) RecommendRewardList(ctx context.Context, req *v1.RecommendRewardListRequest) (*v1.RecommendRewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.RecommendRewardList(ctx, &biz.User{
		ID: userId,
	})
}

func (a *AppService) FeeRewardList(ctx context.Context, req *v1.FeeRewardListRequest) (*v1.FeeRewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.FeeRewardList(ctx, &biz.User{
		ID: userId,
	})
}

func (a *AppService) SetTodayList(ctx context.Context, req *v1.SetTodayListRequest) (*v1.SetTodayListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return &v1.SetTodayListReply{
				Status: "无效TOKEN",
			}, nil
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.SetTodayList(ctx, &biz.User{
		ID: userId,
	})
}

func (a *AppService) WithdrawList(ctx context.Context, req *v1.WithdrawListRequest) (*v1.WithdrawListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return &v1.WithdrawListReply{
				Status: "无效TOKEN",
			}, nil
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.WithdrawList(ctx, req, &biz.User{
		ID: userId,
	}, "usdt")
}

func (a *AppService) OrderList(ctx context.Context, req *v1.OrderListRequest) (*v1.OrderListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return &v1.OrderListReply{
				Status: "无效TOKEN",
			}, nil
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.OrderList(ctx, req, &biz.User{
		ID: userId,
	})
}

func (a *AppService) etTodayList(ctx context.Context, req *v1.WithdrawListRequest) (*v1.WithdrawListReply, error) {
	return nil, nil
}

func (a *AppService) TradeList(ctx context.Context, req *v1.TradeListRequest) (*v1.TradeListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.TradeList(ctx, &biz.User{
		ID: userId,
	})
}

func (a *AppService) TranList(ctx context.Context, req *v1.TranListRequest) (*v1.TranListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.TranList(ctx, &biz.User{
		ID: userId,
	}, req.Type, req.Tran)
}

// PasswordChange withdraw.
func (a *AppService) PasswordChange(ctx context.Context, req *v1.PasswordChangeRequest) (*v1.PasswordChangeReply, error) {
	// 在上下文 context 中取出 claims 对象
	return &v1.PasswordChangeReply{
		Password: fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password))),
	}, nil
}

// Exchange Exchange.
func (a *AppService) Exchange(ctx context.Context, req *v1.ExchangeRequest) (*v1.ExchangeReply, error) {
	return nil, nil
	//var (
	//	userId int64
	//	err    error
	//)
	//if claims, ok := jwt.FromContext(ctx); ok {
	//	c := claims.(jwt2.MapClaims)
	//	if c["UserId"] == nil {
	//		return &v1.ExchangeReply{
	//			Status: "无效TOKEN",
	//		}, nil
	//	}
	//	//if c["Password"] == nil {
	//	//	return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	//	//}
	//	userId = int64(c["UserId"].(float64))
	//	//tokenPassword = c["Password"].(string)
	//}
	//
	//var (
	//	user *biz.User
	//)
	//user, err = a.uuc.GetUserByUserId(ctx, userId)
	//if nil != err {
	//	return &v1.ExchangeReply{
	//		Status: "错误",
	//	}, nil
	//}
	//
	//if 1 == user.IsDelete {
	//	return &v1.ExchangeReply{
	//		Status: "用户已删除",
	//	}, nil
	//}
	//
	//if 1 == user.Lock {
	//	return &v1.ExchangeReply{
	//		Status: "用户已锁定",
	//	}, nil
	//}
	//
	////var (
	////	address string
	////	res     bool
	////)
	////
	////res, address, err = verifySig2(req.SendBody.Sign, req.SendBody.PublicKey, "login")
	////if !res || nil != err || 0 >= len(address) || user.Address != address {
	////	return &v1.ExchangeReply{
	////		Status: "地址签名错误",
	////	}, nil
	////}
	//
	//var (
	//	res             bool
	//	addressFromSign string
	//)
	//if 10 >= len(req.SendBody.Sign) {
	//	return &v1.ExchangeReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//
	//res, addressFromSign = verifySig(req.SendBody.Sign, []byte(user.Address))
	//if !res || addressFromSign != user.Address {
	//	return &v1.ExchangeReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//
	//return a.uuc.Exchange(ctx, req, &biz.User{
	//	ID: userId,
	//})
}

//var lockBuy sync.Mutex

// Buy  buySomething.
func (a *AppService) Buy(ctx context.Context, req *v1.BuyRequest) (*v1.BuyReply, error) {
	return nil, nil
	// 在上下文 context 中取出 claims 对象
	//lockBuy.Lock()
	//defer lockBuy.Unlock()
	//
	//var (
	//	//err           error
	//	userId int64
	//)
	//
	//if claims, ok := jwt.FromContext(ctx); ok {
	//	c := claims.(jwt2.MapClaims)
	//	if c["UserId"] == nil {
	//		return &v1.BuyReply{
	//			Status: "无效TOKEN",
	//		}, nil
	//	}
	//	//if c["Password"] == nil {
	//	//	return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	//	//}
	//	userId = int64(c["UserId"].(float64))
	//	//tokenPassword = c["Password"].(string)
	//}
	//
	//// 验证
	//var (
	//	err error
	//)
	//
	//var (
	//	user *biz.User
	//)
	//user, err = a.uuc.GetUserByUserId(ctx, userId)
	//if nil != err {
	//	return &v1.BuyReply{
	//		Status: "错误",
	//	}, nil
	//}
	//
	//if 1 == user.IsDelete {
	//	return &v1.BuyReply{
	//		Status: "用户已删除",
	//	}, nil
	//}
	//
	//if 1 == user.Lock {
	//	return &v1.BuyReply{
	//		Status: "用户已锁定",
	//	}, nil
	//}
	//
	////fmt.Println(user)
	////res, address, err = verifySig2(req.SendBody.Sign, req.SendBody.PublicKey, "login")
	////if !res || nil != err || 0 >= len(address) || address != user.Address {
	////	return nil, errors.New(500, "AUTHORIZE_ERROR", "地址签名错误")
	////}
	//
	//var (
	//	res             bool
	//	addressFromSign string
	//)
	//if 10 >= len(req.SendBody.Sign) {
	//	return &v1.BuyReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//res, addressFromSign = verifySig(req.SendBody.Sign, []byte(user.Address))
	//if !res || addressFromSign != user.Address {
	//	return &v1.BuyReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//
	////if "" == req.SendBody.Password || 6 > len(req.SendBody.Password) {
	////	return nil, errors.New(500, "AUTHORIZE_ERROR", "账户密码必须大于6位")
	////}
	//// TODO 验证签名
	////password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	//
	//return a.uuc.Buy(ctx, req, user)
}

//var lockSetToday sync.Mutex

// SetToday  SetToday.
func (a *AppService) SetToday(ctx context.Context, req *v1.SetTodayRequest) (*v1.SetTodayReply, error) {
	return nil, nil
	//lockSetToday.Lock()
	//defer lockSetToday.Unlock()
	//
	//// 在上下文 context 中取出 claims 对象
	//var (
	//	//err           error
	//	userId int64
	//)
	//
	//if claims, ok := jwt.FromContext(ctx); ok {
	//	c := claims.(jwt2.MapClaims)
	//	if c["UserId"] == nil {
	//		return &v1.SetTodayReply{
	//			Status: "无效TOKEN",
	//		}, nil
	//	}
	//	//if c["Password"] == nil {
	//	//	return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	//	//}
	//	userId = int64(c["UserId"].(float64))
	//	//tokenPassword = c["Password"].(string)
	//}
	//
	//// 验证
	//var (
	//	err error
	//)
	//
	//var (
	//	user *biz.User
	//)
	//user, err = a.uuc.GetUserByUserId(ctx, userId)
	//if nil != err {
	//	return &v1.SetTodayReply{
	//		Status: "错误",
	//	}, nil
	//}
	//
	//if 1 == user.IsDelete {
	//	return &v1.SetTodayReply{
	//		Status: "用户已删除",
	//	}, nil
	//}
	//
	//if 1 == user.Lock {
	//	return &v1.SetTodayReply{
	//		Status: "用户已锁定",
	//	}, nil
	//}
	//
	////fmt.Println(user)
	////res, address, err = verifySig2(req.SendBody.Sign, req.SendBody.PublicKey, "login")
	////if !res || nil != err || 0 >= len(address) || address != user.Address {
	////	return nil, errors.New(500, "AUTHORIZE_ERROR", "地址签名错误")
	////}
	//
	//var (
	//	res             bool
	//	addressFromSign string
	//)
	//if 10 >= len(req.SendBody.Sign) {
	//	return &v1.SetTodayReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//res, addressFromSign = verifySig(req.SendBody.Sign, []byte(user.Address))
	//if !res || addressFromSign != user.Address {
	//	return &v1.SetTodayReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//
	//return a.uuc.SetToday(ctx, req, user)
}

// AmountTo AmountTo.
func (a *AppService) AmountTo(ctx context.Context, req *v1.AmountToRequest) (*v1.AmountToReply, error) {
	return nil, nil
	//// 在上下文 context 中取出 claims 对象
	//var (
	//	//err           error
	//	userId int64
	//)
	//
	//if claims, ok := jwt.FromContext(ctx); ok {
	//	c := claims.(jwt2.MapClaims)
	//	if c["UserId"] == nil {
	//		return &v1.AmountToReply{
	//			Status: "无效TOKEN",
	//		}, nil
	//	}
	//	//if c["Password"] == nil {
	//	//	return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	//	//}
	//	userId = int64(c["UserId"].(float64))
	//	//tokenPassword = c["Password"].(string)
	//}
	//
	//// 验证
	//var (
	//	err error
	//)
	//
	//var (
	//	user *biz.User
	//)
	//user, err = a.uuc.GetUserByUserId(ctx, userId)
	//if nil != err {
	//	return &v1.AmountToReply{
	//		Status: "错误",
	//	}, nil
	//}
	//
	//if 1 == user.IsDelete {
	//	return &v1.AmountToReply{
	//		Status: "用户已删除",
	//	}, nil
	//}
	//
	//if 1 == user.Lock {
	//	return &v1.AmountToReply{
	//		Status: "用户已锁定",
	//	}, nil
	//}
	//
	////fmt.Println(user)
	////res, address, err = verifySig2(req.SendBody.Sign, req.SendBody.PublicKey, "login")
	////if !res || nil != err || 0 >= len(address) || address != user.Address {
	////	return nil, errors.New(500, "AUTHORIZE_ERROR", "地址签名错误")
	////}
	//
	//var (
	//	res             bool
	//	addressFromSign string
	//)
	//if 10 >= len(req.SendBody.Sign) {
	//	return &v1.AmountToReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//res, addressFromSign = verifySig(req.SendBody.Sign, []byte(user.Address))
	//if !res || addressFromSign != user.Address {
	//	return &v1.AmountToReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//
	////if "" == req.SendBody.Password || 6 > len(req.SendBody.Password) {
	////	return nil, errors.New(500, "AUTHORIZE_ERROR", "账户密码必须大于6位")
	////}
	//// TODO 验证签名
	////password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	//
	//return a.uuc.AmountTo(ctx, req, user)
}

// Stake Stake.
func (a *AppService) Stake(ctx context.Context, req *v1.StakeRequest) (*v1.StakeReply, error) {
	return nil, nil
	//// 在上下文 context 中取出 claims 对象
	//var (
	//	//err           error
	//	userId int64
	//)
	//
	//if claims, ok := jwt.FromContext(ctx); ok {
	//	c := claims.(jwt2.MapClaims)
	//	if c["UserId"] == nil {
	//		return &v1.StakeReply{
	//			Status: "无效TOKEN",
	//		}, nil
	//	}
	//	//if c["Password"] == nil {
	//	//	return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	//	//}
	//	userId = int64(c["UserId"].(float64))
	//	//tokenPassword = c["Password"].(string)
	//}
	//
	//// 验证
	//var (
	//	err error
	//)
	//
	//var (
	//	user *biz.User
	//)
	//user, err = a.uuc.GetUserByUserId(ctx, userId)
	//if nil != err {
	//	return &v1.StakeReply{
	//		Status: "错误",
	//	}, nil
	//}
	//
	//if 1 == user.IsDelete {
	//	return &v1.StakeReply{
	//		Status: "用户已删除",
	//	}, nil
	//}
	//
	//if 1 == user.Lock {
	//	return &v1.StakeReply{
	//		Status: "用户已锁定",
	//	}, nil
	//}
	//
	////fmt.Println(user)
	////res, address, err = verifySig2(req.SendBody.Sign, req.SendBody.PublicKey, "login")
	////if !res || nil != err || 0 >= len(address) || address != user.Address {
	////	return nil, errors.New(500, "AUTHORIZE_ERROR", "地址签名错误")
	////}
	//
	//var (
	//	res             bool
	//	addressFromSign string
	//)
	//if 10 >= len(req.SendBody.Sign) {
	//	return &v1.StakeReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//res, addressFromSign = verifySig(req.SendBody.Sign, []byte(user.Address))
	//if !res || addressFromSign != user.Address {
	//	return &v1.StakeReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//
	////if "" == req.SendBody.Password || 6 > len(req.SendBody.Password) {
	////	return nil, errors.New(500, "AUTHORIZE_ERROR", "账户密码必须大于6位")
	////}
	//// TODO 验证签名
	////password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	//
	//return a.uuc.Stake(ctx, req, user)
}

// UnStake UnStake.
func (a *AppService) UnStake(ctx context.Context, req *v1.UnStakeRequest) (*v1.UnStakeReply, error) {
	return nil, nil
	//// 在上下文 context 中取出 claims 对象
	//var (
	//	//err           error
	//	userId int64
	//)
	//
	//if claims, ok := jwt.FromContext(ctx); ok {
	//	c := claims.(jwt2.MapClaims)
	//	if c["UserId"] == nil {
	//		return &v1.UnStakeReply{
	//			Status: "无效TOKEN",
	//		}, nil
	//	}
	//	//if c["Password"] == nil {
	//	//	return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	//	//}
	//	userId = int64(c["UserId"].(float64))
	//	//tokenPassword = c["Password"].(string)
	//}
	//
	//// 验证
	//var (
	//	err error
	//)
	//
	//var (
	//	user *biz.User
	//)
	//user, err = a.uuc.GetUserByUserId(ctx, userId)
	//if nil != err {
	//	return &v1.UnStakeReply{
	//		Status: "错误",
	//	}, nil
	//}
	//
	//if 1 == user.IsDelete {
	//	return &v1.UnStakeReply{
	//		Status: "用户已删除",
	//	}, nil
	//}
	//
	//if 1 == user.Lock {
	//	return &v1.UnStakeReply{
	//		Status: "用户已锁定",
	//	}, nil
	//}
	//
	////fmt.Println(user)
	////res, address, err = verifySig2(req.SendBody.Sign, req.SendBody.PublicKey, "login")
	////if !res || nil != err || 0 >= len(address) || address != user.Address {
	////	return nil, errors.New(500, "AUTHORIZE_ERROR", "地址签名错误")
	////}
	//
	//var (
	//	res             bool
	//	addressFromSign string
	//)
	//if 10 >= len(req.SendBody.Sign) {
	//	return &v1.UnStakeReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//res, addressFromSign = verifySig(req.SendBody.Sign, []byte(user.Address))
	//if !res || addressFromSign != user.Address {
	//	return &v1.UnStakeReply{
	//		Status: "签名错误",
	//	}, nil
	//}
	//
	////if "" == req.SendBody.Password || 6 > len(req.SendBody.Password) {
	////	return nil, errors.New(500, "AUTHORIZE_ERROR", "账户密码必须大于6位")
	////}
	//// TODO 验证签名
	////password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	//
	//return a.uuc.UnStake(ctx, req, user)
}

var lockWithdraw sync.Mutex

// Withdraw withdraw.
func (a *AppService) Withdraw(ctx context.Context, req *v1.WithdrawRequest) (*v1.WithdrawReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		err           error
		userId        int64
		tokenPassword string
	)

	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return &v1.WithdrawReply{
				Status: "无效TOKEN",
			}, nil
		}
		//if c["Password"] == nil {
		//	return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
		//}
		userId = int64(c["UserId"].(float64))
		//tokenPassword = c["Password"].(string)
	}

	var (
		user *biz.User
	)
	user, err = a.uuc.GetUserByUserId(ctx, userId)
	if nil != err {
		return &v1.WithdrawReply{
			Status: "错误",
		}, nil
	}

	if 1 == user.IsDelete {
		return &v1.WithdrawReply{
			Status: "用户已删除",
		}, nil
	}

	if 1 == user.Lock {
		return &v1.WithdrawReply{
			Status: "用户已锁定",
		}, nil
	}

	lockWithdraw.Lock()
	defer lockWithdraw.Unlock()

	//var (
	//	address string
	//	res     bool
	//)
	//
	//res, address, err = verifySig2(req.SendBody.Sign, req.SendBody.PublicKey, "login")
	//if !res || nil != err || 0 >= len(address) || user.Address != address {
	//	return nil, errors.New(500, "AUTHORIZE_ERROR", "地址签名错误")
	//}

	var (
		res             bool
		addressFromSign string
	)
	if 10 >= len(req.SendBody.Sign) {
		return &v1.WithdrawReply{
			Status: "签名错误",
		}, nil
	}
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(user.Address))
	if !res || addressFromSign != user.Address {
		return &v1.WithdrawReply{
			Status: "签名错误",
		}, nil
	}

	//if "" == req.SendBody.Password || 6 > len(req.SendBody.Password) {
	//	return nil, errors.New(500, "AUTHORIZE_ERROR", "账户密码必须大于6位")
	//}
	// TODO 验证签名
	//password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))

	return a.uuc.Withdraw(ctx, req, &biz.User{
		ID:       userId,
		Password: tokenPassword,
	})
}

// Tran tran .
func (a *AppService) Tran(ctx context.Context, req *v1.TranRequest) (*v1.TranReply, error) {
	return nil, nil
	//// 在上下文 context 中取出 claims 对象
	//var (
	//	userId        int64
	//	tokenPassword string
	//)
	//if claims, ok := jwt.FromContext(ctx); ok {
	//	c := claims.(jwt2.MapClaims)
	//	if c["UserId"] == nil {
	//		return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
	//	}
	//	if c["Password"] == nil {
	//		return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	//	}
	//
	//	userId = int64(c["UserId"].(float64))
	//	tokenPassword = c["Password"].(string)
	//}
	//
	//if "" == req.SendBody.Password || 6 > len(req.SendBody.Password) {
	//	return nil, errors.New(500, "AUTHORIZE_ERROR", "账户密码必须大于6位")
	//}
	//// TODO 验证签名
	//password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	//
	//return a.uuc.Tran(ctx, req, &biz.User{
	//	ID:       userId,
	//	Password: tokenPassword,
	//}, password)
}

func (a *AppService) GetTrade(ctx context.Context, req *v1.GetTradeRequest) (*v1.GetTradeReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		amountB        int64
		tmpValue       int64
		hbs            float64
		amountFloatHbs float64
		amountFloatCsd float64
		csd            string
		err            error
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloatCsd = amountFloat * 10000000000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloatCsd, 'f', -1, 64), 10, 64)
	if 10000000000 > amount {
		return nil, errors.New(500, "ERROR_TOKEN", "输入错误")
	}

	csd, err = GetAmountOut(req.SendBody.Amount + "000000000000000000")
	if nil != err {
		fmt.Println(2)
		return nil, errors.New(500, "ERROR_TOKEN", "查询币价错误")
	}
	lenValue := len(csd)
	if 10 > lenValue {
		return nil, errors.New(500, "ERROR_TOKEN", "币价过低")
	}
	tmpValue, _ = strconv.ParseInt(csd[0:lenValue-8], 10, 64)
	if 0 == tmpValue {
		return nil, errors.New(500, "ERROR_TOKEN", "币价过低")
	}

	hbs, err = requestHbsResult()
	if nil != err {
		fmt.Println(1)
		return nil, errors.New(500, "ERROR_TOKEN", "查询币价错误")
	}
	amountFloatHbs = amountFloat * 10
	amountB = int64(amountFloatHbs / hbs * 10000000000)
	if 0 >= amountB {
		return nil, errors.New(500, "ERROR_TOKEN", "币价错误")
	}

	return &v1.GetTradeReply{
		AmountCsd: fmt.Sprintf("%.4f", float64(tmpValue)/float64(10000000000)),
		AmountHbs: fmt.Sprintf("%.4f", float64(amountB)/float64(10000000000)),
	}, nil
}

func (a *AppService) Trade(ctx context.Context, req *v1.WithdrawRequest) (*v1.WithdrawReply, error) {
	return nil, nil
}

// SetBalanceReward .
func (a *AppService) SetBalanceReward(ctx context.Context, req *v1.SetBalanceRewardRequest) (*v1.SetBalanceRewardReply, error) {
	return nil, nil
	// 在上下文 context 中取出 claims 对象
	//var userId int64
	//if claims, ok := jwt.FromContext(ctx); ok {
	//	c := claims.(jwt2.MapClaims)
	//	if c["UserId"] == nil {
	//		return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
	//	}
	//	userId = int64(c["UserId"].(float64))
	//}
	//
	//return a.uuc.SetBalanceReward(ctx, req, &biz.User{
	//	ID: userId,
	//})
}

func (a *AppService) UserRecommend(ctx context.Context, req *v1.RecommendListRequest) (*v1.RecommendListReply, error) {
	return a.uuc.UserRecommend(ctx, req)
}

// DeleteBalanceReward .
func (a *AppService) DeleteBalanceReward(ctx context.Context, req *v1.DeleteBalanceRewardRequest) (*v1.DeleteBalanceRewardReply, error) {
	return nil, nil
	// 在上下文 context 中取出 claims 对象
	//var userId int64
	//if claims, ok := jwt.FromContext(ctx); ok {
	//	c := claims.(jwt2.MapClaims)
	//	if c["UserId"] == nil {
	//		return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
	//	}
	//	userId = int64(c["UserId"].(float64))
	//}
	//
	//return a.uuc.DeleteBalanceReward(ctx, req, &biz.User{
	//	ID: userId,
	//})
}

func (a *AppService) AdminRewardList(ctx context.Context, req *v1.AdminRewardListRequest) (*v1.AdminRewardListReply, error) {
	return a.uuc.AdminRewardList(ctx, req)
}

func (a *AppService) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListReply, error) {
	return a.uuc.AdminUserList(ctx, req)
}

func (a *AppService) AdminLocationList(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	return a.uuc.AdminLocationList(ctx, req)
}

func (a *AppService) AdminWithdrawList(ctx context.Context, req *v1.AdminWithdrawListRequest) (*v1.AdminWithdrawListReply, error) {
	return a.uuc.AdminWithdrawList(ctx, req)
}

func (a *AppService) AdminWithdraw(ctx context.Context, req *v1.AdminWithdrawRequest) (*v1.AdminWithdrawReply, error) {
	return a.uuc.AdminWithdraw(ctx, req)
}

func (a *AppService) AdminFee(ctx context.Context, req *v1.AdminFeeRequest) (*v1.AdminFeeReply, error) {
	return a.uuc.AdminFee(ctx, req)
}

func (a *AppService) AdminAll(ctx context.Context, req *v1.AdminAllRequest) (*v1.AdminAllReply, error) {
	return a.uuc.AdminAll(ctx, req)
}

func (a *AppService) AdminUserRecommend(ctx context.Context, req *v1.AdminUserRecommendRequest) (*v1.AdminUserRecommendReply, error) {
	return a.uuc.AdminRecommendList(ctx, req)
}

func (a *AppService) AdminMonthRecommend(ctx context.Context, req *v1.AdminMonthRecommendRequest) (*v1.AdminMonthRecommendReply, error) {
	return a.uuc.AdminMonthRecommend(ctx, req)
}

func (a *AppService) AdminConfig(ctx context.Context, req *v1.AdminConfigRequest) (*v1.AdminConfigReply, error) {
	return a.uuc.AdminConfig(ctx, req)
}

func (a *AppService) AdminConfigUpdate(ctx context.Context, req *v1.AdminConfigUpdateRequest) (*v1.AdminConfigUpdateReply, error) {
	return a.uuc.AdminConfigUpdate(ctx, req)
}

func (a *AppService) AdminWithdrawEth(ctx context.Context, req *v1.AdminWithdrawEthRequest) (*v1.AdminWithdrawEthReply, error) {
	return &v1.AdminWithdrawEthReply{}, nil
}

func (a *AppService) TokenWithdraw(ctx context.Context, req *v1.TokenWithdrawRequest) (*v1.TokenWithdrawReply, error) {

	var (
		err error
	)
	for i := 0; i <= 5; i++ {
		tmpUrl1 := "https://bnb-bscnews.rpc.blxrbdn.com/"
		_, err = tokenWithdraw(tmpUrl1, 56)
		if err == nil {
			break
		} else if "insufficient funds for gas * price + value" == err.Error() {
			fmt.Println(5555, err)
		} else if "execution reverted: ERC20: transfer amount exceeds balance" == err.Error() {
			fmt.Println(4444, err)
			break
		} else if "execution reverted: time limit" == err.Error() {
			fmt.Println(4441, err)
			break
		} else {
			if 0 == i {
				tmpUrl1 = "https://bsc-dataseed4.binance.org/"
			} else if 1 == i {
				tmpUrl1 = "https://bsc-dataseed3.binance.org"
			} else if 2 == i {
				tmpUrl1 = "https://bsc-dataseed2.binance.org"
			} else if 3 == i {
				tmpUrl1 = "https://bsc-dataseed1.binance.org"
			} else if 4 == i {
				tmpUrl1 = "https://bsc-dataseed.binance.org"
			}
			fmt.Println(3333, err)
			time.Sleep(3 * time.Second)
		}
	}

	return &v1.TokenWithdrawReply{}, nil
}

func tokenWithdraw(requestUrl string, chainId int64) (bool, error) {

	client, err := ethclient.Dial(requestUrl)
	//client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return false, err
	}

	tokenAddress := common.HexToAddress("0xFC13153Bb4D285939FD23c7899eAdD785fBf6aA2")
	instance, err := NewTokenWithdraw(tokenAddress, client)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	var authUser *bind.TransactOpts

	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.HexToECDSA("")
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(chainId))
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	//var res *types.Transaction
	_, err = instance.WithdrawSx(&bind.TransactOpts{
		From:     authUser.From,
		Signer:   authUser.Signer,
		GasPrice: gasPrice,
		GasLimit: 0,
	})
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	//fmt.Println(res.Hash())

	return true, nil
}

func GetAmountOut(strAmount string) (string, error) {

	var balString string
	url1 := "https://bnb-bscnews.rpc.blxrbdn.com/"

	for i := 4; i < 16; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			return "", err
		}

		tokenAddress := common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
		instance, err := NewPancakerouterv2(tokenAddress, client)
		if err != nil {
			return "", err
		}

		addresses := make([]common.Address, 0)
		addresses = append(addresses, common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"), common.HexToAddress("0xfAd476cd33Ed9213ED0a2F4c20f6865A98bf0a8B"))
		amount, _ := new(big.Int).SetString(strAmount, 10)

		bals, err := instance.GetAmountsOut(&bind.CallOpts{}, amount, addresses)
		if err != nil {
			fmt.Println(err)
			if 0 == i%4 {
				url1 = "https://bsc-dataseed4.binance.org"
			} else if 1 == i%4 {
				url1 = "https://bsc-dataseed1.binance.org"
			} else if 2 == i%4 {
				url1 = "https://bsc-dataseed.binance.org"
			} else if 3 == i%4 {
				url1 = "https://bsc-dataseed3.binance.org"
			}
			continue
		}
		balString = bals[1].String()
		break
	}

	return balString, nil
}

func GetAmountOut1(strAmount string) (string, error) {

	var balString string
	url1 := "https://bnb-bscnews.rpc.blxrbdn.com/"

	client, err := ethclient.Dial(url1)
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
	instance, err := NewPancakerouterv2(tokenAddress, client)
	if err != nil {
		return "", err
	}

	addresses := make([]common.Address, 0)
	addresses = append(addresses, common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"), common.HexToAddress("0x538ac017aa01ba9665052660ea5783ba91a48092"))
	amount, _ := new(big.Int).SetString(strAmount, 10)

	bals, err := instance.GetAmountsOut(&bind.CallOpts{}, amount, addresses)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	balString = bals[1].String()
	return balString, nil
}

type eth struct {
	CoinId string
	Usd    float64
}

func requestHbsResult() (float64, error) {
	//apiUrl := "https://api-testnet.bscscan.com/api"
	apiUrl := "https://be.api.hbsswap.com/market/coin/rates"
	// URL param
	data := url.Values{}

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return 0, err
	}
	u.RawQuery = data.Encode() // URL encode
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(u.String())
	if err != nil {
		return 0, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var i struct {
		Data []*eth `json:"Data"`
	}
	err = json.Unmarshal(b, &i)
	if err != nil {
		return 0, err
	}

	var price float64
	for _, v := range i.Data {
		if "HBS(BEP20)" == v.CoinId { // 接收者
			price = v.Usd
		}
	}

	return price, err
}

func addressCheck(addressParam string) (bool, error) {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(addressParam) {
		return false, nil
	}

	client, err := ethclient.Dial("https://bsc-dataseed4.binance.org/")
	if err != nil {
		return false, err
	}

	// a random user account address
	address := common.HexToAddress(addressParam)
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		return false, err
	}

	if len(bytecode) > 0 {
		return false, nil
	}

	return true, nil
}

//func checkSign() {
//	verifySig(
//		"0x2c1f0a2dc0ba6c9e7a673f4af0a9be4ccebc0377e98120e0878f03d2547be19d5a3efc23dd1057f51968b3861036e23f52c816482d4444e6780690d9f7999ccf1c",
//		[]byte("Hello"),
//	)
//}

func verifySig(sigHex string, msg []byte) (bool, string) {
	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false, ""
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	sigPublicKeyBytes := crypto.FromECDSAPub(recovered)
	signatureNoRecoverID := sig[:len(sig)-1] // remove recovery id
	verified := crypto.VerifySignature(sigPublicKeyBytes, msg, signatureNoRecoverID)
	return verified, recoveredAddr.Hex()
}
