<script setup lang="ts">
import { Bookmark, Heart, MessageCircle } from "lucide-vue-next";
import Avatar from "./Avatar.vue";
import type { ExchangePost } from "@/types";
import { relativeTime } from "@/utils/date";

defineProps<{
  post: ExchangePost;
  featured?: boolean;
}>();
</script>

<template>
  <article class="post-card ui-panel ui-card" :class="{ 'post-card-featured': featured }">
    <RouterLink :to="`/exchange/${post.id}`" class="post-card-link">
      <div v-if="post.imageUrl" class="post-cover">
        <img :src="post.imageUrl" :alt="post.title" loading="lazy" />
      </div>
      <div v-else class="post-cover post-cover-placeholder ui-stack">
        <span>“</span>
        <p>生活的小片段</p>
      </div>
      <div class="post-card-body ui-stack">
        <div class="post-kind-row ui-between">
          <span class="eyebrow ui-inline">
            <span v-for="tag in post.tags" :key="tag" class="tag">{{ tag }}</span>
          </span>
          <span>{{ relativeTime(post.createdAt) }}</span>
        </div>
        <h3>{{ post.title }}</h3>
        <p class="ui-summary">{{ post.content }}</p>
        <div class="post-author-row ui-between">
          <div class="inline-user ui-row">
            <Avatar :user="post.author" :size="28" />
            <span>{{ post.author.nickname }}</span>
          </div>
          <div class="post-metrics">
            <span :class="{ active: post.liked }">
              <Heart :size="15" /> {{ post.likeCount }}
            </span>
            <span>
              <MessageCircle :size="15" /> {{ post.commentCount }}
            </span>
            <span>
              <Bookmark :size="15" /> {{ post.favoriteCount }}
            </span>
          </div>
        </div>
      </div>
    </RouterLink>
  </article>
</template>

