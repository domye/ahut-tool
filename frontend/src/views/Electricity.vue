<template>
  <div class="electricity-page">
    <!-- 宿舍信息卡片 -->
    <a-row class="info-row">
      <a-col :span="24">
        <div class="info-card">

          <div class="info-content">
            <div class="info-value">
              {{ electricityStore.ld_Name && electricityStore.Room_No ? `${electricityStore.ld_Name}${electricityStore.Room_No}号` : '请先点击设置' }}
            </div>
          </div>
          <div class="info-icon clickable-icon" @click="showSettingModal = true" aria-label="设置宿舍信息">
            <SettingOutlined/>
          </div>
        </div>
      </a-col>
    </a-row>

    <!-- 电费信息卡片 - 空调 -->
    <div class="section-title">
      <div class="title-icon">
        <ThunderboltOutlined/>
      </div>
      <span>空调电费</span>
    </div>
    <a-row :gutter="[24, 24]" class="summary-row">
      <a-col :xs="24" :sm="12" :md="8" :lg="8" :xl="8">
        <div class="statistic-card">
          <div class="statistic-icon total">
            <ThunderboltOutlined/>
          </div>
          <div class="statistic-content">
            <div class="statistic-title">总电量</div>
            <div class="statistic-value">{{ electricityStore.airConditioning?.AllAmp || 0 }} kWh</div>
          </div>
        </div>
      </a-col>
      <a-col :xs="24" :sm="12" :md="8" :lg="8" :xl="8">
        <div class="statistic-card">
          <div class="statistic-icon used">
            <LoginOutlined/>
          </div>
          <div class="statistic-content">
            <div class="statistic-title">已用电量</div>
            <div class="statistic-value">{{ electricityStore.airConditioning?.UsedAmp || 0 }} kWh</div>
          </div>
        </div>
      </a-col>
      <a-col :xs="24" :sm="12" :md="8" :lg="8" :xl="8">
        <div class="statistic-card">
          <div class="statistic-icon remain">
            <CheckCircleOutlined/>
          </div>
          <div class="statistic-content">
            <div class="statistic-title">剩余电量</div>
            <div class="statistic-value" :class="{ 'low-balance': isACLowBalance }">
              {{ electricityStore.airConditioning?.RemainAmp || 0 }} kWh
            </div>
          </div>
        </div>
      </a-col>
    </a-row>

    <!-- 电费信息卡片 - 房间 -->
    <div class="section-title">
      <div class="title-icon">
        <HomeOutlined/>
      </div>
      <span>房间电费</span>
    </div>
    <a-row :gutter="[24, 24]" class="summary-row">
      <a-col :xs="24" :sm="12" :md="8" :lg="8" :xl="8">
        <div class="statistic-card">
          <div class="statistic-icon total">
            <ThunderboltOutlined/>
          </div>
          <div class="statistic-content">
            <div class="statistic-title">总电量</div>
            <div class="statistic-value">{{ electricityStore.electricity?.AllAmp || 0 }} kWh</div>
          </div>
        </div>
      </a-col>
      <a-col :xs="24" :sm="12" :md="8" :lg="8" :xl="8">
        <div class="statistic-card">
          <div class="statistic-icon used">
            <LoginOutlined/>
          </div>
          <div class="statistic-content">
            <div class="statistic-title">已用电量</div>
            <div class="statistic-value">{{ electricityStore.electricity?.UsedAmp || 0 }} kWh</div>
          </div>
        </div>
      </a-col>
      <a-col :xs="24" :sm="12" :md="8" :lg="8" :xl="8">
        <div class="statistic-card">
          <div class="statistic-icon remain">
            <CheckCircleOutlined/>
          </div>
          <div class="statistic-content">
            <div class="statistic-title">剩余电量</div>
            <div class="statistic-value" :class="{ 'low-balance': isRoomLowBalance }">
              {{ electricityStore.electricity?.RemainAmp || 0 }} kWh
            </div>
          </div>
        </div>
      </a-col>
    </a-row>


    <!-- 设置弹窗 -->
    <DormSetting
        v-model:open="showSettingModal"
        :loading="electricityStore.loading"
        :initial-data="settingForm"
        @save="handleSettingSave"
    />

    <!-- 错误提示 -->
    <a-alert
        v-if="electricityStore.error"
        :message="electricityStore.error"
        type="error"
        show-icon
        closable
        @close="electricityStore.error = ''"
        style="margin-top: 16px"
    />
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, ref, watch} from 'vue'
import {useElectricityStore} from '../store/electricity'
import {
  CheckCircleOutlined,
  HomeOutlined,
  LoginOutlined,
  SettingOutlined,
  ThunderboltOutlined
} from '@ant-design/icons-vue'
import DormSetting from '../components/DormSetting.vue'

const electricityStore = useElectricityStore()

// 设置弹窗显示状态
const showSettingModal = ref(false)

// 设置表单
const settingForm = ref({
  campus: electricityStore.xiaoqu || 'NewS',
  buildingId: electricityStore.ld_Id || '',
  buildingName: electricityStore.ld_Name || '',
  roomId: electricityStore.Room_No || ''
})

// 监听store中宿舍信息变化，更新设置表单
watch(() => [electricityStore.xiaoqu, electricityStore.ld_Id, electricityStore.ld_Name, electricityStore.Room_No], ([xiaoqu, ldId, ldName, roomNo]) => {
  settingForm.value = {
    campus: xiaoqu || 'NewS',
    buildingId: ldId || '',
    buildingName: ldName || '',
    roomId: roomNo || ''
  }
})

// 判断是否为低电量
const isACLowBalance = computed(() => {
  return (electricityStore.airConditioning?.RemainAmp || 0) < 10
})

const isRoomLowBalance = computed(() => {
  return (electricityStore.electricity?.RemainAmp || 0) < 10
})

// 校区选项
const xiaoquOptions = [
  {value: 'NewS', label: '新校区'},
  {value: 'OldS', label: '老校区'}
]

// 获取校区名称
function getCampusName(campus: string) {
  const option = xiaoquOptions.find(opt => opt.value === campus)
  return option ? option.label : campus
}

// 保存设置
async function handleSettingSave(data: {
  campus: string
  buildingId: string
  buildingName: string
  roomId: string
}) {
  try {
    await electricityStore.saveDormConfig(
        data.campus,
        data.buildingId,
        data.buildingName,
        data.roomId
    )
    // 保存成功后查询电费
    await electricityStore.fetchElectricity()
    showSettingModal.value = false
  } catch (error) {
    console.error('保存设置失败:', error)
  }
}

// 页面加载时初始化
onMounted(async () => {
  try {
    // 加载保存的宿舍配置
    await electricityStore.loadDormConfig()
    // 配置加载成功后查询电费
    await electricityStore.fetchElectricity()
  } catch (error) {
    console.error('初始化失败:', error)
  }
})
</script>

<style scoped>
.electricity-page {
  padding: 32px 24px;
  background-color: #e0e5ec;
  min-height: calc(100vh - 64px);
}

@media (max-width: 768px) {
  .electricity-page {
    padding: 20px 16px;
  }
}

.section-title {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 1.3rem;
  font-weight: 600;
  color: #4a5568;
  margin-bottom: 20px;
  margin-top: 32px;
}

@media (max-width: 768px) {
  .section-title {
    font-size: 1.1rem;
    margin-bottom: 16px;
    margin-top: 24px;
  }
}

.title-icon {
  font-size: 1.5rem;
  color: #409eff;
}

@media (max-width: 768px) {
  .title-icon {
    font-size: 1.3rem;
  }
}

.info-row {
  margin-bottom: 40px;
}

@media (max-width: 768px) {
  .info-row {
    margin-bottom: 24px;
  }
}

.info-card {
  background-color: #e0e5ec;
  border-radius: 20px;
  padding: 24px;
  box-shadow: 8px 8px 16px 0 rgba(163, 177, 198, 0.6), -8px -8px 16px 0 rgba(255, 255, 255, 0.9);
  display: flex;
  align-items: center;
  gap: 20px;
  height: 100%;
  min-height: 50px;
  transition: all 0.3s ease;
}

@media (max-width: 768px) {
  .info-card {
    padding: 20px 16px;
    min-height: 50px;
    gap: 16px;
  }
}

.info-card:hover {
  transform: translateY(-4px);
  box-shadow: 10px 10px 20px 0 rgba(163, 177, 198, 0.7), -10px -10px 20px 0 rgba(255, 255, 255, 1);
}


.info-icon {
  font-size: 1.8rem;
  color: #409eff;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 56px;
  height: 56px;
  background-color: #e0e5ec;
  border-radius: 16px;
  box-shadow: 4px 4px 8px 0 rgba(163, 177, 198, 0.4), -4px -4px 8px 0 rgba(255, 255, 255, 0.9);
}

@media (max-width: 768px) {
  .info-icon {
    width: 48px;
    height: 48px;
    font-size: 1.5rem;
  }
}

/* 可点击图标样式 */
.clickable-icon {
  cursor: pointer;
  transition: all 0.3s ease;
}

.clickable-icon:hover {
  transform: rotate(90deg);
  background-color: rgba(64, 158, 255, 0.1);
  box-shadow: 4px 4px 8px 0 rgba(64, 158, 255, 0.2), -4px -4px 8px 0 rgba(255, 255, 255, 0.9);
}

.clickable-icon:active {
  transform: rotate(90deg) scale(0.95);
}

.action-card .info-icon {
  color: #409eff;
  background-color: rgba(64, 158, 255, 0.1);
  box-shadow: 4px 4px 8px 0 rgba(64, 158, 255, 0.2), -4px -4px 8px 0 rgba(255, 255, 255, 0.9);
  width: 56px;
  height: 56px;
  font-size: 2rem;
}

.info-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

@media (max-width: 768px) {
  .info-content {
    gap: 4px;
  }
}

.info-value {
  font-size: 1.2rem;
  color: #4a5568;
  font-weight: 600;
}

@media (max-width: 768px) {
  .info-value {
    font-size: 1rem;
  }
}

.summary-row {
  margin-bottom: 24px;
}

@media (max-width: 768px) {
  .summary-row {
    margin-bottom: 16px;
  }
}

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

.statistic-icon {
  font-size: 2rem;
  margin-right: 20px;
  color: #4a5568;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 56px;
  height: 56px;
  border-radius: 50%;
  transition: all 0.3s ease;
}

@media (max-width: 768px) {
  .statistic-icon {
    width: 48px;
    height: 48px;
    font-size: 1.5rem;
    margin-right: 16px;
  }
}

.statistic-icon.total {
  color: #409eff;
  background-color: rgba(64, 158, 255, 0.1);
  box-shadow: 4px 4px 8px 0 rgba(64, 158, 255, 0.2), -4px -4px 8px 0 rgba(255, 255, 255, 0.9);
}

.statistic-icon.used {
  color: #e6a23c;
  background-color: rgba(230, 162, 60, 0.1);
  box-shadow: 4px 4px 8px 0 rgba(230, 162, 60, 0.2), -4px -4px 8px 0 rgba(255, 255, 255, 0.9);
}

.statistic-icon.remain {
  color: #67c23a;
  background-color: rgba(103, 194, 58, 0.1);
  box-shadow: 4px 4px 8px 0 rgba(103, 194, 58, 0.2), -4px -4px 8px 0 rgba(255, 255, 255, 0.9);
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


:deep(.filter-form .ant-form-item-label > label) {
  color: #4a5568 !important;
  padding-right: 8px;
  font-weight: 500;
}


:deep(.ant-input-affix-wrapper input) {
  background-color: transparent !important;
  box-shadow: none !important;
  border: none !important;
  padding: 4px 0 !important;
}



</style>
