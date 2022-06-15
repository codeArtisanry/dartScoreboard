const state = () => {
    return {
        users: "",
    };
};

const actions = {
    async getUsers({ commit }) {
        const usersRes = await this.$axios.$get(`/api/v1/users`);
        commit("GET_USERS", usersRes);
    },
};

const mutations = {
    GET_USERS(state, users) {
        state.users = users;
    },
};

const getters = {
    list: (state) => {
        return state.users.list
    }
}
export default {
    state,
    mutations,
    actions,
    getters
};
