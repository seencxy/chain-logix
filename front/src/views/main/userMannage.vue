<template>
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
                                    class="flex items-center gap-3 rounded-lg px-3 py-2 text-gray-500 transition-all hover:text-blue-600 dark:text-gray-400 dark:hover:text-gray-50"><svg
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
                                <a class="flex items-center gap-3 rounded-lg px-3 py-2 text-blue-600 dark:text-gray-50"><svg
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
                <main class="flex flex-col p-6">
                    <h1 class="text-2xl font-semibold mb-4" style="margin-bottom: 20px;">用户管理</h1>
                    <div class="border shadow-sm rounded-lg p-2 overflow-x-auto">
                        <table class="table">
                            <!-- head -->
                            <thead>
                                <tr>
                                    <th>
                                        <label>
                                        </label>
                                    </th>
                                    <th>用户名</th>
                                    <th>身份</th>
                                    <th>状态</th>
                                    <th></th>
                                </tr>
                            </thead>
                            <tbody>
                                <!-- row 1 -->
                                <tr v-for="item in userData" :key="item.email">
                                    <th>
                                        <label>
                                        </label>
                                    </th>
                                    <td>
                                        <div class=" flex items-center gap-3">
                                            <div class="avatar">
                                                <div class="mask mask-squircle w-12 h-12">
                                                    <img :src="item.head" alt="Avatar Tailwind CSS Component" />
                                                </div>
                                            </div>
                                            <div>
                                                <div class="font-bold">{{ item.email }}</div>
                                            </div>
                                        </div>
                                    </td>
                                    <td>{{ getIdentity(item.identity) }}</td>
                                    <td>
                                        <div>{{ getStatus(item.status) }}</div>
                                    </td>

                                    <th>
                                        <button class="btn btn-error btn-xs" v-if="item.status != 3"
                                            @click="updateUserStatus(item.email, 3)">停用账户</button>
                                        <button class="btn btn-ghost btn-xs" v-if="item.status == 0"
                                            @click="updateUserStatus(item.email, 1)">审批</button>
                                        <button class="btn btn-ghost btn-xs" v-if="item.status == 3"
                                            @click="updateUserStatus(item.email, 1)">恢复账号</button>
                                    </th>
                                </tr>
                            </tbody>
                        </table>

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
    name: 'userMannge',
    data() {
        return {
            identity: 0,
            page: 1,
            userData: []
        }
    },
    created() {
        // 从缓存中获取用户头像 如果有则使用 没有则用默认
        this.pic = localStorage.getItem('pic');
        this.identity = localStorage.getItem('identity');
        this.$axios.get("/admin/getAllUser", {
            params: {
                page: this.page,
                number: 10
            }
        }).then((rep) => {
            if (rep.data.code == 200) {
                this.userData = rep.data.data
                return
            }
            this.$message.error(rep.data.message)
        }).catch((rep) => {
            console.log(rep);
        })
    },

    methods: {
        prevPage() {
            if (this.page > 1) {
                this.page -= 1;
                this.$axios.get("/admin/getAllUser", {
                    params: {
                        page: this.page,
                        number: 10
                    }
                }).then((rep) => {
                    if (rep.data.code == 200) {
                        this.userData = rep.data.data
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
            if (this.userData.length == 10) {
                this.page += 1;
            } else {
                this.$message.error('已经是最后一页了')
                return
            }
            this.$axios.get("/admin/getAllUser", {
                params: {
                    page: this.page,
                    number: 10
                }
            }).then((rep) => {
                if (rep.data.code == 200) {
                    this.userData = rep.data.data
                    return
                }
                console.log(rep);
            }).catch((rep) => {
                console.log(rep);
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
        chageHead() {
            window.location.href = '#chage_head';
        },
        getIdentity(identity) {
            let identityList = ['普通用户', '供应链管理员']
            return identityList[identity]
        },
        // 修改用户状态
        updateUserStatus(email, status) {
            this.$axios.post("/admin/UpdateAccountStatus", {
                email: email,
                status: status
            }).then((rep) => {
                console.log(rep)
                if (rep.data.code == 200) {
                    this.$message.success('修改成功')
                    // 遍历列表
                    for (let i = 0; i < this.userData.length; i++) {
                        if (this.userData[i].email == email) {
                            this.userData[i].status = status
                            return
                        }
                    }
                    return
                }
                this.$message.error(rep.data.message)
            }).catch((rep) => {
                console.log(rep);
            })
        },
        getStatus(status) {
            console.log(status)
            let statusList = ['未审核', '已审核', '审核未通过', '停止使用']
            return statusList[status]
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