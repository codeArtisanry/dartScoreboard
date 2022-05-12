import jwtDecode from "jwt-decode";

export const state = () => {
  return {
    auth: "",
    token: "",
    games: "",
    currentGame: "",
    scoreboard: "",
  };
};
export const mutations = {
  setAuth(state, auth, token) {
    state.auth = auth;
    state.token = token;
  },
  currentgame(state, res) {
    state.currentGame = res;
  },
  getScoreboard(state, res) {
    state.scoreboard = res;
  },
  getGames(state, res) {
    state.games = res;
  },
};
export const actions = {
  nuxtServerInit({ commit }, { req, app }) {
    let auth = "";
    let token = "";
    if (req.headers.cookie) {
      auth = app.$cookies.get("user");
      try {
        token = jwtDecode(auth);
      } catch (err) {
        // this.$router.push('/signin')
        // No valid cookie found
      }
    }
    commit("setAuth", { auth, token });
  },
};

export const getters = {
  getWinnerName: (state) => {
    return state.scoreboard.winner;
  },
  getCurrentGame: (state) => {
    return state.currentGame;
  },
  getRound: (state) => {
    const x = state.scoreboard.players_score;
    if (x) {
      return state.scoreboard.players_score[0].rounds;
    }
  },
  getGamesData: (state) => {
    return state.games;
  },
  getPlayerScore: (state) => {
    return state.scoreboard.players_score;
  },
  getNameAndTotal: (state) => {
    return state.scoreboard.players_score;
  },
  getgameDetails: (state) => {
    const players = state.scoreboard.players_score;
    if (players) {
      for (let gameDetails = 0; gameDetails < players.length; gameDetails++) {
        const gamedata = players[gameDetails];
        console.log(gamedata);
      }
      return players;
    }
  },
};
