package major

import (
	"fmt"
	"github.com/blocktree/go-openw-server/open_base/model"
	"github.com/godaddy-x/jorm/amqp"
	"github.com/godaddy-x/jorm/cache/redis"
	"github.com/godaddy-x/jorm/consul"
	"github.com/godaddy-x/jorm/log"
	"github.com/godaddy-x/jorm/sqlc"
	"github.com/godaddy-x/jorm/sqld"
	"github.com/godaddy-x/jorm/util"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	network      = new(Network)
	multiconfMap = map[string]*MultiConfig{}
	Mq_apps      = make(map[string]string, 0)
	Mq_symbols   = make(map[string]string, 0)
	Mq_consuls   = make(map[string]string, 0)
	aeskey       = "QIB@rncytfN5J^f0"
	tradekey     = "Yc@r#iUaXj^dX6vc"
)

type TxQueue struct {
	Exchange string `json:"exchange"`
	Queue    string `json:"queue"`
	Symbol   string `json:"symbol"`
}

type SymbolRSV struct {
	Symbol string
	RSV    bool
}

type MultiConfig struct {
	ApiUrl              string
	AppID               string
	AppKey              string
	AppPriKey           string
	OpenSSL             bool
	CreateAccountExpire int64
	CreateTradeExpire   int64
	UserType            int64
	SymbolRSV           []SymbolRSV
	JPushAppKey         string
	JPushSecretKey      string
	JPushMode           int64 // 0.开发模式 1.测试模式 2.生产模式
}

// 区块地址配置
type Network struct {
	IsLoad       bool
	AddrCount    int64
	AddrIsChange int64
	NodeSeed     int64
	IsTestNet    bool
	Secret       string
	Exptime      int64
	Issub        string
	DxSend       bool
	OWTPAuth     bool
	ScanLogPath  string
	SSLKey       string
	MultiAppID   string
	MQSecretKey  string
}

type ProLogConfig struct {
	Dir     string
	Name    string
	Level   string
	Console bool
}

// 日志配置
type LogConfig struct {
	Webapp   ProLogConfig `json:"webapp"`
	Stdrpc   ProLogConfig `json:"stdrpc"`
	Scanner  ProLogConfig `json:"scanner"`
	Develop  ProLogConfig `json:"develop"`
	Admin    ProLogConfig `json:"admin"`
	Consumer ProLogConfig `json:"consumer"`
	Multisig ProLogConfig `json:"multisig"`
}

// 监控配置
type Dxmonitor struct {
	Open     bool   `json:"open"`
	Symbol   string `json:"symbol"`
	WaitTime int64  `json:"waitTime"`
	SendTime int64  `json:"sendTime"`
	MaxSend  int64  `json:"maxSend"`
	Mobile   string `json:"mobile"`
	Memo     string `json:"memo"`
	Server   string `json:"server"`
	Ipport   string `json:"ipport"`
}

type Dxconf struct {
	Prefix    string `json:"prefix"`
	Symbol    string `json:"symbol"`
	Trysend   int64  `json:"trysend"`
	Height    int64  `json:"height"`
	Lastdate  int64  `json:"lastdate"`
	Blockdate int64  `json:"blockdate"`
}

// account - eos, pia
type CacheValue struct {
	V string
}

func GetTradeLogSignature(trade *model.OwTradeLog) string {
	return util.MD5(util.AddStr(trade.AppID, trade.AccountID, trade.Symbol, trade.Type, trade.Txid, trade.ContractID, trade.Amount, trade.Fees), tradekey)
}

func InitNetwork() *Network {
	if !network.IsLoad {
		consulx, err := new(consul.ConsulManager).Client()
		if err != nil {
			panic(err)
		}
		if err := consulx.ReadJsonConfig("rpc/network", network); err != nil {
			panic("加载NetWork配置失败")
		}
		network.IsLoad = true
	}
	return network
}

func InitMultiConf() {
	consulx, err := new(consul.ConsulManager).Client()
	if err != nil {
		panic(err)
	}
	if err := consulx.ReadJsonConfig("rpc/multiconf", &multiconfMap); err != nil {
		panic("加载Multisig配置失败")
	}
}

func GetMultiConf(appID string) (*MultiConfig, error) {
	if len(appID) == 0 {
		return nil, util.Error("AppID为空")
	}
	if v, ok := multiconfMap[appID]; ok {
		return v, nil
	} else {
		return nil, util.Error("无效的AppID")
	}
}

func GetAllMultiConf() map[string]*MultiConfig {
	return multiconfMap
}

func InitTxQueue() map[string]TxQueue {
	dc := InitDc()
	txQueues := make([]TxQueue, 0)
	if err := dc.ReadJsonConfig("rpc/txqueue", &txQueues); err != nil {
		panic(util.AddStr("读取TxQueue配置失败: ", err.Error()))
	}
	result := make(map[string]TxQueue, 0)
	for e := range txQueues {
		v := txQueues[e]
		result[v.Symbol] = v
	}
	return result
}

func ReadNodeAesData(dc *consul.ConsulManager, node string, result interface{}) error {
	if data, err := dc.GetKV(node); err != nil {
		return err
	} else {
		//fmt.Println(util.AddStr("读取node[", node, "] - ", string(data)))
		rs := AesDecrypt(string(data), aeskey)
		//fmt.Println(util.AddStr("解密node[", node, "] - ", rs))
		if err := util.JsonToObject(rs, result); err != nil {
			log.Println("解析配置文件失败: " + err.Error())
			return err
		}
		return nil
	}
}

func InitDB() {

	dc := InitDc()

	redis := []cache.RedisConfig{}
	if err := ReadNodeAesData(dc, "cache/redis", &redis); err != nil {
		panic(util.AddStr("读取redis配置失败: ", err.Error()))
	}
	new(cache.RedisManager).InitConfig(redis...)
	client, err := new(cache.RedisManager).Client()
	if err != nil {
		panic(err.Error())
	}

	mysql := []sqld.MysqlConfig{}
	if err := ReadNodeAesData(dc, "dbx/mysql", &mysql); err != nil {
		panic(util.AddStr("读取mysql配置失败: ", err.Error()))
	}
	new(sqld.MysqlManager).InitConfigAndCache(client, mysql...)

	mongo := []sqld.MGOConfig{}
	if err := ReadNodeAesData(dc, "dbx/mongo", &mongo); err != nil {
		panic(util.AddStr("读取mongo配置失败: ", err.Error()))
	}
	     
	if len(mongo) > 0  {
			mongo[0].Debug = false
	}
	new(sqld.MGOManager).InitConfigAndCache(client, mongo...)

	amqp := []rabbitmq.AmqpConfig{}
	if err := ReadNodeAesData(dc, "rpc/amqp", &amqp); err != nil {
		panic(util.AddStr("读取mq配置失败: ", err.Error()))
	}
	new(rabbitmq.PublishManager).InitConfig(amqp...)
	if err := ReadNodeAesData(dc, "rpc/amqp", &amqp); err != nil {
		panic(util.AddStr("读取mq配置失败: ", err.Error()))
	}
	new(rabbitmq.PullManager).InitConfig(amqp...)
}

func InitDc() *consul.ConsulManager {
	consulx := []consul.ConsulConfig{
		{Host: consul.DefaultHost, Node: "dc/consul"},
	}
	dc, err := new(consul.ConsulManager).InitConfig(consulx...);
	if err != nil {
		panic(util.AddStr("读取consul配置失败: ", err.Error()))
	}
	dc, err = dc.Client()
	if err != nil {
		panic(err)
	}
	return dc
}

func InitLog(pname string) {
	dc := InitDc()
	result := LogConfig{}
	if err := dc.ReadJsonConfig("rpc/logger", &result); err != nil {
		panic(util.AddStr("读取logger配置失败: ", err.Error()))
	}
	dir := ""
	name := ""
	console := false
	level := "error"
	if pname == "webapp" {
		dir = result.Webapp.Dir
		name = result.Webapp.Name
		console = result.Webapp.Console
		level = result.Webapp.Level
	} else if pname == "stdrpc" {
		dir = result.Stdrpc.Dir
		name = result.Stdrpc.Name
		console = result.Stdrpc.Console
		level = result.Stdrpc.Level
	} else if strings.HasPrefix(pname, "scanner") {
		dir = result.Scanner.Dir
		name = pname + ".log"
		console = result.Scanner.Console
		level = result.Scanner.Level
	} else if pname == "develop" {
		dir = result.Develop.Dir
		name = result.Develop.Name
		console = result.Develop.Console
		level = result.Develop.Level
	} else if pname == "admin" {
		dir = result.Admin.Dir
		name = result.Admin.Name
		console = result.Admin.Console
		level = result.Admin.Level
	} else if pname == "consumer" {
		dir = result.Consumer.Dir
		name = result.Consumer.Name
		console = result.Consumer.Console
		level = result.Consumer.Level
	} else if pname == "multisig" {
		dir = result.Multisig.Dir
		name = result.Multisig.Name
		console = result.Multisig.Console
		level = result.Multisig.Level
	}
	log.InitDefaultLog(&log.ZapConfig{
		Level:   level,
		Console: console,
		FileConfig: &log.FileConfig{
			Compress:   true,
			Filename:   dir + name,
			MaxAge:     7,
			MaxBackups: 7,
			MaxSize:    512,
		}})
	log.Info("日志服务启动成功", 0, log.String("project", pname), log.String("setLevel", level), log.String("filePath", dir+name))
}

func GenMQDataSig(result interface{}) (string, string, error) {
	str, err := util.ObjectToJson(result)
	if err != nil || len(str) == 0 {
		return "", "", util.Error("区块/交易单数据转换JSON失败")
	}
	ret := util.Base64URLEncode(str)
	if len(str) == 0 {
		return "", "", util.Error("区块/交易单数据base64编码失败")
	}
	sig := util.MD5(ret, InitNetwork().MQSecretKey)
	return ret, sig, nil
}

func CheckMQDataSig(message rabbitmq.MsgData, data interface{}) error {
	if len(message.Signature) == 0 {
		return util.Error("MQ客户端消费区块头/交易单数据日志,消息内容无签名")
	}
	if v, b := message.Content.(string); !b {
		return util.Error("MQ客户端消费区块头/交易单数据日志,消息内容非string类型")
	} else {
		if message.Signature != util.MD5(v, InitNetwork().MQSecretKey) {
			return util.Error("MQ客户端消费区块头/交易单数据日志,消息内容签名校验失败")
		}
		str := util.Base64URLDecode(v)
		if len(str) == 0 {
			return util.Error("MQ客户端消费区块头/交易单数据日志,消息内容为空")
		}
		if err := util.JsonToObject(str, data); err != nil {
			return util.Error("MQ客户端消费区块头/交易单数据日志,消息内容JSON转换失败: ", err)
		}
		log.Debug("MQ客户端消费区块头/交易单数据日志", util.Time(), log.Any("message", data))
	}
	return nil
}

func InitDxmonitor() {
	client, err := new(cache.RedisManager).Client()
	if err != nil {
		panic(util.AddStr("加载redis服务失败: ", err.Error()))
	}
	// client.Del("dx.block.coin.ETH")
	go func() {
		for ; ; {
			mongo, err := new(sqld.MGOManager).Get()
			if err != nil {
				continue
			}
			defer mongo.Close()
			conf := model.OwSetting{}
			if err := mongo.FindOne(sqlc.M(model.OwSetting{}).Eq("platform", "scanner"), &conf); err != nil {
				continue
			}
			if conf.Id > 0 && conf.IsOff > 0 {
				continue
			}
			// 延时30秒执行
			time.Sleep(30 * time.Second)
			dc := InitDc()
			vs := []Dxmonitor{}
			if err := dc.ReadJsonConfig("rpc/dxmonitor", &vs); err != nil {
				log.Error(util.AddStr("读取dxmonitor配置失败: ", err.Error()), 0)
			}
			for _, v := range vs {
				if !v.Open {
					continue
				}
				block1 := &model.OwBlock{}
				if err := getMaxBlock(v.Symbol, block1); err != nil {
					log.Error(util.AddStr("读取货币[", v.Symbol, "]最新高度失败: ", err.Error()), 0)
					continue
				}
				blockmax := model.OwBlockMax{}
				if err := mongo.FindOne(sqlc.M(model.OwBlockMax{}).Eq("symbol", v.Symbol), &blockmax); err != nil {
					log.Error(util.AddStr("读取最大区块高度数据失败: ", err.Error()), 0)
					continue
				}
				if blockmax.Id == 0 || blockmax.Utime == 0 {
					continue
				}
				if blockmax.Utime+v.WaitTime*60000 < util.Time() { // 如果区块最后更新时间+最大间隔时间<当前时间,则应该发送报警短信
					dxconf := &Dxconf{}
					if _, err := client.Get("dx.block.coin."+v.Symbol, dxconf); err != nil {
						log.Error(util.AddStr("获取最新监听区块[", blockmax.Symbol, "]缓存异常: ", err.Error()), 0)
						continue
					}
					if dxconf.Lastdate == 0 || blockmax.Utime > dxconf.Blockdate { // 没有找到缓存配置或区块最后更新时间>已有时间,首次发送
						dxconf = &Dxconf{
							Prefix:    "dx.block.coin.",
							Symbol:    blockmax.Symbol,
							Trysend:   0,
							Height:    blockmax.Height,
							Lastdate:  util.Time(),
							Blockdate: blockmax.Utime,
						}
					}
					if dxconf.Trysend >= v.MaxSend {
						continue
					}
					if dxconf.Lastdate+(dxconf.Trysend)*v.SendTime*60000 <= util.Time() {
						dxconf.Trysend = dxconf.Trysend + 1
						dxconf.Lastdate = util.Time()
						if err := client.Put(util.AddStr("dx.block.coin.", v.Symbol), dxconf); err != nil {
							log.Error(util.AddStr("最新监听区块[", v.Symbol, "]写入缓存异常: ", err.Error()), 0)
						}
						s := util.Time2Str(util.Time())
						str := strings.Split(s, " ")
						content := fmt.Sprintf("【openwallet】%s服务器(%s)(%s)发生异常(%d)，请运维人员尽快处理。%s", v.Server, v.Ipport, v.Memo, blockmax.Height, str[0]+"/"+str[1])
						s1, _ := util.ObjectToJson(&blockmax)
						fmt.Println("获取币种最大高度: ", s1)
						s2, _ := util.ObjectToJson(dxconf)
						fmt.Println("获取币种短信对象: ", s2)
						// 短信内容
						fmt.Println("发送短信通知内容: ", content)
						ms := strings.Split(v.Mobile, ",")
						for _, mb := range ms {
							toSendSMS("86", mb, content)
							log.Debug("发送报警短信 -> ", 0, log.String("mobile", mb), log.String("content", content), log.String("sendtime", util.Time2Str(dxconf.Lastdate)))
						}
					}
				}
			}
		}
	}()

}

// 发送短信
func toSendSMS(area, mobile, content string) error {
	url := "http://112.74.66.174:6888/sms.aspx?action=send&userid=702&account=mzrhy&password=localbit@2018&mobile=" + mobile + "&content=" + content + "&sendTime=&extno="
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(""))
	if err != nil {
		return util.Error("短信服务异常: ", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result := string(body)
	if strings.Contains(result, "<returnstatus>Success</returnstatus>") {
		return nil
	}
	return util.Error("发送失败: ", area, mobile, " --- ", result)
}

// 获取最新区块高度
func getMaxBlock(coin string, block *model.OwBlock) error {
	client, err := new(cache.RedisManager).Client()
	if err != nil {
		return err
	}
	if _, err := client.Get(util.AddStr("tx.block.coin.", coin), block); err != nil {
		return util.Error("最新区块[", coin, "]读取缓存异常: ", err.Error())
	}
	if block.Id == 0 {
		mongo, err := new(sqld.MGOManager).Get()
		if err != nil {
			return err
		}
		defer mongo.Close()
		if err := mongo.FindOne(sqlc.M(model.OwBlock{}).Eq("symbol", coin).Eq("fork", "false").Eq("state", 1).Orderby("_id", sqlc.DESC_).Limit(1, 1), block); err != nil {
			return err
		}
		if block.Id == 0 {
			return nil
		}
		if err := client.Put(util.AddStr("tx.block.coin.", coin), block); err != nil {
			return util.Error("最新区块[", coin, "]写入缓存异常: ", err.Error())
		}
	}
	return nil
}

func LoadConsulSymbol(symbol string) error {
	if _, b := Mq_consuls[symbol]; !b {
		key := "dc/consul_" + strings.ToLower(symbol)
		consulx := []consul.ConsulConfig{
			{Host: consul.DefaultHost, Node: key},
		}
		if _, err := new(consul.ConsulManager).InitConfig(consulx...); err != nil {
			log.Error("读取币种consul配置失败", 0, log.String("symbol", key), log.AddError(err))
		}
		Mq_consuls[symbol] = symbol
	}
	return nil
}

// 无限等待
func ForeverWait(msg string) error {
	c := make(chan bool)
	go func() {
		fmt.Println("启动信息: ", msg)
	}()
	<-c
	return nil
}

// 保存redis指定值
func CheckAndSaveCacheV(c *cache.RedisManager, key string, val interface{}) error {
	if b, err := c.Get(key, nil); err != nil {
		return err
	} else if !b {
		if err := c.Put(key, val); err != nil {
			return err
		}
	}
	return nil
}
