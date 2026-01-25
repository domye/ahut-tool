<template>
  <a-modal
    v-model:open="visible"
    title="设置宿舍信息"
    @ok="handleSave"
    :confirm-loading="loading"
    width="500px"
    :footer="null"
    class="dorm-setting-modal"
  >
    <a-form layout="vertical" class="setting-form">
      <a-form-item label="校区">
        <a-select
          v-model:value="formData.campus"
          placeholder="请选择校区"
          :options="xiaoquOptions"
          class="neumorphic-select"
        />
      </a-form-item>
      <a-form-item label="楼栋">
        <a-select
          v-model:value="formData.buildingId"
          placeholder="请选择楼栋"
          :options="buildingOptions"
          @change="handleBuildingChange"
          class="neumorphic-select"
        />
      </a-form-item>
      <a-form-item label="房间号">
        <a-input
          v-model:value="formData.roomId"
          placeholder="请输入房间号"
          class="neumorphic-input"
        />
      </a-form-item>
      <div class="form-actions">
        <a-button @click="handleCancel" class="neumorphic-button">取消</a-button>
        <a-button type="primary" @click="handleSave" :loading="loading" class="neumorphic-button primary-button">
          保存
        </a-button>
      </div>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'

interface Props {
  open: boolean
  loading?: boolean
  initialData?: {
    campus: string
    buildingId: string
    buildingName: string
    roomId: string
  }
}

const props = withDefaults(defineProps<Props>(), {
  loading: false,
  initialData: () => ({
    campus: 'NewS',
    buildingId: '',
    buildingName: '',
    roomId: ''
  })
})

const emit = defineEmits<{
  (e: 'update:open', value: boolean): void
  (e: 'save', data: {
    campus: string
    buildingId: string
    buildingName: string
    roomId: string
  }): void
}>()

// 表单数据
const formData = ref({
  campus: props.initialData.campus,
  buildingId: props.initialData.buildingId,
  buildingName: props.initialData.buildingName,
  roomId: props.initialData.roomId
})

// 弹窗显示状态
const visible = computed({
  get: () => props.open,
  set: (value) => emit('update:open', value)
})

// 校区选项
const xiaoquOptions = [
  { value: 'NewS', label: '新校区' },
  { value: 'OldS', label: '老校区' }
]

// 楼栋数据
const newCampusBuildings = [
  { Id: "01", Name: "东校区1号学生宿舍楼" },
  { Id: "02", Name: "东校区2号学生宿舍楼" },
  { Id: "03", Name: "东校区3号学生宿舍楼" },
  { Id: "04", Name: "东校区4号学生宿舍楼" },
  { Id: "05", Name: "东校区5号学生宿舍楼" },
  { Id: "06", Name: "东校区6号学生宿舍楼" },
  { Id: "07", Name: "东校区7号学生宿舍楼" },
  { Id: "08", Name: "东校区8号学生宿舍楼" },
  { Id: "09", Name: "东校区9号学生宿舍楼" },
  { Id: "10", Name: "东校区A号学生宿舍楼" },
  { Id: "11", Name: "东校区B号学生宿舍楼" },
  { Id: "13", Name: "东校区C号学生宿舍楼" },
  { Id: "14", Name: "东校区D号学生宿舍楼" },
  { Id: "15", Name: "东校区E号学生宿舍楼" },
  { Id: "16", Name: "东校区F号学生宿舍楼" },
  { Id: "17", Name: "东校区H1号学生宿舍楼" },
  { Id: "18", Name: "东校区H2号学生宿舍楼" },
  { Id: "19", Name: "东校区H3号学生宿舍楼" },
  { Id: "20", Name: "东校区G1号学生宿舍楼" },
  { Id: "21", Name: "东校区G2号学生宿舍楼" },
  { Id: "22", Name: "东校区G3号学生宿舍楼" },
  { Id: "23", Name: "东校区J1号学生宿舍楼" },
  { Id: "24", Name: "东校区J2号学生宿舍楼" },
  { Id: "25", Name: "东校区J3号学生宿舍楼" },
  { Id: "26", Name: "东校区K1号学生宿舍楼" },
  { Id: "27", Name: "东校区K2号学生宿舍楼" },
  { Id: "28", Name: "东校区K3号学生宿舍楼" },
  { Id: "29", Name: "东校区L1号学生宿舍楼" },
  { Id: "30", Name: "东校区L2号学生宿舍楼" },
  { Id: "31", Name: "东校区G4号学生宿舍楼" },
  { Id: "32", Name: "东校区研5号学生宿舍楼" },
  { Id: "33", Name: "东校区研6号学生宿舍楼" },
  { Id: "34", Name: "东校区研7号学生宿舍楼" },
  { Id: "35", Name: "东校区研8号学生宿舍楼" },
  { Id: "36", Name: "东校区研1号学生宿舍楼" },
  { Id: "37", Name: "东校区研2号学生宿舍楼" },
  { Id: "38", Name: "东校区研3号学生宿舍楼" },
  { Id: "39", Name: "东校区研4号学生宿舍楼" },
  { Id: "40", Name: "东校区M栋南楼宿舍" },
  { Id: "41", Name: "东校区M栋北楼宿舍" },
  { Id: "42", Name: "东校区N栋南楼宿舍" },
  { Id: "43", Name: "东校区N栋北楼宿舍" }
]

const oldCampusBuildings = [
  { Id: "01", Name: "本部校区01栋学生宿舍" },
  { Id: "02", Name: "本部校区02栋学生宿舍" },
  { Id: "03", Name: "本部校区03栋学生宿舍" },
  { Id: "04", Name: "本部校区04栋学生宿舍" },
  { Id: "05", Name: "本部校区05A栋学生宿舍" },
  { Id: "06", Name: "本部校区05B栋学生宿舍" },
  { Id: "07", Name: "本部校区矿东栋学生宿舍" },
  { Id: "08", Name: "本部校区矿西栋学生宿舍" },
  { Id: "09", Name: "本部校区研A栋学生宿舍" },
  { Id: "10", Name: "本部校区研B栋学生宿舍" },
  { Id: "11", Name: "本部校区8号公寓" },
  { Id: "12", Name: "本部校区7号公寓" }
]

// 楼栋选项（根据校区动态生成）
const buildingOptions = computed(() => {
  const buildings = formData.value.campus === 'OldS' ? oldCampusBuildings : newCampusBuildings
  return buildings.map(building => ({
    value: building.Id,
    label: building.Name
  }))
})

// 处理楼栋选择变化
function handleBuildingChange(value: string) {
  const buildings = formData.value.campus === 'OldS' ? oldCampusBuildings : newCampusBuildings
  const selected = buildings.find(b => b.Id === value)
  if (selected) {
    formData.value.buildingName = selected.Name
  }
}

// 监听校区变化，重置楼栋选择
watch(() => formData.value.campus, () => {
  formData.value.buildingId = ''
  formData.value.buildingName = ''
})

// 监听弹窗打开，初始化表单数据
watch(() => props.open, (newVal) => {
  if (newVal) {
    formData.value = { ...props.initialData }
  }
})

// 保存设置
function handleSave() {
  emit('save', {
    campus: formData.value.campus,
    buildingId: formData.value.buildingId,
    buildingName: formData.value.buildingName,
    roomId: formData.value.roomId
  })
}

// 取消设置
function handleCancel() {
  visible.value = false
}
</script>

<style scoped>
.dorm-setting-modal :deep(.ant-modal-content) {
  background-color: #e0e5ec;
  border-radius: 20px;
  box-shadow: 8px 8px 16px rgba(163, 177, 198, 0.6), -8px -8px 16px rgba(255, 255, 255, 0.8);
  overflow: hidden;
}

.dorm-setting-modal :deep(.ant-modal-header) {
  background-color: transparent;
  border-bottom: none;
  padding: 24px 24px 16px;
}

.dorm-setting-modal :deep(.ant-modal-title) {
  color: #4a5568;
  font-weight: 600;
  font-size: 1.2rem;
}

.dorm-setting-modal :deep(.ant-modal-body) {
  padding: 16px 24px 24px;
}

.dorm-setting-modal :deep(.ant-modal-close) {
  top: 20px;
  right: 20px;
}

.dorm-setting-modal :deep(.ant-modal-close-x) {
  width: 32px;
  height: 32px;
  line-height: 32px;
  color: #718096;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.dorm-setting-modal :deep(.ant-modal-close-x:hover) {
  background-color: rgba(74, 85, 104, 0.1);
  color: #4a5568;
}

.setting-form {
  padding: 8px 0;
}

.setting-form :deep(.ant-form-item) {
  margin-bottom: 24px;
}

.setting-form :deep(.ant-form-item-label > label) {
  color: #4a5568;
  font-weight: 500;
  font-size: 0.95rem;
}

.setting-form :deep(.ant-form-item-explain-error) {
  color: #c53030;
  font-size: 0.85rem;
}

.neumorphic-select {
  width: 100%;
}

.neumorphic-select :deep(.ant-select-selector) {
  background-color: #e0e5ec;
  border: none;
  border-radius: 12px;
  box-shadow: inset 4px 4px 8px rgba(163, 177, 198, 0.3), inset -4px -4px 8px rgba(255, 255, 255, 0.8);
  height: 44px;
  padding: 0 16px;
  transition: all 0.3s ease;
}

.neumorphic-select :deep(.ant-select-selection-placeholder) {
  color: #a0aec0;
  line-height: 44px;
}

.neumorphic-select :deep(.ant-select-selection-item) {
  line-height: 44px;
  color: #4a5568;
}

.neumorphic-select:hover :deep(.ant-select-selector) {
  box-shadow: inset 5px 5px 10px rgba(163, 177, 198, 0.4), inset -5px -5px 10px rgba(255, 255, 255, 0.9);
}

.neumorphic-select-focused :deep(.ant-select-selector) {
  box-shadow: inset 5px 5px 10px rgba(163, 177, 198, 0.4), inset -5px -5px 10px rgba(255, 255, 255, 0.9);
}

.neumorphic-select :deep(.ant-select-arrow) {
  color: #718096;
  right: 16px;
}

.neumorphic-input {
  width: 100%;
}

.neumorphic-input :deep(.ant-input) {
  background-color: #e0e5ec;
  border: none;
  border-radius: 12px;
  box-shadow: inset 4px 4px 8px rgba(163, 177, 198, 0.3), inset -4px -4px 8px rgba(255, 255, 255, 0.8);
  height: 44px;
  padding: 0 16px;
  color: #4a5568;
  transition: all 0.3s ease;
}

.neumorphic-input :deep(.ant-input::placeholder) {
  color: #a0aec0;
}

.neumorphic-input:hover :deep(.ant-input) {
  box-shadow: inset 5px 5px 10px rgba(163, 177, 198, 0.4), inset -5px -5px 10px rgba(255, 255, 255, 0.9);
}

.neumorphic-input :deep(.ant-input:focus) {
  box-shadow: inset 5px 5px 10px rgba(163, 177, 198, 0.4), inset -5px -5px 10px rgba(255, 255, 255, 0.9);
  outline: none;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 8px;
}

.neumorphic-button {
  background-color: #e0e5ec;
  border: none;
  border-radius: 12px;
  color: #4a5568;
  font-weight: 500;
  height: 44px;
  padding: 0 24px;
  box-shadow: 6px 6px 12px rgba(163, 177, 198, 0.4), -6px -6px 12px rgba(255, 255, 255, 0.8);
  transition: all 0.3s ease;
}

.neumorphic-button:hover {
  box-shadow: 8px 8px 16px rgba(163, 177, 198, 0.5), -8px -8px 16px rgba(255, 255, 255, 0.9);
  transform: translateY(-2px);
}

.neumorphic-button:active {
  box-shadow: inset 4px 4px 8px rgba(163, 177, 198, 0.4), inset -4px -4px 8px rgba(255, 255, 255, 0.8);
  transform: translateY(0);
}

.primary-button {
  background-color: #409eff;
  color: white;
  box-shadow: 6px 6px 12px rgba(64, 158, 255, 0.3), -6px -6px 12px rgba(150, 200, 255, 0.2);
}

.primary-button:hover {
  background-color: #66b1ff;
  box-shadow: 8px 8px 16px rgba(64, 158, 255, 0.4), -8px -8px 16px rgba(150, 200, 255, 0.3);
}

.primary-button:active {
  background-color: #3a8ee6;
  box-shadow: inset 4px 4px 8px rgba(64, 158, 255, 0.3), inset -4px -4px 8px rgba(150, 200, 255, 0.2);
}
</style>
