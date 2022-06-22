export default {
  ssr: true,
  target: "server",

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: "web",
    script: [
      {
        src: "https://unpkg.com/dartboard/dist/dartboard.js",
      },
    ],
    meta: [
      {
        charset: "utf-8",
      },
      {
        name: "viewport",
        content: "width=device-width, initial-scale=1",
      },
      {
        hid: "description",
        name: "description",
        content: "",
      },
      {
        name: "format-detection",
        content: "telephone=no",
      },
    ],
    link: [
      {
        rel: "icon",
        type: "image/x-icon",
        href: "/favicon.ico",
      },
    ],
  },
  publicRuntimeConfig: {
    logoutURL: process.env.LOGOUT_URL,
    loginURL: process.env.LOGIN_URL,
  },
  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/eslint
    "@nuxtjs/eslint-module",
    "@nuxtjs/fontawesome",
    "@nuxtjs/dotenv",
  ],
  fontawesome: {
    component: "fa",
    icons: {
      solid: true,
      brands: true,
    },
  },
  eslint: {
    /* module options */
    extensions: ["js", "vue"],
    exclude: ["node_modules"],
    fix: true,
  },

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/bootstrap
    "bootstrap-vue/nuxt",
    // https://go.nuxtjs.dev/axios
    "@nuxtjs/axios",
    // https://go.nuxtjs.dev/pwa
    "@nuxtjs/pwa",
    // https://go.nuxtjs.dev/content
    "@nuxt/content",
    // https://go.nuxtjs.dev/
    "cookie-universal-nuxt",
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    baseURL: `${process.env.BACKEND_URL}/`,
    proxy: true,
  },
  proxy: {
    "/api": `${process.env.BACKEND_URL}/`,
  },

  // PWA module configuration: https://go.nuxtjs.dev/pwa
  pwa: {
    manifest: {
      lang: "en",
    },
  },

  // Content module configuration: https://go.nuxtjs.dev/config-content
  content: {},

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},
  server: {
    host: "0",
  },
  // middleware configuration: https://go.nuxtjs.dev/config-middleware
  // router: {
  //   middleware: 'auth',
  // },
};
