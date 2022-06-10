import { shallowMount, $route, $router, localVue, Vuex } from '@/test/data/test-module.js'
import GamesList from '@/components/GamesList.vue'
import games from '@/test/data/games.test.json'
import _games from '@/store/games.js'


const actions = {
    "games/getGames": jest.fn()
}

const state = {
    games
};

const store = new Vuex.Store({
    actions,
    modules: {
        games: {
            state,
            getters: _games.getters,
            namespaced: true
        }
    }
});

// Inside games-list component, render a games list that user is created and played that game
describe('GamesList', () => {
    let wrapper = null

    // SETUP - run before to all unit test are started
    beforeAll(() => {

        // render the component
        wrapper = shallowMount(GamesList, {
            localVue,
            store,
            mocks: {
                $route,
                $router,
            },
        })
    });

    // TEARDOWN - run after to all unit test is complete
    afterAll(() => {
        wrapper.destroy()
    });

    test("games api called when home page render for get all games", () => {
        expect(actions["games/getGames"]).toHaveBeenCalled();
    });

    test("User can able to see games list", () => {
        const td = wrapper.findAll("td")
        expect(td.at(0).text()).toBe('Game-Name')
        expect(td.at(1).text()).toBe('Target Score-101')
        expect(td.at(2).text()).toBe('Not Started')
    });

    test("When user click on any game in games list user can redirect to game details page", () => {
        const gamesList = wrapper.find("tbody")
        const gameDetails = gamesList.find("tr")
        gameDetails.trigger('click')
        expect($router.push).toHaveBeenCalledWith('/games/1')
    });

    test("Users can see next page to click next button", () => {
        const nextPage = wrapper.find('[data-test="next"]')
        nextPage.trigger('click')
        expect(actions["games/getGames"]).toHaveBeenCalled();
    });
});
