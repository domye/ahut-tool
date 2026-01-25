import { defineStore } from 'pinia'
import { ref } from 'vue'
import { PayLogin, GetIMS, LoadDormSetting, SettingDorm, LoadPayLogin, ExistPayLoginConfig } from '../../wailsjs/go/backend/App'
import { models } from '../../wailsjs/go/models'

export const useElectricityStore = defineStore('electricity', () => {
  const electricity = ref<models.IMSData>({} as models.IMSData)
  const airConditioning = ref<models.IMSData>({} as models.IMSData)
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

  // 宿舍配置
  const dormConfig = ref<models.DormConfig | null>(null)

  // 缴费系统配置
  const payCredentials = ref<models.PayCredentials | null>(null)

  // 加载宿舍配置
  function loadDormConfig(): Promise<void> {
    return new Promise((resolve, reject) => {
      LoadDormSetting()
        .then((config: models.DormConfig) => {
          dormConfig.value = config
          // 应用配置到查询参数
          xiaoqu.value = config.campus
          ld_Name.value = config.buildingName
          ld_Id.value = config.buildingId
          Room_No.value = config.roomId
          resolve()
        })
        .catch((err: any) => {
          error.value = '加载配置失败: ' + err
          reject(err)
        })
    })
  }

  // 保存宿舍配置
  function saveDormConfig(campus: string, buildingId: string, buildingName: string, roomId: string): Promise<void> {
    return new Promise((resolve, reject) => {
      SettingDorm(campus, buildingId, buildingName, roomId)
        .then(() => {
          // 更新本地配置
          dormConfig.value = {
            campus,
            buildingId,
            buildingName,
            roomId
          }
          // 应用配置到查询参数
          xiaoqu.value = campus
          ld_Name.value = buildingName
          ld_Id.value = buildingId
          Room_No.value = roomId
          resolve()
        })
        .catch((err: any) => {
          error.value = '保存配置失败: ' + err
          reject(err)
        })
    })
  }

  // 加载缴费系统配置
  function loadPayConfig(): Promise<void> {
    return new Promise((resolve, reject) => {
      LoadPayLogin()
        .then((config: models.PayCredentials) => {
          payCredentials.value = config
          resolve()
        })
        .catch((error: any) => {
          console.error('加载缴费系统配置失败:', error)
          reject(error)
        })
    })
  }

  function login(): Promise<void> {
    return new Promise((resolve, reject) => {
      // 先检查配置是否存在
      ExistPayLoginConfig()
        .then((exists: boolean) => {
          if (!exists) {
            message.value = '请先配置缴费系统'
            reject(new Error('请先配置缴费系统'))
            return
          }

          // 加载配置
          return loadPayConfig()
        })
        .then(() => {
          message.value = '登录中...'
          loading.value = true
          // 使用配置进行登录
          if (payCredentials.value) {
            return PayLogin(payCredentials.value)
          } else {
            message.value = '未找到配置，请先配置缴费系统'
            loading.value = false
            return Promise.reject(new Error('未找到配置，请先配置缴费系统'))
          }
        })
        .then((result: number) => {
          if (result == 200) {
            message.value = '登录成功'
            isLoggedIn.value = true
            resolve()
          } else {
            message.value = '登录失败'
            reject(new Error('登录失败'))
          }
        })
        .catch((error: any) => {
          if (message.value !== '请先配置缴费系统' && message.value !== '未找到配置，请先配置缴费系统') {
            message.value = '登录失败'
          }
          loading.value = false
          reject(error)
        })
    })
  }

  // 查询空调电费
  function fetchAirConditioning(): Promise<void> {
    return new Promise((resolve, reject) => {
      loading.value = true
      error.value = ''
      GetIMS(xiaoqu.value, ld_Name.value, ld_Id.value, Room_No.value, 'K')
        .then((response: any) => {
          airConditioning.value = response.Data
          resolve()
        })
        .catch((err: any) => {
          error.value = '获取空调电费信息失败: ' + err
          reject(err)
        })
        .finally(() => {
          loading.value = false
        })
    })
  }

  // 查询房间电费
  function fetchRoomElectricity(): Promise<void> {
    return new Promise((resolve, reject) => {
      loading.value = true
      error.value = ''
      GetIMS(xiaoqu.value, ld_Name.value, ld_Id.value, Room_No.value, 'L')
        .then((response: any) => {
          electricity.value = response.Data
          resolve()
        })
        .catch((err: any) => {
          error.value = '获取房间电费信息失败: ' + err
          reject(err)
        })
        .finally(() => {
          loading.value = false
        })
    })
  }

  // 同时查询空调和房间电费
  function fetchAllElectricity(): Promise<void> {
    return new Promise((resolve, reject) => {
      loading.value = true
      error.value = ''

      Promise.all([
        GetIMS(xiaoqu.value, ld_Name.value, ld_Id.value, Room_No.value, 'K'),
        GetIMS(xiaoqu.value, ld_Name.value, ld_Id.value, Room_No.value, 'L')
      ])
        .then(([acResponse, roomResponse]: any[]) => {
          airConditioning.value = acResponse.Data
          electricity.value = roomResponse.Data
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
    loadDormConfig,
    saveDormConfig,
    loadPayConfig,
    login,
    fetchElectricity: fetchAllElectricity,
    fetchAirConditioning,
    fetchRoomElectricity,
    logout
  }
})
