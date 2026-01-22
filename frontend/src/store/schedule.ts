import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { GetClassSchedule } from '../../wailsjs/go/backend/App'
import { models } from '../../wailsjs/go/models'

export const useScheduleStore = defineStore('schedule', () => {
  const classes = ref<models.Class[]>([])
  const loading = ref(false)
  const error = ref('')

  // 当前选中的学期
  const currentSemester = ref('2025-2026-1')

  // 当前选中的周数，默认为第1周
  const currentWeek = ref(1)

  function fetchSchedule(): Promise<void> {
    return new Promise((resolve, reject) => {
      loading.value = true
      error.value = ''
      GetClassSchedule(currentSemester.value)
        .then((response: models.ClassScheduleResponse) => {
          classes.value = response.classes
          resolve()
        })
        .catch((err: any) => {
          error.value = '获取课程表失败: ' + err
          reject(err)
        })
        .finally(() => {
          console.log(classes.value)
          loading.value = false
        })
    })
  }

  // 根据当前周数筛选课程
  const filteredClasses = computed(() => {
    return classes.value.filter((classItem) => {
      // 解析weekNumbers，例如 "1-12,14" 或 "1-11,13-14"
      const weekRanges = classItem.weekNumbers.split(',')
      return weekRanges.some((range) => {
        const [start, end] = range.split('-').map(Number)
        if (isNaN(end)) {
          // 单个周数，如 "14"
          return start === currentWeek.value
        } else {
          // 周数范围，如 "1-12"
          return currentWeek.value >= start && currentWeek.value <= end
        }
      })
    })
  })

  // 设置当前周数
  function setCurrentWeek(week: number) {
    currentWeek.value = week
  }

  return {
    classes,
    filteredClasses,
    loading,
    error,
    currentSemester,
    currentWeek,
    setCurrentWeek,
    fetchSchedule
  }
})
