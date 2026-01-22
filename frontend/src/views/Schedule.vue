<template>
  <div class="schedule-container">
    <div class="schedule-header">
      <h2 class="page-title">课程表</h2>
      <div class="selectors">
        <div class="semester-selector">
          <select v-model="scheduleStore.currentSemester" @change="handleSemesterChange" class="neumorphic-select">
            <option value="2025-2026-1">2025-2026 第一学期</option>
            <option value="2025-2026-2">2025-2026 第二学期</option>
          </select>
        </div>
        <div class="week-selector">
          <div class="week-info">第 {{ scheduleStore.currentWeek }} 周</div>
          <button @click="handlePrevWeek" class="neumorphic-btn week-btn">上一周</button>
          <button @click="handleNextWeek" class="neumorphic-btn week-btn">下一周</button>
        </div>
      </div>
    </div>

    <div v-if="scheduleStore.loading" class="loading-container">
      <div class="spinner"></div>
      <p>加载中...</p>
    </div>

    <div v-else-if="scheduleStore.error" class="error-container">
      <p class="error-message">{{ scheduleStore.error }}</p>
      <button @click="handleRefresh" class="neumorphic-btn">重试</button>
    </div>

    <div v-else class="schedule-content">
      <div class="schedule-grid">
        <!-- 表头 -->
        <div class="schedule-cell header-cell"></div>
        <div v-for="day in weekDays" :key="day" class="schedule-cell header-cell">
          {{ day }}
        </div>

        <!-- 课程表内容 -->
        <template v-for="period in periods" :key="period">
          <div class="schedule-cell period-cell">{{ period }}</div>
          <div
            v-for="day in 5"
            :key="`${period}-${day}`"
            class="schedule-cell course-cell"
            :class="{ 'has-course': hasCourse(day, period) }"
          >
            <div
              v-for="course in getCourse(day, period)"
              :key="course.name"
              class="course-card"
              :style="{ background: getCourseColor(course.name) }"
            >
              <div class="course-name">{{ course.name }}</div>
              <div class="course-info">{{ course.teacher }}</div>
              <div class="course-info">{{ course.classroom }}</div>
              <div class="course-info">周次: {{ course.weekNumbers }}</div>
            </div>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useScheduleStore } from '../store/schedule'

const scheduleStore = useScheduleStore()

const weekDays = ['周一', '周二', '周三', '周四', '周五']
const periods = ['1-2', '3-4', '5-6', '7-8', '9-10-11']

// 处理上一周
function handlePrevWeek() {
  if (scheduleStore.currentWeek > 1) {
    scheduleStore.setCurrentWeek(scheduleStore.currentWeek - 1)
  }
}

// 处理下一周
function handleNextWeek() {
  scheduleStore.setCurrentWeek(scheduleStore.currentWeek + 1)
}

// 为每门课程生成马卡龙色系颜色
const courseColors = computed(() => {
  const colors: { [key: string]: string } = {}
  const uniqueCourses = Array.from(new Set(scheduleStore.classes.map(c => c.name)))

  // 马卡龙色系
  const macaronColors = [
    'linear-gradient(135deg, rgba(255, 183, 178, 0.7) 0%, rgba(255, 158, 158, 0.7) 100%)', // 粉色
    'linear-gradient(135deg, rgba(181, 234, 215, 0.7) 0%, rgba(149, 216, 194, 0.7) 100%)', // 薄荷绿
    'linear-gradient(135deg, rgba(199, 206, 234, 0.7) 0%, rgba(168, 184, 232, 0.7) 100%)', // 淡紫
    'linear-gradient(135deg, rgba(255, 218, 193, 0.7) 0%, rgba(255, 201, 168, 0.7) 100%)', // 桃色
    'linear-gradient(135deg, rgba(226, 240, 203, 0.7) 0%, rgba(208, 230, 184, 0.7) 100%)', // 浅绿
    'linear-gradient(135deg, rgba(181, 234, 215, 0.7) 0%, rgba(149, 216, 194, 0.7) 100%)', // 薄荷蓝
    'linear-gradient(135deg, rgba(255, 154, 162, 0.7) 0%, rgba(255, 133, 143, 0.7) 100%)', // 珊瑚粉
    'linear-gradient(135deg, rgba(226, 240, 203, 0.7) 0%, rgba(205, 232, 181, 0.7) 100%)', // 柠檬绿
    'linear-gradient(135deg, rgba(255, 183, 178, 0.7) 0%, rgba(255, 165, 160, 0.7) 100%)', // 浅粉
    'linear-gradient(135deg, rgba(180, 231, 248, 0.7) 0%, rgba(158, 216, 240, 0.7) 100%)', // 天蓝
    // 'linear-gradient(135deg, rgba(253, 253, 150, 0.7) 0%, rgba(252, 248, 134, 0.7) 100%)', // 淡黄
    'linear-gradient(135deg, rgba(255, 179, 71, 0.7) 0%, rgba(255, 163, 53, 0.7) 100%)', // 橙色
    'linear-gradient(135deg, rgba(201, 174, 235, 0.7) 0%, rgba(183, 158, 224, 0.7) 100%)', // 紫罗兰
    'linear-gradient(135deg, rgba(255, 182, 193, 0.7) 0%, rgba(255, 164, 193, 0.7) 100%)', // 浅红
    'linear-gradient(135deg, rgba(152, 216, 200, 0.7) 0%, rgba(134, 200, 184, 0.7) 100%)', // 青绿
    'linear-gradient(135deg, rgba(165, 214, 167, 0.7) 0%, rgba(102, 187, 106, 0.7) 100%)', // 青色
    'linear-gradient(135deg, rgba(144, 202, 249, 0.7) 0%, rgba(66, 165, 245, 0.7) 100%)', // 蓝色
    'linear-gradient(135deg, rgba(239, 154, 154, 0.7) 0%, rgba(239, 83, 80, 0.7) 100%)', // 红色
    'linear-gradient(135deg, rgba(255, 224, 130, 0.7) 0%, rgba(255, 202, 40, 0.7) 100%)', // 琥珀色
    'linear-gradient(135deg, rgba(188, 170, 164, 0.7) 0%, rgba(141, 110, 99, 0.7) 100%)', // 棕色
    'linear-gradient(135deg, rgba(176, 190, 197, 0.7) 0%, rgba(120, 144, 156, 0.7) 100%)', // 蓝灰
    'linear-gradient(135deg, rgba(255, 204, 128, 0.7) 0%, rgba(255, 167, 38, 0.7) 100%)', // 杏色
    'linear-gradient(135deg, rgba(209, 196, 233, 0.7) 0%, rgba(149, 117, 205, 0.7) 100%)', // 靛青
    'linear-gradient(135deg, rgba(248, 187, 208, 0.7) 0%, rgba(240, 98, 146, 0.7) 100%)', // 玫瑰色
    'linear-gradient(135deg, rgba(197, 202, 233, 0.7) 0%, rgba(121, 134, 203, 0.7) 100%)', // 靛蓝
    'linear-gradient(135deg, rgba(220, 237, 200, 0.7) 0%, rgba(174, 213, 129, 0.7) 100%)', // 草绿
    'linear-gradient(135deg, rgba(255, 171, 64, 0.7) 0%, rgba(255, 145, 0, 0.7) 100%)', // 橙红
  ]

  uniqueCourses.forEach((courseName, index) => {
    colors[courseName] = macaronColors[index % macaronColors.length]
  })

  return colors
})

// 获取课程颜色
function getCourseColor(courseName: string): string {
  return courseColors.value[courseName] || 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)'
}

// 检查某个时间段是否有课程
function hasCourse(day: number, period: string): boolean {
  return scheduleStore.filteredClasses.some(course => {
    return course.dayOfWeek === day && course.period === period
  })
}

// 获取某个时间段的课程
function getCourse(day: number, period: string) {
  return scheduleStore.filteredClasses.filter(course => {
    return course.dayOfWeek === day && course.period === period
  })
}

// 处理学期变化
function handleSemesterChange() {
  scheduleStore.fetchSchedule()
}

// 处理刷新
function handleRefresh() {
  scheduleStore.fetchSchedule()
}

// 组件挂载时加载课程表
onMounted(() => {
  scheduleStore.fetchSchedule()
})
</script>

<style scoped>
.schedule-container {
  padding: 24px;
  max-width: 1400px;
  margin: 0 auto;
}

.schedule-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.page-title {
  font-size: 1.8rem;
  color: #4a5568;
  margin: 0;
  font-weight: 600;
}

.selectors {
  display: flex;
  align-items: center;
  gap: 24px;
}

.semester-selector {
  display: flex;
  align-items: center;
  gap: 12px;
}

.week-selector {
  display: flex;
  align-items: center;
  gap: 12px;
}

.week-info {
  padding: 10px 16px;
  border-radius: 12px;
  background-color: #e0e5ec;
  color: #4a5568;
  font-size: 1rem;
  font-weight: 500;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.7);
}

.week-btn {
  padding: 10px 16px;
  border-radius: 12px;
  border: none;
  background-color: #e0e5ec;
  color: #4a5568;
  font-size: 0.9rem;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.7);
  cursor: pointer;
  transition: all 0.3s ease;
}

.week-btn:hover {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.7);
}

.week-btn:active {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.7);
}

.neumorphic-select {
  padding: 10px 16px;
  border-radius: 12px;
  border: none;
  background-color: #e0e5ec;
  color: #4a5568;
  font-size: 1rem;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.7);
  cursor: pointer;
  transition: all 0.3s ease;
}

.neumorphic-select:hover {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.7);
}

.neumorphic-select:focus {
  outline: none;
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.7);
}

.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 0;
}

.spinner {
  width: 50px;
  height: 50px;
  border: 4px solid #e0e5ec;
  border-top-color: #4a5568;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.error-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 0;
  gap: 16px;
}

.error-message {
  color: #cf1322;
  font-size: 1.1rem;
}

.schedule-content {
  background-color: #e0e5ec;
  border-radius: 20px;
  padding: 24px;
  box-shadow: 8px 8px 16px rgb(163,177,198,0.7), -8px -8px 16px rgba(255,255,255, 0.5);
}

.schedule-grid {
  display: grid;
  grid-template-columns: 80px repeat(5, 1fr);
  gap: 8px;
}

.schedule-cell {
  background-color: #e0e5ec;
  border-radius: 12px;
  padding: 12px;
  min-height: 100px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.7);
}

.header-cell {
  font-weight: 600;
  color: #4a5568;
  min-height: auto;
  padding: 16px;
}

.period-cell {
  font-weight: 500;
  color: #718096;
  writing-mode: vertical-rl;
  text-orientation: mixed;
  min-height: 100px;
}

.course-cell {
  flex-direction: column;
  align-items: flex-start;
  justify-content: flex-start;
  padding: 8px;
  overflow: hidden;
}

.course-cell.has-course {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.7);
}

.course-card {
  width: 100%;
  padding: 8px;
  border-radius: 8px;
  color: white;
  margin-bottom: 4px;
  box-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.course-card:hover {
  transform: translateY(-2px);
  box-shadow: 4px 4px 8px rgba(0, 0, 0, 0.15);
}

.course-name {
  font-weight: 1000;
  font-size: 0.9rem;
  margin-bottom: 4px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.4);
}

.course-info {
  font-weight: 700;
  font-size: 0.9rem;
  opacity: 0.9;
  margin-bottom: 2px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.4);
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .schedule-grid {
    grid-template-columns: 60px repeat(5, 1fr);
    gap: 6px;
  }

  .schedule-cell {
    padding: 8px;
    min-height: 80px;
  }

  .course-name {
    font-size: 0.7rem;
  }

  .course-info {
    font-size: 0.7rem;
  }
}

@media (max-width: 768px) {
  .schedule-container {
    padding: 16px;
  }

  .schedule-header {
    flex-direction: column;
    gap: 16px;
    align-items: flex-start;
  }

  .selectors {
    width: 100%;
    flex-direction: column;
    align-items: flex-start;
  }

  .semester-selector {
    width: 100%;
  }

  .week-selector {
    width: 100%;
    justify-content: space-between;
  }

  .neumorphic-select {
    flex: 1;
  }

  .week-btn {
    flex: 1;
  }

  .schedule-grid {
    grid-template-columns: 40px repeat(5, 1fr);
    gap: 4px;
  }

  .schedule-cell {
    padding: 4px;
    min-height: 60px;
    font-size: 0.7rem;
  }

  .course-name {
    font-size: 0.7rem;
  }

  .course-info {
    font-size: 0.7rem;
  }
}
</style>
