package variable

import (
	LogisticsManagement "backed/internal/contract"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/go-redis/v9"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Viper           *viper.Viper    // 读取配置文件包
	Gorm            *gorm.DB        // 存储mysql连接
	RedisClient     *redis.Client   // 存储redis客户端
	FiscoBcos       *client.Client  // fisco bcos客户端
	AliClient       *alipay.Client  // 支付宝客户端
	ContractAddress *common.Address // 合约地址
	ContractSession *LogisticsManagement.LogisticsManagementSession
	ContractAbi     string

	MessageChannel chan string // 消息发布队列
	RefundChannel  chan string // 退款队列
)
