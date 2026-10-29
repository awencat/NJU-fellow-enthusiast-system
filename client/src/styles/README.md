# 样式目录

入口为 `index.css`，由 `main.ts` 加载。导入顺序决定层叠顺序，调整顺序前请检查桌面与移动端。

| 文件 | 负责范围 |
| --- | --- |
| base.css | 字体、主题变量、元素默认值、辅助类 |
| primitives.css | 横排、纵排、面板、文字截断、卡片交互等公共组合 |
| header.css | 顶部导航、品牌、导航图标及文本链接 |
| shared.css | 按钮、标签、头像、页面容器、标题、空状态、分页 |
| home.css | 首页首屏、照片拼贴、社区介绍、搭子邀请区 |
| cards.css | 帖子卡片、搭子卡片、信息流网格 |
| lists.css | 交流与搭子列表、筛选、标签 |
| forms.css | 表单、编辑器、图片上传、标签选择、搜索框 |
| overlays.css | 弹窗、提示、菜单过渡 |
| exchange-detail.css / mate-detail.css | 两类详情页 |
| auth.css / messages.css / profile.css / admin.css | 登录注册、私信、个人中心、后台 |
| footer.css / not-found.css | 页脚、移动导航默认值、404 |
| responsive-desktop.css | 1050px 及以下适配 |
| responsive-tablet.css | 820px 及以下适配，含移动底部导航 |
| responsive-mobile.css | 620px 及以下适配 |
| motion.css | 减少动态效果的系统偏好 |

基础样式和页面样式先加载，响应式覆盖最后加载，多个断点同时生效时以较窄断点为准。响应式规则保留在统一断点文件，避免分散后改变跨模块的覆盖顺序。

修改时先查对应模块，再查三个响应式文件。组件自己的 `scoped` 样式仍保留在 `.vue` 文件中；`design-archive` 是历史设计，不参与当前页面样式加载。

不要仅凭模板中搜不到类名就删除规则：Vue 的过渡类、`toast-${type}`、路由活动类均由运行时生成。

## 手动美化从哪里改

- 全站颜色、阴影、圆角：修改 `base.css` 的 `:root`。
- 面板边框和底色、横排/纵排、卡片悬浮、文字截断：修改 `primitives.css`，影响使用同一公共类的所有组件。实际使用清单见 [PATTERNS.md](PATTERNS.md)。
- 某张卡片的尺寸或某个页面的间距：修改对应模块，公共布局不再重复抄写。
- 三列卡片的间距：修改 `cards.css` 的 `--feed-gap`；卡片悬浮高度使用 `--card-lift`。
- 搜索框：统一在 `forms.css` 中修改，首页与列表不再分开维护。

例如 `class="mate-card ui-stack ui-panel ui-card"`：`mate-card` 管搭子卡片自己的尺寸，`ui-stack` 管纵向布局，`ui-panel` 管面板外观，`ui-card` 管悬浮交互。新增同类卡片直接组合这些类，不复制公共声明。模块规则后加载，可覆盖公共规则；不要给公共类加 `!important`。
