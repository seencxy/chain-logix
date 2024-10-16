<template>
  <div class="modal" role="dialog" id="chage_head">
    <div class="fixed inset-0 bg-gray-100 bg-opacity-75 flex justify-center items-center">
      <div class="bg-white rounded-lg p-6 w-[600px]">
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-lg font-semibold">修改头像</h2><button aria-label="Close"><svg @click="alert"
              xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
              stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-6 h-6">
              <rect width="18" height="18" x="3" y="3" rx="2" ry="2"></rect>
              <line x1="3" x2="21" y1="9" y2="9"></line>
              <path d="m9 16 3-3 3 3"></path>
            </svg></button>
        </div>
        <div class="mb-4">
          <div class="flex justify-center items-center"><img @click="triggerFileInput" :src="imageBase64 || headImage"
              alt="Profile Picture" class="w-32 h-32 rounded-full border border-gray-300" width="120" height="120"
              style="aspect-ratio: 120 / 120; object-fit: cover;"></div>
          <input type="file" ref="fileInput" @change="onFileChange" hidden />
          <div class="text-center text-sm text-gray-500 mt-1">点击图片以更改头像</div>
        </div>
        <div class="text-sm text-gray-500 mb-6">
          上传新的图片作为您的头像，支持jpg, png, gif格式，大小不超过5MB
        </div>
        <div class="flex justify-end"><button
            class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2 w-[100px]"
            @click="changeUserHead">上传</button>
        </div>
      </div>
    </div>
  </div>

  <router-view v-slot="{ Component }">
    <transition name="fade">
      <component :is="Component" />
    </transition>
  </router-view>
</template>

<script>
import WebSocketService from "./js/WebSocketService"
import headImage from '@/assets/head.svg';
export default {
  data() {
    return {
      imageBase64: '', // 存储图片的 Base64 编码
      headImage: headImage
    }
  },
  mounted() {
    document.addEventListener('websocket-message', this.handleWebsocketMessage);
    WebSocketService.establishConnection(localStorage.getItem("email"));
  },
  methods: {
    handleWebsocketMessage(event) {
      this.$message.success(event.detail);
    },
    alert() {
      window.location.href = '#';
    },
    triggerFileInput() {
      this.$refs.fileInput.click(); // 触发文件输入
    },
    onFileChange(e) {
      const file = e.target.files[0];
      if (file && file.type.match('image.*')) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.imageBase64 = e.target.result; // 将结果存储在 imageBase64 中
        };
        reader.readAsDataURL(file);
      }
    },
    changeUserHead() {
      if (this.imageBase64.length == 0) {
        this.$message.error("更换的头像不能为空")
      } else {
        // 修改用户头像
        this.$axios.post("/user/changeUserHead", {
          'email': localStorage.getItem('email'),
          'head': this.imageBase64
        }).then((rep) => {
          if (rep.data.code == 200) {
            this.$message.success(rep.data.message)
            // 更换缓存头像
            localStorage.setItem('pic', this.imageBase64)
            window.location.href = '#';
            // 刷新页面
            window.location.reload()
            return
          }
          this.$message.error(rep.data.message)
        }).catch((rep) => {
          console.log(rep);
        })
      }
    }
  }
}
</script>

<style scoped>
/* 渐变设置 */
.fade-enter-from,
.fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}

.fade-enter-to,
.fade-leave-from {
  opacity: 1;
}

.fade-enter-active {
  transition: all 0.7s ease;
}

.fade-leave-active {
  transition: all 0.3s cubic-bezier(1, 0.6, 0.6, 1);
}

.radio-group {
  display: flex;
  /* 设置为 flex 布局 */
  flex-wrap: nowrap;
  /* 防止子元素换行 */
}

label {
  display: flex;
  align-items: center;
  margin-right: 10px;
  white-space: nowrap;
}

.radio {
  margin-right: 4px;
}
</style>