import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: '主页',
      component: () => import('@/views/seed/index'),
      meta: { title: '主页', icon: 'dashboard' }
    }]
  },

  // {
  //   path: '/exp',
  //   component: Layout,
  //   redirect: '/exp/list',
  //   name: '升级经验',
  //   meta: { title: '升级经验', icon: 'el-icon-user' },
  //   children: [
  //     {
  //       path: 'list',
  //       name: '全部植物',
  //       component: () => import('@/views/exp/index'),
  //       meta: { title: '全部升级经验', icon: 'table' }
  //     },
  //     {
  //       path: 'create',
  //       name: '新增植物',
  //       component: () => import('@/views/exp/create'),
  //       meta: { title: '新增升级经验', icon: 'el-icon-plus' }
  //     },
  //     {
  //       path: 'edit',
  //       name: '编辑植物',
  //       component: () => import('@/views/exp/edit'),
  //       meta: { title: '编辑升级经验', icon: 'el-icon-edit' },
  //       hidden: true
  //     },
  //   ]
  // },

  // {
  //   path: '/land',
  //   component: Layout,
  //   redirect: '/land/list',
  //   name: '土地等级',
  //   meta: { title: '土地等级', icon: 'el-icon-sunrise' },
  //   children: [
  //     {
  //       path: 'list',
  //       name: '全部土地等级',
  //       component: () => import('@/views/land/index'),
  //       meta: { title: '全部土地等级', icon: 'table' }
  //     },
  //     {
  //       path: 'create',
  //       name: '新增植物',
  //       component: () => import('@/views/land/create'),
  //       meta: { title: '新增土地等级', icon: 'el-icon-plus' }
  //     },
  //     {
  //       path: 'edit',
  //       name: '编辑植物',
  //       component: () => import('@/views/land/edit'),
  //       meta: { title: '编辑土地等级', icon: 'el-icon-edit' },
  //       hidden: true
  //     },
  //   ]
  // },

  // {
  //   path: '/plant',
  //   component: Layout,
  //   redirect: '/plant/list',
  //   name: '植物',
  //   meta: { title: '植物', icon: 'el-icon-cherry' },
  //   children: [
  //     {
  //       path: 'list',
  //       name: '全部植物',
  //       component: () => import('@/views/plant/index'),
  //       meta: { title: '全部植物', icon: 'table' }
  //     },
  //     {
  //       path: 'create',
  //       name: '新增植物',
  //       component: () => import('@/views/plant/create'),
  //       meta: { title: '新增植物', icon: 'el-icon-plus' }
  //     },
  //     {
  //       path: 'edit',
  //       name: '编辑植物',
  //       component: () => import('@/views/plant/edit'),
  //       meta: { title: '编辑植物', icon: 'el-icon-edit' },
  //       hidden: true
  //     },
  //   ]
  // },

  {
    path: '/seed',
    component: Layout,
    redirect: '/seed/list',
    name: '种子',
    meta: { title: '种子', icon: 'el-icon-sunrise-1' },
    children: [
      {
        path: 'list',
        name: '全部种子',
        component: () => import('@/views/seed/index'),
        meta: { title: '全部种子', icon: 'table' }
      },
      {
        path: 'create',
        name: '新增种子',
        component: () => import('@/views/seed/create'),
        meta: { title: '新增种子', icon: 'el-icon-plus' }
      },
      {
        path: 'edit',
        name: '编辑种子',
        component: () => import('@/views/seed/edit'),
        meta: { title: '编辑种子', icon: 'el-icon-plus' },
        hidden: true
      }
    ]
  },

  // {
  //   path: '/muck',
  //   component: Layout,
  //   redirect: '/muck/list',
  //   name: '肥料',
  //   meta: { title: '肥料', icon: 'el-icon-moon' },
  //   children: [
  //     {
  //       path: 'list',
  //       name: '全部肥料',
  //       component: () => import('@/views/muck/index'),
  //       meta: { title: '全部肥料', icon: 'table' }
  //     },
  //     {
  //       path: 'create',
  //       name: '新增肥料',
  //       component: () => import('@/views/muck/create'),
  //       meta: { title: '新增肥料', icon: 'el-icon-plus' }
  //     },
  //     {
  //       path: 'edit',
  //       name: '编辑肥料',
  //       component: () => import('@/views/muck/edit'),
  //       meta: { title: '编辑肥料', icon: 'el-icon-plus' },
  //       hidden: true
  //     }
  //   ]
  // },

  // {
  //   path: '/prop',
  //   component: Layout,
  //   redirect: '/prop/list',
  //   name: '狗粮',
  //   meta: { title: '狗粮', icon: 'el-icon-moon' },
  //   children: [
  //     {
  //       path: 'list',
  //       name: '全部狗粮',
  //       component: () => import('@/views/prop/index'),
  //       meta: { title: '全部狗粮', icon: 'table' }
  //     },
  //     {
  //       path: 'create',
  //       name: '新增狗粮',
  //       component: () => import('@/views/prop/create'),
  //       meta: { title: '新增狗粮', icon: 'el-icon-plus' }
  //     },
  //     {
  //       path: 'edit',
  //       name: '编辑道具',
  //       component: () => import('@/views/prop/edit'),
  //       meta: { title: '编辑狗粮', icon: 'el-icon-plus' },
  //       hidden: true
  //     }
  //   ]
  // },

  {
    path: '/shop',
    component: Layout,
    redirect: '/shop/list',
    name: '商店',
    meta: { title: '商店', icon: 'el-icon-shopping-cart-full' },
    children: [
      {
        path: 'list',
        name: '全部商品',
        component: () => import('@/views/shop/index'),
        meta: { title: '全部商品', icon: 'table' }
      },
      {
        path: 'create',
        name: '新增商品',
        component: () => import('@/views/shop/create'),
        meta: { title: '新增商品', icon: 'el-icon-plus' }
      },
      {
        path: 'edit',
        name: '编辑商品',
        component: () => import('@/views/shop/edit'),
        meta: { title: '编辑商品', icon: 'el-icon-plus' },
        hidden: true
      }
    ]
  },
  
  // {
  //   path: '/pet',
  //   component: Layout,
  //   redirect: '/pet/list',
  //   name: '宠物',
  //   meta: { title: '宠物', icon: 'el-icon-shopping-cart-full' },
  //   children: [
  //     {
  //       path: 'list',
  //       name: '宠物列表',
  //       component: () => import('@/views/pet/index'),
  //       meta: { title: '宠物列表', icon: 'table' }
  //     },
  //     {
  //       path: 'create',
  //       name: '新增宠物',
  //       component: () => import('@/views/pet/create'),
  //       meta: { title: '新增宠物', icon: 'el-icon-plus' }
  //     },
  //     {
  //       path: 'edit',
  //       name: '编辑宠物',
  //       component: () => import('@/views/pet/edit'),
  //       meta: { title: '编辑宠物', icon: 'el-icon-plus' },
  //       hidden: true
  //     }
  //   ]
  // },
  
  // {
  //   path: '/buff',
  //   component: Layout,
  //   redirect: '/buff/list',
  //   name: 'BUFF',
  //   meta: { title: 'BUFF', icon: 'el-icon-shopping-cart-full' },
  //   children: [
  //     {
  //       path: 'list',
  //       name: 'BUFF列表',
  //       component: () => import('@/views/buff/index'),
  //       meta: { title: 'BUFF列表', icon: 'table' }
  //     },
  //     {
  //       path: 'create',
  //       name: '新增BUFF',
  //       component: () => import('@/views/buff/create'),
  //       meta: { title: '新增BUFF', icon: 'el-icon-plus' }
  //     },
  //     {
  //       path: 'edit',
  //       name: '编辑BUFF',
  //       component: () => import('@/views/buff/edit'),
  //       meta: { title: '编辑BUFF', icon: 'el-icon-plus' },
  //       hidden: true
  //     }
  //   ]
  // },

  // {
  //   path: '/monthtask',
  //   component: Layout,
  //   redirect: '/shop/list',
  //   name: '通行证',
  //   meta: { title: '通行证', icon: 'el-icon-shopping-cart-full' },
  //   children: [
  //     {
  //       path: 'list',
  //       name: '等级列表',
  //       component: () => import('@/views/monthtask/index'),
  //       meta: { title: '等级列表', icon: 'table' }
  //     },
  //     {
  //       path: 'create',
  //       name: '新增等级',
  //       component: () => import('@/views/monthtask/create'),
  //       meta: { title: '新增等级', icon: 'el-icon-plus' }
  //     },
  //     {
  //       path: 'edit',
  //       name: '编辑等级',
  //       component: () => import('@/views/monthtask/edit'),
  //       meta: { title: '编辑等级', icon: 'el-icon-plus' },
  //       hidden: true
  //     }
  //   ]
  // },
  
  // {
  //   path: '/exchange',
  //   component: Layout,
  //   redirect: '/exchange/list',
  //   name: '兑换码',
  //   meta: { title: '兑换码', icon: 'el-icon-shopping-cart-full' },
  //   children: [
  //     {
  //       path: 'list',
  //       name: '兑换码列表',
  //       component: () => import('@/views/exchange/index'),
  //       meta: { title: '兑换码列表', icon: 'table' }
  //     },
  //     {
  //       path: 'create',
  //       name: '新增BUFF',
  //       component: () => import('@/views/exchange/create'),
  //       meta: { title: '新增兑换码', icon: 'el-icon-plus' }
  //     }
  //   ]
  // },

  // {
  //   path: '/mail',
  //   component: Layout,
  //   redirect: '/mail/list',
  //   name: '邮件',
  //   meta: { title: '邮件', icon: 'el-icon-shopping-cart-full' },
  //   children: [
  //     {
  //       path: 'create',
  //       name: '系统邮件',
  //       component: () => import('@/views/mail/create'),
  //       meta: { title: '系统邮件', icon: 'table' }
  //     }
  //   ]
  // },

  // {
  //   path: '/nested',
  //   component: Layout,
  //   redirect: '/nested/menu1',
  //   name: 'Nested',
  //   meta: {
  //     title: 'Nested',
  //     icon: 'nested'
  //   },
  //   children: [
  //     {
  //       path: 'menu1',
  //       component: () => import('@/views/nested/menu1/index'), // Parent router-view
  //       name: 'Menu1',
  //       meta: { title: 'Menu1' },
  //       children: [
  //         {
  //           path: 'menu1-1',
  //           component: () => import('@/views/nested/menu1/menu1-1'),
  //           name: 'Menu1-1',
  //           meta: { title: 'Menu1-1' }
  //         },
  //         {
  //           path: 'menu1-2',
  //           component: () => import('@/views/nested/menu1/menu1-2'),
  //           name: 'Menu1-2',
  //           meta: { title: 'Menu1-2' },
  //           children: [
  //             {
  //               path: 'menu1-2-1',
  //               component: () => import('@/views/nested/menu1/menu1-2/menu1-2-1'),
  //               name: 'Menu1-2-1',
  //               meta: { title: 'Menu1-2-1' }
  //             },
  //             {
  //               path: 'menu1-2-2',
  //               component: () => import('@/views/nested/menu1/menu1-2/menu1-2-2'),
  //               name: 'Menu1-2-2',
  //               meta: { title: 'Menu1-2-2' }
  //             }
  //           ]
  //         },
  //         {
  //           path: 'menu1-3',
  //           component: () => import('@/views/nested/menu1/menu1-3'),
  //           name: 'Menu1-3',
  //           meta: { title: 'Menu1-3' }
  //         }
  //       ]
  //     },
  //     {
  //       path: 'menu2',
  //       component: () => import('@/views/nested/menu2/index'),
  //       name: 'Menu2',
  //       meta: { title: 'menu2' }
  //     }
  //   ]
  // },

  // {
  //   path: 'external-link',
  //   component: Layout,
  //   children: [
  //     {
  //       path: 'https://panjiachen.github.io/vue-element-admin-site/#/',
  //       meta: { title: 'External Link', icon: 'link' }
  //     }
  //   ]
  // },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
