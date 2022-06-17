import {
  shallowMount,
  $route,
  $router,
  localVue,
  Vuex,
} from "@/test/data/test-module.js";
import StartGame from "@/components/StartGame.vue";
import games from "@/test/data/games.test.json";
import player from "@/test/data/playerInfo.test.json";
import scoreboard from "@/test/data/scoreboard.test.json";
import currentTurn from "@/test/data/turn.test.json";
import _game from "@/store/game.js";

const actions = {
  "game/getGamePlayerInfo": jest.fn(),
  "game/addScore": jest.fn(),
  "game/getScoreboard": jest.fn(),
  "game/getCurrentTurn": jest.fn(),
  "game/undoScore": jest.fn(),
};

const state = {
  game: games.list[0],
  player,
  scoreboard,
  currentTurn,
};

const store = new Vuex.Store({
  actions,
  modules: {
    game: {
      state,
      getters: _game.getters,
      namespaced: true,
    },
  },
});

// Start game component accept your turn score as input as well as show scoreboard and current player info.
describe("StartGame", () => {
  let wrapper = null;

  // SETUP - run before to all unit test are started
  beforeAll(() => {
    // render the component
    wrapper = shallowMount(StartGame, {
      localVue,
      store,
      window,
      mocks: {
        $route,
        $router,
      },
      stubs: {
        NavBar: true,
        Speak: true,
        Dartboard: true,
      },
    });
  });

  // TEARDOWN - run after to all unit test is complete
  afterAll(() => {
    wrapper.destroy();
  });

  test("", () => {});
});
