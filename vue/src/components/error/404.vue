<template>
  <div class="not-found d-flex align-items-center bg-light">
    <div class="d-flex justify-content-center align-items-center w-100">
      <div class="not-found-left">
        <img class="w-100" :src="Image404" alt="404">
        <img class="not-found-left__cloud left" :src="Image404Cloud" alt="404">
        <img class="not-found-left__cloud mid" :src="Image404Cloud" alt="404">
        <img class="not-found-left__cloud right" :src="Image404Cloud" alt="404">
      </div>
      <div class="not-found-right py-5">
        <div class="not-found-right__oops mb-4 text-primary font-weight-bold">
          如影而来, 如影而去...
        </div>
        <div class="not-found-right__headline mb-3 text-primary">
          您所访问的页面不存在
        </div>
        <div class="not-found-right__info mb-4 text-secondary">
          请检查您输入的网址是否正确，请点击以下按钮返回主页
        </div>
        <router-link class="not-found-right__return-home py-2 bg-primary text-white text-center"
                     :to="{ name: 'dashboard' }">
          返回首页
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import Image404 from '@/assets/images_404/404.png'
import Image404Cloud from '@/assets/images_404/404_cloud.png'

export default {
  data() {
    return {
      Image404,
      Image404Cloud
    }
  }
}
</script>

<style lang="scss" scoped>
.not-found {
  min-height: 100vh;

  .not-found-left {
    position: relative;
    width: 600px;
    overflow: hidden;

    &__cloud {
      position: absolute;
      animation-duration: 2s;
      animation-timing-function: linear;
      animation-fill-mode: forwards;

      &.left {
        width: 80px;
        top: 17px;
        left: 220px;
        opacity: 0;
        animation-name: cloudLeft;
        animation-delay: 1s;
      }

      &.mid {
        width: 46px;
        top: 10px;
        left: 420px;
        opacity: 0;
        animation-name: cloudMid;
        animation-delay: 1.2s;
      }

      &.right {
        width: 62px;
        top: 100px;
        left: 500px;
        opacity: 0;
        animation-name: cloudRight;
        animation-delay: 1s;
      }

      @keyframes cloudLeft {
        0% {
          top: 17px;
          left: 220px;
          opacity: 0;
        }
        20% {
          top: 33px;
          left: 188px;
          opacity: 1;
        }
        80% {
          top: 81px;
          left: 92px;
          opacity: 1;
        }
        100% {
          top: 97px;
          left: 60px;
          opacity: 0;
        }
      }
      @keyframes cloudMid {
        0% {
          top: 10px;
          left: 420px;
          opacity: 0;
        }
        20% {
          top: 40px;
          left: 360px;
          opacity: 1;
        }
        70% {
          top: 130px;
          left: 180px;
          opacity: 1;
        }
        100% {
          top: 160px;
          left: 120px;
          opacity: 0;
        }
      }
      @keyframes cloudRight {
        0% {
          top: 100px;
          left: 500px;
          opacity: 0;
        }
        20% {
          top: 120px;
          left: 460px;
          opacity: 1;
        }
        80% {
          top: 180px;
          left: 340px;
          opacity: 1;
        }
        100% {
          top: 200px;
          left: 300px;
          opacity: 0;
        }
      }
    }
  }

  .not-found-right {
    position: relative;
    width: 300px;

    &__oops {
      font-size: 32px;
      opacity: 0;
      animation-name: slideUp;
      animation-duration: 0.5s;
      animation-fill-mode: forwards;
    }

    &__headline {
      font-size: 20px;
      opacity: 0;
      animation-name: slideUp;
      animation-duration: 0.5s;
      animation-delay: 0.1s;
      animation-fill-mode: forwards;
    }

    &__info {
      font-size: 13px;
      opacity: 0;
      animation-name: slideUp;
      animation-duration: 0.5s;
      animation-delay: 0.2s;
      animation-fill-mode: forwards;
    }

    &__return-home {
      display: block;
      float: left;
      width: 110px;
      border-radius: 100px;
      opacity: 0;
      font-size: 14px;
      text-decoration: none;
      cursor: pointer;
      animation-name: slideUp;
      animation-duration: 0.5s;
      animation-delay: 0.3s;
      animation-fill-mode: forwards;
    }

    @keyframes slideUp {
      0% {
        transform: translateY(60px);
        opacity: 0;
      }
      100% {
        transform: translateY(0);
        opacity: 1;
      }
    }
  }
}
</style>
