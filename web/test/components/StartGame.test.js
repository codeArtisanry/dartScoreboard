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
  player: player[0],
  scoreboard,
  currentTurn: currentTurn[0],
};

const store = new Vuex.Store({
  state: {
    auth: {
      token: "",
    },
  },
  actions,
  modules: {
    game: {
      state,
      getters: _game.getters,
      mutations: _game.mutations,
      namespaced: true,
    },
  },
});

const componentData = {
  localVue,
  store,
  mocks: {
    $route,
    $router,
  },
  stubs: {
    NavBar: true,
    Speak: true,
    Dartboard: true,
  },
};
// Start game component accept your turn score as input as well as show scoreboard and current player info.
describe("StartGame", () => {
  let wrapper = null;

  // SETUP - run before to all unit test are started
  beforeEach(() => {
    // render the component
    wrapper = shallowMount(StartGame, componentData);
  });

  // TEARDOWN - run after to all unit test is complete
  afterEach(() => {
    wrapper.destroy();
  });

  test("When Start game page is render call apis to get required details", () => {
    expect(actions["game/getCurrentTurn"]).toHaveBeenCalled();
    expect(actions["game/getCurrentTurn"].mock.calls[0][1]).toEqual(1);
    expect(wrapper.vm.turn).toBe(currentTurn[0]);

    expect(actions["game/getGamePlayerInfo"]).toHaveBeenCalled();
    expect(actions["game/getGamePlayerInfo"].mock.calls[0][1]).toEqual({
      gameId: 1,
      playerId: 4,
    });
    expect(wrapper.vm.playerInfo).toBe(player[0]);

    expect(actions["game/getScoreboard"]).toHaveBeenCalled();
    expect(actions["game/getScoreboard"].mock.calls[0][1]).toEqual(1);
    expect(wrapper.vm.scoreboard).toBe(scoreboard);
  });

  test("User can able to see dartboard to add scores", () => {
    expect(wrapper.find("#dartboard").exists()).toBe(true);
  });

  test("User can able to see current player details(name, round, turn, score))", () => {
    expect(wrapper.find("[test-data='player-name']").text()).toBe(
      "Jeel Rupapara"
    );
    expect(wrapper.find("[test-data='player-round-turn']").text()).toBe("2, 2");
    expect(wrapper.find("[test-data='player-score']").text()).toBe("50");
  });

  test("User can able to see scoreboard details", () => {
    const scoreboard = wrapper.find("[test-data='scoreboard']");

    // check scoreboard table col names
    const colName = scoreboard.findAll("th");
    expect(colName.at(0).text()).toBe("Name");
    expect(colName.at(1).text()).toBe("R-1");
    expect(colName.at(2).text()).toBe("R-2");
    expect(colName.at(3).text()).toBe("R-3");
    expect(colName.at(4).text()).toBe("Remaining Score");

    // check scoreboard score
    const scoreTable = scoreboard.find("tbody");

    // player name
    expect(scoreTable.findAll("th").at(0).text()).toBe("Jeel R.");
    const round = scoreboard.findAll("td");

    // round 1
    expect(round.at(0).findAll("li").at(0).text()).toBe("10");
    expect(round.at(0).findAll("li").at(1).text()).toBe("12");
    expect(round.at(0).findAll("li").at(2).text()).toBe("13");

    // round 2
    expect(round.at(1).findAll("li").at(0).text()).toBe("1");
    expect(round.at(1).findAll("li").at(1).text()).toBe("2");
    expect(round.at(1).findAll("li").at(2).text()).toBe("3");

    // round 3
    expect(round.at(2).findAll("li").at(0).text()).toBe("2");
    expect(round.at(2).findAll("li").at(1).text()).toBe("32");
    expect(round.at(2).findAll("li").at(2).text()).toBe("12");

    // round total
    expect(round.at(3).text()).toBe("87");

    // player name
    expect(scoreTable.findAll("th").at(1).text()).toBe("Vatsal C.");

    // round 1
    expect(round.at(4).findAll("li").at(0).text()).toBe("25");
    expect(round.at(4).findAll("li").at(1).text()).toBe("23");
    expect(round.at(4).findAll("li").at(2).text()).toBe("21");

    // round 2
    expect(round.at(5).findAll("li").at(0).text()).toBe("32");
    expect(round.at(5).findAll("li").at(1).text()).toBe("12");
    expect(round.at(5).findAll("li").at(2).text()).toBe("12");

    // round 3
    expect(round.at(6).findAll("li").at(0).text()).toBe("32");
    expect(round.at(6).findAll("li").at(1).text()).toBe("44");
    expect(round.at(6).findAll("li").at(2).text()).toBe("43");

    // round total
    expect(round.at(7).text()).toBe("244");
  });

  test("User can undo last entered score", () => {
    const undoButton = wrapper.find("[test-data='undo']");
    undoButton.trigger("click");
    expect(actions["game/undoScore"]).toHaveBeenCalled();

    expect(actions["game/getCurrentTurn"]).toHaveBeenCalled();
    expect(actions["game/getCurrentTurn"].mock.calls[0][1]).toEqual(1);
    wrapper.vm.$store.commit("game/GET_CURRENTTURN", currentTurn[1]);
    expect(wrapper.vm.turn).toBe(currentTurn[1]);

    expect(actions["game/getGamePlayerInfo"]).toHaveBeenCalled();
    expect(actions["game/getGamePlayerInfo"].mock.calls[0][1]).toEqual({
      gameId: 1,
      playerId: 4,
    });
    wrapper.vm.$store.commit("game/GET_GAMEPLAYERINFO", player[1]);
    expect(wrapper.vm.playerInfo).toBe(player[1]);

    expect(actions["game/getScoreboard"]).toHaveBeenCalled();
    expect(actions["game/getScoreboard"].mock.calls[0][1]).toEqual(1);
    expect(wrapper.vm.scoreboard).toBe(scoreboard);
  });

  test("User can add score", async () => {
    // see current player details before add score
    expect(wrapper.find("[test-data='player-name']").text()).toBe(
      "Jeel Rupapara"
    );
    expect(wrapper.find("[test-data='player-round-turn']").text()).toBe("2, 3");
    expect(wrapper.find("[test-data='player-score']").text()).toBe("60");

    // add score
    wrapper.vm.onAddScore({ score: 10 });

    // add score api called
    expect(actions["game/addScore"]).toHaveBeenCalled();
    expect(actions["game/addScore"].mock.calls[0][1]).toEqual({
      gameId: 1,
      playerId: 4,
      roundId: 2,
      score: { score: 10 },
      turnId: 3,
    });

    // current turn api called
    expect(actions["game/getCurrentTurn"]).toHaveBeenCalled();
    expect(actions["game/getCurrentTurn"].mock.lastCall[1]).toEqual(1);

    // change route values
    wrapper.vm.$route.params.roundid = 3;
    wrapper.vm.$route.params.playerid = 3;
    wrapper.vm.$route.params.turnid = 1;

    // change turn value (round, player, turn) from store
    wrapper.vm.$store.commit("game/GET_CURRENTTURN", currentTurn[2]);
    expect(wrapper.vm.turn).toEqual(currentTurn[2]);

    await wrapper.vm.$nextTick();

    // player info api called
    expect(actions["game/getGamePlayerInfo"]).toHaveBeenCalled();
    expect(actions["game/getGamePlayerInfo"].mock.lastCall[1]).toEqual({
      gameId: 1,
      playerId: 3,
    });
    // change state player value
    wrapper.vm.$store.commit("game/GET_GAMEPLAYERINFO", player[2]);
    expect(wrapper.vm.playerInfo).toEqual(player[2]);

    // scoreboard api called
    expect(actions["game/getScoreboard"]).toHaveBeenCalled();
    expect(actions["game/getScoreboard"].mock.calls[0][1]).toEqual(1);
    expect(wrapper.vm.scoreboard).toEqual(scoreboard);

    await wrapper.vm.$nextTick();

    // see current player details before and another score
    expect(wrapper.find("[test-data='player-name']").text()).toBe(
      "Vatsal Chauhan"
    );
    expect(wrapper.find("[test-data='player-round-turn']").text()).toBe("3, 1");
    expect(wrapper.find("[test-data='player-score']").text()).toBe("91");

    // add another score
    wrapper.vm.onAddScore({ score: 5 });

    // add score api called
    expect(actions["game/addScore"]).toHaveBeenCalled();
    expect(actions["game/addScore"].mock.lastCall[1]).toEqual({
      gameId: 1,
      playerId: 3,
      roundId: 3,
      turnId: 1,
      score: { score: 5 },
    });

    // current turn api called
    expect(actions["game/getCurrentTurn"]).toHaveBeenCalled();
    expect(actions["game/getCurrentTurn"].mock.lastCall[1]).toEqual(1);

    // change route values
    wrapper.vm.$route.params.roundid = 3;
    wrapper.vm.$route.params.playerid = 3;
    wrapper.vm.$route.params.turnid = 2;

    // change turn value (round, player, turn) from store
    await wrapper.vm.$store.commit("game/GET_CURRENTTURN", currentTurn[3]);
    expect(wrapper.vm.turn).toEqual(currentTurn[3]);

    await wrapper.vm.$nextTick();

    // player info api called
    expect(actions["game/getGamePlayerInfo"]).toHaveBeenCalled();
    expect(actions["game/getGamePlayerInfo"].mock.lastCall[1]).toEqual({
      gameId: 1,
      playerId: 3,
    });

    // change state player value
    wrapper.vm.$store.commit("game/GET_GAMEPLAYERINFO", player[3]);
    expect(wrapper.vm.playerInfo).toEqual(player[3]);

    // scoreboard api called
    expect(actions["game/getScoreboard"]).toHaveBeenCalled();
    expect(actions["game/getScoreboard"].mock.lastCall[1]).toEqual(1);
    expect(wrapper.vm.scoreboard).toEqual(scoreboard);

    await wrapper.vm.$nextTick();

    // see current player details
    expect(wrapper.find("[test-data='player-name']").text()).toBe(
      "Vatsal Chauhan"
    );
    expect(wrapper.find("[test-data='player-round-turn']").text()).toBe("3, 2");
    expect(wrapper.find("[test-data='player-score']").text()).toBe("96");
  });

  test("When game type is high score then score col name is total score else remaining score", () => {
    wrapper.vm.scoreColumnName();
    expect(wrapper.vm.scoreColName).toBe("Remaining Score");
    wrapper.vm.$store.commit("game/GET_GAMEPLAYERINFO", player[4]);
    wrapper.vm.scoreColumnName();
    expect(wrapper.vm.scoreColName).toBe("Total Score");
  });
});
