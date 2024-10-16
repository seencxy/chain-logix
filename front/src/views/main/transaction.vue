<template>
  <div class="app overflow-auto">
    <div class="grid h-screen min-h-screen w-full overflow-hidden lg:grid-cols-[280px_1fr] overflow-auto">
      <div class="hidden border-r bg-blue-100/40 lg:block dark:bg-gray-800/40 overflow-auto">
        <div class="flex h-full max-h-screen flex-col gap-2">
          <div class="flex h-[60px] items-center border-b px-6"><a class="flex items-center gap-2 font-semibold"
                                                                   href="#">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                 fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                 stroke-linejoin="round" class="h-6 w-6 text-blue-600">
              <path
                  d="M11.767 19.089c4.924.868 6.14-6.025 1.216-6.894m-1.216 6.894L5.86 18.047m5.908 1.042-.347 1.97m1.563-8.864c4.924.869 6.14-6.025 1.215-6.893m-1.215 6.893-3.94-.694m5.155-6.2L8.29 4.26m5.908 1.042.348-1.97M7.48 20.364l3.126-17.727">
              </path>
            </svg>
            <span>
                                ChainLogix
                            </span></a>
            <button
                class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground ml-auto h-8 w-8">
              <svg
                  xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                  stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                  class="h-4 w-4">
                <path d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9"></path>
                <path d="M10.3 21a1.94 1.94 0 0 0 3.4 0"></path>
              </svg>
              <span class="sr-only">
                                Toggle notifications
                            </span></button>
          </div>
          <div class="flex-1 overflow-auto py-2">
            <nav class="grid items-start px-4 text-sm font-medium">
              <router-link to="/index"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-blue-600 dark:text-gray-400 dark:hover:text-gray-50">
                <svg
                    xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                    stroke-linejoin="round" class="h-4 w-4">
                  <rect width="7" height="9" x="3" y="3" rx="1"></rect>
                  <rect width="7" height="5" x="14" y="3" rx="1"></rect>
                  <rect width="7" height="9" x="14" y="12" rx="1"></rect>
                  <rect width="7" height="5" x="3" y="16" rx="1"></rect>
                </svg>
                主页
              </a></router-link>
              <router-link to="/traceShipment"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-blue-600 dark:text-gray-400 dark:hover:text-gray-50">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                     fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                     stroke-linejoin="round" class="h-4 w-4">
                  <path d="M2 17 17 2"></path>
                  <path d="m2 14 8 8"></path>
                  <path d="m5 11 8 8"></path>
                  <path d="m8 8 8 8"></path>
                  <path d="m11 5 8 8"></path>
                  <path d="m14 2 8 8"></path>
                  <path d="M7 22 22 7"></path>
                </svg>
                跟踪货物
              </a></router-link>
              <router-link to="/sendParcel"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-blue-600 dark:text-gray-400 dark:hover:text-gray-50">
                <svg
                    xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                    stroke-linejoin="round" class="h-4 w-4">
                  <path d="m22 2-7 20-4-9-9-4Z"></path>
                  <path d="M22 2 11 13"></path>
                </svg>
                发送包裹
              </a></router-link>
              <router-link to="/transaction"><a
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-blue-600 dark:text-gray-50">
                <svg
                    xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                    stroke-linejoin="round" class="h-4 w-4">
                  <path d="M21 12V7H5a2 2 0 0 1 0-4h14v4"></path>
                  <path d="M3 5v14a2 2 0 0 0 2 2h16v-5"></path>
                  <path d="M18 12a2 2 0 0 0 0 4h4v-4Z"></path>
                </svg>
                交易
              </a></router-link>
              <router-link to="/userMannage" v-show="identity == 2">
                <a
                    class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-blue-600 dark:text-gray-400 dark:hover:text-gray-50">
                  <svg
                      xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                      fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                      stroke-linejoin="round" class="h-4 w-4">
                    <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"></path>
                    <circle cx="12" cy="7" r="4"></circle>
                  </svg>
                  用户管理
                </a></router-link>
              <router-link to="/mannage" v-show="identity != 0">
                <a
                    class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-blue-600 dark:text-gray-400 dark:hover:text-gray-50">
                  <svg
                      xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                      fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                      stroke-linejoin="round" class="h-4 w-4">
                    <path d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4"></path>
                    <polyline points="10 17 15 12 10 7"></polyline>
                    <line x1="15" x2="3" y1="12" y2="12"></line>
                  </svg>
                  物流管理
                </a>
              </router-link>
            </nav>
          </div>
        </div>
      </div>
      <div class="flex flex-col overflow-auto">
        <header class="flex h-14 lg:h-[60px] items-center gap-4 border-b bg-blue-100/40 px-6 dark:bg-gray-800/40"><a
            class="lg:hidden" href="#">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
               viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
               stroke-linejoin="round" class="h-6 w-6">
            <line x1="4" x2="20" y1="12" y2="12"></line>
            <line x1="4" x2="20" y1="6" y2="6"></line>
            <line x1="4" x2="20" y1="18" y2="18"></line>
          </svg>
          <span class="sr-only">
                            Toggle menu
                        </span></a>
          <div class="ml-auto flex gap-2">
            <div class="dropdown dropdown-end " style="margin-right:0px">
              <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
                <div class="w-10 rounded-full">
                  <img alt="Tailwind CSS Navbar component" :src="pic"/>
                </div>
              </div>
              <ul tabindex="0"
                  class="mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52">
                <li @click="chageHead">
                  <a class="justify-between">
                    切换头像
                    <span class="badge">New</span>
                  </a>
                </li>
                <li @click="logout"><a>登出</a></li>
              </ul>
            </div>
          </div>
        </header>
        <main class="flex flex-col p-6 overflow-auto">
          <div class="border shadow-sm rounded-lg p-2">
            <div class="relative w-full overflow-auto">
              <table class="w-full caption-bottom text-sm">
                <thead class="[&amp;_tr]:border-b">
                <tr class="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted">
                  <th
                      class="h-12 px-4 text-center align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 min-w-[150px]">
                    包裹ID
                  </th>
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 min-w-[100px]">
                    发件人
                  </th>
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 min-w-[100px]">
                    收件人
                  </th>
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 min-w-[100px]">
                    发货地址
                  </th>
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 min-w-[100px]">
                    收货地址
                  </th>
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 hidden md:table-cell">
                    供应链管理员
                  </th>
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 hidden md:table-cell">
                    交易哈希
                  </th>
                  <th
                      class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 min-w-[100px]">
                    状态
                  </th>
                  <th
                      class="h-12 px-4 align-middle text-left font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-left">
                    操作
                  </th>
                </tr>
                </thead>
                <tbody>
                <tr class="border-b transition-colors hover:bg-muted/100 data-[state=selected]:bg-muted"
                    v-for="item in     blockData    " :key="item.hash">
                  <td
                      class="p-4 text-center align-middle [&amp;:has([role=checkbox])]:pr-0 md:table-cell">
                    <div class="tooltip" :data-tip="item.PackageId"
                         style="z-index: 9999;margin-left:0px">
                      {{ shortenUUID(item.PackageId) }}
                    </div>
                  </td>
                  <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 md:table-cell"> {{
                      item.Sender
                    }}
                  </td>
                  <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 hidden md:table-cell">
                    {{ item.Receiver }}
                  </td>
                  <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 hidden md:table-cell">
                    {{ item.SenderAddress }}
                  </td>
                  <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 hidden md:table-cell">
                    {{ item.ReceiverAddress }}
                  </td>
                  <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 hidden md:table-cell">
                    {{ item.SupplyChainAdmin }}
                  </td>

                  <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 hidden md:table-cell">
                    {{item.CurrentAddress}}
                  </td>

                  <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 md:table-cell">{{
                      item.Status
                    }}
                  </td>
                  <td class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 min-w-[100px]">

                    <th>
                      <button class="btn btn-error btn-xs" v-if="item.Status == 'await'"
                              @click="cancelPackage(item.PackageId)">取消物流
                      </button>
                      <button class="btn btn-ghost btn-xs"
                              v-if="item.Status != 'await' && item.Status != 'cancelled'"
                              @click="traceShipment(item.PackageId)"> 物流追踪
                      </button>

                    </th>
                  </td>
                </tr>
                </tbody>
              </table>
            </div>
          </div>
          <div class="join" style="margin-left: 40%; margin-right: 40%; margin-top:15px">
            <button class="join-item btn" @click="prevPage">«</button>
            <button class="join-item btn">Page {{ page }}</button>
            <button class="join-item btn" @click="nextPage">»</button>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'transaction',
  data() {
    return {
      identity: 0,
      page: 1,
      blockData: [],
    }
  },

  created() {
    // 从缓存中获取用户头像 如果有则使用 没有则用默认
    this.pic = localStorage.getItem('pic');
    this.identity = localStorage.getItem('identity');
    this.$axios.get("/main/getUserPackage", {
      params: {
        page: this.page,
        size: 10
      }
    }).then((rep) => {
      console.log(rep);
      if (rep.data.code == 200) {
        this.blockData = rep.data.data
        return
      }
      this.$message.error("获取包裹信息失败: " + rep.data.message)
    }).catch((rep) => {
      console.log(rep);
    })
  },
  methods: {
    // 物流追踪
    traceShipment(package_id) {
      console.log(package_id);
      this.$router.push('/traceShipment?id=' + package_id);
    },
    prevPage() {
      if (this.page > 1) {
        this.page -= 1;
        this.$axios.get("/main/getUserPackage", {
          params: {
            page: this.page,
            size: 10
          }
        }).then((rep) => {
          if (rep.data.code == 200) {
            this.blockData = rep.data.data
            return
          }
          console.log(rep);
        }).catch((rep) => {
          console.log(rep);
        })
      } else {
        this.$message.error('已经是第一页了')
        return
      }
    },
    nextPage() {
      // 这里您可能需要一个逻辑来检查是否已经是最后一页
      if (this.blockData.length == 10) {
        this.page += 1;
      } else {
        this.$message.error('已经是最后一页了')
        return
      }
      this.$axios.get("/main/getUserPackage", {
        params: {
          page: this.page,
          size: 10
        }
      }).then((rep) => {
        if (rep.data.code == 200) {
          this.blockData = rep.data.data
          return
        }
        console.log(rep);
      }).catch((rep) => {
        console.log(rep);
      })
    },
    // 接收物流
    cancelPackage(packageId) {
      this.$axios.post("/main/cancelPackage", {
        package_id: packageId
      }).then((rep) => {
        console.log(rep);
        if (rep.data.code == 200) {
          this.$message.success(rep.data.message)
          return
        }
        this.$message.error(rep.data.message)
      }).catch((rep) => {
        console.log(rep);
      })
    },
    shortenUUID(uuid) {
      return uuid.substring(0, 8); // 截取前8个字符
    },
    alert() {
      window.location.href = '#';
    },
    openVerified(id) {
      this.id = id
      window.location.href = '#my_modal_8';
    },
    formatDate(hexTimestamp) {
      const timestamp = parseInt(hexTimestamp, 16);
      const date = new Date(timestamp);
      return date.toLocaleString(); // 或者使用其他格式化方法
    },
    formatNumber(int) {
      return parseInt(int, 16);
    },
    formatHash(hash) {
      return `${hash.substring(0, 5)}*****${hash.substring(hash.length - 5)}`;
    },
    prevPage() {
      if (this.page > 1) {
        this.page -= 1;
        this.$axios.get("/main/getTransactionInfo", {
          params: {
            id: this.page
          }
        }).then((rep) => {
          if (rep.data.code == 200) {
            this.blockData = rep.data.data
          }
          console.log(rep);
        }).catch((rep) => {
          console.log(rep);
        })
      } else {
        this.$message.error('没有上一页了');
      }
    },
    nextPage() {
      // 这里您可能需要一个逻辑来检查是否已经是最后一页
      if (this.blockData.length == 10) {
        this.page += 1;
      } else {
        this.$message.error('没有下一页了');
        return
      }
      this.$axios.get("/main/getTransactionInfo", {
        params: {
          id: this.page
        }
      }).then((rep) => {
        if (rep.data.code == 200) {
          this.blockData = rep.data.data
        }
        console.log(rep);
      }).catch((rep) => {
        console.log(rep);
      })
    },
    logout() {
      // 删除缓存
      localStorage.removeItem('auth_token');
      localStorage.removeItem('email');
      localStorage.removeItem('identity');
      localStorage.removeItem('pic');
      this.$router.push('/login');
      this.$message.success('登出成功!');
      // 设置1s等待
      setTimeout(() => {
        this.$router.push('/login');
      }, 1000)
    },
    extractInformation() {
      console.log(this.text);
      let parts = this.text.split('，');
      if (parts.length < 3) {
        // 处理错误情况，例如显示错误消息或者记录错误
        this.$message.error('输入格式不正确，无法识别信息。请确保格式为：姓名，电话，地址');
        window.location.href = '#';
        this.text = '';
        return;
      }

      if (this.id == 1) {
        this.from_address = parts[2].trim()
        this.from_phone = parts[1].trim()
        this.from_user = parts[0].trim()
        window.location.href = '#';
        this.text = ''
      } else if (this.id == 2) {
        this.to_address = parts[2].trim()
        this.to_phone = parts[1].trim()
        this.to_user = parts[0].trim()
        window.location.href = '#';
        this.text = ''
      } else {
        this.$message.error('智能识别失败');
        window.location.href = '#';
        this.text = ''
      }
    },
    chageHead() {
      window.location.href = '#chage_head';
    }
  }
}
</script>

<style>
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
</style>