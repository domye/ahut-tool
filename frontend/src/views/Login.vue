<template>
  <div class="login-container">
    <div class="login-card">
      <div class="login-title">登录</div>

      <form @submit.prevent="handleLogin" class="login-form">
        <div class="form-group">
          <label class="form-label">学号</label>
          <input
            v-model="formData.userId"
            type="text"
            placeholder="请输入学号"
            class="form-input"
            required
          />
        </div>

        <div class="form-group">
          <label class="form-label">密码</label>
          <input
            v-model="formData.password"
            type="password"
            placeholder="请输入密码"
            class="form-input"
            required
          />
        </div>

        <button type="submit" class="login-button" :disabled="loading">
          {{ loading ? '登录中...' : '登录' }}
        </button>

        <div v-if="userStore.message" :class="['alert', userStore.isLoggedIn ? 'alert-success' : 'alert-error']">
          {{ userStore.message }}
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../store/user'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)

const formData = reactive({
  userId: '',
  password: ''
})

async function handleLogin() {
  loading.value = true
  try {
    await userStore.login(formData.userId, formData.password)
    router.push('/main/home')
  } catch (error) {
    console.error('登录失败:', error)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #e0e5ec;
  padding: 20px;
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: 40px;
  border-radius: 30px;
  background-color: #e0e5ec;
  box-shadow: 20px 20px 60px #bec3c9, -20px -20px 60px #ffffff;
}

.login-title {
  font-size: 1.8rem;
  font-weight: 600;
  color: #4a5568;
  text-align: center;
  margin-bottom: 32px;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-label {
  font-size: 0.95rem;
  color: #4a5568;
  font-weight: 500;
  margin-left: 4px;
}

.form-input {
  padding: 14px 20px;
  border: none;
  border-radius: 12px;
  background-color: #e0e5ec;
  box-shadow: inset 6px 6px 12px #bec3c9, inset -6px -6px 12px #ffffff;
  font-size: 1rem;
  color: #4a5568;
  outline: none;
  transition: all 0.3s ease;
}

.form-input::placeholder {
  color: #8b9bb4;
}

.form-input:focus {
  box-shadow: inset 4px 4px 8px #bec3c9, inset -4px -4px 8px #ffffff;
}

.login-button {
  padding: 14px 20px;
  border: none;
  border-radius: 12px;
  background-color: #e0e5ec;
  color: #4a5568;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  box-shadow: 6px 6px 12px #bec3c9, -6px -6px 12px #ffffff;
  transition: all 0.3s ease;
  margin-top: 8px;
}

.login-button:hover:not(:disabled) {
  box-shadow: 4px 4px 8px #bec3c9, -4px -4px 8px #ffffff;
  transform: translateY(-2px);
}

.login-button:active:not(:disabled) {
  box-shadow: inset 4px 4px 8px #bec3c9, inset -4px -4px 8px #ffffff;
  transform: translateY(0);
}

.login-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.alert {
  padding: 12px 16px;
  border-radius: 10px;
  font-size: 0.9rem;
  margin-top: 8px;
  background-color: #e0e5ec;
  box-shadow: inset 4px 4px 8px #bec3c9, inset -4px -4px 8px #ffffff;
}

.alert-success {
  color: #3f8600;
}

.alert-error {
  color: #cf1322;
}
</style>
