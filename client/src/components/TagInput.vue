<script setup lang="ts">
import { ref } from "vue";
const tags = defineModel<string[]>({ required: true });
const draft = ref("");
const error = ref("");
function add() {
 const additions = draft.value.split(/[,，\s]+/).map(t => t.trim()).filter(Boolean);
 error.value = "";
 if(additions.some(tag => tag.length > 16)) {error.value="每个标签最多16个字符"; return;}
 const next = [...tags.value];
 for (const tag of additions) {
  if (tag.length > 16 || next.some(t => t.toLocaleLowerCase() === tag.toLocaleLowerCase())) continue;
  if (next.length >= 6) {error.value="最多6个标签，请先移除已有标签"; return;}
  next.push(tag);
 }
 tags.value = next;
 draft.value = "";
}
</script>
<template>
 <fieldset class="tag-picker">
  <legend>标签</legend>
  <div class="custom-tag">
   <input v-model="draft" aria-label="输入标签" placeholder="输入标签后按回车添加" @keydown.enter.prevent="add" @blur="add" />
   <button type="button" class="button secondary button-small" @click="add">添加</button>
  </div>
  <p v-if="error" role="alert" class="danger-text">{{ error }}</p>
  <div class="selected-tags ui-wrap">
   <span v-for="tag in tags" :key="tag" class="tag ui-inline">{{ tag }}
    <button type="button" :aria-label="`移除标签 ${tag}`" @click="tags = tags.filter(t => t !== tag)">×</button>
   </span>
  </div>
 </fieldset>
</template>
