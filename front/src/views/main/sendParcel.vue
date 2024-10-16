<template>
    <div class="modal" role="dialog" id="my_modal_8">
        <div class="fixed inset-0 bg-gray-100 bg-opacity-75 flex justify-center items-center">
            <div class="bg-white rounded-lg p-6 w-[600px]">
                <div class="flex justify-between items-center mb-4">
                    <h2 class="text-lg font-semibold">智能填写</h2><button @click="alert" aria-label="Close"><svg
                            xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                            class="w-6 h-6">
                            <rect width="18" height="18" x="3" y="3" rx="2" ry="2"></rect>
                            <line x1="3" x2="21" y1="9" y2="9"></line>
                            <path d="m9 16 3-3 3 3"></path>
                        </svg></button>
                </div>
                <div class="mb-4"><textarea class="w-full h-[120px] p-2 border border-gray-300 rounded resize-none"
                        placeholder="请加信息，自动拼接为结果，电话和地址示例：顾小主，136********，广东省深圳市南山区XX路XX号1栋B座506" v-model="text"></textarea>
                </div>
                <div class="text-sm text-gray-500 mb-6">
                    识别后，选择适合拼接为地址信息最准确，如有遗漏请及时补充
                </div>
                <div class="flex justify-end"><button
                        class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2 w-[100px]"
                        @click="extractInformation">识别</button>
                </div>
            </div>
        </div>
    </div>
    <div class="app">
        <div class="grid h-screen min-h-screen w-full overflow-hidden lg:grid-cols-[280px_1fr]">
            <div class="hidden border-r bg-blue-100/40 lg:block dark:bg-gray-800/40">
                <div class="flex h-full max-h-screen flex-col gap-2">
                    <div class="flex h-[60px] items-center border-b px-6"><a class="flex items-center gap-2 font-semibold"
                            href="#"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round" class="h-6 w-6 text-blue-600">
                                <path
                                    d="M11.767 19.089c4.924.868 6.14-6.025 1.216-6.894m-1.216 6.894L5.86 18.047m5.908 1.042-.347 1.97m1.563-8.864c4.924.869 6.14-6.025 1.215-6.893m-1.215 6.893-3.94-.694m5.155-6.2L8.29 4.26m5.908 1.042.348-1.97M7.48 20.364l3.126-17.727">
                                </path>
                            </svg><span>
                                ChainLogix
                            </span></a><button
                            class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground ml-auto h-8 w-8"><svg
                                xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                                class="h-4 w-4">
                                <path d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9"></path>
                                <path d="M10.3 21a1.94 1.94 0 0 0 3.4 0"></path>
                            </svg><span class="sr-only">
                                Toggle notifications
                            </span></button></div>
                    <div class="flex-1 overflow-auto py-2">
                        <nav class="grid items-start px-4 text-sm font-medium"><router-link to="/index"><a
                                    class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-blue-600 dark:text-gray-400 dark:hover:text-gray-50"><svg
                                        xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round" class="h-4 w-4">
                                        <rect width="7" height="9" x="3" y="3" rx="1"></rect>
                                        <rect width="7" height="5" x="14" y="3" rx="1"></rect>
                                        <rect width="7" height="9" x="14" y="12" rx="1"></rect>
                                        <rect width="7" height="5" x="3" y="16" rx="1"></rect>
                                    </svg>
                                    主页
                                </a></router-link><router-link to="/traceShipment"><a
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
                                </a></router-link> <router-link to="/sendParcel"><a
                                    class="flex items-center gap-3 rounded-lg px-3 py-2 text-blue-600 dark:text-gray-50"><svg
                                        xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round" class="h-4 w-4">
                                        <path d="m22 2-7 20-4-9-9-4Z"></path>
                                        <path d="M22 2 11 13"></path>
                                    </svg>
                                    发送包裹
                                </a></router-link><router-link to="/transaction"><a
                                    class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-blue-600 dark:text-gray-400 dark:hover:text-gray-50"><svg
                                        xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round" class="h-4 w-4">
                                        <path d="M21 12V7H5a2 2 0 0 1 0-4h14v4"></path>
                                        <path d="M3 5v14a2 2 0 0 0 2 2h16v-5"></path>
                                        <path d="M18 12a2 2 0 0 0 0 4h4v-4Z"></path>
                                    </svg>
                                    交易
                                </a></router-link><router-link to="/userMannage" v-show="identity == 2">
                                <a
                                    class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-blue-600 dark:text-gray-400 dark:hover:text-gray-50"><svg
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
                                    class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-blue-600 dark:text-gray-400 dark:hover:text-gray-50"><svg
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
                        class="lg:hidden" href="#"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                            viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                            stroke-linejoin="round" class="h-6 w-6">
                            <line x1="4" x2="20" y1="12" y2="12"></line>
                            <line x1="4" x2="20" y1="6" y2="6"></line>
                            <line x1="4" x2="20" y1="18" y2="18"></line>
                        </svg><span class="sr-only">
                            Toggle menu
                        </span></a>
                    <div class="ml-auto flex gap-2">
                        <div class="dropdown dropdown-end " style="margin-right:0px">
                            <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
                                <div class="w-10 rounded-full">
                                    <img alt="Tailwind CSS Navbar component" :src="pic" />
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
                <main>
                    <div class="bg-gray-100 p-8">
                        <div class="flex justify-between mb-6">
                            <h1 class="text-xl font-bold">常规下单</h1><button
                                class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-primary/90 h-10 px-4 py-2 bg-red-500 text-white">添加订单</button>
                        </div>
                        <div class="grid grid-cols-2 gap-4">
                            <div class="rounded-lg border text-card-foreground shadow-sm w-full bg-white p-4"
                                data-v0-t="card">
                                <div class="flex flex-col space-y-1.5 p-6">
                                    <div class="flex justify-between">
                                        <h2 class="text-lg font-bold">基础信息</h2>
                                        <div><button
                                                class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2"
                                                @click="openVerified(1)">智能识别</button>
                                        </div>
                                    </div>
                                </div>
                                <div class="p-6">
                                    <form>
                                        <div class="grid grid-cols-2 gap-4">
                                            <div class="flex flex-col space-y-1.5"><label for="senderName"
                                                    class="font-bold">*寄件人</label><input
                                                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                                    id="senderName" placeholder="请填写寄件人姓名" v-model="from_user"></div>
                                            <div class="flex flex-col space-y-1.5"><label for="senderPhone"
                                                    class="font-bold">*联系电话</label><input
                                                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                                    id="senderPhone" placeholder="请填写手机号码或固定电话" v-model="from_phone"></div>
                                            <div class="flex flex-col space-y-1.5"><label for="senderAddress"
                                                    class="font-bold">*详细地址</label><input
                                                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                                    id="senderAddress" placeholder="请填写详细地址（例如xx路xx号）"
                                                    v-model="from_address"></div>
                                            <div class="flex flex-col space-y-1.5"><label for="senderRemark"
                                                    class="font-bold">公司名称</label><input
                                                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                                    id="senderRemark" placeholder="请填写公司名称" v-model="from_firm"></div>
                                        </div>
                                    </form>
                                </div>
                            </div>
                            <div class="rounded-lg border text-card-foreground shadow-sm w-full bg-white p-4"
                                data-v0-t="card">
                                <div class="flex flex-col space-y-1.5 p-6">
                                    <div class="flex justify-between">
                                        <h2 class="text-lg font-bold text-red-500">收货信息</h2>
                                        <div><button
                                                class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 text-red-500"
                                                @click="openVerified(2)">智能识别</button>
                                        </div>
                                    </div>
                                </div>
                                <div class="p-6">
                                    <form>
                                        <div class="grid grid-cols-2 gap-4">
                                            <div class="flex flex-col space-y-1.5"><label for="recipientName"
                                                    class="font-bold text-red-500">*收货人</label><input
                                                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                                    id="recipientName" placeholder="请填写收货人姓名" v-model="to_user"></div>
                                            <div class="flex flex-col space-y-1.5"><label for="recipientPhone"
                                                    class="font-bold text-red-500">*联系电话</label><input
                                                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                                    id="recipientPhone" placeholder="请填写手机号码或固定电话" v-model="to_phone"></div>
                                            <div class="flex flex-col space-y-1.5"><label for="recipientAddress"
                                                    class="font-bold text-red-500">*详细地址</label><input
                                                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                                    id="recipientAddress" placeholder="请填写详细地址（例如xx路xx号）"
                                                    v-model="to_address"></div>
                                            <div class="flex flex-col space-y-1.5"><label for="recipientRemark"
                                                    class="font-bold text-red-500">公司名称</label><input
                                                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                                    id="recipientRemark" placeholder="请填写公司名称" v-model="to_firm"></div>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </div>
                        <div class="my-6">
                            <div class="rounded-lg border text-card-foreground shadow-sm w-full bg-white p-4"
                                data-v0-t="card">
                                <div class="flex flex-col space-y-1.5 p-6">
                                    <h2 class="text-lg font-bold">付款方式</h2>
                                </div>
                                <div class="p-6">
                                    <input className="mr-2" id="payOnline" name="paymentMethod" type="radio" checked />
                                    <label className="mr-4" htmlFor="payOnline">
                                        支付宝
                                    </label>
                                </div>
                            </div>
                        </div>
                        <div class="flex justify-end"><button
                                class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-primary/90 h-10 px-4 py-2 bg-red-500 text-white"
                                @click="createPackage">提交订单</button>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'sendParcel',
    data() {
        return {
            identity:0,
            id: 0,
            from_user: '',
            from_phone: '',
            from_address: '',
            from_firm: '',
            to_user: '',
            to_phone: '',
            to_address: '',
            to_firm: '',
            text: '',
            pic: ''
        }
    },
    created() {
        // 从缓存中获取用户头像 如果有则使用 没有则用默认
        this.pic = localStorage.getItem('pic');
        this.identity = localStorage.getItem('identity');
    },
    methods: {
        // 创建包裹
        createPackage() {
            if (this.from_user.length == 0 || this.from_phone.length == 0 || this.from_address.length == 0 || this.to_user.length == 0 || this.to_phone.length == 0 || this.to_address.length == 0) {
                this.$message.error('请填写完整信息');
                return;
            }
            // 创建包裹请求
            this.$axios.post("/main/createPackage", {
                email: localStorage.getItem('email'),
                send_person: this.from_user,
                send_phone_number: this.from_phone,
                send_detailed_address: this.from_address,
                send_company_name: this.from_firm,
                receive_person: this.to_user,
                receive_phone_number: this.to_phone,
                receive_address: this.to_address,
                receive_company_name: this.to_firm,
            }).then(res => {
                if (res.data.code == 200) {
                    this.$message.success('创建成功');
                    this.from_user = ''
                    this.from_phone = ''
                    this.from_address = ''
                    this.from_firm = ''
                    this.to_user = ''
                    this.to_phone = ''
                    this.to_address = ''
                    this.to_firm = ''
                    // 然后打开这个页面
                    window.location.href = res.data.url;
                } else {
                    this.$message.error(res.data.msg);
                }
            }).catch(err => {
                this.$message.error('网络错误');
            })
        },
        alert() {
            window.location.href = '#';
        },
        openVerified(id) {
            this.id = id
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
</style>