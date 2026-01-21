<template>
  <div class="electricity-page">
    <!-- 电费信息卡片 -->
    <a-row :gutter="[24, 24]" class="summary-row">
      <a-col :xs="24" :sm="12" :md="8" :lg="8" :xl="8">
        <div class="statistic-card">
          <div class="statistic-icon">
            <ThunderboltOutlined />
          </div>
          <div class="statistic-content">
            <div class="statistic-title">总电量</div>
            <div class="statistic-value">{{ electricityStore.electricity?.AllAmp || 0 }} kWh</div>
          </div>
        </div>
      </a-col>
      <a-col :xs="24" :sm="12" :md="8" :lg="8" :xl="8">
        <div class="statistic-card">
          <div class="statistic-icon">
            <LoginOutlined />
          </div>
          <div class="statistic-content">
            <div class="statistic-title">已用电量</div>
            <div class="statistic-value">{{ electricityStore.electricity?.UsedAmp || 0 }} kWh</div>
          </div>
        </div>
      </a-col>
      <a-col :xs="24" :sm="12" :md="8" :lg="8" :xl="8">
        <div class="statistic-card">
          <div class="statistic-icon">
            <CheckCircleOutlined />
          </div>
          <div class="statistic-content">
            <div class="statistic-title">剩余电量</div>
            <div class="statistic-value" :class="{ 'low-balance': isLowBalance }">
              {{ electricityStore.electricity?.RemainAmp || 0 }} kWh
            </div>
          </div>
        </div>
      </a-col>
    </a-row>


    <!-- 查询条件 -->
    <div class="filter-section">
      <a-form layout="inline" class="filter-form">
        <a-form-item label="校区">
          <a-select
            v-model:value="electricityStore.xiaoqu"
            placeholder="请选择校区"
            style="width: 200px"
            allow-clear
            :options="xiaoquOptions"
            popup-class-name="neumorphic-select-dropdown"
          />
        </a-form-item>
        <a-form-item label="楼栋名称">
          <a-input
            v-model:value="electricityStore.ld_Name"
            placeholder="请输入楼栋名称"
            style="width: 200px"
            allow-clear
          />
        </a-form-item>
        <a-form-item label="楼栋ID">
          <a-input
            v-model:value="electricityStore.ld_Id"
            placeholder="请输入楼栋ID"
            style="width: 200px"
            allow-clear
          />
        </a-form-item>
        <a-form-item label="房间号">
          <a-input
            v-model:value="electricityStore.Room_No"
            placeholder="请输入房间号"
            style="width: 200px"
            allow-clear
          />
        </a-form-item>
        <a-form-item label="类型">
          <a-select
            v-model:value="electricityStore.etype"
            placeholder="请选择类型"
            style="width: 200px"
            allow-clear
            :options="etypeOptions"
            popup-class-name="neumorphic-select-dropdown"
          />
        </a-form-item>
        <a-form-item>
          <a-space :size="20">
            <a-button
              type="primary"
              @click="handleSearch"
              :loading="electricityStore.loading"
              class="neumorphic-button query-button"
            >
              <template #icon>
                <SearchOutlined />
              </template>
              查询
            </a-button>
            <a-button
              @click="handleReset"
              class="neumorphic-button reset-button"
            >
              <template #icon>
                <RedoOutlined />
              </template>
              重置
            </a-button>
          </a-space>
        </a-form-item>
      </a-form>
      <a-alert
        v-if="electricityStore.error"
        :message="electricityStore.error"
        type="error"
        show-iconw
        closable
        @close="electricityStore.error = ''"
        style="margin-top: 16px"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useElectricityStore } from '../store/electricity'
import { ThunderboltOutlined, LoginOutlined, CheckCircleOutlined, SearchOutlined, RedoOutlined, LogoutOutlined } from '@ant-design/icons-vue'

const electricityStore = useElectricityStore()
const userId = ref('')
const password = ref('')

// 判断是否为低电量
const isLowBalance = computed(() => {
  return (electricityStore.electricity?.RemainAmp || 0) < 10
})

// 校区选项
const xiaoquOptions = [
  { value: 'NewS', label: '新校区' },
  { value: 'OldS', label: '老校区' }
]

// 类型选项
const etypeOptions = [
  { value: 'K', label: '空调' },
  { value: 'L', label: '照明' }
]


// 查询处理
function handleSearch() {
  electricityStore.fetchElectricity()
    .catch((error) => {
      console.error('查询失败:', error)
    })
}

// 重置处理
function handleReset() {
  electricityStore.resetFilters()
}

</script>

<style scoped>
.electricity-page {
  padding: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.summary-row {
  margin-bottom: 32px;
}

.statistic-card {
  background-color: #e0e5ec;
  border-radius: 20px;
  padding: 24px;
  display: flex;
  align-items: center;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
}

.statistic-icon {
  font-size: 2.5rem;
  margin-right: 16px;
  color: #1890ff;
}

.statistic-content {
  flex: 1;
}

.statistic-title {
  font-size: 1rem;
  color: #4a5568;
  margin-bottom: 8px;
}

.statistic-value {
  font-size: 1.5rem;
  color: #2d3748;
  font-weight: 600;
}

.statistic-value.low-balance {
  color: #cf1322;
}

.login-section,
.filter-section {
  margin-bottom: 32px;
}

.neumorphic-card {
  background-color: #e0e5ec;
  border-radius: 20px;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
  border: none;
}

.login-form {
  max-width: 400px;
  margin: 0 auto;
}

.filter-form {
  background-color: #e0e5ec;
  border-radius: 20px;
  padding: 24px;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
}

.neumorphic-button {
  background-color: #e0e5ec;
  border: none;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
  transition: all 0.3s ease;
}

.neumorphic-button:hover {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8);
  transform: translateY(-2px);
}

.neumorphic-button:active {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.8);
  transform: translateY(0);
}

.query-button {
  color: #1890ff;
}

.reset-button {
  color: #faad14;
}

.logout-button {
  color: #cf1322;
}
</style>
