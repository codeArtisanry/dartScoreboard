import { shallowMount, createLocalVue, mount } from '@vue/test-utils'
import BootstrapVue from 'bootstrap-vue';
import Vuex from 'vuex'

const localVue = createLocalVue();
localVue.use(BootstrapVue)
localVue.use(Vuex)

export { localVue, shallowMount, mount, Vuex };

export const $route = {
    params: {
        gameid: 1
    }
};
export const $router = {
    push: jest.fn()
};

export const $config = {
    logoutURL: "http://localhost:8080/logout/google"
}
