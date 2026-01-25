<template>
  <a-modal
    v-model:open="visible"
    title="课程表配置"
    :confirm-loading="loading"
    @ok="handleOk"
    @cancel="handleCancel"
    class="schedule-config-modal"
  >
    <a-form
      :model="formState"
      :label-col="{ span: 6 }"
      :wrapper-col="{ span: 16 }"
    >
      <a-form-item label="当前学期" required>
        <a-select
          v-model:value="formState.semester"
          placeholder="请选择学期"
          :options="semesterOptions"
        />
      </a-form-item>
      <a-form-item label="学期开始日期" required>
        <a-date-picker
          v-model:value="formState.startDate"
          placeholder="请选择学期开始日期"
          style="width: 100%"
        />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { message } from 'ant-design-vue'
import { SettingSchedules } from '../../wailsjs/go/backend/App'
import { generateSemesterOptions } from '../utils/semester'
import dayjs, { Dayjs } from 'dayjs'

interface FormState {
  semester: string
  startDate: Dayjs | null
}

const props = defineProps<{
  open: boolean
}>()

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'success'): void
}>()

const visible = computed({
  get: () => props.open,
  set: (value) => emit('update:open', value)
})

const loading = ref(false)
const formState = reactive<FormState>({
  semester: '',
  startDate: null
})

// 学期选项 - 使用工具类生成前后各2年的学期选项
const semesterOptions = computed(() => generateSemesterOptions(2))

const handleOk = async () => {
  if (!formState.semester) {
    message.error('请选择学期')
    return
  }
  if (!formState.startDate) {
    message.error('请选择学期开始日期')
    return
  }

  loading.value = true
  try {
    // 将日期转换为字符串格式
    const startDateStr = formState.startDate.format('YYYY-MM-DD')
    await SettingSchedules(formState.semester, startDateStr)
    message.success('配置保存成功')
    emit('success')
    emit('update:open', false)
  } catch (error) {
    message.error('配置保存失败: ' + error)
  } finally {
    loading.value = false
  }
}

const handleCancel = () => {
  emit('update:open', false)
}
</script>

<style scoped>
.schedule-config-modal :deep(.ant-modal-content) {
  border-radius: 16px;
  box-shadow: 8px 8px 16px rgba(163, 177, 198, 0.7), -8px -8px 16px rgba(255, 255, 255, 0.5);
}

.schedule-config-modal :deep(.ant-modal-header) {
  border-radius: 16px 16px 0 0;
  background-color: #e0e5ec;
  border-bottom: 1px solid rgba(163, 177, 198, 0.3);
}

.schedule-config-modal :deep(.ant-modal-title) {
  color: #4a5568;
  font-weight: 600;
}

.schedule-config-modal :deep(.ant-modal-body) {
  background-color: #e0e5ec;
  padding: 24px;
}

.schedule-config-modal :deep(.ant-modal-footer) {
  background-color: #e0e5ec;
  border-radius: 0 0 16px 16px;
  border-top: 1px solid rgba(163, 177, 198, 0.3);
}

.schedule-config-modal :deep(.ant-form-item-label > label) {
  color: #4a5568;
  font-weight: 500;
}

.schedule-config-modal :deep(.ant-select-selector) {
  background-color: #e0e5ec !important;
  border: none !important;
  border-radius: 8px !important;
  box-shadow: 4px 4px 8px 0 rgba(163, 177, 198, 0.3), -4px -4px 8px 0 rgba(255, 255, 255, 0.8) !important;
  color: #4a5568 !important;
}

.schedule-config-modal :deep(.ant-picker) {
  background-color: #e0e5ec !important;
  border: none !important;
  border-radius: 8px !important;
  box-shadow: 4px 4px 8px 0 rgba(163, 177, 198, 0.3), -4px -4px 8px 0 rgba(255, 255, 255, 0.8) !important;
}

.schedule-config-modal :deep(.ant-btn) {
  background-color: #e0e5ec !important;
  border: none !important;
  border-radius: 8px !important;
  color: #4a5568 !important;
  font-weight: 500;
  box-shadow: 4px 4px 8px 0 rgba(163, 177, 198, 0.3), -4px -4px 8px 0 rgba(255, 255, 255, 0.8) !important;
}

.schedule-config-modal :deep(.ant-btn-primary) {
  background-color: #409eff !important;
  color: white !important;
  box-shadow: 4px 4px 8px 0 rgba(64, 158, 255, 0.3), -4px -4px 8px 0 rgba(150, 200, 255, 0.2) !important;
}
</style>
