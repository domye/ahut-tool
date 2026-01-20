<template>
  <div class="grades-container">
    <a-card title="成绩查询" :bordered="false">
      <!-- 成绩汇总 -->
      <a-row :gutter="16" class="summary-row">
        <a-col :span="6">
          <a-statistic
            title="所修门数"
            :value="gradesStore.summary.CourseCount || 0"
            :value-style="{ color: '#3f8600' }"
          >
            <template #prefix>
              <BookOutlined />
            </template>
          </a-statistic>
        </a-col>
        <a-col :span="6">
          <a-statistic
            title="所修总学分"
            :value="gradesStore.summary.TotalCredit || 0"
            :value-style="{ color: '#3f8600' }"
          >
            <template #prefix>
              <TrophyOutlined />
            </template>
          </a-statistic>
        </a-col>
        <a-col :span="6">
          <a-statistic
            title="平均学分绩点"
            :value="gradesStore.summary.AvgGPA || 0"
            :precision="2"
            :value-style="{ color: '#1890ff' }"
          >
            <template #prefix>
              <RiseOutlined />
            </template>
          </a-statistic>
        </a-col>
        <a-col :span="6">
          <a-statistic
            title="平均成绩"
            :value="gradesStore.summary.AvgScore || 0"
            :precision="2"
            :value-style="{ color: '#1890ff' }"
          >
            <template #prefix>
              <LineChartOutlined />
            </template>
          </a-statistic>
        </a-col>
      </a-row>

      <!-- 查询条件 -->
      <a-divider />

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

      <!-- 成绩表格 -->
      <a-divider />

      <a-table
        :columns="columns"
        :data-source="gradesStore.grades"
        :loading="gradesStore.loading"
        :scroll="{ x: 1500 }"
        :pagination="{ pageSize: 10 }"
        bordered
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'score'">
            <a-tag :color="getScoreColor(record.Score)">
              {{ record.Score }}
            </a-tag>
          </template>
        </template>
      </a-table>
    </a-card>
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
.grades-container {
  max-width: 1600px;
}

.summary-row {
  margin-bottom: 24px;
}

.filter-form {
  margin-bottom: 24px;
}
</style>
