import { createRouter, createWebHistory } from 'vue-router'

const routes = [
    {
        name: 'welcome',
        path: '/',
        component: () => import('@/views/welcome.vue'),
    },
    {
        name: 'login',
        path: '/login',
        component: () => import('@/views/user/login.vue')
    },
    {
        name: 'register',
        path: '/register',
        component: () => import('@/views/user/register.vue')
    },
    {
        name: 'set-password',
        path: '/set-password',
        component: () => import('@/views/user/reset-password.vue')
    },
    {
        name: '404',
        path: '/:catchAll(.*)',
        component: () => import('@/views/404.vue') //  当路由不存在则到该路由
    },
    {
        name: 'index',
        path: '/index',
        component: () => import('@/views/main/index.vue')
    },
    {
        name: 'sendParcel',
        path: '/sendParcel',
        component: () => import('@/views/main/sendParcel.vue')
    },
    {
        name: 'mannage',
        path: '/mannage',
        component: () => import('@/views/main/mannage.vue')
    },
    {
        name: 'userMannage',
        path: '/userMannage',
        component: () => import('@/views/main/userMannage.vue')
    },
    {
        name: 'transaction',
        path: '/transaction',
        component: () => import('@/views/main/transaction.vue')
    },
    {
        name: 'traceShipment',
        path: '/traceShipment',
        component: () => import('@/views/main/traceShipment.vue')
    },

]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

// 全局导守卫
router.beforeEach((to, from, next) => {
    // to 满足条件的url放过
    if (to.path == "/login" || to.path == "/register" || to.path == "/set-password" || to.path == "/" || to.path == "/404") {
        next()  // 直接放过
    } else {
        // 校验用户的，不受管理员影响
        const auth_token = localStorage.getItem("auth_token")
        // front的需要权限的url
        if (auth_token == "" || auth_token == null || auth_token == '') {
            next("/login")
        } else {
            next()
        }
    }
})

export default router