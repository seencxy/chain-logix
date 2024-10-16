<template>
    <div class="app">
        <!--
// v0 by Vercel.
// https://v0.dev/t/e9OH4wp7Bae
-->
        <main class="flex flex-col items-center justify-center h-screen bg-gray-100 px-4 md:px-6"><svg
                xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                class="w-16 h-16 mb-4 text-red-500">
                <path d="m8 2 1.88 1.88"></path>
                <path d="M14.12 3.88 16 2"></path>
                <path d="M9 7.13v-1a3.003 3.003 0 1 1 6 0v1"></path>
                <path d="M12 20c-3.3 0-6-2.7-6-6v-3a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v3c0 3.3-2.7 6-6 6"></path>
                <path d="M12 20v-9"></path>
                <path d="M6.53 9C4.6 8.8 3 7.1 3 5"></path>
                <path d="M6 13H2"></path>
                <path d="M3 21c0-2.1 1.7-3.9 3.8-4"></path>
                <path d="M20.97 5c0 2.1-1.6 3.8-3.5 4"></path>
                <path d="M22 13h-4"></path>
                <path d="M17.2 17c2.1.1 3.8 1.9 3.8 4"></path>
            </svg>
            <h2 class="text-3xl font-bold tracking-tighter text-gray-900">糟糕!</h2>
            <p
                class="max-w-[600px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400 my-4 text-center">
                您要查找的页面不存在。
            </p><a
                class="inline-flex h-10 items-center justify-center rounded-md border border-gray-200 bg-white px-8 text-sm font-medium shadow-sm transition-colors hover:bg-gray-100 hover:text-gray-900 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-gray-950 disabled:pointer-events-none disabled:opacity-50 dark:border-gray-800 dark:border-gray-800 dark:bg-gray-950 dark:hover:bg-gray-800 dark:hover:text-gray-50 dark:focus-visible:ring-gray-300"
                @click="backHome">
                回到首页
            </a>
        </main>
    </div>
</template>

<script>
export default {
    name: '404',
    methods: {
        // 回到主页
        backHome() {
            // 跳转到主页
            this.$router.push('/index');
        }
    }
}
</script>

<style></style>