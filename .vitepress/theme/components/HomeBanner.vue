<template>
  <div v-if="shouldShowBanner" class="home-banner">
    <span class="banner-text">雨云——新一代云服务提供商，用稳定和性价比，助力您快速上云</span>
    <a :href="bannerLink" class="banner-link" rel="noopener noreferrer" target="_blank">
      点击前往 →
    </a>
  </div>
</template>

<script lang="ts" setup>
import {computed} from 'vue'
import {useData} from 'vitepress'

const { page } = useData()

const shouldShowBanner = computed(() => {
  // 使用 page.relativePath 来判断，这个在首次加载时就是准确的
  const relativePath = page.value.relativePath
  // 只在中文版首页显示
  return relativePath === 'zh/index.md' || relativePath.startsWith('zh/index')
})

const bannerLink = 'https://www.rainyun.com/gocarbon_?s=web'
</script>

<style scoped>
.home-banner {
  margin-top: 16px;
  margin-bottom: -8px;
  padding: 0;
  text-align: left;
  font-size: 14px;
  line-height: 1.8;
}

.banner-text {
  color: var(--vp-c-text-2);
  font-weight: 400;
}

.banner-link {
  color: var(--vp-c-brand-1);
  font-weight: 500;
  text-decoration: none;
  padding: 4px 12px;
  border-radius: 6px;
  transition: all 0.2s;
  white-space: nowrap;
  margin-left: 8px;
  display: inline-block;
}

.banner-link:hover {
  background-color: var(--vp-c-brand-1);
  color: #ffffff;
  transform: translateX(2px);
}

@media (max-width: 768px) {
  .home-banner {
    font-size: 13px;
    margin-top: 12px;
  }
  
  .banner-text {
    font-size: 13px;
  }
  
  .banner-link {
    margin-left: 8px;
    font-size: 13px;
  }
}
</style>
