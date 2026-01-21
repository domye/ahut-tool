import { defineStore } from 'pinia'
import { ref } from 'vue'
import { JwxtLogin } from '../../wailsjs/go/backend/App'

export const useUserStore = defineStore('user', () => {
  const isLoggedIn = ref(false)
  const userId = ref('')
  const message = ref('')

  function login(): Promise<void> {
    return new Promise((resolve, reject) => {
      message.value = '登录中...'
      JwxtLogin()
        .then((result: number) => {
          console.log(result)
          if (result === 302) {
            message.value = '登录成功'
            resolve()
          } else {
            message.value = '登录失败，请检查学号和密码'
            reject(new Error('登录失败，请检查学号和密码'))
          }
        })
        .catch((error: any) => {
          message.value = '登录失败: ' + error
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
    login,
    logout
  }
})
