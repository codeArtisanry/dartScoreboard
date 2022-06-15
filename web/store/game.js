const state = () => {
  return {
    game: "",
    playerInfo: "",
    scoreboard: "",
    currentTurn: "",
  };
};

const actions = {
  async deleteGame({ commit }, gameId) {
    await this.$axios.$delete(`/api/v1/games/${gameId}`);
  },
  async getGame({ commit }, gameId) {
    const gameRes = await this.$axios.$get(`/api/v1/games/${gameId}`);
    commit("GET_GAME", gameRes);
  },
  async createGame({ commit }, game) {
    const gameRes = await this.$axios.$post(`/api/v1/games`, game);
    commit("CREATE_GAME", gameRes);
  },
  async updateGame({ commit }, game) {
    const gameRes = await this.$axios.$put(
      `/api/v1/games/${game.id}`,
      game.update
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

  async undoScore({ commit }, params) {
    await this.$axios
      .$delete(
        `/api/v1/games/${params.gameId}/players/${params.playerId}/rounds/${params.roundId}/turns/${params.turnId}/undo-score`
      )
      .catch((error) => {
        if (error.response.status === 400) {
          alert("Already Undo This Score");
        } else if (error.response.status === 500) {
          alert("Not Enough Score");
        }
      });
  },
};

const mutations = {
  GET_GAME(state, game) {
    state.game = game;
  },
  CREATE_GAME(state, game) {
    state.game = game;
  },
  UPDATE_GAME(state, game) {
    state.game = game;
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

const getters = {
  details: (state) => {
    return state.game
  },
  isOwner: (state, _, rootState) => {
    if (state.game.creater_name === rootState.auth.token.name) {
      return true
    } else {
      return false
    }
  }
}
export default {
  state,
  mutations,
  actions,
  getters
};
