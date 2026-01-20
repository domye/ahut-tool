import { defineStore } from 'pinia'
import { ref } from 'vue'
import { Login } from '../../wailsjs/go/backend/App'

export const useUserStore = defineStore('user', () => {
  const isLoggedIn = ref(false)
  const userId = ref('')
  const message = ref('')

  function login(id: string, password: string): Promise<void> {
    return new Promise((resolve, reject) => {
      if (!id || !password) {
        message.value = '请输入学号和密码'
        reject(new Error('请输入学号和密码'))
        return
      }

      message.value = '登录中...'
      Login(id, password)
        .then((result: number) => {
          console.log(result)
          if (result === 302) {
            message.value = '登录成功'
            isLoggedIn.value = true
            userId.value = id
            localStorage.setItem('isLoggedIn', 'true')
            localStorage.setItem('userId', id)
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
    localStorage.removeItem('isLoggedIn')
    localStorage.removeItem('userId')
  }

  return {
    isLoggedIn,
    userId,
    message,
    login,
    logout
  }
})
