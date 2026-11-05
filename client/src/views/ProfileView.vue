<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import {
  ArrowLeft,
  ArrowRight,
  CalendarDays,
  ChevronRight,
  FileText,
  MessageCircle,
  Save,
  Settings,
  Users,
} from "lucide-vue-next";
import { api } from "@/api";
import { auth } from "@/stores/auth";
import { notify } from "@/stores/toast";
import type { ExchangePost, MatePost, Paginated } from "@/types";
import Avatar from "@/components/Avatar.vue";
import ImageUpload from "@/components/ImageUpload.vue";
import LoadingBlock from "@/components/LoadingBlock.vue";
import EmptyState from "@/components/EmptyState.vue";
import PostCard from "@/components/PostCard.vue";
import MateCard from "@/components/MateCard.vue";

type ApplicationSummary = {
  id: number;
  message: string;
  status: "pending" | "accepted" | "rejected";
  createdAt: string;
  post: { id: number; title: string; status: string };
  applicantNickname: string;
  ownerNickname: string;
  conversationId: number | null;
};

type ProfileTab = "overview" | "posts" | "mates" | "applications" | "edit";
const route = useRoute();
const router = useRouter();
const validTabs: ProfileTab[] = ["overview", "posts", "mates", "applications", "edit"];
const activeTab = ref<ProfileTab>(validTabs.includes(route.query.tab as ProfileTab) ? route.query.tab as ProfileTab : "overview");
const loading = ref(true);
const saving = ref(false);
const applicationFilter = ref<"all" | "sent" | "received">("all");
const posts = ref<ExchangePost[]>([]);
const mates = ref<MatePost[]>([]);
const sentApplications = ref<ApplicationSummary[]>([]);
const receivedApplications = ref<ApplicationSummary[]>([]);
const profileForm = reactive({
  nickname: auth.state.user?.nickname ?? "",
  bio: auth.state.user?.bio ?? "",
  avatar: auth.state.user?.avatar ?? null,
});
const allApplications = computed(() =>
  [...receivedApplications.value, ...sentApplications.value].sort(
    (a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime(),
  ),
);
const visibleApplications = computed(() => {
  if (applicationFilter.value === "sent") return sentApplications.value;
  if (applicationFilter.value === "received") return receivedApplications.value;
  return allApplications.value;
});

async function load() {
  loading.value = true;
  try {
    const [myPosts, myMates, sent, received] = await Promise.all([
      api<Paginated<ExchangePost>>("/api/exchange/posts?mine=true&pageSize=50"),
      api<Paginated<MatePost>>("/api/mates?mine=true&pageSize=50"),
      api<{ items: ApplicationSummary[] }>("/api/mates/applications/mine?scope=sent"),
      api<{ items: ApplicationSummary[] }>("/api/mates/applications/mine?scope=received"),
    ]);
    posts.value = myPosts.items;
    mates.value = myMates.items;
    sentApplications.value = sent.items;
    receivedApplications.value = received.items;
  } catch (error) {
    notify(error instanceof Error ? error.message : "个人内容加载失败", "error");
  } finally {
    loading.value = false;
  }
}

async function saveProfile() {
  saving.value = true;
  try {
    const data = await api<{ user: typeof auth.state.user }>("/api/auth/me", {
      method: "PATCH",
      body: JSON.stringify(profileForm),
    });
    if (data.user) auth.updateUser(data.user);
    selectTab("overview");
    notify("个人信息已更新");
  } catch (error) {
    notify(error instanceof Error ? error.message : "保存失败", "error");
  } finally {
    saving.value = false;
  }
}

function selectTab(tab: ProfileTab) {
  activeTab.value = tab;
  void router.push({ query: { ...route.query, tab: tab === "overview" ? undefined : tab } });
}

watch(() => route.query.tab, (tab) => {
  activeTab.value = validTabs.includes(tab as ProfileTab) ? tab as ProfileTab : "overview";
});

const sectionMeta = computed(() => ({
  overview: { eyebrow: "YOUR CAMPUS SPACE", title: "个人资料", description: "管理你的校园名片，查看在社区中的参与记录。", action: "" },
  posts: { eyebrow: "YOUR STORIES", title: "我的分享", description: "你发布的校园动态与兴趣记录。", action: "写一篇分享" },
  mates: { eyebrow: "YOUR PLANS", title: "我的搭子", description: "查看你发起的同行计划和活动邀请。", action: "发布搭子邀请" },
  applications: { eyebrow: "REQUESTS", title: "申请记录", description: "跟进你发出的申请，以及收到的同行请求。", action: "查看私信" },
  edit: { eyebrow: "PROFILE SETTINGS", title: "编辑资料", description: "更新头像、昵称和个人简介。", action: "" },
}[activeTab.value]));

onMounted(load);
</script>

<template>
  <div class="profile-page page-shell">
    <section v-if="auth.state.user && activeTab === 'overview'" class="profile-hero ui-between">
      <div class="profile-identity ui-row">
        <Avatar :user="auth.state.user" :size="82" />
        <div>
          <span class="eyebrow ui-inline">{{ auth.state.user.role === "admin" ? "社区管理员" : "校园同好" }}</span>
          <h1>{{ auth.state.user.nickname }}</h1>
          <p>@{{ auth.state.user.username }} · {{ auth.state.user.bio || "还没有填写个人简介" }}</p>
        </div>
      </div>
      <div class="profile-actions ui-wrap">
        <RouterLink v-if="auth.state.user.role === 'admin'" class="button secondary button-small ui-inline" to="/admin">
          <Settings :size="15" /> 管理后台
        </RouterLink>
      </div>
    </section>

    <header v-if="activeTab !== 'overview'" class="profile-section-heading ui-between">
      <div>
        <RouterLink class="profile-back-link" to="/profile"><ArrowLeft :size="15" /> 个人概览</RouterLink>
        <span class="eyebrow">{{ sectionMeta.eyebrow }}</span>
        <h1>{{ sectionMeta.title }}</h1>
        <p>{{ sectionMeta.description }}</p>
      </div>
      <RouterLink v-if="activeTab === 'posts'" class="button ui-inline" to="/exchange/new">
        <FileText :size="16" /> {{ sectionMeta.action }}
      </RouterLink>
      <RouterLink v-else-if="activeTab === 'mates'" class="button warm-button ui-inline" to="/mates/new">
        <Users :size="16" /> {{ sectionMeta.action }}
      </RouterLink>
      <RouterLink v-else-if="activeTab === 'applications'" class="button secondary ui-inline" to="/messages">
        <MessageCircle :size="16" /> {{ sectionMeta.action }}
      </RouterLink>
    </header>

    <div v-if="activeTab === 'overview' && !loading" class="profile-overview-grid">
      <article class="profile-overview-card">
        <img src="/images/atlas/journal.svg" alt="" /><strong>{{ posts.length }}</strong><span>篇分享</span>
        <button type="button" @click="selectTab('posts')">查看我的分享 <ArrowRight :size="15" /></button>
      </article>
      <article class="profile-overview-card">
        <img src="/images/atlas/compass.svg" alt="" /><strong>{{ mates.length }}</strong><span>个搭子计划</span>
        <button type="button" @click="selectTab('mates')">查看我的搭子 <ArrowRight :size="15" /></button>
      </article>
      <article class="profile-overview-card">
        <img src="/images/atlas/mail.svg" alt="" /><strong>{{ allApplications.length }}</strong><span>条申请记录</span>
        <button type="button" @click="selectTab('applications')">查看申请记录 <ArrowRight :size="15" /></button>
      </article>
    </div>

    <LoadingBlock v-if="loading" />
    <template v-else>
      <template v-if="activeTab === 'posts'">
        <div v-if="posts.length" class="post-grid">
          <PostCard v-for="post in posts" :key="post.id" :post="post" />
        </div>
        <EmptyState v-else title="你还没有发布分享" text="记录一个校园片段，也许会遇见想法相近的人。">
          <RouterLink class="button button-small ui-inline" to="/exchange/new">写一篇分享</RouterLink>
        </EmptyState>
      </template>

      <template v-if="activeTab === 'mates'">
        <div v-if="mates.length" class="mate-grid">
          <MateCard v-for="mate in mates" :key="mate.id" :mate="mate" />
        </div>
        <EmptyState v-else title="你还没有发布搭子邀请" text="把想做的事写下来，邀请同学一起。">
          <RouterLink class="button button-small ui-inline" to="/mates/new">发布搭子邀请</RouterLink>
        </EmptyState>
      </template>

      <template v-if="activeTab === 'applications'">
        <div class="application-filters" role="group" aria-label="筛选申请记录">
          <button :class="{ active: applicationFilter === 'all' }" @click="applicationFilter = 'all'">全部 <span>{{ allApplications.length }}</span></button>
          <button :class="{ active: applicationFilter === 'sent' }" @click="applicationFilter = 'sent'">我发出的 <span>{{ sentApplications.length }}</span></button>
          <button :class="{ active: applicationFilter === 'received' }" @click="applicationFilter = 'received'">收到的 <span>{{ receivedApplications.length }}</span></button>
        </div>
        <div v-if="visibleApplications.length" class="application-summary-list ui-stack">
          <article v-for="application in visibleApplications" :key="application.id">
            <div class="application-summary-icon ui-center">
              <CalendarDays :size="19" />
            </div>
            <div>
              <RouterLink :to="`/mates/${application.post.id}`">{{ application.post.title }}</RouterLink>
              <p>
                {{
                  receivedApplications.some((item) => item.id === application.id)
                    ? `收到 ${application.applicantNickname} 的申请`
                    : `由你向 ${application.ownerNickname} 发出`
                }}
              </p>
              <span class="status-pill ui-inline" :class="application.status">
                {{ application.status === "pending" ? "等待回应" : application.status === "accepted" ? "已通过" : "未通过" }}
              </span>
            </div>
            <RouterLink
              v-if="application.conversationId"
              class="text-link ui-inline"
              :to="`/messages?c=${application.conversationId}`"
            >
              私信 <ChevronRight :size="15" />
            </RouterLink>
          </article>
        </div>
        <EmptyState v-else :title="applicationFilter === 'sent' ? '还没有发出申请' : applicationFilter === 'received' ? '还没有收到申请' : '还没有申请记录'" text="找到感兴趣的搭子邀请，提交同行申请后可在这里跟进进度。" />
      </template>

      <form v-if="activeTab === 'edit' && auth.state.user" class="profile-edit-card stack-form ui-panel ui-stack" @submit.prevent="saveProfile">
        <div class="profile-edit-heading ui-between">
          <div><strong>编辑个人信息</strong><span>更新你的校园名片</span></div>
          <ImageUpload v-model="profileForm.avatar" />
        </div>
        <label>
          <span>昵称</span>
          <input v-model="profileForm.nickname" maxlength="30" required />
        </label>
        <label>
          <span>个人简介</span>
          <textarea v-model="profileForm.bio" rows="3" maxlength="200" placeholder="介绍一下你的兴趣或最近在做的事" />
        </label>
        <button class="button button-small ui-inline" :disabled="saving">
          <Save :size="15" /> {{ saving ? "保存中…" : "保存资料" }}
        </button>
      </form>
    </template>
  </div>
</template>

