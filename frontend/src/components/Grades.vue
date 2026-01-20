<script lang="ts" setup>
import { reactive, onMounted } from 'vue'
import { GetGrades } from '../../wailsjs/go/backend/App'
import { models } from '../../wailsjs/go/models'

const data = reactive({
  grades: [] as models.Grade[],
  summary: {} as models.GradeSummary,
  loading: false,
  error: '',
  // 查询参数
  kksj: '',  // 开课学期
  kcxz: '',  // 课程性质
  kcmc: '',  // 课程名称
  xsfs: ''   // 显示方式
})

function fetchGrades() {
  data.loading = true
  data.error = ''
  GetGrades(data.kksj, data.kcxz, data.kcmc, data.xsfs)
    .then((response: any) => {
      data.grades = response.grades
      data.summary = response.summary
    })
    .catch((err: any) => {
      data.error = '获取成绩失败: ' + err
    })
    .finally(() => {
      data.loading = false
    })
}

// 组件挂载时自动获取成绩
onMounted(() => {
  fetchGrades()
})
</script>

<template>
  <div class="grades-container">
    <div class="header">
      <h2>成绩查询</h2>
      <div class="summary" v-if="data.summary.CourseCount">
        <div class="summary-item">
          <span class="label">所修门数：</span>
          <span class="value">{{ data.summary.CourseCount }}</span>
        </div>
        <div class="summary-item">
          <span class="label">所修总学分：</span>
          <span class="value">{{ data.summary.TotalCredit }}</span>
        </div>
        <div class="summary-item">
          <span class="label">平均学分绩点：</span>
          <span class="value">{{ data.summary.AvgGPA }}</span>
        </div>
        <div class="summary-item">
          <span class="label">平均成绩：</span>
          <span class="value">{{ data.summary.AvgScore }}</span>
        </div>
      </div>
    </div>

    <div class="filter-bar">
      <input v-model="data.kksj" placeholder="开课学期" class="filter-input" />
      <input v-model="data.kcxz" placeholder="课程性质" class="filter-input" />
      <input v-model="data.kcmc" placeholder="课程名称" class="filter-input" />
      <input v-model="data.xsfs" placeholder="显示方式" class="filter-input" />
      <button @click="fetchGrades" class="search-btn">查询</button>
    </div>

    <div v-if="data.loading" class="loading">加载中...</div>
    <div v-else-if="data.error" class="error">{{ data.error }}</div>
    <div v-else-if="data.grades.length === 0" class="empty">暂无成绩数据</div>
    <div v-else class="table-container">
      <table class="grades-table">
        <thead>
          <tr>
            <th>序号</th>
            <th>开课学期</th>
            <th>课程编号</th>
            <th>课程名称</th>
            <th>分组名</th>
            <th>成绩</th>
            <th>成绩标识</th>
            <th>学分</th>
            <th>总学时</th>
            <th>绩点</th>
            <th>补重学期</th>
            <th>考核方式</th>
            <th>考试性质</th>
            <th>课程属性</th>
            <th>课程性质</th>
            <th>通选课类别</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="grade in data.grades" :key="grade.Index">
            <td>{{ grade.Index }}</td>
            <td>{{ grade.Semester }}</td>
            <td>{{ grade.CourseID }}</td>
            <td>{{ grade.CourseName }}</td>
            <td>{{ grade.GroupName }}</td>
            <td>{{ grade.Score }}</td>
            <td>{{ grade.ScoreFlag }}</td>
            <td>{{ grade.Credit }}</td>
            <td>{{ grade.TotalHours }}</td>
            <td>{{ grade.GPA }}</td>
            <td>{{ grade.RetakeSem }}</td>
            <td>{{ grade.ExamMode }}</td>
            <td>{{ grade.ExamType }}</td>
            <td>{{ grade.CourseAttr }}</td>
            <td>{{ grade.CourseNature }}</td>
            <td>{{ grade.GEType }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.grades-container {
  padding: 2rem;
  max-width: 1600px;
  margin: 0 auto;
}

.header {
  margin-bottom: 2rem;
}

.header h2 {
  color: #333;
  margin-bottom: 1rem;
}

.summary {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  background: #f5f7fa;
  padding: 1.5rem;
  border-radius: 8px;
}

.summary-item {
  display: flex;
  align-items: center;
}

.summary-item .label {
  font-weight: 500;
  color: #666;
  margin-right: 0.5rem;
}

.summary-item .value {
  font-weight: 600;
  color: #333;
}

.filter-bar {
  display: flex;
  gap: 1rem;
  margin-bottom: 2rem;
  flex-wrap: wrap;
}

.filter-input {
  padding: 0.5rem 1rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 0.9rem;
  flex: 1;
  min-width: 150px;
}

.search-btn {
  padding: 0.5rem 2rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: transform 0.2s;
}

.search-btn:hover {
  transform: translateY(-2px);
}

.loading,
.error,
.empty {
  text-align: center;
  padding: 2rem;
  font-size: 1.1rem;
}

.error {
  color: #e74c3c;
}

.table-container {
  overflow-x: auto;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.grades-table {
  width: 100%;
  border-collapse: collapse;
  background: white;
}

.grades-table th,
.grades-table td {
  padding: 1rem;
  text-align: left;
  border-bottom: 1px solid #eee;
}

.grades-table th {
  background: #f5f7fa;
  font-weight: 600;
  color: #333;
  white-space: nowrap;
}

.grades-table tbody tr:hover {
  background: #f9fafb;
}

.grades-table tbody tr:last-child td {
  border-bottom: none;
}
</style>
