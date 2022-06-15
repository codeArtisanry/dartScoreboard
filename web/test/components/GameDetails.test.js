import { shallowMount, $route, $router, localVue, Vuex } from '@/test/data/test-module.js'
import GameDetails from '@/components/GameDetails.vue'
import games from '@/test/data/games.test.json'
import _game from '@/store/game.js'

const actions = {
    "game/getGame": jest.fn(),
    "game/deleteGame": jest.fn(),
}

const state = {
    game: games.list[0],
};

const mutations = {
    CHANGE_STATUS(state, status) {
        state.game.status = status;
    },
    CHANGE_OWNER(state, owner_name) {
        state.game.creater_name = owner_name
    }
}

const store = new Vuex.Store({
    state: {
        auth: {
            token: {
                name: 'Vatsal Chauhan'
            }
        },
    },
    actions,
    modules: {
        game: {
            state,
            mutations,
            getters: _game.getters,
            namespaced: true
        }
    }
});
const componentData = {
    localVue,
    store,
    mocks: {
        $route,
        $router,
    },
    stubs: {
        NavBar: true
    },
};

describe('GameDetails', () => {
    let wrapper = null

    // SETUP - run before to all unit test are started
    beforeEach(() => {

        // render the component
        wrapper = shallowMount(GameDetails, componentData)
    });

    // TEARDOWN - run after to all unit test is complete
    afterEach(() => {
        wrapper.destroy()
    });

    test("User can able to see game details", () => {
        expect(wrapper.find('[data-test="game-name"]').text()).toBe("Game-Name")
        expect(wrapper.find('[data-test="game-type"]').text()).toBe("Target Score-101")
        expect(wrapper.find('[data-test="creater-name"]').text()).toBe("Vatsal Chauhan")
        expect(wrapper.findAll('[data-test="players-name"]').at(0).text()).toBe("Payal")
        expect(wrapper.findAll('[data-test="players-name"]').at(1).text()).toBe("Jeel")
    });

    test("User can able to update game", () => {
        //check update button
        const updateButton = wrapper.find('[data-test="update-button"]')
        expect(updateButton.text()).toBe("Update")
        updateButton.trigger('click')
        expect($router.push).toHaveBeenCalledWith('/games/1/update')
    });

    test("User can able to delete game", () => {
        //check delete button
        const deleteButton = wrapper.find('[data-test="delete-button"]')
        expect(deleteButton.text()).toBe("Delete")
        deleteButton.trigger('click')
        expect(actions["game/deleteGame"]).toHaveBeenCalled();
        expect($router.push).toHaveBeenCalledWith('/')
    });

    test("User can able to start game", () => {
        //check start button
        const startButton = wrapper.find('[data-test="game-state-button"]')
        expect(startButton.text()).toBe("Start")
        startButton.trigger('click')
        expect($router.push).toHaveBeenCalledWith('/games/1/player')
    });

    test("User can go to home page", () => {
        //check back to home button
        const homeButton = wrapper.find('[data-test="home-button"]')
        expect(homeButton.text()).toBe("back to home")
        homeButton.trigger('click')
        expect($router.push).toHaveBeenCalledWith('/')
    });

    test("User can resume game", () => {
        wrapper.vm.$store.commit("game/CHANGE_STATUS", 'In Progress')
        wrapper.vm.updateButtonName()
        expect(wrapper.vm.buttonName).toBe("Resume")
        const resumeButton = wrapper.find('[data-test="game-state-button"]')
        resumeButton.trigger('click')
        expect($router.push).toHaveBeenCalledWith('/games/1/player')
    });

    test("User can able to see scoreboard if game is completed", () => {
        wrapper.vm.$store.commit("game/CHANGE_STATUS", 'Completed')
        wrapper.vm.updateButtonName()
        expect(wrapper.vm.buttonName).toBe("Scoreboard")
        const scoreboardButton = wrapper.find('[data-test="game-state-button"]')
        scoreboardButton.trigger('click')
        expect($router.push).toHaveBeenCalledWith('/games/1/scoreboard')
    });

    test("User is game owner or not?", () => {
        wrapper.vm.$store.commit("game/CHANGE_OWNER", 'Payal Ravya')
        expect(wrapper.vm.isOwner).toBe(false)
    });

    test("If user is not create that game then user not able to update and delete game", () => {
        const deleteButton = wrapper.find('[data-test="delete-button"]')
        expect(deleteButton.exists()).toBe(false)
        const updateButton = wrapper.find('[data-test="update-button"]')
        expect(updateButton.exists()).toBe(false)
    });

});
