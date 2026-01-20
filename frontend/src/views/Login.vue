<template>
  <div class="login-container">
    <a-card class="login-card" :bordered="false">
      <template #title>
        <div class="login-title">安徽工业大学成绩查询</div>
      </template>

      <a-form
        :model="formData"
        @finish="handleLogin"
        layout="vertical"
      >
        <a-form-item
          label="学号"
          name="userId"
          :rules="[{ required: true, message: '请输入学号' }]"
        >
          <a-input
            v-model:value="formData.userId"
            placeholder="请输入学号"
            size="large"
          />
        </a-form-item>

        <a-form-item
          label="密码"
          name="password"
          :rules="[{ required: true, message: '请输入密码' }]"
        >
          <a-input-password
            v-model:value="formData.password"
            placeholder="请输入密码"
            size="large"
          />
        </a-form-item>

        <a-form-item>
          <a-button
            type="primary"
            html-type="submit"
            size="large"
            block
            :loading="loading"
          >
            登录
          </a-button>
        </a-form-item>

        <a-alert
          v-if="userStore.message"
          :message="userStore.message"
          :type="userStore.isLoggedIn ? 'success' : 'error'"
          show-icon
          style="margin-top: 16px"
        />
      </a-form>
    </a-card>
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
    router.push('/main/grades')
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
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-card {
  width: 100%;
  max-width: 400px;
}

.login-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #333;
  text-align: center;
}
</style>
