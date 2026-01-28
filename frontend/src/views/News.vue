
<!-- frontend/src/views/News.vue -->
// 校园新闻页面组件，展示学校新闻、学术通知和公告通知

<template>
  <div class="news-container">
    <h1 class="page-title">校园新闻</h1>

    <div class="tabs">
      <div
        v-for="tab in tabs"
        :key="tab.id"
        class="tab-item"
        :class="{ active: activeTab === tab.id }"
        @click="handleTabChange(tab.id)"
      >
        {{ tab.name }}
      </div>
    </div>

    <div class="news-content">
      <div v-if="loading" class="loading">加载中...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else class="news-list">
        <div
          v-for="news in currentNews"
          :key="news.url"
          class="news-item"
          @click="handleOpenNewsUrl(news.url)"
        >
          <div class="news-title">{{ news.title }}</div>
          <div class="news-date">{{ news.date }}</div>
          <div class="news-preview">{{ news.content }}</div>
        </div>
        <div v-if="currentNews.length === 0" class="empty">暂无新闻</div>
      </div>
    </div>
  </div>
</template>


<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { GetSchoolNews, GetAcademicNotifications, GetAnnouncementNotifications } from '../../wailsjs/go/backend/App'
import { models } from '../../wailsjs/go/models'

interface Tab {
  id: string
  name: string
}

const tabs: Tab[] = [
  { id: 'summary', name: '全部新闻' },
  { id: 'school', name: '学校新闻' },
  { id: 'academic', name: '学术通知' },
  { id: 'announcement', name: '公告通知' }
]

const activeTab = ref<string>('summary')
const loading = ref<boolean>(false)
const error = ref<string>('')
const schoolNews = ref<models.News[]>([])
const academicNews = ref<models.News[]>([])
const announcementNews = ref<models.News[]>([])

const currentNews = ref<models.News[]>([])

async function handleFetchSchoolNews() {
  try {
    loading.value = true
    error.value = ''
    const data = await GetSchoolNews()
    schoolNews.value = data
    updateCurrentNews()
  } catch (err) {
    error.value = '获取学校新闻失败'
    console.error(err)
  } finally {
    loading.value = false
  }
}

async function handleFetchAcademicNews() {
  try {
    loading.value = true
    error.value = ''
    const data = await GetAcademicNotifications()
    academicNews.value = data
    updateCurrentNews()
  } catch (err) {
    error.value = '获取学术通知失败'
    console.error(err)
  } finally {
    loading.value = false
  }
}

async function handleFetchAnnouncementNews() {
  try {
    loading.value = true
    error.value = ''
    const data = await GetAnnouncementNotifications()
    announcementNews.value = data
    updateCurrentNews()
  } catch (err) {
    error.value = '获取公告通知失败'
    console.error(err)
  } finally {
    loading.value = false
  }
}

function updateCurrentNews() {
  switch (activeTab.value) {
    case 'summary':
      const allNews = [...schoolNews.value, ...academicNews.value, ...announcementNews.value]
      currentNews.value = allNews.sort((a, b) => {
        return new Date(b.date).getTime() - new Date(a.date).getTime()
      })
      break
    case 'school':
      currentNews.value = schoolNews.value
      break
    case 'academic':
      currentNews.value = academicNews.value
      break
    case 'announcement':
      currentNews.value = announcementNews.value
      break
  }
}

function handleTabChange(tabId: string) {
  activeTab.value = tabId

  if (tabId === 'summary') {
    if (schoolNews.value.length === 0) {
      handleFetchSchoolNews()
    }
    if (academicNews.value.length === 0) {
      handleFetchAcademicNews()
    }
    if (announcementNews.value.length === 0) {
      handleFetchAnnouncementNews()
    }
  } else {
    if (tabId === 'school' && schoolNews.value.length === 0) {
      handleFetchSchoolNews()
    } else if (tabId === 'academic' && academicNews.value.length === 0) {
      handleFetchAcademicNews()
    } else if (tabId === 'announcement' && announcementNews.value.length === 0) {
      handleFetchAnnouncementNews()
    }
  }

  updateCurrentNews()
}

function handleOpenNewsUrl(url: string) {
  window.open(url, '_blank')
}

onMounted(() => {
  handleFetchSchoolNews()
  handleFetchAcademicNews()
  handleFetchAnnouncementNews()
})
</script>

<style scoped>
.news-container {
  padding: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.page-title {
  font-size: 2rem;
  color: #4a5568;
  margin-bottom: 32px;
  text-align: center;
  font-weight: 600;
}

.tabs {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
  background-color: #e0e5ec;
  padding: 8px;
  border-radius: 16px;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
}

.tab-item {
  flex: 1;
  padding: 12px 24px;
  text-align: center;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  color: #4a5568;
  font-weight: 500;
  box-shadow: 3px 3px 6px 0 rgba(163,177,198, 0.7), -3px -3px 6px 0 rgba(255,255,255, 0.8);
}

.tab-item:hover {
  transform: translateY(-2px);
}

.tab-item.active {
  background-color: #4a5568;
  color: #ffffff;
  box-shadow: inset 3px 3px 6px 0 rgba(0,0,0,0.3), inset -3px -3px 6px 0 rgba(255,255,255, 0.1);
}

.news-content {
  min-height: 400px;
}

.loading,
.error,
.empty {
  text-align: center;
  padding: 40px;
  color: #4a5568;
  font-size: 1.1rem;
}

.error {
  color: #cf1322;
}

.news-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.news-item {
  background-color: #e0e5ec;
  border-radius: 16px;
  padding: 20px;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
}

.news-item:hover {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8);
  transform: translateY(-4px);
}

.news-item:active {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.8);
  transform: translateY(0);
}

.news-title {
  font-size: 1.2rem;
  color: #4a5568;
  font-weight: 600;
  margin-bottom: 8px;
}

.news-date {
  font-size: 0.9rem;
  color: #718096;
  margin-bottom: 12px;
}

.news-preview {
  font-size: 0.95rem;
  color: #718096;
  line-height: 1.6;
  min-height: 60px;
}
</style>
