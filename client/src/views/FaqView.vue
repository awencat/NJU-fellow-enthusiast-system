<template>
  <article class="faq-notes">
    <RouterLink to="/">← 返回首页</RouterLink>
    <h1>常见问题</h1>
    <section>
      <h2>注册需要校园邮箱吗？</h2>
      <p>需要。请使用注册页面提示的校园邮箱后缀，填写邮箱并输入邮件中的 6 位验证码完成注册。</p>
    </section>
    <section>
      <h2>收不到验证码怎么办？</h2>
      <p>验证码有效期为 60 秒。请先检查垃圾邮件，倒计时结束后可重新发送。若仍无法收到，请到 GitHub Issues 反馈，并说明注册页面的提示。</p>
    </section>
    <section>
      <h2>为什么只能和部分同学私信？</h2>
      <p>搭子申请通过后才会开放私信。请先在搭子详情页提交申请，等待对方通过。</p>
    </section>
    <section>
      <h2>发现不合适的内容怎么办？</h2>
      <p>可对帖子或用户提交举报，由管理员核实处理。</p>
    </section>
    <section>
      <h2>账号被封禁后怎么办？</h2>
      <p>封禁后无法正常登录和使用社区功能。若对处理有异议，可通过 GitHub Issues 说明情况，请勿公开密码、验证码等敏感信息。</p>
    </section>
    <section>
      <h2>图片上传有什么限制？</h2>
      <p>支持 JPG、PNG、WebP 和 GIF，单张图片不超过 5MB。</p>
    </section>
    <section>
      <h2>还有问题或改进建议？</h2>
      <p><a href="https://github.com/awencat/NJU-fellow-enthusiast-system/issues">前往 GitHub Issues →</a></p>
    </section>
  </article>
</template>

<style scoped>
.faq-notes {
  max-width: 800px;
  margin: 32px auto;
  padding: 24px;
  background: white;
  color: #222;
  font-family: "Microsoft YaHei", sans-serif;
  line-height: 1.8;
}

.faq-notes h1 {
  margin: 24px 0;
  font-family: inherit;
  font-size: 26px;
}

.faq-notes h2 {
  margin: 24px 0 8px;
  font-family: inherit;
  font-size: 18px;
  letter-spacing: normal;
}

.faq-notes p {
  margin: 0;
}

.faq-notes a {
  color: #2458a6;
  text-decoration: underline;
}
</style>
