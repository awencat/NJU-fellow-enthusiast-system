<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { api } from "@/api";
import type { ExchangePost, MatePost, Paginated } from "@/types";
import { relativeTime } from "@/utils/date";
import LoadingBlock from "@/components/LoadingBlock.vue";

const router = useRouter();
// 每次进入首页时抽取一次，页面内的数据更新不会切换猫咪。
const mascots = [
  "/images/mascots/watermelon.png",
  "/images/mascots/confident.png",
  "/images/mascots/meditating.png",
];
const homeMascot = mascots[Math.floor(Math.random() * mascots.length)];
const search = ref("");
const posts = ref<ExchangePost[]>([]);
const mates = ref<MatePost[]>([]);
const loading = ref(true);
const error = ref(false);

async function loadHome() {
  loading.value = true;
  error.value = false;
  try {
    const [exchange, invitations] = await Promise.all([
      api<Paginated<ExchangePost>>("/api/exchange/posts?pageSize=3"),
      api<Paginated<MatePost>>("/api/mates?pageSize=3"),
    ]);
    posts.value = exchange.items;
    mates.value = invitations.items;
  } catch {
    error.value = true;
  } finally {
    loading.value = false;
  }
}

function searchMates() {
  router.push({ path: "/mates", query: { q: search.value.trim() || undefined } });
}
onMounted(loadHome);
</script>

<template>
  <div class="home-page">
    <header class="atlas-cover">
      <div class="atlas-cover-copy">
        <span class="atlas-imprint">南大 · 社区图鉴</span>
        <h1>校园同好<span>图鉴</span></h1>
        <nav aria-label="本页导航"><a href="#field-posts">同好交流 ↓</a><a href="#field-mates">搭子招募 ↓</a></nav>
      </div>
      <div class="atlas-cover-art" aria-hidden="true"><img :src="homeMascot" alt="" /></div>
    </header>
    <div class="folio-columns">
      <aside class="folio-directory">
        <h2>栏目目录</h2>
        <nav aria-label="栏目目录">
          <RouterLink to="/exchange"><img src="/images/atlas/journal.svg" alt="" />同好交流</RouterLink>
          <RouterLink to="/mates"><img src="/images/atlas/compass.svg" alt="" />找搭子</RouterLink>
          <RouterLink to="/messages"><img src="/images/atlas/mail.svg" alt="" />私信</RouterLink>
        </nav>
        <h2>按兴趣找搭子</h2>
        <form class="folio-search" @submit.prevent="searchMates">
          <input v-model="search" type="search" aria-label="搜索搭子" placeholder="搜索活动或地点" />
          <button>搜索</button>
        </form>
        <div class="folio-tags">
          <RouterLink v-for="tag in ['周末', '学习', '运动']" :key="tag" :to="{ path: '/mates', query: { tag } }">{{ tag }}</RouterLink>
        </div>
      </aside>
      <div class="folio-content">
        <LoadingBlock v-if="loading" />
        <p v-else-if="error" role="status">暂时无法读取社区动态。<button @click="loadHome">重新加载</button></p>
        <template v-else>
          <section id="field-posts" class="folio-section">
            <header><h2><img src="/images/atlas/journal.svg" alt="" />同好交流</h2><RouterLink to="/exchange/new">写分享 ↗</RouterLink></header>
            <ul>
              <li v-for="(post, index) in posts" :key="post.id" :class="{ 'atlas-lead': index === 0 && post.imageUrl }">
                <img v-if="index === 0 && post.imageUrl" class="atlas-lead-cover" :src="post.imageUrl" alt="" />
                <small>{{ post.kind === 'news' ? '校园新闻' : '同好分享' }}</small>
                <RouterLink :to="'/exchange/' + post.id">{{ post.title }}</RouterLink>
                <span>{{ post.author.nickname }} · {{ relativeTime(post.createdAt) }}</span>
              </li>
            </ul>
            <p v-if="!posts.length">暂无分享。</p>
            <RouterLink class="folio-more" to="/exchange">全部交流 →</RouterLink>
          </section>
          <section id="field-mates" class="folio-section">
            <header><h2><img src="/images/atlas/compass.svg" alt="" />搭子招募</h2><RouterLink to="/mates/new">发布邀请 ↗</RouterLink></header>
            <ul>
              <li v-for="mate in mates" :key="mate.id">
                <small>{{ mate.tags[0] || '招募' }}</small>
                <RouterLink :to="'/mates/' + mate.id">{{ mate.title }}</RouterLink>
                <span>{{ mate.location || mate.author.nickname }} · {{ mate.duration }}</span>
              </li>
            </ul>
            <p v-if="!mates.length">暂无搭子邀请。</p>
            <RouterLink class="folio-more" to="/mates">全部邀请 →</RouterLink>
          </section>
        </template>
      </div>
    </div>
    <details class="mascot-note">
      <summary>🐾 校猫手记</summary>
      <figure><img src="/images/campus-mascot.png" alt="身穿红袍、头戴红帽的校园猫咪" /><figcaption>今日值班 · 搭子招募处</figcaption></figure>
    </details>
  </div>
</template>
