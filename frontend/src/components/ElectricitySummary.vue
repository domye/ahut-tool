<!-- frontend/src/components/ElectricitySummary.vue -->
// 电费汇总统计卡片组件

<script setup lang="ts">
import { computed } from 'vue'
import { models } from '../../wailsjs/go/models'

interface Props {
  data: models.IMSData | null
  title: string
  type: 'total' | 'used' | 'remain'
}

const props = defineProps<Props>()

const isLowBalance = computed(() => {
  return (props.data?.RemainAmp || 0) < 10 && props.type === 'remain'
})

const getValue = computed(() => {
  if (props.type === 'total') return props.data?.AllAmp || 0
  if (props.type === 'used') return props.data?.UsedAmp || 0
  if (props.type === 'remain') return props.data?.RemainAmp || 0
  return 0
})

const getIconClass = computed(() => {
  switch (props.type) {
    case 'total': return 'total'
    case 'used': return 'used'
    case 'remain': return 'remain'
    default: return 'total'
  }
})
</script>

<template>
  <div class="statistic-card">
    <div class="statistic-content">
      <div class="statistic-title">{{ title }}</div>
      <div class="statistic-value" :class="{ 'low-balance': isLowBalance }">
        {{ getValue }} kWh
      </div>
    </div>
  </div>
</template>

<style scoped>
.statistic-card {
  background-color: #e0e5ec;
  border-radius: 20px;
  padding: 24px;
  display: flex;
  align-items: center;
  box-shadow: 8px 8px 16px 0 rgba(163, 177, 198, 0.6), -8px -8px 16px 0 rgba(255, 255, 255, 0.9);
  height: 100%;
  min-height: 100px;
  transition: all 0.3s ease;
}

@media (max-width: 768px) {
  .statistic-card {
    padding: 20px 16px;
    min-height: 90px;
  }
}

.statistic-card:hover {
  transform: translateY(-4px);
  box-shadow: 10px 10px 20px 0 rgba(163, 177, 198, 0.7), -10px -10px 20px 0 rgba(255, 255, 255, 1);
}

.statistic-content {
  flex: 1;
}

.statistic-title {
  font-size: 0.95rem;
  color: #718096;
  margin-bottom: 8px;
  font-weight: 500;
}

@media (max-width: 768px) {
  .statistic-title {
    font-size: 0.85rem;
    margin-bottom: 6px;
  }
}

.statistic-value {
  font-size: 1.8rem;
  font-weight: 700;
  color: #4a5568;
  line-height: 1.2;
}

@media (max-width: 768px) {
  .statistic-value {
    font-size: 1.4rem;
  }
}

.statistic-value.low-balance {
  color: #cf1322;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.6;
  }
}
</style>
