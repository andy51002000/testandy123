import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('../components/HelloWorld.vue')
  },
  {
    path: '/test',
    name: 'test',
    component: () => import('../components/TestView.vue')
  },
  {
    path: '/transaction/:memberId',
    name: 'transaction',
    props: (route) => ({
      memberId: route.params.memberId,
      startTime: route.query.startTime,
      endTime: route.query.endTime,
    }),
    component: () => import('../components/TransactionView.vue'), 
  },
  {
    path: '/member-update/:memberId',
    name: 'member-update',
    props: (route) => ({
      memberId: route.params.memberId,
    }),
    component: () => import('../components/MemberUpdate.vue'), 
  },
  {
    path: '/member-get/:memberId',
    name: 'member-get',
    props: (route) => ({
      memberId: route.params.memberId,
    }),
    component: () => import('../components/MemberGet.vue'), 
  },
  {
    path: '/member-create/:memberId',
    name: 'member-create',
    component: () => import('../components/MemberCreate.vue'), 
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router