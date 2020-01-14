<template>
  <div class="container flex-column my-3">
    <block>这个页面在服务器上应该由后端渲染。</block>
    <block>
      <router-link :to="{name: 'dashboard'}">
        <el-button size="small" type="primary">跳转到控制面板</el-button>
      </router-link>
    </block>
  </div>
</template>
