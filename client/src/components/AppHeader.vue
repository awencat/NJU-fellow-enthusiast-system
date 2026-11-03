<script setup lang="ts">
import { computed, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import {
  ChevronDown,
  FileText,
  LogOut,
  Menu,
  MessageCircle,
  Settings,
  UserRound,
  Users,
  X,
} from "lucide-vue-next";
import Avatar from "./Avatar.vue";
import { auth } from "@/stores/auth";
import { notify } from "@/stores/toast";

const route = useRoute();
const router = useRouter();
const mobileOpen = ref(false);
const userMenuOpen = ref(false);

const pageMeta = computed(() => {
  if (route.path.startsWith("/exchange")) return { key: "exchange", kicker: "CAMPUS FEED", title: "同好交流" };
  if (route.path.startsWith("/mates")) return { key: "mates", kicker: "FIND A PARTNER", title: "找搭子" };
  if (route.path.startsWith("/messages")) return { key: "messages", kicker: "PRIVATE SPACE", title: "私信" };
  if (route.path.startsWith("/profile")) {
    const profileMeta: Record<string, { key: string; kicker: string; title: string }> = {
      overview: { key: "profile", kicker: "YOUR CAMPUS SPACE", title: "个人资料" },
      posts: { key: "profile-posts", kicker: "YOUR STORIES", title: "我的分享" },
      mates: { key: "profile-mates", kicker: "YOUR PLANS", title: "我的搭子" },
      applications: { key: "profile-applications", kicker: "REQUESTS", title: "申请记录" },
      edit: { key: "profile-edit", kicker: "PROFILE SETTINGS", title: "编辑资料" },
    };
    return profileMeta[String(route.query.tab || "overview")] || profileMeta.overview;
  }
  if (route.path.startsWith("/admin")) return { key: "admin", kicker: "CONTROL ROOM", title: "管理后台" };
  return { key: "home", kicker: "CAMPUS COMMUNITY", title: "发现" };
});

const nav = computed(() => [
  { to: "/", label: "发现" },
  { to: "/exchange", label: "同好交流" },
  { to: "/mates", label: "找搭子" },
  ...(auth.state.user ? [{ to: "/messages", label: "私信" }] : []),
]);

async function logout() {
  try {
    await auth.logout();
    mobileOpen.value = false;
    userMenuOpen.value = false;
    notify("已安全退出");
    await router.replace({ name: "login" });
  } catch (error) {
    notify(error instanceof Error ? error.message : "退出失败", "error");
  }
}
</script>

<template>
  <header class="site-header" :class="`header-${pageMeta.key}`">
    <div class="header-inner ui-row">
      <RouterLink class="brand ui-row" to="/" @click="mobileOpen = false">
        <span class="brand-mark ui-center">
          <span></span><span></span>
        </span>
        <span class="brand-copy ui-stack">
          <strong>搭个伴</strong>
          <small>校园同好社区</small>
        </span>
      </RouterLink>

      <div class="header-context" aria-live="polite">
        <span>{{ pageMeta.kicker }}</span>
        <strong>{{ pageMeta.title }}</strong>
      </div>

      <nav class="desktop-nav ui-row">
        <RouterLink
          v-for="item in nav"
          :key="item.to"
          :to="item.to"
          :class="{ active: item.to === '/' ? route.path === '/' : route.path.startsWith(item.to) }"
        >
          {{ item.label }}
        </RouterLink>
      </nav>

      <div class="header-actions ui-row">
        <template v-if="auth.state.user">
          <div class="user-menu">
            <button
              class="user-chip user-menu-trigger ui-row ui-panel"
              type="button"
              aria-haspopup="menu"
              :aria-expanded="userMenuOpen"
              @click="userMenuOpen = !userMenuOpen"
            >
              <Avatar :user="auth.state.user" :size="32" />
              <span>{{ auth.state.user.nickname }}</span>
              <ChevronDown :size="16" :class="{ rotated: userMenuOpen }" />
            </button>
            <Transition name="menu">
              <div v-if="userMenuOpen" class="user-menu-panel" role="menu">
                <RouterLink to="/profile" role="menuitem" @click="userMenuOpen = false">
                  <UserRound :size="16" /> 个人资料
                </RouterLink>
                <RouterLink to="/profile?tab=posts" role="menuitem" @click="userMenuOpen = false">
                  <FileText :size="16" /> 我的分享
                </RouterLink>
                <RouterLink to="/profile?tab=mates" role="menuitem" @click="userMenuOpen = false">
                  <Users :size="16" /> 我的搭子
                </RouterLink>
                <RouterLink to="/profile?tab=applications" role="menuitem" @click="userMenuOpen = false">
                  <MessageCircle :size="16" /> 申请记录
                </RouterLink>
                <RouterLink to="/profile?tab=edit" role="menuitem" @click="userMenuOpen = false">
                  <Settings :size="16" /> 编辑资料
                </RouterLink>
                <RouterLink v-if="auth.state.user.role === 'admin'" to="/admin" role="menuitem" @click="userMenuOpen = false">
                  <Settings :size="16" /> 管理后台
                </RouterLink>
                <button type="button" role="menuitem" @click="logout">
                  <LogOut :size="16" /> 退出登录
                </button>
              </div>
            </Transition>
          </div>
        </template>
        <template v-else>
          <RouterLink class="text-link ui-inline" to="/login">登录</RouterLink>
          <RouterLink class="button button-small ui-inline" to="/login?tab=register">
            加入校园
          </RouterLink>
        </template>
        <button
          class="mobile-menu-button"
          :aria-label="mobileOpen ? '关闭菜单' : '打开菜单'"
          @click="mobileOpen = !mobileOpen"
        >
          <X v-if="mobileOpen" :size="22" />
          <Menu v-else :size="22" />
        </button>
      </div>
    </div>

    <Transition name="menu">
      <nav v-if="mobileOpen" class="mobile-nav">
        <RouterLink
          v-for="item in nav"
          :key="item.to"
          :to="item.to"
          @click="mobileOpen = false"
        >
          {{ item.label }}
        </RouterLink>
        <RouterLink v-if="auth.state.user" to="/profile" @click="mobileOpen = false">
          <UserRound :size="18" /> 个人中心
        </RouterLink>
        <RouterLink v-if="auth.state.user" to="/profile?tab=posts" @click="mobileOpen = false">
          <FileText :size="18" /> 我的分享
        </RouterLink>
        <RouterLink v-if="auth.state.user" to="/profile?tab=mates" @click="mobileOpen = false">
          <Users :size="18" /> 我的搭子
        </RouterLink>
        <RouterLink v-if="auth.state.user" to="/profile?tab=applications" @click="mobileOpen = false">
          <MessageCircle :size="18" /> 申请记录
        </RouterLink>
        <RouterLink v-if="auth.state.user" to="/profile?tab=edit" @click="mobileOpen = false">
          <Settings :size="18" /> 编辑资料
        </RouterLink>
        <RouterLink
          v-if="auth.state.user?.role === 'admin'"
          to="/admin"
          @click="mobileOpen = false"
        >
          <Settings :size="18" /> 管理后台
        </RouterLink>
      </nav>
    </Transition>
  </header>
</template>
