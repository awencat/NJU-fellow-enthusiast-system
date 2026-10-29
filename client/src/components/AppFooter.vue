<script setup lang="ts">
import { Github } from "lucide-vue-next";

const repoUrl = "https://github.com/awencat/NJU-fellow-enthusiast-system";
</script>

<template>
  <footer class="site-footer">
    <div class="footer-inner">
      <div class="footer-main-about">
        <nav class="footer-about ui-row">
          <RouterLink to="/faq" class="for_routerlink">常见问题</RouterLink>
        </nav>

      </div>
      <div class="footer-main">
          <nav class="footer-links ui-row">
            <a href="https://github.com/Methyl1644/EL_CampusMate" target="_blank" rel="noreferrer" >
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-github size-4" aria-hidden="true"><path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4">
              </path><path d="M9 18c-4.51 2-5-2-7-2"></path></svg>项目仓库</a>
            <a href="https://github.com/Methyl1644/EL_CampusMate/issues" target="_blank" rel="noreferrer" ><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-circle-help size-4" aria-hidden="true"><circle cx="12" cy="12" r="10"></circle>
              <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"></path>
              <path d="M12 17h.01"></path></svg>问题反馈</a>
          </nav>
      </div>
    </div>
    <div class="footer-thanks">
      <p class="text-thanks">
           NJU-fellow-enthusiast-system
      </p>
    </div>
  </footer>
</template>
