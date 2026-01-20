<template>
  <div class="grades-page">
    <!-- 成绩汇总 -->
    <a-row :gutter="[24, 24]" class="summary-row">
      <a-col :xs="24" :sm="12" :md="6" :lg="6" :xl="6">
        <div class="statistic-card">
          <div class="statistic-icon">
            <BookOutlined />
          </div>
          <div class="statistic-content">
            <div class="statistic-title">所修门数</div>
            <div class="statistic-value">{{ gradesStore.summary.CourseCount || 0 }}</div>
          </div>
        </div>
      </a-col>
      <a-col :xs="24" :sm="12" :md="6" :lg="6" :xl="6">
        <div class="statistic-card">
          <div class="statistic-icon">
            <TrophyOutlined />
          </div>
          <div class="statistic-content">
            <div class="statistic-title">所修总学分</div>
            <div class="statistic-value">{{ gradesStore.summary.TotalCredit || 0 }}</div>
          </div>
        </div>
      </a-col>
      <a-col :xs="24" :sm="12" :md="6" :lg="6" :xl="6">
        <div class="statistic-card">
          <div class="statistic-icon">
            <RiseOutlined />
          </div>
          <div class="statistic-content">
            <div class="statistic-title">平均学分绩点</div>
            <div class="statistic-value">{{ gradesStore.summary.AvgGPA || 0 }}</div>
          </div>
        </div>
      </a-col>
      <a-col :xs="24" :sm="12" :md="6" :lg="6" :xl="6">
        <div class="statistic-card">
          <div class="statistic-icon">
            <LineChartOutlined />
          </div>
          <div class="statistic-content">
            <div class="statistic-title">平均成绩</div>
            <div class="statistic-value">{{ gradesStore.summary.AvgScore || 0 }}</div>
          </div>
        </div>
      </a-col>
    </a-row>

    <!-- 查询条件 -->
    <div class="filter-section">
      <a-form layout="inline" class="filter-form">
        <a-form-item label="开课学期">
          <a-select
            v-model:value="gradesStore.kksj"
            placeholder="请选择开课学期"
            style="width: 200px"
            allow-clear
            :options="semesterOptions"
            popup-class-name="neumorphic-select-dropdown"
          />
        </a-form-item>
<!--        <a-form-item label="课程性质">-->
<!--          <a-input-->
<!--            v-model:value="gradesStore.kcxz"-->
<!--            placeholder="请输入课程性质"-->
<!--            style="width: 200px"-->
<!--            allow-clear-->
<!--          />-->
<!--        </a-form-item>-->
        <a-form-item label="课程名称">
          <a-input
            v-model:value="gradesStore.kcmc"
            placeholder="请输入课程名称"
            style="width: 200px"
            allow-clear
          />
        </a-form-item>
<!--        <a-form-item label="显示方式">-->
<!--          <a-input-->
<!--            v-model:value="gradesStore.xsfs"-->
<!--            placeholder="请输入显示方式"-->
<!--            style="width: 200px"-->
<!--            allow-clear-->
<!--          />-->
<!--        </a-form-item>-->
        <a-form-item>
          <a-space :size="20">
            <a-button 
              type="primary" 
              @click="handleSearch" 
              :loading="gradesStore.loading"
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
    </div>

    <!-- 成绩表格 -->
    <div class="table-section">
      <a-table
        :columns="columns"
        :data-source="gradesStore.grades"
        :loading="gradesStore.loading"
        :scroll="{ x: 'max-content' }"
        :pagination="{ pageSize: 20}"
        size="middle"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'score'">
            <a-tag :color="getScoreColor(getDisplayScore(record.Score, record.GPA))">
              {{ getDisplayScore(record.Score, record.GPA) }}
            </a-tag>
          </template>
        </template>
      </a-table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { useGradesStore } from '../store/grades'
import { BookOutlined, TrophyOutlined, RiseOutlined, LineChartOutlined, SearchOutlined, RedoOutlined } from '@ant-design/icons-vue'
import type { TableColumnsType } from 'ant-design-vue'

const gradesStore = useGradesStore()

// 获取当前学期
function getCurrentSemester(): string {
  const currentDate = new Date()
  const currentYear = currentDate.getFullYear()
  const currentMonth = currentDate.getMonth() + 1 // 1-12

  // 确定基准年份：3月及以后使用当前年份，否则使用上一年份
  const baseYear = currentMonth >= 3 ? currentYear : currentYear - 1

  // 9月到次年3月为第一学期，3月到9月为第二学期
  const semester = currentMonth >= 3 && currentMonth < 9 ? '2' : '1'

  return `${baseYear}-${baseYear + 1}-${semester}`
}

// 计算学期选项
const semesterOptions = computed(() => {
  const currentDate = new Date()
  const currentYear = currentDate.getFullYear()
  const currentMonth = currentDate.getMonth() + 1 // 1-12

  // 确定基准年份：9月及以后使用当前年份，否则使用上一年份
  const baseYear = currentMonth >= 9 ? currentYear : currentYear - 1

  // 生成前三年到后三年的学期选项
  return Array.from({ length: 7 }, (_, i) => {
    const year = baseYear + i - 3
    return [
      {
        value: `${year}-${year + 1}-1`,
        label: `${year}-${year + 1} 第一学期`
      },
      {
        value: `${year}-${year + 1}-2`,
        label: `${year}-${year + 1} 第二学期`
      }
    ]
  }).flat()
})

const columns: TableColumnsType = [
  { title: '序号', dataIndex: 'Index', key: 'Index', width: 80, fixed: 'left' },
  { title: '开课学期', dataIndex: 'Semester', key: 'Semester', width: 120 },
  // { title: '课程编号', dataIndex: 'CourseID', key: 'CourseID', width: 150 },
  { title: '课程名称', dataIndex: 'CourseName', key: 'CourseName', width: 150 },
  // { title: '分组名', dataIndex: 'GroupName', key: 'GroupName', width: 150 },
  { title: '成绩', dataIndex: 'Score', key: 'score', width: 100 },
  // { title: '成绩标识', dataIndex: 'ScoreFlag', key: 'ScoreFlag', width: 100 },
  // { title: '学分', dataIndex: 'Credit', key: 'Credit', width: 80 },
  // { title: '总学时', dataIndex: 'TotalHours', key: 'TotalHours', width: 100 },
  // { title: '绩点', dataIndex: 'GPA', key: 'GPA', width: 80 },
  // { title: '补重学期', dataIndex: 'RetakeSem', key: 'RetakeSem', width: 120 },
  // { title: '考核方式', dataIndex: 'ExamMode', key: 'ExamMode', width: 100 },
  // { title: '考试性质', dataIndex: 'ExamType', key: 'ExamType', width: 100 },
  // { title: '课程属性', dataIndex: 'CourseAttr', key: 'CourseAttr', width: 100 },
  { title: '课程性质', dataIndex: 'CourseNature', key: 'CourseNature', width: 100 },
  // { title: '通选课类别', dataIndex: 'GEType', key: 'GEType', width: 120 }
]

function getScoreColor(score: string): string {
  const numScore = parseFloat(score)
  if (isNaN(numScore)) return 'default'
  if (numScore >= 90) return 'green'
  if (numScore >= 80) return 'blue'
  if (numScore >= 70) return 'orange'
  if (numScore >= 60) return 'gold'
  return 'red'
}

// 根据绩点推算成绩
function getDisplayScore(score: string, gpa: string): string {
  // 如果成绩不是"未评教"，直接返回原成绩
  if (score !== '未评教') {
    return score
  }

  // 如果成绩是"未评教"，根据绩点推算成绩
  const numGpa = parseFloat(gpa)
  if (isNaN(numGpa)) return score

  // 使用数组存储绩点范围和对应的成绩范围
  const gpaRanges = [
    { min: 4.5, score: '96-100' },
    { min: 4.2, score: '90-95' },
    { min: 3.8, score: '85-89' },
    { min: 3.5, score: '80-84' },
    { min: 2.9, score: '75-79' },
    { min: 2.5, score: '70-74' },
    { min: 1.8, score: '60-69' }
  ]

  // 查找匹配的绩点范围
  for (const range of gpaRanges) {
    if (numGpa >= range.min) {
      return range.score
    }
  }

  return '<60'
}

async function handleSearch() {
  await gradesStore.fetchGrades()
}

function handleReset() {
  gradesStore.resetFilters()
  gradesStore.fetchGrades()
}

onMounted(() => {
  // 设置当前学期
  if (!gradesStore.kksj) {
    gradesStore.kksj = getCurrentSemester()
  }
  gradesStore.fetchGrades()
})
</script>

<style scoped>
.grades-page {
  padding: 24px;
  background-color: #e0e5ec;
  min-height: calc(100vh - 48px);
}

.summary-row {
  margin-bottom: 24px;
}

.statistic-card {
  background-color: #e0e5ec;
  border-radius: 16px;
  padding: 16px;
  display: flex;
  align-items: center;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
  height: 100%;
  min-height: 80px;
}

.statistic-icon {
  font-size: 1.8rem;
  margin-right: 16px;
  color: #4a5568;
}

.statistic-content {
  flex: 1;
}

.statistic-title {
  font-size: 0.9rem;
  color: #718096;
  margin-bottom: 4px;
}

.statistic-value {
  font-size: 1.5rem;
  font-weight: 600;
  color: #4a5568;
}

.filter-section {
  background-color: #e0e5ec;
  border-radius: 16px;
  padding: 20px;
  margin-bottom: 24px;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
}

.filter-form {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
  align-items: center;
}

:deep(.filter-form .ant-form-item) {
  margin-bottom: 0;
}

:deep(.filter-form .ant-form-item-label > label) {
  color: #4a5568 !important;
  padding-right: 8px;
  font-weight: 500;
}

:deep(.ant-form-item) {
  margin-bottom: 16px;
}

:deep(.ant-form-item-label) {
  padding-bottom: 4px;
}

:deep(.filter-form .ant-input) {
  background-color: transparent !important;
  box-shadow: none !important;
  border: none !important;
  padding: 4px 0 !important;
  color: #4a5568 !important;
  transition: all 0.3s;
}

:deep(.filter-form .ant-input:hover) {
  background-color: transparent !important;
}

:deep(.filter-form .ant-input:focus) {
  background-color: transparent !important;
  outline: none !important;
  border: none !important;
}

:deep(.ant-input-affix-wrapper) {
  background-color: #e0e5ec !important;
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.3), inset -4px -4px 8px 0 rgba(255,255,255, 0.8) !important;
  border: none !important;
  border-radius: 12px !important;
  color: #4a5568 !important;
  padding: 4px 12px !important;
}

:deep(.ant-input-affix-wrapper:hover) {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.4), inset -4px -4px 8px 0 rgba(255,255,255, 0.9) !important;
}

:deep(.ant-input-affix-wrapper-focused) {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.5), inset -4px -4px 8px 0 rgba(255,255,255, 1) !important;
  outline: none !important;
  border: none !important;
}

:deep(.ant-input-affix-wrapper input) {
  background-color: transparent !important;
  box-shadow: none !important;
  border: none !important;
  padding: 4px 0 !important;
}

:deep(.ant-input-affix-wrapper .ant-input) {
  background-color: transparent !important;
  box-shadow: none !important;
  border: none !important;
}

:deep(.ant-btn) {
  background-color: #e0e5ec !important;
  border: none !important;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8) !important;
  border-radius: 12px !important;
  color: #4a5568 !important;
  font-weight: 500;
  transition: all 0.3s ease;
  padding: 8px 16px;
}

:deep(.ant-btn-primary) {
  background-color: #409eff !important;
  color: white !important;
  box-shadow: 6px 6px 10px 0 rgba(64, 158, 255, 0.3), -6px -6px 10px 0 rgba(150, 200, 255, 0.2) !important;
}

:deep(.ant-btn:hover) {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8) !important;
}

:deep(.ant-btn-primary:hover) {
  box-shadow: 4px 4px 8px 0 rgba(64, 158, 255, 0.3), -4px -4px 8px 0 rgba(150, 200, 255, 0.2) !important;
}

:deep(.ant-btn:active) {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.8) !important;
}

:deep(.ant-btn-primary:active) {
  box-shadow: inset 4px 4px 8px 0 rgba(64, 158, 255, 0.3), inset -4px -4px 8px 0 rgba(150, 200, 255, 0.2) !important;
}

:deep(.ant-input-clear-icon) {
  color: #a0aec0 !important;
  transition: all 0.3s;
}

:deep(.ant-input-clear-icon:hover) {
  color: #4a5568 !important;
}

.table-section {
  background-color: #e0e5ec;
  border-radius: 16px;
  padding: 20px;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
}

/* 表格头部 */
:deep(.ant-table-thead > tr > th) {
  background-color: #e0e5ec !important;
  color: #4a5568 !important;
  font-weight: 600;
  text-align: center;
  border-bottom: 2px solid rgba(163,177,198, 0.4) !important;
  padding: 14px 12px !important;
}

/* 表格单元格 */
:deep(.ant-table-tbody > tr > td) {
  text-align: center;
  background-color: #e0e5ec !important;
  border-bottom: 1px solid rgba(163,177,198, 0.3) !important;
  color: #4a5568 !important;
  padding: 12px !important;
}

/* 表格行悬停效果 */
:deep(.ant-table-tbody > tr:hover > td) {
  background-color: #dce1e8 !important;
}

/* 分页样式 */
:deep(.ant-pagination) {
  margin: 24px 0 0 0;
  gap: 12px !important;
}

:deep(.ant-pagination-item) {
  background-color: #e0e5ec !important;
  border: none !important;
  border-radius: 8px !important;
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.3), -4px -4px 8px 0 rgba(255,255,255, 0.8) !important;
  color: #4a5568 !important;
  transition: all 0.3s ease;
  margin: 0 4px !important;
  min-width: 20px !important;
  height: 20px !important;
  line-height: 20px !important;
}

:deep(.ant-pagination-item-active) {
  background-color: #409eff !important;
  color: white !important;
  box-shadow: 4px 4px 8px 0 rgba(64, 158, 255, 0.3), -4px -4px 8px 0 rgba(150, 200, 255, 0.2) !important;
}

:deep(.ant-pagination-prev),
:deep(.ant-pagination-next) {
  background-color: #e0e5ec !important;
  border: none !important;
  border-radius: 8px !important;
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.3), -4px -4px 8px 0 rgba(255,255,255, 0.8) !important;
  color: #4a5568 !important;
  min-width: 20px !important;
  height: 20px !important;
  line-height: 20px !important;
}

/* 标签样式 */
:deep(.ant-tag) {
  border: none !important;
  font-weight: 500;
  border-radius: 8px;
  padding: 4px 12px;
  box-shadow: 3px 3px 6px 0 rgba(163,177,198, 0.2), -3px -3px 6px 0 rgba(255,255,255, 0.8);
}

/* 下拉框新拟态风格 */
:deep(.ant-select-selector) {
  background-color: #e0e5ec !important;
  border: none !important;
  border-radius: 8px !important;
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.3), -4px -4px 8px 0 rgba(255,255,255, 0.8) !important;
  color: #4a5568 !important;
}

:deep(.ant-select-focused .ant-select-selector) {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.3), -4px -4px 8px 0 rgba(255,255,255, 0.8) !important;
}

:deep(.ant-select-arrow) {
  color: #4a5568 !important;
}

/* 下拉菜单新拟态风格 - 全局样式 */
:global(.neumorphic-select-dropdown) {
  background-color: #e0e5ec !important;
  border: none !important;
  border-radius: 12px !important;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8) !important;
  padding: 8px !important;
}

/* 下拉菜单滚动容器 */
:global(.neumorphic-select-dropdown .rc-virtual-list) {
  scrollbar-width: thin;
  scrollbar-color: #a3b1c6 #e0e5ec;
}

/* Webkit 滚动条样式 */
:global(.neumorphic-select-dropdown .rc-virtual-list::-webkit-scrollbar) {
  width: 6px;
  height: 6px;
}

:global(.neumorphic-select-dropdown .rc-virtual-list::-webkit-scrollbar-track) {
  background-color: #e0e5ec;
  border-radius: 3px;
  box-shadow: inset 1px 1px 2px 0 rgba(163,177,198, 0.3), inset -1px -1px 2px 0 rgba(255,255,255, 0.8);
}

:global(.neumorphic-select-dropdown .rc-virtual-list::-webkit-scrollbar-thumb) {
  background-color: #a3b1c6;
  border-radius: 3px;
  box-shadow: 2px 2px 4px 0 rgba(163,177,198, 0.3), -2px -2px 4px 0 rgba(255,255,255, 0.8);
  transition: all 0.3s ease;
}

:global(.neumorphic-select-dropdown .rc-virtual-list::-webkit-scrollbar-thumb:hover) {
  background-color: #8fa3be;
  box-shadow: 3px 3px 6px 0 rgba(163,177,198, 0.4), -3px -3px 6px 0 rgba(255,255,255, 0.9);
}

:global(.neumorphic-select-dropdown .ant-select-item) {
  background-color: transparent !important;
  color: #4a5568 !important;
  border-radius: 8px;
  margin: 4px 0;
  transition: all 0.3s ease;
  box-shadow: inset 2px 2px 4px 0 rgba(163,177,198, 0.2), inset -2px -2px 4px 0 rgba(255,255,255, 0.8) !important;
}

:global(.neumorphic-select-dropdown .ant-select-item-option-selected) {
  background-color: #409eff !important;
  color: white !important;
  box-shadow: 3px 3px 6px 0 rgba(64, 158, 255, 0.3), -3px -3px 6px 0 rgba(150, 200, 255, 0.2) !important;
}

:global(.neumorphic-select-dropdown .ant-select-item-option-active) {
  background-color: #dce1e8 !important;
  color: #4a5568 !important;
  box-shadow: inset 2px 2px 4px 0 rgba(163,177,198, 0.3), inset -2px -2px 4px 0 rgba(255,255,255, 0.8) !important;
}

:global(.neumorphic-select-dropdown .ant-select-item-option-selected.ant-select-item-option-active) {
  background-color: #409eff !important;
  color: white !important;
  box-shadow: 3px 3px 6px 0 rgba(64, 158, 255, 0.3), -3px -3px 6px 0 rgba(150, 200, 255, 0.2) !important;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .statistic-card {
    margin-bottom: 16px;
  }
  
  .filter-form {
    flex-direction: column;
    align-items: stretch;
  }
  
  :deep(.filter-form .ant-form-item) {
    margin-bottom: 16px;
  }
  
  :deep(.filter-form .ant-form-item-control-wrapper) {
    display: flex;
  }
  
  :deep(.filter-form .ant-input) {
    width: 100% !important;
  }

  /* 下拉框响应式适配 */
  :deep(.filter-form .ant-select) {
    width: 100% !important;
  }

  :deep(.ant-select-selector) {
    min-height: 40px !important;
    padding: 4px 12px !important;
  }

  :deep(.ant-select-selection-item) {
    line-height: 32px !important;
    font-size: 14px;
  }

  :deep(.ant-select-dropdown) {
    max-height: 300px;
  }

  :deep(.ant-select-item) {
    padding: 8px 12px;
    font-size: 14px;
  }
}

/* 按钮新拟态风格 */
.neumorphic-button {
  background-color: #e0e5ec !important;
  border: none !important;
  border-radius: 8px !important;
  color: #4a5568 !important;
  font-weight: 500;
  height: 40px !important;
  padding: 0 24px !important;
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.3), -4px -4px 8px 0 rgba(255,255,255, 0.8) !important;
  transition: all 0.3s ease !important;
  display: inline-flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.neumorphic-button:hover {
  box-shadow: 6px 6px 12px 0 rgba(163,177,198, 0.4), -6px -6px 12px 0 rgba(255,255,255, 0.9) !important;
  transform: translateY(-2px);
}

.neumorphic-button:active {
  box-shadow: inset 2px 2px 4px 0 rgba(163,177,198, 0.3), inset -2px -2px 4px 0 rgba(255,255,255, 0.8) !important;
  transform: translateY(0);
}

/* 查询按钮特殊样式 */
.query-button {
  color: #409eff !important;
}

.query-button:hover {
  color: #66b1ff !important;
}

/* 重置按钮特殊样式 */
.reset-button:hover {
  color: #409eff !important;
}
</style>