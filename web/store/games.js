const state = () => {
    return {
        games: "",
    };
};

const mutations = {
    GET_GAMES(state, games) {
        state.games = games;
    }
};

const actions = {
    async getGames({ commit }, page) {
        const gamesRes = await this.$axios.$get(`/api/v1/games?${page}`);
        commit("GET_GAMES", gamesRes);
    }
};

const getters = {
    list: (state) => {
        return state.games.list
    },
    next: (state) => {
        return state.games.next
    },
    previous: (state) => {
        return state.games.previous
    },
};

export default {
    state,
    getters,
    actions,
    mutations,
};
