package enum

// 说明
// 不能以零值开始设计枚举值，当枚举值为零值时，gorm 默认不更新该值
// 奇偶作为正反

type MenuType uint8

const (
	MenuDir MenuType = iota + 1
	Menu
	MenuBtn
)

// Name 返回枚举值的名称
func (e MenuType) Name() string {
	if e < MenuDir || e > MenuBtn {
		return "未知"
	}
	return [...]string{"目录", "菜单", "按钮"}[e]
}

// 将枚举值转成字符串，便于输出
func (e MenuType) String() string {
	return e.Name()
}