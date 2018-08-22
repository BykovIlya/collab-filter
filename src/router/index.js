import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Products from '@/components/Products'
import Users from '@/components/Users'
import Recommendations from '@/components/Recommendations'

Vue.use(Router)

const routes = [
    {
      name: 'Index',
      path: '/',
      component: HelloWorld,
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
    }
    ];

const router = new Router({routes,mode:'history'});

export default router;
