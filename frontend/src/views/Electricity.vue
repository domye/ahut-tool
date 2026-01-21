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
        <a-form-item label="楼栋">
          <a-select
            v-model:value="selectedBuilding"
            placeholder="请选择楼栋"
            style="width: 200px"
            allow-clear
            :options="buildingOptions"
            @change="handleBuildingChange"
            popup-class-name="neumorphic-select-dropdown"
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
import { ref, computed, watch } from 'vue'
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

// 楼栋数据
const newCampusBuildings = [
  { Id: "01", Name: "东校区1号学生宿舍楼" },
  { Id: "02", Name: "东校区2号学生宿舍楼" },
  { Id: "03", Name: "东校区3号学生宿舍楼" },
  { Id: "04", Name: "东校区4号学生宿舍楼" },
  { Id: "05", Name: "东校区5号学生宿舍楼" },
  { Id: "06", Name: "东校区6号学生宿舍楼" },
  { Id: "07", Name: "东校区7号学生宿舍楼" },
  { Id: "08", Name: "东校区8号学生宿舍楼" },
  { Id: "09", Name: "东校区9号学生宿舍楼" },
  { Id: "10", Name: "东校区A号学生宿舍楼" },
  { Id: "11", Name: "东校区B号学生宿舍楼" },
  { Id: "13", Name: "东校区C号学生宿舍楼" },
  { Id: "14", Name: "东校区D号学生宿舍楼" },
  { Id: "15", Name: "东校区E号学生宿舍楼" },
  { Id: "16", Name: "东校区F号学生宿舍楼" },
  { Id: "17", Name: "东校区H1号学生宿舍楼" },
  { Id: "18", Name: "东校区H2号学生宿舍楼" },
  { Id: "19", Name: "东校区H3号学生宿舍楼" },
  { Id: "20", Name: "东校区G1号学生宿舍楼" },
  { Id: "21", Name: "东校区G2号学生宿舍楼" },
  { Id: "22", Name: "东校区G3号学生宿舍楼" },
  { Id: "23", Name: "东校区J1号学生宿舍楼" },
  { Id: "24", Name: "东校区J2号学生宿舍楼" },
  { Id: "25", Name: "东校区J3号学生宿舍楼" },
  { Id: "26", Name: "东校区K1号学生宿舍楼" },
  { Id: "27", Name: "东校区K2号学生宿舍楼" },
  { Id: "28", Name: "东校区K3号学生宿舍楼" },
  { Id: "29", Name: "东校区L1号学生宿舍楼" },
  { Id: "30", Name: "东校区L2号学生宿舍楼" },
  { Id: "31", Name: "东校区G4号学生宿舍楼" },
  { Id: "32", Name: "东校区研5号学生宿舍楼" },
  { Id: "33", Name: "东校区研6号学生宿舍楼" },
  { Id: "34", Name: "东校区研7号学生宿舍楼" },
  { Id: "35", Name: "东校区研8号学生宿舍楼" },
  { Id: "36", Name: "东校区研1号学生宿舍楼" },
  { Id: "37", Name: "东校区研2号学生宿舍楼" },
  { Id: "38", Name: "东校区研3号学生宿舍楼" },
  { Id: "39", Name: "东校区研4号学生宿舍楼" },
  { Id: "40", Name: "东校区M栋南楼宿舍" },
  { Id: "41", Name: "东校区M栋北楼宿舍" },
  { Id: "42", Name: "东校区N栋南楼宿舍" },
  { Id: "43", Name: "东校区N栋北楼宿舍" }
]

const oldCampusBuildings = [
  { Id: "01", Name: "本部校区01栋学生宿舍" },
  { Id: "02", Name: "本部校区02栋学生宿舍" },
  { Id: "03", Name: "本部校区03栋学生宿舍" },
  { Id: "04", Name: "本部校区04栋学生宿舍" },
  { Id: "05", Name: "本部校区05A栋学生宿舍" },
  { Id: "06", Name: "本部校区05B栋学生宿舍" },
  { Id: "07", Name: "本部校区矿东栋学生宿舍" },
  { Id: "08", Name: "本部校区矿西栋学生宿舍" },
  { Id: "09", Name: "本部校区研A栋学生宿舍" },
  { Id: "10", Name: "本部校区研B栋学生宿舍" },
  { Id: "11", Name: "本部校区8号公寓" },
  { Id: "12", Name: "本部校区7号公寓" }
]

// 当前选中的楼栋
const selectedBuilding = ref('')

// 楼栋选项（根据校区动态生成）
const buildingOptions = computed(() => {
  const buildings = electricityStore.xiaoqu === 'OldS' ? oldCampusBuildings : newCampusBuildings
  return buildings.map(building => ({
    value: building.Id,
    label: building.Name
  }))
})

// 处理楼栋选择变化
function handleBuildingChange(value: string) {
  const buildings = electricityStore.xiaoqu === 'OldS' ? oldCampusBuildings : newCampusBuildings
  const selected = buildings.find(b => b.Id === value)
  if (selected) {
    electricityStore.ld_Id = selected.Id
    electricityStore.ld_Name = selected.Name
  }
}

// 监听校区变化，重置楼栋选择
watch(() => electricityStore.xiaoqu, () => {
  selectedBuilding.value = ''
  electricityStore.ld_Id = ''
  electricityStore.ld_Name = ''
})

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
