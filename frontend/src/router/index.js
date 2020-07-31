import Vue from 'vue'
import Router from 'vue-router'
import store from '@/store'

import IndexPage from '../pages/IndexPage'
import LoginPage from "../pages/LoginPage";
import RegisterPage from "../pages/RegisterPage";
import {routes} from '../router/routes'
import FlowPage from "../pages/FlowPage";
import NotFoundPage from "../pages/NotFoundPage";
import {checkMultipleVue} from "bootstrap-vue/esm/utils/plugins";

Vue.use(Router);

const router = new Router({
    mode: 'history',
    routes: [
        {
            path: routes.index,
            name: 'index',
            component: IndexPage
        },
        {
            path: routes.flow,
            name: 'flow',
            component: FlowPage
        },
        {
            path: routes.login,
            name: 'login',
            component: LoginPage
        },
        {
            path: routes.register,
            name: 'register',
            component: RegisterPage
        },
        {
            path: '*',
            name: '404',
            component: NotFoundPage
        },
    ]
});

/**
 * Валидация перед переходом на внутренню страницу
 *
 * @param {object} to - Слудующий роут
 * @param {object} from - Предыдущий роут
 * @param {function} next - Функция перехода на вызванный
 */
const ifAuthorized = async (to, from, next) => {
    // if (null !== store.getters['user/getUser']) {
    //     next();
    //     return ;
    // }
    //
    // await store.dispatch(`user/checkAuth`);
    //
    // if (null !== store.getters['user/getUser']) {
    //     next();
    // } else {
    //     next({name: `login`});
    // }
};

/**
 * Валидация перед переходом на внешнюю страницу (авторизация, регистрация и тд)
 *
 * @param {object} to - Слудующий роут
 * @param {object} from - Предыдущий роут
 * @param {function} next - Функция перехода на вызванный
 */
const ifNotAuthorized = async (to, from, next) => {
    // if (null !== store.getters['user/getUser']) {
    //
    //     await store.dispatch(`user/checkAuth`);
    //
    //     if (null === store.getters['user/getUser']) {
    //         next();
    //         return;
    //     }
    // }
    //
    // next({name: `home`});
};

const CheckPublicPathsAvailability = async (to, from, next) => {
    if (null !== store.getters['user/getUser']) {
        next({ name: `flow` });
        return;
    }

    await store.dispatch(`user/checkAuth`);
    if (null !== store.getters['user/getUser']) {
        next({ name: `flow` });
        return;
    }

    next();
};

const CheckAuthPagesAvailability = async (to, from, next) => {
    if (null === store.getters['user/getUser']) {
        await store.dispatch(`user/checkAuth`);
        if (null === store.getters['user/getUser']) {
            next();
            return;
        }
    }

    next({name: 'flow'});
};

const CheckPrivetPagesAvailability = async (to, from, next) => {
    if (null !== store.getters['user/getUser']) {
        next();
        return;
    }

    await store.dispatch(`user/checkAuth`);
    if (null !== store.getters['user/getUser']) {
        next();
    } else {
        next({name: `login`});
    }
};



router.beforeEach(async (to, from, next) => {
    switch (to.name) {
        case 'index':
            await CheckPublicPathsAvailability(to, from, next);
            return;

        case 'login':
        case 'register':
            await CheckAuthPagesAvailability(to, from, next);
            return;

        default:
            await CheckPrivetPagesAvailability(to, from, next);
            return;
    }
});

export default router
