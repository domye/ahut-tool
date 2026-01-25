<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../store/user'
import { useElectricityStore } from '../store/electricity'
import { LoadJwxtLogin, LoadPayLogin, SettingJwxtLogin, SettingPayLogin } from '../../wailsjs/go/backend/App'
import { models } from '../../wailsjs/go/models'
import { message } from 'ant-design-vue'

const router = useRouter()
const userStore = useUserStore()
const electricityStore = useElectricityStore()

// 教务系统配置
const jwxtUser = ref('')
const jwxtPassword = ref('')

// 缴费系统配置
const payUser = ref('')
const payPassword = ref('')

// 加载保存的配置
onMounted(() => {
  loadConfigs()
})

function loadConfigs() {
  LoadJwxtLogin()
    .then((config: models.JwxtCredentials) => {
      jwxtUser.value = config.user
      jwxtPassword.value = config.password
    })
    .catch((error: any) => {
      console.error('加载教务系统配置失败:', error)
    })

  LoadPayLogin()
    .then((config: models.PayCredentials) => {
      payUser.value = config.user
      payPassword.value = config.password
    })
    .catch((error: any) => {
      console.error('加载缴费系统配置失败:', error)
    })
}

function saveJwxtConfig() {
  SettingJwxtLogin(jwxtUser.value, jwxtPassword.value)
    .then(() => {
      message.success('教务系统配置保存成功')
    })
    .catch((error: any) => {
      message.error('保存教务系统配置失败: ' + error)
    })
}

function savePayConfig() {
  SettingPayLogin(payUser.value, payPassword.value)
    .then(() => {
      message.success('缴费系统配置保存成功')
    })
    .catch((error: any) => {
      message.error('保存缴费系统配置失败: ' + error)
    })
}

function goBack() {
  router.back()
}
</script>

<template>
  <div class="config-container">
    <div class="config-header">
      <button class="back-btn" @click="goBack">← 返回</button>
      <h2>系统配置</h2>
    </div>

    <div class="config-content">
      <div class="config-card">
        <h3>教务系统配置</h3>
        <div class="form-group">
          <label for="jwxt-user">学号</label>
          <input
            id="jwxt-user"
            v-model="jwxtUser"
            type="text"
            placeholder="请输入学号"
            class="neumorphic-input"
          />
        </div>
        <div class="form-group">
          <label for="jwxt-password">密码</label>
          <input
            id="jwxt-password"
            v-model="jwxtPassword"
            type="password"
            placeholder="请输入密码"
            class="neumorphic-input"
          />
        </div>
        <button class="save-btn" @click="saveJwxtConfig">保存配置</button>
      </div>

      <div class="config-card">
        <h3>缴费系统配置</h3>
        <div class="form-group">
          <label for="pay-user">学号</label>
          <input
            id="pay-user"
            v-model="payUser"
            type="text"
            placeholder="请输入学号"
            class="neumorphic-input"
          />
        </div>
        <div class="form-group">
          <label for="pay-password">密码</label>
          <input
            id="pay-password"
            v-model="payPassword"
            type="password"
            placeholder="请输入密码"
            class="neumorphic-input"
          />
        </div>
        <button class="save-btn" @click="savePayConfig">保存配置</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.config-container {
  padding: 24px;
  max-width: 800px;
  margin: 0 auto;
}

.config-header {
  display: flex;
  align-items: center;
  margin-bottom: 24px;
}

.back-btn {
  background-color: #e0e5ec;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
  border-radius: 8px;
  border: none;
  color: #4a5568;
  padding: 8px 16px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-right: 16px;
}

.back-btn:hover {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.back-btn:active {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.config-header h2 {
  color: #4a5568;
  font-size: 1.5rem;
  font-weight: 600;
}

.config-content {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 24px;
}

.config-card {
  background-color: #e0e5ec;
  box-shadow: 8px 8px 16px rgb(163,177,198,0.6), -8px -8px 16px rgba(255,255,255, 0.5);
  border-radius: 20px;
  padding: 24px;
  transition: all 0.3s ease;
}

.config-card h3 {
  color: #4a5568;
  font-size: 1.2rem;
  font-weight: 600;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  color: #4a5568;
  font-size: 0.95rem;
  font-weight: 500;
  margin-bottom: 8px;
}

.neumorphic-input {
  width: 100%;
  background-color: #e0e5ec;
  box-shadow: inset 6px 6px 10px 0 rgba(163,177,198, 0.7), inset -6px -6px 10px 0 rgba(255,255,255, 0.8);
  border-radius: 12px;
  border: none;
  color: #4a5568;
  padding: 12px 16px;
  font-size: 0.95rem;
  transition: all 0.3s ease;
}

.neumorphic-input:focus {
  outline: none;
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.save-btn {
  width: 100%;
  background-color: #e0e5ec;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
  border-radius: 12px;
  border: none;
  color: #4a5568;
  padding: 12px;
  font-size: 0.95rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.save-btn:hover {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.save-btn:active {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.8);
}
</style>
