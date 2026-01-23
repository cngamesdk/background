package cron_task

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	cron_task2 "github.com/cngamesdk/go-core/model/sql/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderCronTaskConfig = system.InitOrderMyCustom + 1

type initCronTaskConfigMysql struct{}

// auto run
func init() {
	system.RegisterInit(initOrderCronTaskConfig, &initCronTaskConfigMysql{})
}

func (i *initCronTaskConfigMysql) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&cron_task.DimCronTaskConfigModel{})
}

func (i *initCronTaskConfigMysql) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&cron_task.DimCronTaskConfigModel{})
}

func (i *initCronTaskConfigMysql) InitializerName() string {
	return (&cron_task.DimCronTaskConfigModel{}).TableName()
}

func (i *initCronTaskConfigMysql) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []cron_task2.DimCronTaskConfigModel{
		{Name: "天用户登录日志表（按子）", Spec: "* */5 * * * *", Remark: "从登录日志清洗用户每天唯一记录", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `INSERT IGNORE INTO dwd_day_game_reg_uid_login_log (
	platform_id,
	game_id,
	user_id,
	login_date,
	first_login_time,
	last_login_time,
	login_count,
	reg_time,
	agent_id,
	site_id,
	media_site_id,
	idfv,
	imei,
	oaid,
	andriod_id,
	system_version,
	app_version_code,
	sdk_version_code,
	network,
	client_ip,
	ipv4,
	ipv6,
	channel_id,
	model,
	brand,
	user_agent,
	unique_device 
) SELECT
login.platform_id,
login.game_id,
login.user_id,
DATE ( login.login_time ) AS lgin_date,
login.login_time AS first_login_time,
login.login_time AS last_login_timelast_login_time,
1 AS login_count,
reg.reg_time AS reg_time,
login.agent_id,
login.site_id,
login.media_site_id,
login.idfv,
login.imei,
login.oaid,
login.andriod_id,
login.system_version,
login.app_version_code,
login.sdk_version_code,
login.network,
login.client_ip,
login.ipv4,
login.ipv6,
login.channel_id,
login.model,
login.brand,
login.user_agent,
login.unique_device 
FROM
	(
	SELECT
		login.* 
	FROM
		ods_login_logs AS login
		JOIN ( SELECT MIN( id ) AS min_id FROM ods_login_logs WHERE created_at BETWEEN '{{StartDateTime}}' AND '{{EndDateTime}}' GROUP BY platform_id, game_id, user_id ) AS temp ON temp.min_id = login.id 
	) AS login
	LEFT JOIN dwd_game_reg_log AS reg ON login.platform_id = reg.platform_id 
	AND login.game_id = reg.game_id 
	AND login.user_id = reg.user_id;`},
		{Name: "天用户登录日志表（按根）", Spec: "* */5 * * * *", Remark: "从用户天登录日志表(按子)中清洗到按根用户天登录日志表", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `REPLACE INTO dwd_day_root_game_reg_uid_login_log (
	platform_id,
	root_game_id,
	user_id,
	login_date,
	game_id,
	first_login_time,
	last_login_time,
	login_count,
	reg_time,
	agent_id,
	site_id,
	media_site_id,
	idfv,
	imei,
	oaid,
	andriod_id,
	system_version,
	app_version_code,
	sdk_version_code,
	network,
	client_ip,
	ipv4,
	ipv6,
	channel_id,
	model,
	brand,
	user_agent,
	unique_device 
) SELECT
IFNULL( before_login.platform_id, login.platform_id ) AS platform_id,
IFNULL( before_login.root_game_id, login.root_game_id ) AS root_game_id,
IFNULL( before_login.user_id, login.user_id ) AS user_id,
IFNULL( before_login.login_date, login.login_date ) AS lgin_date,
IFNULL( before_login.game_id, login.game_id ) AS game_id,
IFNULL( before_login.first_login_time, login.first_login_time_2 ) AS first_login_time,
IFNULL( before_login.last_login_time, login.last_login_time_2 ) AS last_login_time,
IFNULL( before_login.login_count, 0 ) + login.login_count_2 AS login_count,
IFNULL( reg.reg_time, login.first_login_time ) AS reg_time,
IFNULL( before_login.agent_id, login.agent_id ) AS agent_id,
IFNULL( before_login.site_id, login.site_id ) AS site_id,
IFNULL( before_login.media_site_id, login.media_site_id ) AS media_site_id,
IFNULL( before_login.idfv, login.idfv ) AS idfv,
IFNULL( before_login.imei, login.imei ) AS imei,
IFNULL( before_login.oaid, login.oaid ) AS oaid,
IFNULL( before_login.andriod_id, login.andriod_id ) AS andriod_id,
IFNULL( before_login.system_version, login.system_version ) AS system_version,
IFNULL( before_login.app_version_code, login.app_version_code ) AS app_version_code,
IFNULL( before_login.sdk_version_code, login.sdk_version_code ) AS sdk_version_code,
IFNULL( before_login.network, login.network ) AS network,
IFNULL( before_login.client_ip, login.client_ip ) AS client_ip,
IFNULL( before_login.ipv4, login.ipv4 ) AS ipv4,
IFNULL( before_login.ipv6, login.ipv6 ) AS ipv6,
IFNULL( before_login.channel_id, login.channel_id ) AS channel_id,
IFNULL( before_login.model, login.model ) AS model,
IFNULL( before_login.brand, login.brand ) AS brand,
IFNULL( before_login.user_agent, login.user_agent ) AS user_agent,
IFNULL( before_login.unique_device, login.unique_device ) AS unique_device 
FROM
	(
	SELECT
		login.*,
		tmp.root_game_id,
		tmp.login_date AS login_date_2,
		tmp.login_count AS login_count_2,
		tmp.first_login_time AS first_login_time_2,
		tmp.last_login_time AS last_login_time_2 
	FROM
		dwd_day_game_reg_uid_login_log AS login
		JOIN (
		SELECT
			MIN( login.id ) AS min_id,
			login.platform_id,
			root_game.id AS root_game_id,
			login.user_id,
			login.login_date AS login_date,
			count( 1 ) AS login_count,
			MIN( login.first_login_time ) AS first_login_time,
			MAX( login.last_login_time ) AS last_login_time 
		FROM
			dwd_day_game_reg_uid_login_log AS login
			JOIN dim_game AS game ON login.platform_id = game.platform_id 
			AND login.game_id = game.id
			JOIN dim_main_game AS main_game ON main_game.platform_id = game.platform_id 
			AND main_game.id = game.main_game_id
			JOIN dim_root_game AS root_game ON root_game.platform_id = main_game.platform_id 
			AND root_game.id = main_game.root_game_id 
		WHERE
			login.updated_at BETWEEN '{{StartDateTime}}' 
			AND '{{EndDateTime}}' 
		GROUP BY
			login.platform_id,
			root_game.id,
			login.user_id,
			login.login_date 
		) AS tmp ON login.id = tmp.min_id 
	WHERE
		login.updated_at BETWEEN '{{StartDateTime}}' 
		AND '{{EndDateTime}}' 
	) AS login
	LEFT JOIN dwd_root_game_reg_log AS reg ON login.platform_id = reg.platform_id 
	AND login.root_game_id = reg.root_game_id 
	AND login.user_id = reg.user_id
	LEFT JOIN dwd_day_root_game_reg_uid_login_log AS before_login ON login.platform_id = before_login.platform_id 
	AND login.root_game_id = before_login.root_game_id 
	AND login.user_id = before_login.user_id 
	AND login.login_date = before_login.login_date 
WHERE
	login.updated_at BETWEEN '{{StartDateTime}}' 
	AND '{{EndDateTime}}';`},
		{Name: "子注册清洗", Spec: "* */5 * * * *", Remark: "从天登录日志清洗到子注册表", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `INSERT INTO dwd_game_reg_log (
	platform_id,
	game_id,
	user_id,
	reg_time,
	agent_id,
	site_id,
	media_site_id,
	idfv,
	imei,
	oaid,
	andriod_id,
	system_version,
	app_version_code,
	sdk_version_code,
	network,
	client_ip,
	ipv4,
	ipv6,
	channel_id,
	model,
	brand,
	user_agent 
) SELECT
login.platform_id,
login.game_id,
login.user_id,
login.login_time,
login.agent_id,
login.site_id,
login.media_site_id,
login.idfv,
login.imei,
login.oaid,
login.andriod_id,
login.system_version,
login.app_version_code,
login.sdk_version_code,
login.network,
login.client_ip,
login.ipv4,
login.ipv6,
login.channel_id,
login.model,
login.brand,
login.user_agent 
FROM
	(
	SELECT
		login.* 
	FROM
		ods_login_logs AS login
		JOIN ( SELECT platform_id, game_id, user_id, min( id ) AS min_id FROM ods_login_logs WHERE updated_at BETWEEN '{{StartDateTime}}' AND '{{EndDateTime}}' GROUP BY platform_id, game_id, user_id ) AS user_min_log ON login.id = user_min_log.min_id 
	) AS login
	LEFT JOIN dwd_game_reg_log AS reg ON login.platform_id = reg.platform_id 
	AND login.game_id = reg.game_id 
	AND login.user_id = reg.user_id 
WHERE
	reg.platform_id IS NULL;`},
		{Name: "根注册清洗", Spec: "* */5 * * * *", Remark: "从子注册清洗到根注册表", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `INSERT INTO dwd_root_game_reg_log (
	platform_id,
	root_game_id,
	user_id,
	reg_time,
	game_id,
	agent_id,
	site_id,
	media_site_id,
	idfv,
	imei,
	oaid,
	andriod_id,
	system_version,
	app_version_code,
	sdk_version_code,
	network,
	client_ip,
	ipv4,
	ipv6,
	channel_id,
	model,
	brand,
	user_agent 
) SELECT
game_reg.platform_id,
game_reg.root_game_id,
game_reg.user_id,
game_reg.reg_time,
game_reg.game_id,
game_reg.agent_id,
game_reg.site_id,
game_reg.media_site_id,
game_reg.idfv,
game_reg.imei,
game_reg.oaid,
game_reg.andriod_id,
game_reg.system_version,
game_reg.app_version_code,
game_reg.sdk_version_code,
game_reg.network,
game_reg.client_ip,
game_reg.ipv4,
game_reg.ipv6,
game_reg.channel_id,
game_reg.model,
game_reg.brand,
game_reg.user_agent 
FROM
	(
	SELECT
		reg.*,
		main_game.root_game_id 
	FROM
		dwd_game_reg_log AS reg
		JOIN dim_game AS game ON reg.platform_id = game.platform_id 
		AND reg.game_id = game.id
		JOIN dim_main_game AS main_game ON game.platform_id = main_game.platform_id 
		AND game.main_id = main_game.id
		JOIN dim_root_game AS root_game ON main_game.platform_id = root_game.platform_id 
		AND main_game.root_game_id = root_game.id 
	WHERE
		reg.updated_at BETWEEN '{{StartDateTime}}' 
		AND '{{EndDateTime}}' 
	) AS game_reg
	LEFT JOIN dwd_root_game_reg_log AS root_reg ON game_reg.platform_id = root_reg.platform_id 
	AND game_reg.root_game_id = root_reg.root_game_id 
	AND game_reg.user_id = root_reg.user_id 
WHERE
	root_reg.platform_id IS NULL;`},
		{Name: "根注册写入30天回流表", Spec: "* */5 * * * *", Remark: "从根注册清洗到30天回流表", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `INSERT INTO dwd_root_game_back_reg_log (
	platform_id,
	root_game_id,
	user_id,
	reg_time,
	last_time,
	game_id,
	agent_id,
	site_id,
	media_site_id,
	idfv,
	imei,
	oaid,
	andriod_id,
	system_version,
	app_version_code,
	sdk_version_code,
	network,
	client_ip,
	ipv4,
	ipv6,
	channel_id,
	model,
	brand,
	user_agent 
)
SELECT
	root_reg.platform_id,
	root_reg.root_game_id,
	root_reg.user_id,
	root_reg.reg_time,
	CONCAT(date(DATE_ADD(root_reg.reg_time, INTERVAL 30 DAY)), " 23:59:59") AS last_time,
	root_reg.game_id,
	root_reg.agent_id,
	root_reg.site_id,
	root_reg.media_site_id,
	root_reg.idfv,
	root_reg.imei,
	root_reg.oaid,
	root_reg.andriod_id,
	root_reg.system_version,
	root_reg.app_version_code,
	root_reg.sdk_version_code,
	root_reg.network,
	root_reg.client_ip,
	root_reg.ipv4,
	root_reg.ipv6,
	root_reg.channel_id,
	root_reg.model,
	root_reg.brand,
	root_reg.user_agent 
FROM
	dwd_root_game_reg_log AS root_reg
	LEFT JOIN dwd_root_game_back_reg_log AS back_reg ON root_reg.platform_id = back_reg.platform_id 
	AND root_reg.root_game_id = back_reg.root_game_id 
	AND root_reg.user_id = back_reg.user_id
WHERE
	 root_reg.updated_at BETWEEN '{{StartDateTime}}' AND '{{EndDateTime}}' 
	AND back_reg.platform_id IS NULL;`},
		{Name: "30天回流用户写入回流表", Spec: "* */5 * * * *", Remark: "登录日志中流失用户写入30天回流表", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `INSERT INTO dwd_root_game_back_reg_log (
	platform_id,
	root_game_id,
	user_id,
	reg_time,
	last_time,
	game_id,
	agent_id,
	site_id,
	media_site_id,
	idfv,
	imei,
	oaid,
	andriod_id,
	system_version,
	app_version_code,
	sdk_version_code,
	network,
	client_ip,
	ipv4,
	ipv6,
	channel_id,
	model,
	brand,
	user_agent 
) SELECT
login.platform_id,
login.root_game_id,
login.user_id,
login.login_time,
CONCAT(date(DATE_ADD(login.login_time, INTERVAL 30 DAY)), " 23:59:59") AS last_time,
login.game_id,
login.agent_id,
login.site_id,
login.media_site_id,
login.idfv,
login.imei,
login.oaid,
login.andriod_id,
login.system_version,
login.app_version_code,
login.sdk_version_code,
login.network,
login.client_ip,
login.ipv4,
login.ipv6,
login.channel_id,
login.model,
login.brand,
login.user_agent 
FROM
	(
	SELECT
		login.*,
		main_game.root_game_id 
	FROM
		ods_login_logs AS login
		JOIN ( SELECT platform_id, game_id, user_id, min( id ) AS min_id FROM ods_login_logs WHERE updated_at BETWEEN '{{StartDateTime}}' AND '{{EndDateTime}}' GROUP BY platform_id, game_id, user_id ) AS user_min_log ON login.id = user_min_log.min_id
		JOIN dim_game AS game ON user_min_log.platform_id = game.platform_id 
		AND user_min_log.game_id = game.id
		JOIN dim_main_game AS main_game ON game.platform_id = main_game.platform_id 
		AND game.main_id = main_game.id
		JOIN dim_root_game AS root_game ON main_game.platform_id = root_game.platform_id 
		AND main_game.root_game_id = root_game.id 
	) AS login
	JOIN (
	SELECT
		back_reg.* 
	FROM
		dwd_root_game_back_reg_log AS back_reg
		JOIN ( SELECT platform_id, root_game_id, user_id, max( id ) AS max_id FROM dwd_root_game_back_reg_log GROUP BY platform_id, root_game_id, user_id ) AS tmp_back_reg ON back_reg.id = tmp_back_reg.max_id 
	) AS reg ON login.platform_id = reg.platform_id 
	AND login.root_game_id = reg.root_game_id 
	AND login.user_id = reg.user_id 
	AND login.login_time > reg.last_time
	LEFT JOIN dwd_root_game_back_reg_log AS back_reg ON back_reg.platform_id = reg.platform_id 
	AND back_reg.root_game_id = reg.root_game_id 
	AND back_reg.user_id = reg.user_id 
WHERE
	back_reg.platform_id IS NULL;`},
		{Name: "更新30天回流用户last_time", Spec: "* */5 * * * *", Remark: "通过登录日志更新30天回流last_time", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `REPLACE INTO dwd_root_game_back_reg_log (
	platform_id,
	root_game_id,
	user_id,
	reg_time,
	last_time,
	game_id,
	agent_id,
	site_id,
	media_site_id,
	idfv,
	imei,
	oaid,
	andriod_id,
	system_version,
	app_version_code,
	sdk_version_code,
	network,
	client_ip,
	ipv4,
	ipv6,
	channel_id,
	model,
	brand,
	user_agent 
) SELECT
reg.platform_id,
reg.root_game_id,
reg.user_id,
reg.reg_time,
CONCAT(date(DATE_ADD(login.login_time, INTERVAL 30 DAY)), " 23:59:59") AS last_time,
reg.game_id,
reg.agent_id,
reg.site_id,
reg.media_site_id,
reg.idfv,
reg.imei,
reg.oaid,
reg.andriod_id,
reg.system_version,
reg.app_version_code,
reg.sdk_version_code,
reg.network,
reg.client_ip,
reg.ipv4,
reg.ipv6,
reg.channel_id,
reg.model,
reg.brand,
reg.user_agent 
FROM
	(
	SELECT
		login.*,
		main_game.root_game_id 
	FROM
		ods_login_logs AS login
		JOIN ( SELECT platform_id, game_id, user_id, max( id ) AS max_id FROM ods_login_logs WHERE updated_at BETWEEN '{{StartDateTime}}' AND '{{EndDateTime}}' GROUP BY platform_id, game_id, user_id ) AS user_min_log ON login.id = user_min_log.max_id
		JOIN dim_game AS game ON user_min_log.platform_id = game.platform_id 
		AND user_min_log.game_id = game.id
		JOIN dim_main_game AS main_game ON game.platform_id = main_game.platform_id 
		AND game.main_id = main_game.id
		JOIN dim_root_game AS root_game ON main_game.platform_id = root_game.platform_id 
		AND main_game.root_game_id = root_game.id 
	) AS login
	JOIN (
	SELECT
		back_reg.* 
	FROM
		dwd_root_game_back_reg_log AS back_reg
		JOIN ( SELECT platform_id, root_game_id, user_id, max( id ) AS max_id FROM dwd_root_game_back_reg_log GROUP BY platform_id, root_game_id, user_id ) AS tmp_back_reg ON back_reg.id = tmp_back_reg.max_id 
	) AS reg ON login.platform_id = reg.platform_id 
	AND login.root_game_id = reg.root_game_id 
	AND login.user_id = reg.user_id 
	AND login.login_time BETWEEN reg.reg_time AND reg.last_time
	;`},
		{Name: "30天回流激活统计", Spec: "* */5 * * * *", Remark: "激活日志表洗入每天30天回流overview表", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `REPLACE INTO dws_day_root_game_back_overview_log (
	platform_id,
	stat_date,
	game_id,
	agent_id,
	site_id,
	ad3_id,
	activation,
	activation_device,
	launch,
	launch_device,
	reg,
	reg_device,
	login,
	login_user,
	login_device,
	role,
	role_user,
	role_device,
	pay,
	pay_user,
	pay_device,
	pay_money 
) SELECT
launch.platform_id,
launch.stat_date,
launch.game_id,
launch.agent_id,
launch.site_id,
launch.ad3_id,
launch.activation,
launch.activation_device,
IFNULL( overview.launch, 0 ) AS launch,
IFNULL( overview.launch_device, 0 ) AS launch_device,
IFNULL( overview.reg, 0 ) AS reg,
IFNULL( overview.reg_device, 0 ) AS reg_device,
IFNULL( overview.login, 0 ) AS login,
IFNULL( overview.login_user, 0 ) AS login_user,
IFNULL( overview.login_device, 0 ) AS login_device,
IFNULL( overview.role, 0 ) AS role,
IFNULL( overview.role_user, 0 ) AS role_user,
IFNULL( overview.role_device, 0 ) AS role_device,
IFNULL( overview.pay, 0 ) AS pay,
IFNULL( overview.pay_user, 0 ) AS pay_user,
IFNULL( overview.pay_device, 0 ) AS pay_device,
IFNULL( overview.pay_money, 0 ) AS pay_money 
FROM
	(
	SELECT
		launch.platform_id,
		DATE ( launch.action_time ) AS stat_date,
		launch.game_id,
		launch.agent_id,
		launch.site_id,
		launch.ad3_id,
		count( 1 ) AS activation,
		count( DISTINCT unique_device ) AS activation_device 
	FROM
		ods_launch_log AS launch
	WHERE
		launch.action_time BETWEEN '{{StartDate}}' 
		AND '{{EndDate}}' 
		AND launch.action = 'active' 
	GROUP BY
		launch.platform_id,
		DATE ( launch.action_time ),
		launch.game_id,
		launch.agent_id,
		launch.site_id,
		launch.ad3_id 
	) AS launch
	LEFT JOIN dws_day_root_game_back_overview_log AS overview ON launch.platform_id = overview.platform_id 
	AND launch.stat_date = overview.stat_date 
	AND launch.game_id = overview.game_id 
	AND launch.agent_id = overview.agent_id 
	AND launch.site_id = overview.site_id 
	AND launch.ad3_id = overview.ad3_id 
WHERE
	launch.activation != overview.activation 
	OR launch.activation_device != overview.activation_device
;`},
		{Name: "30天回流注册写入overview天表", Spec: "* */5 * * * *", Remark: "30天回流注册日志表写入overview天表", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `REPLACE INTO dws_day_root_game_back_overview_log (
	platform_id,
	stat_date,
	game_id,
	agent_id,
	site_id,
	ad3_id,
	activation,
	activation_device,
	launch,
	launch_device,
	reg,
	reg_device,
	login,
	login_user,
	login_device,
	role,
	role_user,
	role_device,
	pay,
	pay_user,
	pay_device,
	pay_money 
) SELECT
tmp.platform_id,
tmp.stat_date,
tmp.game_id,
tmp.agent_id,
tmp.site_id,
tmp.ad3_id,
IFNULL( overview.activation, 0 ) AS activation,
IFNULL( overview.activation_device, 0 ) AS activation_device,
IFNULL( overview.launch, 0 ) AS launch,
IFNULL( overview.launch_device, 0 ) AS launch_device,
tmp.reg AS reg,
tmp.reg_device AS reg_device,
IFNULL( overview.login, 0 ) AS login,
IFNULL( overview.login_user, 0 ) AS login_user,
IFNULL( overview.login_device, 0 ) AS login_device,
IFNULL( overview.role, 0 ) AS role,
IFNULL( overview.role_user, 0 ) AS role_user,
IFNULL( overview.role_device, 0 ) AS role_device,
IFNULL( overview.pay, 0 ) AS pay,
IFNULL( overview.pay_user, 0 ) AS pay_user,
IFNULL( overview.pay_device, 0 ) AS pay_device,
IFNULL( overview.pay_money, 0 ) AS pay_money 
FROM
	(
	SELECT
		reg.platform_id,
		DATE ( reg.reg_time ) AS stat_date,
		reg.game_id,
		reg.agent_id,
		reg.site_id,
		reg.ad3_id,
		count( 1 ) AS reg,
		count( DISTINCT unique_device ) AS reg_device 
	FROM
		dwd_root_game_back_reg_log AS reg
	WHERE
		reg.reg_time BETWEEN '{{StartDate}}' AND '{{EndDate}}' 
	GROUP BY
		reg.platform_id,
		DATE ( reg.reg_time ),
		reg.game_id,
		reg.agent_id,
		reg.site_id,
		reg.ad3_id 
	) AS tmp
	LEFT JOIN dws_day_root_game_back_overview_log AS overview ON tmp.platform_id = overview.platform_id 
	AND tmp.stat_date = overview.stat_date 
	AND tmp.game_id = overview.game_id 
	AND tmp.agent_id = overview.agent_id 
	AND tmp.site_id = overview.site_id 
	AND tmp.ad3_id = overview.ad3_id 
WHERE
	tmp.reg != overview.reg 
	OR tmp.reg_device != overview.reg_device
	;`},
		{Name: "登录数据写入30天回流overview表", Spec: "* */5 * * * *", Remark: "从登录日志中写入30天回流overview天表", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `REPLACE INTO dws_day_root_game_back_overview_log (
	platform_id,
	stat_date,
	game_id,
	agent_id,
	site_id,
	ad3_id,
	activation,
	activation_device,
	launch,
	launch_device,
	reg,
	reg_device,
	login,
	login_user,
	login_device,
	role,
	role_user,
	role_device,
	pay,
	pay_user,
	pay_device,
	pay_money 
) SELECT
tmp.platform_id,
tmp.stat_date,
tmp.game_id,
tmp.agent_id,
tmp.site_id,
tmp.ad3_id,
IFNULL( overview.activation, 0 ) AS activation,
IFNULL( overview.activation_device, 0 ) AS activation_device,
IFNULL( overview.launch, 0 ) AS launch,
IFNULL( overview.launch_device, 0 ) AS launch_device,
IFNULL( overview.reg, 0 ) AS reg,
IFNULL( overview.reg_device, 0 ) AS reg_device,
tmp.login AS login,
tmp.login_user AS login_user,
tmp.login_device AS login_device,
IFNULL( overview.role, 0 ) AS role,
IFNULL( overview.role_user, 0 ) AS role_user,
IFNULL( overview.role_device, 0 ) AS role_device,
IFNULL( overview.pay, 0 ) AS pay,
IFNULL( overview.pay_user, 0 ) AS pay_user,
IFNULL( overview.pay_device, 0 ) AS pay_device,
IFNULL( overview.pay_money, 0 ) AS pay_money 
FROM
	(
	SELECT
		reg.platform_id,
		DATE ( login.login_time ) AS stat_date,
		reg.game_id,
		reg.agent_id,
		reg.site_id,
		reg.ad3_id,
		count( 1 ) AS login,
		count( DISTINCT login.user_id ) AS login_user,
		count( DISTINCT login.unique_device ) AS login_device 
	FROM
		dwd_root_game_back_reg_log AS reg
		JOIN dim_root_game AS root_game
		ON reg.platform_id = root_game.platform_id AND reg.root_game_id = root_game.id
		JOIN dim_main_game AS main_game
		ON root_game.platform_id = main_game.platform_id AND root_game.id = main_game.root_game_id
		JOIN dim_game AS game
		ON main_game.platform_id = game.platform_id AND main_game.id = game.main_id
		JOIN ods_login_logs AS login
		ON reg.platform_id = login.platform_id AND reg.user_id = login.user_id AND game.id = login.game_id
	WHERE
		login.login_time BETWEEN '{{StartDate}}' AND '{{EndDate}}' 
	GROUP BY
		reg.platform_id,
		DATE ( login.login_time ),
		reg.game_id,
		reg.agent_id,
		reg.site_id,
		reg.ad3_id 
	) AS tmp
	LEFT JOIN dws_day_root_game_back_overview_log AS overview ON tmp.platform_id = overview.platform_id 
	AND tmp.stat_date = overview.stat_date 
	AND tmp.game_id = overview.game_id 
	AND tmp.agent_id = overview.agent_id 
	AND tmp.site_id = overview.site_id 
	AND tmp.ad3_id = overview.ad3_id 
WHERE
	tmp.login != overview.login 
	OR tmp.login_user != overview.login_user
    OR tmp.login_device != overview.login_device
	;`},
		{Name: "创角数据写入30天回流overview表", Spec: "* */5 * * * *", Remark: "从创角日志中写入30天回流overview表", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `REPLACE INTO dws_day_root_game_back_overview_log (
	platform_id,
	stat_date,
	game_id,
	agent_id,
	site_id,
	ad3_id,
	activation,
	activation_device,
	launch,
	launch_device,
	reg,
	reg_device,
	login,
	login_user,
	login_device,
	role,
	role_user,
	role_device,
	pay,
	pay_user,
	pay_device,
	pay_money 
) SELECT
tmp.platform_id,
tmp.stat_date,
tmp.game_id,
tmp.agent_id,
tmp.site_id,
tmp.ad3_id,
IFNULL( overview.activation, 0 ) AS activation,
IFNULL( overview.activation_device, 0 ) AS activation_device,
IFNULL( overview.launch, 0 ) AS launch,
IFNULL( overview.launch_device, 0 ) AS launch_device,
IFNULL( overview.reg, 0 ) AS reg,
IFNULL( overview.reg_device, 0 ) AS reg_device,
IFNULL( overview.login, 0 ) AS login,
IFNULL( overview.login_user, 0 ) AS login_user,
IFNULL( overview.login_device, 0 ) AS login_device,
tmp.role AS role,
tmp.role_user AS role_user,
tmp.role_device AS role_device,
IFNULL( overview.pay, 0 ) AS pay,
IFNULL( overview.pay_user, 0 ) AS pay_user,
IFNULL( overview.pay_device, 0 ) AS pay_device,
IFNULL( overview.pay_money, 0 ) AS pay_money 
FROM
	(
	SELECT
		reg.platform_id,
		DATE ( action.action_time ) AS stat_date,
		reg.game_id,
		reg.agent_id,
		reg.site_id,
		reg.ad3_id,
		count( 1 ) AS role,
		count( DISTINCT action.user_id ) AS role_user,
		count( DISTINCT action.unique_device ) AS role_device 
	FROM
		dwd_root_game_back_reg_log AS reg
		JOIN dim_root_game AS root_game
		ON reg.platform_id = root_game.platform_id AND reg.root_game_id = root_game.id
		JOIN dim_main_game AS main_game
		ON root_game.platform_id = main_game.platform_id AND root_game.id = main_game.root_game_id
		JOIN dim_game AS game
		ON main_game.platform_id = game.platform_id AND main_game.id = game.main_id
		JOIN ods_game_behavior_log AS action
		ON reg.platform_id = action.platform_id AND reg.user_id = action.user_id AND game.id = action.game_id
	WHERE
		action.action = 'role-create'
		AND action.action_time BETWEEN '{{StartDate}}' AND '{{EndDate}}' 
	GROUP BY
		reg.platform_id,
		DATE ( action.action_time ),
		reg.game_id,
		reg.agent_id,
		reg.site_id,
		reg.ad3_id 
	) AS tmp
	LEFT JOIN dws_day_root_game_back_overview_log AS overview ON tmp.platform_id = overview.platform_id 
	AND tmp.stat_date = overview.stat_date 
	AND tmp.game_id = overview.game_id 
	AND tmp.agent_id = overview.agent_id 
	AND tmp.site_id = overview.site_id 
	AND tmp.ad3_id = overview.ad3_id 
WHERE
	tmp.role != overview.role 
	OR tmp.role_user != overview.role_user
	OR tmp.role_device != overview.role_device
	;`},
		{Name: "付费数据写入30天回流overview表", Spec: "* */5 * * * *", Remark: "从付费日志中写入30天回流overview天表", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `REPLACE INTO dws_day_root_game_back_overview_log (
	platform_id,
	stat_date,
	game_id,
	agent_id,
	site_id,
	ad3_id,
	activation,
	activation_device,
	launch,
	launch_device,
	reg,
	reg_device,
	login,
	login_user,
	login_device,
	role,
	role_user,
	role_device,
	pay,
	pay_user,
	pay_device,
	pay_money 
) SELECT
tmp.platform_id,
tmp.stat_date,
tmp.game_id,
tmp.agent_id,
tmp.site_id,
tmp.ad3_id,
IFNULL( overview.activation, 0 ) AS activation,
IFNULL( overview.activation_device, 0 ) AS activation_device,
IFNULL( overview.launch, 0 ) AS launch,
IFNULL( overview.launch_device, 0 ) AS launch_device,
IFNULL( overview.reg, 0 ) AS reg,
IFNULL( overview.reg_device, 0 ) AS reg_device,
IFNULL( overview.login, 0 ) AS login,
IFNULL( overview.login_user, 0 ) AS login_user,
IFNULL( overview.login_device, 0 ) AS login_device,
IFNULL( overview.role, 0 ) AS role,
IFNULL( overview.role_user, 0 ) AS role_user,
IFNULL( overview.role_device, 0 ) AS role_device,
tmp.pay AS pay,
tmp.pay_user AS pay_user,
tmp.pay_device AS pay_device,
tmp.pay_money AS pay_money 
FROM
	(
	SELECT
		reg.platform_id,
		DATE ( pay.pay_time ) AS stat_date,
		reg.game_id,
		reg.agent_id,
		reg.site_id,
		reg.ad3_id,
		count( 1 ) AS pay,
		count( DISTINCT pay.user_id ) AS pay_user,
		count( DISTINCT pay.unique_device ) AS pay_device,
		sum( pay.money ) AS pay_money
	FROM
		dwd_root_game_back_reg_log AS reg
		JOIN dim_root_game AS root_game
		ON reg.platform_id = root_game.platform_id AND reg.root_game_id = root_game.id
		JOIN dim_main_game AS main_game
		ON root_game.platform_id = main_game.platform_id AND root_game.id = main_game.root_game_id
		JOIN dim_game AS game
		ON main_game.platform_id = game.platform_id AND main_game.id = game.main_id
		JOIN ods_pay_log AS pay
		ON reg.platform_id = pay.platform_id AND reg.user_id = pay.user_id AND game.id = pay.game_id
	WHERE
		pay.pay_status = 'success'
		AND pay.test_order = 0
		AND pay.pay_time BETWEEN '{{StartDate}}' AND '{{EndDate}}' 
	GROUP BY
		reg.platform_id,
		DATE ( pay.pay_time ),
		reg.root_game_id,
		reg.agent_id,
		reg.site_id,
		reg.ad3_id 
	) AS tmp
	LEFT JOIN dws_day_root_game_back_overview_log AS overview ON tmp.platform_id = overview.platform_id 
	AND tmp.stat_date = overview.stat_date 
	AND tmp.game_id = overview.game_id 
	AND tmp.agent_id = overview.agent_id 
	AND tmp.site_id = overview.site_id 
	AND tmp.ad3_id = overview.ad3_id 
WHERE
	tmp.pay != overview.pay 
	OR tmp.pay_user != overview.pay_user
	OR tmp.pay_device != overview.pay_device
	OR tmp.pay_money != overview.pay_money
	;`},
		{Name: "30天回流付费情况", Spec: "* */5 * * * *", Remark: "通过注册和付费日志表清洗为N日付费情况", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `REPLACE INTO dws_day_root_game_back_pay_active_log ( platform_id, pay_date, game_id, agent_id, site_id, ad3_id, reg_date, active_days, pay_count, active_user, pay_amount ) SELECT
reg.platform_id AS platform_id,
DATE ( pay.pay_time ) AS pay_date,
reg.game_id AS game_id,
reg.agent_id AS agent_id,
reg.site_id AS site_id,
reg.ad3_id AS ad3_id,
DATE ( reg.reg_time ) AS reg_date,
DATEDIFF( pay.pay_time, reg.reg_time ) + 1 AS active_days,
count( 1 ) AS pay_count,
count( DISTINCT pay.user_id ) AS active_user,
sum( pay.money ) AS pay_amount 
FROM
	ods_pay_log AS pay
	JOIN dim_game AS game ON pay.platform_id = game.platform_id 
	AND pay.game_id = game.id
	JOIN dim_main_game AS main_game ON game.platform_id = main_game.platform_id 
	AND game.main_id = main_game.id
	JOIN dim_root_game AS root_game ON main_game.platform_id = root_game.platform_id 
	AND main_game.root_game_id = root_game.id
	JOIN dwd_root_game_back_reg_log AS reg ON pay.platform_id = reg.platform_id 
	AND pay.user_id = reg.user_id 
	AND main_game.root_game_id = reg.root_game_id 
WHERE
	pay.pay_time BETWEEN '{{StartDate}}' 
	AND '{{EndDate}}' 
	AND pay.test_order = 0
	AND pay.pay_status = 'success' 
GROUP BY
	reg.platform_id,
	DATE ( pay.pay_time ),
	reg.root_game_id,
	reg.agent_id,
	reg.site_id,
	reg.ad3_id,
	DATE (reg.reg_time)
;`},
		{Name: "30天回流留存情况", Spec: "* */5 * * * *", Remark: "从注册日志和登录日志清洗30天活跃数据", Status: sql.StatusNormal, TaskType: cron_task2.TaskTypeSqlCleaning, ExecutionMode: cron_task2.ExecutionModeAsync, Content: `REPLACE INTO dws_day_root_game_back_login_active_log ( platform_id, login_date, game_id, agent_id, site_id, ad3_id, reg_date, active_days, active_count )
 SELECT
reg.platform_id AS platform_id,
DATE ( login.login_time ) AS login_date,
reg.game_id AS game_id,
reg.agent_id AS agent_id,
reg.site_id AS site_id,
reg.ad3_id AS ad3_id,
DATE ( reg.reg_time ) AS reg_date,
DATEDIFF( login.login_time, reg.reg_time ) + 1 AS active_days,
count( DISTINCT login.user_id ) AS active_count 
FROM
	ods_login_logs AS login
	JOIN dim_game AS game ON login.platform_id = game.platform_id 
	AND login.game_id = game.id
	JOIN dim_main_game AS main_game ON game.platform_id = main_game.platform_id 
	AND game.main_id = main_game.id
	JOIN dim_root_game AS root_game ON main_game.platform_id = root_game.platform_id 
	AND main_game.root_game_id = root_game.id
	JOIN dwd_root_game_back_reg_log AS reg ON login.platform_id = reg.platform_id 
	AND login.user_id = reg.user_id 
	AND root_game.id = reg.root_game_id 
WHERE
	login.login_time BETWEEN '{{StartDate}}' 
	AND '{{EndDate}}' 
GROUP BY
	reg.platform_id,
	DATE ( login.login_time ),
	reg.root_game_id,
	reg.agent_id,
	reg.site_id,
	reg.ad3_id,
	DATE ( reg.reg_time );`},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, (&cron_task.DimCronTaskConfigModel{}).TableName()+"表数据初始化失败!")
	}
	return ctx, nil
}

func (i *initCronTaskConfigMysql) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	lookup := cron_task2.DimCronTaskConfigModel{Name: "子注册清洗"}
	if errors.Is(db.First(&lookup, &lookup).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
