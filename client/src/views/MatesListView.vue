<script setup lang="ts">
import { ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { Plus } from "lucide-vue-next";
import { api } from "@/api";
import { auth } from "@/stores/auth";
import { notify } from "@/stores/toast";
import TagInput from "@/components/TagInput.vue";
import type { MatePost, Paginated } from "@/types";
import MateCard from "@/components/MateCard.vue";
import EmptyState from "@/components/EmptyState.vue";
import LoadingBlock from "@/components/LoadingBlock.vue";
import PaginationBar from "@/components/PaginationBar.vue";
const route = useRoute();
const router = useRouter();
const mates = ref<MatePost[]>([]);
const loading = ref(true);
const q = ref("");
const tags = ref<string[]>([]);
const page = ref(1);
const total = ref(0);
const pageSize = 9;
let request = 0;
watch(() => route.query, async query => {
 q.value = String(query.q ?? "");
 tags.value = (Array.isArray(query.tag) ? query.tag : query.tag ? [query.tag] : []).filter((t): t is string => typeof t === "string");
 page.value = Math.max(1, Number(query.page) || 1);
 const params = new URLSearchParams({ page: String(page.value), pageSize: String(pageSize) });
 if(q.value) params.set("q", q.value);
 tags.value.forEach(tag => params.append("tag",tag));
 const current = ++request;
 loading.value = true;
 try {
  const data = await api<Paginated<MatePost>>(`/api/mates?${params}`);
  if(current === request) {mates.value = data.items; total.value = data.total;}
 } catch(error) {
  if(current === request) {mates.value = []; total.value = 0; notify(error instanceof Error ? error.message : "读取失败", "error");}
 } finally {if(current === request) loading.value = false;}
}, { immediate: true });
function applySearch() {
 router.replace({query: {...(q.value.trim() ? {q: q.value.trim()} : {}), ...(tags.value.length ? {tag: tags.value} : {})}});
}
function clearFilters() { q.value = ""; tags.value = []; applySearch(); }
function changePage(value: number) {
 router.replace({query: {...route.query,page: value}});
 window.scrollTo({top:0, behavior:"smooth"});
}
</script>

<template>
  <div class="page-shell mates-page listing-page">


    <form class="tag-search-panel" role="search" aria-label="筛选搭子" @submit.prevent="applySearch">
      <div class="search-input-row">
      <label class="search-keyword"><span>关键词</span><input v-model="q" type="search" placeholder="搜索标题或正文" /></label>
      <TagInput v-model="tags" />
      <RouterLink
        class="search-publish button warm-button ui-inline"
        :to="auth.state.user ? '/mates/new' : '/login?redirect=/mates/new'"
      >
        <Plus :size="18" /> 发布搭子邀请
      </RouterLink>
      </div>
      <div class="search-actions">
        <button type="submit" class="button">筛选</button>
        <button type="button" class="button secondary" @click="clearFilters">清空</button>
        <small>标签支持部分文字匹配，匹配标签越多越靠前。</small>
      </div>
    </form>

    <div class="list-summary ui-between">
      <strong>{{ total }} 个正在寻找同伴的邀请</strong>
      <span>请尊重彼此边界，线下见面优先选择公共场所</span>
    </div>

    <LoadingBlock v-if="loading" />
    <template v-else>
      <div v-if="mates.length" class="mate-grid mates-grid">
        <MateCard v-for="mate in mates" :key="mate.id" :mate="mate" />
      </div>
      <EmptyState
        v-else
        title="暂时没有匹配的搭子"
        text="换个关键词，或者发布一个属于你的邀请。"
      >
        <RouterLink class="button button-small ui-inline" to="/mates/new">发布邀请</RouterLink>
      </EmptyState>
      <PaginationBar
        :page="page"
        :total="total"
        :page-size="pageSize"
        @change="changePage"
      />
    </template>
  </div>
</template>

