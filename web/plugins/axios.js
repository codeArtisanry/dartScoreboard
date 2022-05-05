export default function ({ $axios, redirect }) {
    $axios.onError((error) => {
      if (
        error.response.status_code === 400 ||
        error.response.status_code === 500
      ) {
        console.log(error.response.msg)
        redirect('/signin')
      }
    })
  }
