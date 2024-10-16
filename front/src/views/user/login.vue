<template>
  <div class="app">
    <main class="flex items-center justify-center  h-screen">
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
                id="username" required="" type="email" placeholder="邮箱" v-model="email"></div>
            <div class="space-y-2"><label
                class="font-medium peer-disabled:cursor-not-allowed peer-disabled:opacity-70 text-lg"
                for="password">密码</label><input
                class="flex h-10 w-full rounded-md border bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 border-blue-300 focus:ring-2 focus:ring-blue-600 focus:border-transparent"
                id="password" required="" type="password" placeholder="密码" v-model="password"></div>
            <div class="space-y-2"><label
                class="font-medium peer-disabled:cursor-not-allowed peer-disabled:opacity-70 text-lg"
                for="role">选择身份</label>
              <select class="select select-info w-full max-w-xs" style="width: 50%;margin-left:10%"
                      v-model="role">
                <option disabled selected>选择身份</option>
                <option>用户</option>
                <option>供应链管理员</option>
                <option>总管理员</option>
              </select>
            </div>
            <button
                class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 h-10 px-4 py-2 w-full bg-blue-600 hover:bg-blue-700 text-white"
                type="button" @click="login">登录
            </button>
          </form>
        </div>
        <div class="p-4 bg-gray-200 text-right">
          <a style="font-weight: bolder;cursor:pointer" class="text-blue-700" @click="gotoRegister">注册</a>
          <a>&nbsp&nbsp&nbsp&nbsp&nbsp</a>
          <a style="font-weight: bolder;cursor:pointer" class="text-blue-700" @click="gotoSetPassword">
            忘记密码
          </a>
        </div>
      </div>

    </main>
  </div>
</template>

<script>
import WebSocketService from "../../js/WebSocketService"

export default {
  name: 'login',
  data() {
    return {
      email: '',
      password: '',
      role: '选择身份',
      reconnectAttempts: 0,
      maxReconnectAttempts: 10,
    }
  },
  methods: {
    gotoRegister() {
      this.$router.push('/register');
    },
    gotoSetPassword() {
      this.$router.push('/set-password');
    },
    validateEmail() { // 检查邮箱是否合法
      const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(this.email.toLowerCase());
    },

    login() {
      if (this.role != "选择身份" && this.password.length >= 6 && this.validateEmail(this.email)) {
        let role;
        if (this.role == '用户') {
          role = 0
        } else if (this.role == '供应链管理员') {
          role = 1
        } else {
          role = 2
        }
        this.$axios.post('/user/login', {
          "email": this.email,
          "password": this.password,
          "identity": role,
        }).then((rep) => {
          if (rep.data.code == 200) {
            // 将token以及email存储起来
            localStorage.setItem("auth_token", rep.data.token)
            localStorage.setItem("email", rep.data.email)
            localStorage.setItem('identity', rep.data.identity)
            localStorage.setItem('pic', rep.data.pic)
            this.$message.success(rep.data.message);


            WebSocketService.establishConnection(this.email);

            // 设置1s等待
            setTimeout(() => {
              this.$router.push('/index');
            }, 1000);
          } else {
            this.$message.error(rep.data.message);
          }
        }).catch((rep) => {
          console.log(rep);
        })
      } else if (this.role == "选择身份") {
        this.$message.error('请选择身份');
      } else if (this.password.length < 6) {
        this.$message.error('密码长度不得小于6位');
      } else if (!this.validateEmail(this.email)) {
        this.$message.error('邮箱格式不正确');
      }
    }
  }
}
</script>

<style scoped>
* {
  -webkit-user-select: none;
  /* Safari */
  -moz-user-select: none;
  /* Firefox */
  -ms-user-select: none;
  /* IE/Edge */
  user-select: none;
  /* 标准语法 */
}

.app {
  background-image: url('@/assets/350.png'); /* 使用相对路径指向你的图片文件 */
  background-size: cover; /* 覆盖整个容器 */
  background-repeat: no-repeat; /* 不重复背景图片 */
  background-position: center center; /* 居中背景图片 */
}
</style>