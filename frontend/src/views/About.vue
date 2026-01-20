<template>
  <div class="about-container">
    <div class="about-card">
      <h1 class="about-title">关于安徽工业大学生活工具</h1>

      <div class="about-section">
        <h2 class="section-title">项目介绍</h2>
        <p class="section-content">
          安徽工业大学生活工具是一个专为安徽工业大学师生设计的综合服务平台，旨在提供便捷的校园生活服务。
          本项目整合了校园内常用的功能模块，为师生提供一站式的校园生活解决方案。
        </p>
      </div>

      <div class="about-section">
        <h2 class="section-title">主要功能</h2>
        <ul class="feature-list">
          <li>成绩查询 - 快速查询学业成绩和学分绩点</li>
          <li>课程表 - 查看和管理个人课程安排</li>
          <li>图书馆 - 图书查询、借阅状态查看</li>
          <li>校园地图 - 校园建筑和设施导航</li>
          <li>食堂查询 - 食堂菜品和营业时间</li>
          <li>校车时刻 - 校园班车时刻表</li>
          <li>校园活动 - 校园活动和社团信息</li>
          <li>校园新闻 - 最新校园资讯和通知</li>
        </ul>
      </div>

      <div class="about-section">
        <h2 class="section-title">技术特点</h2>
        <p class="section-content">
          本项目采用现代化的Web技术栈开发，界面采用新拟态(Neumorphism)设计风格，提供流畅的用户体验。
          前端使用Vue 3框架，后端采用高性能的服务架构，确保系统稳定可靠。
        </p>
      </div>

      <div class="about-section">
        <h2 class="section-title">联系我们</h2>
        <p class="section-content">
          如有任何问题或建议，欢迎通过以下方式联系我们：
        </p>
        <div class="contact-info">
          <div class="contact-item">
            <span class="contact-icon">📧</span>
            <span class="contact-text">support@ahut.edu.cn</span>
          </div>
          <div class="contact-item">
            <span class="contact-icon">📱</span>
            <span class="contact-text">微信公众号：安徽工业大学</span>
          </div>
        </div>
      </div>

      <div class="about-footer">
        <p>© 2024 安徽工业大学 版权所有</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// 关于页面逻辑
</script>

<style scoped>
.about-container {
  padding: 24px;
  max-width: 900px;
  margin: 0 auto;
}

.about-card {
  background-color: #e0e5ec;
  border-radius: 24px;
  padding: 32px;
  box-shadow: 8px 8px 16px rgb(163,177,198,0.6), -8px -8px 16px rgba(255,255,255, 0.5);
}

.about-title {
  font-size: 2rem;
  color: #4a5568;
  margin-bottom: 32px;
  text-align: center;
  font-weight: 600;
}

.about-section {
  margin-bottom: 32px;
}

.section-title {
  font-size: 1.4rem;
  color: #4a5568;
  margin-bottom: 16px;
  font-weight: 600;
  padding-bottom: 8px;
  border-bottom: 2px solid #e0e5ec;
  box-shadow: 0 1px 0 rgba(163,177,198,0.7);
}

.section-content {
  font-size: 1rem;
  color: #4a5568;
  line-height: 1.6;
  margin-bottom: 16px;
}

.feature-list {
  list-style: none;
  padding: 0;
}

.feature-list li {
  padding: 12px 16px;
  margin-bottom: 8px;
  background-color: #e0e5ec;
  border-radius: 12px;
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8);
  color: #4a5568;
  font-size: 1rem;
  transition: all 0.3s ease;
}

.feature-list li:hover {
  box-shadow: 2px 2px 4px 0 rgba(163,177,198, 0.7), -2px -2px 4px 0 rgba(255,255,255, 0.8);
  transform: translateX(4px);
}

.contact-info {
  margin-top: 16px;
}

.contact-item {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  margin-bottom: 12px;
  background-color: #e0e5ec;
  border-radius: 12px;
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.contact-icon {
  font-size: 1.5rem;
  margin-right: 16px;
}

.contact-text {
  font-size: 1rem;
  color: #4a5568;
}

.about-footer {
  margin-top: 32px;
  padding-top: 16px;
  border-top: 2px solid #e0e5ec;
  box-shadow: 0 1px 0 rgba(163,177,198,0.7);
  text-align: center;
}

.about-footer p {
  font-size: 0.9rem;
  color: #718096;
  margin: 0;
}
</style>
