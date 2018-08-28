import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Products from '@/components/Products'
import Users from '@/components/Users'
import Recommendations from '@/components/Recommendations'
import Tests from '@/components/Tests'

Vue.use(Router)

const routes = [
    {
      name: 'Home',
      path: '/',
      component: Home,
    },
    {
      name: 'Users',
      path: '/users',
      component: Users,
    },
    {
      name: 'Products',
      path: '/products',
      component: Products,
    },
    {
      name: 'Recommendations',
      path: '/recommendations',
      component: Recommendations,
    },
    {
      name: 'Tests',
      path: '/tests',
      component: Tests,
    }
    ];

const router = new Router({routes,mode:'history'});

export default router;
