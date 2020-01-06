package model

// Article 文章
type Article struct {
	ID          string   `bson:"_id,omitempty" json:"_id,omitempty"`                 // ID
	UserID      string   `bson:"userId,omitempty" json:"userId,omitempty"`           // 用户ID
	CompanyID   string   `bson:"companyId,omitempty" json:"companyId,omitempty"`     // 公司ID
	CreateTime  string   `bson:"createTime,omitempty" json:"createTime,omitempty"`   // 创建时间
	ShowTime    string   `bson:"showTime,omitempty" json:"showTime,omitempty"`       // 用于显示
	UpdateTime  string   `bson:"updateTime,omitempty" json:"updateTime,omitempty"`   // 最近一次修改时间
	Cover       string   `bson:"cover,omitempty" json:"cover,omitempty"`             // 封面URL
	Title       string   `bson:"title,omitempty" json:"title,omitempty"`             // 标题
	PV          int64    `bson:"pv,omitempty" json:"pv,omitempty"`                   // 浏览总数
	UV          int64    `bson:"uv,omitempty" json:"uv,omitempty"`                   // 浏览人数
	Tags        []string `bson:"tags" json:"tags"`                                   // 标签ID
	Context     string   `bson:"context,omitempty" json:"context,omitempty"`         // 文章内容
	Status      int      `bson:"status,omitempty" json:"status,omitempty"`           // 发布状态 1：已发布，0：未发布，2：已下线
	URL         string   `bson:"url,omitempty" json:"url,omitempty"`                 // 发布url。保留字段
	ReleaseTime string   `bson:"releaseTime,omitempty" json:"releaseTime,omitempty"` // 发布时间
}
