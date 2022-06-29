package apiService

/*
DROP TABLE IF EXISTS `t_api_service`;
CREATE TABLE `t_api_service`  (
`id` bigint(20) NOT NULL COMMENT '主键',
`module_id` bigint(20) NOT NULL COMMENT '模块id',
`api_name` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'api名称',
`api_des` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'api描述',
`http_method` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'http请求方式',
`url` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '后端请求地址',
`port` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '后端请求端口号',
`audience_ident` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'audience标识',
`path` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'api访问路径',
`is_check` char(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '是否检测',
`param_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '传参形式',
`is_param` char(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '是否有传参',
`request_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '传参格式',
`status` tinyint(4) NOT NULL COMMENT '系统状态 -1-停用，1-启用',
`create_user` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
`create_time` datetime(3) NOT NULL COMMENT '创建时间',
`update_user` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '更新人',
`update_time` datetime(3) NOT NULL COMMENT '创建时间',
PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'API服务表' ROW_FORMAT = DYNAMIC;

*/

type Api struct {
	Id            int    `json:"id"`
	ModuleId      int    `json:"moduleId"`
	ApiName       string `json:"apiName"`
	ApiDes        string `json:"apiDes"`
	HttpMethod    string `json:"httpMethod"`
	Url           string `json:"url"`
	Port          string `json:"port"`
	AudienceIdent string `json:"audienceIdent"`
	Path          string `json:"path"`
	IsCheck       string `json:"isCheck"`
	ParamType     string `json:"paramType"`
	IsParam       string `json:"isParam"`
	RequestParam  string `json:"requestParam"`
	Status        string `json:"status"`
	CreatedUser   string `json:"createdUser"`
	CreatedTime   string `json:"createdTime"`
	UpdatedUser   string `json:"updatedUser"`
	UpdatedTime   string `json:"updatedTime"`
}

func (model *Api) TableName() string {
	return "t_api_service"
}
