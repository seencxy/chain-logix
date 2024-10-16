<template>
    <div class="app">
        <div class="flex flex-col min-h-screen py-6 bg-white dark:bg-gray-900">
            <header class="px-4 sm:px-6 lg:px-8">
                <div class="flex items-center justify-between py-6">
                    <div class="flex items-center"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                            viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                            stroke-linejoin="round" class="w-8 h-8">
                            <path
                                d="M17.8 19.2 16 11l3.5-3.5C21 6 21.5 4 21 3c-1-.5-3 0-4.5 1.5L13 8 4.8 6.2c-.5-.1-.9.1-1.1.5l-.3.5c-.2.5-.1 1 .3 1.3L9 12l-2 3H4l-1 1 3 2 2 3 1-1v-3l3-2 3.5 5.3c.3.4.8.5 1.3.3l.5-.2c.4-.3.6-.7.5-1.2z">
                            </path>
                        </svg><span class="sr-only">SaaS App</span></div>
                    <div class="hidden md:flex items-center justify-end md:flex-1 lg:w-0"><a
                            class="whitespace-nowrap text-base font-medium text-gray-500 hover:text-gray-900"
                            @click="gotoLogin">
                            登录
                        </a><a
                            class="ml-8 whitespace-nowrap inline-flex items-center justify-center px-4 py-2 border border-transparent rounded-md shadow-sm text-base font-medium text-white bg-indigo-600 hover:bg-indigo-700"
                            @click="gotoRegister">
                            注册
                        </a></div>
                </div>
            </header>
            <main class="mt-16 mx-auto max-w-7xl px-4 sm:mt-24">
                <div class="text-center">
                    <h1 class="text-4xl tracking-tight font-extrabold text-gray-900 sm:text-5xl md:text-6xl"><span
                            class="block xl:inline">值得信赖的</span><span class="block text-indigo-600 xl:inline">
                            物流管理平台</span></h1>
                    <p class="mt-3 max-w-md mx-auto text-base text-gray-500 sm:text-lg md:mt-5 md:text-xl md:max-w-3xl">
                        基于区块链的物流管理平台通过其不可篡改的数据记录、高度透明度、智能合约自动化和去中心化特性，相比传统物流管理系统提供了更高的安全性、效率和可追溯性。
                    </p>
                </div>
            </main>
            <section class="mt-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <h2 class="text-3xl font-extrabold tracking-tight text-gray-900">
                    身份分配
                </h2>
                <div class="mt-6 grid gap-5 max-w-lg mx-auto lg:grid-cols-3 lg:max-w-none">
                    <!-- 供应链管理员（Supply Chain Manager）卡片 -->
                    <div class="card w-96 bg-base-100 shadow-xl">
                        <figure><img src="@/assets/manage.webp" alt="Shoes" /></figure>
                        <div class="card-body">
                            <h2 class="card-title">
                                供应链管理员（Supply Chain Manager）
                                <div class="badge badge-secondary">Approve</div>
                            </h2>
                            <p>1.监控和管理整个供应链流程。</p>
                            <p>2.使用区块链平台来追踪货物从生产到交付的每个阶段。</p>
                            <p>3.确保所有交易和记录在区块链上准确无误地执行和更新。</p>
                            <p>4.分析供应链数据，以优化流程和减少浪费。</p>
                            <div class="card-actions justify-end">
                                <button class="btn btn-primary" @click="gotoRegister(1)">注册（Approve）</button>
                            </div>
                        </div>
                    </div>
                    <!-- 用户（User）卡片 -->
                    <div class="card w-96 bg-base-100 shadow-xl" style="margin-left: 10%;">
                        <figure><img src="@/assets/user.webp" alt="Shoes" /></figure>
                        <div class="card-body">
                            <h2 class="card-title">
                                用户（User）
                                <div class="badge badge-secondary">Any</div>
                            </h2>
                            <p>1.在平台上注册和管理个人账户。</p>
                            <p>2.追踪货物运输状态和位置。</p>
                            <p>3.查看和管理个人交易和记录。</p>
                            <p>4.与供应链管理员和其他用户交流协作。</p>
                            <div class="card-actions justify-end">
                                <button class="btn btn-primary" @click="gotoRegister(2)">注册</button>
                            </div>
                        </div>
                    </div>
                </div>

            </section>
            <footer class="mt-12 border-t border-gray-200 py-8">
                <p class="mt-1 text-center text-base text-gray-500" style="font-weight: bolder;font-size:10px">
                    © 2023 湖南科技职业学院. All rights reserved.
                </p>
            </footer>
        </div>
    </div>
</template>

<script>
export default {
    name: 'welcome',
    methods: {
        gotoRegister(id) {
            if (id == 1) {
                this.$router.push('/register?id=1');
            } else if (id == 2) {
                this.$router.push('/register?id=2');
            } else {
                this.$router.push('/register');
            }
        },
        gotoLogin() {
            this.$router.push('/login');
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