import { shallowMount, $config } from '@/test/data/test-module.js'
import SignIn from '@/components/SignIn.vue'
import { BCard } from 'bootstrap-vue';

// SignIn component include sign-in button to authorize our application.
describe('SignIn', () => {
    let wrapper = null

    // SETUP - run before to all unit test are started
    beforeAll(() => {
        // render the component
        wrapper = shallowMount(SignIn, {
            mocks: {
                $config
            },
            stubs: {
                "b-card": BCard
            }
        })
    });

    // TEARDOWN - run after to all unit test is complete
    afterAll(() => {
        wrapper.destroy()
    });

    test("User can able to access home page using sign in button", async () => {
        const signInButton = wrapper.find('[data-test="login-button"]');
        expect(signInButton.text()).toBe('Sign In With Google')
        expect(signInButton.attributes().href).toBe($config.loginURL);
    });
})
