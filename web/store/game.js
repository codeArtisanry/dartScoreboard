const state = () => {
  return {
    game: "",
    games: "",
    users: "",
    playerInfo: "",
    scoreboard: "",
    currentTurn: "",
  };
};

const actions = {
  async deleteGame({ commit }, gameId) {
    await this.$axios.$delete(`/api/v1/games/${gameId}`);
  },
  async getUsers({ commit }) {
    const usersRes = await this.$axios.$get(`/api/v1/users`);
    commit("GET_USERS", usersRes);
  },
  async getPrePage({ commit }, prePageLink) {
    const prePageRes = await this.$axios.$get(`${prePageLink}`);
    commit("GET_PREVIOUS_GAMES", prePageRes);
  },
  async getPostPage({ commit }, postPageLink) {
    const postPageRes = await this.$axios.$get(`${postPageLink}`);
    commit("GET_NEXT_GAMES", postPageRes);
  },
  async getGames({ commit }) {
    const gamesRes = await this.$axios.$get(`/api/v1/games?page=1`);
    commit("GET_GAMES", gamesRes);
  },
  async getGame({ commit }, gameId) {
    const gameRes = await this.$axios.$get(`/api/v1/games/${gameId}`);
    commit("GET_GAME", gameRes);
  },
  async postGame({ commit }, gameData) {
    const gameRes = await this.$axios.$post(`/api/v1/games`, gameData);
    commit("POST_GAME", gameRes);
  },
  async updateGame({ commit }, gameData) {
    const gameRes = await this.$axios.$put(
      `/api/v1/games/${gameData.gameId}`,
      gameData.update
    );
    commit("UPDATE_GAME", gameRes);
  },
  async getScoreboard({ commit }, gameId) {
    const scoreboardRes = await this.$axios.$get(
      `/api/v1/games/${gameId}/scoreboard`
    );
    commit("GET_SCOREBOARD", scoreboardRes);
  },
  async getGamePlayerInfo({ commit }, params) {
    const gamePlayerRes = await this.$axios.$get(
      `api/v1/games/${params.gameId}/players/${params.playerId}/player-info`
    );
    commit("GET_GAMEPLAYERINFO", gamePlayerRes);
  },
  async getCurrentTurn({ commit }, gameId) {
    const curruntTurnRes = await this.$axios.$get(
      `api/v1/games/${gameId}/active-status`
    );
    commit("GET_CURRENTTURN", curruntTurnRes);
  },
  async postScore({ commit }, params) {
    await this.$axios
      .$post(
        `/api/v1/games/${params.gameId}/players/${params.playerId}/rounds/${params.roundId}/turns/${params.turnId}/score`,
        params.score
      )
      .catch((error) => {
        if (error.response.status === 400) {
          alert("Already Entered Score For This Dart");
        }
      });
  },
};

const mutations = {
  GET_GAME(state, game) {
    state.game = game;
  },
  POST_GAME(state, game) {
    state.game = game;
  },
  GET_GAMES(state, games) {
    state.games = games;
  },
  UPDATE_GAME(state, game) {
    state.game = game;
  },
  GET_USERS(state, users) {
    state.users = users;
  },
  GET_NEXT_GAMES(state, games) {
    state.games = games;
  },
  GET_PREVIOUS_GAMES(state, games) {
    state.games = games;
  },
  GET_SCOREBOARD(state, scoreboard) {
    state.scoreboard = scoreboard;
  },
  GET_GAMEPLAYERINFO(state, playerInfo) {
    state.playerInfo = playerInfo;
  },
  GET_CURRENTTURN(state, currentTurn) {
    state.currentTurn = currentTurn;
  },
};

export default {
  state,
  mutations,
  actions,
};
