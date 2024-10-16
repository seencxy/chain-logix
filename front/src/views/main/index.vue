<template>
  <div class="modal" role="dialog" id="my_modal_8">
    <div class="modal-box">
      <h3 class="font-bold text-lg">团队联系方式</h3>
      <p class="py-4">Email: golandcxy@gmail.com</p>
      <div class="modal-action">
        <a href="#" class="btn">确定</a>
      </div>
    </div>
  </div>
  <div class="app">
    <div class="grid h-screen min-h-screen w-full overflow-hidden lg:grid-cols-[280px_1fr]">
      <div class="hidden border-r bg-blue-100/40 lg:block dark:bg-gray-800/40">
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
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-blue-600 dark:text-gray-50">
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
                  class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-blue-600 dark:text-gray-400 dark:hover:text-gray-50">
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
      <div class="flex flex-col">
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
        <main class="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-6">
          <div class="flex items-center gap-4">
            <h1 class="font-semibold text-lg md:text-xl">Welcome to ChainLogix</h1>
          </div>
          <div class="flex flex-col md:grid md:grid-cols-6 gap-6">
            <div class="md:col-span-4 lg:col-span-3 xl:col-span-4 flex flex-col gap-6">
              <div class="rounded-lg border bg-card text-card-foreground shadow-sm" data-v0-t="card">
                <div class="flex flex-col space-y-1.5 p-6">
                  <h3 class="text-2xl font-semibold leading-none tracking-tight">近期订单</h3>
                </div>
                <div class="p-6">
                  <div class="relative w-full overflow-auto">
                    <table class="w-full caption-bottom text-sm">
                      <thead class="[&amp;_tr]:border-b">
                      <tr
                          class="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted">
                        <th
                            class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0">
                          物流ID
                        </th>
                        <th
                            class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0">
                          最新哈希
                        </th>
                        <th
                            class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0">
                          发往地
                        </th>
                        <th
                            class="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0">
                          状态
                        </th>
                      </tr>
                      </thead>
                      <tbody class="[&amp;_tr:last-child]:border-0 overflow-auto">
                      <tr v-for="element in data" v-key="data.PackageId"
                          class="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted">
                        <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0">{{
                            element.PackageId
                          }}
                        </td>
                        <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0">{{
                            element.SupplyChainAdmin
                          }}
                        </td>
                        <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0">{{
                            element.ReceiverAddress
                          }}
                        </td>
                        <td class="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0">{{
                            element.Status
                          }}
                        </td>
                      </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
              <div class="rounded-lg border bg-card text-card-foreground shadow-sm" data-v0-t="card">
                <div class="flex flex-col space-y-1.5 p-6">
                  <h3 class="text-2xl font-semibold leading-none tracking-tight">支持</h3>
                </div>
                <div class="p-6">
                  <p>如果您需要帮助或有任何疑问，请联系我们的支持团队。</p>
                </div>
                <div class="flex items-center p-6">
                  <button @click="openVerified"
                          class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2">
                    联系支持人员
                  </button>
                </div>
              </div>
            </div>
            <div class="md:col-span-2 lg:col-span-3 xl:col-span-2 flex flex-col gap-6">
              <div class="rounded-lg border bg-card text-card-foreground shadow-sm" data-v0-t="card">
                <div class="flex flex-col space-y-1.5 p-6">
                  <h3 class="text-2xl font-semibold leading-none tracking-tight">查询交易</h3>
                </div>
                <div class="p-6">
                  <form>
                    <div class="relative">
                      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                           viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                           stroke-linecap="round" stroke-linejoin="round"
                           class="absolute left-2.5 top-2.5 h-4 w-4 text-gray-500 dark:text-gray-400">
                        <circle cx="11" cy="11" r="8"></circle>
                        <path d="m21 21-4.3-4.3"></path>
                      </svg>
                      <input
                          class="flex h-10 rounded-md border border-input px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 w-full bg-white shadow-none appearance-none pl-8 md:w-2/3 lg:w-1/3 dark:bg-gray-950"
                          placeholder="输入交易哈希" type="search" style="width:90%" v-model="search"/>
                    </div>
                  </form>
                </div>
                <div class="flex items-center p-6">
                  <button @click="toSearch"
                          class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2">
                    查询
                  </button>
                </div>
              </div>
              <div class="rounded-lg border bg-card text-card-foreground shadow-sm" data-v0-t="card">
                <div class="flex flex-col space-y-1.5 p-6">
                  <h3 class="text-2xl font-semibold leading-none tracking-tight">分析</h3>
                  <h4>各省物流情况</h4>
                  <div class="w-full lg:w-1/2">
                    <div class="w-full aspect-[20/9]" style="width: 200px;">
                      <div class="flex flex-col space-y-1.5 p-6"
                           style="padding-left:0%;margin-bottom:-100%;margin-top:30%;width:160%">
                        <BarChart :chartData="baseData" :chartOptions="barChartOptions"/>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="p-6">
                  <div class="w-full aspect-[2/1]">
                    <div style="width: 100%; height: 100%;"></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script>
import BarChart from '../../components/chart.vue';
import WebSocketService from "../../js/WebSocketService"

export default {
  name: 'index',
  data() {
    return {
      identity: 0,
      data: [],
      search: '',
      baseData: {
        labels: [
          "北京市", "天津市", "河北省", "山西省", "内蒙古自治区",
          "辽宁省", "吉林省", "黑龙江省", "上海市", "江苏省",
          "浙江省", "安徽省", "福建省", "江西省", "山东省",
          "河南省", "湖北省", "湖南省", "广东省", "广西壮族自治区",
          "海南省", "重庆市", "四川省", "贵州省", "云南省",
          "西藏自治区", "陕西省", "甘肃省", "青海省", "宁夏回族自治区",
          "新疆维吾尔自治区", "台湾省", "香港特别行政区", "澳门特别行政区"
        ],
        datasets: [{
          label: 'Number',
          data: [
            820, 446, 523, 221, 687, 551, 693, 914, 392, 825,
            482, 597, 544, 515, 625, 846, 475, 596, 111, 269,
            889, 570, 293, 545, 791, 525, 233, 804, 117, 994,
            257, 238, 506, 965
          ],
          backgroundColor: 'rgba(54, 162, 235, 0.2)',
          borderColor: 'rgba(54, 162, 235, 1)',
          borderWidth: 1
        }]
      },

      barChartOptions: {
        scales: {
          x: {
            ticks: {
              // 设置X轴标签倾斜
              maxRotation: 45, // 最大旋转角度
              minRotation: 45  // 最小旋转角度
            }
          },
          y: {
            beginAtZero: true
          }
        }
      },
    }
  },
  created() {
    // 从缓存中获取用户头像 如果有则使用 没有则用默认
    this.pic = localStorage.getItem('pic');
    this.identity = localStorage.getItem('identity');
    WebSocketService.establishConnection(this.email);
    this.$axios.get('/main/getIndexInfos').then((rep) => {
      if (rep.data.code == 200) {
        console.log(rep)
        this.data = rep.data.data
        return
      }
      this.$message.error(rep.data.msg);
    }).catch((rep) => {
      console.log(rep)
    })
  },
  methods: {
    toSearch() {
      if (this.search.length == 0) {
        this.$message.error('请输入交易哈希');
        return;
      }
      let url = "http://192.168.40.190:5100/#/transaction?pageSize=10&pageNumber=1&blockNumber=&transHash=" + this.search;
      window.open(url, '_blank'); // 在新标签页中打开
    },
    handleWebsocketMessage(event) {
      this.$message.success(event.detail);
    },
    alert() {
      window.location.href = '#';
    },
    openVerified() {
      window.location.href = '#my_modal_8';
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
  },
  components: {
    BarChart,
  },
}
</script>

<style>

.flex-col-bottom {
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  /* 推送子元素到底部 */
  height: 100%;
  /* 确保容器有足够的高度 */
}
</style>