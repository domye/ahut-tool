import { defineStore } from 'pinia'
import { ref } from 'vue'
import { GetClassSchedule } from '../../wailsjs/go/backend/App'
import { models } from '../../wailsjs/go/models'

export const useScheduleStore = defineStore('schedule', () => {
  const classes = ref<models.Class[]>([])
  const loading = ref(false)
  const error = ref('')

  // 当前选中的学期
  const currentSemester = ref('2025-2026-1')

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
          loading.value = false
        })
    })
  }

  return {
    classes,
    loading,
    error,
    currentSemester,
    fetchSchedule
  }
})
