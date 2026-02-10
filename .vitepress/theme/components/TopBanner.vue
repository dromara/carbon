<template>
  <div v-if="shouldShowBanner" class="top-banner">
    <div class="banner-content">
      <svg class="banner-icon" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
        <path d="M18 10h-1.26A8 8 0 1 0 9 20h9a5 5 0 0 0 0-10z"></path>
      </svg>
      <span class="banner-text">
        雨云——新一代云服务提供商，用稳定和性价比，助力您快速上云
      </span>
      <a :href="bannerLink" class="banner-link" rel="noopener noreferrer" target="_blank">
        点击前往 →
      </a>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {computed} from 'vue'
import {useRoute} from 'vitepress'

const route = useRoute()

const shouldShowBanner = computed(() => {
  const path = route.path
  // 只在中文版显示，且排除首页
  return path.startsWith('/zh/') && path !== '/zh/' && path !== '/zh/index.html'
})

const bannerLink = 'https://www.rainyun.com/gopkg_?s=carbon'
</script>

<style scoped>
.top-banner {
  margin: -24px -24px 16px;
  padding: 12px 24px;
  background-color: var(--vp-c-bg-soft);
  border: 1px solid var(--vp-c-divider);
  border-radius: 8px;
  transition: all 0.3s;
}

.top-banner:hover {
  border-color: var(--vp-c-brand-1);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.banner-content {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  flex-wrap: wrap;
}

.banner-icon {
  width: 18px;
  height: 18px;
  color: var(--vp-c-brand-1);
  flex-shrink: 0;
}

.banner-text {
  color: var(--vp-c-text-1);
  font-size: 14px;
  line-height: 1.5;
}

.banner-link {
  color: var(--vp-c-brand-1);
  font-size: 14px;
  font-weight: 500;
  text-decoration: none;
  padding: 4px 12px;
  border-radius: 4px;
  transition: all 0.2s;
  white-space: nowrap;
  background-color: var(--vp-c-brand-soft);
}

.banner-link:hover {
  background-color: var(--vp-c-brand-1);
  color: #ffffff;
}

@media (max-width: 768px) {
  .top-banner {
    margin: -24px -16px 16px;
    padding: 10px 16px;
  }
  
  .banner-content {
    gap: 8px;
  }
  
  .banner-icon {
    width: 16px;
    height: 16px;
  }
  
  .banner-text {
    font-size: 13px;
    flex: 1;
    min-width: 100%;
    text-align: center;
  }
  
  .banner-link {
    font-size: 13px;
    margin: 0 auto;
  }
}

@media (max-width: 960px) {
  .top-banner {
    margin: -24px -24px 16px;
  }
}
</style>
