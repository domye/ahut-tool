<template>
  <div class="coming-soon-container">
    <div class="coming-soon-card">
      <div class="icon-container">
        <div class="icon">🚧</div>
      </div>
      <h1 class="title">功能开发中</h1>
      <p class="description">
        感谢您对安徽工业大学生活工具的关注！<br>
        此功能正在紧锣密鼓地开发中，敬请期待！
      </p>
      <div class="progress-container">
        <div class="progress-bar">
          <div class="progress"></div>
        </div>
        <span class="progress-text">开发进度: 60%</span>
      </div>
      <button class="back-btn" @click="goBack">返回主页</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

function goBack() {
  router.push('/main/home')
}
</script>

<style scoped>
.coming-soon-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: calc(100vh - 112px);
  padding: 24px;
}

.coming-soon-card {
  background-color: #e0e5ec;
  border-radius: 24px;
  padding: 48px;
  box-shadow: 8px 8px 16px rgb(163,177,198,0.6), -8px -8px 16px rgba(255,255,255, 0.5);
  max-width: 500px;
  width: 100%;
  text-align: center;
}

.icon-container {
  margin-bottom: 24px;
}

.icon {
  font-size: 5rem;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.1);
  }
  100% {
    transform: scale(1);
  }
}

.title {
  font-size: 2rem;
  color: #4a5568;
  margin-bottom: 16px;
  font-weight: 600;
}

.description {
  font-size: 1.1rem;
  color: #4a5568;
  line-height: 1.6;
  margin-bottom: 32px;
}

.progress-container {
  margin-bottom: 32px;
}

.progress-bar {
  height: 12px;
  background-color: #e0e5ec;
  border-radius: 6px;
  box-shadow: inset 3px 3px 6px 0 rgba(163,177,198, 0.7), inset -3px -3px 6px 0 rgba(255,255,255, 0.8);
  margin-bottom: 8px;
  overflow: hidden;
}

.progress {
  height: 100%;
  width: 60%;
  background: linear-gradient(90deg, #667eea 0%, #764ba2 100%);
  border-radius: 6px;
  animation: progress 2s ease-in-out infinite;
}

@keyframes progress {
  0% {
    width: 60%;
  }
  50% {
    width: 65%;
  }
  100% {
    width: 60%;
  }
}

.progress-text {
  font-size: 0.9rem;
  color: #718096;
  display: block;
}

.back-btn {
  background-color: #e0e5ec;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
  border-radius: 12px;
  border: none;
  color: #4a5568;
  padding: 12px 32px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.back-btn:hover {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.back-btn:active {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.8);
}
</style>
