// frontend/src/store/electricity.ts
// 电费查询状态管理 Store

import { defineStore } from 'pinia'
import { ref } from 'vue'
import { PayLogin, GetIMS, LoadDormSetting, SettingDorm, LoadPayLogin, ExistPayLoginConfig } from '../../wailsjs/go/backend/App'
import { models } from '../../wailsjs/go/models'

interface IMSResponse {
  Data: models.IMSData
}

export const useElectricityStore = defineStore('electricity', () => {
  const electricity = ref<models.IMSData>({} as models.IMSData)
  const airConditioning = ref<models.IMSData>({} as models.IMSData)
  const loading = ref(false)
  const error = ref('')
  const isLoggedIn = ref(false)
  const message = ref('')

  const xiaoqu = ref('NewS')
  const ld_Name = ref('')
  const ld_Id = ref('')
  const Room_No = ref('')
  const etype = ref('K')

  const dormConfig = ref<models.DormConfig | null>(null)
  const payCredentials = ref<models.PayCredentials | null>(null)

  const handleLoadDormConfig = (): Promise<void> => {
    return new Promise((resolve, reject) => {
      LoadDormSetting()
        .then((config: models.DormConfig) => {
          dormConfig.value = config
          xiaoqu.value = config.campus
          ld_Name.value = config.buildingName
          ld_Id.value = config.buildingId
          Room_No.value = config.roomId
          resolve()
        })
        .catch((err: unknown) => {
          const errorMessage = err instanceof Error ? err.message : String(err)
          error.value = '加载配置失败: ' + errorMessage
          reject(err)
        })
    })
  }

  const handleSaveDormConfig = (
    campus: string,
    buildingId: string,
    buildingName: string,
    roomId: string
  ): Promise<void> => {
    return new Promise((resolve, reject) => {
      SettingDorm(campus, buildingId, buildingName, roomId)
        .then(() => {
          dormConfig.value = { campus, buildingId, buildingName, roomId }
          xiaoqu.value = campus
          ld_Name.value = buildingName
          ld_Id.value = buildingId
          Room_No.value = roomId
          resolve()
        })
        .catch((err: unknown) => {
          const errorMessage = err instanceof Error ? err.message : String(err)
          error.value = '保存配置失败: ' + errorMessage
          reject(err)
        })
    })
  }

  const handleLoadPayConfig = (): Promise<void> => {
    return new Promise((resolve, reject) => {
      LoadPayLogin()
        .then((config: models.PayCredentials) => {
          payCredentials.value = config
          resolve()
        })
        .catch((error: unknown) => {
          console.error('加载缴费系统配置失败:', error)
          reject(error)
        })
    })
  }

  const handleLogin = (): Promise<void> => {
    return new Promise((resolve, reject) => {
      ExistPayLoginConfig()
        .then((exists: boolean) => {
          if (!exists) {
            message.value = '请先配置缴费系统'
            reject(new Error('请先配置缴费系统'))
            return
          }

          return handleLoadPayConfig()
        })
        .then(() => {
          message.value = '登录中...'
          loading.value = true
          if (payCredentials.value) {
            return PayLogin(payCredentials.value)
          } else {
            message.value = '未找到配置，请先配置缴费系统'
            loading.value = false
            return Promise.reject(new Error('未找到配置，请先配置缴费系统'))
          }
        })
        .then((result: number) => {
          if (result === 200) {
            message.value = '登录成功'
            isLoggedIn.value = true
            resolve()
          } else {
            message.value = '登录失败'
            reject(new Error('登录失败'))
          }
        })
        .catch((error: unknown) => {
          if (message.value !== '请先配置缴费系统' && message.value !== '未找到配置，请先配置缴费系统') {
            message.value = '登录失败'
          }
          loading.value = false
          reject(error)
        })
    })
  }

  const handleFetchAirConditioning = (): Promise<void> => {
    return new Promise((resolve, reject) => {
      loading.value = true
      error.value = ''
      GetIMS(xiaoqu.value, ld_Name.value, ld_Id.value, Room_No.value, 'K')
        .then((response: IMSResponse) => {
          airConditioning.value = response.Data
          resolve()
        })
        .catch((err: unknown) => {
          const errorMessage = err instanceof Error ? err.message : String(err)
          error.value = '获取空调电费信息失败: ' + errorMessage
          reject(err)
        })
        .finally(() => {
          loading.value = false
        })
    })
  }

  const handleFetchRoomElectricity = (): Promise<void> => {
    return new Promise((resolve, reject) => {
      loading.value = true
      error.value = ''
      GetIMS(xiaoqu.value, ld_Name.value, ld_Id.value, Room_No.value, 'L')
        .then((response: IMSResponse) => {
          electricity.value = response.Data
          resolve()
        })
        .catch((err: unknown) => {
          const errorMessage = err instanceof Error ? err.message : String(err)
          error.value = '获取房间电费信息失败: ' + errorMessage
          reject(err)
        })
        .finally(() => {
          loading.value = false
        })
    })
  }

  const handleFetchAllElectricity = (): Promise<void> => {
    return new Promise((resolve, reject) => {
      loading.value = true
      error.value = ''

      Promise.all([
        GetIMS(xiaoqu.value, ld_Name.value, ld_Id.value, Room_No.value, 'K'),
        GetIMS(xiaoqu.value, ld_Name.value, ld_Id.value, Room_No.value, 'L')
      ])
        .then(([acResponse, roomResponse]: IMSResponse[]) => {
          airConditioning.value = acResponse.Data
          electricity.value = roomResponse.Data
          resolve()
        })
        .catch((err: unknown) => {
          const errorMessage = err instanceof Error ? err.message : String(err)
          error.value = '获取电费信息失败: ' + errorMessage
          reject(err)
        })
        .finally(() => {
          loading.value = false
        })
    })
  }

  const handleLogout = (): void => {
    isLoggedIn.value = false
    localStorage.removeItem('electricityLoggedIn')
  }

  return {
    electricity,
    airConditioning,
    loading,
    error,
    isLoggedIn,
    message,
    xiaoqu,
    ld_Name,
    ld_Id,
    Room_No,
    etype,
    dormConfig,
    payCredentials,
    loadDormConfig: handleLoadDormConfig,
    saveDormConfig: handleSaveDormConfig,
    loadPayConfig: handleLoadPayConfig,
    login: handleLogin,
    fetchElectricity: handleFetchAllElectricity,
    fetchAirConditioning: handleFetchAirConditioning,
    fetchRoomElectricity: handleFetchRoomElectricity,
    logout: handleLogout
  }
})
