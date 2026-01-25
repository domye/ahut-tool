<template>
  <div class="about-container">
    <div class="about-card">
      <h1 class="about-title">关于安徽工业大学生活工具</h1>

      <div class="about-section">
        <h2 class="section-title">项目介绍</h2>
        <p class="section-content">
          安徽工业大学生活工具是一款专为安徽工业大学师生打造的智慧校园服务平台，致力于用科技赋能校园生活。
          我们深入理解师生日常需求，将分散的校园服务整合于一体，打造全方位、个性化的数字生活助手。
          通过简洁直观的界面设计与流畅的交互体验，让每一位师生都能轻松享受便捷高效的校园服务。
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
          本项目采用Wails框架构建，将Web技术栈与Go语言后端完美结合，打造高性能的桌面应用程序。
          前端基于Vue 3框架，运用Composition API进行组件化开发，配合TypeScript提升代码质量与可维护性。
          采用Ant Design Vue组件库，提供美观一致的用户界面，并融入新拟态(Neumorphism)设计风格，带来流畅自然的交互体验。
          后端使用Go语言开发，利用其高效的并发处理能力和简洁的语法，确保系统运行稳定可靠。
          通过Wails的绑定机制，实现前后端无缝通信，数据交互快速高效，为师生提供便捷的桌面端校园服务体验。
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
            <span class="contact-text">1523610551@qq.com</span>
          </div>
<!--          <div class="contact-item">-->
<!--            <span class="contact-icon">📱</span>-->
<!--            <span class="contact-text">微信公众号：安徽工业大学</span>-->
<!--          </div>-->
        </div>
      </div>

      <div class="about-footer">
        <p>© 2026 Domye 版权所有</p>
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
