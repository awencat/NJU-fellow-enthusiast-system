# 公共样式使用索引

这些类在 `primitives.css` 中定义，先于模块加载；模块只定义差异。

## ui-between：横排，两端对齐

`admin-panel-head`、`post-kind-row`、`post-author-row`、`mate-card-footer`、`filter-bar`、`list-summary`、`conversation-head`、`modal-header`、`profile-hero`、`profile-edit-heading`

## ui-heading-row：标题与操作区底部对齐

`admin-heading`、`split-heading`、`section-heading`

## ui-stack：纵向排列

`report-admin-list`、`post-cover-placeholder`、`post-card-body`、`mate-card`、`mate-meta`、`mate-callout-list`、`stack-form`、`upload-placeholder`、`brand-copy`、`hero-note`、`hero-float-card`、`chat-panel`、`not-found`、`toast-stack`、`application-summary-list`、`empty-state`

## ui-wrap：可换行的标签或操作区

`report-resolution`、`auth-points`、`footer-links`、`tag-options`、`selected-tags`、`hero-quick-links`、`popular-tags`、`application-actions`、`profile-actions`、`tag-row`

## ui-row：横排，垂直居中

`report-card-head`、`report-reason`、`auth-story`、`article-meta`、`article-author`、`article-actions`、`owner-actions`、`comment-head`、`footer-heading`、`header-inner`、`brand`、`desktop-nav`、`header-actions`、`user-chip`、`hero-search`、`search-field`、`mate-author-line`、`status-callout`、`application-head`、`chat-header`、`chat-intro`、`message-time`、`toast`、`profile-identity`、`inline-user`、`loading-block`、`pagination`

## ui-inline：行内图标与文字

`floating-badge`、`application-count`、`footer-brand`、`text-link`、`hero-kicker`、`online-dot`、`button`、`text-button`、`eyebrow`、`tag`、`back-link`、`status-pill`

## ui-center：图标居中

`auth-form-wrap`、`remove-image`、`brand-mark`、`modal-backdrop`、`application-summary-icon`

## ui-panel：标准面板：边框、底色与圆角

`admin-panel`、`report-admin-card`、`post-card`、`mate-card`、`exchange-article`、`comment-composer`、`editor-card`、`user-chip`、`hero-search`、`search-field`、`mate-search-panel`、`mate-detail-main`、`mate-apply-panel`、`application-card`、`chat-shell`、`toast`、`profile-edit-card`、`profile-stat-row`

## ui-ellipsis：单行省略

会话最后一条消息、来源帖子、聊天标题和后台内容摘要。

## ui-summary：三行摘要

`mate-summary`、帖子卡片正文。

## ui-card：卡片悬浮交互

`post-card`、`mate-card`。阴影统一使用 `--shadow-small`；悬浮高度使用 `--card-lift`，默认 `-3px`。

## 信息流网格

`post-grid` 和 `mate-grid` 共用列数规则；间距使用 `--feed-gap`，默认 `20px`，帖子列表为 `22px`。断点仍由三个 responsive 文件控制。
