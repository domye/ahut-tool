/**
 * 获取当前学期
 * @returns 当前学期字符串，格式为 "YYYY-YYYY-X"，例如 "2024-2025-1"
 */
export function getCurrentSemester(): string {
  const currentDate = new Date()
  const currentYear = currentDate.getFullYear()
  const currentMonth = currentDate.getMonth() + 1 // 1-12

  // 确定基准年份：3月及以后使用当前年份，否则使用上一年份
  const baseYear = currentMonth >= 3 ? currentYear : currentYear - 1

  // 9月到次年3月为第一学期，3月到9月为第二学期
  const semester = currentMonth >= 3 && currentMonth < 9 ? '2' : '1'

  return `${baseYear}-${baseYear + 1}-${semester}`
}

/**
 * 生成学期选项
 * @param range 生成学期的范围，默认为前三年到后三年
 * @returns 学期选项数组
 */
export function generateSemesterOptions(range: number = 3): Array<{ value: string; label: string }> {
  const currentDate = new Date()
  const currentYear = currentDate.getFullYear()
  const currentMonth = currentDate.getMonth() + 1 // 1-12

  // 确定基准年份：9月及以后使用当前年份，否则使用上一年份
  const baseYear = currentMonth >= 9 ? currentYear : currentYear - 1

  // 生成前range年到后range年的学期选项
  return Array.from({ length: range * 2 + 1 }, (_, i) => {
    const year = baseYear + i - range
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
}
