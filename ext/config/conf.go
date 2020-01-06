package config

import (
	// "git.dustess.com/mk-base/es-driver/es"
	"t-mk-opentrace/ext/constant"
	"t-mk-opentrace/ext/es-driver/es"
	mongo2 "t-mk-opentrace/ext/mongo-driver/mongo"

	// "git.dustess.com/mk-base/redis-driver/redis"
	"github.com/stevenroose/gonfig"
)

// MongoAuth mongodb认证信息结构
type MongoAuth struct {
	Username    string `id:"username" default:"foo"` // username
	Password    string `id:"password" default:"bar"` // password
	AuthSource  string `id:"source" default:"admin"` // auth source
	PasswordSet bool   `id:"on" default:"true"`      // is auth on
}

// GranType 授权类型
type GranType struct {
	AuthorizationCode string `id:"authorizationCode" default:"authorization_code"`
	ClientCredential  string `id:"clientCredential" default:"client_credential"`
	RefreshToken      string `id:"refreshToken" default:"refreshToken"`
}

// Config 配置结构
type Config struct {
	ConfigFile string `short:"c" default:"config.json"`

	Server struct {
		Host string `id:"host" default:"" desc:"listen addr"`                 // 监听地址，配置为空字符串表示监听0.0.0.0
		Port string `id:"port" default:"5000" desc:"http server listen port"` // 启动端口
	} `id:"server" desc:"server config"`

	RPC struct {
		Host string `id:"host" default:"" desc:"listen rpc addr"`              // rpc监听地址，配置为空字符串表示监听0.0.0.0
		Port string `id:"port" default:"50000" desc:"http server listen port"` // 启动端口
	} `id:"rpc" desc:"rpc config"`

	OpenTrace struct {
		Host string `id:"host" default:"localhost"`
		Port string `id:"port" default:"6831"`
	}
	RPCService struct {
		MKPmPlanSVC string `id:"mk-pm-plan-svc" default:"" desc:"推广任务rpc服务"`
	} `id:"rpc-service" desc:"rpc service"`
	HTTPService struct {
		MKJobSrv string `id:"mk-job-srv" default:"" desc:"定时任务服务"`
	} `id:"http-service" desc:"http service"`
	Mongo struct {
		MongoMK struct {
			Addr       string    `id:"addr" default:"mongodb://127.0.0.1:27017"` // 地址
			Auth       MongoAuth `id:"auth" desc:"mongo auth config"`            // 认证信息
			MKManageDB struct {
				Name string `id:"name" default:"mk_manage" desc:"营销运营管理系统数据库名"`
				Coll struct {
					User          string `id:"user" default:"user"`
					SMSCodeRecord string `id:"sms_code_record" default:"sms_code_record"`
				} `id:"coll" desc:"营销运营管理系统系统用户"`
			} `id:"mk_manage" desc:"营销运营管理系统数据库"`
			MKDB struct {
				Name string `id:"name" default:"mk" desc:"markting db name"`
				Coll struct {
					MaterialConfig                                string `id:"material_config" default:"material_config"`
					MaterialForm                                  string `id:"material_form" default:"material_form"`
					MaterialFormFeedback                          string `id:"material_form_feedback" default:"material_form_feedback"`
					MaterialFormFeedbackExportTask                string `id:"material_form_feedbackexport_task" default:"material_form_feedbackexport_task"`
					MaterialFormMeta                              string `id:"material_form_meta" default:"material_form_meta"`
					MaterialFormComponentGroup                    string `id:"material_form_component_group" default:"material_form_component_group"`
					MaterialArticle                               string `id:"material_article" default:"material_article"`
					MaterialCircleFriends                         string `id:"material_cof" default:"material_cof"`
					MaterialH5                                    string `id:"material_h5" default:"material_h5"`
					MaterialActivityH5                            string `id:"material_activity_h5" default:"material_activity_h5"`
					MarketingTimingPlan                           string `id:"timing_plan" default:"timing_plan"`
					MarketingTimingPlanMeta                       string `id:"timing_plan_meta" default:"timing_plan_meta"`
					MarketingReleaseTask                          string `id:"timing_plan_task" default:"timing_plan_task"`
					MarketingReleasePlan                          string `id:"timing_plan_bridge_user" default:"timing_plan_bridge_user"`
					BridgeMaterial                                string `id:"material_bridge" default:"material_bridge"`
					PlatFormUser                                  string `id:"wp_platform_user" default:"wp_platform_user"`
					ACLTree                                       string `id:"acl_tree" default:"acl_tree"`
					ACLRole                                       string `id:"acl_role" default:"acl_role"`
					Tag                                           string `id:"tag" default:"tag"`
					TagMaterialBridge                             string `id:"tag_material_bridge" default:"tag_material_bridge"`
					TagWXCustBridge                               string `id:"tag_wxcust_bridge" default:"tag_wxcust_bridge"`
					TagWXClueBridge                               string `id:"tag_activity_wxclue_bridge" default:"tag_activity_wxclue_bridge"`
					TagWXClueAttributeBridge                      string `id:"tag_attribute_wxclue_bridge" default:"tag_attribute_wxclue_bridge"`
					WXBehave                                      string `id:"wx_clue_activity" default:"wx_clue_activity"`
					WXPlanClue                                    string `id:"wx_plan_clue" default:"wx_plan_clue"`
					WXClue                                        string `id:"wx_clue" default:"wx_clue"`
					WXClueAll                                     string `id:"wx_all_clue" default:"wx_all_clue"`
					WXClueTrace                                   string `id:"wx_clue_trace" default:"wx_clue_trace"`
					WXClueActionInfo                              string `id:"wx_clue_action_info" default:"wx_clue_action_info"`
					WXSaleActionInfo                              string `id:"wx_sale_action_info" default:"wx_sale_action_info"`
					WXClueContact                                 string `id:"wx_clue_contact" default:"wx_clue_contact"`
					WXUser                                        string `id:"wx_user" default:"wx_user"`
					WXScene                                       string `id:"wx_scene" default:"wx_scene"`
					WXEvent                                       string `id:"wx_clue_activity_config" default:"wx_clue_activity_config"`
					SearchMeta                                    string `id:"search_meta" default:"search_meta"`
					SearchConfig                                  string `id:"search_config" default:"search_config"`
					Area                                          string `id:"prov_city" default:"prov_city"`
					TaskNewClueStatistics                         string `id:"statistics_task_clue" default:"statistics_task_clue"`
					NewCluePV                                     string `id:"statistics_task_clue_pv" default:"statistics_task_clue_pv"`
					NewClueUV                                     string `id:"statistics_task_clue_uv" default:"statistics_task_clue_uv"`
					NewClueStatistics                             string `id:"statistics_clue" default:"statistics_clue"`
					ActionTagStatistics                           string `id:"statistics_action_tag" default:"statistics_action_tag"`
					PlanClueReport                                string `id:"timing_plan_clue_day_report" default:"timing_plan_clue_day_report"`
					PlanTagReport                                 string `id:"timing_plan_tag_day_report" default:"timing_plan_tag_day_report"`
					PlanPVReport                                  string `id:"timing_plan_pv_day_report" default:"timing_plan_pv_day_report"`
					PlanUVReport                                  string `id:"timing_plan_uv_day_report" default:"timing_plan_uv_day_report"`
					UserClueReport                                string `id:"user_clue_day_report" default:"user_clue_day_report"`
					UserPVReport                                  string `id:"user_pv_day_report" default:"user_pv_day_report"`
					MaterialPVReport                              string `id:"material_pv_day_report" default:"material_pv_day_report"`
					MaterialUVReport                              string `id:"material_uv_day_report" default:"material_uv_day_report"`
					CompanyClueReport                             string `id:"company_clue_day_report" default:"company_clue_day_report"`
					CompanyPVReport                               string `id:"company_pv_day_report" default:"company_pv_day_report"`
					ClueLivenessReport                            string `id:"clue_liveness_day_report" default:"clue_liveness_day_report"`
					WXComponent                                   string `id:"wx_component" default:"wx_component"`
					WXMP                                          string `id:"wx_mp" default:"wx_mp"`
					WXOPEN                                        string `id:"wx_open" default:"wx_open"`
					WXOPENMP                                      string `id:"wx_open_mp" default:"wx_open_mp"`
					WXAuthoredMP                                  string `id:"wx_authored_mp" default:"wx_authored_mp"`
					WXAuthoredMini                                string `id:"wx_authored_mini" default:"wx_authored_mini"`
					WxMpFunc                                      string `id:"wx_mp_func" default:"wx_mp_func"`
					WxMiniFunc                                    string `id:"wx_mini_func" default:"wx_mini_func"`
					SmsCodeRecord                                 string `id:"sms_code_record" default:"sms_code_record"`
					MaterialSourcePVReport                        string `id:"material_source_pv_report" default:"material_source_pv_report"`
					MaterialSourceUVReport                        string `id:"material_source_uv_report" default:"material_source_uv_report"`
					MaterialShareCountReport                      string `id:"material_share_count_report" default:"material_share_count_report"`
					SaleMsTotalUVReport                           string `id:"sale_ms_total_uv_report" default:"sale_ms_total_uv_report"`
					MaterialSourceTotalUVReport                   string `id:"material_source_total_uv_report" default:"material_source_total_uv_report"`
					TimingPlanMaterialReportSourceUVTotal         string `id:"timing_plan_ms_total_uv_report" default:"timing_plan_ms_total_uv_report"`
					TimingPlanTaskMaterialReportSourceUVTotal     string `id:"timing_plan_task_ms_total_uv_report" default:"timing_plan_task_ms_total_uv_report"`
					TimingPlanMReportSourceUVTotal                string `id:"timing_plan_material_total_uv_report" default:"timing_plan_material_total_uv_report"`
					TimingPlanMSaleReportSourceUVTotal            string `id:"timing_plan_material_sale_total_uv_report" default:"timing_plan_material_sale_total_uv_report"`
					TimingPlanMaterialReportSourceUV              string `id:"timing_plan_ms_uv_report" default:"timing_plan_ms_uv_report"`
					TimingPlanMReportSourceUV                     string `id:"timing_plan_material_uv_report" default:"timing_plan_material_uv_report"`
					TimingPlanMSaleReportSourceUV                 string `id:"timing_plan_material_sale_uv_report" default:"timing_plan_material_sale_uv_report"`
					TimingPlanTaskMaterialReportSourceUV          string `id:"timing_plan_task_ms_uv_report" default:"timing_plan_task_ms_uv_report"`
					TimingPlanSaleMaterialReportSourceUVTotal     string `id:"timing_plan_sale_ms_total_uv_report" default:"timing_plan_sale_ms_total_uv_report"`
					TimingPlanTaskSaleMaterialReportSourceUVTotal string `id:"timing_plan_task_sale_ms_total_uv_report" default:"timing_plan_task_sale_ms_total_uv_report"`
					TimingPlanSaleMaterialReportSourceUV          string `id:"timing_plan_sale_ms_uv_report" default:"timing_plan_sale_ms_uv_report"`
					TimingPlanTaskSaleMaterialReportSourceUV      string `id:"timing_plan_task_sale_ms_uv_report" default:"timing_plan_task_sale_ms_uv_report"`
					SalePlanMaterialReportSourceUV                string `id:"sale_ms_uv_report" default:"sale_ms_uv_report"`
					PromotionPlanWhole                            string `id:"pm_plan_whole" default:"pm_plan_whole"`
					PromotionPlanPart                             string `id:"pm_plan_part" default:"pm_plan_part"`
					PromotionPlanNewTask                          string `id:"pm_plan_new_task" default:"pm_plan_new_task"`
					PmPVReport                                    string `id:"pm_pv_report" default:"pm_pv_report"`
					PmUVReport                                    string `id:"pm_uv_report" default:"pm_uv_report"`
					PmCompanyIncomeReport                         string `id:"pm_company_income_report" default:"pm_company_income_report"`
					PmCompanyIncomeDetailReport                   string `id:"pm_company_income_detail_report" default:"pm_company_income_detail_report"`
					PmQRCodeIdentReport                           string `id:"pm_qr_code_ident_report" default:"pm_qr_code_ident_report"`
					PmSubmitReport                                string `id:"pm_submit_report" default:"pm_submit_report"`
					PmValidEventReport                            string `id:"pm_valid_event_report" default:"pm_valid_event_report"`
					PmBehaveEvent                                 string `id:"pm_behave_event" default:"pm_behave_event"`
					PmFeedback                                    string `id:"pm_feedback" default:"pm_feedback"`
					PmQRCodeIdent                                 string `id:"pm_qr_code_ident" default:"pm_qr_code_ident"`
				} `id:"coll" desc:"markting db colls config"`
			} `id:"mk" desc:"markting db config"`
			WPBill struct {
				Name string `id:"name" default:"wpbill" desc:"wpbill db name"`
				Coll struct {
					PlatformUser    string `id:"wp_platform_user" default:"wp_platform_user"`
					PlatformAccount string `id:"wp_platform_account" default:"wp_platform_account"`
				} `id:"coll" desc:"wpbill db colls config"`
			} `id:"wpbill" desc:"wpbill db config"`
			WPDATA20190612 struct {
				Name string `id:"name" default:"wp_data_20190612" desc:"wp_data_20190612 db name"`
			} `id:"wp_data_20190612" desc:"wp_data_20190612 db config"`
		} `id:"mk_biz" desc:"mk mongo config"`
		MongoMKWat struct {
			Addr string    `id:"addr" default:"mongodb://127.0.0.1:27017"` // 地址
			Auth MongoAuth `id:"auth" desc:"mongo auth config"`            // 认证信息
		} `id:"mk_wat" desc:"mk_wat mongo config"`
		MongoWP struct {
			Addr   string    `id:"addr" default:"mongodb://127.0.0.1:27017"` // 地址
			Auth   MongoAuth `id:"auth" desc:"mongo auth config"`            // 认证信息
			WPBill struct {
				Name string `id:"name" default:"wpbill" desc:"wpbill db name"`
				Coll struct {
					PlatformUser    string `id:"wp_platform_user" default:"wp_platform_user"`
					PlatformAccount string `id:"wp_platform_account" default:"wp_platform_account"`
				} `id:"coll" desc:"wpbill db colls config"`
			} `id:"wpbill" desc:"wpbill db config"`
			WPDATA20190612 struct {
				Name string `id:"name" default:"wp_data_20190612" desc:"wp_data_20190612 db name"`
			} `id:"wp_data_20190612" desc:"wp_data_20190612 db config"`
			WP struct {
				Name string `id:"name" default:"wp" desc:"工作手机 db name"`
				Coll struct {
					WpUserDevice string `id:"wp_user_device" default:"wp_user_device"`
					WpWxUser     string `id:"wp_wx_user" default:"wp_wx_user"`
					WpWxContact  string `id:"_wp_wx_contact" default:"_wp_wx_contact"`
					WpWxChat     string `id:"_wp_wx_chat" default:"_wp_wx_chat"`
					WpWxMessage  string `id:"_wp_wx_message" default:"_wp_wx_message"`
					WpDepartment string `id:"wp_department" default:"wp_department"`
					WpRole       string `id:"wp_role" default:"wp_role"`
				} `id:"coll" desc:"wp db colls config"`
			} `id:"wp" desc:"wp工作手机 db config"`
		} `id:"wp" desc:"wp mongo config"`
	} `id:"mongo" desc:"mongo cluster config"`

	Redis struct {
		Addr     string `id:"addr" default:"127.0.0.1"` // redis addr，例如 127.0.0.1:6379
		Password string `id:"password" default:"6379"`  // redis password
		PoolSize int    `id:"poolsize" default:"10"`    // redis poolsize
		CacheDB  int    `id:"cachedb" default:"6"`      // redis cachedb
	} `id:"redis" desc:"redis config"`

	ES struct {
		Addr     string `id:"addr" default:"http://127.0.0.1:9200"` // redis addr，例如 127.0.0.1:6379
		Username string `id:"username" default:""`                  // es basic auth username
		Password string `id:"password" default:""`                  // es basic auth password
	} `id:"es" desc:"es config"`

	Kafka struct {
		Addrs []string `id:"addrs" default:"[127.0.0.1:9092]"` // kafka addrs
		Group struct {
			MKWatWXLoad string `id:"mk_wat_wx_load" desc:"kafka consumer group mk_wat_wx_load: load data to mongo"`
			MKPmFDBLoad string `id:"mk_pm_fdb_load" desc:"kafka consumer group mk_pm_fdb_load: load data to mongo"`
		} `id:"group" desc:"kafka consumer group config"`
		Topic struct {
			MKPmEvent struct {
				Name      string `id:"name" default:"mk_pm_event"` // mk_promotion_feedback topic name
				Partition int    `id:"partition" default:"8"`      // mk_promotion_feedback partition
			} `id:"mk_pm_event" desc:"kafka pm feedback topic name"` // wat topic
			MKPmEventReport struct {
				Name      string `id:"name" default:"mk_pm_event_report"` // mk_promotion_feedback report topic name
				Partition int    `id:"partition" default:"8"`             // mk_promotion_feedback partition
			} `id:"mk_pm_event_report" desc:"kafka pm feedback topic name"` // mk_promotion_feedback topic
			MKPmEventReportRetry struct {
				Name      string `id:"name" default:"mk_pm_event_report_retry"` // mk_promotion_feedback report for retry topic name
				Partition int    `id:"partition" default:"8"`                   // mk_promotion_feedback partition
			} `id:"mk_pm_event_report_retry" desc:"kafka pm feedback topic name for retry"` // mk_promotion_feedback err topic for retry
			MKPlanReport struct {
				Name      string `id:"name" default:"mk_plan_report"` // mk_wat topic name
				Partition int    `id:"partition" default:"8"`         // mk_wat partition
			} `id:"mk_plan_report" desc:"kafka wat topic name"` // wat topic
			MKWorkbenchReport struct {
				Name      string `id:"name" default:"mk_workbench_report"` // mk_plan_report topic name
				Partition int    `id:"partition" default:"8"`              // mk_plan_report partition
			} `id:"mk_workbench_report" desc:"kafka mk_workbench_report topic"` // mk_plan_report topic
			MKClueTrace struct {
				Name      string `id:"name" default:"mk_clue_trace"` // mk_clue_trace topic name
				Partition int    `id:"partition" default:"8"`        // mk_clue_trace partition
			} `id:"mk_clue_trace" desc:"kafka mk_clue_trace topic"` // mk_clue_trace topic
		} `id:"topic" desc:"kafka topic config"`
	} `id:"kafka" desc:"kafka config"`

	AliOss struct {
		OssKey     string `id:"ossKey" default:""`     // AliOss key
		OssSecret  string `id:"ossSecret" default:""`  // AliOss secret
		OssRoleAcs string `id:"ossRoleAcs" default:""` // AliOss 角色
		Bucket     string `id:"bucket" default:""`     // AliOss 存储空间名称
		EndPoint   string `id:"endPoint" default:""`   // 地域节点
	} `id:"aliOss" desc:"ali oss config"`

	WX struct {
		AuthRedirectURI    string `id:"authRedirectURI" default:"http://mk-wx-dev.dustess.com/wxopen/v1/redirect"`
		CreateOpenPlatform string `id:"createOpenPlatform" default:"https://api.weixin.qq.com/cgi-bin/open/create"`
		BindOpenPlatform   string `id:"bindOpenPlatform" default:"https://api.weixin.qq.com/cgi-bin/open/bind"`
		UnBindOpenPlatform string `id:"unBindOpenPlatform" default:"https://api.weixin.qq.com/cgi-bin/open/unbind"`
		GetOpenPlatform    string `id:"getOpenPlatform" default:"https://api.weixin.qq.com/cgi-bin/open/get"`
		Token              string `id:"token" default:""`
		AESKey             string `id:"encodingAESKey" default:""`
		AppID              string `id:"appId" default:""`
		AppSecret          string `id:"appSecret" default:""`
	} `id:"wx" desc:"wx config"`

	MiniProgram struct {
		Secret             string   `id:"secret" default:"ca0e2cd4b04bb5d61e204088de223edd"`
		AppID              string   `id:"appId" default:"LTAIrzZVfpOUN68L"`
		GranType           GranType `id:"granType" desc:"weixin auth type"`
		AuthURL            string   `id:"authUrl" default:"https://api.weixin.qq.com/sns/jscode2session"`
		TokenURL           string   `id:"tokenUrl" default:"https://api.weixin.qq.com/cgi-bin/token"`
		QRCodeURLUnlimited string   `id:"qrCodeUrlUnlimited" default:"https://api.weixin.qq.com/wxa/getwxacodeunlimit"`
		QRCodeURL          string   `id:"qrCodeUrl" default:"https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode"`
	} `id:"miniProgram" desc:"小程序参数结构体"`

	Wechat struct {
		Secret           string   `id:"secret" default:"919ec841a599afd10f9263282f52183f"`
		AppID            string   `id:"appId" default:"wx965ef62990b09099"`
		GranType         GranType `id:"granType" desc:"weixin auth type"`
		AuthURL          string   `id:"authUrl" default:"https://open.weixin.qq.com/connect/oauth2/authorize"`
		AccessToken      string   `id:"accessToken" default:"https://api.weixin.qq.com/sns/oauth2/component/access_token"`
		CheckAccessToken string   `id:"checkAccessToken" default:"https://api.weixin.qq.com/sns/auth"`
		GetTicketURL     string   `id:"getTicketUrl" default:"https://api.weixin.qq.com/cgi-bin/ticket/getticket"`
		UserInfo         string   `id:"userInfo" default:"https://api.weixin.qq.com/sns/userinfo"`
		RefreshToken     string   `id:"refreshToken" default:"https://api.weixin.qq.com/sns/oauth2/refresh_token"`
		RedirectURL      string   `id:"redirectUrl" default:"https://mk-wx-dev.dustess.com/wxopen/v1/wxpublic/authback"`
		Token            string   `id:"token" default:"asdfghjkl123456789" desc:"公众号接入验证token，填写在服务器配置上的参数，自定义"`
	} `id:"wechat" desc:"公众号参数结构体"`
	Mode constant.RunTimeMode `id:"mode" default:"debug"` // 运行模式 debug|test|release
}

// Init 初始化
func (config *Config) Init() error {
	return gonfig.Load(config, gonfig.Conf{
		ConfigFileVariable:  constant.ConfigFileVariable, // enables passing --configfile myfile.conf
		FileDefaultFilename: constant.ConfigFileDefaultName,
		FileDecoder:         gonfig.DecoderJSON,
		EnvPrefix:           constant.ConfigEnvPrefix,
	})
}

// // ToCacheConfig 转为redis cache配置
// func (config *Config) ToCacheConfig() *redis.Config {
// 	return &redis.Config{
// 		ClientName: redis.MKCache,
// 		Addr:       config.Redis.Addr,
// 		Password:   config.Redis.Password,
// 		DB:         config.Redis.CacheDB,
// 		PoolSize:   config.Redis.PoolSize,
// 	}
// }

// // ToSessionConfig 转为redis session配置
// func (config *Config) ToSessionConfig() *redis.Config {
// 	return &redis.Config{
// 		ClientName: redis.MKSession,
// 		Addr:       config.Redis.Addr,
// 		Password:   config.Redis.Password,
// 		DB:         0,
// 		PoolSize:   config.Redis.PoolSize,
// 	}
// }

// ToMongoMKBizConfig mk-biz 集群配置
func (config *Config) ToMongoMKBizConfig() *mongo2.Config {
	return &mongo2.Config{
		ClientName: mongo2.MKBiz,
		Addr:       config.Mongo.MongoMK.Addr,
		Auth: mongo2.Auth{
			Username:    config.Mongo.MongoMK.Auth.Username,
			Password:    config.Mongo.MongoMK.Auth.Password,
			AuthSource:  config.Mongo.MongoMK.Auth.AuthSource,
			PasswordSet: config.Mongo.MongoMK.Auth.PasswordSet,
		},
	}
}

// ToMongoMKWatConfig mk-wat 集群配置
func (config *Config) ToMongoMKWatConfig() *mongo2.Config {
	return &mongo2.Config{
		ClientName: mongo2.MKWat,
		Addr:       config.Mongo.MongoMKWat.Addr,
		Auth: mongo2.Auth{
			Username:    config.Mongo.MongoMKWat.Auth.Username,
			Password:    config.Mongo.MongoMKWat.Auth.Password,
			AuthSource:  config.Mongo.MongoMKWat.Auth.AuthSource,
			PasswordSet: config.Mongo.MongoMKWat.Auth.PasswordSet,
		},
	}
}

// ToMongoWPConfig wp 集群配置
func (config *Config) ToMongoWPConfig() *mongo2.Config {
	return &mongo2.Config{
		ClientName: mongo2.WP,
		Addr:       config.Mongo.MongoWP.Addr,
		Auth: mongo2.Auth{
			Username:    config.Mongo.MongoWP.Auth.Username,
			Password:    config.Mongo.MongoWP.Auth.Password,
			AuthSource:  config.Mongo.MongoWP.Auth.AuthSource,
			PasswordSet: config.Mongo.MongoWP.Auth.PasswordSet,
		},
	}
}

// ToESConfig 转换为es配置
func (config *Config) ToESConfig() *es.Config {
	return &es.Config{
		Addr:     config.ES.Addr,
		Username: config.ES.Username,
		Password: config.ES.Password,
	}
}
