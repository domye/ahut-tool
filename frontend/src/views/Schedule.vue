<template>
  <div class="schedule-page">
    <!-- 查询条件 -->
    <div class="filter-section">
      <a-form layout="inline" class="filter-form">
        <a-form-item label="学期">
          <a-select
            v-model:value="scheduleStore.currentSemester"
            @change="handleSemesterChange"
            placeholder="请选择学期"
            style="width: 200px"
            allow-clear
            :options="semesterOptions"
            popup-class-name="neumorphic-select-dropdown"
          />
        </a-form-item>
        <a-form-item label="当前周次">
          <a-space>
            <a-button @click="handlePrevWeek" class="neumorphic-button">
              <template #icon>
                <LeftOutlined />
              </template>
              上一周
            </a-button>
            <a-tag class="week-tag">第 {{ scheduleStore.currentWeek }} 周</a-tag>
            <a-button @click="handleNextWeek" class="neumorphic-button">
              下一周
              <template #icon>
                <RightOutlined />
              </template>
            </a-button>
          </a-space>
        </a-form-item>
        <a-form-item>
          <a-space :size="20">
            <a-button
              type="primary"
              @click="handleRefresh"
              :loading="scheduleStore.loading"
              class="neumorphic-button query-button"
            >
              <template #icon>
                <ReloadOutlined />
              </template>
              刷新
            </a-button>
          </a-space>
        </a-form-item>
      </a-form>
      <a-alert
        v-if="scheduleStore.error"
        :message="scheduleStore.error"
        type="error"
        show-icon
        closable
        @close="scheduleStore.error = ''"
        style="margin-top: 16px"
      />
    </div>



    <div v-if="!scheduleStore.loading && !scheduleStore.error" class="schedule-content">
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
              v-for="(course, index) in getCourse(day, period)"
              :key="course.name + '-' + index"
              class="course-card"
              :class="{ 'single-course': getCourse(day, period).length === 1, 'multiple-courses': getCourse(day, period).length > 1 }"
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
import { LeftOutlined, RightOutlined, ReloadOutlined } from '@ant-design/icons-vue'
import { generateSemesterOptions } from '../utils/semester'

const scheduleStore = useScheduleStore()

const weekDays = ['周一', '周二', '周三', '周四', '周五']
const periods = ['1-2', '3-4', '5-6', '7-8', '9-10-11']

// 学期选项 - 使用工具类生成前后各2年的学期选项
const semesterOptions = computed(() => generateSemesterOptions(2))

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
.schedule-page {
  padding: 24px;
  background-color: #e0e5ec;
  min-height: calc(100vh - 48px);
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

/* 按钮新拟态风格 */
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

/* 标签样式 */
:deep(.ant-tag) {
  border: none !important;
  font-weight: 500;
  border-radius: 8px;
  padding: 4px 12px;
  box-shadow: 3px 3px 6px 0 rgba(163,177,198, 0.2), -3px -3px 6px 0 rgba(255,255,255, 0.8);
}

.week-tag {
  padding: 8px 16px;
  font-size: 1rem;
  font-weight: 500;
  border-radius: 12px;
  background-color: #e0e5ec !important;
  color: #4a5568 !important;
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.3), -4px -4px 8px 0 rgba(255,255,255, 0.8) !important;
}

/* 下拉菜单新拟态风格 - 全局样式 */
:global(.neumorphic-select-dropdown) {
  background-color: #e0e5ec !important;
  border: none !important;
  border-radius: 12px !important;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8) !important;
  padding: 8px !important;
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

.schedule-content {
  background-color: #e0e5ec;
  border-radius: 20px;
  padding: 24px;
  box-shadow: 8px 8px 16px rgb(163,177,198,0.7), -8px -8px 16px rgba(255,255,255, 0.5);
}

.schedule-grid {
  display: grid;
  grid-template-columns: 80px repeat(5, 1fr);
  grid-auto-rows: minmax(120px, auto);
  gap: 8px;
}

.schedule-cell {
  background-color: #e0e5ec;
  border-radius: 12px;
  padding: 12px;
  min-height: 120px;
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
  min-height: 120px;
}

.course-cell {
  flex-direction: column;
  align-items: stretch;
  justify-content: flex-start;
  padding: 8px;
  overflow: hidden;
  display: flex;
}

.course-cell.has-course {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.7);
}

.course-card {
  width: 100%;
  padding: 10px;
  border-radius: 8px;
  color: white;
  margin-bottom: 4px;
  box-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.course-card.single-course {
  height: calc(100% - 4px); /* 减去margin-bottom */
}

.course-card.multiple-courses {
  height: auto;
}

.course-card:hover {
  transform: translateY(-2px);
  box-shadow: 4px 4px 8px rgba(0, 0, 0, 0.15);
}

.course-name {
  font-weight: 1000;
  font-size: 0.9rem;
  margin-bottom: 4px;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.4);
  word-break: break-all;
  white-space: normal;
  line-height: 1.2;
}

.course-info {
  font-weight: 700;
  font-size: 0.9rem;
  opacity: 0.9;
  margin-bottom: 2px;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.4);
  word-break: break-all;
  white-space: normal;
  line-height: 1.2;
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .schedule-grid {
    grid-template-columns: 60px repeat(5, 1fr);
    grid-auto-rows: minmax(100px, auto);
    gap: 6px;
  }

  .schedule-cell {
    padding: 8px;
    min-height: 100px;
  }

  .course-name {
    font-size: 0.7rem;
    line-height: 1.2;
  }

  .course-info {
    font-size: 0.7rem;
    line-height: 1.2;
  }
}

@media (max-width: 768px) {
  .schedule-page {
    padding: 16px;
  }

  .filter-form {
    flex-direction: column;
    align-items: stretch;
  }

  :deep(.filter-form .ant-form-item) {
    width: 100%;
    margin-bottom: 16px;
  }

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

  .schedule-grid {
    grid-template-columns: 40px repeat(5, 1fr);
    grid-auto-rows: minmax(80px, auto);
    gap: 4px;
  }

  .schedule-cell {
    padding: 4px;
    min-height: 80px;
    font-size: 0.7rem;
  }

  .course-name {
    font-size: 0.7rem;
    line-height: 1.2;
  }

  .course-info {
    font-size: 0.7rem;
    line-height: 1.2;
  }
}
</style>
