<script setup lang="ts">
import { Home, MessageCircle, UserRound, Users } from "lucide-vue-next";
import AppHeader from "@/components/AppHeader.vue";
import AppFooter from "@/components/AppFooter.vue";
import ToastHost from "@/components/ToastHost.vue";
import { auth } from "@/stores/auth";
</script>

<template>
  <AppHeader />
  <main class="app-main">
    <RouterView />
  </main>
  <AppFooter />

  <nav v-if="auth.state.user" class="bottom-nav">
    <RouterLink to="/">
      <Home :size="20" /><span>发现</span>
    </RouterLink>
    <RouterLink to="/mates">
      <Users :size="20" /><span>找搭子</span>
    </RouterLink>
    <RouterLink to="/messages">
      <MessageCircle :size="20" /><span>私信</span>
    </RouterLink>
    <RouterLink to="/profile">
      <UserRound :size="20" /><span>我的</span>
    </RouterLink>
  </nav>
  <ToastHost />
</template>

