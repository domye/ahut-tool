import { defineStore } from 'pinia'
import { ref } from 'vue'
import { PayLogin, GetIMS } from '../../wailsjs/go/backend/App'
import { models } from '../../wailsjs/go/models'

export const useElectricityStore = defineStore('electricity', () => {
  const electricity = ref<models.IMSData>({} as models.IMSData)
  const loading = ref(false)
  const error = ref('')
  const isLoggedIn = ref(false)
  const message = ref('')

  // 查询参数
  const xiaoqu = ref('NewS')  // 校区
  const ld_Name = ref('')     // 楼栋名称
  const ld_Id = ref('')       // 楼栋ID
  const Room_No = ref('')     // 房间号
  const etype = ref('K')      // 类型

  function login(): Promise<void> {
    return new Promise((resolve, reject) => {

      message.value = '登录中...'
      loading.value = true
      PayLogin()
        .then((result: number) => {
          if (result == 200) {
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
        .finally(() => {
          loading.value = false
        })
    })
  }

  function fetchElectricity(): Promise<void> {
    return new Promise((resolve, reject) => {
      loading.value = true
      error.value = ''
      GetIMS(xiaoqu.value, ld_Name.value, ld_Id.value, Room_No.value, etype.value)
        .then((response: any) => {
          electricity.value = response.Data
          resolve()
        })
        .catch((err: any) => {
          error.value = '获取电费信息失败: ' + err
          reject(err)
        })
        .finally(() => {
          loading.value = false
        })
    })
  }


  function logout() {
    isLoggedIn.value = false
    localStorage.removeItem('electricityLoggedIn')
  }

  return {
    electricity,
    loading,
    error,
    isLoggedIn,
    message,
    xiaoqu,
    ld_Name,
    ld_Id,
    Room_No,
    etype,
    login,
    fetchElectricity,
    logout
  }
})
