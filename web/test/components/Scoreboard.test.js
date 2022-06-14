import { shallowMount, $route, $router, localVue, Vuex } from '@/test/data/test-module.js'
import Scoreboard from '@/components/Scoreboard.vue'
import games from '@/test/data/games.test.json'
import scoreboard from '@/test/data/scoreboard.test.json'
import _game from '@/store/game.js'
import { BCollapse } from 'bootstrap-vue'

const state = {
    game: games.list[0],
    scoreboard: scoreboard
};

const actions = {
    "game/getGame": jest.fn(),
    "game/getScoreboard": jest.fn()
}

const store = new Vuex.Store({
    actions,
    modules: {
        game: {
            state,
            getters: _game.getters,
            namespaced: true
        }
    }
});
describe('Scoreboard', () => {
    let wrapper = null

    // SETUP - run before to all unit test are started
    beforeAll(() => {

        // render the component
        wrapper = shallowMount(Scoreboard, {
            store,
            localVue,
            mocks: {
                $router,
                $route,
            },
            stubs: {
                'b-collapse': BCollapse,
                'Speak': true
            }
        });
    });

    // TEARDOWN - run after to all unit test is complete
    afterAll(() => {
        wrapper.destroy()
    });

    test("Game and scoreboard api called when scoreboard render for get all details of a current game scoreboard", () => {
        expect(actions["game/getGame"]).toHaveBeenCalled();
        expect(actions["game/getScoreboard"]).toHaveBeenCalled();
    });

    test("User can able to see game Details", () => {
        // check game details
        expect(wrapper.find('[data-test="game-title"]').text()).toBe('Game-Name')
        expect(wrapper.find('[data-test="game-type"]').text()).toBe('Target Score-101')
    });

    test("User can able to see scoreboard details", () => {
        // check scoreboard details (player-name, total)
        const playerTotal = wrapper.find('[data-test="player-details"]')
        expect(playerTotal.findAll('td').at(1).text()).toBe('Jeel Rupapara')
        expect(playerTotal.findAll('td').at(2).text()).toBe('14') //remaining score

        // check scoreboard details (player scores in all rounds)
        const playerScore = wrapper.find('[data-test="player-score"]')
        const scores = playerScore.findAll('div')

        expect(scores.at(1).text()).toBe('1')        // round-1
        expect(scores.at(2).text()).toBe('10')
        expect(scores.at(3).text()).toBe('12')
        expect(scores.at(4).text()).toBe('13')
        expect(scores.at(5).text()).toBe('35')       // round-total

        expect(scores.at(6).text()).toBe('2')        // round-2
        expect(scores.at(7).text()).toBe('1')
        expect(scores.at(8).text()).toBe('2')
        expect(scores.at(9).text()).toBe('3')
        expect(scores.at(10).text()).toBe('6')       // round-total

        expect(scores.at(11).text()).toBe('3')       // round-3
        expect(scores.at(12).text()).toBe('2')
        expect(scores.at(13).text()).toBe('32')
        expect(scores.at(14).text()).toBe('12')
        expect(scores.at(15).text()).toBe('46')      // round-total

        // check winner
        expect(wrapper.find('[data-test="winner-name"]').text()).toBe('Jeel Rupapara')
    });


    test("User able to go home page", () => {
        const homeButton = wrapper.find("button")
        expect(homeButton.text()).toBe('Home Page')
        homeButton.trigger('click')
        expect($router.push).toHaveBeenCalledWith('/')
    });
})
