import { defineStore } from 'pinia'
import { ref } from 'vue'
import { JwxtLogin, LoadJwxtLogin, ExistJwxtLoginConfig } from '../../wailsjs/go/backend/App'
import { models } from '../../wailsjs/go/models'

export const useUserStore = defineStore('user', () => {
  const isLoggedIn = ref(false)
  const userId = ref('')
  const message = ref('')
  const jwxtCredentials = ref<models.JwxtCredentials | null>(null)

  // 加载教务系统配置
  function loadJwxtConfig(): Promise<void> {
    return new Promise((resolve, reject) => {
      LoadJwxtLogin()
        .then((config: models.JwxtCredentials) => {
          jwxtCredentials.value = config
          userId.value = config.user
          resolve()
        })
        .catch((error: any) => {
          console.error('加载教务系统配置失败:', error)
          reject(error)
        })
    })
  }

  function login(): Promise<void> {
    return new Promise((resolve, reject) => {
      // 先检查配置是否存在
      ExistJwxtLoginConfig()
        .then((exists: boolean) => {
          if (!exists) {
            message.value = '请先配置教务系统'
            reject(new Error('请先配置教务系统'))
            return
          }

          // 加载配置
          return loadJwxtConfig()
        })
        .then(() => {
          message.value = '登录中...'
          // 使用配置进行登录
          if (jwxtCredentials.value) {
            return JwxtLogin(jwxtCredentials.value)
          } else {
            message.value = '未找到配置，请先配置教务系统'
            return Promise.reject(new Error('未找到配置，请先配置教务系统'))
          }
        })
        .then((result: number) => {
          console.log(result)
          if (result === 302) {
            message.value = '登录成功'
            isLoggedIn.value = true
            resolve()
          } else {
            message.value = '登录失败'
            reject(new Error('登录失败'))
          }
        })
        .catch((error: any) => {
          if (message.value !== '请先配置教务系统' && message.value !== '未找到配置，请先配置教务系统') {
            message.value = '登录失败'
          }
          reject(error)
        })
    })
  }

  function logout() {
    isLoggedIn.value = false
    userId.value = ''
  }

  return {
    isLoggedIn,
    userId,
    message,
    jwxtCredentials,
    loadJwxtConfig,
    login,
    logout
  }
})
