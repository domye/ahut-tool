import { defineStore } from 'pinia'
import { ref } from 'vue'
import { GetGrades } from '../../wailsjs/go/backend/App'
import { models } from '../../wailsjs/go/models'

export const useGradesStore = defineStore('grades', () => {
  const grades = ref<models.Grade[]>([])
  const summary = ref<models.GradeSummary>({} as models.GradeSummary)
  const loading = ref(false)
  const error = ref('')

  // 查询参数
  const kksj = ref('')  // 开课学期
  const kcxz = ref('')  // 课程性质
  const kcmc = ref('')  // 课程名称
  const xsfs = ref('')  // 显示方式

  function fetchGrades(): Promise<void> {
    return new Promise((resolve, reject) => {
      loading.value = true
      error.value = ''
      GetGrades(kksj.value, kcxz.value, kcmc.value, xsfs.value)
        .then((response: any) => {
          grades.value = response.grades
          summary.value = response.summary
          resolve()
        })
        .catch((err: any) => {
          error.value = '获取成绩失败: ' + err
          reject(err)
        })
        .finally(() => {
          loading.value = false
        })
    })
  }

  function resetFilters() {
    kksj.value = ''
    kcxz.value = ''
    kcmc.value = ''
    xsfs.value = ''
  }

  return {
    grades,
    summary,
    loading,
    error,
    kksj,
    kcxz,
    kcmc,
    xsfs,
    fetchGrades,
    resetFilters
  }
})
