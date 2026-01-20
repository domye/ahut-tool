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
          <a-input
            v-model:value="gradesStore.kksj"
            placeholder="请输入开课学期"
            style="width: 200px"
            allow-clear
          />
        </a-form-item>
        <a-form-item label="课程性质">
          <a-input
            v-model:value="gradesStore.kcxz"
            placeholder="请输入课程性质"
            style="width: 200px"
            allow-clear
          />
        </a-form-item>
        <a-form-item label="课程名称">
          <a-input
            v-model:value="gradesStore.kcmc"
            placeholder="请输入课程名称"
            style="width: 200px"
            allow-clear
          />
        </a-form-item>
        <a-form-item label="显示方式">
          <a-input
            v-model:value="gradesStore.xsfs"
            placeholder="请输入显示方式"
            style="width: 200px"
            allow-clear
          />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" @click="handleSearch" :loading="gradesStore.loading">
              查询
            </a-button>
            <a-button @click="handleReset">
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
        :pagination="{ pageSize: 10 }"
        size="middle"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'score'">
            <a-tag :color="getScoreColor(record.Score)">
              {{ record.Score }}
            </a-tag>
          </template>
        </template>
      </a-table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useGradesStore } from '../store/grades'
import { BookOutlined, TrophyOutlined, RiseOutlined, LineChartOutlined } from '@ant-design/icons-vue'
import type { TableColumnsType } from 'ant-design-vue'

const gradesStore = useGradesStore()

const columns: TableColumnsType = [
  { title: '序号', dataIndex: 'Index', key: 'Index', width: 80, fixed: 'left' },
  { title: '开课学期', dataIndex: 'Semester', key: 'Semester', width: 120 },
  { title: '课程编号', dataIndex: 'CourseID', key: 'CourseID', width: 150 },
  { title: '课程名称', dataIndex: 'CourseName', key: 'CourseName', width: 200 },
  { title: '分组名', dataIndex: 'GroupName', key: 'GroupName', width: 150 },
  { title: '成绩', dataIndex: 'Score', key: 'score', width: 100 },
  { title: '成绩标识', dataIndex: 'ScoreFlag', key: 'ScoreFlag', width: 100 },
  { title: '学分', dataIndex: 'Credit', key: 'Credit', width: 80 },
  { title: '总学时', dataIndex: 'TotalHours', key: 'TotalHours', width: 100 },
  { title: '绩点', dataIndex: 'GPA', key: 'GPA', width: 80 },
  { title: '补重学期', dataIndex: 'RetakeSem', key: 'RetakeSem', width: 120 },
  { title: '考核方式', dataIndex: 'ExamMode', key: 'ExamMode', width: 100 },
  { title: '考试性质', dataIndex: 'ExamType', key: 'ExamType', width: 100 },
  { title: '课程属性', dataIndex: 'CourseAttr', key: 'CourseAttr', width: 100 },
  { title: '课程性质', dataIndex: 'CourseNature', key: 'CourseNature', width: 100 },
  { title: '通选课类别', dataIndex: 'GEType', key: 'GEType', width: 120 }
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

async function handleSearch() {
  await gradesStore.fetchGrades()
}

function handleReset() {
  gradesStore.resetFilters()
  gradesStore.fetchGrades()
}

onMounted(() => {
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

:deep(.ant-table) {
  background-color: #e0e5ec;
  border-radius: 12px;
}

:deep(.ant-table-wrapper) {
  border-radius: 12px;
}

:deep(.ant-table-thead > tr > th) {
  background-color: #e0e5ec !important;
  color: #4a5568 !important;
  font-weight: 600;
  text-align: center;
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.3), inset -4px -4px 8px 0 rgba(255,255,255, 0.8) !important;
  border-color: #e0e5ec !important;
  border: none !important;
  padding: 16px !important;
}

:deep(.ant-table-thead > tr > th:first-child) {
  border-top-left-radius: 12px;
  border-bottom-left-radius: 12px;
}

:deep(.ant-table-thead > tr > th:last-child) {
  border-top-right-radius: 12px;
  border-bottom-right-radius: 12px;
}

:deep(.ant-table-tbody > tr > td) {
  text-align: center;
  background-color: #e0e5ec;
  border-color: #e0e5ec;
  color: #4a5568;
  border: none;
  padding: 12px 16px;
  transition: all 0.3s ease;
}

:deep(.ant-table-tbody > tr:hover > td) {
  background-color: #e0e5ec !important;
}

:deep(.ant-table-tbody > tr > td:first-child) {
  border-top-left-radius: 8px;
  border-bottom-left-radius: 8px;
}

:deep(.ant-table-tbody > tr > td:last-child) {
  border-top-right-radius: 8px;
  border-bottom-right-radius: 8px;
}

:deep(.ant-table-pagination.ant-pagination) {
  margin: 24px 0 0 0;
  text-align: center;
}

:deep(.ant-pagination-item) {
  background-color: #e0e5ec;
  border: none;
  border-radius: 8px;
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.3), -4px -4px 8px 0 rgba(255,255,255, 0.8);
  color: #4a5568;
  transition: all 0.3s ease;
}

:deep(.ant-pagination-item:hover) {
  box-shadow: 3px 3px 6px 0 rgba(163,177,198, 0.4), -3px -3px 6px 0 rgba(255,255,255, 0.9);
}

:deep(.ant-pagination-item-active) {
  background-color: #409eff;
  color: white;
  box-shadow: 4px 4px 8px 0 rgba(64, 158, 255, 0.3), -4px -4px 8px 0 rgba(150, 200, 255, 0.2);
}

:deep(.ant-pagination-prev),
:deep(.ant-pagination-next) {
  background-color: #e0e5ec;
  border: none;
  border-radius: 8px;
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.3), -4px -4px 8px 0 rgba(255,255,255, 0.8);
  color: #4a5568;
  transition: all 0.3s ease;
}

:deep(.ant-pagination-prev:hover),
:deep(.ant-pagination-next:hover) {
  box-shadow: 3px 3px 6px 0 rgba(163,177,198, 0.4), -3px -3px 6px 0 rgba(255,255,255, 0.9);
}

:deep(.ant-pagination-disabled) {
  box-shadow: inset 3px 3px 6px 0 rgba(163,177,198, 0.3), inset -3px -3px 6px 0 rgba(255,255,255, 0.8);
  color: #a0aec0;
}

:deep(.ant-tag) {
  border: none;
  font-weight: 500;
  border-radius: 8px;
  padding: 4px 12px;
  box-shadow: 3px 3px 6px 0 rgba(163,177,198, 0.2), -3px -3px 6px 0 rgba(255,255,255, 0.8);
}

:deep(.ant-spin-container) {
  border-radius: 12px;
}

:deep(.ant-table-container) {
  border-radius: 12px;
}

:deep(.ant-table-wrapper) {
  padding: 8px;
}

:deep(.ant-table-body) {
  margin: 8px;
  border-radius: 12px;
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.15), inset -4px -4px 8px 0 rgba(255,255,255, 0.8);
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
}
</style>