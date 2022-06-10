import { shallowMount, $router, localVue } from '@/test/data/test-module.js'
import Header from '@/components/Header.vue'

// Header component include app name and new-game button
describe('Header', () => {
    let wrapper = null

    // SETUP - run before to all unit test are started
    beforeAll(() => {
        // render the component
        wrapper = shallowMount(Header, {
            localVue,
            mocks: {
                $router
            },
            stubs: {
                fa: true
            }
        })
    });

    // TEARDOWN - run after to all unit test is complete
    afterAll(() => {
        wrapper.destroy()
    });

    test("When user click 'new-game' button then user redirect to new game page", () => {
        const newGameButton = wrapper.find("button")
        expect(newGameButton.text()).toBe('New Game')
        newGameButton.trigger('click')

        expect($router.push).toHaveBeenCalledWith('/games/new')
    });
})
