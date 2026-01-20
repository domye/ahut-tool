<script lang="ts" setup>
import { reactive } from 'vue'
import { Login } from '../../wailsjs/go/backend/App'

const formData = reactive({
  userId: '',
  password: '',
  message: ''
})

function handleLogin() {
  if (!formData.userId || !formData.password) {
    formData.message = '请输入学号和密码'
    return
  }

  formData.message = '登录中...'
  Login(formData.userId, formData.password).then((result: number) => {
    console.log(result)
    if (result == 302) {
      formData.message = '登录成功'
      // 登录成功后，触发自定义事件通知父组件
      const event = new CustomEvent('login-success')
      window.dispatchEvent(event)
    } else {
      formData.message = '登录失败，请检查学号和密码'
    }
  }).catch((error: any) => {
    formData.message = '登录失败: ' + error
  })
}
</script>

<template>
  <div class="login-container">
    <div class="login-box">
      <h2>安徽工业大学成绩查询</h2>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="userId">学号</label>
          <input 
            id="userId" 
            v-model="formData.userId" 
            type="text" 
            placeholder="请输入学号"
            autocomplete="off"
          />
        </div>
        <div class="form-group">
          <label for="password">密码</label>
          <input 
            id="password" 
            v-model="formData.password" 
            type="password" 
            placeholder="请输入密码"
          />
        </div>
        <div class="message">{{ formData.message }}</div>
        <button type="submit" class="login-btn">登录</button>
      </form>
    </div>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-box {
  background: white;
  padding: 2rem;
  border-radius: 10px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
  width: 100%;
  max-width: 400px;
}

h2 {
  text-align: center;
  color: #333;
  margin-bottom: 2rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  color: #555;
  font-weight: 500;
}

input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 5px;
  font-size: 1rem;
  transition: border-color 0.3s;
}

input:focus {
  outline: none;
  border-color: #667eea;
}

.message {
  text-align: center;
  margin-bottom: 1rem;
  min-height: 1.5rem;
  color: #e74c3c;
}

.login-btn {
  width: 100%;
  padding: 0.75rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 5px;
  font-size: 1rem;
  cursor: pointer;
  transition: transform 0.2s;
}

.login-btn:hover {
  transform: translateY(-2px);
}

.login-btn:active {
  transform: translateY(0);
}
</style>
