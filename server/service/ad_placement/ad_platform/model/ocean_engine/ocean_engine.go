package ocean_engine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/pkg/errors"
	"slices"
)

const (
	AccountTypeEbp               = "EBP"            // 账户类型 - EBP 升级版巨量引擎工作台
	AssetManagementScopeDirect   = "DIRECT"         // 资产范围 - 仅查询入参组织创建或被共享的资产
	AssetManagementScopeTraverse = "TRAVERSE"       // 资产范围 - 查询入参组织及下属组织创建或被共享的资产
	AssetOwnershipCreate         = "CREATE"         // 资产来源 - 仅查询组织及下级组织创建的资产
	AssetOwnershipShare          = "SHARE"          // 资产来源 - 仅查询组织及下级组织被共享的资产
	StatusOffline                = "OFFLINE"        // 筛选应用状态 - 已下架
	StatusPunished               = "PUNISHED"       // 筛选应用状态 - 违规惩罚
	StatusReviewing              = "REVIEWING"      // 筛选应用状态 - 审核中
	StatusReviewFail             = "REVIEW_FAIL"    // 筛选应用状态 - 审核失败
	StatusReviewSuccess          = "REVIEW_SUCCESS" // 筛选应用状态 - 待发布
	StatusReviewInuse            = "IN_USE"         // 筛选应用状态 - 已发布

	AccountRoleAdvertiser                       = "ADVERTISER"                           // 账户类型 - 客户
	AccountRoleCustomerAdmin                    = "CUSTOMER_ADMIN"                       // 账户类型 - 管理员
	AccountRoleCustomerOperator                 = "CUSTOMER_OPERATOR"                    // 账户类型 - 协作者
	AccountRoleCustomerAgent                    = "AGENT"                                // 账户类型 - 代理商
	AccountRoleChildAgent                       = "CHILD_AGENT"                          // 账户类型 - 二级代理商
	AccountRolePlatformRoleStar                 = "PLATFORM_ROLE_STAR"                   // 账户类型 - 星图账户
	AccountRolePlatformRoleQianchuanAgent       = "PLATFORM_ROLE_QIANCHUAN_AGENT"        // 账户类型 - 千川代理商
	AccountRolePlatformRoleStarAgent            = "PLATFORM_ROLE_STAR_AGENT"             // 账户类型 - 星图代理商
	AccountRolePlatformRoleAweme                = "PLATFORM_ROLE_AWEME"                  // 账户类型 - 抖音号
	AccountRolePlatformRoleStarMcn              = "PLATFORM_ROLE_STAR_MCN"               // 账户类型 - 星图MCN机构
	AccountRolePlatformRoleStarIsv              = "PLATFORM_ROLE_STAR_ISV"               // 账户类型 - 星图服务商
	AccountRoleAgentSystemAccount               = "AGENT_SYSTEM_ACCOUNT"                 // 账户类型 - 代理商系统账户
	AccountRolePlatformRoleLocalAgent           = "PLATFORM_ROLE_LOCAL_AGENT"            // 账户类型 - 本地推代理商
	AccountRolePlatformRoleYuntuBrandIsvAdmin   = "PLATFORM_ROLE_YUNTU_BRAND_ISV_ADMIN"  // 账户类型 - 云图品牌服务商管理员
	AccountRolePlatformRoleLife                 = "PLATFORM_ROLE_LIFE"                   // 账户类型 - 抖音来客账户
	AccountRolePlatformRoleEnterpriseBpAdmin    = "PLATFORM_ROLE_ENTERPRISE_BP_ADMIN"    // 账户类型 - 升级版工作台管理员
	AccountRolePlatformRoleEnterpriseBpOperator = "PLATFORM_ROLE_ENTERPRISE_BP_OPERATOR" // 账户类型 - 升级版工作台协作者

	AccountSourceAd    = "AD"    // 账户类型 - AD 巨量营销客户账号
	AccountSourceLocal = "LOCAL" // 账户类型 - LOCAL 本地推
)

type PageInfo struct {
	Page        int `json:"page"`
	PageSize    int `json:"page_size"`
	TotalPage   int `json:"total_page"`
	TotalNumber int `json:"total_number"`
}

type AppListReq struct {
	AccountId            int64  `json:"account_id"`
	AccountType          string `json:"account_type"`
	AssetManagementScope string `json:"asset_management_scope"`
	Filtering            struct {
		SearchKey      string `json:"search_key"`
		AssetOwnership string `json:"asset_ownership"`
		Status         string `json:"status"`
		PublishTime    struct {
			Start string `json:"start"` // 发布起始时间，格式：%Y-%m-%d
			End   string `json:"end"`   // 发布结束时间，格式：%Y-%m-%d
		} `json:"publish_time"`
	} `json:"filtering"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (receiver *AppListReq) Format() {
}

func (receiver *AppListReq) Validate() (err error) {
	if receiver.AccountId <= 0 {
		err = errors.New("AccountId为空")
		return
	}
	if receiver.AccountType != AccountTypeEbp {
		err = errors.New("AccountType非法。" + receiver.AccountType)
		return
	}
	if !slices.Contains([]string{
		AssetManagementScopeDirect,
		AssetManagementScopeTraverse,
	}, receiver.AssetManagementScope) {
		err = errors.New("AssetManagementScope非法。" + receiver.AssetManagementScope)
		return
	}
	if receiver.Filtering.AssetOwnership != "" && !slices.Contains([]string{
		AssetOwnershipCreate,
		AssetOwnershipShare,
	}, receiver.Filtering.AssetOwnership) {
		err = errors.New("Filtering.AssetOwnership非法。" + receiver.Filtering.AssetOwnership)
		return
	}
	return
}

type AppListResp struct {
	BasicAppList []advertising.OdsAdvertisingOceanengineAppLogModel `json:"basic_app_list"`
	PageInfo     PageInfo                                           `json:"page_info"`
}

type EbpAdvertiserListReq struct {
	EnterpriseOrganizationId int64  `json:"enterprise_organization_id"`
	AccountSource            string `json:"account_source"`
	Filtering                struct {
		AccountName   string `json:"account_name"`
		ActiveAccount bool   `json:"active_account"`
	} `json:"filtering"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (receiver *EbpAdvertiserListReq) Format() {
	if receiver.Page <= 0 {
		receiver.Page = 1
	}
	if receiver.PageSize <= 0 || receiver.PageSize > 100 {
		receiver.PageSize = 100
	}
}

func (receiver *EbpAdvertiserListReq) Validate() (err error) {
	if receiver.EnterpriseOrganizationId <= 0 {
		err = errors.New("enterprise_organization_id 为空")
		return
	}
	if !slices.Contains([]string{
		AccountSourceAd,
		AccountSourceLocal,
	}, receiver.AccountSource) {
		err = errors.New("AccountSource 非法." + receiver.AccountSource)
		return
	}
	return
}

type EbpAdvertiserListResp struct {
	AccountList []struct {
		AccountId   int64  `json:"account_id"`
		AccountType string `json:"account_type"`
		AccountName string `json:"account_name"`
	} `json:"account_list"`
	PageInfo PageInfo `json:"page_info"`
}
