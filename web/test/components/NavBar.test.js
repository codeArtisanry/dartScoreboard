import { mount, localVue, $router, Vuex, $config } from '@/test/data/test-module.js'
import NavBar from '@/components/NavBar.vue'

const state = {
    auth: {
        token: {
            name: 'Jeel Rupapara',
            picture: 'https://lh3.googleusercontent.com/a/AATXAJzziht9CYV5whqpEGVDjLDp_h0oaWW_NXNPg4TS=s96-c',
            email: 'jeel@improwised.com'
        }
    }
};

const store = new Vuex.Store({
    state
});

describe('NavBar', () => {
    let wrapper = null

    // SETUP - run before to all unit test are started
    beforeAll(() => {
        // render the component
        wrapper = mount(NavBar, {
            localVue,
            store,
            mocks: {
                $router,
                $config
            }
        })
    });

    // TEARDOWN - run after to all unit test is complete
    afterAll(() => {
        wrapper.destroy()
    });

    test("User can able to open side bar menu using sandwich button on navbar", () => {
        // check navbar sandwich button is working or not?
        const navButton = wrapper.find('button');
        expect(navButton.exists()).toBe(true);
        expect(navButton.text()).toBe('');
        navButton.trigger('click'); // open side bar

        // check sidebar exits or not?
        const sideBar = wrapper.find('#sidebar-no-header');
        expect(sideBar.exists()).toBe(true);
    });

    test("User can see profile details in side bar menu", () => {
        expect(wrapper.find('[data-test="profile-pic"]').attributes().src).toBe("https://lh3.googleusercontent.com/a/AATXAJzziht9CYV5whqpEGVDjLDp_h0oaWW_NXNPg4TS=s96-c");
        expect(wrapper.find('[data-test="user-name"]').text()).toBe('Jeel Rupapara')
        expect(wrapper.find('[data-test="user-email"]').text()).toBe('jeel@improwised.com')
    })

    test("User can able to go home page using home button", () => {
        // check navbar home button
        const homeButton = wrapper.find('[data-test="home-button"]');
        expect(homeButton.exists()).toBe(true);
        expect(homeButton.text()).toBe('Home');
        homeButton.trigger('click');
        expect($router.push).toHaveBeenCalledWith('/')
    });

    test("User can logout app using logout button", () => {
        // check navbar logout button
        const LogoutButton = wrapper.find('[data-test="logout-button"]');
        expect(LogoutButton.exists()).toBe(true);
        expect(LogoutButton.text()).toBe('Sign Out');
        LogoutButton.trigger('click');
        expect(LogoutButton.attributes().href).toBe($config.logoutURL);
    })

});
