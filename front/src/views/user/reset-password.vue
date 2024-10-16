<template>
  <div class="app">
    <main class="flex items-center justify-center h-screen">
      <div class="title text-blue-700" style="margin-left: 200px">
        <div class="box-header">
          <p style="font-weight: bolder;font-size: 30px">ChainLogix</p>
          <img class="logo" src="@/assets/logo.svg" alt=""/>
        </div>
      </div>
      <div class="border text-card-foreground w-full max-w-md mx-auto bg-white shadow-xl rounded-lg overflow-hidden"
           data-v0-t="card">
        <div class="p-6">
          <h1 class="text-3xl font-bold text-center text-blue-700">区块链物流平台</h1>
          <form class="space-y-6 mt-4">
            <div class="space-y-2"><label
                class="font-medium peer-disabled:cursor-not-allowed peer-disabled:opacity-70 text-lg"
                for="username">邮箱</label><input
                class="flex h-10 w-full rounded-md border bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 border-blue-300 focus:ring-2 focus:ring-blue-600 focus:border-transparent"
                id="username" required="" placeholder="邮箱" v-model="email"></div>
            <div class="space-y-2"><label
                class="font-medium peer-disabled:cursor-not-allowed peer-disabled:opacity-70 text-lg"
                for="password">密码</label><input type="password"
                                                  class="flex h-10 w-full rounded-md border bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 border-blue-300 focus:ring-2 focus:ring-blue-600 focus:border-transparent"
                                                  id="password" required="" placeholder="密码" v-model="password"></div>
            <div class="space-y-2"><label
                class="font-medium peer-disabled:cursor-not-allowed peer-disabled:opacity-70 text-lg"
                for="password">再次输入密码</label><input type="password"
                                                          class="flex h-10 w-full rounded-md border bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 border-blue-300 focus:ring-2 focus:ring-blue-600 focus:border-transparent"
                                                          id="password" required="" placeholder="再次输入密码"
                                                          v-model="againPass"></div>
            <div class="space-y-2"><label
                class="font-medium peer-disabled:cursor-not-allowed peer-disabled:opacity-70 text-lg"
                for="verification">验证码</label><input
                class="flex h-10 w-full rounded-md border bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 border-blue-300 focus:ring-2 focus:ring-blue-600 focus:border-transparent"
                id="verification" required="" placeholder="验证码" v-model="code">
              <button
                  class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 px-4 py-2 w-full bg-blue-600 hover:bg-blue-700 text-white"
                  type="button" @click="sendCode">发送验证码
              </button>
            </div>
            <button
                class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 px-4 py-2 w-full bg-blue-600 hover:bg-blue-700 text-white"
                type="button" @click="ResetPass">设置密码
            </button>
          </form>
        </div>
        <div class="p-4 bg-gray-200 text-right">
          <a style="font-weight: bolder;cursor:pointer" class="text-blue-700" @click="gotoLogin">登录</a>
          <a>&nbsp&nbsp&nbsp&nbsp&nbsp</a>
          <a style="font-weight: bolder;cursor:pointer" class="text-blue-700" @click="gotoRegister">
            注册
          </a>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
export default {
  name: "reset-password",
  data() {
    return {
      email: '',
      password: '',
      againPass: '',
      code: '',
    }
  },
  methods: {
    gotoRegister() {
      this.$router.push('/register');
    },
    gotoLogin() {
      this.$router.push('/login');
    },
    validateEmail() { // 检查邮箱是否合法
      const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(this.email.toLowerCase());
    },
    sendCode() {
      if (this.validateEmail(this.email)) {
        this.$axios.post('/user/sendCode', {
          email: this.email
        }).then((rep) => {
          if (rep.data.code == 200) {
            this.$message.success(rep.data.message);
          } else {
            this.$message.error(rep.data.message);
          }
        }).catch((rep) => {
          console.log(rep)
        })
      } else {
        this.$message.error('请输入正确的邮箱');
      }
    },
    ResetPass() {
      if (this.validateEmail(this.email) && this.password.length >= 6 && this.password == this.againPass && this.code.length == 6) {
        this.$axios.post('/user/resetPassword', {
          email: this.email,
          password: this.password,
          code: this.code
        }).then((rep) => {
          if (rep.data.code == 200) {
            this.$message.success(rep.data.message);
            // 设置1s等待
            setTimeout(() => {
              this.$router.push('/login');
            }, 1000);
          } else {
            this.$message.error(rep.data.message);
          }
        }).catch((rep) => {
          console.log(rep)
        })
      } else if (!this.validateEmail(this.email)) {
        this.$message.error('请输入合法邮箱');
      } else if (this.password.length < 6) {
        this.$message.error('密码长度不能小于6位');
      } else if (this.password != this.againPass) {
        this.$message.error('两次密码输入不一致');
      } else if (this.code.length != 6) {
        this.$message.error('请输入正确的验证码');
      }
    }

  }
}
</script>

<style scoped>
.app {
  background-image: url('@/assets/350.png'); /* 使用相对路径指向你的图片文件 */
  background-size: cover; /* 覆盖整个容器 */
  background-repeat: no-repeat; /* 不重复背景图片 */
  background-position: center center; /* 居中背景图片 */
}
</style>