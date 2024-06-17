<script>
  // @ts-nocheck
  export let users;
  export let title;
  import axios from "axios";
  import Add from "../Vote/Add.svelte";
  let data = "";

  async function GetData() {
    if (title) {
      try {
        const response = await axios.post(`http://localhost:3000/api/anime`, {
          name: title,
        });

        if (response.data.status === "OK") {
          data = response.data.data;
          console.log(response.data);
        }
        if (response.data.status === 404) {
          setTimeout(function () {
            window.location.href = "/404";
          }, 1);
        }
      } catch (err) {
        console.error(err);
      }
    } else {
    }
  }
  GetData();
</script>

<main class="bg-white rounded-lg shadow m-4 dark:bg-gray-800">
  <div class="w-full mx-auto max-w-screen-xl p-4">
    <span class="text-5xl text-gray-100 sm:text-center dark:text-gray-100"
      >{title}</span
    >
    <span
      class="block py-2 px-3 text-gray-900 rounded md:hover:bg-transparent md:border-0 md:p-0 dark:text-white md:dark:hover:bg-transparent"
      >Açıklama kısmı : {data.description}</span
    >
    <span
      class="block py-2 px-3 text-gray-900 rounded md:hover:bg-transparent md:border-0 md:p-0 dark:text-white md:dark:hover:bg-transparent"
      >Küçük resim : {data.thumbnail}</span
    >
    <span
      class="block py-2 px-3 text-gray-900 rounded md:hover:bg-transparent md:border-0 md:p-0 dark:text-white md:dark:hover:bg-transparent"
      >🌟: {data.rating}</span
    >
    {#if users}
      <span
        class="block py-2 px-3 text-gray-900 rounded md:hover:bg-transparent md:border-0 md:p-0 dark:text-white md:dark:hover:bg-transparent"
        ><Add anime={data.name} /></span
      >
    {/if}
  </div>
</main>
