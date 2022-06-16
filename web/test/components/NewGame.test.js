import { shallowMount, $router, $route, localVue, Vuex } from '@/test/data/test-module.js'
import NewGame from '@/components/NewGame.vue'
import games from '@/test/data/games.test.json'
import users from '@/test/data/users.test.json'
import _game from '@/store/game.js'
import _users from '@/store/users.js'


const actions = {
    "game/getGame": jest.fn(),
    "game/createGame": jest.fn(),
    "game/updateGame": jest.fn(),
    "users/getUsers": jest.fn()
}

const stateGame = {
    game: games.list[0]
};

const stateUsers = {
    users
}

const store = new Vuex.Store({
    actions,
    modules: {
        game: {
            state: stateGame,
            getters: _game.getters,
            namespaced: true
        },
        users: {
            state: stateUsers,
            getters: _users.getters,
            namespaced: true
        }
    }
});

const componentData = {
    localVue,
    store,
    mocks: {
        $route: {
            params: {
                gameid: undefined
            }
        },
        $router,
    },
    stubs: {
        NavBar: true,
    },
}

// NewGame component use for create a new game
describe('NewGame', () => {
    let wrapper = null

    // SETUP - run before to each unit test are started
    beforeEach(() => {

        // render the component
        wrapper = shallowMount(NewGame, componentData)
    });

    // TEARDOWN - run after to each unit test is complete
    afterEach(() => {
        wrapper.destroy()
    });

    test("User can able to register a new game", () => {
        // call users api
        expect(actions["users/getUsers"]).toHaveBeenCalled();

        //check game name input box
        const gameNameInput = wrapper.find('input[type="text"]')
        gameNameInput.setValue('Game-Name')
        expect(wrapper.find('input[type="text"]').element.value).toBe('Game-Name')

        //check game type select options
        const gameTypeInput = wrapper.find('select')
        gameTypeInput.setValue('High Score')
        expect(wrapper.find('select').element.value).toBe('High Score')

        // select players in multiselect options
        wrapper.vm.players = wrapper.vm.usersList

        // check register button
        const registerButton = wrapper.find('[data-test="register"]')
        expect(registerButton.text()).toBe('Register')
        registerButton.trigger('click')

        // call create game api
        expect(actions["game/createGame"]).toHaveBeenCalled();

        // check payload for create game
        expect(wrapper.vm.game).toStrictEqual({
            "name": "Game-Name",
            "players": [1, 2, 3],
            "type": "High Score",
        })
    });

    test("User can see option in format for select players", () => {
        const optionFormat = wrapper.vm.optionsFormat({ first_name: "Jeel", last_name: "Rupapara", email: "jeel@improwised.com" })
        expect(optionFormat).toBe("Jeel Rupapara — [jeel@improwised.com]")
    });

    test("If user not give any name of game when create a new game then generate name", () => {
        // select players in multiselect options
        wrapper.vm.players = wrapper.vm.usersList

        // click register button
        const registerButton = wrapper.find('[data-test="register"]')
        registerButton.trigger('click')

        // call create game api
        expect(actions["game/createGame"]).toHaveBeenCalled();

        expect($router.push).toHaveBeenCalledWith('/')

        // check generated game name
        expect(wrapper.vm.name).toBe("Jeel (+2 others)")
    });

    test("If user not select any player then show alert", () => {
        // click register button
        const registerButton = wrapper.find('[data-test="register"]')
        registerButton.trigger('click')

        // show alert
        expect(wrapper.vm.alert).toBe(true)
    });

    test("User can able to update a game", () => {
        componentData.mocks.$route = $route
        const wrapperD = shallowMount(NewGame, componentData)

        // update game name
        const gameNameInput = wrapperD.find('input[type="text"]')
        expect(wrapperD.vm.name).toBe("")
        gameNameInput.setValue('Update-Name')

        // update game type
        const gameTypeInput = wrapperD.find('select')
        expect(wrapperD.vm.type).toBe('High Score')
        gameTypeInput.setValue('Target Score-501')

        // update players
        wrapperD.vm.players = [{
            "id": 1,
            "first_name": "Jeel",
            "last_name": "Rupapara",
            "email": "jeel@improwised.com"
        }]

        // update game details
        const updateButton = wrapperD.find('[data-test="update"]')
        expect(updateButton.text()).toBe('Update')
        updateButton.trigger('click')

        // call update game api
        expect(actions["game/updateGame"]).toHaveBeenCalled();

        expect($router.push).toHaveBeenCalledWith('/')

        // check received update payload
        expect(wrapperD.vm.game).toStrictEqual({
            "name": "Update-Name",
            "players": [1],
            "type": "Target Score-501",
        })
    });

    test("If user deselect all players when update a game then show alert", () => {
        // deselect all players
        wrapper.vm.players = []

        // click update button
        const updateButton = wrapper.find('[data-test="update"]')
        expect(updateButton.text()).toBe('Update')
        updateButton.trigger('click')
        expect($router.push).toHaveBeenCalledWith('/')

        // show alert
        expect(wrapper.vm.alert).toBe(true)
    });
});
