// frontend/src/store/user.ts
// 教务系统用户状态管理 Store

import { defineStore } from 'pinia'
import { ref } from 'vue'
import { JwxtLogin, LoadJwxtLogin, ExistJwxtLoginConfig } from '../../wailsjs/go/backend/App'
import { models } from '../../wailsjs/go/models'

interface LoginResponse {
  status: number
  error?: string
}

export const useUserStore = defineStore('user', () => {
  const isLoggedIn = ref(false)
  const userId = ref('')
  const message = ref('')
  const jwxtCredentials = ref<models.JwxtCredentials | null>(null)

  const handleLoadJwxtConfig = (): Promise<void> => {
    return new Promise((resolve, reject) => {
      LoadJwxtLogin()
        .then((config: models.JwxtCredentials) => {
          jwxtCredentials.value = config
          userId.value = config.user
          resolve()
        })
        .catch((error: unknown) => {
          console.error('加载教务系统配置失败:', error)
          reject(error)
        })
    })
  }

  const handleLogin = (): Promise<void> => {
    return new Promise((resolve, reject) => {
      ExistJwxtLoginConfig()
        .then((exists: boolean) => {
          if (!exists) {
            message.value = '请先配置教务系统'
            reject(new Error('请先配置教务系统'))
            return
          }

          return handleLoadJwxtConfig()
        })
        .then(() => {
          message.value = '登录中...'
          if (jwxtCredentials.value) {
            return JwxtLogin(jwxtCredentials.value)
          } else {
            message.value = '未找到配置，请先配置教务系统'
            return Promise.reject(new Error('未找到配置，请先配置教务系统'))
          }
        })
        .then((result: number) => {
          if (result === 302) {
            message.value = '登录成功'
            isLoggedIn.value = true
            resolve()
          } else {
            message.value = '登录失败'
            reject(new Error('登录失败'))
          }
        })
        .catch((error: unknown) => {
          if (message.value !== '请先配置教务系统' && message.value !== '未找到配置，请先配置教务系统') {
            message.value = '登录失败'
          }
          reject(error)
        })
    })
  }

  const handleLogout = (): void => {
    isLoggedIn.value = false
    userId.value = ''
  }

  return {
    isLoggedIn,
    userId,
    message,
    jwxtCredentials,
    loadJwxtConfig: handleLoadJwxtConfig,
    login: handleLogin,
    logout: handleLogout
  }
})
