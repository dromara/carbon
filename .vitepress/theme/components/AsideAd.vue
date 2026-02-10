<template>
  <div v-if="isZhLocale" class="aside-ad">
    <div class="ad-list">
      <a 
        v-for="(ad, index) in adList" 
        :key="index"
        :href="ad.link" 
        class="ad-item"
        rel="noopener noreferrer" 
        target="_blank"
      >
        <div class="ad-item-header">
          <span class="ad-title">{{ ad.title }}</span>
        </div>
        <div class="ad-desc">{{ ad.desc }}</div>
        <div class="ad-footer">
          <span class="ad-action">点击购买</span>
        </div>
      </a>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {computed} from 'vue'
import {useRoute} from 'vitepress'

const route = useRoute()

const isZhLocale = computed(() => {
  return route.path.startsWith('/zh/')
})

// 配置文字广告列表
const adList = [
  {
    title: '雨云',
    desc: 'KVM高配版，4核 8G 300M，仅需68元/月起，半年付八折，年付七折，七天无理由退订（送5折券）',
    link: 'https://www.rainyun.com/gopkg_?s=carbon',
  },
  {
    title: '林枫云',
    desc: '9950X高防，4核 8G 10M，仅需168元/月，续费同价',
    link: 'https://www.dkdun.cn/aff/NVHPHCEF', // 替换为实际链接
  },
  {
    title: '莱卡云',
    desc: '浙江电信云，2核 4G 10M，仅需38元/月，续费同价',
    link: 'https://www.lcayun.com/aff/MBVNVNFX', // 替换为实际链接
  }
]
</script>

<style scoped>
.aside-ad {
  margin-top: 16px;
}

.ad-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.ad-item {
  display: block;
  text-decoration: none;
  padding: 10px;
  background-color: var(--vp-c-bg-soft);
  border: 1px solid var(--vp-c-divider);
  border-radius: 8px;
  transition: all 0.3s;
}

.ad-item:hover {
  border-color: var(--vp-c-brand-1);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  transform: translateY(-2px);
}

.ad-item-header {
  margin-bottom: 6px;
}

.ad-title {
  color: var(--vp-c-text-1);
  font-size: 14px;
  font-weight: 600;
  line-height: 1.4;
}

.ad-desc {
  color: var(--vp-c-text-2);
  font-size: 12px;
  line-height: 1.5;
  margin-bottom: 6px;
}

.ad-footer {
  display: flex;
  align-items: center;
}

.ad-action {
  color: var(--vp-c-brand-1);
  font-size: 12px;
  font-weight: 500;
}
</style>
